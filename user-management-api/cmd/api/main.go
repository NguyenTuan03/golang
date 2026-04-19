package main

import (
	"log"
	"user-management-api/internal/config"
	"user-management-api/internal/database"
	"user-management-api/internal/handler"
	"user-management-api/internal/models"
	"user-management-api/internal/repository"
	"user-management-api/internal/routes"
	"user-management-api/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Load config
	cfg := config.LoadConfig()

	// 2. Initialize database
	db, err := database.InitDB(cfg.DB.DSN)
	if err != nil {
		log.Fatalf("Could not initialize database: %v", err)
	}

	// 3. Auto Migrate
	log.Println("Running auto migration...")
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Auto migration failed: %v", err)
	}

	// 4. Initialize layers
	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// 5. Setup Gin
	r := gin.Default()

	// 6. Setup routes
	routes.SetupRoutes(r, userHandler)

	// 7. Start server
	port := cfg.App.Port
	log.Printf("Server starting on port %s", port)
	if err := r.Run(port); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
