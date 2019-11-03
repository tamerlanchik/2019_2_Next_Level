// Code generated by MockGen. DO NOT EDIT.
// Source: server/User/usecase.go

// Package mock is a generated GoMock package.
package mock

import (
	model "2019_2_Next_Level/internal/model"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockUserUsecase is a mock of UserUsecase interface
type MockUserUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockUserUsecaseMockRecorder
}

// MockUserUsecaseMockRecorder is the mock recorder for MockUserUsecase
type MockUserUsecaseMockRecorder struct {
	mock *MockUserUsecase
}

// NewMockUserUsecase creates a new mock instance
func NewMockUserUsecase(ctrl *gomock.Controller) *MockUserUsecase {
	mock := &MockUserUsecase{ctrl: ctrl}
	mock.recorder = &MockUserUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserUsecase) EXPECT() *MockUserUsecaseMockRecorder {
	return m.recorder
}

// GetUser mocks base method
func (m *MockUserUsecase) GetUser(login string) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", login)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser
func (mr *MockUserUsecaseMockRecorder) GetUser(login interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockUserUsecase)(nil).GetUser), login)
}

// EditUser mocks base method
func (m *MockUserUsecase) EditUser(arg0 *model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditUser", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// EditUser indicates an expected call of EditUser
func (mr *MockUserUsecaseMockRecorder) EditUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditUser", reflect.TypeOf((*MockUserUsecase)(nil).EditUser), arg0)
}

// EditPassword mocks base method
func (m *MockUserUsecase) EditPassword(login, currPass, newPass string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditPassword", login, currPass, newPass)
	ret0, _ := ret[0].(error)
	return ret0
}

// EditPassword indicates an expected call of EditPassword
func (mr *MockUserUsecaseMockRecorder) EditPassword(login, currPass, newPass interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditPassword", reflect.TypeOf((*MockUserUsecase)(nil).EditPassword), login, currPass, newPass)
}
