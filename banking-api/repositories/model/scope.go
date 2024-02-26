package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (account *Account) BeforeCreate(tx *gorm.DB) (err error) {
	account.AccountNumber = uuid.NewString()
	account.CreatedAt = time.Now()
	return
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.CreatedAt = time.Now()
	return
}
