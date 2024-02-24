package pg

import (
	"gorm.io/gorm"
)

type SeedRunner struct {
	g *gorm.DB
}

type Seed struct {
	Name string
	Run  func(*gorm.DB) error
}

func NewSeedRunner(g *gorm.DB) SeedRunner {
	return SeedRunner{
		g: g,
	}
}

func (r SeedRunner) RunSeeds(seeds ...Seed) error {
	for _, seed := range seeds {
		if err := seed.Run(r.g); err != nil {
			return err
		}
	}
	return nil
}
