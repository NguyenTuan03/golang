package routes

import (
	"user-management-api/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, userHandler *handler.UserHandler) {
	api := r.Group("/api/v1")
	{
		users := api.Group("/users")
		{
			users.POST("", userHandler.CreateUser)
			users.GET("", userHandler.GetUsers)
			users.GET("/:uuid", userHandler.GetUser)
			users.PUT("/:uuid", userHandler.UpdateUser)
			users.DELETE("/:uuid", userHandler.DeleteUser)
		}
	}
}