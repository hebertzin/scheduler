package usecases

import (
	"context"

	"github.com/hebertzin/scheduler/internal/core"
	"github.com/hebertzin/scheduler/internal/domain"
	"github.com/hebertzin/scheduler/internal/domain/ports/inbound"
	"github.com/hebertzin/scheduler/internal/domain/ports/outbound"
	"github.com/sirupsen/logrus"
)

type ServicesManger struct {
	repository outbound.ServicesRepository
	logger     *logrus.Logger
}

func NewServices(repository outbound.ServicesRepository, logger *logrus.Logger) inbound.ServicesUseCase {
	return &ServicesManger{repository: repository, logger: logger}
}

func (manager *ServicesManger) FindServiceById(ctx context.Context, id string) (*domain.Services, *core.Exception) {
	service, err := manager.repository.FindServiceById(ctx, id)
	if err != nil {
		manager.logger.Error("Error finding service by id.", "services_use_case", "err", err.Error())

		return nil, core.Unexpected(core.WithMessage("error finding service"), core.WithError(err))
	}

	manager.logger.Info("Service found successfully.", "services_use_case")

	return service, nil
}

func (manager *ServicesManger) Add(ctx context.Context, payload *domain.Services) (*domain.Services, *core.Exception) {
	if payload.Name == "" || payload.Duration == "" {
		manager.logger.Error("Missing required fields to create service.", "services_use_case")

		return nil, core.BadRequest(core.WithMessage("some fields are missing"))
	}

	service, err := manager.repository.Add(ctx, payload)
	if err != nil {
		manager.logger.Error("Error creating service.", "services_use_case", "err", err.Error())

		return nil, core.Unexpected()
	}

	manager.logger.Info("Service created successfully.", "services_use_case")

	return service, nil
}

func (manager *ServicesManger) GetAllServicesByProfessionalId(ctx context.Context, professionalId string) ([]domain.Services, *core.Exception) {
	services, err := manager.repository.GetAllServicesByProfessionalId(ctx, professionalId)
	if err != nil {
		manager.logger.Error("Error getting services by professional id.", "services_use_case", "err", err.Error())

		return nil, core.Unexpected()
	}

	manager.logger.Info("Services retrieved successfully.", "services_use_case")

	return services, nil
}
