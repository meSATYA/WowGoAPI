package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	var err error

	config := zap.NewProductionConfig()
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig = encoderConfig

	log, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}

func Info(message string, fields ...zap.Field) *zap.Logger {
	log.Info(message, fields...)
	return log
}

func Debug(message string, fields ...zap.Field) *zap.Logger {
	log.Debug(message, fields...)
	return log
}

func Error(message string, fields ...zap.Field) *zap.Logger {
	log.Error(message, fields...)
	return log
}

func Fatal(message string, fields ...zap.Field) *zap.Logger {
	log.Fatal(message, fields...)
	return log
}
