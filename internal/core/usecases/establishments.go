package usecases

import (
	"context"

	"github.com/hebertzin/scheduler/internal/core"
	"github.com/hebertzin/scheduler/internal/domain"
	"github.com/sirupsen/logrus"
)

type EstablishmentUserUseCase struct {
	repository domain.EstablishmentRepository
	logger     *logrus.Logger
}

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

func (s *EstablishmentUserUseCase) GetEstablishmentReport(ctx context.Context, establishmentId string) (*domain.EstablishmentReport, *core.Exception) {
	stats, err := s.repository.GetEstablishmentReport(ctx, establishmentId)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("some error has been occurred trying update a establishment"))
	}

	return stats, nil
}
