package main

import (
	"github.com/sk0gen/sleep-tracking-api/internal/api"
	"github.com/sk0gen/sleep-tracking-api/internal/database"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := database.New()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// Start the server
	api.Run()
}
