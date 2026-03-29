package middleware

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func ApiKeyMiddleWare() gin.HandlerFunc {
	expectedKey := os.Getenv("API_KEY")
	if expectedKey == "" {
		log.Fatal("API_KEY is not set")
	}
	return func(ctx *gin.Context) {
		apiKey := ctx.GetHeader("x-api-key")
		log.Println("API Key: ", apiKey)

		if apiKey != expectedKey {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid API key",
			})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
