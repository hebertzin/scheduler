package usecases

import (
	"context"

	"github.com/hebertzin/scheduler/internal/core"
	"github.com/hebertzin/scheduler/internal/domain"
	"github.com/hebertzin/scheduler/internal/domain/ports/inbound"
	"github.com/hebertzin/scheduler/internal/domain/ports/outbound"
	"github.com/sirupsen/logrus"
)

type ClientUseCase struct {
	repository inbound.ClientUseCase
	logger     *logrus.Logger
}

func NewClientUseCase(repository outbound.ClientRepository, logger *logrus.Logger) inbound.ClientUseCase {
	return &ClientUseCase{repository: repository, logger: logger}
}

func (s *ClientUseCase) Add(ctx context.Context, account *domain.Client) (*domain.Client, *core.Exception) {
	client, err := s.repository.Add(ctx, account)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("error creating client"), core.WithError(err))
	}

	return client, nil
}
