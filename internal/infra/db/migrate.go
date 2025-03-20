package db

import (
	"fmt"

	"github.com/hebertzin/scheduler/internal/infra/db/models"
	"gorm.io/gorm"
)

func Migrate(database *gorm.DB) error {
	if err := database.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"").Error; err != nil {
		return fmt.Errorf("failed to enable uuid-ossp extension: %w", err)
	}
	err := database.AutoMigrate(
		&models.Users{},
		&models.Establishment{},
		&models.Professional{},
		&models.Services{},
		&models.ProfessionalAvailability{},
		&models.Appointment{},
	)
	if err != nil {
		return fmt.Errorf("failed to migrate models: %w", err)
	}

	fmt.Println("Database migrated successfully")
	return nil
}
