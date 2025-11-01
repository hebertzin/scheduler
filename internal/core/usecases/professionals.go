package usecases

import (
	"context"

	"github.com/hebertzin/scheduler/internal/core"
	"github.com/hebertzin/scheduler/internal/domain"
	"github.com/sirupsen/logrus"
)

type ProfessionalsUseCase struct {
	repository domain.ProfessionalsRepository
	logger     *logrus.Logger
}

func NewProfissionalUseCase(repository domain.ProfessionalsRepository, logger *logrus.Logger) domain.ProfessionalsUseCase {
	return &ProfessionalsUseCase{repository: repository, logger: logger}
}

func (s *ProfessionalsUseCase) FindProfessionalById(ctx context.Context, professionalId string) (*domain.Professionals, *core.Exception) {
	professional, err := s.repository.FindProfessionalById(ctx, professionalId)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("error finding professional"), core.WithError(err))
	}

	return professional, nil
}

func (s *ProfessionalsUseCase) Add(ctx context.Context, payload *domain.Professionals) (*domain.Professionals, *core.Exception) {
	if payload.Name == "" || payload.Role == "" || payload.EstablishmentId == "" {
		return nil, core.BadRequest(core.WithMessage("some fields are missing"))
	}

	professional, err := s.repository.Add(ctx, payload)
	if err != nil {
		return nil, core.Unexpected()
	}

	return professional, nil
}

func (s *ProfessionalsUseCase) UpdateProfessionalById(ctx context.Context, professionalId string, professionalData *domain.Professionals) (*domain.Professionals, *core.Exception) {
	professional, err := s.repository.UpdateProfessionalById(ctx, professionalId, professionalData)
	if err != nil {
		return nil, core.Unexpected()
	}

	return professional, nil
}
