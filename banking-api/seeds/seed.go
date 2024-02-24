package seeds

import (
	"github.com/plutus/banking-api/pkg/pg"
	"github.com/plutus/banking-api/repositories/model"
	"gorm.io/gorm"
)

func GetSeeds() []pg.Seed {
	return []pg.Seed{
		{
			Name: "create users",
			Run: func(d *gorm.DB) error {
				if res := d.Save(&model.User{Name: "a", Email: "a@a"}); res.Error != nil {
					return res.Error
				}
				if res := d.Save(&model.User{Name: "b", Email: "b@b"}); res.Error != nil {
					return res.Error
				}
				return nil
			},
		},
		{
			Name: "create accounts",
			Run: func(d *gorm.DB) error {
				if res := d.Save(&model.Account{UserID: 1}); res.Error != nil {
					return res.Error
				}
				if res := d.Save(&model.Account{UserID: 2}); res.Error != nil {
					return res.Error
				}
				return nil
			},
		},
	}
}
