package domain

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hebertzin/scheduler/internal/core"
)

type (
	Establishment struct {
		Name       string    `json:"name"`        // Name of the establishment
		City       string    `json:"city"`        // City where the establishment is located
		State      string    `json:"state"`       // State where the establishment is located
		PostalCode string    `json:"postal_code"` // Postal code of the establishment
		Number     string    `json:"number"`      // Street number of the establishment
		UserId     string    `json:"user_id"`     // ID of the user who owns the establishment
		CreatedAt  time.Time `json:"created_at"`  // Timestamp of creation
		UpdatedAt  time.Time `json:"updated_at"`  // Timestamp of last update
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
		TotalProfessionals int64 `json:"total_professionals"`
		TotalServices      int64 `json:"total_services"`
	}

	// EstablishmentUseCase defines the business logic for establishments.
	EstablishmentUseCase interface {
		// Add creates a new establishment.
		Add(ctx context.Context, payload *Establishment) (*Establishment, *core.Exception)
		// GetAllProfessionalsByEstablishmentId retrieves all professionals linked to a given establishment.
		GetAllProfessionalsByEstablishmentId(ctx context.Context, establishment_id string) ([]Professionals, *core.Exception)
		// FindEstablishmentById retrieves an establishment by its ID.
		FindEstablishmentById(ctx context.Context, establishment_id string) (*Establishment, *core.Exception)
		// GetEstablishmentReport generates a report for a given establishment.
		GetEstablishmentReport(ctx context.Context, establishment_id string) (*EstablishmentMetrics, *core.Exception)
		// UpdateEstablishmentById updates establishment data by its ID.
		UpdateEstablishmentById(ctx context.Context, establishment_id string, establishmentData *Establishment) (*Establishment, *core.Exception)
	}

	// EstablishmentRepository defines the data access methods for establishments.
	EstablishmentRepository interface {
		// Add inserts a new establishment into the database.
		Add(ctx context.Context, establishment *Establishment) (*Establishment, error)
		// GetAllProfessionalsByEstablishmentId fetches all professionals associated with a specific establishment.
		GetAllProfessionalsByEstablishmentId(ctx context.Context, establishment_id string) ([]Professionals, error)
		// FindEstablishmentById retrieves an establishment by its ID from the database.
		FindEstablishmentById(ctx context.Context, email string) (*Establishment, error)
		// GetEstablishmentReport retrieves a report for the specified establishment.
		GetEstablishmentReport(ctx context.Context, establishment_id string) (*EstablishmentReport, error)
		// UpdateEstablishmentById updates the data of an existing establishment in the database.
		UpdateEstablishmentById(ctx context.Context, establishment_id string, establishmentData *Establishment) (*Establishment, error)
	}

	// EstablishmentController defines the HTTP handlers for establishment operations.
	EstablishmentController interface {
		// Add handles the HTTP request to create a new establishment.
		Add(ctx *gin.Context)
		// FindEstablishmentById handles the HTTP request to retrieve an establishment by ID.
		FindEstablishmentById(ctx *gin.Context)
		// GetAllProfessinalsByEstablishmentId handles the HTTP request to get all professionals linked to an establishment.
		GetAllProfessinalsByEstablishmentId(ctx *gin.Context)
		// UpdateEstablishmentById handles the HTTP request to update establishment data by ID.
		UpdateEstablishmentById(ctx *gin.Context)
		// GetEstablishmentReport handles the HTTP request to generate a report for a specific establishment.
		GetEstablishmentReport(ctx *gin.Context)
	}
)
