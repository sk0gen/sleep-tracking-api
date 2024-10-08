package api

import (
	"github.com/caarlos0/env/v11"
	"log"
	"time"
)

type Config struct {
	HttpServerPort    string        `env:"HTTP_SERVER_PORT" ,envDefault:"8080"`
	HttpServerTimeout time.Duration `env:"HTTP_SERVER_TIMEOUT" ,envDefault:"10s"`
}

func loadConfig() Config {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("Error parsing environment variables: %s", err)
	}
	return cfg
}
