package domain

import (
	"time"
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
)
