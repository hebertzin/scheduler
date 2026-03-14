package outbound

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/hebertzin/scheduler/internal/domain"
)

type (
	EstablishmentRepository interface {
		Add(ctx context.Context, establishment *domain.Establishment) (*domain.Establishment, error)
		GetAllProfessionalsByEstablishmentId(ctx context.Context, id string) ([]domain.Professionals, error)
		FindEstablishmentById(ctx context.Context, email string) (*domain.Establishment, error)
		GetEstablishmentReport(ctx context.Context, id string) (*domain.EstablishmentReport, error)
		UpdateEstablishmentById(ctx context.Context, id string, payload *domain.Establishment) (*domain.Establishment, error)
	}

	EstablishmentController interface {
		Add(ctx *gin.Context)
		FindEstablishmentById(ctx *gin.Context)
		GetAllProfessinalsByEstablishmentId(ctx *gin.Context)
		UpdateEstablishmentById(ctx *gin.Context)
		GetEstablishmentReport(ctx *gin.Context)
	}
)
