package repository

import (
	"context"

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

func (repo *AppointmentDatabaseRepository) GetAllAppointmentsByProfessionalId(ctx context.Context) ([]domain.Appointment, error) {
	var appointments []domain.Appointment
	err := repo.db.WithContext(ctx).Find(&appointments).Error
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
	startTime string,
	endTime string,
) (bool, error) {

	var count int64

	err := repo.db.WithContext(ctx).
		Model(&domain.Appointment{}).
		Where("startTime = ? AND endTime = ?", startTime, endTime).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
