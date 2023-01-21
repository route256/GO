package main

import (
	"log"

	"go.uber.org/zap"
)

func initLogger() *zap.Logger {
	var logger *zap.Logger
	var err error
	if *develMode {
		logger, err = zap.NewDevelopment()
	} else {
		cfg := zap.NewProductionConfig()
		cfg.DisableCaller = true
		cfg.DisableStacktrace = true
		cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
		logger, err = cfg.Build()
	}
	if err != nil {
		log.Fatal("cannot init zap", err)
	}

	return logger
}
