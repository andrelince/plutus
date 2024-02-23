package models

import (
	"time"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"type:varchar(100)"`
	Email     string `gorm:"type:varchar(100);uniqueIndex"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Accounts  []Account `gorm:"foreignKey:UserID"`
}

type Account struct {
	ID            uint    `gorm:"primaryKey"`
	UserID        uint    `gorm:"index"`
	AccountNumber string  `gorm:"type:varchar(100);uniqueIndex"`
	Balance       float64 `gorm:"type:decimal(10,2)"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Transactions  []Transaction `gorm:"foreignKey:AccountID"`
}

type Transaction struct {
	ID             uint    `gorm:"primaryKey"`
	AccountID      uint    `gorm:"index"`
	Type           string  `gorm:"type:varchar(50)"`
	Amount         float64 `gorm:"type:decimal(10,2)"`
	TransactionFee float64 `gorm:"type:decimal(10,2)"`
	CurrencyCode   string  `gorm:"type:varchar(3)"`
	Status         string  `gorm:"type:varchar(50)"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type Currency struct {
	CurrencyCode   string  `gorm:"primaryKey;type:varchar(3)"`
	Name           string  `gorm:"type:varchar(100)"`
	Symbol         string  `gorm:"type:varchar(10)"`
	ConversionRate float64 `gorm:"type:decimal(10,2)"`
}
