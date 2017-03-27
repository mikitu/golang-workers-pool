// Automatically generated by MockGen. DO NOT EDIT!
// Source: src/tasks/task_interface.go

package tests

import (
	gomock "github.com/golang/mock/gomock"
)

// Mock of Task interface
type MockTask struct {
	ctrl     *gomock.Controller
	recorder *_MockTaskRecorder
}

// Recorder for MockTask (not exported)
type _MockTaskRecorder struct {
	mock *MockTask
}

func NewMockTask(ctrl *gomock.Controller) *MockTask {
	mock := &MockTask{ctrl: ctrl}
	mock.recorder = &_MockTaskRecorder{mock}
	return mock
}

func (_m *MockTask) EXPECT() *_MockTaskRecorder {
	return _m.recorder
}

func (_m *MockTask) Execute() {
	_m.ctrl.Call(_m, "Execute")
}

func (_mr *_MockTaskRecorder) Execute() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Execute")
}