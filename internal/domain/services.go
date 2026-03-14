package domain

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hebertzin/scheduler/internal/core"
)

type (
	Services struct {
		Name           string    `json:"name"`
		Value          string    `json:"value"`
		Duration       string    `json:"duration"`
		ProfessionalId string    `json:"professionalId"`
		CreatedAt      time.Time `json:"createdAt"`
		UpdatedAt      time.Time `json:"updatedAt"`
	}

	ServicesUseCase interface {
		Add(ctx context.Context, payload *Services) (*Services, *core.Exception)
		FindServiceById(ctx context.Context, id string) (*Services, *core.Exception)
		GetAllServicesByProfessionalId(ctx context.Context, professional_id string) ([]Services, *core.Exception)
	}

	ServicesRepository interface {
		Add(ctx context.Context, establishment *Services) (*Services, error)
		FindServiceById(ctx context.Context, service_id string) (*Services, error)
		GetAllServicesByProfessionalId(ctx context.Context, professional_id string) ([]Services, error)
	}

	ServicesController interface {
		Add(ctx *gin.Context)
		FindServiceById(ctx *gin.Context)
		GetAllServicesByProfessionalId(ctx *gin.Context)
	}
)
