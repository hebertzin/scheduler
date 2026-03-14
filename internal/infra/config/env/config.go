package env

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/hebertzin/scheduler/internal/domain"
	"github.com/hebertzin/scheduler/internal/infra/config/logging"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func LoadEnvConfig() *domain.ServiceConfig {
	err := godotenv.Load()
	if err != nil {
		logging.Log.WithFields(logrus.Fields{
			"method": "LoadEnvConfig",
		}).Error("Error loading .env file")
	}

	c := &domain.ServiceConfig{
		Port:                os.Getenv("PORT"),
		RunMigrationEnabled: false,
		SwaggerEnabled:      true,
		DevModeEnabled:      false,
		GrafanaEnabled:      false,
		LoggingEnabled:      true,
		Database: domain.DatabaseConfig{
			User:     os.Getenv("USER"),
			Port:     os.Getenv("PORT"),
			Password: os.Getenv("PASSWORD"),
			Host:     os.Getenv("HOST"),
			Database: os.Getenv("DATABASE"),
		},
	}

	if c.Database.User == "" || c.Database.Password == "" || c.Database.Database == "" || c.Database.Port == "" {
		println("Some env fields are missing")
	}

	return c
}

func LoadJSONConfig(path string) (*domain.ServiceConfig, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	var closeErr error
	defer func() {
		closeErr = file.Close()
	}()
	if closeErr != nil {
		return nil, fmt.Errorf("failed to close file: %w", closeErr)
	}

	decoder := json.NewDecoder(file)
	var config domain.ServiceConfig
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func LoadConfiguration(configPath string) (*domain.ServiceConfig, error) {
	serviceConfig, err := LoadJSONConfig(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load service config: %w", err)
	}

	if serviceConfig.DevModeEnabled {
		logging.Log.Info("DevMode enabled — loading configuration from JSON")
		return serviceConfig, nil
	}

	logging.Log.Info("DevMode disabled — loading configuration from .env")
	return LoadEnvConfig(), nil
}
