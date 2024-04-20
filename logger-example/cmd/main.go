package main

import (
	greeting_handler "logger-example/internal/modules/greeting/handler"
	"logger-example/internal/pkg/configs"
	pkg_http "logger-example/pkg/http"
	pkg_http_middleware "logger-example/pkg/http/middleware"
	pkg_logger "logger-example/pkg/logger"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := configs.LoadConfig()

	pkg_logger.InitLogger(cfg.AppEnv, "logger-app", pkg_http.LOGFILE)
	srv := pkg_http.NewHTTPServer(cfg.AppEnv)

	srv.Use(pkg_http_middleware.LogHandlerMiddleware(), gin.Logger(), gin.Recovery())
	srv.Use(pkg_http_middleware.CorsHandlerMiddleware())
	srv.Use(pkg_http_middleware.ErrorHandlerMiddleware())

	srv.GET("", func(ctx *gin.Context) {
		ctx.JSON(200, "Welcome to Greetings Api")
	})

	greeting_handler.BindGreetingHandler(srv.Group("/greeting"))

	if err := srv.Run(":8080"); err != nil {
		panic(err)
	}
}
