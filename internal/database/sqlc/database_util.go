package db

import (
	"context"
	"errors"
	"fmt"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"log"
	"path/filepath"
	"runtime"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // used by migrator
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type TestDatabase struct {
	container testcontainers.Container
	Config    Config
}

func NewTestDatabase() *TestDatabase {
	ctx := context.Background()

	dbName := "local"
	dbUser := "user"
	dbPassword := "password"

	postgresContainer, err := postgres.Run(ctx,
		"postgres:16.4-alpine",
		postgres.WithDatabase(dbName),
		postgres.WithUsername(dbUser),
		postgres.WithPassword(dbPassword),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		log.Fatalf("failed to start container: %s", err)
	}

	port, err := postgresContainer.MappedPort(ctx, "5432")
	if err != nil {
		log.Fatalf("failed to get mapped port: %s", err)
	}

	host, err := postgresContainer.Host(ctx)
	if err != nil {
		log.Fatalf("failed to get host: %s", err)
	}

	config := Config{
		Name:     dbName,
		User:     dbUser,
		Password: dbPassword,
		Host:     host,
		Port:     port.Port(),
	}

	RunDBMigration(config.DSN())

	return &TestDatabase{
		container: postgresContainer,
		Config:    config,
	}

}

func (td *TestDatabase) Close() {
	td.container.Terminate(context.Background())
}

func RunDBMigration(dbSource string) error {
	migrationsUrl := fmt.Sprintf("file://%s", dbMigrationsDir())

	m, err := migrate.New(migrationsUrl, dbSource)
	if err != nil {
		return fmt.Errorf("failed create migrate: %w", err)
	}
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("failed run migrate: %w", err)
	}
	srcErr, dbErr := m.Close()
	if srcErr != nil {
		return fmt.Errorf("migrate source error: %w", srcErr)
	}
	if dbErr != nil {
		return fmt.Errorf("migrate database error: %w", dbErr)
	}
	return nil
}

// dbMigrationsDir returns the path on disk to the migrations. It uses
// runtime.Caller() to get the path to the caller, since this package is
// imported by multiple others at different levels.
func dbMigrationsDir() string {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		return ""
	}
	return filepath.Join(filepath.Dir(filename), "../migrations")
}
