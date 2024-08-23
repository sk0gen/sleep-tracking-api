/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"github.com/sk0gen/sleep-tracking-api/internal/api"
	db "github.com/sk0gen/sleep-tracking-api/internal/database/sqlc"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the API server",
	Long:  `Starts the API server`,
	RunE:  runServe,
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

var interruptSignals = []os.Signal{
	os.Interrupt,
	syscall.SIGTERM,
	syscall.SIGINT,
}

func runServe(cmd *cobra.Command, args []string) error {
	if err := realMain(); err != nil {
		log.Printf("ERROR: %v", err)
		os.Exit(1)
	}
	return nil
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
