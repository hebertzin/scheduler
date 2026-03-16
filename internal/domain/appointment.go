package domain

import "time"

type Appointment struct {
	ProfessionalId string    `json:"professionalId"`
	ServiceId      string    `json:"serviceId"`
	StartTime      time.Time `json:"startTime"`
	EndTime        time.Time `json:"endTime"`
	DayOfWeek      string    `json:"dayOfWeek"`
	Email          string    `json:"email"`
	Phone          string    `json:"phone"`
	Notes          string    `json:"notes"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}
