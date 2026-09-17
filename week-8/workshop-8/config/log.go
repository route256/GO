package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap/zapcore"
)

type LogConfig struct {
	Level zapcore.Level
}

func InitLogConfig() *LogConfig {
	var logLevel zapcore.Level

	logLevelStr := viper.GetString("log.level")
	err := logLevel.UnmarshalText([]byte(logLevelStr))
	if err != nil {
		panic(err)
	}

	return &LogConfig{
		Level: logLevel,
	}
}

func init() {
	viper.SetDefault("log.level", "DEBUG")
	InitError(viper.BindEnv("log.level", "LOG_LEVEL"))
}
