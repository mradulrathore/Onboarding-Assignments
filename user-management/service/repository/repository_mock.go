// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package repository is a generated GoMock package.
package repository

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	user "github.com/mradulrathore/user-management/service/user"
)

// MockRepositoryI is a mock of RepositoryI interface.
type MockRepositoryI struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryIMockRecorder
}

// MockRepositoryIMockRecorder is the mock recorder for MockRepositoryI.
type MockRepositoryIMockRecorder struct {
	mock *MockRepositoryI
}

// NewMockRepositoryI creates a new mock instance.
func NewMockRepositoryI(ctrl *gomock.Controller) *MockRepositoryI {
	mock := &MockRepositoryI{ctrl: ctrl}
	mock.recorder = &MockRepositoryIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepositoryI) EXPECT() *MockRepositoryIMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockRepositoryI) Add(arg0 user.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockRepositoryIMockRecorder) Add(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockRepositoryI)(nil).Add), arg0)
}

// CheckDataExistence mocks base method.
func (m *MockRepositoryI) CheckDataExistence(arg0 int) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckDataExistence", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CheckDataExistence indicates an expected call of CheckDataExistence.
func (mr *MockRepositoryIMockRecorder) CheckDataExistence(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckDataExistence", reflect.TypeOf((*MockRepositoryI)(nil).CheckDataExistence), arg0)
}

// Close mocks base method.
func (m *MockRepositoryI) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockRepositoryIMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockRepositoryI)(nil).Close))
}

// DeleteByRollNo mocks base method.
func (m *MockRepositoryI) DeleteByRollNo(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByRollNo", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByRollNo indicates an expected call of DeleteByRollNo.
func (mr *MockRepositoryIMockRecorder) DeleteByRollNo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByRollNo", reflect.TypeOf((*MockRepositoryI)(nil).DeleteByRollNo), arg0)
}

// GetAll mocks base method.
func (m *MockRepositoryI) GetAll(arg0 string, arg1 int) ([]user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", arg0, arg1)
	ret0, _ := ret[0].([]user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockRepositoryIMockRecorder) GetAll(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockRepositoryI)(nil).GetAll), arg0, arg1)
}

// Load mocks base method.
func (m *MockRepositoryI) Load() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Load")
	ret0, _ := ret[0].(error)
	return ret0
}

// Load indicates an expected call of Load.
func (mr *MockRepositoryIMockRecorder) Load() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Load", reflect.TypeOf((*MockRepositoryI)(nil).Load))
}

// Save mocks base method.
func (m *MockRepositoryI) Save(arg0 []user.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockRepositoryIMockRecorder) Save(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockRepositoryI)(nil).Save), arg0)
}
