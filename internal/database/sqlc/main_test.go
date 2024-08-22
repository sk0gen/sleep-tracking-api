package db

import (
	"github.com/caarlos0/env/v11"
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

	var cfg Config
	env.Parse(&cfg)

	testStore, err = NewStore(cfg)
	if err != nil {
		return
	}

	os.Exit(m.Run())
}
