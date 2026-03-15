package inbound

import (
	"context"

	"github.com/hebertzin/scheduler/internal/core"
	"github.com/hebertzin/scheduler/internal/domain"
)

type (
	AppointmentUseCase interface {
		Add(ctx context.Context, appointment *domain.Appointment) (*domain.Appointment, *core.Exception)
		GetAllAppointmentsByProfessionalId(ctx context.Context, professionalId string) ([]domain.Appointment, *core.Exception)
		GetAppointmentById(ctx context.Context, id string) (*domain.Appointment, *core.Exception)
		DeleteAppointment(ctx context.Context, id string) *core.Exception
	}
)
