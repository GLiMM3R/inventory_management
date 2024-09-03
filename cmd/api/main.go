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
}
