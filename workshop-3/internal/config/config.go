package config

import (
	"os"
	"time"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

const configFile = "data/config.yaml"

type Config struct {
	Token            string           `yaml:"token"`
	CurrencySettings CurrencySettings `yaml:"currency_settings"`
}

type CurrencySettings struct {
	BaseCurrency                 string   `yaml:"base_currency"`
	SupportedCodes               []string `yaml:"supported_codes"`
	FrequencyExchangeRateUpdates int      `yaml:"frequency_exchange_rate_updates"`
}

type Service struct {
	config Config
}

func New() (*Service, error) {
	s := &Service{}

	rawYAML, err := os.ReadFile(configFile)
	if err != nil {
		return nil, errors.Wrap(err, "reading config file")
	}

	err = yaml.Unmarshal(rawYAML, &s.config)
	if err != nil {
		return nil, errors.Wrap(err, "parsing yaml")
	}

	return s, nil
}

func (s *Service) Token() string {
	return s.config.Token
}

func (s *Service) SupportedCurrencyCodes() []string {
	return s.config.CurrencySettings.SupportedCodes
}

func (s *Service) GetBaseCurrency() string {
	return s.config.CurrencySettings.BaseCurrency
}

func (s *Service) GetFrequencyExchangeRate() time.Duration {
	return time.Duration(s.config.CurrencySettings.FrequencyExchangeRateUpdates) * time.Second
}
