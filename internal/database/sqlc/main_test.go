package db

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"
)

var testQueries *Queries

func TestMain(m *testing.M) {

	err := godotenv.Load("../../../.env.test")
	if err != nil {
		log.Fatalf("Error loading .env.test file: %s", err)
	}

	db, err := NewStore()
	if err != nil {
		return
	}

	testQueries = db.Queries
	os.Exit(m.Run())
}
