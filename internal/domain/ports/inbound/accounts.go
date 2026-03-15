package inbound

import (
	"context"

	"github.com/hebertzin/scheduler/internal/domain"
	"github.com/hebertzin/scheduler/internal/errors"
)

type (
	AccountUseCase interface {
		Add(ctx context.Context, payload *domain.Account) (*domain.Account, *errors.Exception)
		FindAccountById(ctx context.Context, id string) (*domain.Account, *errors.Exception)
		FindAllAccounts(ctx context.Context) ([]domain.Account, *errors.Exception)
		FindAllEstablishmentsByAccountId(ctx context.Context, accountId string) ([]domain.Establishment, *errors.Exception)
	}
)
