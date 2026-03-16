package repository

import (
	"context"
	"time"

	"github.com/hebertzin/scheduler/internal/domain"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AppointmentDatabaseRepository struct {
	db     *gorm.DB
	logger *logrus.Logger
}

func NewAppointmentRepository(db *gorm.DB, logger *logrus.Logger) *AppointmentDatabaseRepository {
	return &AppointmentDatabaseRepository{
		db:     db,
		logger: logger,
	}
}

func (repo *AppointmentDatabaseRepository) Add(ctx context.Context, appointment *domain.Appointment) (*domain.Appointment, error) {
	if err := repo.db.WithContext(ctx).
		Create(appointment).Error; err != nil {
		return nil, err
	}
	return appointment, nil
}

func (repo *AppointmentDatabaseRepository) GetAllAppointmentsByProfessionalId(ctx context.Context, professionalId string) ([]domain.Appointment, error) {
	var appointments []domain.Appointment
	err := repo.db.WithContext(ctx).Find(&appointments).Where("professionalId = ?", professionalId).Error
	if err != nil {
		return nil, err
	}
	return appointments, nil
}

func (repo *AppointmentDatabaseRepository) GetAppointmentById(ctx context.Context, appointmentId string) (*domain.Appointment, error) {
	var appointment domain.Appointment
	if err := repo.db.WithContext(ctx).
		Where("id = ?", appointmentId).
		First(&appointment).Error; err != nil {
		return nil, err
	}
	return &appointment, nil
}

func (repo *AppointmentDatabaseRepository) DeleteAppointment(ctx context.Context, appointmentId string) error {
	err := repo.db.WithContext(ctx).Where("id = ?", appointmentId).Delete(&appointmentId).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *AppointmentDatabaseRepository) ExistsByStartAndEndTime(
	ctx context.Context,
	startTime time.Time,
	endTime time.Time,
	dayOfWeek string,
) (bool, error) {

	var count int64

	err := repo.db.WithContext(ctx).
		Model(&domain.Appointment{}).
		Where("day_of_week = ? AND start_time < ? AND end_time > ?", dayOfWeek, endTime, startTime).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
