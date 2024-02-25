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

func Test_GetAccountByUserIDAndID(t *testing.T) {
	testCases := map[string]struct {
		ctx        context.Context
		userID     uint
		accountID  uint
		serviceOut entities.Account
		serviceErr error
		repoOut    model.Account
		repoErr    error
	}{
		"success": {
			ctx:        context.Background(),
			userID:     1,
			accountID:  1,
			serviceOut: entities.Account{ID: 1, AccountNumber: "12", UserID: 1},
			repoOut:    model.Account{ID: 1, AccountNumber: "12", UserID: 1},
		},
		"error": {
			ctx:        context.Background(),
			userID:     1,
			accountID:  1,
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
				GetAccountByUserIDAndID(args.ctx, args.userID, args.accountID).
				Return(args.repoOut, args.repoErr)

			svc := NewAccountService(mctr)

			out, err := svc.GetAccountByUserIDAndID(args.ctx, args.userID, args.accountID)
			assert.Equal(args.serviceOut, out)
			assert.Equal(args.serviceErr, err)
		})
	}
}

func Test_CreateTransaction(t *testing.T) {

	testCases := map[string]struct {
		ctx        context.Context
		accountID  uint
		serviceIn  entities.Transaction
		serviceOut entities.Transaction
		serviceErr error
		repoIn     model.Transaction
		repoOut    model.Transaction
		repoErr    error
	}{
		"success": {
			ctx:        context.Background(),
			accountID:  1,
			serviceIn:  entities.Transaction{Type: "debit"},
			serviceOut: entities.Transaction{ID: 1, AccountID: 1},
			repoIn:     model.Transaction{Type: "debit"},
			repoOut:    model.Transaction{ID: 1, AccountID: 1},
		},
		"error": {
			ctx:        context.Background(),
			accountID:  1,
			serviceIn:  entities.Transaction{Type: "debit"},
			serviceErr: errors.New("error"),
			repoIn:     model.Transaction{Type: "debit"},
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
				CreateTransaction(args.ctx, args.accountID, args.repoIn).
				Return(args.repoOut, args.repoErr)

			svc := NewAccountService(mctr)

			out, err := svc.CreateTransaction(args.ctx, args.accountID, args.serviceIn)
			assert.Equal(args.serviceOut, out)
			assert.Equal(args.serviceErr, err)
		})
	}
}

func Test_GetTransactions(t *testing.T) {

	testCases := map[string]struct {
		ctx        context.Context
		accountID  uint
		serviceOut []entities.Transaction
		serviceErr error
		repoOut    []model.Transaction
		repoErr    error
	}{
		"success": {
			ctx:        context.Background(),
			accountID:  1,
			serviceOut: []entities.Transaction{{ID: 1}},
			repoOut:    []model.Transaction{{ID: 1}},
		},
		"error": {
			ctx:        context.Background(),
			accountID:  1,
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
				GetAccountTransactions(args.ctx, args.accountID).
				Return(args.repoOut, args.repoErr)

			svc := NewAccountService(mctr)

			out, err := svc.GetAccountTransactions(args.ctx, args.accountID)
			assert.ElementsMatch(args.serviceOut, out)
			assert.Equal(args.serviceErr, err)
		})
	}
}
