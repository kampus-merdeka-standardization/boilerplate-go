package main

import (
	"fmt"
	hello_handler "simple-golang-monitoring/internal/modules/hello/handler"
	internal_configs "simple-golang-monitoring/internal/pkg/configs"
	pkg_http "simple-golang-monitoring/pkg/http"
	pkg_otel "simple-golang-monitoring/pkg/otel"
	pkg_otel_gin_metrics "simple-golang-monitoring/pkg/otel/otelginmetrics"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func main() {
	confLoader := internal_configs.NewConfigLoader()
	configs := confLoader.LoadConfig()

	app := pkg_http.NewHTTPServer(configs.AppEnv)

	otel, err := pkg_otel.NewOpenTelemetry(&configs.OtelAddress, configs.AppName, configs.AppEnv, "golang-standarization")
	if err != nil {
		panic(err)
	}

	app.Use(gin.Logger(), gin.Recovery(), otelgin.Middleware(configs.AppName), pkg_otel_gin_metrics.Middleware(configs.AppName))

	helloGroup := app.Group("/hello")
	hello_handler.BindHelloHandler(helloGroup, otel)

	if err := app.Run(fmt.Sprintf(":%s", configs.AppPort)); err != nil {
		otel.EndAPM()
		panic(err)
	}
	otel.EndAPM()
}
