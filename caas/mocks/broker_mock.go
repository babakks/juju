// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/caas (interfaces: Broker)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	caas "github.com/juju/juju/caas"
	constraints "github.com/juju/juju/core/constraints"
	secrets "github.com/juju/juju/core/secrets"
	docker "github.com/juju/juju/docker"
	environs "github.com/juju/juju/environs"
	config "github.com/juju/juju/environs/config"
	context0 "github.com/juju/juju/environs/context"
	proxy "github.com/juju/juju/proxy"
	storage "github.com/juju/juju/storage"
	names "github.com/juju/names/v4"
	version "github.com/juju/version/v2"
	gomock "go.uber.org/mock/gomock"
)

// MockBroker is a mock of Broker interface.
type MockBroker struct {
	ctrl     *gomock.Controller
	recorder *MockBrokerMockRecorder
}

// MockBrokerMockRecorder is the mock recorder for MockBroker.
type MockBrokerMockRecorder struct {
	mock *MockBroker
}

// NewMockBroker creates a new mock instance.
func NewMockBroker(ctrl *gomock.Controller) *MockBroker {
	mock := &MockBroker{ctrl: ctrl}
	mock.recorder = &MockBrokerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBroker) EXPECT() *MockBrokerMockRecorder {
	return m.recorder
}

// APIVersion mocks base method.
func (m *MockBroker) APIVersion() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "APIVersion")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// APIVersion indicates an expected call of APIVersion.
func (mr *MockBrokerMockRecorder) APIVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "APIVersion", reflect.TypeOf((*MockBroker)(nil).APIVersion))
}

// AdoptResources mocks base method.
func (m *MockBroker) AdoptResources(arg0 context0.ProviderCallContext, arg1 string, arg2 version.Number) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AdoptResources", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AdoptResources indicates an expected call of AdoptResources.
func (mr *MockBrokerMockRecorder) AdoptResources(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdoptResources", reflect.TypeOf((*MockBroker)(nil).AdoptResources), arg0, arg1, arg2)
}

// AnnotateUnit mocks base method.
func (m *MockBroker) AnnotateUnit(arg0, arg1 string, arg2 names.UnitTag) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AnnotateUnit", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AnnotateUnit indicates an expected call of AnnotateUnit.
func (mr *MockBrokerMockRecorder) AnnotateUnit(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AnnotateUnit", reflect.TypeOf((*MockBroker)(nil).AnnotateUnit), arg0, arg1, arg2)
}

// Application mocks base method.
func (m *MockBroker) Application(arg0 string, arg1 caas.DeploymentType) caas.Application {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Application", arg0, arg1)
	ret0, _ := ret[0].(caas.Application)
	return ret0
}

// Application indicates an expected call of Application.
func (mr *MockBrokerMockRecorder) Application(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Application", reflect.TypeOf((*MockBroker)(nil).Application), arg0, arg1)
}

// Bootstrap mocks base method.
func (m *MockBroker) Bootstrap(arg0 environs.BootstrapContext, arg1 context0.ProviderCallContext, arg2 environs.BootstrapParams) (*environs.BootstrapResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Bootstrap", arg0, arg1, arg2)
	ret0, _ := ret[0].(*environs.BootstrapResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Bootstrap indicates an expected call of Bootstrap.
func (mr *MockBrokerMockRecorder) Bootstrap(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bootstrap", reflect.TypeOf((*MockBroker)(nil).Bootstrap), arg0, arg1, arg2)
}

// CheckCloudCredentials mocks base method.
func (m *MockBroker) CheckCloudCredentials() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckCloudCredentials")
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckCloudCredentials indicates an expected call of CheckCloudCredentials.
func (mr *MockBrokerMockRecorder) CheckCloudCredentials() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckCloudCredentials", reflect.TypeOf((*MockBroker)(nil).CheckCloudCredentials))
}

// Config mocks base method.
func (m *MockBroker) Config() *config.Config {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Config")
	ret0, _ := ret[0].(*config.Config)
	return ret0
}

// Config indicates an expected call of Config.
func (mr *MockBrokerMockRecorder) Config() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Config", reflect.TypeOf((*MockBroker)(nil).Config))
}

// ConstraintsValidator mocks base method.
func (m *MockBroker) ConstraintsValidator(arg0 context0.ProviderCallContext) (constraints.Validator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConstraintsValidator", arg0)
	ret0, _ := ret[0].(constraints.Validator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConstraintsValidator indicates an expected call of ConstraintsValidator.
func (mr *MockBrokerMockRecorder) ConstraintsValidator(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConstraintsValidator", reflect.TypeOf((*MockBroker)(nil).ConstraintsValidator), arg0)
}

// Create mocks base method.
func (m *MockBroker) Create(arg0 context0.ProviderCallContext, arg1 environs.CreateParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockBrokerMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockBroker)(nil).Create), arg0, arg1)
}

