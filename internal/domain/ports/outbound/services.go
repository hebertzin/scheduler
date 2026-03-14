package outbound

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/hebertzin/scheduler/internal/domain"
)

type (
	ServicesRepository interface {
		Add(ctx context.Context, services *domain.Services) (*domain.Services, error)
		FindServiceById(ctx context.Context, serviceId string) (*domain.Services, error)
		GetAllServicesByProfessionalId(ctx context.Context, professionalId string) ([]domain.Services, error)
	}

	ServicesController interface {
		Add(ctx *gin.Context)
		FindServiceById(ctx *gin.Context)
		GetAllServicesByProfessionalId(ctx *gin.Context)
	}
)
