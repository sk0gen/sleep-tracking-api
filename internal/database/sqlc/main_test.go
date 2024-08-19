package db

import (
	"github.com/joho/godotenv"
	"github.com/sk0gen/sleep-tracking-api/internal/database"
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

	pool, err := database.New()
	if err != nil {
		return
	}

	testQueries = New(pool)
	os.Exit(m.Run())
}
