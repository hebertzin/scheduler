package domain

import (
	"time"
)

type (
	Professionals struct {
		Name            string    `json:"name"`
		Role            string    `json:"role"`
		EstablishmentId string    `json:"establishment"`
		CreatedAt       time.Time `json:"createdAt"`
		UpdatedAt       time.Time `json:"updatedAt"`
	}
)
