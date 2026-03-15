package factory

import (
	"github.com/hebertzin/scheduler/internal/core/usecases"
	"github.com/hebertzin/scheduler/internal/domain/ports/outbound"
	"github.com/hebertzin/scheduler/internal/infra/db/repository"
	"github.com/hebertzin/scheduler/internal/infra/security"
	"github.com/hebertzin/scheduler/internal/presentation/controllers"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func AccountFactory(db *gorm.DB, logger *logrus.Logger) outbound.AccountController {
	repo := repository.NewAccountsRepository(db, logger)
	bcryptHasher := security.NewBcryptHasher(bcrypt.DefaultCost)
	accountManager := usecases.NewAccount(repo, nil, bcryptHasher, logger)

	return controllers.NewAccount(accountManager)
}
