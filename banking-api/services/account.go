package services

import (
	"context"

	"github.com/plutus/banking-api/repositories"
	"github.com/plutus/banking-api/services/entities"
	"github.com/plutus/banking-api/services/transformer"
)

type AccountConnector interface {
	CreateAccount(ctx context.Context, userID uint) (entities.Account, error)
}

type AccountService struct {
	accountRepo repositories.AccountConnector
}

func NewAccountService(accountRepo repositories.AccountConnector) AccountService {
	return AccountService{
		accountRepo: accountRepo,
	}
}

func (r AccountService) CreateAccount(ctx context.Context, userID uint) (entities.Account, error) {
	u, err := r.accountRepo.CreateAccount(ctx, userID)
	if err != nil {
		return entities.Account{}, err
	}
	return transformer.FromAccountModelToEntity(u), nil
}
