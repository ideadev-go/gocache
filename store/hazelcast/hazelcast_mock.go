// Code generated by MockGen. DO NOT EDIT.
// Source: store/hazelcast.go

// Package mocks is a generated GoMock package.
package hazelcast

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	types "github.com/hazelcast/hazelcast-go-client/types"
)

// MockHazelcastMapInterface is a mock of HazelcastMapInterface interface.
type MockHazelcastMapInterface struct {
	ctrl     *gomock.Controller
	recorder *MockHazelcastMapInterfaceMockRecorder
}

// MockHazelcastMapInterfaceMockRecorder is the mock recorder for MockHazelcastMapInterface.
type MockHazelcastMapInterfaceMockRecorder struct {
	mock *MockHazelcastMapInterface
}

// NewMockHazelcastMapInterface creates a new mock instance.
func NewMockHazelcastMapInterface(ctrl *gomock.Controller) *MockHazelcastMapInterface {
	mock := &MockHazelcastMapInterface{ctrl: ctrl}
	mock.recorder = &MockHazelcastMapInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHazelcastMapInterface) EXPECT() *MockHazelcastMapInterfaceMockRecorder {
	return m.recorder
}

// Clear mocks base method.
func (m *MockHazelcastMapInterface) Clear(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clear", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Clear indicates an expected call of Clear.
func (mr *MockHazelcastMapInterfaceMockRecorder) Clear(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clear", reflect.TypeOf((*MockHazelcastMapInterface)(nil).Clear), ctx)
}

// Get mocks base method.
func (m *MockHazelcastMapInterface) Get(ctx context.Context, key any) (any, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, key)
	ret0, _ := ret[0].(any)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockHazelcastMapInterfaceMockRecorder) Get(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockHazelcastMapInterface)(nil).Get), ctx, key)
}

// GetEntryView mocks base method.
func (m *MockHazelcastMapInterface) GetEntryView(ctx context.Context, key any) (*types.SimpleEntryView, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEntryView", ctx, key)
	ret0, _ := ret[0].(*types.SimpleEntryView)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEntryView indicates an expected call of GetEntryView.
func (mr *MockHazelcastMapInterfaceMockRecorder) GetEntryView(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEntryView", reflect.TypeOf((*MockHazelcastMapInterface)(nil).GetEntryView), ctx, key)
}

// Remove mocks base method.
func (m *MockHazelcastMapInterface) Remove(ctx context.Context, key any) (any, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", ctx, key)
	ret0, _ := ret[0].(any)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Remove indicates an expected call of Remove.
func (mr *MockHazelcastMapInterfaceMockRecorder) Remove(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockHazelcastMapInterface)(nil).Remove), ctx, key)
}

// SetWithTTL mocks base method.
func (m *MockHazelcastMapInterface) SetWithTTL(ctx context.Context, key, value any, ttl time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetWithTTL", ctx, key, value, ttl)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetWithTTL indicates an expected call of SetWithTTL.
func (mr *MockHazelcastMapInterfaceMockRecorder) SetWithTTL(ctx, key, value, ttl interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetWithTTL", reflect.TypeOf((*MockHazelcastMapInterface)(nil).SetWithTTL), ctx, key, value, ttl)
}