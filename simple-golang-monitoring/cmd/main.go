package main

import (
	hello_handler "simple-golang-monitoring/internal/modules/hello/handler"
	pkg_http "simple-golang-monitoring/pkg/http"
	pkg_otel "simple-golang-monitoring/pkg/otel"

	"github.com/gin-gonic/gin"
)

func main() {
	app := pkg_http.NewHTTPServer(gin.DebugMode)

	otelAddress := "otel-collector:4317"
	otelTracer, err := pkg_otel.NewOpenTelemetryTracer(&otelAddress, "simple-golang-monitoring", gin.Mode(), "golang-standarization")
	if err != nil {
		panic(err)
	}

	helloGroup := app.Group("/hello")
	hello_handler.BindHelloHandler(helloGroup, otelTracer)

	if err := app.Run(":8080"); err != nil {
		panic(err)
	}
}
