// Code generated by MockGen. DO NOT EDIT.
// Source: google.golang.org/genproto/googleapis/longrunning (interfaces: OperationsClient)

// Package mock_longrunning is a generated GoMock package.
package mock_longrunning

import (
	gomock "github.com/golang/mock/gomock"
	empty "github.com/golang/protobuf/ptypes/empty"
	context "golang.org/x/net/context"
	longrunning "google.golang.org/genproto/googleapis/longrunning"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockOperationsClient is a mock of OperationsClient interface
type MockOperationsClient struct {
	ctrl     *gomock.Controller
	recorder *MockOperationsClientMockRecorder
}

// MockOperationsClientMockRecorder is the mock recorder for MockOperationsClient
type MockOperationsClientMockRecorder struct {
	mock *MockOperationsClient
}

// NewMockOperationsClient creates a new mock instance
func NewMockOperationsClient(ctrl *gomock.Controller) *MockOperationsClient {
	mock := &MockOperationsClient{ctrl: ctrl}
	mock.recorder = &MockOperationsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOperationsClient) EXPECT() *MockOperationsClientMockRecorder {
	return m.recorder
}

// CancelOperation mocks base method
func (m *MockOperationsClient) CancelOperation(arg0 context.Context, arg1 *longrunning.CancelOperationRequest, arg2 ...grpc.CallOption) (*empty.Empty, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CancelOperation", varargs...)
	ret0, _ := ret[0].(*empty.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CancelOperation indicates an expected call of CancelOperation
func (mr *MockOperationsClientMockRecorder) CancelOperation(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelOperation", reflect.TypeOf((*MockOperationsClient)(nil).CancelOperation), varargs...)
}

// DeleteOperation mocks base method
func (m *MockOperationsClient) DeleteOperation(arg0 context.Context, arg1 *longrunning.DeleteOperationRequest, arg2 ...grpc.CallOption) (*empty.Empty, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteOperation", varargs...)
	ret0, _ := ret[0].(*empty.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteOperation indicates an expected call of DeleteOperation
func (mr *MockOperationsClientMockRecorder) DeleteOperation(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOperation", reflect.TypeOf((*MockOperationsClient)(nil).DeleteOperation), varargs...)
}

// GetOperation mocks base method
func (m *MockOperationsClient) GetOperation(arg0 context.Context, arg1 *longrunning.GetOperationRequest, arg2 ...grpc.CallOption) (*longrunning.Operation, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetOperation", varargs...)
	ret0, _ := ret[0].(*longrunning.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOperation indicates an expected call of GetOperation
func (mr *MockOperationsClientMockRecorder) GetOperation(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperation", reflect.TypeOf((*MockOperationsClient)(nil).GetOperation), varargs...)
}

// ListOperations mocks base method
func (m *MockOperationsClient) ListOperations(arg0 context.Context, arg1 *longrunning.ListOperationsRequest, arg2 ...grpc.CallOption) (*longrunning.ListOperationsResponse, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListOperations", varargs...)
	ret0, _ := ret[0].(*longrunning.ListOperationsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOperations indicates an expected call of ListOperations
func (mr *MockOperationsClientMockRecorder) ListOperations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOperations", reflect.TypeOf((*MockOperationsClient)(nil).ListOperations), varargs...)
}
