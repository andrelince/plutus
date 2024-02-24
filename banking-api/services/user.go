package services

import (
	"context"

	"github.com/plutus/banking-api/repositories"
	"github.com/plutus/banking-api/services/entities"
	"github.com/plutus/banking-api/services/transformer"
)

type UserConnector interface {
	CreateUser(ctx context.Context, user entities.User) (entities.User, error)
	UpdateUser(ctx context.Context, id uint, user entities.User) (entities.User, error)
	DeleteUser(ctx context.Context, id uint) error
	GetUserByID(ctx context.Context, id uint) (entities.User, error)
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

func (r UserService) UpdateUser(ctx context.Context, id uint, user entities.User) (entities.User, error) {
	u, err := r.userRepo.UpdateUser(ctx, id, transformer.FromUserEntityToModel(user))
	if err != nil {
		return entities.User{}, err
	}
	return transformer.FromUserModelToEntity(u), nil
}

func (r UserService) DeleteUser(ctx context.Context, id uint) error {
	return r.userRepo.DeleteUser(ctx, id)
}

func (r UserService) GetUserByID(ctx context.Context, id uint) (entities.User, error) {
	u, err := r.userRepo.GetUserByID(ctx, id)
	if err != nil {
		return entities.User{}, err
	}
	return transformer.FromUserModelToEntity(u), nil
}
