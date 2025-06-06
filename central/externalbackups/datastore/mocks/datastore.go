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

// ForEachBackup mocks base method.
func (m *MockDataStore) ForEachBackup(ctx context.Context, fn func(*storage.ExternalBackup) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForEachBackup", ctx, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// ForEachBackup indicates an expected call of ForEachBackup.
func (mr *MockDataStoreMockRecorder) ForEachBackup(ctx, fn any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForEachBackup", reflect.TypeOf((*MockDataStore)(nil).ForEachBackup), ctx, fn)
}

// GetBackup mocks base method.
func (m *MockDataStore) GetBackup(ctx context.Context, id string) (*storage.ExternalBackup, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBackup", ctx, id)
	ret0, _ := ret[0].(*storage.ExternalBackup)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetBackup indicates an expected call of GetBackup.
func (mr *MockDataStoreMockRecorder) GetBackup(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBackup", reflect.TypeOf((*MockDataStore)(nil).GetBackup), ctx, id)
}

// RemoveBackup mocks base method.
func (m *MockDataStore) RemoveBackup(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveBackup", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveBackup indicates an expected call of RemoveBackup.
func (mr *MockDataStoreMockRecorder) RemoveBackup(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveBackup", reflect.TypeOf((*MockDataStore)(nil).RemoveBackup), ctx, id)
}

// UpsertBackup mocks base method.
func (m *MockDataStore) UpsertBackup(ctx context.Context, backup *storage.ExternalBackup) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertBackup", ctx, backup)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertBackup indicates an expected call of UpsertBackup.
func (mr *MockDataStoreMockRecorder) UpsertBackup(ctx, backup any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertBackup", reflect.TypeOf((*MockDataStore)(nil).UpsertBackup), ctx, backup)
}
