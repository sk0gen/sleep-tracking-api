package db

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"
)

var testStore Store

func TestMain(m *testing.M) {

	err := godotenv.Load("../../../.env.test")
	if err != nil {
		log.Fatalf("Error loading .env.test file: %s", err)
	}

	testStore, err = NewStore()
	if err != nil {
		return
	}

	os.Exit(m.Run())
}
