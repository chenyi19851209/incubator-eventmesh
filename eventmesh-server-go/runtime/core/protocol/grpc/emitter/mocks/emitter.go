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

	grpc "github.com/apache/incubator-eventmesh/eventmesh-server-go/pkg/common/protocol/grpc"
	pb "github.com/apache/incubator-eventmesh/eventmesh-server-go/runtime/proto/pb"
	gomock "github.com/golang/mock/gomock"
)

// MockEventEmitter is a mock of EventEmitter interface.
type MockEventEmitter struct {
	ctrl     *gomock.Controller
	recorder *MockEventEmitterMockRecorder
}

// MockEventEmitterMockRecorder is the mock recorder for MockEventEmitter.
type MockEventEmitterMockRecorder struct {
	mock *MockEventEmitter
}

// NewMockEventEmitter creates a new mock instance.
func NewMockEventEmitter(ctrl *gomock.Controller) *MockEventEmitter {
	mock := &MockEventEmitter{ctrl: ctrl}
	mock.recorder = &MockEventEmitterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventEmitter) EXPECT() *MockEventEmitterMockRecorder {
	return m.recorder
}

// SendStreamResp mocks base method.
func (m *MockEventEmitter) SendStreamResp(arg0 *pb.RequestHeader, arg1 *grpc.StatusCode) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendStreamResp", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendStreamResp indicates an expected call of SendStreamResp.
func (mr *MockEventEmitterMockRecorder) SendStreamResp(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendStreamResp", reflect.TypeOf((*MockEventEmitter)(nil).SendStreamResp), arg0, arg1)
}
