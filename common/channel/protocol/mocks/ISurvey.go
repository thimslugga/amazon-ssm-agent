// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	message "github.com/aws/amazon-ssm-agent/common/message"
	mock "github.com/stretchr/testify/mock"

	utils "github.com/aws/amazon-ssm-agent/common/channel/utils"
)

// ISurvey is an autogenerated mock type for the ISurvey type
type ISurvey struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *ISurvey) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Dial provides a mock function with given fields: path
func (_m *ISurvey) Dial(path string) error {
	ret := _m.Called(path)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetCommProtocolInfo provides a mock function with given fields:
func (_m *ISurvey) GetCommProtocolInfo() utils.SocketType {
	ret := _m.Called()

	var r0 utils.SocketType
	if rf, ok := ret.Get(0).(func() utils.SocketType); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(utils.SocketType)
	}

	return r0
}

// Initialize provides a mock function with given fields:
func (_m *ISurvey) Initialize() {
	_m.Called()
}

// Listen provides a mock function with given fields: address
func (_m *ISurvey) Listen(address string) error {
	ret := _m.Called(address)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(address)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Recv provides a mock function with given fields:
func (_m *ISurvey) Recv() ([]byte, error) {
	ret := _m.Called()

	var r0 []byte
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Send provides a mock function with given fields: _a0
func (_m *ISurvey) Send(_a0 *message.Message) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*message.Message) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetOption provides a mock function with given fields: name, value
func (_m *ISurvey) SetOption(name string, value interface{}) error {
	ret := _m.Called(name, value)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, interface{}) error); ok {
		r0 = rf(name, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}