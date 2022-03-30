package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	config := zap.NewProductionConfig()
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.StacktraceKey = ""
	config.EncoderConfig = encoderConfig

	logger, err := config.Build(zap.AddCallerSkip(1)) // skip is needed to skip wrapper level
	log = logger
	if err != nil {
		panic(err)
	}
}

func Info(template string, fields ...zap.Field) {
	log.Info(template, fields...)
}

func Debug(template string, fields ...zap.Field) {
	log.Debug(template, fields...)
}

func Error(template string, fields ...zap.Field) {
	log.Error(template, fields...)
}

func Infof(template string, args ...interface{}) {
	log.Sugar().Infof(template, args...)
}
