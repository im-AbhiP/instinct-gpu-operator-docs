/*
Copyright (c) Advanced Micro Devices, Inc. All rights reserved.

Licensed under the Apache License, Version 2.0 (the \"License\");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an \"AS IS\" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by MockGen. DO NOT EDIT.
// Source: testrunner.go
//
// Generated by this command:
//
//	mockgen -source=testrunner.go -package=testrunner -destination=mock_testrunner.go TestRunner
//
// Package testrunner is a generated GoMock package.
package testrunner

import (
	reflect "reflect"

	v1alpha1 "github.com/ROCm/gpu-operator/api/v1alpha1"
	gomock "go.uber.org/mock/gomock"
	v1 "k8s.io/api/apps/v1"
)

// MockTestRunner is a mock of TestRunner interface.
type MockTestRunner struct {
	ctrl     *gomock.Controller
	recorder *MockTestRunnerMockRecorder
}

// MockTestRunnerMockRecorder is the mock recorder for MockTestRunner.
type MockTestRunnerMockRecorder struct {
	mock *MockTestRunner
}

// NewMockTestRunner creates a new mock instance.
func NewMockTestRunner(ctrl *gomock.Controller) *MockTestRunner {
	mock := &MockTestRunner{ctrl: ctrl}
	mock.recorder = &MockTestRunnerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTestRunner) EXPECT() *MockTestRunnerMockRecorder {
	return m.recorder
}

// SetTestRunnerAsDesired mocks base method.
func (m *MockTestRunner) SetTestRunnerAsDesired(ds *v1.DaemonSet, devConfig *v1alpha1.DeviceConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetTestRunnerAsDesired", ds, devConfig)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetTestRunnerAsDesired indicates an expected call of SetTestRunnerAsDesired.
func (mr *MockTestRunnerMockRecorder) SetTestRunnerAsDesired(ds, devConfig any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTestRunnerAsDesired", reflect.TypeOf((*MockTestRunner)(nil).SetTestRunnerAsDesired), ds, devConfig)
}
