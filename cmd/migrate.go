/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/sk0gen/sleep-tracking-api/internal/config"
	db "github.com/sk0gen/sleep-tracking-api/internal/database/sqlc"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Apply database migrations",
	RunE:  runMigrate,
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}

func runMigrate(cmd *cobra.Command, args []string) error {
	cfg := config.NewConfig()
	err := db.RunDBMigrationUp(cfg.Database.DSN())
	if err != nil {
		return fmt.Errorf("error running migration: %v", err)
	}
	fmt.Println("Migration successful")
	return nil
}
