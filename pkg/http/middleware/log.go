package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func LogHandler(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now() // Start timer
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Fill the params
		param := gin.LogFormatterParams{}

		param.TimeStamp = time.Now() // Stop timer
		param.Latency = param.TimeStamp.Sub(start)
		if param.Latency > time.Minute {
			param.Latency = param.Latency.Truncate(time.Second)
		}

		param.ClientIP = c.ClientIP()
		param.Method = c.Request.Method
		param.StatusCode = c.Writer.Status()
		param.BodySize = c.Writer.Size()
		if raw != "" {
			path = path + "?" + raw
		}
		param.Path = path

		if param.StatusCode >= 500 {
			logger.Warn(
				"Internal Server Error",
				zap.String("client_id", param.ClientIP),
				zap.String("method", param.Method),
				zap.Int("body_size", param.BodySize),
				zap.String("path", path),
				zap.String("latency", param.Latency.String()),
				zap.String("error", param.ErrorMessage),
			)
		} else {
			logger.Info(
				"Request",
				zap.String("client_id", param.ClientIP),
				zap.String("method", param.Method),
				zap.Int("body_size", param.BodySize),
				zap.String("path", path),
				zap.String("latency", param.Latency.String()),
			)
		}
	}
}
