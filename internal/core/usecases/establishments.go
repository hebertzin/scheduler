package usecases

import (
	"context"

	"github.com/hebertzin/scheduler/internal/core"
	"github.com/hebertzin/scheduler/internal/domain"
	"github.com/hebertzin/scheduler/internal/domain/ports/inbound"
	"github.com/hebertzin/scheduler/internal/domain/ports/outbound"
	"github.com/sirupsen/logrus"
)

type EstablishmentManager struct {
	repository outbound.EstablishmentRepository
	logger     *logrus.Logger
}

func NewEstablishment(repository outbound.EstablishmentRepository, logger *logrus.Logger) inbound.EstablishmentUseCase {
	return &EstablishmentManager{repository: repository, logger: logger}
}

func (e *EstablishmentManager) Add(ctx context.Context, payload *domain.Establishment) (*domain.Establishment, *core.Exception) {
	establishment, err := e.repository.Add(ctx, payload)
	if err != nil {
		return nil, core.Unexpected()
	}

	return establishment, nil
}

func (e *EstablishmentManager) FindEstablishmentById(ctx context.Context, id string) (*domain.Establishment, *core.Exception) {
	establishment, err := e.repository.FindEstablishmentById(ctx, id)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("error finding establishment"), core.WithError(err))
	}

	return establishment, nil
}

func (e *EstablishmentManager) GetAllProfessionalsByEstablishmentId(ctx context.Context, id string) ([]domain.Professionals, *core.Exception) {
	professionals, err := e.repository.GetAllProfessionalsByEstablishmentId(ctx, id)
	if err != nil {
		return nil, core.Unexpected()
	}

	return professionals, nil
}

func (e *EstablishmentManager) UpdateEstablishmentById(ctx context.Context, id string, payload *domain.Establishment) (*domain.Establishment, *core.Exception) {
	establishment, err := e.repository.UpdateEstablishmentById(ctx, id, payload)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("some error has been occurred trying update a establishment"))
	}

	return establishment, nil
}

func (e *EstablishmentManager) GetEstablishmentReport(ctx context.Context, id string) (*domain.EstablishmentMetrics, *core.Exception) {
	stats, err := e.repository.GetEstablishmentReport(ctx, id)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("some error has been occurred trying update a establishment"))
	}

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
