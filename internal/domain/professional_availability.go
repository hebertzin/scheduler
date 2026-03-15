package domain

import (
	"time"

	"github.com/google/uuid"
)

type (
	ProfessionalAvailability struct {
		ProfessionalId uuid.UUID `json:"professionalId"`
		DayOfWeek      string    `json:"dayOfWeek"`
		StartTime      string    `json:"startTime"`
		EndTime        string    `json:"endTime"`
		CreatedAt      time.Time `json:"createdAt"`
		UpdatedAt      time.Time `json:"updatedAt"`
	}
)
