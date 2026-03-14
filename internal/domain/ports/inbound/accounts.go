package inbound

import (
	"context"

	"github.com/hebertzin/scheduler/internal/core"
	"github.com/hebertzin/scheduler/internal/domain"
)

type (
	AccountUseCase interface {
		Add(ctx context.Context, payload *domain.Account) (*domain.Account, *core.Exception)
		FindAccountById(ctx context.Context, id string) (*domain.Account, *core.Exception)
		FindAllAccounts(ctx context.Context) ([]domain.Account, *core.Exception)
		FindAllEstablishmentsByAccountId(ctx context.Context, accountId string) ([]domain.Establishment, *core.Exception)
	}
)
