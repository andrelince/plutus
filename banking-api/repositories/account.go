package repositories

import (
	"context"

	"github.com/plutus/banking-api/repositories/model"
	"gorm.io/gorm"
)

//go:generate mockgen -destination=./mocks/mock_account_connector.go -package=mocks github.com/plutus/banking-api/repositories AccountConnector
type AccountConnector interface {
	CreateAccount(ctx context.Context, userID uint) (model.Account, error)
}

type AccountRepo struct {
	g *gorm.DB
}

func NewAccountRepo(g *gorm.DB) AccountRepo {
	return AccountRepo{
		g: g,
	}
}

func (r AccountRepo) CreateAccount(ctx context.Context, userID uint) (model.Account, error) {
	account := model.Account{UserID: userID}
	res := r.g.
		WithContext(ctx).
		Create(&account)
	return account, res.Error
}
