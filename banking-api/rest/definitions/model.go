package definitions

import "time"

type User struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserInput struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

type Account struct {
	ID            uint      `json:"id"`
	UserID        uint      `json:"user_id"`
	AccountNumber string    `json:"account_number"`
	Balance       float64   `json:"balance"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type Transaction struct {
	ID             uint      `json:"id"`
	AccountID      uint      `json:"account_id"`
	Type           string    `json:"type"`
	Amount         float64   `json:"amount"`
	CurrencyCode   string    `json:"currency_code"`
	TransactionFee float64   `json:"transaction_fee"`
	Status         string    `json:"status"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type TransactionInput struct {
	Type         string  `json:"type" validate:"oneof=debit credit"`
	Amount       float64 `json:"amount" validate:"required,gt=0"`
	CurrencyCode string  `json:"currency_code" validate:"oneof=USD EUR"`
}
