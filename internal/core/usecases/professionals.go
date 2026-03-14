package usecases

import (
	"context"

	"github.com/hebertzin/scheduler/internal/core"
	"github.com/hebertzin/scheduler/internal/domain"
	"github.com/sirupsen/logrus"
)

type ProfessionalsManager struct {
	repository domain.ProfessionalsRepository
	logger     *logrus.Logger
}

func NewProfissional(repository domain.ProfessionalsRepository, logger *logrus.Logger) domain.ProfessionalsUseCase {
	return &ProfessionalsManager{repository: repository, logger: logger}
}

func (s *ProfessionalsManager) FindProfessionalById(ctx context.Context, professionalId string) (*domain.Professionals, *core.Exception) {
	professional, err := s.repository.FindProfessionalById(ctx, professionalId)
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

func (s *ProfessionalsManager) UpdateProfessionalById(ctx context.Context, professionalId string, professionalData *domain.Professionals) (*domain.Professionals, *core.Exception) {
	professional, err := s.repository.UpdateProfessionalById(ctx, professionalId, professionalData)
	if err != nil {
		return nil, core.Unexpected()
	}

	return professional, nil
}
