// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/worker/dbaccessor (interfaces: Logger,DBApp,NodeManager,TrackedDB,Hub,Client)

// Package dbaccessor is a generated GoMock package.
package dbaccessor

import (
	context "context"
	sql "database/sql"
	reflect "reflect"

	sqlair "github.com/canonical/sqlair"
	app "github.com/juju/juju/database/app"
	dqlite "github.com/juju/juju/database/dqlite"
	loggo "github.com/juju/loggo"
	gomock "go.uber.org/mock/gomock"
)

// MockLogger is a mock of Logger interface.
type MockLogger struct {
	ctrl     *gomock.Controller
	recorder *MockLoggerMockRecorder
}

// MockLoggerMockRecorder is the mock recorder for MockLogger.
type MockLoggerMockRecorder struct {
	mock *MockLogger
}

// NewMockLogger creates a new mock instance.
func NewMockLogger(ctrl *gomock.Controller) *MockLogger {
	mock := &MockLogger{ctrl: ctrl}
	mock.recorder = &MockLoggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogger) EXPECT() *MockLoggerMockRecorder {
	return m.recorder
}

// Debugf mocks base method.
func (m *MockLogger) Debugf(arg0 string, arg1 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debugf", varargs...)
}

// Debugf indicates an expected call of Debugf.
func (mr *MockLoggerMockRecorder) Debugf(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debugf", reflect.TypeOf((*MockLogger)(nil).Debugf), varargs...)
}

// Errorf mocks base method.
func (m *MockLogger) Errorf(arg0 string, arg1 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Errorf", varargs...)
}

// Errorf indicates an expected call of Errorf.
func (mr *MockLoggerMockRecorder) Errorf(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Errorf", reflect.TypeOf((*MockLogger)(nil).Errorf), varargs...)
}

// Infof mocks base method.
func (m *MockLogger) Infof(arg0 string, arg1 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Infof", varargs...)
}

// Infof indicates an expected call of Infof.
func (mr *MockLoggerMockRecorder) Infof(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Infof", reflect.TypeOf((*MockLogger)(nil).Infof), varargs...)
}

// IsTraceEnabled mocks base method.
func (m *MockLogger) IsTraceEnabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsTraceEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsTraceEnabled indicates an expected call of IsTraceEnabled.
func (mr *MockLoggerMockRecorder) IsTraceEnabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsTraceEnabled", reflect.TypeOf((*MockLogger)(nil).IsTraceEnabled))
}

// Logf mocks base method.
func (m *MockLogger) Logf(arg0 loggo.Level, arg1 string, arg2 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Logf", varargs...)
}

// Logf indicates an expected call of Logf.
func (mr *MockLoggerMockRecorder) Logf(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logf", reflect.TypeOf((*MockLogger)(nil).Logf), varargs...)
}

// Tracef mocks base method.
func (m *MockLogger) Tracef(arg0 string, arg1 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Tracef", varargs...)
}

// Tracef indicates an expected call of Tracef.
func (mr *MockLoggerMockRecorder) Tracef(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tracef", reflect.TypeOf((*MockLogger)(nil).Tracef), varargs...)
}

// Warningf mocks base method.
func (m *MockLogger) Warningf(arg0 string, arg1 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Warningf", varargs...)
}

// Warningf indicates an expected call of Warningf.
func (mr *MockLoggerMockRecorder) Warningf(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warningf", reflect.TypeOf((*MockLogger)(nil).Warningf), varargs...)
}

// MockDBApp is a mock of DBApp interface.
type MockDBApp struct {
	ctrl     *gomock.Controller
	recorder *MockDBAppMockRecorder
}

// MockDBAppMockRecorder is the mock recorder for MockDBApp.
type MockDBAppMockRecorder struct {
	mock *MockDBApp
}

// NewMockDBApp creates a new mock instance.
func NewMockDBApp(ctrl *gomock.Controller) *MockDBApp {
	mock := &MockDBApp{ctrl: ctrl}
	mock.recorder = &MockDBAppMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDBApp) EXPECT() *MockDBAppMockRecorder {
	return m.recorder
}

// Client mocks base method.
func (m *MockDBApp) Client(arg0 context.Context) (Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Client", arg0)
	ret0, _ := ret[0].(Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Client indicates an expected call of Client.
func (mr *MockDBAppMockRecorder) Client(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Client", reflect.TypeOf((*MockDBApp)(nil).Client), arg0)
}

// Close mocks base method.
func (m *MockDBApp) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockDBAppMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockDBApp)(nil).Close))
}

// Handover mocks base method.
func (m *MockDBApp) Handover(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Handover", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Handover indicates an expected call of Handover.
func (mr *MockDBAppMockRecorder) Handover(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handover", reflect.TypeOf((*MockDBApp)(nil).Handover), arg0)
}

// ID mocks base method.
func (m *MockDBApp) ID() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// ID indicates an expected call of ID.
func (mr *MockDBAppMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockDBApp)(nil).ID))
}

