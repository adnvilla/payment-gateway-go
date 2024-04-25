// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/stripe/stripe-go/v78 (interfaces: Backend)
//
// Generated by this command:
//
//	mockgen -destination=src/pkg/stripe/mock/backend.go -package=mock github.com/stripe/stripe-go/v78 Backend
//

// Package mock is a generated GoMock package.
package mock

import (
	bytes "bytes"
	reflect "reflect"

	stripe "github.com/stripe/stripe-go/v78"
	form "github.com/stripe/stripe-go/v78/form"
	gomock "go.uber.org/mock/gomock"
)

// MockBackend is a mock of Backend interface.
type MockBackend struct {
	ctrl     *gomock.Controller
	recorder *MockBackendMockRecorder
}

// MockBackendMockRecorder is the mock recorder for MockBackend.
type MockBackendMockRecorder struct {
	mock *MockBackend
}

// NewMockBackend creates a new mock instance.
func NewMockBackend(ctrl *gomock.Controller) *MockBackend {
	mock := &MockBackend{ctrl: ctrl}
	mock.recorder = &MockBackendMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBackend) EXPECT() *MockBackendMockRecorder {
	return m.recorder
}

// Call mocks base method.
func (m *MockBackend) Call(arg0, arg1, arg2 string, arg3 stripe.ParamsContainer, arg4 stripe.LastResponseSetter) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Call", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// Call indicates an expected call of Call.
func (mr *MockBackendMockRecorder) Call(arg0, arg1, arg2, arg3, arg4 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Call", reflect.TypeOf((*MockBackend)(nil).Call), arg0, arg1, arg2, arg3, arg4)
}

// CallMultipart mocks base method.
func (m *MockBackend) CallMultipart(arg0, arg1, arg2, arg3 string, arg4 *bytes.Buffer, arg5 *stripe.Params, arg6 stripe.LastResponseSetter) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CallMultipart", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(error)
	return ret0
}

// CallMultipart indicates an expected call of CallMultipart.
func (mr *MockBackendMockRecorder) CallMultipart(arg0, arg1, arg2, arg3, arg4, arg5, arg6 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CallMultipart", reflect.TypeOf((*MockBackend)(nil).CallMultipart), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// CallRaw mocks base method.
func (m *MockBackend) CallRaw(arg0, arg1, arg2 string, arg3 *form.Values, arg4 *stripe.Params, arg5 stripe.LastResponseSetter) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CallRaw", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(error)
	return ret0
}

// CallRaw indicates an expected call of CallRaw.
func (mr *MockBackendMockRecorder) CallRaw(arg0, arg1, arg2, arg3, arg4, arg5 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CallRaw", reflect.TypeOf((*MockBackend)(nil).CallRaw), arg0, arg1, arg2, arg3, arg4, arg5)
}

// CallStreaming mocks base method.
func (m *MockBackend) CallStreaming(arg0, arg1, arg2 string, arg3 stripe.ParamsContainer, arg4 stripe.StreamingLastResponseSetter) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CallStreaming", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// CallStreaming indicates an expected call of CallStreaming.
func (mr *MockBackendMockRecorder) CallStreaming(arg0, arg1, arg2, arg3, arg4 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CallStreaming", reflect.TypeOf((*MockBackend)(nil).CallStreaming), arg0, arg1, arg2, arg3, arg4)
}

// SetMaxNetworkRetries mocks base method.
func (m *MockBackend) SetMaxNetworkRetries(arg0 int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetMaxNetworkRetries", arg0)
}

// SetMaxNetworkRetries indicates an expected call of SetMaxNetworkRetries.
func (mr *MockBackendMockRecorder) SetMaxNetworkRetries(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMaxNetworkRetries", reflect.TypeOf((*MockBackend)(nil).SetMaxNetworkRetries), arg0)
}