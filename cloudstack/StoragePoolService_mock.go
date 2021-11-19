//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: ./cloudstack/StoragePoolService.go

// Package cloudstack is a generated GoMock package.
package cloudstack

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockStoragePoolServiceIface is a mock of StoragePoolServiceIface interface.
type MockStoragePoolServiceIface struct {
	ctrl     *gomock.Controller
	recorder *MockStoragePoolServiceIfaceMockRecorder
}

// MockStoragePoolServiceIfaceMockRecorder is the mock recorder for MockStoragePoolServiceIface.
type MockStoragePoolServiceIfaceMockRecorder struct {
	mock *MockStoragePoolServiceIface
}

// NewMockStoragePoolServiceIface creates a new mock instance.
func NewMockStoragePoolServiceIface(ctrl *gomock.Controller) *MockStoragePoolServiceIface {
	mock := &MockStoragePoolServiceIface{ctrl: ctrl}
	mock.recorder = &MockStoragePoolServiceIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStoragePoolServiceIface) EXPECT() *MockStoragePoolServiceIfaceMockRecorder {
	return m.recorder
}

// CancelStorageMaintenance mocks base method.
func (m *MockStoragePoolServiceIface) CancelStorageMaintenance(p *CancelStorageMaintenanceParams) (*CancelStorageMaintenanceResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelStorageMaintenance", p)
	ret0, _ := ret[0].(*CancelStorageMaintenanceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CancelStorageMaintenance indicates an expected call of CancelStorageMaintenance.
func (mr *MockStoragePoolServiceIfaceMockRecorder) CancelStorageMaintenance(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelStorageMaintenance", reflect.TypeOf((*MockStoragePoolServiceIface)(nil).CancelStorageMaintenance), p)
}

// EnableStorageMaintenance mocks base method.
func (m *MockStoragePoolServiceIface) EnableStorageMaintenance(p *EnableStorageMaintenanceParams) (*EnableStorageMaintenanceResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableStorageMaintenance", p)
	ret0, _ := ret[0].(*EnableStorageMaintenanceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnableStorageMaintenance indicates an expected call of EnableStorageMaintenance.
func (mr *MockStoragePoolServiceIfaceMockRecorder) EnableStorageMaintenance(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableStorageMaintenance", reflect.TypeOf((*MockStoragePoolServiceIface)(nil).EnableStorageMaintenance), p)
}

// ListStorageProviders mocks base method.
func (m *MockStoragePoolServiceIface) ListStorageProviders(p *ListStorageProvidersParams) (*ListStorageProvidersResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListStorageProviders", p)
	ret0, _ := ret[0].(*ListStorageProvidersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListStorageProviders indicates an expected call of ListStorageProviders.
func (mr *MockStoragePoolServiceIfaceMockRecorder) ListStorageProviders(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListStorageProviders", reflect.TypeOf((*MockStoragePoolServiceIface)(nil).ListStorageProviders), p)
}

// NewCancelStorageMaintenanceParams mocks base method.
func (m *MockStoragePoolServiceIface) NewCancelStorageMaintenanceParams(id string) *CancelStorageMaintenanceParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewCancelStorageMaintenanceParams", id)
	ret0, _ := ret[0].(*CancelStorageMaintenanceParams)
	return ret0
}

// NewCancelStorageMaintenanceParams indicates an expected call of NewCancelStorageMaintenanceParams.
func (mr *MockStoragePoolServiceIfaceMockRecorder) NewCancelStorageMaintenanceParams(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewCancelStorageMaintenanceParams", reflect.TypeOf((*MockStoragePoolServiceIface)(nil).NewCancelStorageMaintenanceParams), id)
}

// NewEnableStorageMaintenanceParams mocks base method.
func (m *MockStoragePoolServiceIface) NewEnableStorageMaintenanceParams(id string) *EnableStorageMaintenanceParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewEnableStorageMaintenanceParams", id)
	ret0, _ := ret[0].(*EnableStorageMaintenanceParams)
	return ret0
}

// NewEnableStorageMaintenanceParams indicates an expected call of NewEnableStorageMaintenanceParams.
func (mr *MockStoragePoolServiceIfaceMockRecorder) NewEnableStorageMaintenanceParams(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewEnableStorageMaintenanceParams", reflect.TypeOf((*MockStoragePoolServiceIface)(nil).NewEnableStorageMaintenanceParams), id)
}

// NewListStorageProvidersParams mocks base method.
func (m *MockStoragePoolServiceIface) NewListStorageProvidersParams(storagePoolType string) *ListStorageProvidersParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListStorageProvidersParams", storagePoolType)
	ret0, _ := ret[0].(*ListStorageProvidersParams)
	return ret0
}

// NewListStorageProvidersParams indicates an expected call of NewListStorageProvidersParams.
func (mr *MockStoragePoolServiceIfaceMockRecorder) NewListStorageProvidersParams(storagePoolType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListStorageProvidersParams", reflect.TypeOf((*MockStoragePoolServiceIface)(nil).NewListStorageProvidersParams), storagePoolType)
}