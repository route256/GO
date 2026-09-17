package log

import (
	"os"
	"workshopcache/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(level zapcore.Level) (*zap.Logger, error) {
	var logConfig zap.Config

	if os.Getenv("ENVIRONMENT") == "DEV" {
		logConfig = zap.NewDevelopmentConfig()
	} else {
		logConfig = zap.NewProductionConfig()
	}

	logConfig.Level = zap.NewAtomicLevelAt(level)

	logger, err := logConfig.Build()
	if err != nil {
		return logger, err
	}
	return logger.Named(config.ServiceName), nil
}
