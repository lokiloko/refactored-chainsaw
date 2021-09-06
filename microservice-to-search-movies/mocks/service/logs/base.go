// Code generated by MockGen. DO NOT EDIT.
// Source: service/logs/base.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockLogsService is a mock of LogsService interface
type MockLogsService struct {
	ctrl     *gomock.Controller
	recorder *MockLogsServiceMockRecorder
}

// MockLogsServiceMockRecorder is the mock recorder for MockLogsService
type MockLogsServiceMockRecorder struct {
	mock *MockLogsService
}

// NewMockLogsService creates a new mock instance
func NewMockLogsService(ctrl *gomock.Controller) *MockLogsService {
	mock := &MockLogsService{ctrl: ctrl}
	mock.recorder = &MockLogsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLogsService) EXPECT() *MockLogsServiceMockRecorder {
	return m.recorder
}

// WriteLog mocks base method
func (m *MockLogsService) WriteLog(request, response interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteLog", request, response)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteLog indicates an expected call of WriteLog
func (mr *MockLogsServiceMockRecorder) WriteLog(request, response interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteLog", reflect.TypeOf((*MockLogsService)(nil).WriteLog), request, response)
}
