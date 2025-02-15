// Code generated by MockGen. DO NOT EDIT.
// Source: backend.go
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/backend.go -source backend.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	http "net/http"
	reflect "reflect"

	authproviders "github.com/stackrox/rox/pkg/auth/authproviders"
	tokens "github.com/stackrox/rox/pkg/auth/tokens"
	requestinfo "github.com/stackrox/rox/pkg/grpc/requestinfo"
	gomock "go.uber.org/mock/gomock"
)

// MockBackend is a mock of Backend interface.
type MockBackend struct {
	ctrl     *gomock.Controller
	recorder *MockBackendMockRecorder
	isgomock struct{}
}

// MockBackendMockRecorder is the mock recorder for MockBackend.
type MockBackendMockRecorder struct {
	mock *MockBackend
}

// NewMockBackend creates a new mock instance.
func NewMockBackend(ctrl *gomock.Controller) *MockBackend {
	mock := &MockBackend{ctrl: ctrl}
	mock.recorder = &MockBackendMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBackend) EXPECT() *MockBackendMockRecorder {
	return m.recorder
}

// Config mocks base method.
func (m *MockBackend) Config() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Config")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// Config indicates an expected call of Config.
func (mr *MockBackendMockRecorder) Config() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Config", reflect.TypeOf((*MockBackend)(nil).Config))
}

// ExchangeToken mocks base method.
func (m *MockBackend) ExchangeToken(ctx context.Context, externalToken, state string) (*authproviders.AuthResponse, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExchangeToken", ctx, externalToken, state)
	ret0, _ := ret[0].(*authproviders.AuthResponse)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ExchangeToken indicates an expected call of ExchangeToken.
func (mr *MockBackendMockRecorder) ExchangeToken(ctx, externalToken, state any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExchangeToken", reflect.TypeOf((*MockBackend)(nil).ExchangeToken), ctx, externalToken, state)
}

// LoginURL mocks base method.
func (m *MockBackend) LoginURL(clientState string, ri *requestinfo.RequestInfo) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoginURL", clientState, ri)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoginURL indicates an expected call of LoginURL.
func (mr *MockBackendMockRecorder) LoginURL(clientState, ri any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginURL", reflect.TypeOf((*MockBackend)(nil).LoginURL), clientState, ri)
}

// OnDisable mocks base method.
func (m *MockBackend) OnDisable(provider authproviders.Provider) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnDisable", provider)
}

// OnDisable indicates an expected call of OnDisable.
func (mr *MockBackendMockRecorder) OnDisable(provider any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnDisable", reflect.TypeOf((*MockBackend)(nil).OnDisable), provider)
}

// OnEnable mocks base method.
func (m *MockBackend) OnEnable(provider authproviders.Provider) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnEnable", provider)
}

// OnEnable indicates an expected call of OnEnable.
func (mr *MockBackendMockRecorder) OnEnable(provider any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnEnable", reflect.TypeOf((*MockBackend)(nil).OnEnable), provider)
}

// ProcessHTTPRequest mocks base method.
func (m *MockBackend) ProcessHTTPRequest(w http.ResponseWriter, r *http.Request) (*authproviders.AuthResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessHTTPRequest", w, r)
	ret0, _ := ret[0].(*authproviders.AuthResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProcessHTTPRequest indicates an expected call of ProcessHTTPRequest.
func (mr *MockBackendMockRecorder) ProcessHTTPRequest(w, r any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessHTTPRequest", reflect.TypeOf((*MockBackend)(nil).ProcessHTTPRequest), w, r)
}

// RefreshURL mocks base method.
func (m *MockBackend) RefreshURL() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshURL")
	ret0, _ := ret[0].(string)
	return ret0
}

// RefreshURL indicates an expected call of RefreshURL.
func (mr *MockBackendMockRecorder) RefreshURL() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshURL", reflect.TypeOf((*MockBackend)(nil).RefreshURL))
}

// Validate mocks base method.
func (m *MockBackend) Validate(ctx context.Context, claims *tokens.Claims) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate", ctx, claims)
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate.
func (mr *MockBackendMockRecorder) Validate(ctx, claims any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockBackend)(nil).Validate), ctx, claims)
}

// MockRefreshTokenEnabledBackend is a mock of RefreshTokenEnabledBackend interface.
type MockRefreshTokenEnabledBackend struct {
	ctrl     *gomock.Controller
	recorder *MockRefreshTokenEnabledBackendMockRecorder
	isgomock struct{}
}

// MockRefreshTokenEnabledBackendMockRecorder is the mock recorder for MockRefreshTokenEnabledBackend.
type MockRefreshTokenEnabledBackendMockRecorder struct {
	mock *MockRefreshTokenEnabledBackend
}

// NewMockRefreshTokenEnabledBackend creates a new mock instance.
func NewMockRefreshTokenEnabledBackend(ctrl *gomock.Controller) *MockRefreshTokenEnabledBackend {
	mock := &MockRefreshTokenEnabledBackend{ctrl: ctrl}
	mock.recorder = &MockRefreshTokenEnabledBackendMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRefreshTokenEnabledBackend) EXPECT() *MockRefreshTokenEnabledBackendMockRecorder {
	return m.recorder
}