// DeleteJujuSecret mocks base method.
func (m *MockBroker) DeleteJujuSecret(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteJujuSecret", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteJujuSecret indicates an expected call of DeleteJujuSecret.
func (mr *MockBrokerMockRecorder) DeleteJujuSecret(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteJujuSecret", reflect.TypeOf((*MockBroker)(nil).DeleteJujuSecret), arg0, arg1)
}

// Destroy mocks base method.
func (m *MockBroker) Destroy(arg0 context0.ProviderCallContext) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Destroy", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Destroy indicates an expected call of Destroy.
func (mr *MockBrokerMockRecorder) Destroy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Destroy", reflect.TypeOf((*MockBroker)(nil).Destroy), arg0)
}

// DestroyController mocks base method.
func (m *MockBroker) DestroyController(arg0 context0.ProviderCallContext, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DestroyController", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DestroyController indicates an expected call of DestroyController.
func (mr *MockBrokerMockRecorder) DestroyController(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DestroyController", reflect.TypeOf((*MockBroker)(nil).DestroyController), arg0, arg1)
}

// EnsureImageRepoSecret mocks base method.
func (m *MockBroker) EnsureImageRepoSecret(arg0 docker.ImageRepoDetails) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureImageRepoSecret", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureImageRepoSecret indicates an expected call of EnsureImageRepoSecret.
func (mr *MockBrokerMockRecorder) EnsureImageRepoSecret(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureImageRepoSecret", reflect.TypeOf((*MockBroker)(nil).EnsureImageRepoSecret), arg0)
}

// EnsureModelOperator mocks base method.
func (m *MockBroker) EnsureModelOperator(arg0, arg1 string, arg2 *caas.ModelOperatorConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureModelOperator", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureModelOperator indicates an expected call of EnsureModelOperator.
func (mr *MockBrokerMockRecorder) EnsureModelOperator(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureModelOperator", reflect.TypeOf((*MockBroker)(nil).EnsureModelOperator), arg0, arg1, arg2)
}

// EnsureSecretAccessToken mocks base method.
func (m *MockBroker) EnsureSecretAccessToken(arg0 names.Tag, arg1, arg2, arg3 []string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureSecretAccessToken", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnsureSecretAccessToken indicates an expected call of EnsureSecretAccessToken.
func (mr *MockBrokerMockRecorder) EnsureSecretAccessToken(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureSecretAccessToken", reflect.TypeOf((*MockBroker)(nil).EnsureSecretAccessToken), arg0, arg1, arg2, arg3)
}

// GetJujuSecret mocks base method.
func (m *MockBroker) GetJujuSecret(arg0 context.Context, arg1 string) (secrets.SecretValue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetJujuSecret", arg0, arg1)
	ret0, _ := ret[0].(secrets.SecretValue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJujuSecret indicates an expected call of GetJujuSecret.
func (mr *MockBrokerMockRecorder) GetJujuSecret(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJujuSecret", reflect.TypeOf((*MockBroker)(nil).GetJujuSecret), arg0, arg1)
}

// GetSecretToken mocks base method.
func (m *MockBroker) GetSecretToken(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecretToken", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecretToken indicates an expected call of GetSecretToken.
func (mr *MockBrokerMockRecorder) GetSecretToken(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecretToken", reflect.TypeOf((*MockBroker)(nil).GetSecretToken), arg0)
}

// GetService mocks base method.
func (m *MockBroker) GetService(arg0 string, arg1 bool) (*caas.Service, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetService", arg0, arg1)
	ret0, _ := ret[0].(*caas.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetService indicates an expected call of GetService.
func (mr *MockBrokerMockRecorder) GetService(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetService", reflect.TypeOf((*MockBroker)(nil).GetService), arg0, arg1)
}

// ModelOperator mocks base method.
func (m *MockBroker) ModelOperator() (*caas.ModelOperatorConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelOperator")
	ret0, _ := ret[0].(*caas.ModelOperatorConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModelOperator indicates an expected call of ModelOperator.
func (mr *MockBrokerMockRecorder) ModelOperator() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelOperator", reflect.TypeOf((*MockBroker)(nil).ModelOperator))
}

// ModelOperatorExists mocks base method.
func (m *MockBroker) ModelOperatorExists() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelOperatorExists")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModelOperatorExists indicates an expected call of ModelOperatorExists.
func (mr *MockBrokerMockRecorder) ModelOperatorExists() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelOperatorExists", reflect.TypeOf((*MockBroker)(nil).ModelOperatorExists))
}

// PrecheckInstance mocks base method.
func (m *MockBroker) PrecheckInstance(arg0 context0.ProviderCallContext, arg1 environs.PrecheckInstanceParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrecheckInstance", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// PrecheckInstance indicates an expected call of PrecheckInstance.
func (mr *MockBrokerMockRecorder) PrecheckInstance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrecheckInstance", reflect.TypeOf((*MockBroker)(nil).PrecheckInstance), arg0, arg1)
}

// PrepareForBootstrap mocks base method.
func (m *MockBroker) PrepareForBootstrap(arg0 environs.BootstrapContext, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrepareForBootstrap", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// PrepareForBootstrap indicates an expected call of PrepareForBootstrap.
func (mr *MockBrokerMockRecorder) PrepareForBootstrap(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareForBootstrap", reflect.TypeOf((*MockBroker)(nil).PrepareForBootstrap), arg0, arg1)
}

// Provider mocks base method.
func (m *MockBroker) Provider() caas.ContainerEnvironProvider {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Provider")
	ret0, _ := ret[0].(caas.ContainerEnvironProvider)
	return ret0
}

// Provider indicates an expected call of Provider.
func (mr *MockBrokerMockRecorder) Provider() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Provider", reflect.TypeOf((*MockBroker)(nil).Provider))
}

// ProxyToApplication mocks base method.
func (m *MockBroker) ProxyToApplication(arg0, arg1 string) (proxy.Proxier, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProxyToApplication", arg0, arg1)
	ret0, _ := ret[0].(proxy.Proxier)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProxyToApplication indicates an expected call of ProxyToApplication.
func (mr *MockBrokerMockRecorder) ProxyToApplication(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProxyToApplication", reflect.TypeOf((*MockBroker)(nil).ProxyToApplication), arg0, arg1)
}

// RemoveSecretAccessToken mocks base method.
func (m *MockBroker) RemoveSecretAccessToken(arg0 names.Tag) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveSecretAccessToken", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveSecretAccessToken indicates an expected call of RemoveSecretAccessToken.
func (mr *MockBrokerMockRecorder) RemoveSecretAccessToken(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveSecretAccessToken", reflect.TypeOf((*MockBroker)(nil).RemoveSecretAccessToken), arg0)
}

// SaveJujuSecret mocks base method.
func (m *MockBroker) SaveJujuSecret(arg0 context.Context, arg1 string, arg2 secrets.SecretValue) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveJujuSecret", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveJujuSecret indicates an expected call of SaveJujuSecret.
func (mr *MockBrokerMockRecorder) SaveJujuSecret(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveJujuSecret", reflect.TypeOf((*MockBroker)(nil).SaveJujuSecret), arg0, arg1, arg2)
}

// SetConfig mocks base method.
func (m *MockBroker) SetConfig(arg0 *config.Config) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetConfig", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetConfig indicates an expected call of SetConfig.
func (mr *MockBrokerMockRecorder) SetConfig(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetConfig", reflect.TypeOf((*MockBroker)(nil).SetConfig), arg0)
}

// StorageProvider mocks base method.
func (m *MockBroker) StorageProvider(arg0 storage.ProviderType) (storage.Provider, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StorageProvider", arg0)
	ret0, _ := ret[0].(storage.Provider)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StorageProvider indicates an expected call of StorageProvider.
func (mr *MockBrokerMockRecorder) StorageProvider(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StorageProvider", reflect.TypeOf((*MockBroker)(nil).StorageProvider), arg0)
}

// StorageProviderTypes mocks base method.
func (m *MockBroker) StorageProviderTypes() ([]storage.ProviderType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StorageProviderTypes")
	ret0, _ := ret[0].([]storage.ProviderType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StorageProviderTypes indicates an expected call of StorageProviderTypes.
func (mr *MockBrokerMockRecorder) StorageProviderTypes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StorageProviderTypes", reflect.TypeOf((*MockBroker)(nil).StorageProviderTypes))
}

// Units mocks base method.
func (m *MockBroker) Units(arg0 string) ([]caas.Unit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Units", arg0)
	ret0, _ := ret[0].([]caas.Unit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Units indicates an expected call of Units.
func (mr *MockBrokerMockRecorder) Units(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Units", reflect.TypeOf((*MockBroker)(nil).Units), arg0)
}

// Upgrade mocks base method.
func (m *MockBroker) Upgrade(arg0 string, arg1 version.Number) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upgrade", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upgrade indicates an expected call of Upgrade.
func (mr *MockBrokerMockRecorder) Upgrade(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upgrade", reflect.TypeOf((*MockBroker)(nil).Upgrade), arg0, arg1)
}

// ValidateStorageClass mocks base method.
func (m *MockBroker) ValidateStorageClass(arg0 map[string]interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateStorageClass", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateStorageClass indicates an expected call of ValidateStorageClass.
func (mr *MockBrokerMockRecorder) ValidateStorageClass(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateStorageClass", reflect.TypeOf((*MockBroker)(nil).ValidateStorageClass), arg0)
}

// Version mocks base method.
func (m *MockBroker) Version() (*version.Number, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Version")
	ret0, _ := ret[0].(*version.Number)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Version indicates an expected call of Version.
func (mr *MockBrokerMockRecorder) Version() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*MockBroker)(nil).Version))
}
