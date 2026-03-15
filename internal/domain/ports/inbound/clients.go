package inbound

import (
	"context"

	"github.com/hebertzin/scheduler/internal/core"
	"github.com/hebertzin/scheduler/internal/domain"
)

type (
	ClientUseCase interface {
		Add(ctx context.Context, payload *domain.Client) (*domain.Client, *core.Exception)
	}
)
