package db

import (
	"context"
	"fmt"
	"github.com/caarlos0/env/v11"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

type config struct {
	Name     string `env:"DB_DATABASE"`
	User     string `env:"DB_USERNAME"`
	Password string `env:"DB_PASSWORD"`
	Host     string `env:"DB_HOST"`
	Port     string `env:"DB_PORT"`
}

type Store interface {
	Querier
	Close()
}

type SqlStore struct {
	connPool *pgxpool.Pool
	*Queries
}

func NewStore() (Store, error) {

	var cfg config
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("Error parsing environment variables: %s", err)
	}

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)
	dbPool, err := pgxpool.New(context.Background(), connStr)
	if nil != err {
		log.Fatal(err)
	}

	return &SqlStore{
		connPool: dbPool,
		Queries:  New(dbPool),
	}, nil
}

func (store *SqlStore) Close() {
	store.connPool.Close()
}
