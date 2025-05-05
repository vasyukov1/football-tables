package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/vasyukov1/football-tables/backend/internal/config"
	"strings"
)

func CORS(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", strings.Join(cfg.CORS.AllowedOrigins, ","))
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers",
			"Content-Type, Content-Length, Accept-Encoding, Authorization, Accept, Origin, Cache-Control")
		c.Writer.Header().Set("Access-Control-Allow-Methods",
			"POST, OPTIONS, GET, PUT, PATCH, DELETE")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
