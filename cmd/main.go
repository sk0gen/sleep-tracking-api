package main

import (
	"context"
	"github.com/sk0gen/sleep-tracking-api/internal/api"
	"github.com/sk0gen/sleep-tracking-api/internal/database/sqlc"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var interruptSignals = []os.Signal{
	os.Interrupt,
	syscall.SIGTERM,
	syscall.SIGINT,
}

func main() {

	if err := realMain(); err != nil {
		log.Printf("ERROR: %v", err)
		os.Exit(1)
	}
}

func realMain() error {
	cfg := api.NewConfig()

	ctx, done := signal.NotifyContext(context.Background(), interruptSignals...)

	defer done()

	db, err := db.NewStore(ctx, cfg.Database)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Start the server
	server := api.NewServer(cfg, db)
	err = server.Start(ctx)
	if err != nil {
		return err
	}
	return nil
}
