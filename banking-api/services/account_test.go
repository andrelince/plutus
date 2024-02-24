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

func Test_CreateAccount(t *testing.T) {

	testCases := map[string]struct {
		ctx        context.Context
		userID     uint
		serviceOut entities.Account
		serviceErr error
		repoOut    model.Account
		repoErr    error
	}{
		"success": {
			ctx:        context.Background(),
			userID:     1,
			serviceOut: entities.Account{ID: 1, UserID: 1},
			repoOut:    model.Account{ID: 1, UserID: 1},
		},
		"error": {
			ctx:        context.Background(),
			userID:     1,
			serviceErr: errors.New("error"),
			repoErr:    errors.New("error"),
		},
	}

	for name, args := range testCases {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)
			ctrl := gomock.NewController(t)
			mctr := mocks.NewMockAccountConnector(ctrl)
			defer ctrl.Finish()

			mctr.
				EXPECT().
				CreateAccount(args.ctx, args.userID).
				Return(args.repoOut, args.repoErr)

			svc := NewAccountService(mctr)

			out, err := svc.CreateAccount(args.ctx, args.userID)
			assert.Equal(args.serviceOut, out)
			assert.Equal(args.serviceErr, err)
		})
	}
}
