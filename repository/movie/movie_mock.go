// Code generated by MockGen. DO NOT EDIT.
// Source: types.go

// Package repository is a generated GoMock package.
package repository

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIRepoMovie is a mock of IRepoMovie interface.
type MockIRepoMovie struct {
	ctrl     *gomock.Controller
	recorder *MockIRepoMovieMockRecorder
}

// MockIRepoMovieMockRecorder is the mock recorder for MockIRepoMovie.
type MockIRepoMovieMockRecorder struct {
	mock *MockIRepoMovie
}

// NewMockIRepoMovie creates a new mock instance.
func NewMockIRepoMovie(ctrl *gomock.Controller) *MockIRepoMovie {
	mock := &MockIRepoMovie{ctrl: ctrl}
	mock.recorder = &MockIRepoMovieMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRepoMovie) EXPECT() *MockIRepoMovieMockRecorder {
	return m.recorder
}

// GetOneMovie mocks base method.
func (m *MockIRepoMovie) GetOneMovie(imdbID string) Movie {
//	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOneMovie", imdbID)
	ret0, _ := ret[0].(Movie)
	return ret0
}

// GetOneMovie indicates an expected call of GetOneMovie.
func (mr *MockIRepoMovieMockRecorder) GetOneMovie(imdbID interface{}) *gomock.Call {
//	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOneMovie", reflect.TypeOf((*MockIRepoMovie)(nil).GetOneMovie), imdbID)
}

// SearchMovie mocks base method.
func (m *MockIRepoMovie) SearchMovie(key, page string) SeachMovie {
//	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchMovie", key, page)
	ret0, _ := ret[0].(SeachMovie)
	return ret0
}

// SearchMovie indicates an expected call of SearchMovie.
func (mr *MockIRepoMovieMockRecorder) SearchMovie(key, page interface{}) *gomock.Call {
//	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchMovie", reflect.TypeOf((*MockIRepoMovie)(nil).SearchMovie), key, page)
}