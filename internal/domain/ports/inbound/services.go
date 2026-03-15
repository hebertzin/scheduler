package inbound

import (
	"context"

	"github.com/hebertzin/scheduler/internal/core"
	"github.com/hebertzin/scheduler/internal/domain"
)

type (
	ServicesUseCase interface {
		Add(ctx context.Context, payload *domain.Services) (*domain.Services, *core.Exception)
		FindServiceById(ctx context.Context, id string) (*domain.Services, *core.Exception)
		GetAllServicesByProfessionalId(ctx context.Context, professionalId string) ([]domain.Services, *core.Exception)
	}
)
