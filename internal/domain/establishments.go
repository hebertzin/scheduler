package domain

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hebertzin/scheduler/internal/core"
)

type (
	Establishment struct {
		Name       string    `json:"name"`
		City       string    `json:"city"`
		State      string    `json:"state"`
		PostalCode string    `json:"postalCode"`
		Number     string    `json:"number"`
		UserId     string    `json:"userId"`
		CreatedAt  time.Time `json:"createdAt"`
		UpdatedAt  time.Time `json:"updatedAt"`
	}

	EstablishmentMetrics struct {
		TotalProfessionals    int64
		TotalServices         int64
		TotalRevenue          int
		TotalAppointments     int
		TotalAppointsCanceled int
		TotalClients          int
	}

	EstablishmentReport struct {
		TotalProfessionals int64 `json:"professionalsCount"`
		TotalServices      int64 `json:"serviceCount"`
	}

	EstablishmentUseCase interface {
		Add(ctx context.Context, payload *Establishment) (*Establishment, *core.Exception)
		GetAllProfessionalsByEstablishmentId(ctx context.Context, id string) ([]Professionals, *core.Exception)
		FindEstablishmentById(ctx context.Context, id string) (*Establishment, *core.Exception)
		GetEstablishmentReport(ctx context.Context, id string) (*EstablishmentMetrics, *core.Exception)
		UpdateEstablishmentById(ctx context.Context, id string, payload *Establishment) (*Establishment, *core.Exception)
	}

	EstablishmentRepository interface {
		Add(ctx context.Context, establishment *Establishment) (*Establishment, error)
		GetAllProfessionalsByEstablishmentId(ctx context.Context, id string) ([]Professionals, error)
		FindEstablishmentById(ctx context.Context, email string) (*Establishment, error)
		GetEstablishmentReport(ctx context.Context, id string) (*EstablishmentReport, error)
		UpdateEstablishmentById(ctx context.Context, id string, payload *Establishment) (*Establishment, error)
	}

	EstablishmentController interface {
		Add(ctx *gin.Context)
		FindEstablishmentById(ctx *gin.Context)
		GetAllProfessinalsByEstablishmentId(ctx *gin.Context)
		UpdateEstablishmentById(ctx *gin.Context)
		GetEstablishmentReport(ctx *gin.Context)
	}
)
