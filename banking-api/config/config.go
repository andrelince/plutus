package config

import (
	"github.com/plutus/banking-api/pkg/logger"
	"github.com/plutus/banking-api/pkg/pg"
)

type Config struct {
	Port                string `env:"SRV_PORT" envDefault:"3000"`
	LoggerSettings      logger.Settings
	PostgresSettings    pg.PostgresSettings
	TransactionSettings TransactionSettings
}

type TransactionSettings struct {
	BaseCurrency   string  `env:"BASE_CURRENCY" envDefault:"EUR"`
	TransactionFee float64 `env:"TRANSACTION_FEE" envDefault:"0.1"`
}

func NewLoggerSettings(config Config) logger.Settings {
	return config.LoggerSettings
}

func NewPostgresSettings(config Config) pg.PostgresSettings {
	return config.PostgresSettings
}

func NewTransactionSettings(config Config) TransactionSettings {
	return config.TransactionSettings
}
