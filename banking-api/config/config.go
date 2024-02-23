package config

import (
	"github.com/plutus/banking-api/pkg/logger"
)

type Config struct {
	Port           string `env:"SRV_PORT" envDefault:"3000"`
	LoggerSettings logger.Settings
}

func NewLoggerSettings(config Config) logger.Settings {
	return config.LoggerSettings
}
