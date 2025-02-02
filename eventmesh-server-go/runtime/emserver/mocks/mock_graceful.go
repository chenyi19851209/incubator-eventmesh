/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockGracefulServer is a mock of GracefulServer interface.
type MockGracefulServer struct {
	ctrl     *gomock.Controller
	recorder *MockGracefulServerMockRecorder
}

// MockGracefulServerMockRecorder is the mock recorder for MockGracefulServer.
type MockGracefulServerMockRecorder struct {
	mock *MockGracefulServer
}

// NewMockGracefulServer creates a new mock instance.
func NewMockGracefulServer(ctrl *gomock.Controller) *MockGracefulServer {
	mock := &MockGracefulServer{ctrl: ctrl}
	mock.recorder = &MockGracefulServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGracefulServer) EXPECT() *MockGracefulServerMockRecorder {
	return m.recorder
}

// Serve mocks base method.
func (m *MockGracefulServer) Serve() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Serve")
	ret0, _ := ret[0].(error)
	return ret0
}

// Serve indicates an expected call of Serve.
func (mr *MockGracefulServerMockRecorder) Serve() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Serve", reflect.TypeOf((*MockGracefulServer)(nil).Serve))
}

// Stop mocks base method.
func (m *MockGracefulServer) Stop() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop.
func (mr *MockGracefulServerMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockGracefulServer)(nil).Stop))
}
