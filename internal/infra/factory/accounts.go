package factory

import (
	"github.com/hebertzin/scheduler/internal/core/usecases"
	"github.com/hebertzin/scheduler/internal/domain/ports/outbound"
	"github.com/hebertzin/scheduler/internal/infra/db/repository"
	"github.com/hebertzin/scheduler/internal/presentation/controllers"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func AccountFactory(db *gorm.DB, logger *logrus.Logger) outbound.AccountController {
	repo := repository.NewAccountsRepository(db, logger)
	accountManager := usecases.NewAccount(repo, nil, logger)

	return controllers.NewAccount(accountManager)
}
