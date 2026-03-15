package usecases

import (
	"context"

	"github.com/hebertzin/scheduler/internal/domain"
	"github.com/hebertzin/scheduler/internal/domain/ports/inbound"
	"github.com/hebertzin/scheduler/internal/domain/ports/outbound"
	"github.com/hebertzin/scheduler/internal/errors"
	"github.com/sirupsen/logrus"
)

type EstablishmentManager struct {
	repository outbound.EstablishmentRepository
	logger     *logrus.Logger
}

func NewEstablishment(repository outbound.EstablishmentRepository, logger *logrus.Logger) inbound.EstablishmentUseCase {
	return &EstablishmentManager{repository: repository, logger: logger}
}

func (manager *EstablishmentManager) Add(ctx context.Context, payload *domain.Establishment) (*domain.Establishment, *errors.Exception) {
	establishment, err := manager.repository.Add(ctx, payload)
	if err != nil {
		manager.logger.Error("Error creating establishment.", "establishment_use_case", "err", err.Error())

		return nil, errors.Unexpected()
	}

	manager.logger.Info("Establishment created successfully.", "establishment_use_case")

	return establishment, nil
}

func (manager *EstablishmentManager) FindEstablishmentById(ctx context.Context, id string) (*domain.Establishment, *errors.Exception) {
	establishment, err := manager.repository.FindEstablishmentById(ctx, id)
	if err != nil {
		manager.logger.Error("Error finding establishment by id.", "establishment_use_case", "err", err.Error())

		return nil, errors.Unexpected(errors.WithMessage("error finding establishment"), errors.WithError(err))
	}

	manager.logger.Info("Establishment found successfully.", "establishment_use_case")

	return establishment, nil
}

func (manager *EstablishmentManager) GetAllProfessionalsByEstablishmentId(ctx context.Context, id string) ([]domain.Professionals, *errors.Exception) {
	professionals, err := manager.repository.GetAllProfessionalsByEstablishmentId(ctx, id)
	if err != nil {
		manager.logger.Error("Error getting professionals by establishment id.", "establishment_use_case", "err", err.Error())

		return nil, errors.Unexpected()
	}

	manager.logger.Info("Professionals retrieved successfully.", "establishment_use_case")

	return professionals, nil
}

func (manager *EstablishmentManager) UpdateEstablishmentById(ctx context.Context, id string, payload *domain.Establishment) (*domain.Establishment, *errors.Exception) {
	establishment, err := manager.repository.UpdateEstablishmentById(ctx, id, payload)
	if err != nil {
		manager.logger.Error("Error updating establishment.", "establishment_use_case", "err", err.Error())

		return nil, errors.Unexpected(errors.WithMessage("some error has been occurred trying update a establishment"))
	}

	manager.logger.Info("Establishment updated successfully.", "establishment_use_case")

	return establishment, nil
}

func (manager *EstablishmentManager) GetEstablishmentReport(ctx context.Context, id string) (*domain.EstablishmentMetrics, *errors.Exception) {
	stats, err := manager.repository.GetEstablishmentReport(ctx, id)
	if err != nil {
		manager.logger.Error("Error getting establishment report.", "establishment_use_case", "err", err.Error())

		return nil, errors.Unexpected(errors.WithMessage("some error has been occurred trying update a establishment"))
	}

	manager.logger.Info("Establishment report retrieved successfully.", "establishment_use_case")

	metrics := domain.EstablishmentMetrics{
		TotalProfessionals:    stats.TotalProfessionals,
		TotalServices:         stats.TotalServices,
		TotalRevenue:          0,
		TotalAppointments:     0,
		TotalAppointsCanceled: 0,
		TotalClients:          0,
	}

	return &metrics, nil
}