// Config mocks base method.
func (m *MockRefreshTokenEnabledBackend) Config() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Config")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// Config indicates an expected call of Config.
func (mr *MockRefreshTokenEnabledBackendMockRecorder) Config() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Config", reflect.TypeOf((*MockRefreshTokenEnabledBackend)(nil).Config))
}

// ExchangeToken mocks base method.
func (m *MockRefreshTokenEnabledBackend) ExchangeToken(ctx context.Context, externalToken, state string) (*authproviders.AuthResponse, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExchangeToken", ctx, externalToken, state)
	ret0, _ := ret[0].(*authproviders.AuthResponse)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ExchangeToken indicates an expected call of ExchangeToken.
func (mr *MockRefreshTokenEnabledBackendMockRecorder) ExchangeToken(ctx, externalToken, state any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExchangeToken", reflect.TypeOf((*MockRefreshTokenEnabledBackend)(nil).ExchangeToken), ctx, externalToken, state)
}

// LoginURL mocks base method.
func (m *MockRefreshTokenEnabledBackend) LoginURL(clientState string, ri *requestinfo.RequestInfo) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoginURL", clientState, ri)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoginURL indicates an expected call of LoginURL.
func (mr *MockRefreshTokenEnabledBackendMockRecorder) LoginURL(clientState, ri any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginURL", reflect.TypeOf((*MockRefreshTokenEnabledBackend)(nil).LoginURL), clientState, ri)
}

// OnDisable mocks base method.
func (m *MockRefreshTokenEnabledBackend) OnDisable(provider authproviders.Provider) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnDisable", provider)
}

// OnDisable indicates an expected call of OnDisable.
func (mr *MockRefreshTokenEnabledBackendMockRecorder) OnDisable(provider any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnDisable", reflect.TypeOf((*MockRefreshTokenEnabledBackend)(nil).OnDisable), provider)
}

// OnEnable mocks base method.
func (m *MockRefreshTokenEnabledBackend) OnEnable(provider authproviders.Provider) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnEnable", provider)
}

// OnEnable indicates an expected call of OnEnable.
func (mr *MockRefreshTokenEnabledBackendMockRecorder) OnEnable(provider any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnEnable", reflect.TypeOf((*MockRefreshTokenEnabledBackend)(nil).OnEnable), provider)
}

// ProcessHTTPRequest mocks base method.
func (m *MockRefreshTokenEnabledBackend) ProcessHTTPRequest(w http.ResponseWriter, r *http.Request) (*authproviders.AuthResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessHTTPRequest", w, r)
	ret0, _ := ret[0].(*authproviders.AuthResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProcessHTTPRequest indicates an expected call of ProcessHTTPRequest.
func (mr *MockRefreshTokenEnabledBackendMockRecorder) ProcessHTTPRequest(w, r any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessHTTPRequest", reflect.TypeOf((*MockRefreshTokenEnabledBackend)(nil).ProcessHTTPRequest), w, r)
}

// RefreshAccessToken mocks base method.
func (m *MockRefreshTokenEnabledBackend) RefreshAccessToken(ctx context.Context, refreshTokenData authproviders.RefreshTokenData) (*authproviders.AuthResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshAccessToken", ctx, refreshTokenData)
	ret0, _ := ret[0].(*authproviders.AuthResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RefreshAccessToken indicates an expected call of RefreshAccessToken.
func (mr *MockRefreshTokenEnabledBackendMockRecorder) RefreshAccessToken(ctx, refreshTokenData any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshAccessToken", reflect.TypeOf((*MockRefreshTokenEnabledBackend)(nil).RefreshAccessToken), ctx, refreshTokenData)
}

// RefreshURL mocks base method.
func (m *MockRefreshTokenEnabledBackend) RefreshURL() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshURL")
	ret0, _ := ret[0].(string)
	return ret0
}

// RefreshURL indicates an expected call of RefreshURL.
func (mr *MockRefreshTokenEnabledBackendMockRecorder) RefreshURL() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshURL", reflect.TypeOf((*MockRefreshTokenEnabledBackend)(nil).RefreshURL))
}

// RevokeRefreshToken mocks base method.
func (m *MockRefreshTokenEnabledBackend) RevokeRefreshToken(ctx context.Context, refreshTokenData authproviders.RefreshTokenData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RevokeRefreshToken", ctx, refreshTokenData)
	ret0, _ := ret[0].(error)
	return ret0
}

// RevokeRefreshToken indicates an expected call of RevokeRefreshToken.
func (mr *MockRefreshTokenEnabledBackendMockRecorder) RevokeRefreshToken(ctx, refreshTokenData any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokeRefreshToken", reflect.TypeOf((*MockRefreshTokenEnabledBackend)(nil).RevokeRefreshToken), ctx, refreshTokenData)
}

// Validate mocks base method.
func (m *MockRefreshTokenEnabledBackend) Validate(ctx context.Context, claims *tokens.Claims) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate", ctx, claims)
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate.
func (mr *MockRefreshTokenEnabledBackendMockRecorder) Validate(ctx, claims any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockRefreshTokenEnabledBackend)(nil).Validate), ctx, claims)
}
