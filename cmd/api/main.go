package main

import (
	"context"
	"fmt"
	"inverntory_management/config"
	"inverntory_management/internal/app"
	"inverntory_management/internal/database/migration"
	"os"
	"os/signal"

	"github.com/spf13/cobra"
)

func main() {
	// Initialize the application
	cfg := config.LoadConfig(".", ".env")

	cmd := &cobra.Command{
		Short: "Start the inventory management application",
	}

	start := &cobra.Command{
		Use:   "start",
		Short: "Start the application",
		Run: func(cmd *cobra.Command, args []string) {
			app := app.New(cfg)

			ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
			defer cancel()

			app.Start(ctx)
		},
	}

	database_migration := &cobra.Command{
		Use:   "migrate",
		Short: "Run Database Migration",
		Run: func(cmd *cobra.Command, args []string) {
			migrate := migration.New(cfg)
			migrate.Run()
		},
	}

	cmd.AddCommand(start, database_migration)

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
