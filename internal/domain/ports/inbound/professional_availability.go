package inbound

import (
	"context"

	"github.com/hebertzin/scheduler/internal/domain"
	"github.com/hebertzin/scheduler/internal/errors"
)

type (
	ProfessionalsAvailabilityUseCase interface {
		Add(ctx context.Context, availability *domain.ProfessionalAvailability) (*domain.ProfessionalAvailability, *errors.Exception)
		GetProfessionalAvailabilityById(ctx context.Context, id string) ([]domain.ProfessionalAvailability, *errors.Exception)
	}
)
