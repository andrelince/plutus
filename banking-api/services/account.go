package services

import (
	"context"

	"github.com/plutus/banking-api/pkg/slice"
	"github.com/plutus/banking-api/repositories"
	"github.com/plutus/banking-api/services/entities"
	"github.com/plutus/banking-api/services/transformer"
)

type AccountConnector interface {
	CreateAccount(ctx context.Context, userID uint) (entities.Account, error)
	GetAccountByUserIDAndID(ctx context.Context, userID, accountID uint) (entities.Account, error)
	CreateTransaction(ctx context.Context, accountID uint, transaction entities.Transaction) (entities.Transaction, error)
	GetAccountTransactions(ctx context.Context, accountID uint) ([]entities.Transaction, error)
	DeleteAccount(ctx context.Context, userID, accountID uint) error
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
	a, err := r.accountRepo.CreateAccount(ctx, userID)
	if err != nil {
		return entities.Account{}, err
	}
	return transformer.FromAccountModelToEntity(a), nil
}

func (r AccountService) GetAccountByUserIDAndID(ctx context.Context, userID, accountID uint) (entities.Account, error) {
	a, err := r.accountRepo.GetAccountByUserIDAndID(ctx, userID, accountID)
	if err != nil {
		return entities.Account{}, err
	}
	return transformer.FromAccountModelToEntity(a), nil
}

func (r AccountService) CreateTransaction(ctx context.Context, accountID uint, transaction entities.Transaction) (entities.Transaction, error) {
	t, err := r.accountRepo.CreateTransaction(ctx, accountID, transformer.FromTransactionEntityToModel(transaction))
	if err != nil {
		return entities.Transaction{}, err
	}
	return transformer.FromTransactionModelToEntity(t), nil
}

func (r AccountService) GetAccountTransactions(ctx context.Context, accountID uint) ([]entities.Transaction, error) {
	t, err := r.accountRepo.GetAccountTransactions(ctx, accountID)
	if err != nil {
		return []entities.Transaction{}, err
	}
	return slice.FromManyToMany(t, transformer.FromTransactionModelToEntity), nil
}

func (r AccountService) DeleteAccount(ctx context.Context, userID, accountID uint) error {
	return r.accountRepo.DeleteAccount(ctx, userID, accountID)
}
