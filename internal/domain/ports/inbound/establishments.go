package inbound

import (
	"context"

	"github.com/hebertzin/scheduler/internal/core"
	"github.com/hebertzin/scheduler/internal/domain"
)

type (
	EstablishmentUseCase interface {
		Add(ctx context.Context, payload *domain.Establishment) (*domain.Establishment, *core.Exception)
		GetAllProfessionalsByEstablishmentId(ctx context.Context, id string) ([]domain.Professionals, *core.Exception)
		FindEstablishmentById(ctx context.Context, id string) (*domain.Establishment, *core.Exception)
		GetEstablishmentReport(ctx context.Context, id string) (*domain.EstablishmentMetrics, *core.Exception)
		UpdateEstablishmentById(ctx context.Context, id string, payload *domain.Establishment) (*domain.Establishment, *core.Exception)
	}
)
