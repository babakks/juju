// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/environs (interfaces: Environ,PrecheckJujuUpgradeStep)

// Package commands is a generated GoMock package.
package commands

import (
	reflect "reflect"

	constraints "github.com/juju/juju/core/constraints"
	instance "github.com/juju/juju/core/instance"
	environs "github.com/juju/juju/environs"
	config "github.com/juju/juju/environs/config"
	context "github.com/juju/juju/environs/context"
	instances "github.com/juju/juju/environs/instances"
	storage "github.com/juju/juju/internal/storage"
	version "github.com/juju/version/v2"
	gomock "go.uber.org/mock/gomock"
)

// MockEnviron is a mock of Environ interface.
type MockEnviron struct {
	ctrl     *gomock.Controller
	recorder *MockEnvironMockRecorder
}

// MockEnvironMockRecorder is the mock recorder for MockEnviron.
type MockEnvironMockRecorder struct {
	mock *MockEnviron
}

// NewMockEnviron creates a new mock instance.
func NewMockEnviron(ctrl *gomock.Controller) *MockEnviron {
	mock := &MockEnviron{ctrl: ctrl}
	mock.recorder = &MockEnvironMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEnviron) EXPECT() *MockEnvironMockRecorder {
	return m.recorder
}

// AdoptResources mocks base method.
func (m *MockEnviron) AdoptResources(arg0 context.ProviderCallContext, arg1 string, arg2 version.Number) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AdoptResources", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AdoptResources indicates an expected call of AdoptResources.
func (mr *MockEnvironMockRecorder) AdoptResources(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdoptResources", reflect.TypeOf((*MockEnviron)(nil).AdoptResources), arg0, arg1, arg2)
}

