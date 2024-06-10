package pkg_otel

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Example using OTLP exporters + collector + third-party backends. For
// information about using the exporter, see:
// https://pkg.go.dev/go.opentelemetry.io/otel/exporters/otlp?tab=doc#example-package-Insecure

import (
	"context"
	"testing"
	"time"
)

func TestOtel(t *testing.T) {
	host := "103.127.137.201:30001"
	otel, err := NewOpenTelemetryTracer(&host, "simple-golang-monitoring", "development", "golang-standarization")
	if err != nil {
		panic(err)
	}
	defer otel.EndAPM()

	ctx := context.Background()
	ctx, spanController := otel.StartTransaction(ctx, "Test Controller")
	time.Sleep(time.Second / 10)

	ctx, spanService := otel.StartTransaction(ctx, "Test Service")
	time.Sleep(time.Second / 50)

	_, spanRepository := otel.StartTransaction(ctx, "Test Repository")
	time.Sleep(time.Millisecond * 300)
	otel.EndTransaction(spanRepository)
	otel.EndTransaction(spanService)
	otel.EndTransaction(spanController)
}
