package usecases

import (
	"context"
	"time"

	"github.com/hebertzin/scheduler/internal/domain/ports/outbound"
	"github.com/sirupsen/logrus"
)

type AvailabilityManager struct {
	repo   outbound.AppointmentRepository
	logger *logrus.Logger
}

type Availability interface {
	ExistByStartAndEndTime(ctx context.Context, startTime, endTime time.Time, dayOfWeek string) (bool, error)
}

func NewAvailability(repo outbound.AppointmentRepository, logger *logrus.Logger) *AvailabilityManager {
	return &AvailabilityManager{
		repo:   repo,
		logger: logger,
	}
}

func (manager AvailabilityManager) ExistByStartAndEndTime(ctx context.Context, startTime, endTime time.Time, dayOfWeek string) (bool, error) {
	exist, err := manager.repo.ExistsByStartAndEndTime(ctx, startTime, endTime, dayOfWeek)
	if err != nil {
		manager.logger.Error("Error checking if appointment exists", "use_case_manager", "err", err.Error())

		return false, err
	}

	return exist, nil
}
