package factory

import (
	"github.com/hebertzin/scheduler/internal/core/usecases"
	"github.com/hebertzin/scheduler/internal/domain"
	"github.com/hebertzin/scheduler/internal/infra/db/repository"
	"github.com/hebertzin/scheduler/internal/presentation/controllers"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func ProfessionalAvailabilityFactory(db *gorm.DB, logger *logrus.Logger) domain.ProfessionalAvailabilityController {
	repo := repository.NewProfessionalsAvailabilityRepository(db, logger)
	profissionalAvailabilityManager := usecases.NewProfessionalsAvailability(repo, logger)

	return controllers.NewProfessionalAvailability(profissionalAvailabilityManager)
}
