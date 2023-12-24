package config

import (
	"fmt"
	"os"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type AppConfig struct {
	DBURL              string `env:"DB_URL"`
	DBDriver           string `env:"DB_DRIVER"`
	ServeAddress       string `env:"SERVE_ADDRESS"`
	ServePort          string `env:"SERVE_PORT"`
	RedirectURL        string `env:"REDIRECT_URL"`
	GoogleClientID     string `env:"GOOGLE_CLIENT_ID"`
	GoogleClientSecret string `env:"GOOGLE_CLIENT_SECRET"`
	JWTSecret          string `env:"JWT_SECRET"`
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
