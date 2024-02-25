package transformer

import (
	"github.com/plutus/banking-api/pkg/slice"
	"github.com/plutus/banking-api/repositories/model"
	"github.com/plutus/banking-api/services/entities"
)

func FromUserModelToEntity(in model.User) entities.User {
	out := entities.User{
		ID:        in.ID,
		Name:      in.Name,
		Email:     in.Email,
		CreatedAt: in.CreatedAt,
		UpdatedAt: in.UpdatedAt,
		Accounts:  slice.FromManyToMany(in.Accounts, FromAccountModelToEntity),
	}

	if in.DeletedAt.Valid {
		out.DeletedAt = &in.DeletedAt.Time
	}

	return out
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
