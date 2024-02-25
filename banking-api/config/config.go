package config

import (
	"github.com/plutus/banking-api/pkg/pg"
)

type Config struct {
	Port                string `env:"SRV_PORT" envDefault:"3000"`
	PostgresSettings    pg.PostgresSettings
	TransactionSettings TransactionSettings
}

type TransactionSettings struct {
	BaseCurrency   string  `env:"BASE_CURRENCY" envDefault:"EUR"`
	TransactionFee float64 `env:"TRANSACTION_FEE" envDefault:"0.1"`
}

func NewPostgresSettings(config Config) pg.PostgresSettings {
	return config.PostgresSettings
}

func NewTransactionSettings(config Config) TransactionSettings {
	return config.TransactionSettings
}
