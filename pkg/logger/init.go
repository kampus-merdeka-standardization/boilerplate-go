package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func InitLogger(mode string) {
	var config zap.Config
	if mode == "release" {
		config = zap.NewProductionConfig()
	} else {
		config = zap.NewDevelopmentConfig()
	}
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, err := config.Build()
	if err != nil {
		panic(err)
	}
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
