/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"github.com/sk0gen/sleep-tracking-api/internal/api"
	"github.com/sk0gen/sleep-tracking-api/internal/config"
	db "github.com/sk0gen/sleep-tracking-api/internal/database/sqlc"
	"github.com/sk0gen/sleep-tracking-api/internal/gapi"
	"github.com/sk0gen/sleep-tracking-api/internal/logging"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
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
	cfg := config.NewConfig()
	logger, err := logging.InitZap(cfg.LogConfig)
	defer logger.Sync()

	ctx, done := signal.NotifyContext(context.Background(), interruptSignals...)

	defer done()

	db, err := db.NewStore(ctx, cfg.Database)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	waitGroup, ctx := errgroup.WithContext(ctx)
	startAPIServer(ctx, waitGroup, cfg, db, logger)
	startGRPCServer(ctx, waitGroup, cfg, db, logger)

	return waitGroup.Wait()
}

func startAPIServer(ctx context.Context, waitGroup *errgroup.Group, cfg config.Config, db db.Store, logger *zap.Logger) {
	server := api.NewServer(cfg, db, logger)
	waitGroup.Go(func() error {
		return server.Start(ctx)
	})
}

func startGRPCServer(ctx context.Context, waitGroup *errgroup.Group, cfg config.Config, db db.Store, logger *zap.Logger) {
	server := gapi.NewServer(cfg, db)
	waitGroup.Go(func() error {
		return server.Start(ctx)
	})
}
