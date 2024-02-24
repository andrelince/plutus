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
	GetUserByID(ctx context.Context, id uint) (model.User, error)
	GetUsers(ctx context.Context) ([]model.User, error)
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
	if _, err := r.GetUserByID(ctx, id); err != nil {
		return model.User{}, err
	}

	user.ID = id
	user.UpdatedAt = time.Now()
	res := r.g.
		WithContext(ctx).
		Save(&user)
	return user, res.Error
}

func (r UserRepo) DeleteUser(ctx context.Context, id uint) error {
	if _, err := r.GetUserByID(ctx, id); err != nil {
		return err
	}

	res := r.g.
		WithContext(ctx).
		Delete(&model.User{}, id)
	return res.Error
}

func (r UserRepo) GetUserByID(ctx context.Context, id uint) (model.User, error) {
	var found model.User
	foundRes := r.g.First(&found, id)
	return found, foundRes.Error
}

func (r UserRepo) GetUsers(ctx context.Context) ([]model.User, error) {
	var out []model.User
	res := r.g.
		WithContext(ctx).
		Find(&out)
	return out, res.Error
}
