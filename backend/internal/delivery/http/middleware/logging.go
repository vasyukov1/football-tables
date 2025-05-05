package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"time"
)

func Logging() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		latency := time.Since(start)
		log.Info().
			Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Int("status", c.Writer.Status()).
			Dur("latency", latency).
			Msg("request")
	}
}
