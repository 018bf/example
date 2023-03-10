// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/018bf/example/internal/domain/interceptors (interfaces: PlanInterceptor)

// Package mock_interceptors is a generated GoMock package.
package mock_interceptors

import (
	context "context"
	reflect "reflect"

	models "github.com/018bf/example/internal/domain/models"
	gomock "github.com/golang/mock/gomock"
)

// MockPlanInterceptor is a mock of PlanInterceptor interface.
type MockPlanInterceptor struct {
	ctrl     *gomock.Controller
	recorder *MockPlanInterceptorMockRecorder
}

// MockPlanInterceptorMockRecorder is the mock recorder for MockPlanInterceptor.
type MockPlanInterceptorMockRecorder struct {
	mock *MockPlanInterceptor
}

// NewMockPlanInterceptor creates a new mock instance.
func NewMockPlanInterceptor(ctrl *gomock.Controller) *MockPlanInterceptor {
	mock := &MockPlanInterceptor{ctrl: ctrl}
	mock.recorder = &MockPlanInterceptorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPlanInterceptor) EXPECT() *MockPlanInterceptorMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockPlanInterceptor) Create(arg0 context.Context, arg1 *models.PlanCreate, arg2 *models.User) (*models.Plan, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2)
	ret0, _ := ret[0].(*models.Plan)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockPlanInterceptorMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPlanInterceptor)(nil).Create), arg0, arg1, arg2)
}

// Delete mocks base method.
func (m *MockPlanInterceptor) Delete(arg0 context.Context, arg1 models.UUID, arg2 *models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockPlanInterceptorMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockPlanInterceptor)(nil).Delete), arg0, arg1, arg2)
}

// Get mocks base method.
func (m *MockPlanInterceptor) Get(arg0 context.Context, arg1 models.UUID, arg2 *models.User) (*models.Plan, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(*models.Plan)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockPlanInterceptorMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPlanInterceptor)(nil).Get), arg0, arg1, arg2)
}

// List mocks base method.
func (m *MockPlanInterceptor) List(arg0 context.Context, arg1 *models.PlanFilter, arg2 *models.User) ([]*models.Plan, uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*models.Plan)
	ret1, _ := ret[1].(uint64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List.
func (mr *MockPlanInterceptorMockRecorder) List(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockPlanInterceptor)(nil).List), arg0, arg1, arg2)
}

// Update mocks base method.
func (m *MockPlanInterceptor) Update(arg0 context.Context, arg1 *models.PlanUpdate, arg2 *models.User) (*models.Plan, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(*models.Plan)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockPlanInterceptorMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockPlanInterceptor)(nil).Update), arg0, arg1, arg2)
}
