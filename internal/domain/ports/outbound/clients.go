package outbound

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/hebertzin/scheduler/internal/domain"
	"github.com/hebertzin/scheduler/internal/errors"
)

type (
	ClientRepository interface {
		Add(ctx context.Context, payload *domain.Client) (*domain.Client, *errors.Exception)
	}

	ClientController interface {
		Add(ctx *gin.Context)
	}
)
