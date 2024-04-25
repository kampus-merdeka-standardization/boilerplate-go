package main

import (
	internal_configs "config-example/internal/pkg/configs"
	pkg_http "config-example/pkg/http"
	pkg_http_middleware "config-example/pkg/http/middleware"
	pkg_http_wrapper "config-example/pkg/http/wrapper"
	pkg_logger "config-example/pkg/logger"
	"context"

	"github.com/gin-gonic/gin"
)

func main() {
	config := internal_configs.LoadConfig()
	pkg_logger.InitLogger(config.AppEnv, config.AppName, config.LogPath)
	srv := pkg_http.NewHTTPServer(gin.DebugMode)

	srv.Use(
		pkg_http_middleware.LogHandlerMiddleware(),
		gin.Logger(),
		gin.Recovery(),
		pkg_http_middleware.TraceIdAssignmentMiddleware(),
		pkg_http_middleware.CorsHandlerMiddleware(),
		pkg_http_middleware.ErrorHandlerMiddleware(),
	)

	srv.GET("", func(ctx *gin.Context) {
		ctx.JSON(200, pkg_http_wrapper.NewResponse("Success Request"))
	})

	if err := srv.Run(":" + config.AppPort); err != nil {
		pkg_logger.Fatal(context.Background(), err.Error())
	}
}
