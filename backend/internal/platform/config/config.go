// Package config loads application settings from environment variables.
package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

// Config holds every setting the API needs to run.
type Config struct {
	AppEnv             string   `env:"APP_ENV" envDefault:"development"`
	HTTPPort           int      `env:"HTTP_PORT" envDefault:"8080"`
	DatabaseURL        string   `env:"DATABASE_URL,required"`
	LogLevel           string   `env:"LOG_LEVEL" envDefault:"info"`
	CORSAllowedOrigins []string `env:"CORS_ALLOWED_ORIGINS" envSeparator:"," envDefault:"http://localhost:3000"`
}

// IsDevelopment reports whether the app is running locally.
func (c Config) IsDevelopment() bool {
	return c.AppEnv == "development"
}

// Load reads a .env file (if present) and parses environment variables into Config.
func Load() (Config, error) {
	// A missing .env is fine: in production, variables come from the environment.
	_ = godotenv.Load()

	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return Config{}, fmt.Errorf("parse env: %w", err)
	}
	return cfg, nil
}
