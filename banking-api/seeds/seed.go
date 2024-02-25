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
				if res := d.Save(&model.User{ID: 1, Name: "a", Email: "a@a"}); res.Error != nil {
					return res.Error
				}
				if res := d.Save(&model.User{ID: 2, Name: "b", Email: "b@b"}); res.Error != nil {
					return res.Error
				}
				return nil
			},
		},
		{
			Name: "create accounts",
			Run: func(d *gorm.DB) error {
				if res := d.Save(&model.Account{ID: 1, UserID: 1, AccountNumber: "1"}); res.Error != nil {
					return res.Error
				}
				if res := d.Save(&model.Account{ID: 2, UserID: 2, AccountNumber: "2"}); res.Error != nil {
					return res.Error
				}
				return nil
			},
		},
		{
			Name: "create currencies",
			Run: func(d *gorm.DB) error {
				if res := d.Save(&model.Currency{Name: "Dollar", CurrencyCode: "USD", Symbol: "$"}); res.Error != nil {
					return res.Error
				}
				if res := d.Save(&model.Currency{Name: "Euro", CurrencyCode: "EUR", Symbol: "€"}); res.Error != nil {
					return res.Error
				}
				if res := d.Save(&model.CurrencyConversionRate{
					ID: 1, FromCurrencyCode: "EUR", ToCurrencyCode: "USD", ConversionRate: 1.08,
				}); res.Error != nil {
					return res.Error
				}
				if res := d.Save(&model.CurrencyConversionRate{
					ID: 2, FromCurrencyCode: "USD", ToCurrencyCode: "EUR", ConversionRate: 0.92,
				}); res.Error != nil {
					return res.Error
				}
				return nil
			},
		},
	}
}
