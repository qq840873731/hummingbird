// Code generated by MockGen. DO NOT EDIT.
// Source: plugin/federation/federation.pb.go

// Package federation is a generated GoMock package.
package federation

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockisEvent_Event is a mock of isEvent_Event interface
type MockisEvent_Event struct {
	ctrl     *gomock.Controller
	recorder *MockisEvent_EventMockRecorder
}

// MockisEvent_EventMockRecorder is the mock recorder for MockisEvent_Event
type MockisEvent_EventMockRecorder struct {
	mock *MockisEvent_Event
}

// NewMockisEvent_Event creates a new mock instance
func NewMockisEvent_Event(ctrl *gomock.Controller) *MockisEvent_Event {
	mock := &MockisEvent_Event{ctrl: ctrl}
	mock.recorder = &MockisEvent_EventMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockisEvent_Event) EXPECT() *MockisEvent_EventMockRecorder {
	return m.recorder
}

// isEvent_Event mocks base method
func (m *MockisEvent_Event) isEvent_Event() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "isEvent_Event")
}

// isEvent_Event indicates an expected call of isEvent_Event
func (mr *MockisEvent_EventMockRecorder) isEvent_Event() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "isEvent_Event", reflect.TypeOf((*MockisEvent_Event)(nil).isEvent_Event))
}
