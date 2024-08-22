package api

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	db "github.com/sk0gen/sleep-tracking-api/internal/database/sqlc"
	"log"
	"time"
)

type ApiConfig struct {
	HttpServerPort     string        `env:"HTTP_SERVER_PORT,required"`
	HttpServerTimeout  time.Duration `env:"HTTP_SERVER_TIMEOUT" ,envDefault:"10s"`
	JWTSecret          string        `env:"JWT_SECRET"`
	JWTTokenExpiration time.Duration `env:"JWT_TOKEN_EXPIRATION" ,envDefault:"1h"`
}

type Config struct {
	ApiConfig ApiConfig
	Database  db.Config
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
