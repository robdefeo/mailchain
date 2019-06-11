// Code generated by MockGen. DO NOT EDIT.
// Source: pubkey_finder.go

// Package configtest is a generated GoMock package.
package configtest

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockPubKeyFinderSetter is a mock of PubKeyFinderSetter interface
type MockPubKeyFinderSetter struct {
	ctrl     *gomock.Controller
	recorder *MockPubKeyFinderSetterMockRecorder
}

// MockPubKeyFinderSetterMockRecorder is the mock recorder for MockPubKeyFinderSetter
type MockPubKeyFinderSetterMockRecorder struct {
	mock *MockPubKeyFinderSetter
}

// NewMockPubKeyFinderSetter creates a new mock instance
func NewMockPubKeyFinderSetter(ctrl *gomock.Controller) *MockPubKeyFinderSetter {
	mock := &MockPubKeyFinderSetter{ctrl: ctrl}
	mock.recorder = &MockPubKeyFinderSetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPubKeyFinderSetter) EXPECT() *MockPubKeyFinderSetterMockRecorder {
	return m.recorder
}

// Set mocks base method
func (m *MockPubKeyFinderSetter) Set(chain, network, pubkeyFinder string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", chain, network, pubkeyFinder)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set
func (mr *MockPubKeyFinderSetterMockRecorder) Set(chain, network, pubkeyFinder interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockPubKeyFinderSetter)(nil).Set), chain, network, pubkeyFinder)
}
