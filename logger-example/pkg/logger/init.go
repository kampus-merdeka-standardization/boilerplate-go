package pkg_logger

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
		MaxSize:    200,
		MaxBackups: 10,
		// MaxAge: 30,
	})

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(config),
		writer,
		level,
	)

	logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	appName = name
	defer logger.Sync()
}

func Info(ctx context.Context, methodName, msg string, fields ...zap.Field) {
	traceID, ok := ctx.Value(TraceID).(string)
	if !ok {
		traceID = uuid.NewString()
	}

	fields = append(
		fields,
		zap.String("app_name", appName),
		zap.String("trace_id", traceID),
		zap.String("method_name", methodName),
	)

	logger.Info(msg,
		fields...,
	)
}

func Debug(ctx context.Context, methodName, msg string, fields ...zap.Field) {
	traceID, ok := ctx.Value(TraceID).(string)
	if !ok {
		traceID = uuid.NewString()
	}

	fields = append(
		fields,
		zap.String("app_name", appName),
		zap.String("trace_id", traceID),
		zap.String("method_name", methodName),
	)

	logger.Debug(msg,
		fields...,
	)
}

func Warn(ctx context.Context, methodName, msg string, fields ...zap.Field) {
	traceID, ok := ctx.Value(TraceID).(string)
	if !ok {
		traceID = uuid.NewString()
	}

	fields = append(
		fields,
		zap.String("app_name", appName),
		zap.String("trace_id", traceID),
		zap.String("method_name", methodName),
	)

	logger.Warn(msg, fields...)
}

func Error(ctx context.Context, methodName, msg string, fields ...zap.Field) {
	traceID, ok := ctx.Value(TraceID).(string)
	if !ok {
		traceID = uuid.NewString()
	}

	fields = append(
		fields,
		zap.String("app_name", appName),
		zap.String("trace_id", traceID),
		zap.String("method_name", methodName),
	)

	logger.Error(msg,
		fields...,
	)
}

func Fatal(ctx context.Context, methodName, msg string, fields ...zap.Field) {
	traceID, ok := ctx.Value(TraceID).(string)
	if !ok {
		traceID = uuid.NewString()
	}

	fields = append(
		fields,
		zap.String("app_name", appName),
		zap.String("trace_id", traceID),
		zap.String("method_name", methodName),
	)

	logger.Fatal(msg,
		fields...,
	)

	//	func traceMethodName() string {
	//		pc := make([]uintptr, 1) // at least 1 entry needed
	//		runtime.Callers(4, pc)
	//		m := runtime.FuncForPC(pc[0])
	//		return m.Name()
	//	}
}
