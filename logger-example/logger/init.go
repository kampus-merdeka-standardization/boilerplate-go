package logger

import (
	"context"
	"runtime"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

var appName string

func InitLogger(mode, name, path string) {
	var config zap.Config
	if mode == "release" {
		config = zap.NewProductionConfig()
	} else {
		config = zap.NewDevelopmentConfig()
	}
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339Nano)
	config.DisableCaller = true

	config.OutputPaths = []string{path}

	log, err := config.Build()
	if err != nil {
		panic(err)
	}

	logger = log
	appName = name
	defer logger.Sync()
}

func Info(ctx context.Context, methodName, msg string) {
	_, fileName, _, _ := runtime.Caller(0)

	traceID, ok := ctx.Value(TraceID).(string)
	if !ok {
		traceID = uuid.NewString()
	}

	logger.Info(msg,
		zap.String("app_name", appName),
		zap.String("trace_id", traceID),
		zap.String("file_name", fileName),
		zap.String("method_name", methodName),
	)
}

func Debug(ctx context.Context, methodName, msg string) {
	_, fileName, _, _ := runtime.Caller(0)

	traceID, ok := ctx.Value(TraceID).(string)
	if !ok {
		traceID = uuid.NewString()
	}

	logger.Debug(msg,
		zap.String("app_name", appName),
		zap.String("trace_id", traceID),
		zap.String("file_name", fileName),
		zap.String("method_name", methodName),
	)
}

func Warn(ctx context.Context, methodName, msg string) {
	_, fileName, _, _ := runtime.Caller(0)

	traceID, ok := ctx.Value(TraceID).(string)
	if !ok {
		traceID = uuid.NewString()
	}

	logger.Warn(msg,
		zap.String("app_name", appName),
		zap.String("trace_id", traceID),
		zap.String("file_name", fileName),
		zap.String("method_name", methodName),
	)
}

func Error(ctx context.Context, methodName, msg string) {
	_, fileName, _, _ := runtime.Caller(0)

	traceID, ok := ctx.Value(TraceID).(string)
	if !ok {
		traceID = uuid.NewString()
	}

	logger.Error(msg,
		zap.String("app_name", appName),
		zap.String("trace_id", traceID),
		zap.String("file_name", fileName),
		zap.String("method_name", methodName),
	)
}

func Fatal(ctx context.Context, methodName, msg string) {
	_, fileName, _, _ := runtime.Caller(0)

	traceID, ok := ctx.Value(TraceID).(string)
	if !ok {
		traceID = uuid.NewString()
	}

	logger.Fatal(msg,
		zap.String("app_name", appName),
		zap.String("trace_id", traceID),
		zap.String("file_name", fileName),
		zap.String("method_name", methodName),
	)
}
