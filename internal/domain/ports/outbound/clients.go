package outbound

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/hebertzin/scheduler/internal/core"
	"github.com/hebertzin/scheduler/internal/domain"
)

type (
	ClientRepository interface {
		Add(ctx context.Context, payload *domain.Client) (*domain.Client, *core.Exception)
	}

	ClientController interface {
		Add(ctx *gin.Context)
	}
)
