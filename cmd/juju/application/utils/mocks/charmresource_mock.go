// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/cmd/juju/application/utils (interfaces: CharmClient)
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/charmresource_mock.go github.com/juju/juju/cmd/juju/application/utils CharmClient
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	charm "github.com/juju/charm/v13"
	resource "github.com/juju/charm/v13/resource"
	charms "github.com/juju/juju/api/common/charms"
	gomock "go.uber.org/mock/gomock"
)

// MockCharmClient is a mock of CharmClient interface.
type MockCharmClient struct {
	ctrl     *gomock.Controller
	recorder *MockCharmClientMockRecorder
}

// MockCharmClientMockRecorder is the mock recorder for MockCharmClient.
type MockCharmClientMockRecorder struct {
	mock *MockCharmClient
}

// NewMockCharmClient creates a new mock instance.
func NewMockCharmClient(ctrl *gomock.Controller) *MockCharmClient {
	mock := &MockCharmClient{ctrl: ctrl}
	mock.recorder = &MockCharmClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCharmClient) EXPECT() *MockCharmClientMockRecorder {
	return m.recorder
}

// CharmInfo mocks base method.
func (m *MockCharmClient) CharmInfo(arg0 string) (*charms.CharmInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CharmInfo", arg0)
	ret0, _ := ret[0].(*charms.CharmInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CharmInfo indicates an expected call of CharmInfo.
func (mr *MockCharmClientMockRecorder) CharmInfo(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CharmInfo", reflect.TypeOf((*MockCharmClient)(nil).CharmInfo), arg0)
}

// ListCharmResources mocks base method.
func (m *MockCharmClient) ListCharmResources(arg0 string, arg1 charm.Origin) ([]resource.Resource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCharmResources", arg0, arg1)
	ret0, _ := ret[0].([]resource.Resource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCharmResources indicates an expected call of ListCharmResources.
func (mr *MockCharmClientMockRecorder) ListCharmResources(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCharmResources", reflect.TypeOf((*MockCharmClient)(nil).ListCharmResources), arg0, arg1)
}
