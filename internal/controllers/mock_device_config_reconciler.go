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
// Source: device_config_reconciler.go
//
// Generated by this command:
//
//	mockgen -source=device_config_reconciler.go -package=controllers -destination=mock_device_config_reconciler.go deviceConfigReconcilerHelperAPI
//
// Package controllers is a generated GoMock package.
package controllers

import (
	context "context"
	reflect "reflect"

	v1alpha1 "github.com/ROCm/gpu-operator/api/v1alpha1"
	v1beta1 "github.com/rh-ecosystem-edge/kernel-module-management/api/v1beta1"
	gomock "go.uber.org/mock/gomock"
	v1 "k8s.io/api/core/v1"
	v10 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	controllerruntime "sigs.k8s.io/controller-runtime"
	client "sigs.k8s.io/controller-runtime/pkg/client"
	reconcile "sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// MockdeviceConfigReconcilerHelperAPI is a mock of deviceConfigReconcilerHelperAPI interface.
type MockdeviceConfigReconcilerHelperAPI struct {
	ctrl     *gomock.Controller
	recorder *MockdeviceConfigReconcilerHelperAPIMockRecorder
}

// MockdeviceConfigReconcilerHelperAPIMockRecorder is the mock recorder for MockdeviceConfigReconcilerHelperAPI.
type MockdeviceConfigReconcilerHelperAPIMockRecorder struct {
	mock *MockdeviceConfigReconcilerHelperAPI
}

// NewMockdeviceConfigReconcilerHelperAPI creates a new mock instance.
func NewMockdeviceConfigReconcilerHelperAPI(ctrl *gomock.Controller) *MockdeviceConfigReconcilerHelperAPI {
	mock := &MockdeviceConfigReconcilerHelperAPI{ctrl: ctrl}
	mock.recorder = &MockdeviceConfigReconcilerHelperAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockdeviceConfigReconcilerHelperAPI) EXPECT() *MockdeviceConfigReconcilerHelperAPIMockRecorder {
	return m.recorder
}

// buildDeviceConfigStatus mocks base method.
func (m *MockdeviceConfigReconcilerHelperAPI) buildDeviceConfigStatus(ctx context.Context, devConfig *v1alpha1.DeviceConfig, nodes *v1.NodeList) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "buildDeviceConfigStatus", ctx, devConfig, nodes)
	ret0, _ := ret[0].(error)
	return ret0
}

// buildDeviceConfigStatus indicates an expected call of buildDeviceConfigStatus.
func (mr *MockdeviceConfigReconcilerHelperAPIMockRecorder) buildDeviceConfigStatus(ctx, devConfig, nodes any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "buildDeviceConfigStatus", reflect.TypeOf((*MockdeviceConfigReconcilerHelperAPI)(nil).buildDeviceConfigStatus), ctx, devConfig, nodes)
}

// buildNodeAssignments mocks base method.
func (m *MockdeviceConfigReconcilerHelperAPI) buildNodeAssignments(deviceConfigList *v1alpha1.DeviceConfigList) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "buildNodeAssignments", deviceConfigList)
	ret0, _ := ret[0].(error)
	return ret0
}

// buildNodeAssignments indicates an expected call of buildNodeAssignments.
func (mr *MockdeviceConfigReconcilerHelperAPIMockRecorder) buildNodeAssignments(deviceConfigList any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "buildNodeAssignments", reflect.TypeOf((*MockdeviceConfigReconcilerHelperAPI)(nil).buildNodeAssignments), deviceConfigList)
}

// deleteCondition mocks base method.
func (m *MockdeviceConfigReconcilerHelperAPI) deleteCondition(ctx context.Context, condition string, devConfig *v1alpha1.DeviceConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "deleteCondition", ctx, condition, devConfig)
	ret0, _ := ret[0].(error)
	return ret0
}

// deleteCondition indicates an expected call of deleteCondition.
func (mr *MockdeviceConfigReconcilerHelperAPIMockRecorder) deleteCondition(ctx, condition, devConfig any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "deleteCondition", reflect.TypeOf((*MockdeviceConfigReconcilerHelperAPI)(nil).deleteCondition), ctx, condition, devConfig)
}

