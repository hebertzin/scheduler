package domain

import (
	"context"
	"database/sql"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hebertzin/scheduler/internal/core"
)

type (
	Account struct {
		Name        string       `json:"name"`
		Email       string       `json:"email"`
		Password    string       `json:"password"`
		ActivatedAt sql.NullTime `json:"activatedAt"`
		CreatedAt   time.Time    `json:"createdAt"`
		UpdatedAt   time.Time    `json:"updatedAt"`
	}

	AccountUseCase interface {
		Add(ctx context.Context, payload *Account) (*Account, *core.Exception)
		FindAccountById(ctx context.Context, id string) (*Account, *core.Exception)
		FindAllAccounts(ctx context.Context) ([]Account, *core.Exception)
		FindAllEstablishmentsByAccountId(ctx context.Context, accountId string) ([]Establishment, *core.Exception)
	}

	AccountRepository interface {
		Add(ctx context.Context, account *Account) (*Account, error)
		FindAccountByEmail(ctx context.Context, email string) (*Account, error)
		FindAccountById(ctx context.Context, id string) (*Account, error)
		FindAllEstablishmentsByAccountId(ctx context.Context, id string) ([]Establishment, error)
		FindAllAccounts(ctx context.Context) ([]Account, error)
	}

	AccountController interface {
		Add(ctx *gin.Context)
		FindAccountById(ctx *gin.Context)
		FindAllAccounts(ctx *gin.Context)
		FindAllEstablishmentsByAccountId(ctx *gin.Context)
	}
)
