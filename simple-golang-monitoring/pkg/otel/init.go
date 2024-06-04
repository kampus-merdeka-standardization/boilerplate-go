package pkg_otel

import (
	"context"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdkTrace "go.opentelemetry.io/otel/sdk/trace"
	semConv "go.opentelemetry.io/otel/semconv/v1.12.0"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type OpenTelemetryTracer struct {
	tracer trace.Tracer
}

func NewOpenTelemetryTracer(serviceHost *string, serviceName string, serviceEnv, serviceTribe string) (*OpenTelemetryTracer, error) {
	var (
		err  error
		conn *grpc.ClientConn
	)
	conn, err = grpc.DialContext(context.Background(), *serviceHost,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()))
	if err != nil {
		return nil, err
	}

	// Create a new OTLP exporter over gRPC
	exp, err := otlptracegrpc.New(context.Background(),
		otlptracegrpc.WithGRPCConn(conn))
	if err != nil {
		return nil, err
	}

	traceExporter := sdkTrace.SpanExporter(exp)
	// Create a new trace provider with the exporter.
	tp := sdkTrace.NewTracerProvider(
		sdkTrace.WithBatcher(traceExporter),
		sdkTrace.WithResource(resource.NewWithAttributes(
			semConv.SchemaURL,
			semConv.ServiceNameKey.String(serviceName),
			attribute.String("tribe", serviceTribe),
			attribute.String("env", serviceEnv),
			attribute.String("version", "ver.1"),
			attribute.String("platform", "go"),
		)),
	)

	// Set the global trace provider and the propagation.
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			propagation.Baggage{}, propagation.TraceContext{}))

	tracer := otel.GetTracerProvider().Tracer(serviceName)

	return &OpenTelemetryTracer{tracer: tracer}, nil
}

func (o *OpenTelemetryTracer) StartTransaction(ctx context.Context, name string) (context.Context, interface{}) {
	// Start a new OpenTelemetry span with the given name from a context.
	ctx, span := o.tracer.Start(ctx, name)
	return ctx, span
}

func (o *OpenTelemetryTracer) EndTransaction(txn interface{}) {
	// End the given OpenTelemetry span.
	span := txn.(trace.Span)
	span.End()
}

func (o *OpenTelemetryTracer) EndAPM() {
	// shutdown the tracer
	if tp, ok := otel.GetTracerProvider().(*sdkTrace.TracerProvider); ok {
		tp.Shutdown(context.Background())
	}
}
