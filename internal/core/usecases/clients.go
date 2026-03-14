package usecases

import (
	"context"

	"github.com/hebertzin/scheduler/internal/core"
	"github.com/hebertzin/scheduler/internal/domain"
	"github.com/sirupsen/logrus"
)

type ClientUseCase struct {
	repository domain.ClientUseCase
	logger     *logrus.Logger
}

func NewClientUseCase(repository domain.ClientRepository, logger *logrus.Logger) domain.ClientUseCase {
	return &ClientUseCase{repository: repository, logger: logger}
}

func (s *ClientUseCase) Add(ctx context.Context, account *domain.Client) (*domain.Client, *core.Exception) {
	client, err := s.repository.Add(ctx, account)
	if err != nil {
		return nil, core.Unexpected(core.WithMessage("error creating client"), core.WithError(err))
	}

	return client, nil
}
