package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hebertzin/scheduler/internal/domain"
	"github.com/hebertzin/scheduler/internal/domain/eventconstants"
	"github.com/hebertzin/scheduler/internal/infra/broker"
	"github.com/hebertzin/scheduler/internal/infra/config/env"
	"github.com/hebertzin/scheduler/internal/infra/config/logging"
	"github.com/hebertzin/scheduler/internal/infra/db"
	"github.com/hebertzin/scheduler/internal/infra/factory"
	"github.com/hebertzin/scheduler/internal/presentation/middlewares"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

// @title Scheduler app
// @version 1.0
// @description Saas plataform

// @contact.name Hebert Santos
// @contact.url https://www.hebertzin.com/
// @contact.email hebertsantosdeveloper@gmail.com

// @BasePath /api/v1
func main() {
	config, _ := env.LoadConfiguration("/configs/config.json")

	logger := configureLogging(config)

	database := db.ConnectDatabase(config)

	r := createRouter()

	b := broker.NewRabbitMQ("amqp://guest:guest@localhost:5672")
	err := b.Connect()
	if err != nil {
		logger.Println("error connecting to the broker")
	}
	defer b.Close()

	configureSwagger(config, r)

	configureMigration(config, database)

	cofigureMetrics(config, r)

	startAppointmentAPI(r, database, logger, b.Channel)

	startAccountAPI(r, database, logger, b.Channel)

	startEstablishmentAPI(r, database, logger)

	startProfessionalsAPI(r, database, logger)

	startServicesAPI(r, database, logger)

	startProfessionalAvailabilityAPI(r, database, logger)

	srv := http.Server{
		Addr:    config.Port,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Println("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Println("Server Shutdown:", err)
	}

	<-ctx.Done()
	logger.Println("Server exiting")
}

func createRouter() *gin.Engine {
	return gin.Default()
}

func startAccountAPI(router *gin.Engine, db *gorm.DB, logger *logrus.Logger, channel *amqp091.Channel) {
	accountPublisherConfig := broker.PublishingConfig{
		RoutingKey:   eventconstants.AccountRoutingKey,
		ExchangeName: eventconstants.AccountExchangeName,
	}

	accountPublisher := broker.NewPublisher(channel, accountPublisherConfig)

	accountFactory := factory.AccountFactory(db, logger, accountPublisher)

	v1 := router.Group("/api/v1")
	{
		v1.GET("/accounts", accountFactory.FindAllAccounts)
		v1.POST("/accounts", accountFactory.Add)
		v1.GET("/accounts/:id", accountFactory.FindAccountById)
		v1.GET("/accounts/:id/establishments", accountFactory.FindAllEstablishmentsByAccountId)
	}

}

func startAppointmentAPI(router *gin.Engine, db *gorm.DB, logger *logrus.Logger, channel *amqp091.Channel) {
	appointmentPublisherConfig := broker.PublishingConfig{
		RoutingKey:   eventconstants.AppointmentRoutingKey,
		ExchangeName: eventconstants.AppointmentExchange,
	}

	appointmentPublisher := broker.NewPublisher(channel, appointmentPublisherConfig)

	appointmentFactory := factory.AppointmentFactory(db, logger, appointmentPublisher)
	v1 := router.Group("/api/v1")
	{
		v1.POST("/appointments", appointmentFactory.Add)
		v1.GET("/appointments/:id/professional", appointmentFactory.GetAllAppointmentsByProfessionalId)
		v1.GET("/appointments/:id", appointmentFactory.GetAppointmentById)
	}

}

func startEstablishmentAPI(router *gin.Engine, db *gorm.DB, logger *logrus.Logger) {
	establishmentFactory := factory.EstablishmentFactory(db, logger)
	v1 := router.Group("/api/v1")
	{
		v1.POST("/establishments", establishmentFactory.Add)
		v1.GET("/establishments/:id", establishmentFactory.FindEstablishmentById)
		v1.GET("/establishments/:id/professionals", establishmentFactory.GetAllProfessinalsByEstablishmentId)
		v1.GET("/establishments/:id/report", establishmentFactory.GetEstablishmentReport)
		v1.PUT("/establishments/:id/update", establishmentFactory.UpdateEstablishmentById)
	}

}

func startProfessionalAvailabilityAPI(router *gin.Engine, db *gorm.DB, logger *logrus.Logger) {
	professionalsAvailabilityFactory := factory.ProfessionalAvailabilityFactory(db, logger)
	v1 := router.Group("/api/v1")
	{

		v1.GET("/availability/:id/professional", middlewares.ValidateParamRequest(), professionalsAvailabilityFactory.GetProfessionalAvailabilityById)
		v1.POST("/availability", professionalsAvailabilityFactory.Add)
	}

}

func startProfessionalsAPI(router *gin.Engine, db *gorm.DB, logger *logrus.Logger) {
	professionalFactory := factory.ProfessionalFactory(db, logger)
	v1 := router.Group("/api/v1")
	{
		v1.POST("/professionals", professionalFactory.Add)
		v1.GET("/professionals/:id", professionalFactory.FindProfessionalById)
		v1.PUT("/professionals/:id", professionalFactory.UpdateProfessionalById)
	}
}

func startServicesAPI(router *gin.Engine, db *gorm.DB, logger *logrus.Logger) {
	servicesFactory := factory.ServiceFactory(db, logger)
	v1 := router.Group("/api/v1")
	{
		v1.POST("/services", servicesFactory.Add)
		v1.GET("/services/:id", servicesFactory.FindServiceById)
		v1.GET("/services/:id/all", servicesFactory.GetAllServicesByProfessionalId)
	}

}

func configureMigration(config *domain.ServiceConfig, database *gorm.DB) {
	if config.RunMigrationEnabled {
		if err := db.Migrate(database); err != nil {
			panic("Database migration failed: " + err.Error())
		}
	}
}

func configureSwagger(config *domain.ServiceConfig, router *gin.Engine) {
	if config.SwaggerEnabled {
		router.GET("/api/v1/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}

func grafanaHandler(c *gin.Context) {
	var requests *prometheus.CounterVec
	requests.WithLabelValues(c.Request.Method, c.FullPath()).Inc()
}

func cofigureMetrics(config *domain.ServiceConfig, router *gin.Engine) {
	if config.GrafanaEnabled {
		requests := prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "http_requests_total",
				Help: "Total number of HTTP requests",
			},
			[]string{"method", "endpoint"},
		)
		prometheus.MustRegister(requests)

		router.GET("/grafana", grafanaHandler)
		router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	}
}

func configureLogging(config *domain.ServiceConfig) *logrus.Logger {
	if config.LoggingEnabled {
		return logging.InitLogger()
	}
	return nil
}
