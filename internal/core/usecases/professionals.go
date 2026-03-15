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

func (manager *ProfessionalsManager) FindProfessionalById(ctx context.Context, id string) (*domain.Professionals, *core.Exception) {
	professional, err := manager.repository.FindProfessionalById(ctx, id)
	if err != nil {
		manager.logger.Error("Error finding professional by id.", "professionals_use_case", "err", err.Error())

		return nil, core.Unexpected(core.WithMessage("error finding professional"), core.WithError(err))
	}

	manager.logger.Info("Professional found successfully.", "professionals_use_case")

	return professional, nil
}

func (manager *ProfessionalsManager) Add(ctx context.Context, payload *domain.Professionals) (*domain.Professionals, *core.Exception) {
	if payload.Name == "" || payload.Role == "" || payload.EstablishmentId == "" {
		manager.logger.Error("Missing required fields to create professional.", "professionals_use_case")

		return nil, core.BadRequest(core.WithMessage("some fields are missing"))
	}

	professional, err := manager.repository.Add(ctx, payload)
	if err != nil {
		manager.logger.Error("Error creating professional.", "professionals_use_case", "err", err.Error())

		return nil, core.Unexpected()
	}

	manager.logger.Info("Professional created successfully.", "professionals_use_case")

	return professional, nil
}

func (manager *ProfessionalsManager) UpdateProfessionalById(ctx context.Context, id string, professionalData *domain.Professionals) (*domain.Professionals, *core.Exception) {
	professional, err := manager.repository.UpdateProfessionalById(ctx, id, professionalData)
	if err != nil {
		manager.logger.Error("Error updating professional.", "professionals_use_case", "err", err.Error())

		return nil, core.Unexpected()
	}

	manager.logger.Info("Professional updated successfully.", "professionals_use_case")

	return professional, nil
}
