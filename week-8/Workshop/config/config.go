package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

const ServiceName = "miner"

type Config struct {
	App    *AppConfig
	Log    *LogConfig
	Crypto *CryptoConfig
	Cache  *CacheConfig
}

func InitError(err error) {
	if err != nil {
		panic(err)
	}
}

func InitConfig() *Config {
	return &Config{
		App:    InitAppConfig(),
		Log:    InitLogConfig(),
		Crypto: InitCryptoConfig(),
		Cache:  InitCacheConfig(),
	}
}

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("There is no .env file provided!")
	}
	viper.SetDefault("environment", "DEV")
	InitError(viper.BindEnv("environment", "ENVIRONMENT"))
}