// finalizeDeviceConfig mocks base method.
func (m *MockdeviceConfigReconcilerHelperAPI) finalizeDeviceConfig(ctx context.Context, devConfig *v1alpha1.DeviceConfig, nodes *v1.NodeList) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "finalizeDeviceConfig", ctx, devConfig, nodes)
	ret0, _ := ret[0].(error)
	return ret0
}

// finalizeDeviceConfig indicates an expected call of finalizeDeviceConfig.
func (mr *MockdeviceConfigReconcilerHelperAPIMockRecorder) finalizeDeviceConfig(ctx, devConfig, nodes any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "finalizeDeviceConfig", reflect.TypeOf((*MockdeviceConfigReconcilerHelperAPI)(nil).finalizeDeviceConfig), ctx, devConfig, nodes)
}

// findDeviceConfigsForNMC mocks base method.
func (m *MockdeviceConfigReconcilerHelperAPI) findDeviceConfigsForNMC(ctx context.Context, nmc client.Object) []reconcile.Request {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "findDeviceConfigsForNMC", ctx, nmc)
	ret0, _ := ret[0].([]reconcile.Request)
	return ret0
}

// findDeviceConfigsForNMC indicates an expected call of findDeviceConfigsForNMC.
func (mr *MockdeviceConfigReconcilerHelperAPIMockRecorder) findDeviceConfigsForNMC(ctx, nmc any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "findDeviceConfigsForNMC", reflect.TypeOf((*MockdeviceConfigReconcilerHelperAPI)(nil).findDeviceConfigsForNMC), ctx, nmc)
}

// findDeviceConfigsForSecret mocks base method.
func (m *MockdeviceConfigReconcilerHelperAPI) findDeviceConfigsForSecret(ctx context.Context, secret client.Object) []reconcile.Request {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "findDeviceConfigsForSecret", ctx, secret)
	ret0, _ := ret[0].([]reconcile.Request)
	return ret0
}

// findDeviceConfigsForSecret indicates an expected call of findDeviceConfigsForSecret.
func (mr *MockdeviceConfigReconcilerHelperAPIMockRecorder) findDeviceConfigsForSecret(ctx, secret any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "findDeviceConfigsForSecret", reflect.TypeOf((*MockdeviceConfigReconcilerHelperAPI)(nil).findDeviceConfigsForSecret), ctx, secret)
}

// findDeviceConfigsWithKMM mocks base method.
func (m *MockdeviceConfigReconcilerHelperAPI) findDeviceConfigsWithKMM(ctx context.Context, node client.Object) []reconcile.Request {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "findDeviceConfigsWithKMM", ctx, node)
	ret0, _ := ret[0].([]reconcile.Request)
	return ret0
}

// findDeviceConfigsWithKMM indicates an expected call of findDeviceConfigsWithKMM.
func (mr *MockdeviceConfigReconcilerHelperAPIMockRecorder) findDeviceConfigsWithKMM(ctx, node any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "findDeviceConfigsWithKMM", reflect.TypeOf((*MockdeviceConfigReconcilerHelperAPI)(nil).findDeviceConfigsWithKMM), ctx, node)
}

