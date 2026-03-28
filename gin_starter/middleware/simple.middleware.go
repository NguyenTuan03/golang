package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

func SimpleMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Println("Before request")

		ctx.Next()

		log.Println("After request")
	}
}
