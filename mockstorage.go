// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/lovoo/goka/storage (interfaces: Storage)

// Package goka is a generated GoMock package.
package goka

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	storage "github.com/lovoo/goka/storage"
)

// MockStorage is a mock of Storage interface.
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage.
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance.
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockStorage) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockStorageMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockStorage)(nil).Close))
}

// Delete mocks base method.
func (m *MockStorage) Delete(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockStorageMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockStorage)(nil).Delete), arg0)
}

// Get mocks base method.
func (m *MockStorage) Get(arg0 string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockStorageMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockStorage)(nil).Get), arg0)
}

// GetOffset mocks base method.
func (m *MockStorage) GetOffset(arg0 int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOffset", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOffset indicates an expected call of GetOffset.
func (mr *MockStorageMockRecorder) GetOffset(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOffset", reflect.TypeOf((*MockStorage)(nil).GetOffset), arg0)
}

// Has mocks base method.
func (m *MockStorage) Has(arg0 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Has indicates an expected call of Has.
func (mr *MockStorageMockRecorder) Has(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockStorage)(nil).Has), arg0)
}

// Iterator mocks base method.
func (m *MockStorage) Iterator() (storage.Iterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Iterator")
	ret0, _ := ret[0].(storage.Iterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Iterator indicates an expected call of Iterator.
func (mr *MockStorageMockRecorder) Iterator() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Iterator", reflect.TypeOf((*MockStorage)(nil).Iterator))
}

// IteratorWithRange mocks base method.
func (m *MockStorage) IteratorWithRange(arg0, arg1 []byte) (storage.Iterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IteratorWithRange", arg0, arg1)
	ret0, _ := ret[0].(storage.Iterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IteratorWithRange indicates an expected call of IteratorWithRange.
func (mr *MockStorageMockRecorder) IteratorWithRange(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IteratorWithRange", reflect.TypeOf((*MockStorage)(nil).IteratorWithRange), arg0, arg1)
}

// MarkRecovered mocks base method.
func (m *MockStorage) MarkRecovered() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarkRecovered")
	ret0, _ := ret[0].(error)
	return ret0
}

// MarkRecovered indicates an expected call of MarkRecovered.
func (mr *MockStorageMockRecorder) MarkRecovered() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkRecovered", reflect.TypeOf((*MockStorage)(nil).MarkRecovered))
}

// Open mocks base method.
func (m *MockStorage) Open() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Open")
	ret0, _ := ret[0].(error)
	return ret0
}

// Open indicates an expected call of Open.
func (mr *MockStorageMockRecorder) Open() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Open", reflect.TypeOf((*MockStorage)(nil).Open))
}

// Set mocks base method.
func (m *MockStorage) Set(arg0 string, arg1 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockStorageMockRecorder) Set(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockStorage)(nil).Set), arg0, arg1)
}

// SetOffset mocks base method.
func (m *MockStorage) SetOffset(arg0 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetOffset", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetOffset indicates an expected call of SetOffset.
func (mr *MockStorageMockRecorder) SetOffset(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetOffset", reflect.TypeOf((*MockStorage)(nil).SetOffset), arg0)
}
