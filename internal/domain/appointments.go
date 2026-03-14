package domain

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hebertzin/scheduler/internal/core"
)

type (
	Appointment struct {
		ProfessionalId string    `json:"professionalId"`
		ServiceId      string    `json:"serviceId"`
		StartTime      string    `json:"startTime"`
		EndTime        string    `json:"endTime"`
		Email          string    `json:"email"`
		Phone          string    `json:"phone"`
		Notes          string    `json:"notes"`
		CreatedAt      time.Time `json:"createdAt"`
		UpdatedAt      time.Time `json:"updatedAt"`
	}

	AppointmentUseCase interface {
		Add(ctx context.Context, appointment *Appointment) (*Appointment, *core.Exception)
		GetAllAppointmentsByProfessionalId(ctx context.Context, professionalId string) ([]Appointment, *core.Exception)
		GetAppointmentById(ctx context.Context, id string) (*Appointment, *core.Exception)
		DeleteAppointment(ctx context.Context, id string) *core.Exception
	}

	AppointmentRepository interface {
		Add(ctx context.Context, appointment *Appointment) (*Appointment, error)
		GetAllAppointmentsByProfessionalId(ctx context.Context, professionalId string) ([]Appointment, error)
		GetAppointmentById(ctx context.Context, id string) (*Appointment, error)
		DeleteAppointment(ctx context.Context, id string) error
		ExistsByStartAndEndTime(ctx context.Context, startTime string, endTime string) (bool, error)
	}

	AppointmentController interface {
		Add(ctx *gin.Context)
		GetAllAppointmentsByProfessionalId(ctx *gin.Context)
		GetAppointmentById(ctx *gin.Context)
		DeleteAppointment(ctx *gin.Context)
	}
)
