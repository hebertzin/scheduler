package factory

import (
	"github.com/hebertzin/scheduler/internal/core/usecases"
	"github.com/hebertzin/scheduler/internal/domain"
	"github.com/hebertzin/scheduler/internal/infra/db/repository"
	"github.com/hebertzin/scheduler/internal/presentation/controllers"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func ProfessionalFactory(db *gorm.DB, logger *logrus.Logger) domain.ProfessionalsController {
	repo := repository.NewProfessionalsRepository(db, logger)
	profissionalManager := usecases.NewProfissional(repo, logger)

	return controllers.NewProfessional(profissionalManager)
}
