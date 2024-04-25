package pkg_http_middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CorsHandlerMiddleware(origin ...string) gin.HandlerFunc {
	allowAllOrigin := false
	if len(origin) == 0 {
		allowAllOrigin = true
	}
	return cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization", "Content-Disposition"},
		ExposeHeaders:    []string{"Content-Disposition"},
		AllowCredentials: true,
		AllowAllOrigins:  allowAllOrigin,
		AllowOrigins:     origin,
		MaxAge:           12 * time.Hour,
	})
}
