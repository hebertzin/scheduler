package domain

import (
	"time"
)

type (
	Services struct {
		Name           string    `json:"name"`
		Value          string    `json:"value"`
		Duration       string    `json:"duration"`
		ProfessionalId string    `json:"professionalId"`
		CreatedAt      time.Time `json:"createdAt"`
		UpdatedAt      time.Time `json:"updatedAt"`
	}
)
