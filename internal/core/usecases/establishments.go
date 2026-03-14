package usecases

import (
	"context"

	"github.com/hebertzin/scheduler/internal/core"
	"github.com/hebertzin/scheduler/internal/domain"
	"github.com/sirupsen/logrus"
)

type EstablishmentManager struct {
	repository domain.EstablishmentRepository
	logger     *logrus.Logger
}

func NewEstablishment(repository domain.EstablishmentRepository, logger *logrus.Logger) domain.EstablishmentUseCase {
	return &EstablishmentManager{repository: repository, logger: logger}
}

func (e *EstablishmentManager) FindEstablishmentById(ctx context.Context, establishmentId string) (*domain.Establishment, *core.Exception) {
	establishment, err := e.repository.FindEstablishmentById(ctx, establishmentId)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("error finding establishment"), core.WithError(err))
	}

	return establishment, nil
}

func (e *EstablishmentManager) Add(ctx context.Context, payload *domain.Establishment) (*domain.Establishment, *core.Exception) {
	establishment, err := e.repository.Add(ctx, payload)
	if err != nil {
		return nil, core.Unexpected()
	}

	return establishment, nil
}

func (e *EstablishmentManager) GetAllProfessionalsByEstablishmentId(ctx context.Context, establishmentId string) ([]domain.Professionals, *core.Exception) {
	professionals, err := e.repository.GetAllProfessionalsByEstablishmentId(ctx, establishmentId)
	if err != nil {
		return nil, core.Unexpected()
	}

	return professionals, nil
}

func (e *EstablishmentManager) UpdateEstablishmentById(ctx context.Context, establishmentId string, establishmentData *domain.Establishment) (*domain.Establishment, *core.Exception) {
	establishment, err := e.repository.UpdateEstablishmentById(ctx, establishmentId, establishmentData)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("some error has been occurred trying update a establishment"))
	}

	return establishment, nil
}

func (e *EstablishmentManager) GetEstablishmentReport(ctx context.Context, establishmentId string) (*domain.EstablishmentMetrics, *core.Exception) {
	stats, err := e.repository.GetEstablishmentReport(ctx, establishmentId)
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
