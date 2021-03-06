// Code generated by MockGen. DO NOT EDIT.
// Source: handler/base.go

// Package mock_handler is a generated GoMock package.
package mock_handler

import (
	gomock "github.com/golang/mock/gomock"
	dto "github.com/lokiloko/refactored-chainsaw/microservice-to-search-movies/model/dto"
	reflect "reflect"
)

// MockHandler is a mock of Handler interface
type MockHandler struct {
	ctrl     *gomock.Controller
	recorder *MockHandlerMockRecorder
}

// MockHandlerMockRecorder is the mock recorder for MockHandler
type MockHandlerMockRecorder struct {
	mock *MockHandler
}

// NewMockHandler creates a new mock instance
func NewMockHandler(ctrl *gomock.Controller) *MockHandler {
	mock := &MockHandler{ctrl: ctrl}
	mock.recorder = &MockHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHandler) EXPECT() *MockHandlerMockRecorder {
	return m.recorder
}

// GetByIMDBID mocks base method
func (m *MockHandler) GetByIMDBID(id string) (dto.GetByIMDBIDResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByIMDBID", id)
	ret0, _ := ret[0].(dto.GetByIMDBIDResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByIMDBID indicates an expected call of GetByIMDBID
func (mr *MockHandlerMockRecorder) GetByIMDBID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByIMDBID", reflect.TypeOf((*MockHandler)(nil).GetByIMDBID), id)
}

// GetMoviesPaginated mocks base method
func (m *MockHandler) GetMoviesPaginated(page uint64, keyword string) (dto.GetMoviesPaginatedResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMoviesPaginated", page, keyword)
	ret0, _ := ret[0].(dto.GetMoviesPaginatedResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMoviesPaginated indicates an expected call of GetMoviesPaginated
func (mr *MockHandlerMockRecorder) GetMoviesPaginated(page, keyword interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMoviesPaginated", reflect.TypeOf((*MockHandler)(nil).GetMoviesPaginated), page, keyword)
}
