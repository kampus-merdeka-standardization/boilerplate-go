package hello_handler

import (
	pkg_otel "simple-golang-monitoring/pkg/otel"

	"github.com/gin-gonic/gin"
)

func BindHelloHandler(router *gin.RouterGroup, otelTracer *pkg_otel.OpenTelemetryTracer) {
	// root path : /hello
	controller := &helloController{
		tracer: otelTracer,
	}

	router.GET("/:name", controller.GetHelloByName)
	router.POST("", controller.CreateHello)
	router.PUT("", controller.ReplaceHello)
	router.PATCH("", controller.UpdateHello)
	router.DELETE("/:id", controller.DeleteHello)
}
