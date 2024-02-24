package repositories

import (
	"context"

	"github.com/plutus/banking-api/repositories/model"
	"gorm.io/gorm"
)

//go:generate mockgen -destination=./mocks/mock_user_connector.go -package=mocks github.com/plutus/banking-api/repositories UserConnector
type UserConnector interface {
	CreateUser(ctx context.Context, user model.User) (model.User, error)
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
