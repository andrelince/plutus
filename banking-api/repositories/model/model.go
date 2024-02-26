package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"type:varchar(100)"`
	Email     string `gorm:"type:varchar(100);uniqueIndex"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Accounts  []Account      `gorm:"foreignKey:UserID"`
}

type Account struct {
	ID            uint    `gorm:"primaryKey"`
	UserID        uint    `gorm:"index"`
	AccountNumber string  `gorm:"type:varchar(100);uniqueIndex"`
	Balance       float64 `gorm:"type:decimal(10,2)"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	Transactions  []Transaction  `gorm:"foreignKey:AccountID"`
}

type Transaction struct {
	ID              uint    `gorm:"primaryKey"`
	AccountID       uint    `gorm:"index"`
	Type            string  `gorm:"type:varchar(50)"`   // debit or credit
	Amount          float64 `gorm:"type:decimal(10,2)"` // original amount
	CurrencyCode    string  `gorm:"type:varchar(3)"`
	ConvertedAmount float64 `gorm:"type:decimal(10,2)"` // amount after currency conversion
	ConversionRate  float64 `gorm:"type:decimal(10,6)"` // rate used for conversion
	TransactionFee  float64 `gorm:"type:decimal(10,2)"`
	Status          string  `gorm:"type:varchar(50)"` // completed / failed
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type Currency struct {
	CurrencyCode string `gorm:"primaryKey;type:varchar(3)"`
	Name         string `gorm:"type:varchar(100)"`
	Symbol       string `gorm:"type:varchar(10)"`
}

type CurrencyConversionRate struct {
	ID               uint    `gorm:"primaryKey"`
	FromCurrencyCode string  `gorm:"type:varchar(3);index"`
	ToCurrencyCode   string  `gorm:"type:varchar(3);index"`
	ConversionRate   float64 `gorm:"type:decimal(10,6)"`
	EffectiveDate    time.Time
}
