// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/internal/servicefactory (interfaces: ProviderServiceFactory,ProviderServiceFactoryGetter)
//
// Generated by this command:
//
//	mockgen -package providerservicefactory -destination servicefactory_mock_test.go github.com/juju/juju/internal/servicefactory ProviderServiceFactory,ProviderServiceFactoryGetter
//

// Package providerservicefactory is a generated GoMock package.
package providerservicefactory

import (
	reflect "reflect"

	service "github.com/juju/juju/domain/cloud/service"
	service0 "github.com/juju/juju/domain/credential/service"
	service1 "github.com/juju/juju/domain/model/service"
	service2 "github.com/juju/juju/domain/modelconfig/service"
	servicefactory "github.com/juju/juju/internal/servicefactory"
	gomock "go.uber.org/mock/gomock"
)

// MockProviderServiceFactory is a mock of ProviderServiceFactory interface.
type MockProviderServiceFactory struct {
	ctrl     *gomock.Controller
	recorder *MockProviderServiceFactoryMockRecorder
}

// MockProviderServiceFactoryMockRecorder is the mock recorder for MockProviderServiceFactory.
type MockProviderServiceFactoryMockRecorder struct {
	mock *MockProviderServiceFactory
}

// NewMockProviderServiceFactory creates a new mock instance.
func NewMockProviderServiceFactory(ctrl *gomock.Controller) *MockProviderServiceFactory {
	mock := &MockProviderServiceFactory{ctrl: ctrl}
	mock.recorder = &MockProviderServiceFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProviderServiceFactory) EXPECT() *MockProviderServiceFactoryMockRecorder {
	return m.recorder
}

// Cloud mocks base method.
func (m *MockProviderServiceFactory) Cloud() *service.WatchableProviderService {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cloud")
	ret0, _ := ret[0].(*service.WatchableProviderService)
	return ret0
}

// Cloud indicates an expected call of Cloud.
func (mr *MockProviderServiceFactoryMockRecorder) Cloud() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cloud", reflect.TypeOf((*MockProviderServiceFactory)(nil).Cloud))
}

// Config mocks base method.
func (m *MockProviderServiceFactory) Config() *service2.WatchableProviderService {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Config")
	ret0, _ := ret[0].(*service2.WatchableProviderService)
	return ret0
}

// Config indicates an expected call of Config.
func (mr *MockProviderServiceFactoryMockRecorder) Config() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Config", reflect.TypeOf((*MockProviderServiceFactory)(nil).Config))
}

// Credential mocks base method.
func (m *MockProviderServiceFactory) Credential() *service0.WatchableProviderService {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Credential")
	ret0, _ := ret[0].(*service0.WatchableProviderService)
	return ret0
}

// Credential indicates an expected call of Credential.
func (mr *MockProviderServiceFactoryMockRecorder) Credential() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Credential", reflect.TypeOf((*MockProviderServiceFactory)(nil).Credential))
}

// Model mocks base method.
func (m *MockProviderServiceFactory) Model() *service1.ProviderService {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Model")
	ret0, _ := ret[0].(*service1.ProviderService)
	return ret0
}

// Model indicates an expected call of Model.
func (mr *MockProviderServiceFactoryMockRecorder) Model() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Model", reflect.TypeOf((*MockProviderServiceFactory)(nil).Model))
}

// MockProviderServiceFactoryGetter is a mock of ProviderServiceFactoryGetter interface.
type MockProviderServiceFactoryGetter struct {
	ctrl     *gomock.Controller
	recorder *MockProviderServiceFactoryGetterMockRecorder
}

// MockProviderServiceFactoryGetterMockRecorder is the mock recorder for MockProviderServiceFactoryGetter.
type MockProviderServiceFactoryGetterMockRecorder struct {
	mock *MockProviderServiceFactoryGetter
}

// NewMockProviderServiceFactoryGetter creates a new mock instance.
func NewMockProviderServiceFactoryGetter(ctrl *gomock.Controller) *MockProviderServiceFactoryGetter {
	mock := &MockProviderServiceFactoryGetter{ctrl: ctrl}
	mock.recorder = &MockProviderServiceFactoryGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProviderServiceFactoryGetter) EXPECT() *MockProviderServiceFactoryGetterMockRecorder {
	return m.recorder
}

// FactoryForModel mocks base method.
func (m *MockProviderServiceFactoryGetter) FactoryForModel(arg0 string) servicefactory.ProviderServiceFactory {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FactoryForModel", arg0)
	ret0, _ := ret[0].(servicefactory.ProviderServiceFactory)
	return ret0
}

// FactoryForModel indicates an expected call of FactoryForModel.
func (mr *MockProviderServiceFactoryGetterMockRecorder) FactoryForModel(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FactoryForModel", reflect.TypeOf((*MockProviderServiceFactoryGetter)(nil).FactoryForModel), arg0)
}
