package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var (
	databaseName = os.Getenv("DB_DATABASE")
	password     = os.Getenv("DB_PASSWORD")
	username     = os.Getenv("DB_USERNAME")
	port         = os.Getenv("DB_PORT")
	host         = os.Getenv("DB_HOST")
	schema       = os.Getenv("DB_SCHEMA")
)

func New() (*sql.DB, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s", username, password, host, port, databaseName, schema)
	db, err := sql.Open("pgx", connStr)
	if nil != err {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
