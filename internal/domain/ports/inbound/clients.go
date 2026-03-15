package inbound

import (
	"context"

	"github.com/hebertzin/scheduler/internal/domain"
	"github.com/hebertzin/scheduler/internal/errors"
)

type (
	ClientUseCase interface {
		Add(ctx context.Context, payload *domain.Client) (*domain.Client, *errors.Exception)
	}
)
