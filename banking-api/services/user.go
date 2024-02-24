package services

import (
	"context"

	"github.com/plutus/banking-api/repositories"
	"github.com/plutus/banking-api/services/entities"
	"github.com/plutus/banking-api/services/transformer"
)

type UserConnector interface {
	CreateUser(ctx context.Context, user entities.User) (entities.User, error)
}

type UserService struct {
	userRepo repositories.UserConnector
}

func NewUserService(userRepo repositories.UserConnector) UserService {
	return UserService{
		userRepo: userRepo,
	}
}

func (r UserService) CreateUser(ctx context.Context, user entities.User) (entities.User, error) {
	u, err := r.userRepo.CreateUser(ctx, transformer.FromUserEntityToModel(user))
	if err != nil {
		return entities.User{}, err
	}
	return transformer.FromUserModelToEntity(u), nil
}
