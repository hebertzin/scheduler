package domain

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hebertzin/scheduler/internal/core"
)

type (
	Professionals struct {
		Name            string    `json:"name"`
		Role            string    `json:"role"`
		EstablishmentId string    `json:"establishment"`
		CreatedAt       time.Time `json:"createdAt"`
		UpdatedAt       time.Time `json:"updatedAt"`
	}

	ProfessionalsUseCase interface {
		Add(ctx context.Context, payload *Professionals) (*Professionals, *core.Exception)
		FindProfessionalById(ctx context.Context, id string) (*Professionals, *core.Exception)
		UpdateProfessionalById(ctx context.Context, professional_id string, professionalData *Professionals) (*Professionals, *core.Exception)
	}

	ProfessionalsRepository interface {
		Add(ctx context.Context, establishment *Professionals) (*Professionals, error)
		FindProfessionalById(ctx context.Context, email string) (*Professionals, error)
		UpdateProfessionalById(ctx context.Context, professional_id string, professionalData *Professionals) (*Professionals, error)
	}

	ProfessionalsController interface {
		Add(ctx *gin.Context)
		FindProfessionalById(ctx *gin.Context)
		UpdateProfessionalById(ctx *gin.Context)
	}
)
