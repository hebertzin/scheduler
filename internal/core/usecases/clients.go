package usecases

import (
	"context"

	"github.com/hebertzin/scheduler/internal/core"
	"github.com/hebertzin/scheduler/internal/domain"
	"github.com/hebertzin/scheduler/internal/domain/ports/inbound"
	"github.com/hebertzin/scheduler/internal/domain/ports/outbound"
	"github.com/sirupsen/logrus"
)

type ClientManager struct {
	repository inbound.ClientUseCase
	logger     *logrus.Logger
}

func NewClientUseCase(repository outbound.ClientRepository, logger *logrus.Logger) inbound.ClientUseCase {
	return &ClientManager{repository: repository, logger: logger}
}

func (s *ClientManager) Add(ctx context.Context, account *domain.Client) (*domain.Client, *core.Exception) {
	client, err := s.repository.Add(ctx, account)
	if err != nil {
		s.logger.Error("Error creating client.", "client_use_case", "err", err.Error())

		return nil, core.Unexpected(core.WithMessage("error creating client"), core.WithError(err))
	}

	s.logger.Info("Client created successfully.", "client_use_case")

	return client, nil
}
