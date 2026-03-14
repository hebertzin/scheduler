package usecases

import (
	"context"

	"github.com/hebertzin/scheduler/internal/core"
	"github.com/hebertzin/scheduler/internal/domain"
	"github.com/sirupsen/logrus"
)

type (
	EstablishmentUserUseCase struct {
		repository domain.EstablishmentRepository
		logger     *logrus.Logger
	}

	EstablishmentMetrics struct {
		TotalProfessionals    int64
		TotalServices         int64
		TotalRevenue          int
		TotalAppointments     int
		TotalAppointsCanceled int
		TotalClients          int
	}
)

func NewEstablishmentUseCase(repository domain.EstablishmentRepository, logger *logrus.Logger) domain.EstablishmentUseCase {
	return &EstablishmentUserUseCase{repository: repository, logger: logger}
}

func (s *EstablishmentUserUseCase) FindEstablishmentById(ctx context.Context, establishmentId string) (*domain.Establishment, *core.Exception) {
	establishment, err := s.repository.FindEstablishmentById(ctx, establishmentId)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("error finding establishment"), core.WithError(err))
	}

	return establishment, nil
}

func (s *EstablishmentUserUseCase) Add(ctx context.Context, payload *domain.Establishment) (*domain.Establishment, *core.Exception) {
	establishment, err := s.repository.Add(ctx, payload)
	if err != nil {
		return nil, core.Unexpected()
	}

	return establishment, nil
}

func (s *EstablishmentUserUseCase) GetAllProfessionalsByEstablishmentId(ctx context.Context, establishmentId string) ([]domain.Professionals, *core.Exception) {
	professionals, err := s.repository.GetAllProfessionalsByEstablishmentId(ctx, establishmentId)
	if err != nil {
		return nil, core.Unexpected()
	}

	return professionals, nil
}

func (s *EstablishmentUserUseCase) UpdateEstablishmentById(ctx context.Context, establishmentId string, establishmentData *domain.Establishment) (*domain.Establishment, *core.Exception) {
	establishment, err := s.repository.UpdateEstablishmentById(ctx, establishmentId, establishmentData)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("some error has been occurred trying update a establishment"))
	}

	return establishment, nil
}

func (s *EstablishmentUserUseCase) GetEstablishmentReport(ctx context.Context, establishmentId string) (*EstablishmentMetrics, *core.Exception) {
	stats, err := s.repository.GetEstablishmentReport(ctx, establishmentId)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("some error has been occurred trying update a establishment"))
	}

	metrics := EstablishmentMetrics{
		TotalProfessionals:    stats.TotalProfessionals,
		TotalServices:         stats.TotalServices,
		TotalRevenue:          0,
		TotalAppointments:     0,
		TotalAppointsCanceled: 0,
		TotalClients:          0,
	}

	return &metrics, nil
}
