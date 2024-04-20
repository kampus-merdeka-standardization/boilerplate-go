package main

import (
	pkg_http "logger-example/pkg/http"
	pkg_http_middleware "logger-example/pkg/http/middleware"
	pkg_logger "logger-example/pkg/logger"

	"github.com/gin-gonic/gin"
)

func main() {
	pkg_logger.InitLogger(gin.DebugMode, "logger-app", pkg_http.LOGFILE)
	srv := pkg_http.NewHTTPServer(gin.DebugMode)

	srv.Use(pkg_http_middleware.LogHandlerMiddleware(), gin.Logger(), gin.Recovery())
	srv.Use(pkg_http_middleware.CorsHandlerMiddleware())
	srv.Use(pkg_http_middleware.ErrorHandlerMiddleware())

	srv.GET("", func(ctx *gin.Context) {
		ctx.JSON(200, pkg_http.Response{
			Message: "Successfully Connected to Logger Example API",
		})
	})

	if err := srv.Run(":8080"); err != nil {
		panic(err)
	}
}
