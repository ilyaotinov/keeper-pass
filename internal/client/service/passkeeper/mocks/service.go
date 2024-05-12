// Code generated by MockGen. DO NOT EDIT.
// Source: internal/client/service/passkeeper/service.go

// Package passkeeper_mock is a generated GoMock package.
package passkeeper_mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/ilya372317/pass-keeper/internal/client/domain"
)

// MockpassClient is a mock of passClient interface.
type MockpassClient struct {
	ctrl     *gomock.Controller
	recorder *MockpassClientMockRecorder
}

// MockpassClientMockRecorder is the mock recorder for MockpassClient.
type MockpassClientMockRecorder struct {
	mock *MockpassClient
}

// NewMockpassClient creates a new mock instance.
func NewMockpassClient(ctrl *gomock.Controller) *MockpassClient {
	mock := &MockpassClient{ctrl: ctrl}
	mock.recorder = &MockpassClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockpassClient) EXPECT() *MockpassClientMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockpassClient) All(ctx context.Context) ([]domain.ShortData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx)
	ret0, _ := ret[0].([]domain.ShortData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockpassClientMockRecorder) All(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockpassClient)(nil).All), ctx)
}

// Delete mocks base method.
func (m *MockpassClient) Delete(arg0 context.Context, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockpassClientMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockpassClient)(nil).Delete), arg0, arg1)
}

// Login mocks base method.
func (m *MockpassClient) Login(arg0 context.Context, arg1, arg2 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockpassClientMockRecorder) Login(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockpassClient)(nil).Login), arg0, arg1, arg2)
}

// Register mocks base method.
func (m *MockpassClient) Register(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register.
func (mr *MockpassClientMockRecorder) Register(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockpassClient)(nil).Register), arg0, arg1, arg2)
}

// SaveBinary mocks base method.
func (m *MockpassClient) SaveBinary(arg0 context.Context, arg1 string, arg2 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveBinary", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveBinary indicates an expected call of SaveBinary.
func (mr *MockpassClientMockRecorder) SaveBinary(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveBinary", reflect.TypeOf((*MockpassClient)(nil).SaveBinary), arg0, arg1, arg2)
}

// SaveCard mocks base method.
func (m *MockpassClient) SaveCard(arg0 context.Context, arg1, arg2 string, arg3 int, arg4 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveCard", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveCard indicates an expected call of SaveCard.
func (mr *MockpassClientMockRecorder) SaveCard(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveCard", reflect.TypeOf((*MockpassClient)(nil).SaveCard), arg0, arg1, arg2, arg3, arg4)
}

// SaveLogin mocks base method.
func (m *MockpassClient) SaveLogin(arg0 context.Context, arg1, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveLogin", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveLogin indicates an expected call of SaveLogin.
func (mr *MockpassClientMockRecorder) SaveLogin(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveLogin", reflect.TypeOf((*MockpassClient)(nil).SaveLogin), arg0, arg1, arg2, arg3)
}

// SaveText mocks base method.
func (m *MockpassClient) SaveText(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveText", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveText indicates an expected call of SaveText.
func (mr *MockpassClientMockRecorder) SaveText(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveText", reflect.TypeOf((*MockpassClient)(nil).SaveText), arg0, arg1, arg2)
}

// MocktokenStorage is a mock of tokenStorage interface.
type MocktokenStorage struct {
	ctrl     *gomock.Controller
	recorder *MocktokenStorageMockRecorder
}

// MocktokenStorageMockRecorder is the mock recorder for MocktokenStorage.
type MocktokenStorageMockRecorder struct {
	mock *MocktokenStorage
}

// NewMocktokenStorage creates a new mock instance.
func NewMocktokenStorage(ctrl *gomock.Controller) *MocktokenStorage {
	mock := &MocktokenStorage{ctrl: ctrl}
	mock.recorder = &MocktokenStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MocktokenStorage) EXPECT() *MocktokenStorageMockRecorder {
	return m.recorder
}

// SetAccessToken mocks base method.
func (m *MocktokenStorage) SetAccessToken(token string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetAccessToken", token)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetAccessToken indicates an expected call of SetAccessToken.
func (mr *MocktokenStorageMockRecorder) SetAccessToken(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAccessToken", reflect.TypeOf((*MocktokenStorage)(nil).SetAccessToken), token)
}
