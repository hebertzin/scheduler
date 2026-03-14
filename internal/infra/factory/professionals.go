package factory

import (
	"github.com/hebertzin/scheduler/internal/core/usecases"
	"github.com/hebertzin/scheduler/internal/domain/ports/outbound"
	"github.com/hebertzin/scheduler/internal/infra/db/repository"
	"github.com/hebertzin/scheduler/internal/presentation/controllers"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func ProfessionalFactory(db *gorm.DB, logger *logrus.Logger) outbound.ProfessionalsController {
	repo := repository.NewProfessionalsRepository(db, logger)
	professionalManager := usecases.NewProfessional(repo, logger)

	return controllers.NewProfessional(professionalManager)
}
