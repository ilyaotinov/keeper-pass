// Code generated by MockGen. DO NOT EDIT.
// Source: internal/server/service/loginpass/service.go

// Package loginpass_mock is a generated GoMock package.
package loginpass_mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/ilya372317/pass-keeper/internal/server/domain"
)

// MockdataService is a mock of dataService interface.
type MockdataService struct {
	ctrl     *gomock.Controller
	recorder *MockdataServiceMockRecorder
}

// MockdataServiceMockRecorder is the mock recorder for MockdataService.
type MockdataServiceMockRecorder struct {
	mock *MockdataService
}

// NewMockdataService creates a new mock instance.
func NewMockdataService(ctrl *gomock.Controller) *MockdataService {
	mock := &MockdataService{ctrl: ctrl}
	mock.recorder = &MockdataServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockdataService) EXPECT() *MockdataServiceMockRecorder {
	return m.recorder
}

// EncryptAndSaveData mocks base method.
func (m *MockdataService) EncryptAndSaveData(arg0 context.Context, arg1 domain.Data) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EncryptAndSaveData", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// EncryptAndSaveData indicates an expected call of EncryptAndSaveData.
func (mr *MockdataServiceMockRecorder) EncryptAndSaveData(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EncryptAndSaveData", reflect.TypeOf((*MockdataService)(nil).EncryptAndSaveData), arg0, arg1)
}

// EncryptAndUpdateData mocks base method.
func (m *MockdataService) EncryptAndUpdateData(arg0 context.Context, arg1 domain.Data) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EncryptAndUpdateData", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// EncryptAndUpdateData indicates an expected call of EncryptAndUpdateData.
func (mr *MockdataServiceMockRecorder) EncryptAndUpdateData(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EncryptAndUpdateData", reflect.TypeOf((*MockdataService)(nil).EncryptAndUpdateData), arg0, arg1)
}

// GetAndDecryptData mocks base method.
func (m *MockdataService) GetAndDecryptData(arg0 context.Context, arg1 int64) (domain.Data, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAndDecryptData", arg0, arg1)
	ret0, _ := ret[0].(domain.Data)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAndDecryptData indicates an expected call of GetAndDecryptData.
func (mr *MockdataServiceMockRecorder) GetAndDecryptData(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAndDecryptData", reflect.TypeOf((*MockdataService)(nil).GetAndDecryptData), arg0, arg1)
}
