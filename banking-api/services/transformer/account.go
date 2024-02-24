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
