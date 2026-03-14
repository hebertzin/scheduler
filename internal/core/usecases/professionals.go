package usecases

import (
	"context"

	"github.com/hebertzin/scheduler/internal/core"
	"github.com/hebertzin/scheduler/internal/domain"
	"github.com/hebertzin/scheduler/internal/domain/ports/inbound"
	"github.com/hebertzin/scheduler/internal/domain/ports/outbound"
	"github.com/sirupsen/logrus"
)

type ProfessionalsManager struct {
	repository outbound.ProfessionalsRepository
	logger     *logrus.Logger
}

func NewProfessional(repository outbound.ProfessionalsRepository, logger *logrus.Logger) inbound.ProfessionalsUseCase {
	return &ProfessionalsManager{repository: repository, logger: logger}
}

func (s *ProfessionalsManager) FindProfessionalById(ctx context.Context, id string) (*domain.Professionals, *core.Exception) {
	professional, err := s.repository.FindProfessionalById(ctx, id)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("error finding professional"), core.WithError(err))
	}

	return professional, nil
}

func (s *ProfessionalsManager) Add(ctx context.Context, payload *domain.Professionals) (*domain.Professionals, *core.Exception) {
	if payload.Name == "" || payload.Role == "" || payload.EstablishmentId == "" {
		return nil, core.BadRequest(core.WithMessage("some fields are missing"))
	}

	professional, err := s.repository.Add(ctx, payload)
	if err != nil {
		return nil, core.Unexpected()
	}

	return professional, nil
}

func (s *ProfessionalsManager) UpdateProfessionalById(ctx context.Context, id string, professionalData *domain.Professionals) (*domain.Professionals, *core.Exception) {
	professional, err := s.repository.UpdateProfessionalById(ctx, id, professionalData)
	if err != nil {
		return nil, core.Unexpected()
	}

	return professional, nil
}
