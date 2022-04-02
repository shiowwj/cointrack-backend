package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger *zap.Logger
)

func init() {
	config := zap.NewDevelopmentConfig()
	config.DisableCaller = true
	config.EncoderConfig.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	config.EncoderConfig.TimeKey = ""
	config.EncoderConfig.StacktraceKey = ""
	config.EncoderConfig.CallerKey = ""
	logger, _ = config.Build()
	// sysout, _ = config.Build()
}

func Info(st string, fields ...zapcore.Field) {
	logger.Info(st, fields...)
}

func Debug(st string, fields ...zapcore.Field) {
	logger.Debug(st, fields...)
}

func Warn(st string, fields ...zapcore.Field) {
	logger.Warn(st, fields...)
}

func Error(st string, fields ...zapcore.Field) {
	logger.Error(st, fields...)
}

func Fatal(st string, fields ...zapcore.Field) {
	logger.Fatal(st, fields...)
}
