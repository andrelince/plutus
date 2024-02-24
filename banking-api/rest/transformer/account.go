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
	}
}
