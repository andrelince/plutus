package services

import (
	"context"
	"errors"
	"testing"

	"github.com/plutus/banking-api/repositories/mocks"
	"github.com/plutus/banking-api/repositories/model"
	"github.com/plutus/banking-api/services/entities"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_CreateUser(t *testing.T) {

	testCases := map[string]struct {
		ctx        context.Context
		serviceIn  entities.User
		serviceOut entities.User
		serviceErr error
		repoIn     model.User
		repoOut    model.User
		repoErr    error
	}{
		"success": {
			ctx:        context.Background(),
			serviceIn:  entities.User{Name: "a", Email: "a@a.com"},
			serviceOut: entities.User{ID: 1, Name: "a", Email: "a@a.com"},
			repoIn:     model.User{Name: "a", Email: "a@a.com"},
			repoOut:    model.User{ID: 1, Name: "a", Email: "a@a.com"},
		},
		"error": {
			ctx:        context.Background(),
			serviceIn:  entities.User{Name: "a", Email: "a@a.com"},
			serviceErr: errors.New("error"),
			repoIn:     model.User{Name: "a", Email: "a@a.com"},
			repoErr:    errors.New("error"),
		},
	}

	for name, args := range testCases {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)
			ctrl := gomock.NewController(t)
			mctr := mocks.NewMockUserConnector(ctrl)
			defer ctrl.Finish()

			mctr.
				EXPECT().
				CreateUser(args.ctx, args.repoIn).
				Return(args.repoOut, args.repoErr)

			svc := NewUserService(mctr)

			out, err := svc.CreateUser(args.ctx, args.serviceIn)
			assert.Equal(args.serviceOut, out)
			assert.Equal(args.serviceErr, err)
		})
	}
}

func Test_UpdateUser(t *testing.T) {

	testCases := map[string]struct {
		ctx        context.Context
		id         uint
		serviceIn  entities.User
		serviceOut entities.User
		serviceErr error
		repoIn     model.User
		repoOut    model.User
		repoErr    error
	}{
		"success": {
			ctx:        context.Background(),
			id:         1,
			serviceIn:  entities.User{Name: "a", Email: "a@a.com"},
			serviceOut: entities.User{ID: 1, Name: "a", Email: "a@a.com"},
			repoIn:     model.User{Name: "a", Email: "a@a.com"},
			repoOut:    model.User{ID: 1, Name: "a", Email: "a@a.com"},
		},
		"error": {
			ctx:        context.Background(),
			serviceIn:  entities.User{Name: "a", Email: "a@a.com"},
			serviceErr: errors.New("error"),
			repoIn:     model.User{Name: "a", Email: "a@a.com"},
			repoErr:    errors.New("error"),
		},
	}

	for name, args := range testCases {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)
			ctrl := gomock.NewController(t)
			mctr := mocks.NewMockUserConnector(ctrl)
			defer ctrl.Finish()

			mctr.
				EXPECT().
				UpdateUser(args.ctx, args.id, args.repoIn).
				Return(args.repoOut, args.repoErr)

			svc := NewUserService(mctr)

			out, err := svc.UpdateUser(args.ctx, args.id, args.serviceIn)
			assert.Equal(args.serviceOut, out)
			assert.Equal(args.serviceErr, err)
		})
	}
}
