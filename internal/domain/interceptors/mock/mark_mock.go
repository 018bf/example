// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/018bf/example/internal/domain/interceptors (interfaces: MarkInterceptor)

// Package mock_interceptors is a generated GoMock package.
package mock_interceptors

import (
	context "context"
	reflect "reflect"

	models "github.com/018bf/example/internal/domain/models"
	gomock "github.com/golang/mock/gomock"
)

// MockMarkInterceptor is a mock of MarkInterceptor interface.
type MockMarkInterceptor struct {
	ctrl     *gomock.Controller
	recorder *MockMarkInterceptorMockRecorder
}

// MockMarkInterceptorMockRecorder is the mock recorder for MockMarkInterceptor.
type MockMarkInterceptorMockRecorder struct {
	mock *MockMarkInterceptor
}

// NewMockMarkInterceptor creates a new mock instance.
func NewMockMarkInterceptor(ctrl *gomock.Controller) *MockMarkInterceptor {
	mock := &MockMarkInterceptor{ctrl: ctrl}
	mock.recorder = &MockMarkInterceptorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMarkInterceptor) EXPECT() *MockMarkInterceptorMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockMarkInterceptor) Create(arg0 context.Context, arg1 *models.MarkCreate, arg2 *models.User) (*models.Mark, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2)
	ret0, _ := ret[0].(*models.Mark)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockMarkInterceptorMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMarkInterceptor)(nil).Create), arg0, arg1, arg2)
}

// Delete mocks base method.
func (m *MockMarkInterceptor) Delete(arg0 context.Context, arg1 string, arg2 *models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockMarkInterceptorMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockMarkInterceptor)(nil).Delete), arg0, arg1, arg2)
}

// Get mocks base method.
func (m *MockMarkInterceptor) Get(arg0 context.Context, arg1 string, arg2 *models.User) (*models.Mark, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(*models.Mark)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockMarkInterceptorMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockMarkInterceptor)(nil).Get), arg0, arg1, arg2)
}

// List mocks base method.
func (m *MockMarkInterceptor) List(arg0 context.Context, arg1 *models.MarkFilter, arg2 *models.User) ([]*models.Mark, uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*models.Mark)
	ret1, _ := ret[1].(uint64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List.
func (mr *MockMarkInterceptorMockRecorder) List(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockMarkInterceptor)(nil).List), arg0, arg1, arg2)
}

// Update mocks base method.
func (m *MockMarkInterceptor) Update(arg0 context.Context, arg1 *models.MarkUpdate, arg2 *models.User) (*models.Mark, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(*models.Mark)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockMarkInterceptorMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockMarkInterceptor)(nil).Update), arg0, arg1, arg2)
}