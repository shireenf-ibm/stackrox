// Code generated by MockGen. DO NOT EDIT.
// Source: waiter.go
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/waiter.go -source waiter.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockWaiter is a mock of Waiter interface.
type MockWaiter[T any] struct {
	ctrl     *gomock.Controller
	recorder *MockWaiterMockRecorder[T]
	isgomock struct{}
}

// MockWaiterMockRecorder is the mock recorder for MockWaiter.
type MockWaiterMockRecorder[T any] struct {
	mock *MockWaiter[T]
}

// NewMockWaiter creates a new mock instance.
func NewMockWaiter[T any](ctrl *gomock.Controller) *MockWaiter[T] {
	mock := &MockWaiter[T]{ctrl: ctrl}
	mock.recorder = &MockWaiterMockRecorder[T]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWaiter[T]) EXPECT() *MockWaiterMockRecorder[T] {
	return m.recorder
}

// Close mocks base method.
func (m *MockWaiter[T]) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockWaiterMockRecorder[T]) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockWaiter[T])(nil).Close))
}

// ID mocks base method.
func (m *MockWaiter[T]) ID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ID indicates an expected call of ID.
func (mr *MockWaiterMockRecorder[T]) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockWaiter[T])(nil).ID))
}

// Wait mocks base method.
func (m *MockWaiter[T]) Wait(ctx context.Context) (T, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Wait", ctx)
	ret0, _ := ret[0].(T)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Wait indicates an expected call of Wait.
func (mr *MockWaiterMockRecorder[T]) Wait(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Wait", reflect.TypeOf((*MockWaiter[T])(nil).Wait), ctx)
}