// getDeviceConfigOwnedKMMModule mocks base method.
func (m *MockdeviceConfigReconcilerHelperAPI) getDeviceConfigOwnedKMMModule(ctx context.Context, devConfig *v1alpha1.DeviceConfig) (*v1beta1.Module, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getDeviceConfigOwnedKMMModule", ctx, devConfig)
	ret0, _ := ret[0].(*v1beta1.Module)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// getDeviceConfigOwnedKMMModule indicates an expected call of getDeviceConfigOwnedKMMModule.
func (mr *MockdeviceConfigReconcilerHelperAPIMockRecorder) getDeviceConfigOwnedKMMModule(ctx, devConfig any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getDeviceConfigOwnedKMMModule", reflect.TypeOf((*MockdeviceConfigReconcilerHelperAPI)(nil).getDeviceConfigOwnedKMMModule), ctx, devConfig)
}

// getRequestedDeviceConfig mocks base method.
func (m *MockdeviceConfigReconcilerHelperAPI) getRequestedDeviceConfig(ctx context.Context, namespacedName types.NamespacedName) (*v1alpha1.DeviceConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getRequestedDeviceConfig", ctx, namespacedName)
	ret0, _ := ret[0].(*v1alpha1.DeviceConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// getRequestedDeviceConfig indicates an expected call of getRequestedDeviceConfig.
func (mr *MockdeviceConfigReconcilerHelperAPIMockRecorder) getRequestedDeviceConfig(ctx, namespacedName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getRequestedDeviceConfig", reflect.TypeOf((*MockdeviceConfigReconcilerHelperAPI)(nil).getRequestedDeviceConfig), ctx, namespacedName)
}

// handleBuildConfigMap mocks base method.
func (m *MockdeviceConfigReconcilerHelperAPI) handleBuildConfigMap(ctx context.Context, devConfig *v1alpha1.DeviceConfig, nodes *v1.NodeList) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "handleBuildConfigMap", ctx, devConfig, nodes)
	ret0, _ := ret[0].(error)
	return ret0
}

// handleBuildConfigMap indicates an expected call of handleBuildConfigMap.
func (mr *MockdeviceConfigReconcilerHelperAPIMockRecorder) handleBuildConfigMap(ctx, devConfig, nodes any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "handleBuildConfigMap", reflect.TypeOf((*MockdeviceConfigReconcilerHelperAPI)(nil).handleBuildConfigMap), ctx, devConfig, nodes)
}

// handleConfigManager mocks base method.
func (m *MockdeviceConfigReconcilerHelperAPI) handleConfigManager(ctx context.Context, devConfig *v1alpha1.DeviceConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "handleConfigManager", ctx, devConfig)
	ret0, _ := ret[0].(error)
	return ret0
}

// handleConfigManager indicates an expected call of handleConfigManager.
func (mr *MockdeviceConfigReconcilerHelperAPIMockRecorder) handleConfigManager(ctx, devConfig any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "handleConfigManager", reflect.TypeOf((*MockdeviceConfigReconcilerHelperAPI)(nil).handleConfigManager), ctx, devConfig)
}

// handleDevicePlugin mocks base method.
func (m *MockdeviceConfigReconcilerHelperAPI) handleDevicePlugin(ctx context.Context, devConfig *v1alpha1.DeviceConfig, nodes *v1.NodeList) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "handleDevicePlugin", ctx, devConfig, nodes)
	ret0, _ := ret[0].(error)
	return ret0
}

// handleDevicePlugin indicates an expected call of handleDevicePlugin.
func (mr *MockdeviceConfigReconcilerHelperAPIMockRecorder) handleDevicePlugin(ctx, devConfig, nodes any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "handleDevicePlugin", reflect.TypeOf((*MockdeviceConfigReconcilerHelperAPI)(nil).handleDevicePlugin), ctx, devConfig, nodes)
}

// handleKMMModule mocks base method.
func (m *MockdeviceConfigReconcilerHelperAPI) handleKMMModule(ctx context.Context, devConfig *v1alpha1.DeviceConfig, nodes *v1.NodeList) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "handleKMMModule", ctx, devConfig, nodes)
	ret0, _ := ret[0].(error)
	return ret0
}

// handleKMMModule indicates an expected call of handleKMMModule.
func (mr *MockdeviceConfigReconcilerHelperAPIMockRecorder) handleKMMModule(ctx, devConfig, nodes any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "handleKMMModule", reflect.TypeOf((*MockdeviceConfigReconcilerHelperAPI)(nil).handleKMMModule), ctx, devConfig, nodes)
}

// handleKMMVersionLabel mocks base method.
func (m *MockdeviceConfigReconcilerHelperAPI) handleKMMVersionLabel(ctx context.Context, devConfig *v1alpha1.DeviceConfig, nodes *v1.NodeList) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "handleKMMVersionLabel", ctx, devConfig, nodes)
	ret0, _ := ret[0].(error)
	return ret0
}

