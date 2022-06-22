// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// DatetimeInterface is an autogenerated mock type for the DatetimeInterface type
type DatetimeInterface struct {
	mock.Mock
}

// FormatDate provides a mock function with given fields: _a0, format
func (_m *DatetimeInterface) FormatDate(_a0 string, format string) (string, error) {
	ret := _m.Called(_a0, format)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(_a0, format)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(_a0, format)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FormatDatetime provides a mock function with given fields: _a0, format
func (_m *DatetimeInterface) FormatDatetime(_a0 time.Time, format string) string {
	ret := _m.Called(_a0, format)

	var r0 string
	if rf, ok := ret.Get(0).(func(time.Time, string) string); ok {
		r0 = rf(_a0, format)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetDateNow provides a mock function with given fields: format
func (_m *DatetimeInterface) GetDateNow(format string) string {
	ret := _m.Called(format)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(format)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type NewDatetimeInterfaceT interface {
	mock.TestingT
	Cleanup(func())
}

// NewDatetimeInterface creates a new instance of DatetimeInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDatetimeInterface(t NewDatetimeInterfaceT) *DatetimeInterface {
	mock := &DatetimeInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
