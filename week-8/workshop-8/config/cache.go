package config

import "github.com/spf13/viper"

type CacheConfig struct {
	Host string
	Port string
}

func InitCacheConfig() *CacheConfig {
	cacheConfig := CacheConfig{
		Host: viper.GetString("redis.host"),
		Port: viper.GetString("redis.port"),
	}
	return &cacheConfig
}

func init() {
	InitError(viper.BindEnv("redis.host", "REDIS_HOST"))
	InitError(viper.BindEnv("redis.port", "REDIS_PORT"))
}
