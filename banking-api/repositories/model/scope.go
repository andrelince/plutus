package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (account *Account) BeforeCreate(tx *gorm.DB) (err error) {
	account.AccountNumber = uuid.NewString()
	return
}
