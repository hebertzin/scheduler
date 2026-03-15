package domain

import "time"

type Appointment struct {
	ProfessionalId string    `json:"professionalId"`
	ServiceId      string    `json:"serviceId"`
	StartTime      string    `json:"startTime"`
	DayOfWeek      string    `json:"dayOfWeek"`
	EndTime        string    `json:"endTime"`
	Email          string    `json:"email"`
	Phone          string    `json:"phone"`
	Notes          string    `json:"notes"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}
