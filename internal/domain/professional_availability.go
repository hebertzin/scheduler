package domain

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hebertzin/scheduler/internal/core"
)

type (
	ProfessionalAvailability struct {
		ProfessionalId uuid.UUID `json:"professionalId"`
		DayOfWeek      string    `json:"dayOfWeek"`
		StartTime      string    `json:"startTime"`
		EndTime        string    `json:"endTime"`
		CreatedAt      time.Time `json:"createdAt"`
		UpdatedAt      time.Time `json:"updatedAt"`
	}

	ProfessionalsAvailabilityUseCase interface {
		Add(ctx context.Context, availability *ProfessionalAvailability) (*ProfessionalAvailability, *core.Exception)
		GetProfessionalAvailabilityById(ctx context.Context, id string) ([]ProfessionalAvailability, *core.Exception)
	}

	ProfessionalsAvailabilityRepository interface {
		Add(ctx context.Context, availability *ProfessionalAvailability) (*ProfessionalAvailability, error)
		GetProfessionalAvailabilityById(ctx context.Context, id string) ([]ProfessionalAvailability, error)
	}

	ProfessionalAvailabilityController interface {
		Add(ctx *gin.Context)
		GetProfessionalAvailabilityById(ctx *gin.Context)
	}
)
