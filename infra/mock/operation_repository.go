// Code generated by MockGen. DO NOT EDIT.
// Source: domain/operation.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/piovani/digital_wallet_2/domain"
)

// MockOperationRepository is a mock of OperationRepository interface.
type MockOperationRepository struct {
	ctrl     *gomock.Controller
	recorder *MockOperationRepositoryMockRecorder
}

// MockOperationRepositoryMockRecorder is the mock recorder for MockOperationRepository.
type MockOperationRepositoryMockRecorder struct {
	mock *MockOperationRepository
}

// NewMockOperationRepository creates a new mock instance.
func NewMockOperationRepository(ctrl *gomock.Controller) *MockOperationRepository {
	mock := &MockOperationRepository{ctrl: ctrl}
	mock.recorder = &MockOperationRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOperationRepository) EXPECT() *MockOperationRepositoryMockRecorder {
	return m.recorder
}

// FindByUserName mocks base method.
func (m *MockOperationRepository) FindByUserName(username string) ([]domain.Operation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByUserName", username)
	ret0, _ := ret[0].([]domain.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByUserName indicates an expected call of FindByUserName.
func (mr *MockOperationRepositoryMockRecorder) FindByUserName(username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByUserName", reflect.TypeOf((*MockOperationRepository)(nil).FindByUserName), username)
}

// Insert mocks base method.
func (m *MockOperationRepository) Insert(opt *domain.Operation) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", opt)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockOperationRepositoryMockRecorder) Insert(opt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockOperationRepository)(nil).Insert), opt)
}
