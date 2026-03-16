package inbound

import (
	"context"

	"github.com/hebertzin/scheduler/internal/domain"
	"github.com/hebertzin/scheduler/internal/errors"
)

type (
	AppointmentUseCase interface {
		Add(ctx context.Context, appointment *domain.Appointment) (*domain.Appointment, *errors.Exception)
		GetAllAppointmentsByProfessionalId(ctx context.Context, professionalId string) ([]domain.Appointment, *errors.Exception)
		GetAppointmentById(ctx context.Context, id string) (*domain.Appointment, *errors.Exception)
		DeleteAppointment(ctx context.Context, id string) *errors.Exception
	}
)
