package models

import (
	"time"

	"github.com/google/uuid"
)

type Appointment struct {
	ID             uuid.UUID    `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	ProfessionalID uuid.UUID    `gorm:"type:uuid;not null"`
	Professional   Professional `gorm:"foreignKey:ProfessionalID"`
	ServiceID      uuid.UUID    `gorm:"type:uuid;not null"`
	Service        Services     `gorm:"foreignKey:ServiceID"`
	DayOfWeek      string       `gorm:"not null"`
	StartTime      string       `gorm:"size:100"`
	EndTime        string       `gorm:"size:20"`
	Email          string       `gorm:"size:100"`
	Phone          string       `gorm:"size:20"`
	Notes          string       `gorm:"size:255"`
	CreatedAt      time.Time    `json:"created_at"`
	UpdatedAt      time.Time    `json:"updated_at"`
}
