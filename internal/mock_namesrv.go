/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by MockGen. DO NOT EDIT.
// Source: namesrv.go

// Package internal is a generated GoMock package.
package internal

import (
	primitive "github.com/apache/rocketmq-client-go/primitive"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockNamesrvs is a mock of Namesrvs interface
type MockNamesrvs struct {
	ctrl     *gomock.Controller
	recorder *MockNamesrvsMockRecorder
}

// MockNamesrvsMockRecorder is the mock recorder for MockNamesrvs
type MockNamesrvsMockRecorder struct {
	mock *MockNamesrvs
}

// NewMockNamesrvs creates a new mock instance
func NewMockNamesrvs(ctrl *gomock.Controller) *MockNamesrvs {
	mock := &MockNamesrvs{ctrl: ctrl}
	mock.recorder = &MockNamesrvsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNamesrvs) EXPECT() *MockNamesrvsMockRecorder {
	return m.recorder
}

// AddBroker mocks base method
func (m *MockNamesrvs) AddBroker(routeData *TopicRouteData) {
	m.ctrl.Call(m, "AddBroker", routeData)
}

// AddBroker indicates an expected call of AddBroker
func (mr *MockNamesrvsMockRecorder) AddBroker(routeData interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddBroker", reflect.TypeOf((*MockNamesrvs)(nil).AddBroker), routeData)
}

// cleanOfflineBroker mocks base method
func (m *MockNamesrvs) cleanOfflineBroker() {
	m.ctrl.Call(m, "cleanOfflineBroker")
}

// cleanOfflineBroker indicates an expected call of cleanOfflineBroker
func (mr *MockNamesrvsMockRecorder) cleanOfflineBroker() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "cleanOfflineBroker", reflect.TypeOf((*MockNamesrvs)(nil).cleanOfflineBroker))
}

// UpdateTopicRouteInfo mocks base method
func (m *MockNamesrvs) UpdateTopicRouteInfo(topic string) *TopicRouteData {
	ret := m.ctrl.Call(m, "UpdateTopicRouteInfo", topic)
	ret0, _ := ret[0].(*TopicRouteData)
	return ret0
}

// UpdateTopicRouteInfo indicates an expected call of UpdateTopicRouteInfo
func (mr *MockNamesrvsMockRecorder) UpdateTopicRouteInfo(topic interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTopicRouteInfo", reflect.TypeOf((*MockNamesrvs)(nil).UpdateTopicRouteInfo), topic)
}

// FetchPublishMessageQueues mocks base method
func (m *MockNamesrvs) FetchPublishMessageQueues(topic string) ([]*primitive.MessageQueue, error) {
	ret := m.ctrl.Call(m, "FetchPublishMessageQueues", topic)
	ret0, _ := ret[0].([]*primitive.MessageQueue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchPublishMessageQueues indicates an expected call of FetchPublishMessageQueues
func (mr *MockNamesrvsMockRecorder) FetchPublishMessageQueues(topic interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchPublishMessageQueues", reflect.TypeOf((*MockNamesrvs)(nil).FetchPublishMessageQueues), topic)
}

// FindBrokerAddrByTopic mocks base method
func (m *MockNamesrvs) FindBrokerAddrByTopic(topic string) string {
	ret := m.ctrl.Call(m, "FindBrokerAddrByTopic", topic)
	ret0, _ := ret[0].(string)
	return ret0
}

// FindBrokerAddrByTopic indicates an expected call of FindBrokerAddrByTopic
func (mr *MockNamesrvsMockRecorder) FindBrokerAddrByTopic(topic interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindBrokerAddrByTopic", reflect.TypeOf((*MockNamesrvs)(nil).FindBrokerAddrByTopic), topic)
}

// FindBrokerAddrByName mocks base method
func (m *MockNamesrvs) FindBrokerAddrByName(brokerName string) string {
	ret := m.ctrl.Call(m, "FindBrokerAddrByName", brokerName)
	ret0, _ := ret[0].(string)
	return ret0
}

// FindBrokerAddrByName indicates an expected call of FindBrokerAddrByName
func (mr *MockNamesrvsMockRecorder) FindBrokerAddrByName(brokerName interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindBrokerAddrByName", reflect.TypeOf((*MockNamesrvs)(nil).FindBrokerAddrByName), brokerName)
}

// FindBrokerAddressInSubscribe mocks base method
func (m *MockNamesrvs) FindBrokerAddressInSubscribe(brokerName string, brokerId int64, onlyThisBroker bool) *FindBrokerResult {
	ret := m.ctrl.Call(m, "FindBrokerAddressInSubscribe", brokerName, brokerId, onlyThisBroker)
	ret0, _ := ret[0].(*FindBrokerResult)
	return ret0
}

// FindBrokerAddressInSubscribe indicates an expected call of FindBrokerAddressInSubscribe
func (mr *MockNamesrvsMockRecorder) FindBrokerAddressInSubscribe(brokerName, brokerId, onlyThisBroker interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindBrokerAddressInSubscribe", reflect.TypeOf((*MockNamesrvs)(nil).FindBrokerAddressInSubscribe), brokerName, brokerId, onlyThisBroker)
}

// FetchSubscribeMessageQueues mocks base method
func (m *MockNamesrvs) FetchSubscribeMessageQueues(topic string) ([]*primitive.MessageQueue, error) {
	ret := m.ctrl.Call(m, "FetchSubscribeMessageQueues", topic)
	ret0, _ := ret[0].([]*primitive.MessageQueue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchSubscribeMessageQueues indicates an expected call of FetchSubscribeMessageQueues
func (mr *MockNamesrvsMockRecorder) FetchSubscribeMessageQueues(topic interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchSubscribeMessageQueues", reflect.TypeOf((*MockNamesrvs)(nil).FetchSubscribeMessageQueues), topic)
}