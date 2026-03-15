package outbound

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/hebertzin/scheduler/internal/domain"
)

type (
	ProfessionalsRepository interface {
		Add(ctx context.Context, establishment *domain.Professionals) (*domain.Professionals, error)
		FindProfessionalById(ctx context.Context, email string) (*domain.Professionals, error)
		UpdateProfessionalById(ctx context.Context, professional_id string, professionalData *domain.Professionals) (*domain.Professionals, error)
	}

	ProfessionalsController interface {
		Add(ctx *gin.Context)
		FindProfessionalById(ctx *gin.Context)
		UpdateProfessionalById(ctx *gin.Context)
	}
)
