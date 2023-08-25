// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/jujuclient (interfaces: ProxyFactory)

// Package jujuclient_test is a generated GoMock package.
package jujuclient_test

import (
	reflect "reflect"

	proxy "github.com/juju/juju/internal/proxy"
	gomock "go.uber.org/mock/gomock"
)

// MockProxyFactory is a mock of ProxyFactory interface.
type MockProxyFactory struct {
	ctrl     *gomock.Controller
	recorder *MockProxyFactoryMockRecorder
}

// MockProxyFactoryMockRecorder is the mock recorder for MockProxyFactory.
type MockProxyFactoryMockRecorder struct {
	mock *MockProxyFactory
}

// NewMockProxyFactory creates a new mock instance.
func NewMockProxyFactory(ctrl *gomock.Controller) *MockProxyFactory {
	mock := &MockProxyFactory{ctrl: ctrl}
	mock.recorder = &MockProxyFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProxyFactory) EXPECT() *MockProxyFactoryMockRecorder {
	return m.recorder
}

// ProxierFromConfig mocks base method.
func (m *MockProxyFactory) ProxierFromConfig(arg0 string, arg1 map[string]interface{}) (proxy.Proxier, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProxierFromConfig", arg0, arg1)
	ret0, _ := ret[0].(proxy.Proxier)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProxierFromConfig indicates an expected call of ProxierFromConfig.
func (mr *MockProxyFactoryMockRecorder) ProxierFromConfig(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProxierFromConfig", reflect.TypeOf((*MockProxyFactory)(nil).ProxierFromConfig), arg0, arg1)
}
