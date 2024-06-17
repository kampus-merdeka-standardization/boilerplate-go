package pkg_otel

import (
	"context"
	"log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdkTrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
)

type OpenTelemetry struct {
	tracer trace.Tracer
	meter  metric.Meter
}

func NewOpenTelemetry(serviceHost *string, serviceName string, serviceEnv, serviceTribe string) (*OpenTelemetry, error) {
	ctx := context.Background()

	conn, err := initConn(*serviceHost)
	if err != nil {
		log.Fatal(err)
	}

	var attributeName = semconv.ServiceNameKey.String(serviceName)

	res, err := resource.New(ctx,
		resource.WithAttributes(
			// The service name used to display traces in backends
			attributeName,
			attribute.String("tribe", serviceTribe),
			attribute.String("env", serviceEnv),
			attribute.String("version", "ver.1"),
			attribute.String("platform", "go"),
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	err = initTracerProvider(ctx, res, conn)
	if err != nil {
		log.Fatal(err)
	}
	err = initMeterProvider(ctx, res, conn)
	if err != nil {
		log.Fatal(err)
	}

	tracer := otel.Tracer(serviceName)
	meter := otel.Meter(serviceName)

	return &OpenTelemetry{
		tracer: tracer,
		meter:  meter,
	}, nil
}

func (o *OpenTelemetry) StartTransaction(ctx context.Context, name string, attributes ...trace.SpanStartOption) (context.Context, interface{}) {
	// Start a new OpenTelemetry span with the given name from a context.
	ctx, span := o.tracer.Start(ctx, name, attributes...)
	return ctx, span
}

func (o *OpenTelemetry) EndTransaction(txn interface{}) {
	// End the given OpenTelemetry span.
	span := txn.(trace.Span)
	span.End()
}

func (o *OpenTelemetry) EndAPM() {
	// shutdown the tracer
	if tp, ok := otel.GetTracerProvider().(*sdkTrace.TracerProvider); ok {
		tp.Shutdown(context.Background())
	}

	// shutdown the meter
	if mp, ok := otel.GetMeterProvider().(*sdkmetric.MeterProvider); ok {
		mp.Shutdown(context.Background())
	}
}
