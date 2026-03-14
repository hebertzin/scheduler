package outbound

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/hebertzin/scheduler/internal/domain"
)

type (
	ProfessionalsAvailabilityRepository interface {
		Add(ctx context.Context, availability *domain.ProfessionalAvailability) (*domain.ProfessionalAvailability, error)
		GetProfessionalAvailabilityById(ctx context.Context, id string) ([]domain.ProfessionalAvailability, error)
	}

	ProfessionalAvailabilityController interface {
		Add(ctx *gin.Context)
		GetProfessionalAvailabilityById(ctx *gin.Context)
	}
)
