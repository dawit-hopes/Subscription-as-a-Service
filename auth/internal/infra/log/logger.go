// Package log provides a simple wrapper around zap for structured logging.
package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func InitLogger() error {
	config := zap.NewDevelopmentConfig()

	// Customize time format to ISO8601 for readability
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// Enable colorized level output (Info, Error, etc.)
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	// Use console encoder (human-readable, not JSON)
	config.Encoding = "console"

	// Optional: Customize message key or caller key if needed
	config.EncoderConfig.MessageKey = "msg"
	config.EncoderConfig.CallerKey = "caller"
	config.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	var err error
	Logger, err = config.Build()
	return err
}

func Sync() {
	_ = Logger.Sync()
}
