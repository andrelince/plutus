// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/plutus/banking-api/repositories (interfaces: UserConnector)
//
// Generated by this command:
//
//	mockgen -destination=./mocks/mock_user_connector.go -package=mocks github.com/plutus/banking-api/repositories UserConnector
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	model "github.com/plutus/banking-api/repositories/model"
	gomock "go.uber.org/mock/gomock"
)

// MockUserConnector is a mock of UserConnector interface.
type MockUserConnector struct {
	ctrl     *gomock.Controller
	recorder *MockUserConnectorMockRecorder
}

// MockUserConnectorMockRecorder is the mock recorder for MockUserConnector.
type MockUserConnectorMockRecorder struct {
	mock *MockUserConnector
}

// NewMockUserConnector creates a new mock instance.
func NewMockUserConnector(ctrl *gomock.Controller) *MockUserConnector {
	mock := &MockUserConnector{ctrl: ctrl}
	mock.recorder = &MockUserConnectorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserConnector) EXPECT() *MockUserConnectorMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUserConnector) CreateUser(arg0 context.Context, arg1 model.User) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserConnectorMockRecorder) CreateUser(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserConnector)(nil).CreateUser), arg0, arg1)
}

// DeleteUser mocks base method.
func (m *MockUserConnector) DeleteUser(arg0 context.Context, arg1 uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockUserConnectorMockRecorder) DeleteUser(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUserConnector)(nil).DeleteUser), arg0, arg1)
}

// UpdateUser mocks base method.
func (m *MockUserConnector) UpdateUser(arg0 context.Context, arg1 uint, arg2 model.User) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1, arg2)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockUserConnectorMockRecorder) UpdateUser(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockUserConnector)(nil).UpdateUser), arg0, arg1, arg2)
}
