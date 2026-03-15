package factory

import (
	"github.com/hebertzin/scheduler/internal/core/usecases"

	"github.com/hebertzin/scheduler/internal/domain/ports/outbound"

	"github.com/hebertzin/scheduler/internal/infra/db/repository"
	"github.com/hebertzin/scheduler/internal/presentation/controllers"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func AppointmentFactory(db *gorm.DB, logger *logrus.Logger, publisher outbound.Publisher) outbound.AppointmentController {
	repo := repository.NewAppointmentRepository(db, logger)
	availabilityManager := usecases.NewAvailability(repo)
	appointmentManager := usecases.NewAppointment(repo, availabilityManager, publisher, logger)

	return controllers.NewAppointment(appointmentManager)

}
