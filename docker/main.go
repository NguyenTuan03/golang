package main

import (
	"docker/internal/database"
	"docker/internal/handlers"
	"docker/internal/repository"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Initialize database
	err = database.InitDB()
	if err != nil {
		log.Fatalf("Error loading database: %v", err)
	}

	r := gin.Default()

	userRepository := repository.NewSQLUserRepository(database.DB)
	userHandler := handlers.NewUserHandler(userRepository)

	api := r.Group("/api")
	{
		//version 1
		v1 := api.Group("/v1")
		{
			v1.GET("/users", userHandler.GetListUsers)
			v1.GET("/users/:id", userHandler.GetUserById)
			v1.POST("/users", userHandler.CreateUser)
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
