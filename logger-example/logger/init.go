package logger

import (
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
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)

	config.OutputPaths = []string{path}

	log, err := config.Build()
	if err != nil {
		panic(err)
	}

	logger = log
	appName = name
	defer logger.Sync()
}

func Info(fileName, methodName, msg string) {
	logger.Info(msg,
		zap.String("app_name", appName),
		zap.String("trace_id", uuid.NewString()),
		zap.String("file_name", fileName),
		zap.String("method_name", methodName),
	)
}

func Debug(fileName, methodName, msg string) {
	logger.Debug(msg,
		zap.String("app_name", appName),
		zap.String("trace_id", uuid.NewString()),
		zap.String("file_name", fileName),
		zap.String("method_name", methodName),
	)
}

func Warn(fileName, methodName, msg string) {
	logger.Warn(msg,
		zap.String("app_name", appName),
		zap.String("trace_id", uuid.NewString()),
		zap.String("file_name", fileName),
		zap.String("method_name", methodName),
	)
}

func Error(fileName, methodName, msg string) {
	logger.Error(msg,
		zap.String("app_name", appName),
		zap.String("trace_id", uuid.NewString()),
		zap.String("file_name", fileName),
		zap.String("method_name", methodName),
	)
}

func Fatal(fileName, methodName, msg string) {
	logger.Fatal(msg,
		zap.String("app_name", appName),
		zap.String("trace_id", uuid.NewString()),
		zap.String("file_name", fileName),
		zap.String("method_name", methodName),
	)
}
