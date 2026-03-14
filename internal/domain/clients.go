package domain

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/hebertzin/scheduler/internal/core"
)

// represent client in the system
type (
	Client struct {
		Email    string `json:"email"`
		Phone    string `json:"phone"`
		Location string `json:"location"`
	}

	ClientUseCase interface {
		Add(ctx context.Context, payload *Client) (*Client, *core.Exception)
	}

	ClientRepository interface {
		Add(ctx context.Context, payload *Client) (*Client, *core.Exception)
	}

	ClientController interface {
		Add(ctx *gin.Context)
	}
)
