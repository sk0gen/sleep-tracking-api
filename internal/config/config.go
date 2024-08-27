package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"github.com/sk0gen/sleep-tracking-api/internal/database/sqlc"
	"github.com/sk0gen/sleep-tracking-api/internal/logging"
	"github.com/sk0gen/sleep-tracking-api/internal/token"
	"log"
)

type Config struct {
	Database   db.Config
	AuthConfig token.Config
	LogConfig  logging.Config
}

func NewConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("Error parsing environment variables: %s", err)
	}

	return cfg
}
