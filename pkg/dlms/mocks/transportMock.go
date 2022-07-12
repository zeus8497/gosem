// Code generated by mockery v2.5.1. DO NOT EDIT.

package mocks

import (
	log "log"

	mock "github.com/stretchr/testify/mock"
)

// TransportMock is an autogenerated mock type for the Transport type
type TransportMock struct {
	mock.Mock
}

// Connect provides a mock function with given fields:
func (_m *TransportMock) Connect() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Disconnect provides a mock function with given fields:
func (_m *TransportMock) Disconnect() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IsConnected provides a mock function with given fields:
func (_m *TransportMock) IsConnected() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Send provides a mock function with given fields: src
func (_m *TransportMock) Send(src []byte) ([]byte, error) {
	ret := _m.Called(src)

	var r0 []byte
	if rf, ok := ret.Get(0).(func([]byte) []byte); ok {
		r0 = rf(src)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(src)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetAddress provides a mock function with given fields: client, server
func (_m *TransportMock) SetAddress(client int, server int) {
	_m.Called(client, server)
}

// SetLogger provides a mock function with given fields: logger
func (_m *TransportMock) SetLogger(logger *log.Logger) {
	_m.Called(logger)
}
