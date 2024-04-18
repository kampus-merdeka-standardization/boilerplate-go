package main

import (
	"context"
	"logger-example/logger"

	"github.com/google/uuid"
)

func main() {
	logger.InitLogger("release", "example-app", "./log/example.log")

	ctx := context.Background()

	logger.Info(ctx, "create_product", "can't connect to db")
	logger.Debug(ctx, "create_product", "can't connect to db")
	logger.Warn(ctx, "create_product", "can't connect to db")
	logger.Error(ctx, "create_product", "can't connect to db")
	// logger.Fatal(ctx, "create_product", "can't connect to db")

	ctx = context.WithValue(ctx, logger.TraceID, uuid.NewString())
	logger.Info(ctx, "create_product", "failed to create product")
	logger.Debug(ctx, "create_product", "failed to create product")
	logger.Warn(ctx, "create_product", "failed to create product")
	logger.Error(ctx, "create_product", "failed to create product")
	// logger.Fatal(ctx, "create_product", "failed to create product")
}
