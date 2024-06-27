package main

import (
	greeting_handler "simple-golang-log/internal/modules/greeting/handler"
	"simple-golang-log/internal/pkg/configs"
	pkg_http "simple-golang-log/pkg/http"
	pkg_http_middleware "simple-golang-log/pkg/http/middleware"
	pkg_logger "simple-golang-log/pkg/logger"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := configs.LoadConfig()

	pkg_logger.InitLogger(cfg.AppEnv, "logger-app", pkg_http.LOGFILE)
	srv := pkg_http.NewHTTPServer(cfg.AppEnv)

	srv.Use(
		pkg_http_middleware.TraceIdAssignmentMiddleware(),
		pkg_http_middleware.LogHandlerMiddleware(),
		gin.Logger(),
		gin.Recovery(),
		pkg_http_middleware.CorsHandlerMiddleware(),
		pkg_http_middleware.ErrorHandlerMiddleware(),
	)

	srv.GET("", func(ctx *gin.Context) {
		ctx.JSON(200, "Welcome to Greetings Api")
	})

	greeting_handler.BindGreetingHandler(srv.Group("/greeting"))

	if err := srv.Run(":8080"); err != nil {
		panic(err)
	}
}
