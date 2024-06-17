package hello_handler_test

import (
	hello_controller "simple-golang-monitoring/internal/modules/hello/handler"
	pkg_http "simple-golang-monitoring/pkg/http"
	pkg_otel "simple-golang-monitoring/pkg/otel"

	"github.com/gin-gonic/gin"
)

func setupTest() *gin.Engine {
	srv := pkg_http.NewHTTPServer(gin.TestMode)

	otelAddress := "otel-collector:4317"
	otelTracer, err := pkg_otel.NewOpenTelemetry(&otelAddress, "simple-golang-monitoring", gin.Mode(), "golang-standarization")
	if err != nil {
		panic(err)
	}

	hello_controller.BindHelloHandler(srv.Group("/hello"), otelTracer)
	return srv
}