// handleKMMVersionLabel indicates an expected call of handleKMMVersionLabel.
func (mr *MockdeviceConfigReconcilerHelperAPIMockRecorder) handleKMMVersionLabel(ctx, devConfig, nodes any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "handleKMMVersionLabel", reflect.TypeOf((*MockdeviceConfigReconcilerHelperAPI)(nil).handleKMMVersionLabel), ctx, devConfig, nodes)
}

// handleMetricsExporter mocks base method.
func (m *MockdeviceConfigReconcilerHelperAPI) handleMetricsExporter(ctx context.Context, devConfig *v1alpha1.DeviceConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "handleMetricsExporter", ctx, devConfig)
	ret0, _ := ret[0].(error)
	return ret0
}

// handleMetricsExporter indicates an expected call of handleMetricsExporter.
func (mr *MockdeviceConfigReconcilerHelperAPIMockRecorder) handleMetricsExporter(ctx, devConfig any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "handleMetricsExporter", reflect.TypeOf((*MockdeviceConfigReconcilerHelperAPI)(nil).handleMetricsExporter), ctx, devConfig)
}

// handleModuleUpgrade mocks base method.
func (m *MockdeviceConfigReconcilerHelperAPI) handleModuleUpgrade(ctx context.Context, devConfig *v1alpha1.DeviceConfig, nodes *v1.NodeList, delete bool) (controllerruntime.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "handleModuleUpgrade", ctx, devConfig, nodes, delete)
	ret0, _ := ret[0].(controllerruntime.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// handleModuleUpgrade indicates an expected call of handleModuleUpgrade.
func (mr *MockdeviceConfigReconcilerHelperAPIMockRecorder) handleModuleUpgrade(ctx, devConfig, nodes, delete any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "handleModuleUpgrade", reflect.TypeOf((*MockdeviceConfigReconcilerHelperAPI)(nil).handleModuleUpgrade), ctx, devConfig, nodes, delete)
}

// handleNodeLabeller mocks base method.
func (m *MockdeviceConfigReconcilerHelperAPI) handleNodeLabeller(ctx context.Context, devConfig *v1alpha1.DeviceConfig, nodes *v1.NodeList) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "handleNodeLabeller", ctx, devConfig, nodes)
	ret0, _ := ret[0].(error)
	return ret0
}

// handleNodeLabeller indicates an expected call of handleNodeLabeller.
func (mr *MockdeviceConfigReconcilerHelperAPIMockRecorder) handleNodeLabeller(ctx, devConfig, nodes any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "handleNodeLabeller", reflect.TypeOf((*MockdeviceConfigReconcilerHelperAPI)(nil).handleNodeLabeller), ctx, devConfig, nodes)
}

// handleTestRunner mocks base method.
func (m *MockdeviceConfigReconcilerHelperAPI) handleTestRunner(ctx context.Context, devConfig *v1alpha1.DeviceConfig, nodes *v1.NodeList) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "handleTestRunner", ctx, devConfig, nodes)
	ret0, _ := ret[0].(error)
	return ret0
}

// handleTestRunner indicates an expected call of handleTestRunner.
func (mr *MockdeviceConfigReconcilerHelperAPIMockRecorder) handleTestRunner(ctx, devConfig, nodes any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "handleTestRunner", reflect.TypeOf((*MockdeviceConfigReconcilerHelperAPI)(nil).handleTestRunner), ctx, devConfig, nodes)
}

// listDeviceConfigs mocks base method.
func (m *MockdeviceConfigReconcilerHelperAPI) listDeviceConfigs(ctx context.Context) (*v1alpha1.DeviceConfigList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "listDeviceConfigs", ctx)
	ret0, _ := ret[0].(*v1alpha1.DeviceConfigList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// listDeviceConfigs indicates an expected call of listDeviceConfigs.
func (mr *MockdeviceConfigReconcilerHelperAPIMockRecorder) listDeviceConfigs(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "listDeviceConfigs", reflect.TypeOf((*MockdeviceConfigReconcilerHelperAPI)(nil).listDeviceConfigs), ctx)
}

