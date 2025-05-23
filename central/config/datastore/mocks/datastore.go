// Code generated by MockGen. DO NOT EDIT.
// Source: datastore.go
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/datastore.go -source datastore.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	storage "github.com/stackrox/rox/generated/storage"
	gomock "go.uber.org/mock/gomock"
)

// MockDataStore is a mock of DataStore interface.
type MockDataStore struct {
	ctrl     *gomock.Controller
	recorder *MockDataStoreMockRecorder
	isgomock struct{}
}

// MockDataStoreMockRecorder is the mock recorder for MockDataStore.
type MockDataStoreMockRecorder struct {
	mock *MockDataStore
}

// NewMockDataStore creates a new mock instance.
func NewMockDataStore(ctrl *gomock.Controller) *MockDataStore {
	mock := &MockDataStore{ctrl: ctrl}
	mock.recorder = &MockDataStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataStore) EXPECT() *MockDataStoreMockRecorder {
	return m.recorder
}

// GetConfig mocks base method.
func (m *MockDataStore) GetConfig(arg0 context.Context) (*storage.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfig", arg0)
	ret0, _ := ret[0].(*storage.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConfig indicates an expected call of GetConfig.
func (mr *MockDataStoreMockRecorder) GetConfig(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfig", reflect.TypeOf((*MockDataStore)(nil).GetConfig), arg0)
}

// GetPrivateConfig mocks base method.
func (m *MockDataStore) GetPrivateConfig(arg0 context.Context) (*storage.PrivateConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPrivateConfig", arg0)
	ret0, _ := ret[0].(*storage.PrivateConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPrivateConfig indicates an expected call of GetPrivateConfig.
func (mr *MockDataStoreMockRecorder) GetPrivateConfig(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPrivateConfig", reflect.TypeOf((*MockDataStore)(nil).GetPrivateConfig), arg0)
}

// GetPublicConfig mocks base method.
func (m *MockDataStore) GetPublicConfig() (*storage.PublicConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPublicConfig")
	ret0, _ := ret[0].(*storage.PublicConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPublicConfig indicates an expected call of GetPublicConfig.
func (mr *MockDataStoreMockRecorder) GetPublicConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPublicConfig", reflect.TypeOf((*MockDataStore)(nil).GetPublicConfig))
}

// GetVulnerabilityExceptionConfig mocks base method.
func (m *MockDataStore) GetVulnerabilityExceptionConfig(ctx context.Context) (*storage.VulnerabilityExceptionConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVulnerabilityExceptionConfig", ctx)
	ret0, _ := ret[0].(*storage.VulnerabilityExceptionConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVulnerabilityExceptionConfig indicates an expected call of GetVulnerabilityExceptionConfig.
func (mr *MockDataStoreMockRecorder) GetVulnerabilityExceptionConfig(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVulnerabilityExceptionConfig", reflect.TypeOf((*MockDataStore)(nil).GetVulnerabilityExceptionConfig), ctx)
}

// UpsertConfig mocks base method.
func (m *MockDataStore) UpsertConfig(arg0 context.Context, arg1 *storage.Config) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertConfig", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertConfig indicates an expected call of UpsertConfig.
func (mr *MockDataStoreMockRecorder) UpsertConfig(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertConfig", reflect.TypeOf((*MockDataStore)(nil).UpsertConfig), arg0, arg1)
}
