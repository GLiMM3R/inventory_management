package main

import (
	"fmt"
	"inverntory_management/config"
	"inverntory_management/internal/app"
)

func main() {
	// Initialize the application
	e, err := app.Initialize()
	if err != nil {
		panic(fmt.Sprintf("Failed to initialize application: %v", err))
	}

	port := fmt.Sprintf(":%d", config.AppConfig.PORT)

	e.Logger.Fatal(e.Start(port))
	// ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	// defer stop()
	// // Start server
	// go func() {
	// 	if err := e.Start(port); err != nil && err != http.ErrServerClosed {
	// 		e.Logger.Fatal("shutting down the server")
	// 	}
	// }()

	// // Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// <-ctx.Done()
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()
	// if err := e.Shutdown(ctx); err != nil {
	// 	e.Logger.Fatal(err)
	// }
}
