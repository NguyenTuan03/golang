package main

import "github.com/gin-gonic/gin"

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
	userGroup := router.Group("/user")
	{
		userGroup.Group("/v1")
		{
			userGroup.GET("/", func(ctx *gin.Context) {
				ctx.JSON(200, gin.H{
					"message": "pong",
				})
			})
		}
	}

	router.Run(":8080")
}
