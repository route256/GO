package log

import (
	"go.uber.org/zap"
)

func InitLogger(serviceName string) (*zap.Logger, error) {
	var logConfig zap.Config

	logConfig = zap.NewDevelopmentConfig()

	logConfig.EncoderConfig.LineEnding = "\n\n"
	logConfig.Level = zap.NewAtomicLevelAt(zap.InfoLevel)

	logger, _ := logConfig.Build()

	return logger.Named(serviceName), nil
}
