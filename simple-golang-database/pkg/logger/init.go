package pkg_logger

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func InitLogger(mode string, path string) {
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
	defer logger.Sync()
}

func Info(msg string, field ...zap.Field) {
	logger.Info(msg, field...)
}

func Debug(msg string, field ...zap.Field) {
	logger.Debug(msg, field...)
}

func Warn(msg string, field ...zap.Field) {
	logger.Warn(msg, field...)
}

func Error(msg string, field ...zap.Field) {
	logger.Error(msg, field...)
}

func Fatal(msg string, field ...zap.Field) {
	logger.Fatal(msg, field...)
}
