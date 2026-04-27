package routes

import (
	"user-management-api/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.RouterGroup, userHandler *handler.UserHandler) {
	users := r.Group("/users")
	{
		users.POST("", userHandler.CreateUser)
		users.GET("", userHandler.GetUsers)
		users.GET("/:uuid", userHandler.GetUser)
		users.PUT("/:uuid", userHandler.UpdateUser)
		users.DELETE("/:uuid", userHandler.DeleteUser)
	}
}