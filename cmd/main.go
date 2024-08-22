package main

import (
	"github.com/sk0gen/sleep-tracking-api/internal/api"
	"github.com/sk0gen/sleep-tracking-api/internal/database/sqlc"
	"log"
)

func main() {

	cfg := api.NewConfig()

	db, err := db.NewStore(cfg.Database)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// Start the server
	server := api.NewServer(cfg, db)
	err = server.Start()
	if err != nil {
		return
	}
}
