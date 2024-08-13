package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

type Database interface {
	GetDb() *sql.DB
	Close() error
}

type database struct {
	db *sql.DB
}

var (
	databaseName = os.Getenv("DB_DATABASE")
	password     = os.Getenv("DB_PASSWORD")
	username     = os.Getenv("DB_USERNAME")
	port         = os.Getenv("DB_PORT")
	host         = os.Getenv("DB_HOST")
	schema       = os.Getenv("DB_SCHEMA")
	dbInstance   *database
)

func Init() Database {
	if dbInstance != nil {
		return dbInstance
	}
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s", username, password, host, port, databaseName, schema)
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	dbInstance = &database{
		db: db,
	}
	return dbInstance
}

func (database *database) GetDb() *sql.DB {
	return database.db
}

func (database *database) Close() error {
	return database.db.Close()
}
