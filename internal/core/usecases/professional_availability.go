package usecases

import (
	"context"

	"github.com/hebertzin/scheduler/internal/domain"
	"github.com/hebertzin/scheduler/internal/domain/ports/inbound"
	"github.com/hebertzin/scheduler/internal/domain/ports/outbound"
	"github.com/hebertzin/scheduler/internal/errors"
	"github.com/sirupsen/logrus"
)

type ProfessionalsAvailabilityManager struct {
	repository outbound.ProfessionalsAvailabilityRepository
	logger     *logrus.Logger
}

func NewProfessionalsAvailability(repository outbound.ProfessionalsAvailabilityRepository, logger *logrus.Logger) inbound.ProfessionalsAvailabilityUseCase {
	return &ProfessionalsAvailabilityManager{repository: repository, logger: logger}
}

func (manager *ProfessionalsAvailabilityManager) Add(ctx context.Context, availability *domain.ProfessionalAvailability) (*domain.ProfessionalAvailability, *errors.Exception) {
	availability, err := manager.repository.Add(ctx, availability)
	if err != nil {
		manager.logger.Error("Error creating professional availability.", "professional_availability_use_case", "err", err.Error())

		return nil, errors.Unexpected(errors.WithMessage("error creating availability"), errors.WithError(err))
	}

	manager.logger.Info("Professional availability created successfully.", "professional_availability_use_case")

	return availability, nil
}

func (manager *ProfessionalsAvailabilityManager) GetProfessionalAvailabilityById(ctx context.Context, id string) ([]domain.ProfessionalAvailability, *errors.Exception) {
	availability, err := manager.repository.GetProfessionalAvailabilityById(ctx, id)
	if err != nil {
		manager.logger.Error("Error getting professional availability.", "professional_availability_use_case", "err", err.Error())

		return nil, errors.Unexpected(errors.WithMessage("error get professional availability"), errors.WithError(err))
	}

	manager.logger.Info("Professional availability retrieved successfully.", "professional_availability_use_case")

	return availability, nil
}
