package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

type Config struct {
	Name     string `env:"DB_DATABASE"`
	User     string `env:"DB_USERNAME"`
	Password string `env:"DB_PASSWORD"`
	Host     string `env:"DB_HOST"`
	Port     string `env:"DB_PORT"`
}

func (cfg Config) DSN() string {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)
	return connStr
}

type Store interface {
	Querier
	Close()
}

type SqlStore struct {
	connPool *pgxpool.Pool
	*Queries
}

func NewStore(cfg Config) (Store, error) {
	dbPool, err := pgxpool.New(context.Background(), cfg.DSN())
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
