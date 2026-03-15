package inbound

import (
	"context"

	"github.com/hebertzin/scheduler/internal/domain"
	"github.com/hebertzin/scheduler/internal/errors"
)

type (
	ProfessionalsUseCase interface {
		Add(ctx context.Context, payload *domain.Professionals) (*domain.Professionals, *errors.Exception)
		FindProfessionalById(ctx context.Context, id string) (*domain.Professionals, *errors.Exception)
		UpdateProfessionalById(ctx context.Context, professional_id string, professionalData *domain.Professionals) (*domain.Professionals, *errors.Exception)
	}
)
