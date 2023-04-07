// Code generated by MockGen. DO NOT EDIT.
// Source: crypt.go

// Package crypt is a generated GoMock package.
package crypt

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCrypt is a mock of Crypt interface.
type MockCrypt struct {
	ctrl     *gomock.Controller
	recorder *MockCryptMockRecorder
}

// MockCryptMockRecorder is the mock recorder for MockCrypt.
type MockCryptMockRecorder struct {
	mock *MockCrypt
}

// NewMockCrypt creates a new mock instance.
func NewMockCrypt(ctrl *gomock.Controller) *MockCrypt {
	mock := &MockCrypt{ctrl: ctrl}
	mock.recorder = &MockCryptMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCrypt) EXPECT() *MockCryptMockRecorder {
	return m.recorder
}

// NewHashSha1 mocks base method.
func (m *MockCrypt) NewHashSha1(str string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewHashSha1", str)
	ret0, _ := ret[0].(string)
	return ret0
}

// NewHashSha1 indicates an expected call of NewHashSha1.
func (mr *MockCryptMockRecorder) NewHashSha1(str interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewHashSha1", reflect.TypeOf((*MockCrypt)(nil).NewHashSha1), str)
}

// NewHashSha256 mocks base method.
func (m *MockCrypt) NewHashSha256(str string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewHashSha256", str)
	ret0, _ := ret[0].(string)
	return ret0
}

// NewHashSha256 indicates an expected call of NewHashSha256.
func (mr *MockCryptMockRecorder) NewHashSha256(str interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewHashSha256", reflect.TypeOf((*MockCrypt)(nil).NewHashSha256), str)
}

// PasswordCompare mocks base method.
func (m *MockCrypt) PasswordCompare(encryptedPassword, password []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PasswordCompare", encryptedPassword, password)
	ret0, _ := ret[0].(error)
	return ret0
}

// PasswordCompare indicates an expected call of PasswordCompare.
func (mr *MockCryptMockRecorder) PasswordCompare(encryptedPassword, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PasswordCompare", reflect.TypeOf((*MockCrypt)(nil).PasswordCompare), encryptedPassword, password)
}

// PasswordEncrypt mocks base method.
func (m *MockCrypt) PasswordEncrypt(password string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PasswordEncrypt", password)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PasswordEncrypt indicates an expected call of PasswordEncrypt.
func (mr *MockCryptMockRecorder) PasswordEncrypt(password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PasswordEncrypt", reflect.TypeOf((*MockCrypt)(nil).PasswordEncrypt), password)
}

// RandomBytes mocks base method.
func (m *MockCrypt) RandomBytes(size int) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RandomBytes", size)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RandomBytes indicates an expected call of RandomBytes.
func (mr *MockCryptMockRecorder) RandomBytes(size interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RandomBytes", reflect.TypeOf((*MockCrypt)(nil).RandomBytes), size)
}