package inbound

import (
	"context"

	"github.com/hebertzin/scheduler/internal/core"
	"github.com/hebertzin/scheduler/internal/domain"
)

type (
	ProfessionalsUseCase interface {
		Add(ctx context.Context, payload *domain.Professionals) (*domain.Professionals, *core.Exception)
		FindProfessionalById(ctx context.Context, id string) (*domain.Professionals, *core.Exception)
		UpdateProfessionalById(ctx context.Context, professional_id string, professionalData *domain.Professionals) (*domain.Professionals, *core.Exception)
	}
)