// setCondition mocks base method.
func (m *MockdeviceConfigReconcilerHelperAPI) setCondition(ctx context.Context, condition string, devConfig *v1alpha1.DeviceConfig, status v10.ConditionStatus, reason, message string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "setCondition", ctx, condition, devConfig, status, reason, message)
	ret0, _ := ret[0].(error)
	return ret0
}

// setCondition indicates an expected call of setCondition.
func (mr *MockdeviceConfigReconcilerHelperAPIMockRecorder) setCondition(ctx, condition, devConfig, status, reason, message any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "setCondition", reflect.TypeOf((*MockdeviceConfigReconcilerHelperAPI)(nil).setCondition), ctx, condition, devConfig, status, reason, message)
}

// setFinalizer mocks base method.
func (m *MockdeviceConfigReconcilerHelperAPI) setFinalizer(ctx context.Context, devConfig *v1alpha1.DeviceConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "setFinalizer", ctx, devConfig)
	ret0, _ := ret[0].(error)
	return ret0
}

// setFinalizer indicates an expected call of setFinalizer.
func (mr *MockdeviceConfigReconcilerHelperAPIMockRecorder) setFinalizer(ctx, devConfig any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "setFinalizer", reflect.TypeOf((*MockdeviceConfigReconcilerHelperAPI)(nil).setFinalizer), ctx, devConfig)
}

// updateDeviceConfigStatus mocks base method.
func (m *MockdeviceConfigReconcilerHelperAPI) updateDeviceConfigStatus(ctx context.Context, devConfig *v1alpha1.DeviceConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "updateDeviceConfigStatus", ctx, devConfig)
	ret0, _ := ret[0].(error)
	return ret0
}

// updateDeviceConfigStatus indicates an expected call of updateDeviceConfigStatus.
func (mr *MockdeviceConfigReconcilerHelperAPIMockRecorder) updateDeviceConfigStatus(ctx, devConfig any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "updateDeviceConfigStatus", reflect.TypeOf((*MockdeviceConfigReconcilerHelperAPI)(nil).updateDeviceConfigStatus), ctx, devConfig)
}

// updateNodeAssignments mocks base method.
func (m *MockdeviceConfigReconcilerHelperAPI) updateNodeAssignments(namespacedName string, nodes *v1.NodeList, isFinalizer bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "updateNodeAssignments", namespacedName, nodes, isFinalizer)
}

// updateNodeAssignments indicates an expected call of updateNodeAssignments.
func (mr *MockdeviceConfigReconcilerHelperAPIMockRecorder) updateNodeAssignments(namespacedName, nodes, isFinalizer any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "updateNodeAssignments", reflect.TypeOf((*MockdeviceConfigReconcilerHelperAPI)(nil).updateNodeAssignments), namespacedName, nodes, isFinalizer)
}

// validateDeviceConfig mocks base method.
func (m *MockdeviceConfigReconcilerHelperAPI) validateDeviceConfig(ctx context.Context, devConfig *v1alpha1.DeviceConfig) []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "validateDeviceConfig", ctx, devConfig)
	ret0, _ := ret[0].([]string)
	return ret0
}

// validateDeviceConfig indicates an expected call of validateDeviceConfig.
func (mr *MockdeviceConfigReconcilerHelperAPIMockRecorder) validateDeviceConfig(ctx, devConfig any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "validateDeviceConfig", reflect.TypeOf((*MockdeviceConfigReconcilerHelperAPI)(nil).validateDeviceConfig), ctx, devConfig)
}

// validateNodeAssignments mocks base method.
func (m *MockdeviceConfigReconcilerHelperAPI) validateNodeAssignments(namespacedName string, nodes *v1.NodeList) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "validateNodeAssignments", namespacedName, nodes)
	ret0, _ := ret[0].(error)
	return ret0
}

// validateNodeAssignments indicates an expected call of validateNodeAssignments.
func (mr *MockdeviceConfigReconcilerHelperAPIMockRecorder) validateNodeAssignments(namespacedName, nodes any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "validateNodeAssignments", reflect.TypeOf((*MockdeviceConfigReconcilerHelperAPI)(nil).validateNodeAssignments), namespacedName, nodes)
}
