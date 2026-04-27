package routes

import (
	"user-management-api/internal/handler"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	userHandler *handler.UserHandler
}

func NewUserRouter(userHandler *handler.UserHandler) *UserRouter {
	return &UserRouter{
		userHandler: userHandler,
	}
}

func (r *UserRouter) Register(v1api *gin.RouterGroup) {
	users := v1api.Group("/users")
	{
		users.POST("", r.userHandler.CreateUser)
		users.GET("", r.userHandler.GetUsers)
		users.GET("/:uuid", r.userHandler.GetUser)
		users.PUT("/:uuid", r.userHandler.UpdateUser)
		users.DELETE("/:uuid", r.userHandler.DeleteUser)
	}
}