// Open mocks base method.
func (m *MockDBApp) Open(arg0 context.Context, arg1 string) (*sql.DB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Open", arg0, arg1)
	ret0, _ := ret[0].(*sql.DB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Open indicates an expected call of Open.
func (mr *MockDBAppMockRecorder) Open(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Open", reflect.TypeOf((*MockDBApp)(nil).Open), arg0, arg1)
}

// Ready mocks base method.
func (m *MockDBApp) Ready(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ready", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Ready indicates an expected call of Ready.
func (mr *MockDBAppMockRecorder) Ready(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ready", reflect.TypeOf((*MockDBApp)(nil).Ready), arg0)
}

// MockNodeManager is a mock of NodeManager interface.
type MockNodeManager struct {
	ctrl     *gomock.Controller
	recorder *MockNodeManagerMockRecorder
}

// MockNodeManagerMockRecorder is the mock recorder for MockNodeManager.
type MockNodeManagerMockRecorder struct {
	mock *MockNodeManager
}

// NewMockNodeManager creates a new mock instance.
func NewMockNodeManager(ctrl *gomock.Controller) *MockNodeManager {
	mock := &MockNodeManager{ctrl: ctrl}
	mock.recorder = &MockNodeManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNodeManager) EXPECT() *MockNodeManagerMockRecorder {
	return m.recorder
}

// ClusterServers mocks base method.
func (m *MockNodeManager) ClusterServers(arg0 context.Context) ([]dqlite.NodeInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterServers", arg0)
	ret0, _ := ret[0].([]dqlite.NodeInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClusterServers indicates an expected call of ClusterServers.
func (mr *MockNodeManagerMockRecorder) ClusterServers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterServers", reflect.TypeOf((*MockNodeManager)(nil).ClusterServers), arg0)
}

// EnsureDataDir mocks base method.
func (m *MockNodeManager) EnsureDataDir() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureDataDir")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnsureDataDir indicates an expected call of EnsureDataDir.
func (mr *MockNodeManagerMockRecorder) EnsureDataDir() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureDataDir", reflect.TypeOf((*MockNodeManager)(nil).EnsureDataDir))
}

// IsBootstrappedNode mocks base method.
func (m *MockNodeManager) IsBootstrappedNode(arg0 context.Context) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsBootstrappedNode", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsBootstrappedNode indicates an expected call of IsBootstrappedNode.
func (mr *MockNodeManagerMockRecorder) IsBootstrappedNode(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsBootstrappedNode", reflect.TypeOf((*MockNodeManager)(nil).IsBootstrappedNode), arg0)
}

// IsExistingNode mocks base method.
func (m *MockNodeManager) IsExistingNode() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsExistingNode")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsExistingNode indicates an expected call of IsExistingNode.
func (mr *MockNodeManagerMockRecorder) IsExistingNode() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsExistingNode", reflect.TypeOf((*MockNodeManager)(nil).IsExistingNode))
}

// SetClusterServers mocks base method.
func (m *MockNodeManager) SetClusterServers(arg0 context.Context, arg1 []dqlite.NodeInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetClusterServers", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetClusterServers indicates an expected call of SetClusterServers.
func (mr *MockNodeManagerMockRecorder) SetClusterServers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetClusterServers", reflect.TypeOf((*MockNodeManager)(nil).SetClusterServers), arg0, arg1)
}

// SetNodeInfo mocks base method.
func (m *MockNodeManager) SetNodeInfo(arg0 dqlite.NodeInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetNodeInfo", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetNodeInfo indicates an expected call of SetNodeInfo.
func (mr *MockNodeManagerMockRecorder) SetNodeInfo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetNodeInfo", reflect.TypeOf((*MockNodeManager)(nil).SetNodeInfo), arg0)
}

// WithAddressOption mocks base method.
func (m *MockNodeManager) WithAddressOption(arg0 string) app.Option {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithAddressOption", arg0)
	ret0, _ := ret[0].(app.Option)
	return ret0
}

// WithAddressOption indicates an expected call of WithAddressOption.
func (mr *MockNodeManagerMockRecorder) WithAddressOption(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithAddressOption", reflect.TypeOf((*MockNodeManager)(nil).WithAddressOption), arg0)
}

// WithClusterOption mocks base method.
func (m *MockNodeManager) WithClusterOption(arg0 []string) app.Option {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithClusterOption", arg0)
	ret0, _ := ret[0].(app.Option)
	return ret0
}

// WithClusterOption indicates an expected call of WithClusterOption.
func (mr *MockNodeManagerMockRecorder) WithClusterOption(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithClusterOption", reflect.TypeOf((*MockNodeManager)(nil).WithClusterOption), arg0)
}

// WithLogFuncOption mocks base method.
func (m *MockNodeManager) WithLogFuncOption() app.Option {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithLogFuncOption")
	ret0, _ := ret[0].(app.Option)
	return ret0
}

// WithLogFuncOption indicates an expected call of WithLogFuncOption.
func (mr *MockNodeManagerMockRecorder) WithLogFuncOption() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithLogFuncOption", reflect.TypeOf((*MockNodeManager)(nil).WithLogFuncOption))
}

// WithTLSOption mocks base method.
func (m *MockNodeManager) WithTLSOption() (app.Option, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithTLSOption")
	ret0, _ := ret[0].(app.Option)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WithTLSOption indicates an expected call of WithTLSOption.
func (mr *MockNodeManagerMockRecorder) WithTLSOption() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithTLSOption", reflect.TypeOf((*MockNodeManager)(nil).WithTLSOption))
}

// WithTracingOption mocks base method.
func (m *MockNodeManager) WithTracingOption() app.Option {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithTracingOption")
	ret0, _ := ret[0].(app.Option)
	return ret0
}

// WithTracingOption indicates an expected call of WithTracingOption.
func (mr *MockNodeManagerMockRecorder) WithTracingOption() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithTracingOption", reflect.TypeOf((*MockNodeManager)(nil).WithTracingOption))
}

