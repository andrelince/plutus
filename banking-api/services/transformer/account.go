package transformer

import (
	"github.com/plutus/banking-api/repositories/model"
	"github.com/plutus/banking-api/services/entities"
)

func FromAccountModelToEntity(in model.Account) entities.Account {
	return entities.Account{
		ID:            in.ID,
		UserID:        in.UserID,
		AccountNumber: in.AccountNumber,
		Balance:       in.Balance,
		CreatedAt:     in.CreatedAt,
		UpdatedAt:     in.UpdatedAt,
	}
}

func FromAccountEntityToModel(in entities.Account) model.Account {
	return model.Account{
		ID:            in.ID,
		UserID:        in.UserID,
		AccountNumber: in.AccountNumber,
		Balance:       in.Balance,
	}
}

func FromTransactionEntityToModel(in entities.Transaction) model.Transaction {
	return model.Transaction{
		Type:         in.Type,
		Amount:       in.Amount,
		CurrencyCode: in.CurrencyCode,
	}
}

func FromTransactionModelToEntity(in model.Transaction) entities.Transaction {
	return entities.Transaction{
		ID:              in.ID,
		Type:            in.Type,
		Amount:          in.Amount,
		CurrencyCode:    in.CurrencyCode,
		AccountID:       in.AccountID,
		ConvertedAmount: in.ConvertedAmount,
		ConversionRate:  in.ConversionRate,
		TransactionFee:  in.TransactionFee,
		Status:          in.Status,
		CreatedAt:       in.CreatedAt,
		UpdatedAt:       in.UpdatedAt,
	}
}
