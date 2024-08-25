package gapi

import (
	"github.com/caarlos0/env/v11"
	"log"
	"time"
)

type Config struct {
	GrpcServerPort    string        `env:"GRPC_SERVER_PORT" ,envDefault:"9090"`
	GrpcServerTimeout time.Duration `env:"GRPC_SERVER_TIMEOUT" ,envDefault:"10s"`
}

func loadConfig() Config {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("Error parsing environment variables: %s", err)
	}
	return cfg
}
