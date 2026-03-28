package main

import (
	"gin/internal/api/v1/handler"
	handler2 "gin/internal/api/v2/handler"
	"gin/middleware"
	"net/http"

	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func getPing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
func main() {
	// Redirect standard log to stdout to align with Gin's output
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	go middleware.RemoveExpiredClients()

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
			product := v1.Group("/products").Use(middleware.ApiKeyMiddleWare(), middleware.RateLimitingMiddleware())
			{
				productHandlerV1 := handler.NewProductHandler()
				product.GET("", productHandlerV1.GetListProducts)
				product.GET("/:id", productHandlerV1.GetProductById)
				product.GET("/admin/:slug", productHandlerV1.GetProductionBySlug)
				product.GET("/categories/:category", productHandlerV1.GetProductsByCategory)
				product.POST("", productHandlerV1.CreateNewProduct)
				product.PUT("/:id", productHandlerV1.UpdateProductById)
				product.DELETE("/:id", productHandlerV1.DeleteProductById)
			}
		}
		v2 := r.Group("/v2")
		{
			userv2 := v2.Group("/users")
			{
				userHandlerV2 := handler2.NewUserHandler()
				userv2.GET("", userHandlerV2.GetListUsers)
				userv2.GET("/:id", userHandlerV2.GetUserById)
				userv2.GET("/admin/:uuid", userHandlerV2.GetUserByUuid)
				userv2.POST("", userHandlerV2.CreateNewUser)
				userv2.PUT("/:id", userHandlerV2.UpdateUserById)
				userv2.DELETE("/:id", userHandlerV2.DeleteUserById)
			}
			product := v2.Group("/products")
			{
				productHandlerV2 := handler2.NewProductHandler()
				product.GET("/:id", productHandlerV2.GetProductById)

			}
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8386"
	}
	r.Run(":" + port)
}
