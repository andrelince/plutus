package transformer

import (
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
	}
}

func FromUserInputDefToEntity(in definitions.UserInput) entities.User {
	return entities.User{
		Name:  in.Name,
		Email: in.Email,
	}
}
