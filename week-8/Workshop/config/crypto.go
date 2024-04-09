package config

import "github.com/spf13/viper"

type CryptoConfig struct {
	Rule int64
}

func InitCryptoConfig() *CryptoConfig {
	cryptoConfig := CryptoConfig{
		Rule: viper.GetInt64("crypto.rule"),
	}
	return &cryptoConfig
}

func init() {
	InitError(viper.BindEnv("crypto.rule", "CRYPTO_RULE"))
}
