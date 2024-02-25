package entities

import (
	"time"
)

type User struct {
	ID        uint
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Account struct {
	ID            uint
	UserID        uint
	AccountNumber string
	Balance       float64
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Transaction struct {
	ID              uint
	Type            string
	Amount          float64
	CurrencyCode    string
	AccountID       uint
	ConvertedAmount float64
	ConversionRate  float64
	TransactionFee  float64
	Status          string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
