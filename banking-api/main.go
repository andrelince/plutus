package main

import (
	"context"
	"net/http"
	"time"

	"github.com/plutus/banking-api/config"
	"github.com/plutus/banking-api/di"
	"github.com/plutus/banking-api/pkg/pg"
	"github.com/plutus/banking-api/repositories/model"
	"github.com/plutus/banking-api/rest"
	"github.com/plutus/banking-api/seeds"
	"github.com/rs/cors"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
)

func main() {
	c, err := di.NewDI()
	if err != nil {
		panic(err)
	}

	if err := c.Invoke(func(g *gorm.DB) error {
		return g.AutoMigrate(
			&model.User{},
			&model.Account{},
			&model.Transaction{},
			&model.Currency{},
			&model.CurrencyConversionRate{},
		)
	}); err != nil {
		panic(err)
	}

	if err := c.Invoke(func(runner pg.SeedRunner) error {
		return runner.RunSeeds(seeds.GetSeeds()...)
	}); err != nil {
		panic(err)
	}

	if err := c.Invoke(rest.NewRest); err != nil {
		panic(err)
	}

	if err := c.Invoke(start); err != nil {
		panic(err)
	}
}

func start(c config.Config, router *http.ServeMux) error {
	g, _ := errgroup.WithContext(context.Background())

	corss := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "HEAD", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	srv := &http.Server{
		Handler:      corss.Handler(router),
		Addr:         ":" + c.Port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	g.Go(func() error { return srv.ListenAndServe() })

	if err := g.Wait(); err != nil {
		return err
	}

	return nil
}