// MockTrackedDB is a mock of TrackedDB interface.
type MockTrackedDB struct {
	ctrl     *gomock.Controller
	recorder *MockTrackedDBMockRecorder
}

// MockTrackedDBMockRecorder is the mock recorder for MockTrackedDB.
type MockTrackedDBMockRecorder struct {
	mock *MockTrackedDB
}

// NewMockTrackedDB creates a new mock instance.
func NewMockTrackedDB(ctrl *gomock.Controller) *MockTrackedDB {
	mock := &MockTrackedDB{ctrl: ctrl}
	mock.recorder = &MockTrackedDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTrackedDB) EXPECT() *MockTrackedDBMockRecorder {
	return m.recorder
}

// Kill mocks base method.
func (m *MockTrackedDB) Kill() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Kill")
}

// Kill indicates an expected call of Kill.
func (mr *MockTrackedDBMockRecorder) Kill() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Kill", reflect.TypeOf((*MockTrackedDB)(nil).Kill))
}

// StdTxn mocks base method.
func (m *MockTrackedDB) StdTxn(arg0 context.Context, arg1 func(context.Context, *sql.Tx) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StdTxn", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// StdTxn indicates an expected call of StdTxn.
func (mr *MockTrackedDBMockRecorder) StdTxn(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StdTxn", reflect.TypeOf((*MockTrackedDB)(nil).StdTxn), arg0, arg1)
}

// Txn mocks base method.
func (m *MockTrackedDB) Txn(arg0 context.Context, arg1 func(context.Context, *sqlair.TX) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Txn", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Txn indicates an expected call of Txn.
func (mr *MockTrackedDBMockRecorder) Txn(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Txn", reflect.TypeOf((*MockTrackedDB)(nil).Txn), arg0, arg1)
}

// Wait mocks base method.
func (m *MockTrackedDB) Wait() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Wait")
	ret0, _ := ret[0].(error)
	return ret0
}

// Wait indicates an expected call of Wait.
func (mr *MockTrackedDBMockRecorder) Wait() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Wait", reflect.TypeOf((*MockTrackedDB)(nil).Wait))
}

// MockHub is a mock of Hub interface.
type MockHub struct {
	ctrl     *gomock.Controller
	recorder *MockHubMockRecorder
}

// MockHubMockRecorder is the mock recorder for MockHub.
type MockHubMockRecorder struct {
	mock *MockHub
}

// NewMockHub creates a new mock instance.
func NewMockHub(ctrl *gomock.Controller) *MockHub {
	mock := &MockHub{ctrl: ctrl}
	mock.recorder = &MockHubMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHub) EXPECT() *MockHubMockRecorder {
	return m.recorder
}

// Publish mocks base method.
func (m *MockHub) Publish(arg0 string, arg1 interface{}) (func(), error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Publish", arg0, arg1)
	ret0, _ := ret[0].(func())
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Publish indicates an expected call of Publish.
func (mr *MockHubMockRecorder) Publish(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockHub)(nil).Publish), arg0, arg1)
}

// Subscribe mocks base method.
func (m *MockHub) Subscribe(arg0 string, arg1 interface{}) (func(), error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscribe", arg0, arg1)
	ret0, _ := ret[0].(func())
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Subscribe indicates an expected call of Subscribe.
func (mr *MockHubMockRecorder) Subscribe(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockHub)(nil).Subscribe), arg0, arg1)
}

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// Cluster mocks base method.
func (m *MockClient) Cluster(arg0 context.Context) ([]dqlite.NodeInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cluster", arg0)
	ret0, _ := ret[0].([]dqlite.NodeInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Cluster indicates an expected call of Cluster.
func (mr *MockClientMockRecorder) Cluster(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cluster", reflect.TypeOf((*MockClient)(nil).Cluster), arg0)
}

// Leader mocks base method.
func (m *MockClient) Leader(arg0 context.Context) (*dqlite.NodeInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Leader", arg0)
	ret0, _ := ret[0].(*dqlite.NodeInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Leader indicates an expected call of Leader.
func (mr *MockClientMockRecorder) Leader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Leader", reflect.TypeOf((*MockClient)(nil).Leader), arg0)
}
