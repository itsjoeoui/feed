package config

import (
	"fmt"
	"os"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type AppConfig struct {
	DbURL        string `env:"DB_URL"`
	DbDriver     string `env:"DB_DRIVER"`
	ServeAddress string `env:"SERVE_ADDRESS"`
	ServePort    string `env:"SERVE_PORT"`
}

func LoadAppConfig() (AppConfig, error) {
	// Only need to load from .env if we are running locally
	if os.Getenv("FLY_APP_NAME") == "" {
		err := godotenv.Load()
		if err != nil {
			return AppConfig{}, fmt.Errorf("error loading .env file: %w", err)
		}
	}

	cfg := AppConfig{}

	if err := env.Parse(&cfg); err != nil {
		return AppConfig{}, fmt.Errorf("%+v", err)
	}

	return cfg, nil
}
