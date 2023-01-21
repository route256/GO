package logger

import (
	"log"

	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {
	localLogger, err := zap.NewProduction()
	if err != nil {
		log.Fatal("logger init", err)
	}

	logger = localLogger
}

func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}
