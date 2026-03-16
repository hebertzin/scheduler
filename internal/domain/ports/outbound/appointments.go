package outbound

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hebertzin/scheduler/internal/domain"
)

type (
	AppointmentRepository interface {
		Add(ctx context.Context, appointment *domain.Appointment) (*domain.Appointment, error)
		GetAllAppointmentsByProfessionalId(ctx context.Context, professionalId string) ([]domain.Appointment, error)
		GetAppointmentById(ctx context.Context, id string) (*domain.Appointment, error)
		DeleteAppointment(ctx context.Context, id string) error
		ExistsByStartAndEndTime(ctx context.Context, startTime, endTime time.Time, dayOfWeek string) (bool, error)
	}

	AppointmentController interface {
		Add(ctx *gin.Context)
		GetAllAppointmentsByProfessionalId(ctx *gin.Context)
		GetAppointmentById(ctx *gin.Context)
		DeleteAppointment(ctx *gin.Context)
	}
)
