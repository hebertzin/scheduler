package outbound

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/hebertzin/scheduler/internal/domain"
)

type (
	AccountRepository interface {
		Add(ctx context.Context, account *domain.Account) (*domain.Account, error)
		FindAccountByEmail(ctx context.Context, email string) (*domain.Account, error)
		FindAccountById(ctx context.Context, id string) (*domain.Account, error)
		FindAllEstablishmentsByAccountId(ctx context.Context, id string) ([]domain.Establishment, error)
		FindAllAccounts(ctx context.Context) ([]domain.Account, error)
	}

	AccountController interface {
		Add(ctx *gin.Context)
		FindAccountById(ctx *gin.Context)
		FindAllAccounts(ctx *gin.Context)
	}
)
