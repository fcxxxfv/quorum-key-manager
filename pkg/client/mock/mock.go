// Code generated by MockGen. DO NOT EDIT.
// Source: client.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	jsonrpc "github.com/ConsenSysQuorum/quorum-key-manager/pkg/jsonrpc"
	types "github.com/ConsenSysQuorum/quorum-key-manager/src/api/types"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockSecretsClient is a mock of SecretsClient interface
type MockSecretsClient struct {
	ctrl     *gomock.Controller
	recorder *MockSecretsClientMockRecorder
}

// MockSecretsClientMockRecorder is the mock recorder for MockSecretsClient
type MockSecretsClientMockRecorder struct {
	mock *MockSecretsClient
}

// NewMockSecretsClient creates a new mock instance
func NewMockSecretsClient(ctrl *gomock.Controller) *MockSecretsClient {
	mock := &MockSecretsClient{ctrl: ctrl}
	mock.recorder = &MockSecretsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSecretsClient) EXPECT() *MockSecretsClientMockRecorder {
	return m.recorder
}

// SetSecret mocks base method
func (m *MockSecretsClient) SetSecret(ctx context.Context, storeName string, request *types.SetSecretRequest) (*types.SecretResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetSecret", ctx, storeName, request)
	ret0, _ := ret[0].(*types.SecretResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetSecret indicates an expected call of SetSecret
func (mr *MockSecretsClientMockRecorder) SetSecret(ctx, storeName, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSecret", reflect.TypeOf((*MockSecretsClient)(nil).SetSecret), ctx, storeName, request)
}

// GetSecret mocks base method
func (m *MockSecretsClient) GetSecret(ctx context.Context, storeName, id, version string) (*types.SecretResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecret", ctx, storeName, id, version)
	ret0, _ := ret[0].(*types.SecretResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecret indicates an expected call of GetSecret
func (mr *MockSecretsClientMockRecorder) GetSecret(ctx, storeName, id, version interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecret", reflect.TypeOf((*MockSecretsClient)(nil).GetSecret), ctx, storeName, id, version)
}

// ListSecrets mocks base method
func (m *MockSecretsClient) ListSecrets(ctx context.Context, storeName string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSecrets", ctx, storeName)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSecrets indicates an expected call of ListSecrets
func (mr *MockSecretsClientMockRecorder) ListSecrets(ctx, storeName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSecrets", reflect.TypeOf((*MockSecretsClient)(nil).ListSecrets), ctx, storeName)
}

// MockKeysClient is a mock of KeysClient interface
type MockKeysClient struct {
	ctrl     *gomock.Controller
	recorder *MockKeysClientMockRecorder
}

// MockKeysClientMockRecorder is the mock recorder for MockKeysClient
type MockKeysClientMockRecorder struct {
	mock *MockKeysClient
}

// NewMockKeysClient creates a new mock instance
func NewMockKeysClient(ctrl *gomock.Controller) *MockKeysClient {
	mock := &MockKeysClient{ctrl: ctrl}
	mock.recorder = &MockKeysClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockKeysClient) EXPECT() *MockKeysClientMockRecorder {
	return m.recorder
}

// CreateKey mocks base method
func (m *MockKeysClient) CreateKey(ctx context.Context, storeName string, request *types.CreateKeyRequest) (*types.KeyResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateKey", ctx, storeName, request)
	ret0, _ := ret[0].(*types.KeyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateKey indicates an expected call of CreateKey
func (mr *MockKeysClientMockRecorder) CreateKey(ctx, storeName, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateKey", reflect.TypeOf((*MockKeysClient)(nil).CreateKey), ctx, storeName, request)
}

// ImportKey mocks base method
func (m *MockKeysClient) ImportKey(ctx context.Context, storeName string, request *types.ImportKeyRequest) (*types.KeyResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImportKey", ctx, storeName, request)
	ret0, _ := ret[0].(*types.KeyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ImportKey indicates an expected call of ImportKey
func (mr *MockKeysClientMockRecorder) ImportKey(ctx, storeName, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImportKey", reflect.TypeOf((*MockKeysClient)(nil).ImportKey), ctx, storeName, request)
}

// Sign mocks base method
func (m *MockKeysClient) Sign(ctx context.Context, storeName, id string, request *types.SignBase64PayloadRequest) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sign", ctx, storeName, id, request)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Sign indicates an expected call of Sign
func (mr *MockKeysClientMockRecorder) Sign(ctx, storeName, id, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sign", reflect.TypeOf((*MockKeysClient)(nil).Sign), ctx, storeName, id, request)
}

// GetKey mocks base method
func (m *MockKeysClient) GetKey(ctx context.Context, storeName, id string) (*types.KeyResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetKey", ctx, storeName, id)
	ret0, _ := ret[0].(*types.KeyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetKey indicates an expected call of GetKey
func (mr *MockKeysClientMockRecorder) GetKey(ctx, storeName, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKey", reflect.TypeOf((*MockKeysClient)(nil).GetKey), ctx, storeName, id)
}

// ListKeys mocks base method
func (m *MockKeysClient) ListKeys(ctx context.Context, storeName string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListKeys", ctx, storeName)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListKeys indicates an expected call of ListKeys
func (mr *MockKeysClientMockRecorder) ListKeys(ctx, storeName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListKeys", reflect.TypeOf((*MockKeysClient)(nil).ListKeys), ctx, storeName)
}

// DestroyKey mocks base method
func (m *MockKeysClient) DestroyKey(ctx context.Context, storeName, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DestroyKey", ctx, storeName, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DestroyKey indicates an expected call of DestroyKey
func (mr *MockKeysClientMockRecorder) DestroyKey(ctx, storeName, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DestroyKey", reflect.TypeOf((*MockKeysClient)(nil).DestroyKey), ctx, storeName, id)
}

// MockEth1Client is a mock of Eth1Client interface
type MockEth1Client struct {
	ctrl     *gomock.Controller
	recorder *MockEth1ClientMockRecorder
}

// MockEth1ClientMockRecorder is the mock recorder for MockEth1Client
type MockEth1ClientMockRecorder struct {
	mock *MockEth1Client
}

// NewMockEth1Client creates a new mock instance
func NewMockEth1Client(ctrl *gomock.Controller) *MockEth1Client {
	mock := &MockEth1Client{ctrl: ctrl}
	mock.recorder = &MockEth1ClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEth1Client) EXPECT() *MockEth1ClientMockRecorder {
	return m.recorder
}

// CreateAccount mocks base method
func (m *MockEth1Client) CreateAccount(ctx context.Context, storeName string, request *types.CreateEth1AccountRequest) (*types.Eth1AccountResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccount", ctx, storeName, request)
	ret0, _ := ret[0].(*types.Eth1AccountResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccount indicates an expected call of CreateAccount
func (mr *MockEth1ClientMockRecorder) CreateAccount(ctx, storeName, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockEth1Client)(nil).CreateAccount), ctx, storeName, request)
}

// MockJSONRPC is a mock of JSONRPC interface
type MockJSONRPC struct {
	ctrl     *gomock.Controller
	recorder *MockJSONRPCMockRecorder
}

// MockJSONRPCMockRecorder is the mock recorder for MockJSONRPC
type MockJSONRPCMockRecorder struct {
	mock *MockJSONRPC
}

// NewMockJSONRPC creates a new mock instance
func NewMockJSONRPC(ctrl *gomock.Controller) *MockJSONRPC {
	mock := &MockJSONRPC{ctrl: ctrl}
	mock.recorder = &MockJSONRPCMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockJSONRPC) EXPECT() *MockJSONRPCMockRecorder {
	return m.recorder
}

// Call mocks base method
func (m *MockJSONRPC) Call(ctx context.Context, nodeID, method string, args ...interface{}) (*jsonrpc.ResponseMsg, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, nodeID, method}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Call", varargs...)
	ret0, _ := ret[0].(*jsonrpc.ResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Call indicates an expected call of Call
func (mr *MockJSONRPCMockRecorder) Call(ctx, nodeID, method interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, nodeID, method}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Call", reflect.TypeOf((*MockJSONRPC)(nil).Call), varargs...)
}

// MockKeyManagerClient is a mock of KeyManagerClient interface
type MockKeyManagerClient struct {
	ctrl     *gomock.Controller
	recorder *MockKeyManagerClientMockRecorder
}

// MockKeyManagerClientMockRecorder is the mock recorder for MockKeyManagerClient
type MockKeyManagerClientMockRecorder struct {
	mock *MockKeyManagerClient
}

// NewMockKeyManagerClient creates a new mock instance
func NewMockKeyManagerClient(ctrl *gomock.Controller) *MockKeyManagerClient {
	mock := &MockKeyManagerClient{ctrl: ctrl}
	mock.recorder = &MockKeyManagerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockKeyManagerClient) EXPECT() *MockKeyManagerClientMockRecorder {
	return m.recorder
}

// SetSecret mocks base method
func (m *MockKeyManagerClient) SetSecret(ctx context.Context, storeName string, request *types.SetSecretRequest) (*types.SecretResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetSecret", ctx, storeName, request)
	ret0, _ := ret[0].(*types.SecretResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetSecret indicates an expected call of SetSecret
func (mr *MockKeyManagerClientMockRecorder) SetSecret(ctx, storeName, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSecret", reflect.TypeOf((*MockKeyManagerClient)(nil).SetSecret), ctx, storeName, request)
}

// GetSecret mocks base method
func (m *MockKeyManagerClient) GetSecret(ctx context.Context, storeName, id, version string) (*types.SecretResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecret", ctx, storeName, id, version)
	ret0, _ := ret[0].(*types.SecretResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecret indicates an expected call of GetSecret
func (mr *MockKeyManagerClientMockRecorder) GetSecret(ctx, storeName, id, version interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecret", reflect.TypeOf((*MockKeyManagerClient)(nil).GetSecret), ctx, storeName, id, version)
}

// ListSecrets mocks base method
func (m *MockKeyManagerClient) ListSecrets(ctx context.Context, storeName string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSecrets", ctx, storeName)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSecrets indicates an expected call of ListSecrets
func (mr *MockKeyManagerClientMockRecorder) ListSecrets(ctx, storeName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSecrets", reflect.TypeOf((*MockKeyManagerClient)(nil).ListSecrets), ctx, storeName)
}

// CreateKey mocks base method
func (m *MockKeyManagerClient) CreateKey(ctx context.Context, storeName string, request *types.CreateKeyRequest) (*types.KeyResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateKey", ctx, storeName, request)
	ret0, _ := ret[0].(*types.KeyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateKey indicates an expected call of CreateKey
func (mr *MockKeyManagerClientMockRecorder) CreateKey(ctx, storeName, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateKey", reflect.TypeOf((*MockKeyManagerClient)(nil).CreateKey), ctx, storeName, request)
}

// ImportKey mocks base method
func (m *MockKeyManagerClient) ImportKey(ctx context.Context, storeName string, request *types.ImportKeyRequest) (*types.KeyResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImportKey", ctx, storeName, request)
	ret0, _ := ret[0].(*types.KeyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ImportKey indicates an expected call of ImportKey
func (mr *MockKeyManagerClientMockRecorder) ImportKey(ctx, storeName, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImportKey", reflect.TypeOf((*MockKeyManagerClient)(nil).ImportKey), ctx, storeName, request)
}

// Sign mocks base method
func (m *MockKeyManagerClient) Sign(ctx context.Context, storeName, id string, request *types.SignBase64PayloadRequest) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sign", ctx, storeName, id, request)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Sign indicates an expected call of Sign
func (mr *MockKeyManagerClientMockRecorder) Sign(ctx, storeName, id, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sign", reflect.TypeOf((*MockKeyManagerClient)(nil).Sign), ctx, storeName, id, request)
}

// GetKey mocks base method
func (m *MockKeyManagerClient) GetKey(ctx context.Context, storeName, id string) (*types.KeyResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetKey", ctx, storeName, id)
	ret0, _ := ret[0].(*types.KeyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetKey indicates an expected call of GetKey
func (mr *MockKeyManagerClientMockRecorder) GetKey(ctx, storeName, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKey", reflect.TypeOf((*MockKeyManagerClient)(nil).GetKey), ctx, storeName, id)
}

// ListKeys mocks base method
func (m *MockKeyManagerClient) ListKeys(ctx context.Context, storeName string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListKeys", ctx, storeName)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListKeys indicates an expected call of ListKeys
func (mr *MockKeyManagerClientMockRecorder) ListKeys(ctx, storeName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListKeys", reflect.TypeOf((*MockKeyManagerClient)(nil).ListKeys), ctx, storeName)
}

// DestroyKey mocks base method
func (m *MockKeyManagerClient) DestroyKey(ctx context.Context, storeName, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DestroyKey", ctx, storeName, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DestroyKey indicates an expected call of DestroyKey
func (mr *MockKeyManagerClientMockRecorder) DestroyKey(ctx, storeName, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DestroyKey", reflect.TypeOf((*MockKeyManagerClient)(nil).DestroyKey), ctx, storeName, id)
}

// CreateAccount mocks base method
func (m *MockKeyManagerClient) CreateAccount(ctx context.Context, storeName string, request *types.CreateEth1AccountRequest) (*types.Eth1AccountResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccount", ctx, storeName, request)
	ret0, _ := ret[0].(*types.Eth1AccountResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccount indicates an expected call of CreateAccount
func (mr *MockKeyManagerClientMockRecorder) CreateAccount(ctx, storeName, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockKeyManagerClient)(nil).CreateAccount), ctx, storeName, request)
}

// Call mocks base method
func (m *MockKeyManagerClient) Call(ctx context.Context, nodeID, method string, args ...interface{}) (*jsonrpc.ResponseMsg, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, nodeID, method}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Call", varargs...)
	ret0, _ := ret[0].(*jsonrpc.ResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Call indicates an expected call of Call
func (mr *MockKeyManagerClientMockRecorder) Call(ctx, nodeID, method interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, nodeID, method}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Call", reflect.TypeOf((*MockKeyManagerClient)(nil).Call), varargs...)
}
