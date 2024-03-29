package transformer

import (
	"github.com/plutus/banking-api/pkg/slice"
	"github.com/plutus/banking-api/rest/definitions"
	"github.com/plutus/banking-api/services/entities"
)

func FromUserEntityToDef(in entities.User) definitions.User {
	return definitions.User{
		ID:        in.ID,
		Name:      in.Name,
		Email:     in.Email,
		CreatedAt: in.CreatedAt,
		UpdatedAt: in.UpdatedAt,
		DeletedAt: in.DeletedAt,
		Accounts:  slice.FromManyToMany(in.Accounts, FromAccountEntityToDef),
	}
}

func FromUserInputDefToEntity(in definitions.UserInput) entities.User {
	return entities.User{
		Name:  in.Name,
		Email: in.Email,
	}
}
