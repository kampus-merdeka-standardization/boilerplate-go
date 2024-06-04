package hello_handler

import pkg_otel "simple-golang-monitoring/pkg/otel"

type helloController struct {
	tracer *pkg_otel.OpenTelemetryTracer
}
