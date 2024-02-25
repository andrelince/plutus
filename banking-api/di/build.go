package di

import (
	"net/http"

	"github.com/plutus/banking-api/config"
	"github.com/plutus/banking-api/pkg/env"
	"github.com/plutus/banking-api/pkg/pg"
	"github.com/plutus/banking-api/repositories"
	"github.com/plutus/banking-api/rest"
	"github.com/plutus/banking-api/services"
	"go.uber.org/dig"
	"gorm.io/gorm"
)

func buildConfig(c *dig.Container) error {
	if err := c.Provide(func() (config.Config, error) {
		return env.New(config.Config{})
	}); err != nil {
		return err
	}

	if err := c.Provide(config.NewPostgresSettings); err != nil {
		return err
	}

	if err := c.Provide(config.NewTransactionSettings); err != nil {
		return err
	}

	if err := c.Provide(pg.NewPostgres); err != nil {
		return err
	}

	if err := c.Provide(func(postgres *pg.Postgres) (*gorm.DB, error) {
		return pg.NewGorm(postgres.DB())
	}); err != nil {
		return err
	}

	if err := c.Provide(pg.NewSeedRunner); err != nil {
		return err
	}

	if err := c.Provide(func(g *gorm.DB) repositories.UserConnector {
		return repositories.NewUserRepo(g)
	}); err != nil {
		return err
	}

	if err := c.Provide(func(r repositories.UserConnector) services.UserConnector {
		return services.NewUserService(r)
	}); err != nil {
		return err
	}

	if err := c.Provide(func(g *gorm.DB, s config.TransactionSettings) repositories.AccountConnector {
		return repositories.NewAccountRepo(g, repositories.TransactionSettings{
			BaseCurrency:   s.BaseCurrency,
			TransactionFee: s.TransactionFee,
		})
	}); err != nil {
		return err
	}

	if err := c.Provide(func(r repositories.AccountConnector) services.AccountConnector {
		return services.NewAccountService(r)
	}); err != nil {
		return err
	}

	if err := c.Provide(http.NewServeMux); err != nil {
		return err
	}

	if err := c.Provide(rest.NewHandler); err != nil {
		return err
	}

	return nil
}
