package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hebertzin/scheduler/internal/infra/factory"
	"github.com/hebertzin/scheduler/internal/presentation/middlewares"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func ProfessionalAvailabilityGroupRouter(router *gin.Engine, db *gorm.DB, logger *logrus.Logger) {
	professionalsAvailabilityFactory := factory.ProfessionalAvailabilityFactory(db, logger)
	v1 := router.Group("/api/v1")
	{
		v1.GET("/availability/:id/professional", middlewares.ValidateParamRequest(), professionalsAvailabilityFactory.GetProfessionalAvailabilityById)
		v1.POST("/availability", professionalsAvailabilityFactory.Add)
	}
}
