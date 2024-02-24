package transformer

import (
	"github.com/plutus/banking-api/repositories/model"
	"github.com/plutus/banking-api/services/entities"
)

func FromUserModelToEntity(in model.User) entities.User {
	return entities.User{
		ID:        in.ID,
		Name:      in.Name,
		Email:     in.Email,
		CreatedAt: in.CreatedAt,
		UpdatedAt: in.UpdatedAt,
	}
}

func FromUserEntityToModel(in entities.User) model.User {
	return model.User{
		ID:        in.ID,
		Name:      in.Name,
		Email:     in.Email,
		CreatedAt: in.CreatedAt,
		UpdatedAt: in.UpdatedAt,
	}
}
