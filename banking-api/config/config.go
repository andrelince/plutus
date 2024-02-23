package config

import (
	"github.com/plutus/banking-api/pkg/logger"
	"github.com/plutus/banking-api/pkg/pg"
)

type Config struct {
	Port             string `env:"SRV_PORT" envDefault:"3000"`
	LoggerSettings   logger.Settings
	PostgresSettings pg.PostgresSettings
}

func NewLoggerSettings(config Config) logger.Settings {
	return config.LoggerSettings
}

func NewPostgresSettings(config Config) pg.PostgresSettings {
	return config.PostgresSettings
}
