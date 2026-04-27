package routes

import (
	"user-management-api/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, userHandler *handler.UserHandler) {
	api := r.Group("/api/v1")
	{
		SetupUserRoutes(api, userHandler)
	}
}