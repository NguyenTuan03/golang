package routes

import (
	"github.com/gin-gonic/gin"
)

type Route interface {
	Register(r *gin.RouterGroup)
}

func SetupRoutes(r *gin.Engine, routes ...Route) {
	v1api := r.Group("/api/v1")

	for _, route := range routes {
		route.Register(v1api)
	}
}
