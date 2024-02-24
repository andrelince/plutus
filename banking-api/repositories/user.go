package repositories

import (
	"context"
	"time"

	"github.com/plutus/banking-api/repositories/model"
	"gorm.io/gorm"
)

//go:generate mockgen -destination=./mocks/mock_user_connector.go -package=mocks github.com/plutus/banking-api/repositories UserConnector
type UserConnector interface {
	CreateUser(ctx context.Context, user model.User) (model.User, error)
	UpdateUser(ctx context.Context, id uint, user model.User) (model.User, error)
	DeleteUser(ctx context.Context, id uint) error
}

type UserRepo struct {
	g *gorm.DB
}

func NewUserRepo(g *gorm.DB) UserRepo {
	return UserRepo{
		g: g,
	}
}

func (r UserRepo) CreateUser(ctx context.Context, user model.User) (model.User, error) {
	res := r.g.
		WithContext(ctx).
		Create(&user)
	return user, res.Error
}

func (r UserRepo) UpdateUser(ctx context.Context, id uint, user model.User) (model.User, error) {
	var found model.User
	if foundRes := r.g.First(&found, id); foundRes.Error != nil {
		return model.User{}, foundRes.Error
	}

	user.ID = id
	user.UpdatedAt = time.Now()
	res := r.g.
		WithContext(ctx).
		Save(&user)
	return user, res.Error
}

func (r UserRepo) DeleteUser(ctx context.Context, id uint) error {
	var found model.User
	if foundRes := r.g.First(&found, id); foundRes.Error != nil {
		return foundRes.Error
	}

	res := r.g.
		WithContext(ctx).
		Delete(&model.User{}, id)
	return res.Error
}
