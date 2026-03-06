package main

import (
	"gin/internal/api/v1/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getPing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
func main() {
	r := gin.Default()
	{
		v1 := r.Group("/v1")
		{
			user := v1.Group("/users")
			{
				userHandlerV1 := handler.NewUserHandler()
				user.GET("", userHandlerV1.GetListUsers)
				user.GET("/:id", userHandlerV1.GetUserById)
				user.GET("/admin/:uuid", userHandlerV1.GetUserByUuid)
				user.POST("", userHandlerV1.CreateNewUser)
				user.PUT("/:id", userHandlerV1.UpdateUserById)
				user.DELETE("/:id", userHandlerV1.DeleteUserById)
			}
			product := v1.Group("/products")
			{
				productHandlerV1 := handler.NewProductHandler()
				product.GET("", productHandlerV1.GetListProducts)
				product.GET("/:id", productHandlerV1.GetProductById)
				product.POST("", productHandlerV1.CreateNewProduct)
				product.PUT("/:id", productHandlerV1.UpdateProductById)
				product.DELETE("/:id", productHandlerV1.DeleteProductById)
			}
		}
	}

	r.Run(":8386")
}
