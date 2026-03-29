package middleware

import (
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func LoggerMiddleware() gin.HandlerFunc {
	logPath := "logs/http.log"
	if err := os.MkdirAll(filepath.Dir(logPath), os.ModePerm); err != nil {
		panic(err)
	}

	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}

	logger := zerolog.New(file).With().Timestamp().Logger()

	return func(ctx *gin.Context) {
		logEvent := logger.Info()
		statusCode := ctx.Writer.Status()
		if statusCode >= 500 {
			logEvent = logger.Error()
		} else if statusCode >= 400 {
			logEvent = logger.Warn()
		}

		logEvent.
			Str("method", ctx.Request.Method).
			Str("path", ctx.Request.URL.Path).
			Str("ip", ctx.ClientIP()).
			Str("host", ctx.Request.Host).
			Str("user_agent", ctx.Request.UserAgent()).
			Int("status", ctx.Writer.Status()).
			Dur("latency", time.Since(time.Now())).
			Send()

		ctx.Next()
	}
}
