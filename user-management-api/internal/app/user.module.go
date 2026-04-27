package app

import (
	"user-management-api/internal/handler"
	"user-management-api/internal/repository"
	"user-management-api/internal/routes"
	"user-management-api/internal/services"

	"gorm.io/gorm"
)

type UserModule struct {
	routes routes.Route
}

func NewUserModule(db *gorm.DB) *UserModule {
	// Initialize layers
	userRepo := repository.NewUserRepository(db)       // repo
	userService := services.NewUserService(userRepo)   // service
	userHandler := handler.NewUserHandler(userService) // handler

	return &UserModule{
		routes: routes.NewUserRouter(userHandler), // router
	}
}

func (m *UserModule) Routes() routes.Route {
	return m.routes
}