// AllInstances mocks base method.
func (m *MockEnviron) AllInstances(arg0 context.ProviderCallContext) ([]instances.Instance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllInstances", arg0)
	ret0, _ := ret[0].([]instances.Instance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllInstances indicates an expected call of AllInstances.
func (mr *MockEnvironMockRecorder) AllInstances(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllInstances", reflect.TypeOf((*MockEnviron)(nil).AllInstances), arg0)
}

// AllRunningInstances mocks base method.
func (m *MockEnviron) AllRunningInstances(arg0 context.ProviderCallContext) ([]instances.Instance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllRunningInstances", arg0)
	ret0, _ := ret[0].([]instances.Instance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllRunningInstances indicates an expected call of AllRunningInstances.
func (mr *MockEnvironMockRecorder) AllRunningInstances(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllRunningInstances", reflect.TypeOf((*MockEnviron)(nil).AllRunningInstances), arg0)
}

// Bootstrap mocks base method.
func (m *MockEnviron) Bootstrap(arg0 environs.BootstrapContext, arg1 context.ProviderCallContext, arg2 environs.BootstrapParams) (*environs.BootstrapResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Bootstrap", arg0, arg1, arg2)
	ret0, _ := ret[0].(*environs.BootstrapResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Bootstrap indicates an expected call of Bootstrap.
func (mr *MockEnvironMockRecorder) Bootstrap(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bootstrap", reflect.TypeOf((*MockEnviron)(nil).Bootstrap), arg0, arg1, arg2)
}

// Config mocks base method.
func (m *MockEnviron) Config() *config.Config {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Config")
	ret0, _ := ret[0].(*config.Config)
	return ret0
}

// Config indicates an expected call of Config.
func (mr *MockEnvironMockRecorder) Config() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Config", reflect.TypeOf((*MockEnviron)(nil).Config))
}

// ConstraintsValidator mocks base method.
func (m *MockEnviron) ConstraintsValidator(arg0 context.ProviderCallContext) (constraints.Validator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConstraintsValidator", arg0)
	ret0, _ := ret[0].(constraints.Validator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConstraintsValidator indicates an expected call of ConstraintsValidator.
func (mr *MockEnvironMockRecorder) ConstraintsValidator(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConstraintsValidator", reflect.TypeOf((*MockEnviron)(nil).ConstraintsValidator), arg0)
}

// ControllerInstances mocks base method.
func (m *MockEnviron) ControllerInstances(arg0 context.ProviderCallContext, arg1 string) ([]instance.Id, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerInstances", arg0, arg1)
	ret0, _ := ret[0].([]instance.Id)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ControllerInstances indicates an expected call of ControllerInstances.
func (mr *MockEnvironMockRecorder) ControllerInstances(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerInstances", reflect.TypeOf((*MockEnviron)(nil).ControllerInstances), arg0, arg1)
}

// Create mocks base method.
func (m *MockEnviron) Create(arg0 context.ProviderCallContext, arg1 environs.CreateParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockEnvironMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockEnviron)(nil).Create), arg0, arg1)
}

// Destroy mocks base method.
func (m *MockEnviron) Destroy(arg0 context.ProviderCallContext) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Destroy", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Destroy indicates an expected call of Destroy.
func (mr *MockEnvironMockRecorder) Destroy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Destroy", reflect.TypeOf((*MockEnviron)(nil).Destroy), arg0)
}

// DestroyController mocks base method.
func (m *MockEnviron) DestroyController(arg0 context.ProviderCallContext, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DestroyController", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DestroyController indicates an expected call of DestroyController.
func (mr *MockEnvironMockRecorder) DestroyController(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DestroyController", reflect.TypeOf((*MockEnviron)(nil).DestroyController), arg0, arg1)
}

// InstanceTypes mocks base method.
func (m *MockEnviron) InstanceTypes(arg0 context.ProviderCallContext, arg1 constraints.Value) (instances.InstanceTypesWithCostMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstanceTypes", arg0, arg1)
	ret0, _ := ret[0].(instances.InstanceTypesWithCostMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InstanceTypes indicates an expected call of InstanceTypes.
func (mr *MockEnvironMockRecorder) InstanceTypes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstanceTypes", reflect.TypeOf((*MockEnviron)(nil).InstanceTypes), arg0, arg1)
}

// Instances mocks base method.
func (m *MockEnviron) Instances(arg0 context.ProviderCallContext, arg1 []instance.Id) ([]instances.Instance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Instances", arg0, arg1)
	ret0, _ := ret[0].([]instances.Instance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Instances indicates an expected call of Instances.
func (mr *MockEnvironMockRecorder) Instances(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Instances", reflect.TypeOf((*MockEnviron)(nil).Instances), arg0, arg1)
}

// PrecheckInstance mocks base method.
func (m *MockEnviron) PrecheckInstance(arg0 context.ProviderCallContext, arg1 environs.PrecheckInstanceParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrecheckInstance", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// PrecheckInstance indicates an expected call of PrecheckInstance.
func (mr *MockEnvironMockRecorder) PrecheckInstance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrecheckInstance", reflect.TypeOf((*MockEnviron)(nil).PrecheckInstance), arg0, arg1)
}

// PrepareForBootstrap mocks base method.
func (m *MockEnviron) PrepareForBootstrap(arg0 environs.BootstrapContext, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrepareForBootstrap", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// PrepareForBootstrap indicates an expected call of PrepareForBootstrap.
func (mr *MockEnvironMockRecorder) PrepareForBootstrap(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareForBootstrap", reflect.TypeOf((*MockEnviron)(nil).PrepareForBootstrap), arg0, arg1)
}

// Provider mocks base method.
func (m *MockEnviron) Provider() environs.EnvironProvider {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Provider")
	ret0, _ := ret[0].(environs.EnvironProvider)
	return ret0
}

// Provider indicates an expected call of Provider.
func (mr *MockEnvironMockRecorder) Provider() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Provider", reflect.TypeOf((*MockEnviron)(nil).Provider))
}

// SetConfig mocks base method.
func (m *MockEnviron) SetConfig(arg0 *config.Config) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetConfig", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetConfig indicates an expected call of SetConfig.
func (mr *MockEnvironMockRecorder) SetConfig(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetConfig", reflect.TypeOf((*MockEnviron)(nil).SetConfig), arg0)
}

// StartInstance mocks base method.
func (m *MockEnviron) StartInstance(arg0 context.ProviderCallContext, arg1 environs.StartInstanceParams) (*environs.StartInstanceResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartInstance", arg0, arg1)
	ret0, _ := ret[0].(*environs.StartInstanceResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartInstance indicates an expected call of StartInstance.
func (mr *MockEnvironMockRecorder) StartInstance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartInstance", reflect.TypeOf((*MockEnviron)(nil).StartInstance), arg0, arg1)
}

// StopInstances mocks base method.
func (m *MockEnviron) StopInstances(arg0 context.ProviderCallContext, arg1 ...instance.Id) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StopInstances", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// StopInstances indicates an expected call of StopInstances.
func (mr *MockEnvironMockRecorder) StopInstances(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopInstances", reflect.TypeOf((*MockEnviron)(nil).StopInstances), varargs...)
}

// StorageProvider mocks base method.
func (m *MockEnviron) StorageProvider(arg0 storage.ProviderType) (storage.Provider, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StorageProvider", arg0)
	ret0, _ := ret[0].(storage.Provider)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StorageProvider indicates an expected call of StorageProvider.
func (mr *MockEnvironMockRecorder) StorageProvider(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StorageProvider", reflect.TypeOf((*MockEnviron)(nil).StorageProvider), arg0)
}

// StorageProviderTypes mocks base method.
func (m *MockEnviron) StorageProviderTypes() ([]storage.ProviderType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StorageProviderTypes")
	ret0, _ := ret[0].([]storage.ProviderType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StorageProviderTypes indicates an expected call of StorageProviderTypes.
func (mr *MockEnvironMockRecorder) StorageProviderTypes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StorageProviderTypes", reflect.TypeOf((*MockEnviron)(nil).StorageProviderTypes))
}

// MockPrecheckJujuUpgradeStep is a mock of PrecheckJujuUpgradeStep interface.
type MockPrecheckJujuUpgradeStep struct {
	ctrl     *gomock.Controller
	recorder *MockPrecheckJujuUpgradeStepMockRecorder
}

// MockPrecheckJujuUpgradeStepMockRecorder is the mock recorder for MockPrecheckJujuUpgradeStep.
type MockPrecheckJujuUpgradeStepMockRecorder struct {
	mock *MockPrecheckJujuUpgradeStep
}

// NewMockPrecheckJujuUpgradeStep creates a new mock instance.
func NewMockPrecheckJujuUpgradeStep(ctrl *gomock.Controller) *MockPrecheckJujuUpgradeStep {
	mock := &MockPrecheckJujuUpgradeStep{ctrl: ctrl}
	mock.recorder = &MockPrecheckJujuUpgradeStepMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPrecheckJujuUpgradeStep) EXPECT() *MockPrecheckJujuUpgradeStepMockRecorder {
	return m.recorder
}

// Description mocks base method.
func (m *MockPrecheckJujuUpgradeStep) Description() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Description")
	ret0, _ := ret[0].(string)
	return ret0
}

// Description indicates an expected call of Description.
func (mr *MockPrecheckJujuUpgradeStepMockRecorder) Description() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Description", reflect.TypeOf((*MockPrecheckJujuUpgradeStep)(nil).Description))
}

// Run mocks base method.
func (m *MockPrecheckJujuUpgradeStep) Run() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run")
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run.
func (mr *MockPrecheckJujuUpgradeStepMockRecorder) Run() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockPrecheckJujuUpgradeStep)(nil).Run))
}
