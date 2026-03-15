package inbound

import (
	"context"

	"github.com/hebertzin/scheduler/internal/core"
	"github.com/hebertzin/scheduler/internal/domain"
)

type (
	ProfessionalsAvailabilityUseCase interface {
		Add(ctx context.Context, availability *domain.ProfessionalAvailability) (*domain.ProfessionalAvailability, *core.Exception)
		GetProfessionalAvailabilityById(ctx context.Context, id string) ([]domain.ProfessionalAvailability, *core.Exception)
	}
)
