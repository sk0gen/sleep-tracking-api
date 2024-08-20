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

	db, err := database.NewDatabase()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Pool.Close()

	// Start the server
	server := api.NewServer(db)
	err = server.Start()
	if err != nil {
		return
	}
}
