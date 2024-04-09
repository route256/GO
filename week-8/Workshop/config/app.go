package config

import "github.com/spf13/viper"

type AppConfig struct {
	Host string
	Port string
}

func InitAppConfig() *AppConfig {
	appConfig := AppConfig{
		Host: viper.GetString("app.host"),
		Port: viper.GetString("app.port"),
	}
	return &appConfig
}

func init() {
	InitError(viper.BindEnv("app.host", "APP_HOST"))
	InitError(viper.BindEnv("app.port", "APP_PORT"))
}
