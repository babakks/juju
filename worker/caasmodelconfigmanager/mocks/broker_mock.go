// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/worker/caasmodelconfigmanager (interfaces: CAASBroker)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	docker "github.com/juju/juju/internal/docker"
	gomock "go.uber.org/mock/gomock"
)

// MockCAASBroker is a mock of CAASBroker interface.
type MockCAASBroker struct {
	ctrl     *gomock.Controller
	recorder *MockCAASBrokerMockRecorder
}

// MockCAASBrokerMockRecorder is the mock recorder for MockCAASBroker.
type MockCAASBrokerMockRecorder struct {
	mock *MockCAASBroker
}

// NewMockCAASBroker creates a new mock instance.
func NewMockCAASBroker(ctrl *gomock.Controller) *MockCAASBroker {
	mock := &MockCAASBroker{ctrl: ctrl}
	mock.recorder = &MockCAASBrokerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCAASBroker) EXPECT() *MockCAASBrokerMockRecorder {
	return m.recorder
}

// EnsureImageRepoSecret mocks base method.
func (m *MockCAASBroker) EnsureImageRepoSecret(arg0 docker.ImageRepoDetails) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureImageRepoSecret", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureImageRepoSecret indicates an expected call of EnsureImageRepoSecret.
func (mr *MockCAASBrokerMockRecorder) EnsureImageRepoSecret(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureImageRepoSecret", reflect.TypeOf((*MockCAASBroker)(nil).EnsureImageRepoSecret), arg0)
}
