package inbound

import (
	"context"

	"github.com/hebertzin/scheduler/internal/domain"
	"github.com/hebertzin/scheduler/internal/errors"
)

type (
	ServicesUseCase interface {
		Add(ctx context.Context, payload *domain.Services) (*domain.Services, *errors.Exception)
		FindServiceById(ctx context.Context, id string) (*domain.Services, *errors.Exception)
		GetAllServicesByProfessionalId(ctx context.Context, professionalId string) ([]domain.Services, *errors.Exception)
	}
)
