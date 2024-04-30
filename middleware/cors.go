package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mikezzb/steam-trading-server/pkg/setting"
)

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		// check if origin is allowed
		if setting.Server.AllowedOrigins[origin] {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Type, Content-Length, X-CSRF-Token")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Set("Content-Type", "application/json")

			if c.Request.Method == "OPTIONS" {
				c.AbortWithStatus(http.StatusOK)
				return
			}
		}

		c.Next()
	}
}
