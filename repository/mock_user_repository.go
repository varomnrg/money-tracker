// Code generated by MockGen. DO NOT EDIT.
// Source: repository/user_interface.go
//
// Generated by this command:
//
//	mockgen -source repository/user_interface.go -destination repository/mock_user_repository.go -package repository
//
// Package repository is a generated GoMock package.
package repository

import (
	reflect "reflect"

	model "github.com/varomnrg/money-tracker/model"
	gomock "go.uber.org/mock/gomock"
)

// MockIUserRepository is a mock of IUserRepository interface.
type MockIUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIUserRepositoryMockRecorder
}

// MockIUserRepositoryMockRecorder is the mock recorder for MockIUserRepository.
type MockIUserRepositoryMockRecorder struct {
	mock *MockIUserRepository
}

// NewMockIUserRepository creates a new mock instance.
func NewMockIUserRepository(ctrl *gomock.Controller) *MockIUserRepository {
	mock := &MockIUserRepository{ctrl: ctrl}
	mock.recorder = &MockIUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUserRepository) EXPECT() *MockIUserRepositoryMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockIUserRepository) CreateUser(movie model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", movie)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockIUserRepositoryMockRecorder) CreateUser(movie any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockIUserRepository)(nil).CreateUser), movie)
}

// DeleteUser mocks base method.
func (m *MockIUserRepository) DeleteUser(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockIUserRepositoryMockRecorder) DeleteUser(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockIUserRepository)(nil).DeleteUser), id)
}

// GetUser mocks base method.
func (m *MockIUserRepository) GetUser(id string) (model.UserResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", id)
	ret0, _ := ret[0].(model.UserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockIUserRepositoryMockRecorder) GetUser(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockIUserRepository)(nil).GetUser), id)
}

// GetUsers mocks base method.
func (m *MockIUserRepository) GetUsers() ([]model.UserResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers")
	ret0, _ := ret[0].([]model.UserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsers indicates an expected call of GetUsers.
func (mr *MockIUserRepositoryMockRecorder) GetUsers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockIUserRepository)(nil).GetUsers))
}

// IsUsernameExist mocks base method.
func (m *MockIUserRepository) IsUsernameExist(username string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsUsernameExist", username)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsUsernameExist indicates an expected call of IsUsernameExist.
func (mr *MockIUserRepositoryMockRecorder) IsUsernameExist(username any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsUsernameExist", reflect.TypeOf((*MockIUserRepository)(nil).IsUsernameExist), username)
}

// UpdateUser mocks base method.
func (m *MockIUserRepository) UpdateUser(id string, movie model.UserRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", id, movie)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockIUserRepositoryMockRecorder) UpdateUser(id, movie any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockIUserRepository)(nil).UpdateUser), id, movie)
}