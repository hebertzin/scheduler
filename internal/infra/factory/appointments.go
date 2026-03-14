package factory

import (
	"github.com/hebertzin/scheduler/internal/core/usecases"
	"github.com/hebertzin/scheduler/internal/domain"

	"github.com/hebertzin/scheduler/internal/infra/db/repository"
	"github.com/hebertzin/scheduler/internal/presentation/controllers"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func AppointmentFactory(db *gorm.DB, logger *logrus.Logger) domain.AppointmentController {
	repo := repository.NewAppointmentRepository(db, logger)
	appointmentManager := usecases.NewAppointment(repo, logger, nil)

	return controllers.NewAppointment(appointmentManager)

}
