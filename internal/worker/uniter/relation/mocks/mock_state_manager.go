// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/internal/worker/uniter/relation (interfaces: StateManager)
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/mock_state_manager.go github.com/juju/juju/internal/worker/uniter/relation StateManager
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	relation "github.com/juju/juju/internal/worker/uniter/relation"
	gomock "go.uber.org/mock/gomock"
)

// MockStateManager is a mock of StateManager interface.
type MockStateManager struct {
	ctrl     *gomock.Controller
	recorder *MockStateManagerMockRecorder
}

// MockStateManagerMockRecorder is the mock recorder for MockStateManager.
type MockStateManagerMockRecorder struct {
	mock *MockStateManager
}

// NewMockStateManager creates a new mock instance.
func NewMockStateManager(ctrl *gomock.Controller) *MockStateManager {
	mock := &MockStateManager{ctrl: ctrl}
	mock.recorder = &MockStateManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStateManager) EXPECT() *MockStateManagerMockRecorder {
	return m.recorder
}

// KnownIDs mocks base method.
func (m *MockStateManager) KnownIDs() []int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KnownIDs")
	ret0, _ := ret[0].([]int)
	return ret0
}

// KnownIDs indicates an expected call of KnownIDs.
func (mr *MockStateManagerMockRecorder) KnownIDs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KnownIDs", reflect.TypeOf((*MockStateManager)(nil).KnownIDs))
}

// Relation mocks base method.
func (m *MockStateManager) Relation(arg0 int) (*relation.State, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Relation", arg0)
	ret0, _ := ret[0].(*relation.State)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Relation indicates an expected call of Relation.
func (mr *MockStateManagerMockRecorder) Relation(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Relation", reflect.TypeOf((*MockStateManager)(nil).Relation), arg0)
}

// RelationFound mocks base method.
func (m *MockStateManager) RelationFound(arg0 int) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RelationFound", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// RelationFound indicates an expected call of RelationFound.
func (mr *MockStateManagerMockRecorder) RelationFound(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RelationFound", reflect.TypeOf((*MockStateManager)(nil).RelationFound), arg0)
}

// RemoveRelation mocks base method.
func (m *MockStateManager) RemoveRelation(arg0 context.Context, arg1 int, arg2 relation.UnitGetter, arg3 map[string]bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveRelation", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveRelation indicates an expected call of RemoveRelation.
func (mr *MockStateManagerMockRecorder) RemoveRelation(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveRelation", reflect.TypeOf((*MockStateManager)(nil).RemoveRelation), arg0, arg1, arg2, arg3)
}

// SetRelation mocks base method.
func (m *MockStateManager) SetRelation(arg0 *relation.State) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetRelation", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetRelation indicates an expected call of SetRelation.
func (mr *MockStateManagerMockRecorder) SetRelation(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRelation", reflect.TypeOf((*MockStateManager)(nil).SetRelation), arg0)
}
