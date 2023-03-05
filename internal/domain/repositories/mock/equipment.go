// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/018bf/example/internal/domain/repositories (interfaces: EquipmentRepository)

// Package mock_repositories is a generated GoMock package.
package mock_repositories

import (
	context "context"
	reflect "reflect"

	models "github.com/018bf/example/internal/domain/models"
	gomock "github.com/golang/mock/gomock"
)

// MockEquipmentRepository is a mock of EquipmentRepository interface.
type MockEquipmentRepository struct {
	ctrl     *gomock.Controller
	recorder *MockEquipmentRepositoryMockRecorder
}

// MockEquipmentRepositoryMockRecorder is the mock recorder for MockEquipmentRepository.
type MockEquipmentRepositoryMockRecorder struct {
	mock *MockEquipmentRepository
}

// NewMockEquipmentRepository creates a new mock instance.
func NewMockEquipmentRepository(ctrl *gomock.Controller) *MockEquipmentRepository {
	mock := &MockEquipmentRepository{ctrl: ctrl}
	mock.recorder = &MockEquipmentRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEquipmentRepository) EXPECT() *MockEquipmentRepositoryMockRecorder {
	return m.recorder
}

// Count mocks base method.
func (m *MockEquipmentRepository) Count(arg0 context.Context, arg1 *models.EquipmentFilter) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", arg0, arg1)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockEquipmentRepositoryMockRecorder) Count(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockEquipmentRepository)(nil).Count), arg0, arg1)
}

// Create mocks base method.
func (m *MockEquipmentRepository) Create(arg0 context.Context, arg1 *models.Equipment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockEquipmentRepositoryMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockEquipmentRepository)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockEquipmentRepository) Delete(arg0 context.Context, arg1 models.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockEquipmentRepositoryMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockEquipmentRepository)(nil).Delete), arg0, arg1)
}

// Get mocks base method.
func (m *MockEquipmentRepository) Get(arg0 context.Context, arg1 models.UUID) (*models.Equipment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*models.Equipment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockEquipmentRepositoryMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockEquipmentRepository)(nil).Get), arg0, arg1)
}

// List mocks base method.
func (m *MockEquipmentRepository) List(arg0 context.Context, arg1 *models.EquipmentFilter) ([]*models.Equipment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]*models.Equipment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockEquipmentRepositoryMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockEquipmentRepository)(nil).List), arg0, arg1)
}

// Update mocks base method.
func (m *MockEquipmentRepository) Update(arg0 context.Context, arg1 *models.Equipment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockEquipmentRepositoryMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockEquipmentRepository)(nil).Update), arg0, arg1)
}