package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/ping/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		price := ctx.Query("price")
		ctx.JSON(200, gin.H{
			"message": "pong",
			"id":      id,
			"price":   price,
		})
	})
	// Router group
	routerGroup := router.Group("/api")
	{
		routerGroup.Group("/v1/user")
		{
			routerGroup.GET("/", func(ctx *gin.Context) {
				ctx.JSON(http.StatusOK, gin.H{
					"message": "success",
				})
			})
		}
	}
	// http://localhost:8080/api/v1/user
	router.Run(":8080")
}
