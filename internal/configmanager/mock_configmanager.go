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
// Source: configmanager.go
//
// Generated by this command:
//
//	mockgen -source=configmanager.go -package=configmanager -destination=mock_configmanager.go ConfigManager
//
// Package configmanager is a generated GoMock package.
package configmanager

import (
	reflect "reflect"

	v1alpha1 "github.com/ROCm/gpu-operator/api/v1alpha1"
	gomock "go.uber.org/mock/gomock"
	v1 "k8s.io/api/apps/v1"
)

// MockConfigManager is a mock of ConfigManager interface.
type MockConfigManager struct {
	ctrl     *gomock.Controller
	recorder *MockConfigManagerMockRecorder
}

// MockConfigManagerMockRecorder is the mock recorder for MockConfigManager.
type MockConfigManagerMockRecorder struct {
	mock *MockConfigManager
}

// NewMockConfigManager creates a new mock instance.
func NewMockConfigManager(ctrl *gomock.Controller) *MockConfigManager {
	mock := &MockConfigManager{ctrl: ctrl}
	mock.recorder = &MockConfigManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConfigManager) EXPECT() *MockConfigManagerMockRecorder {
	return m.recorder
}

// SetConfigManagerAsDesired mocks base method.
func (m *MockConfigManager) SetConfigManagerAsDesired(ds *v1.DaemonSet, devConfig *v1alpha1.DeviceConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetConfigManagerAsDesired", ds, devConfig)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetConfigManagerAsDesired indicates an expected call of SetConfigManagerAsDesired.
func (mr *MockConfigManagerMockRecorder) SetConfigManagerAsDesired(ds, devConfig any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetConfigManagerAsDesired", reflect.TypeOf((*MockConfigManager)(nil).SetConfigManagerAsDesired), ds, devConfig)
}
