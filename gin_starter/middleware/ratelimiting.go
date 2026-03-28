package middleware

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// 127.0.0.1 => {limiter, sender}
type Client struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

var (
	clients = make(map[string]*Client)
	mx      sync.Mutex
)

func GetClientIP(ctx *gin.Context) string {
	ip := ctx.ClientIP()
	if ip == "" {
		ip = ctx.Request.RemoteAddr
	}
	return ip
}

func GetRateLimiter(ip string) *rate.Limiter {
	mx.Lock()
	defer mx.Unlock()

	client, exist := clients[ip]
	if !exist {
		newClient := &Client{
			limiter:  rate.NewLimiter(10, 10),
			lastSeen: time.Now(),
		}
		clients[ip] = newClient
	}

	client.lastSeen = time.Now()

	return client.limiter
}

func RemoveExpiredClients() {
	for {
		time.Sleep(time.Minute)
		mx.Lock()
		for ip, client := range clients {
			if time.Since(client.lastSeen) > 3*time.Minute {
				delete(clients, ip)
			}
		}
		mx.Unlock()
	}
}

func RateLimitingMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ip := GetClientIP(ctx)
		limiter := GetRateLimiter(ip)

		if !limiter.Allow() {
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"code":    http.StatusTooManyRequests,
				"message": "Too many requests",
			})
			return
		}
		log.Println("IP: ", ip)
		ctx.Next()
	}
}
