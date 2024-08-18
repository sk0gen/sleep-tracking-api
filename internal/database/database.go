package database

import (
	"context"
	"fmt"
	"github.com/caarlos0/env/v11"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log"
)

type Config struct {
	Name     string `env:"DB_DATABASE"`
	User     string `env:"DB_USERNAME"`
	Password string `env:"DB_PASSWORD"`
	Host     string `env:"DB_HOST"`
	Port     string `env:"DB_PORT"`
	Schema   string `env:"DB_SCHEMA"`
}

func New() (*pgxpool.Pool, error) {

	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("Error parsing environment variables: %s", err)
	}

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name, cfg.Schema)
	dbPool, err := pgxpool.New(context.Background(), connStr)
	if nil != err {
		log.Fatal(err)
	}
	return dbPool, nil
}
