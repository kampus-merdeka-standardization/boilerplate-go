package pkg_http_middleware

import (
	pkg_logger "simple-golang-monitoring/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func LogHandler(logPath string) gin.HandlerFunc {
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

		pkg_logger.InitLogger(gin.Mode(), logPath)

		if param.StatusCode >= 500 {
			pkg_logger.Error(
				"Internal Server Error",
				zap.String("client_id", param.ClientIP),
				zap.String("method", param.Method),
				zap.Int("body_size", param.BodySize),
				zap.String("path", path),
				zap.String("latency", param.Latency.String()),
				zap.String("error", param.ErrorMessage),
			)
		} else {
			pkg_logger.Debug(
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
