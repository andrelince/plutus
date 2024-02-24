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
