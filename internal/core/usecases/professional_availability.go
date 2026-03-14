package usecases

import (
	"context"

	"github.com/hebertzin/scheduler/internal/core"
	"github.com/hebertzin/scheduler/internal/domain"
	"github.com/hebertzin/scheduler/internal/domain/ports/inbound"
	"github.com/hebertzin/scheduler/internal/domain/ports/outbound"
	"github.com/sirupsen/logrus"
)

type ProfessionalsAvailabilityManager struct {
	repository outbound.ProfessionalsAvailabilityRepository
	logger     *logrus.Logger
}

func NewProfessionalsAvailability(repository outbound.ProfessionalsAvailabilityRepository, logger *logrus.Logger) inbound.ProfessionalsAvailabilityUseCase {
	return &ProfessionalsAvailabilityManager{repository: repository, logger: logger}
}

func (s *ProfessionalsAvailabilityManager) Add(ctx context.Context, availability *domain.ProfessionalAvailability) (*domain.ProfessionalAvailability, *core.Exception) {
	availability, err := s.repository.Add(ctx, availability)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("error creating availability"), core.WithError(err))
	}

	return availability, nil
}

func (s *ProfessionalsAvailabilityManager) GetProfessionalAvailabilityById(ctx context.Context, id string) ([]domain.ProfessionalAvailability, *core.Exception) {
	availability, err := s.repository.GetProfessionalAvailabilityById(ctx, id)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("error get professional availability"), core.WithError(err))
	}

	return availability, nil
}
