package transformer

import (
	"github.com/plutus/banking-api/rest/definitions"
	"github.com/plutus/banking-api/services/entities"
)

func FromAccountEntityToDef(in entities.Account) definitions.Account {
	return definitions.Account{
		ID:            in.ID,
		UserID:        in.UserID,
		AccountNumber: in.AccountNumber,
		Balance:       in.Balance,
		CreatedAt:     in.CreatedAt,
		UpdatedAt:     in.UpdatedAt,
		DeletedAt:     in.DeletedAt,
	}
}

func FromTransactionInputDefToModel(in definitions.TransactionInput) entities.Transaction {
	return entities.Transaction{
		Type:         in.Type,
		Amount:       in.Amount,
		CurrencyCode: in.CurrencyCode,
	}
}

func FromTransactionEntityToDef(in entities.Transaction) definitions.Transaction {
	return definitions.Transaction{
		ID:             in.ID,
		AccountID:      in.AccountID,
		Type:           in.Type,
		Amount:         in.Amount,
		CurrencyCode:   in.CurrencyCode,
		TransactionFee: in.TransactionFee,
		Status:         in.Status,
		CreatedAt:      in.CreatedAt,
		UpdatedAt:      in.UpdatedAt,
	}
}
