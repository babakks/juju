// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/facades/controller/migrationtarget (interfaces: ControllerConfigService,ExternalControllerService)

// Package migrationtarget_test is a generated GoMock package.
package migrationtarget_test

import (
	context "context"
	reflect "reflect"

	controller "github.com/juju/juju/controller"
	crossmodel "github.com/juju/juju/core/crossmodel"
	gomock "go.uber.org/mock/gomock"
)

// MockControllerConfigService is a mock of ControllerConfigService interface.
type MockControllerConfigService struct {
	ctrl     *gomock.Controller
	recorder *MockControllerConfigServiceMockRecorder
}

// MockControllerConfigServiceMockRecorder is the mock recorder for MockControllerConfigService.
type MockControllerConfigServiceMockRecorder struct {
	mock *MockControllerConfigService
}

// NewMockControllerConfigService creates a new mock instance.
func NewMockControllerConfigService(ctrl *gomock.Controller) *MockControllerConfigService {
	mock := &MockControllerConfigService{ctrl: ctrl}
	mock.recorder = &MockControllerConfigServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockControllerConfigService) EXPECT() *MockControllerConfigServiceMockRecorder {
	return m.recorder
}

// ControllerConfig mocks base method.
func (m *MockControllerConfigService) ControllerConfig(arg0 context.Context) (controller.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerConfig", arg0)
	ret0, _ := ret[0].(controller.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ControllerConfig indicates an expected call of ControllerConfig.
func (mr *MockControllerConfigServiceMockRecorder) ControllerConfig(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerConfig", reflect.TypeOf((*MockControllerConfigService)(nil).ControllerConfig), arg0)
}

// MockExternalControllerService is a mock of ExternalControllerService interface.
type MockExternalControllerService struct {
	ctrl     *gomock.Controller
	recorder *MockExternalControllerServiceMockRecorder
}

// MockExternalControllerServiceMockRecorder is the mock recorder for MockExternalControllerService.
type MockExternalControllerServiceMockRecorder struct {
	mock *MockExternalControllerService
}

// NewMockExternalControllerService creates a new mock instance.
func NewMockExternalControllerService(ctrl *gomock.Controller) *MockExternalControllerService {
	mock := &MockExternalControllerService{ctrl: ctrl}
	mock.recorder = &MockExternalControllerServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExternalControllerService) EXPECT() *MockExternalControllerServiceMockRecorder {
	return m.recorder
}

// ControllerForModel mocks base method.
func (m *MockExternalControllerService) ControllerForModel(arg0 context.Context, arg1 string) (*crossmodel.ControllerInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerForModel", arg0, arg1)
	ret0, _ := ret[0].(*crossmodel.ControllerInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ControllerForModel indicates an expected call of ControllerForModel.
func (mr *MockExternalControllerServiceMockRecorder) ControllerForModel(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerForModel", reflect.TypeOf((*MockExternalControllerService)(nil).ControllerForModel), arg0, arg1)
}

// UpdateExternalController mocks base method.
func (m *MockExternalControllerService) UpdateExternalController(arg0 context.Context, arg1 crossmodel.ControllerInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateExternalController", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateExternalController indicates an expected call of UpdateExternalController.
func (mr *MockExternalControllerServiceMockRecorder) UpdateExternalController(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateExternalController", reflect.TypeOf((*MockExternalControllerService)(nil).UpdateExternalController), arg0, arg1)
}
