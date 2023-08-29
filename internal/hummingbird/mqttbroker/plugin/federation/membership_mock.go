// Code generated by MockGen. DO NOT EDIT.
// Source: plugin/federation/membership.go

// Package federation is a generated GoMock package.
package federation

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	serf "github.com/hashicorp/serf/serf"
)

// MockiSerf is a mock of iSerf interface
type MockiSerf struct {
	ctrl     *gomock.Controller
	recorder *MockiSerfMockRecorder
}

// MockiSerfMockRecorder is the mock recorder for MockiSerf
type MockiSerfMockRecorder struct {
	mock *MockiSerf
}

// NewMockiSerf creates a new mock instance
func NewMockiSerf(ctrl *gomock.Controller) *MockiSerf {
	mock := &MockiSerf{ctrl: ctrl}
	mock.recorder = &MockiSerfMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockiSerf) EXPECT() *MockiSerfMockRecorder {
	return m.recorder
}

// Join mocks base method
func (m *MockiSerf) Join(existing []string, ignoreOld bool) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Join", existing, ignoreOld)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Join indicates an expected call of Join
func (mr *MockiSerfMockRecorder) Join(existing, ignoreOld interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Join", reflect.TypeOf((*MockiSerf)(nil).Join), existing, ignoreOld)
}

// RemoveFailedNode mocks base method
func (m *MockiSerf) RemoveFailedNode(node string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveFailedNode", node)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveFailedNode indicates an expected call of RemoveFailedNode
func (mr *MockiSerfMockRecorder) RemoveFailedNode(node interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveFailedNode", reflect.TypeOf((*MockiSerf)(nil).RemoveFailedNode), node)
}

// Leave mocks base method
func (m *MockiSerf) Leave() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Leave")
	ret0, _ := ret[0].(error)
	return ret0
}

// Leave indicates an expected call of Leave
func (mr *MockiSerfMockRecorder) Leave() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Leave", reflect.TypeOf((*MockiSerf)(nil).Leave))
}

// Members mocks base method
func (m *MockiSerf) Members() []serf.Member {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Members")
	ret0, _ := ret[0].([]serf.Member)
	return ret0
}

// Members indicates an expected call of Members
func (mr *MockiSerfMockRecorder) Members() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Members", reflect.TypeOf((*MockiSerf)(nil).Members))
}

// Shutdown mocks base method
func (m *MockiSerf) Shutdown() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Shutdown")
	ret0, _ := ret[0].(error)
	return ret0
}

// Shutdown indicates an expected call of Shutdown
func (mr *MockiSerfMockRecorder) Shutdown() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockiSerf)(nil).Shutdown))
}
