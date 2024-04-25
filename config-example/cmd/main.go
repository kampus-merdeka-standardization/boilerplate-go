package main

import (
	pkg_http "config-example/pkg/http"
	pkg_http_middleware "config-example/pkg/http/middleware"
	pkg_http_wrapper "config-example/pkg/http/wrapper"
	pkg_logger "config-example/pkg/logger"

	"github.com/gin-gonic/gin"
)

func main() {
	pkg_logger.InitLogger(gin.DebugMode, "config-example", "./log/application.log")
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

	srv.Run(":5000")
}
