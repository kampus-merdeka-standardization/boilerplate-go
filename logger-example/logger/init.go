package logger

import (
	"context"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

var logger *zap.Logger

var appName string

func InitLogger(mode, name, path string) {
	var config zapcore.EncoderConfig
	var level zapcore.Level
	if mode == "release" {
		config = zap.NewProductionEncoderConfig()
		level = zap.InfoLevel
	} else {
		config = zap.NewDevelopmentEncoderConfig()
		level = zap.DebugLevel
	}
	config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339Nano)

	writer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   path,
		MaxSize:    500,
		MaxBackups: 10,
		// MaxAge: 30,
	})

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(config),
		writer,
		level,
	)

	logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(level))
	appName = name
	defer logger.Sync()
}

func Info(ctx context.Context, methodName, msg string) {
	traceID, ok := ctx.Value(TraceID).(string)
	if !ok {
		traceID = uuid.NewString()
	}

	logger.Info(msg,
		zap.String("app_name", appName),
		zap.String("trace_id", traceID),
		zap.String("method_name", methodName),
	)
}

func Debug(ctx context.Context, methodName, msg string) {
	traceID, ok := ctx.Value(TraceID).(string)
	if !ok {
		traceID = uuid.NewString()
	}

	logger.Debug(msg,
		zap.String("app_name", appName),
		zap.String("trace_id", traceID),
		zap.String("method_name", methodName),
	)
}

func Warn(ctx context.Context, methodName, msg string) {
	traceID, ok := ctx.Value(TraceID).(string)
	if !ok {
		traceID = uuid.NewString()
	}

	logger.Warn(msg,
		zap.String("app_name", appName),
		zap.String("trace_id", traceID),
		zap.String("method_name", methodName),
	)
}

func Error(ctx context.Context, methodName, msg string) {
	traceID, ok := ctx.Value(TraceID).(string)
	if !ok {
		traceID = uuid.NewString()
	}

	logger.Error(msg,
		zap.String("app_name", appName),
		zap.String("trace_id", traceID),
		zap.String("method_name", methodName),
	)
}

func Fatal(ctx context.Context, methodName, msg string) {
	traceID, ok := ctx.Value(TraceID).(string)
	if !ok {
		traceID = uuid.NewString()
	}

	logger.Fatal(msg,
		zap.String("app_name", appName),
		zap.String("trace_id", traceID),
		zap.String("method_name", methodName),
	)
}
