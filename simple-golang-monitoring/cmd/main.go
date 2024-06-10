package main

import (
	hello_handler "simple-golang-monitoring/internal/modules/hello/handler"
	internal_configs "simple-golang-monitoring/internal/pkg/configs"
	pkg_http "simple-golang-monitoring/pkg/http"
	pkg_otel "simple-golang-monitoring/pkg/otel"

	"github.com/gin-gonic/gin"
)

func main() {
	confLoader := internal_configs.NewConfigLoader()
	configs := confLoader.LoadConfig()

	app := pkg_http.NewHTTPServer(configs.AppEnv)

	otel, err := pkg_otel.NewOpenTelemetryTracer(&configs.OtelAddress, configs.AppName, configs.AppEnv, "golang-standarization")
	if err != nil {
		panic(err)
	}

	app.Use(gin.Logger(), gin.Recovery())

	helloGroup := app.Group("/hello")
	hello_handler.BindHelloHandler(helloGroup, otel)

	if err := app.Run(":8080"); err != nil {
		otel.EndAPM()
		panic(err)
	}
	otel.EndAPM()
}
