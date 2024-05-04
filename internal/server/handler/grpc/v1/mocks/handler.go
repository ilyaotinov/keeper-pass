// Code generated by MockGen. DO NOT EDIT.
// Source: internal/server/handler/grpc/v1/handler.go

// Package v1_mock is a generated GoMock package.
package v1_mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/ilya372317/pass-keeper/internal/server/domain"
	dto "github.com/ilya372317/pass-keeper/internal/server/dto"
)

// MockAuthService is a mock of AuthService interface.
type MockAuthService struct {
	ctrl     *gomock.Controller
	recorder *MockAuthServiceMockRecorder
}

// MockAuthServiceMockRecorder is the mock recorder for MockAuthService.
type MockAuthServiceMockRecorder struct {
	mock *MockAuthService
}

// NewMockAuthService creates a new mock instance.
func NewMockAuthService(ctrl *gomock.Controller) *MockAuthService {
	mock := &MockAuthService{ctrl: ctrl}
	mock.recorder = &MockAuthServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthService) EXPECT() *MockAuthServiceMockRecorder {
	return m.recorder
}

// Login mocks base method.
func (m *MockAuthService) Login(arg0 context.Context, arg1 dto.LoginDTO) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockAuthServiceMockRecorder) Login(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockAuthService)(nil).Login), arg0, arg1)
}

// Register mocks base method.
func (m *MockAuthService) Register(arg0 context.Context, arg1 dto.RegisterDTO) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", arg0, arg1)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register.
func (mr *MockAuthServiceMockRecorder) Register(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockAuthService)(nil).Register), arg0, arg1)
}

// MockloginPassService is a mock of loginPassService interface.
type MockloginPassService struct {
	ctrl     *gomock.Controller
	recorder *MockloginPassServiceMockRecorder
}

// MockloginPassServiceMockRecorder is the mock recorder for MockloginPassService.
type MockloginPassServiceMockRecorder struct {
	mock *MockloginPassService
}

// NewMockloginPassService creates a new mock instance.
func NewMockloginPassService(ctrl *gomock.Controller) *MockloginPassService {
	mock := &MockloginPassService{ctrl: ctrl}
	mock.recorder = &MockloginPassServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockloginPassService) EXPECT() *MockloginPassServiceMockRecorder {
	return m.recorder
}

// Save mocks base method.
func (m *MockloginPassService) Save(ctx context.Context, d dto.SaveLoginPassDTO) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", ctx, d)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockloginPassServiceMockRecorder) Save(ctx, d interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockloginPassService)(nil).Save), ctx, d)
}

// Show mocks base method.
func (m *MockloginPassService) Show(ctx context.Context, id int) (domain.LoginPassData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Show", ctx, id)
	ret0, _ := ret[0].(domain.LoginPassData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Show indicates an expected call of Show.
func (mr *MockloginPassServiceMockRecorder) Show(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Show", reflect.TypeOf((*MockloginPassService)(nil).Show), ctx, id)
}

// Update mocks base method.
func (m *MockloginPassService) Update(ctx context.Context, d dto.UpdateLoginPassDTO) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, d)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockloginPassServiceMockRecorder) Update(ctx, d interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockloginPassService)(nil).Update), ctx, d)
}

// MockcreditCardService is a mock of creditCardService interface.
type MockcreditCardService struct {
	ctrl     *gomock.Controller
	recorder *MockcreditCardServiceMockRecorder
}

// MockcreditCardServiceMockRecorder is the mock recorder for MockcreditCardService.
type MockcreditCardServiceMockRecorder struct {
	mock *MockcreditCardService
}

// NewMockcreditCardService creates a new mock instance.
func NewMockcreditCardService(ctrl *gomock.Controller) *MockcreditCardService {
	mock := &MockcreditCardService{ctrl: ctrl}
	mock.recorder = &MockcreditCardServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockcreditCardService) EXPECT() *MockcreditCardServiceMockRecorder {
	return m.recorder
}

// Save mocks base method.
func (m *MockcreditCardService) Save(ctx context.Context, d dto.SaveCreditCardDTO) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", ctx, d)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockcreditCardServiceMockRecorder) Save(ctx, d interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockcreditCardService)(nil).Save), ctx, d)
}

// Show mocks base method.
func (m *MockcreditCardService) Show(ctx context.Context, id int64) (domain.CreditCardData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Show", ctx, id)
	ret0, _ := ret[0].(domain.CreditCardData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Show indicates an expected call of Show.
func (mr *MockcreditCardServiceMockRecorder) Show(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Show", reflect.TypeOf((*MockcreditCardService)(nil).Show), ctx, id)
}

// Update mocks base method.
func (m *MockcreditCardService) Update(ctx context.Context, d dto.UpdateCreditCardDTO) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, d)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockcreditCardServiceMockRecorder) Update(ctx, d interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockcreditCardService)(nil).Update), ctx, d)
}
