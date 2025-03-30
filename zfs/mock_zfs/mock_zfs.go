// Code generated by MockGen. DO NOT EDIT.
// Source: zfs.go

// Package mock_zfs is a generated GoMock package.
package mock_zfs

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	zfs "github.com/waitingsong/zfs_exporter/v3/zfs"
)

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

// Datasets mocks base method.
func (m *MockClient) Datasets(pool string, kind zfs.DatasetKind) zfs.Datasets {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Datasets", pool, kind)
	ret0, _ := ret[0].(zfs.Datasets)
	return ret0
}

// Datasets indicates an expected call of Datasets.
func (mr *MockClientMockRecorder) Datasets(pool, kind interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Datasets", reflect.TypeOf((*MockClient)(nil).Datasets), pool, kind)
}

// Pool mocks base method.
func (m *MockClient) Pool(name string) zfs.Pool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pool", name)
	ret0, _ := ret[0].(zfs.Pool)
	return ret0
}

// Pool indicates an expected call of Pool.
func (mr *MockClientMockRecorder) Pool(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pool", reflect.TypeOf((*MockClient)(nil).Pool), name)
}

// PoolNames mocks base method.
func (m *MockClient) PoolNames() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PoolNames")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PoolNames indicates an expected call of PoolNames.
func (mr *MockClientMockRecorder) PoolNames() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PoolNames", reflect.TypeOf((*MockClient)(nil).PoolNames))
}

// MockPool is a mock of Pool interface.
type MockPool struct {
	ctrl     *gomock.Controller
	recorder *MockPoolMockRecorder
}

// MockPoolMockRecorder is the mock recorder for MockPool.
type MockPoolMockRecorder struct {
	mock *MockPool
}

// NewMockPool creates a new mock instance.
func NewMockPool(ctrl *gomock.Controller) *MockPool {
	mock := &MockPool{ctrl: ctrl}
	mock.recorder = &MockPoolMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPool) EXPECT() *MockPoolMockRecorder {
	return m.recorder
}

// Name mocks base method.
func (m *MockPool) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockPoolMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockPool)(nil).Name))
}

// Properties mocks base method.
func (m *MockPool) Properties(props ...string) (zfs.PoolProperties, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range props {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Properties", varargs...)
	ret0, _ := ret[0].(zfs.PoolProperties)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Properties indicates an expected call of Properties.
func (mr *MockPoolMockRecorder) Properties(props ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Properties", reflect.TypeOf((*MockPool)(nil).Properties), props...)
}

// MockPoolProperties is a mock of PoolProperties interface.
type MockPoolProperties struct {
	ctrl     *gomock.Controller
	recorder *MockPoolPropertiesMockRecorder
}

// MockPoolPropertiesMockRecorder is the mock recorder for MockPoolProperties.
type MockPoolPropertiesMockRecorder struct {
	mock *MockPoolProperties
}

// NewMockPoolProperties creates a new mock instance.
func NewMockPoolProperties(ctrl *gomock.Controller) *MockPoolProperties {
	mock := &MockPoolProperties{ctrl: ctrl}
	mock.recorder = &MockPoolPropertiesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPoolProperties) EXPECT() *MockPoolPropertiesMockRecorder {
	return m.recorder
}

// Properties mocks base method.
func (m *MockPoolProperties) Properties() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Properties")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// Properties indicates an expected call of Properties.
func (mr *MockPoolPropertiesMockRecorder) Properties() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Properties", reflect.TypeOf((*MockPoolProperties)(nil).Properties))
}

// MockDatasets is a mock of Datasets interface.
type MockDatasets struct {
	ctrl     *gomock.Controller
	recorder *MockDatasetsMockRecorder
}

// MockDatasetsMockRecorder is the mock recorder for MockDatasets.
type MockDatasetsMockRecorder struct {
	mock *MockDatasets
}

// NewMockDatasets creates a new mock instance.
func NewMockDatasets(ctrl *gomock.Controller) *MockDatasets {
	mock := &MockDatasets{ctrl: ctrl}
	mock.recorder = &MockDatasetsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDatasets) EXPECT() *MockDatasetsMockRecorder {
	return m.recorder
}

// Kind mocks base method.
func (m *MockDatasets) Kind() zfs.DatasetKind {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Kind")
	ret0, _ := ret[0].(zfs.DatasetKind)
	return ret0
}

// Kind indicates an expected call of Kind.
func (mr *MockDatasetsMockRecorder) Kind() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Kind", reflect.TypeOf((*MockDatasets)(nil).Kind))
}

// Pool mocks base method.
func (m *MockDatasets) Pool() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pool")
	ret0, _ := ret[0].(string)
	return ret0
}

// Pool indicates an expected call of Pool.
func (mr *MockDatasetsMockRecorder) Pool() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pool", reflect.TypeOf((*MockDatasets)(nil).Pool))
}

// Properties mocks base method.
func (m *MockDatasets) Properties(props ...string) ([]zfs.DatasetProperties, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range props {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Properties", varargs...)
	ret0, _ := ret[0].([]zfs.DatasetProperties)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Properties indicates an expected call of Properties.
func (mr *MockDatasetsMockRecorder) Properties(props ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Properties", reflect.TypeOf((*MockDatasets)(nil).Properties), props...)
}

// MockDatasetProperties is a mock of DatasetProperties interface.
type MockDatasetProperties struct {
	ctrl     *gomock.Controller
	recorder *MockDatasetPropertiesMockRecorder
}

// MockDatasetPropertiesMockRecorder is the mock recorder for MockDatasetProperties.
type MockDatasetPropertiesMockRecorder struct {
	mock *MockDatasetProperties
}

// NewMockDatasetProperties creates a new mock instance.
func NewMockDatasetProperties(ctrl *gomock.Controller) *MockDatasetProperties {
	mock := &MockDatasetProperties{ctrl: ctrl}
	mock.recorder = &MockDatasetPropertiesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDatasetProperties) EXPECT() *MockDatasetPropertiesMockRecorder {
	return m.recorder
}

// DatasetName mocks base method.
func (m *MockDatasetProperties) DatasetName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DatasetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// DatasetName indicates an expected call of DatasetName.
func (mr *MockDatasetPropertiesMockRecorder) DatasetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DatasetName", reflect.TypeOf((*MockDatasetProperties)(nil).DatasetName))
}

// Properties mocks base method.
func (m *MockDatasetProperties) Properties() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Properties")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// Properties indicates an expected call of Properties.
func (mr *MockDatasetPropertiesMockRecorder) Properties() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Properties", reflect.TypeOf((*MockDatasetProperties)(nil).Properties))
}

// Mockhandler is a mock of handler interface.
type Mockhandler struct {
	ctrl     *gomock.Controller
	recorder *MockhandlerMockRecorder
}

// MockhandlerMockRecorder is the mock recorder for Mockhandler.
type MockhandlerMockRecorder struct {
	mock *Mockhandler
}

// NewMockhandler creates a new mock instance.
func NewMockhandler(ctrl *gomock.Controller) *Mockhandler {
	mock := &Mockhandler{ctrl: ctrl}
	mock.recorder = &MockhandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockhandler) EXPECT() *MockhandlerMockRecorder {
	return m.recorder
}

// processLine mocks base method.
func (m *Mockhandler) processLine(pool string, line []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "processLine", pool, line)
	ret0, _ := ret[0].(error)
	return ret0
}

// processLine indicates an expected call of processLine.
func (mr *MockhandlerMockRecorder) processLine(pool, line interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "processLine", reflect.TypeOf((*Mockhandler)(nil).processLine), pool, line)
}
