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
// Source: ./cloudstack/OauthService.go

// Package cloudstack is a generated GoMock package.
package cloudstack

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockOauthServiceIface is a mock of OauthServiceIface interface.
type MockOauthServiceIface struct {
	ctrl     *gomock.Controller
	recorder *MockOauthServiceIfaceMockRecorder
}

// MockOauthServiceIfaceMockRecorder is the mock recorder for MockOauthServiceIface.
type MockOauthServiceIfaceMockRecorder struct {
	mock *MockOauthServiceIface
}

// NewMockOauthServiceIface creates a new mock instance.
func NewMockOauthServiceIface(ctrl *gomock.Controller) *MockOauthServiceIface {
	mock := &MockOauthServiceIface{ctrl: ctrl}
	mock.recorder = &MockOauthServiceIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOauthServiceIface) EXPECT() *MockOauthServiceIfaceMockRecorder {
	return m.recorder
}

// DeleteOauthProvider mocks base method.
func (m *MockOauthServiceIface) DeleteOauthProvider(p *DeleteOauthProviderParams) (*DeleteOauthProviderResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOauthProvider", p)
	ret0, _ := ret[0].(*DeleteOauthProviderResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteOauthProvider indicates an expected call of DeleteOauthProvider.
func (mr *MockOauthServiceIfaceMockRecorder) DeleteOauthProvider(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOauthProvider", reflect.TypeOf((*MockOauthServiceIface)(nil).DeleteOauthProvider), p)
}

// GetOauthProviderByID mocks base method.
func (m *MockOauthServiceIface) GetOauthProviderByID(id string, opts ...OptionFunc) (*OauthProvider, int, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{id}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetOauthProviderByID", varargs...)
	ret0, _ := ret[0].(*OauthProvider)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetOauthProviderByID indicates an expected call of GetOauthProviderByID.
func (mr *MockOauthServiceIfaceMockRecorder) GetOauthProviderByID(id interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{id}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOauthProviderByID", reflect.TypeOf((*MockOauthServiceIface)(nil).GetOauthProviderByID), varargs...)
}

// GetOauthProviderByName mocks base method.
func (m *MockOauthServiceIface) GetOauthProviderByName(name string, opts ...OptionFunc) (*OauthProvider, int, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{name}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetOauthProviderByName", varargs...)
	ret0, _ := ret[0].(*OauthProvider)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetOauthProviderByName indicates an expected call of GetOauthProviderByName.
func (mr *MockOauthServiceIfaceMockRecorder) GetOauthProviderByName(name interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{name}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOauthProviderByName", reflect.TypeOf((*MockOauthServiceIface)(nil).GetOauthProviderByName), varargs...)
}

// GetOauthProviderID mocks base method.
func (m *MockOauthServiceIface) GetOauthProviderID(keyword string, opts ...OptionFunc) (string, int, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{keyword}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetOauthProviderID", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetOauthProviderID indicates an expected call of GetOauthProviderID.
func (mr *MockOauthServiceIfaceMockRecorder) GetOauthProviderID(keyword interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{keyword}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOauthProviderID", reflect.TypeOf((*MockOauthServiceIface)(nil).GetOauthProviderID), varargs...)
}

// ListOauthProvider mocks base method.
func (m *MockOauthServiceIface) ListOauthProvider(p *ListOauthProviderParams) (*ListOauthProviderResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListOauthProvider", p)
	ret0, _ := ret[0].(*ListOauthProviderResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOauthProvider indicates an expected call of ListOauthProvider.
func (mr *MockOauthServiceIfaceMockRecorder) ListOauthProvider(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOauthProvider", reflect.TypeOf((*MockOauthServiceIface)(nil).ListOauthProvider), p)
}

// NewDeleteOauthProviderParams mocks base method.
func (m *MockOauthServiceIface) NewDeleteOauthProviderParams(id string) *DeleteOauthProviderParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewDeleteOauthProviderParams", id)
	ret0, _ := ret[0].(*DeleteOauthProviderParams)
	return ret0
}

// NewDeleteOauthProviderParams indicates an expected call of NewDeleteOauthProviderParams.
func (mr *MockOauthServiceIfaceMockRecorder) NewDeleteOauthProviderParams(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewDeleteOauthProviderParams", reflect.TypeOf((*MockOauthServiceIface)(nil).NewDeleteOauthProviderParams), id)
}

// NewListOauthProviderParams mocks base method.
func (m *MockOauthServiceIface) NewListOauthProviderParams() *ListOauthProviderParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListOauthProviderParams")
	ret0, _ := ret[0].(*ListOauthProviderParams)
	return ret0
}

// NewListOauthProviderParams indicates an expected call of NewListOauthProviderParams.
func (mr *MockOauthServiceIfaceMockRecorder) NewListOauthProviderParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListOauthProviderParams", reflect.TypeOf((*MockOauthServiceIface)(nil).NewListOauthProviderParams))
}

// NewUpdateOauthProviderParams mocks base method.
func (m *MockOauthServiceIface) NewUpdateOauthProviderParams(id string) *UpdateOauthProviderParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewUpdateOauthProviderParams", id)
	ret0, _ := ret[0].(*UpdateOauthProviderParams)
	return ret0
}

// NewUpdateOauthProviderParams indicates an expected call of NewUpdateOauthProviderParams.
func (mr *MockOauthServiceIfaceMockRecorder) NewUpdateOauthProviderParams(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewUpdateOauthProviderParams", reflect.TypeOf((*MockOauthServiceIface)(nil).NewUpdateOauthProviderParams), id)
}

// UpdateOauthProvider mocks base method.
func (m *MockOauthServiceIface) UpdateOauthProvider(p *UpdateOauthProviderParams) (*UpdateOauthProviderResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOauthProvider", p)
	ret0, _ := ret[0].(*UpdateOauthProviderResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOauthProvider indicates an expected call of UpdateOauthProvider.
func (mr *MockOauthServiceIfaceMockRecorder) UpdateOauthProvider(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOauthProvider", reflect.TypeOf((*MockOauthServiceIface)(nil).UpdateOauthProvider), p)
}
