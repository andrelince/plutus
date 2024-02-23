package di

import (
	"net/http"

	"github.com/plutus/banking-api/config"
	"github.com/plutus/banking-api/pkg/env"
	"github.com/plutus/banking-api/pkg/logger"
	"github.com/plutus/banking-api/rest"
	"go.uber.org/dig"
)

func buildConfig(c *dig.Container) error {
	if err := c.Provide(func() (config.Config, error) {
		return env.New(config.Config{})
	}); err != nil {
		return err
	}

	if err := c.Provide(config.NewLoggerSettings); err != nil {
		return err
	}

	if err := c.Provide(func(s logger.Settings) (logger.Logger, error) {
		return logger.New(s), nil
	}); err != nil {
		return err
	}

	if err := c.Provide(func() *http.ServeMux {
		return http.NewServeMux()
	}); err != nil {
		return err
	}

	if err := c.Provide(rest.NewHandler); err != nil {
		return err
	}

	return nil
}
