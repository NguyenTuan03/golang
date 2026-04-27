package app

import (
	"log"

	"user-management-api/internal/config"
	"user-management-api/internal/database"
	"user-management-api/internal/models"
	"user-management-api/internal/routes"

	"github.com/gin-gonic/gin"
)

type Module interface {
	Routes() routes.Route
}

type Application struct {
	cfg    *config.Config
	engine *gin.Engine
	module []Module
}

func NewApplication(cfg *config.Config) (*Application, error) {
	// Initialize database
	db, err := database.InitDB(cfg.DB.DSN)
	if err != nil {
		return nil, err
	}

	// Auto Migrate
	log.Println("Running auto migration...")
	if err := db.AutoMigrate(&models.User{}); err != nil {
		return nil, err
	}

	// Setup Gin
	engine := gin.Default()

	// Setup routes
	modules := []Module{
		NewUserModule(db),
	}
	routes.SetupRoutes(engine, getModuleRoutes(modules)...)

	return &Application{
		cfg:    cfg,
		engine: engine,
	}, nil
}

func (a *Application) Run() error {
	port := a.cfg.App.Port
	log.Printf("Server starting on port %s", port)
	return a.engine.Run(port)
}

func getModuleRoutes(module []Module) []routes.Route {
	routesList := make([]routes.Route, len(module))

	for i, mod := range module {
		routesList[i] = mod.Routes()
	}

	return routesList
}
