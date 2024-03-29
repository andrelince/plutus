// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/plutus/banking-api/repositories (interfaces: AccountConnector)
//
// Generated by this command:
//
//	mockgen -destination=./mocks/mock_account_connector.go -package=mocks github.com/plutus/banking-api/repositories AccountConnector
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	model "github.com/plutus/banking-api/repositories/model"
	gomock "go.uber.org/mock/gomock"
)

// MockAccountConnector is a mock of AccountConnector interface.
type MockAccountConnector struct {
	ctrl     *gomock.Controller
	recorder *MockAccountConnectorMockRecorder
}

// MockAccountConnectorMockRecorder is the mock recorder for MockAccountConnector.
type MockAccountConnectorMockRecorder struct {
	mock *MockAccountConnector
}

// NewMockAccountConnector creates a new mock instance.
func NewMockAccountConnector(ctrl *gomock.Controller) *MockAccountConnector {
	mock := &MockAccountConnector{ctrl: ctrl}
	mock.recorder = &MockAccountConnectorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountConnector) EXPECT() *MockAccountConnectorMockRecorder {
	return m.recorder
}

// CreateAccount mocks base method.
func (m *MockAccountConnector) CreateAccount(arg0 context.Context, arg1 uint) (model.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccount", arg0, arg1)
	ret0, _ := ret[0].(model.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccount indicates an expected call of CreateAccount.
func (mr *MockAccountConnectorMockRecorder) CreateAccount(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockAccountConnector)(nil).CreateAccount), arg0, arg1)
}

// CreateTransaction mocks base method.
func (m *MockAccountConnector) CreateTransaction(arg0 context.Context, arg1 uint, arg2 model.Transaction) (model.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTransaction", arg0, arg1, arg2)
	ret0, _ := ret[0].(model.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTransaction indicates an expected call of CreateTransaction.
func (mr *MockAccountConnectorMockRecorder) CreateTransaction(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTransaction", reflect.TypeOf((*MockAccountConnector)(nil).CreateTransaction), arg0, arg1, arg2)
}

// DeleteAccount mocks base method.
func (m *MockAccountConnector) DeleteAccount(arg0 context.Context, arg1, arg2 uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAccount", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAccount indicates an expected call of DeleteAccount.
func (mr *MockAccountConnectorMockRecorder) DeleteAccount(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAccount", reflect.TypeOf((*MockAccountConnector)(nil).DeleteAccount), arg0, arg1, arg2)
}

// GetAccountByUserIDAndID mocks base method.
func (m *MockAccountConnector) GetAccountByUserIDAndID(arg0 context.Context, arg1, arg2 uint) (model.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountByUserIDAndID", arg0, arg1, arg2)
	ret0, _ := ret[0].(model.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountByUserIDAndID indicates an expected call of GetAccountByUserIDAndID.
func (mr *MockAccountConnectorMockRecorder) GetAccountByUserIDAndID(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountByUserIDAndID", reflect.TypeOf((*MockAccountConnector)(nil).GetAccountByUserIDAndID), arg0, arg1, arg2)
}

// GetAccountTransactions mocks base method.
func (m *MockAccountConnector) GetAccountTransactions(arg0 context.Context, arg1 uint) ([]model.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountTransactions", arg0, arg1)
	ret0, _ := ret[0].([]model.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountTransactions indicates an expected call of GetAccountTransactions.
func (mr *MockAccountConnectorMockRecorder) GetAccountTransactions(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountTransactions", reflect.TypeOf((*MockAccountConnector)(nil).GetAccountTransactions), arg0, arg1)
}
