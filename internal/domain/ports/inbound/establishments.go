package inbound

import (
	"context"

	"github.com/hebertzin/scheduler/internal/domain"
	"github.com/hebertzin/scheduler/internal/errors"
)

type (
	EstablishmentUseCase interface {
		Add(ctx context.Context, payload *domain.Establishment) (*domain.Establishment, *errors.Exception)
		GetAllProfessionalsByEstablishmentId(ctx context.Context, id string) ([]domain.Professionals, *errors.Exception)
		FindEstablishmentById(ctx context.Context, id string) (*domain.Establishment, *errors.Exception)
		GetEstablishmentReport(ctx context.Context, id string) (*domain.EstablishmentMetrics, *errors.Exception)
		UpdateEstablishmentById(ctx context.Context, id string, payload *domain.Establishment) (*domain.Establishment, *errors.Exception)
	}
)
