package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hebertzin/scheduler/internal/infra/factory"
	"github.com/hebertzin/scheduler/internal/presentation/middlewares"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func AppointmentGroupRouter(router *gin.Engine, db *gorm.DB, logger *logrus.Logger) {
	appointmentFactory := factory.AppointmentFactory(db, logger)
	v1 := router.Group("/api/v1")
	{
		v1.GET("/appointments/:id/professional", middlewares.ValidateParamRequest(), appointmentFactory.GetAllAppointmentsByProfessionalId)
		v1.GET("/appointments/:id", middlewares.ValidateParamRequest(), appointmentFactory.GetAppointmentById)
		v1.POST("/appointments", appointmentFactory.Add)
	}
}
