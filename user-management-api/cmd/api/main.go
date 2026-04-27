package main

import (
	"log"
	"user-management-api/internal/app"
	"user-management-api/internal/config"
)

func main() {
	cfg := config.LoadConfig()

	application, err := app.NewApplication(cfg)
	if err != nil {
		log.Fatalf("Could not initialize application: %v", err)
	}

	if err := application.Run(); err != nil {
		log.Fatalf("Could not run application: %v", err)
	}
}
