// Code generated by MockGen. DO NOT EDIT.
// Source: service/interfaces.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	service "github.com/LetsFocus/goLF/service"
	gomock "github.com/golang/mock/gomock"
)

// Mockservice is a mock of service interface.
type Mockservice struct {
	ctrl     *gomock.Controller
	recorder *MockserviceMockRecorder
}

// MockserviceMockRecorder is the mock recorder for Mockservice.
type MockserviceMockRecorder struct {
	mock *Mockservice
}

// NewMockservice creates a new mock instance.
func NewMockservice(ctrl *gomock.Controller) *Mockservice {
	mock := &Mockservice{ctrl: ctrl}
	mock.recorder = &MockserviceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockservice) EXPECT() *MockserviceMockRecorder {
	return m.recorder
}

// Bind mocks base method.
func (m *Mockservice) Bind(data []byte, i interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Bind", data, i)
	ret0, _ := ret[0].(error)
	return ret0
}

// Bind indicates an expected call of Bind.
func (mr *MockserviceMockRecorder) Bind(data, i interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bind", reflect.TypeOf((*Mockservice)(nil).Bind), data, i)
}

// Delete mocks base method.
func (m *Mockservice) Delete(ctx context.Context, target string, params map[string]interface{}, headers map[string]string) (service.HTTPResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, target, params, headers)
	ret0, _ := ret[0].(service.HTTPResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockserviceMockRecorder) Delete(ctx, target, params, headers interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*Mockservice)(nil).Delete), ctx, target, params, headers)
}

// Get mocks base method.
func (m *Mockservice) Get(ctx context.Context, target string, params map[string]interface{}, headers map[string]string) (service.HTTPResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, target, params, headers)
	ret0, _ := ret[0].(service.HTTPResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockserviceMockRecorder) Get(ctx, target, params, headers interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*Mockservice)(nil).Get), ctx, target, params, headers)
}

// Patch mocks base method.
func (m *Mockservice) Patch(ctx context.Context, target string, body []byte, params map[string]interface{}, headers map[string]string) (service.HTTPResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Patch", ctx, target, body, params, headers)
	ret0, _ := ret[0].(service.HTTPResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Patch indicates an expected call of Patch.
func (mr *MockserviceMockRecorder) Patch(ctx, target, body, params, headers interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*Mockservice)(nil).Patch), ctx, target, body, params, headers)
}

// Post mocks base method.
func (m *Mockservice) Post(ctx context.Context, target string, body []byte, headers map[string]string) (service.HTTPResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Post", ctx, target, body, headers)
	ret0, _ := ret[0].(service.HTTPResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Post indicates an expected call of Post.
func (mr *MockserviceMockRecorder) Post(ctx, target, body, headers interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Post", reflect.TypeOf((*Mockservice)(nil).Post), ctx, target, body, headers)
}

// Put mocks base method.
func (m *Mockservice) Put(ctx context.Context, target string, body []byte, params map[string]interface{}, headers map[string]string) (service.HTTPResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Put", ctx, target, body, params, headers)
	ret0, _ := ret[0].(service.HTTPResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Put indicates an expected call of Put.
func (mr *MockserviceMockRecorder) Put(ctx, target, body, params, headers interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*Mockservice)(nil).Put), ctx, target, body, params, headers)
}
