package main

import (
	"github.com/sk0gen/sleep-tracking-api/internal/api"
	"github.com/sk0gen/sleep-tracking-api/internal/config"
	"github.com/sk0gen/sleep-tracking-api/internal/database"
	"log"
)

func main() {
	config.Init()

	db, err := database.New()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// Start the server
	api.Run()
}
