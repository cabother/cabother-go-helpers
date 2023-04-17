// Code generated by mockery v2.24.0. DO NOT EDIT.

package mocks

import (
	mysql "github.com/cabother/cabother-go-helpers/pkg/mysql"
	mock "github.com/stretchr/testify/mock"
)

// ConfigurationInterface is an autogenerated mock type for the ConfigurationInterface type
type ConfigurationInterface struct {
	mock.Mock
}

// GetCustomMySQLConfigs provides a mock function with given fields: database, host, username, password
func (_m *ConfigurationInterface) GetCustomMySQLConfigs(database string, host string, username string, password string) *mysql.MySQLConfig {
	ret := _m.Called(database, host, username, password)

	var r0 *mysql.MySQLConfig
	if rf, ok := ret.Get(0).(func(string, string, string, string) *mysql.MySQLConfig); ok {
		r0 = rf(database, host, username, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mysql.MySQLConfig)
		}
	}

	return r0
}

// GetDefaultMySQLConfigs provides a mock function with given fields:
func (_m *ConfigurationInterface) GetDefaultMySQLConfigs() *mysql.MySQLConfig {
	ret := _m.Called()

	var r0 *mysql.MySQLConfig
	if rf, ok := ret.Get(0).(func() *mysql.MySQLConfig); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mysql.MySQLConfig)
		}
	}

	return r0
}

type mockConstructorTestingTNewConfigurationInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewConfigurationInterface creates a new instance of ConfigurationInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewConfigurationInterface(t mockConstructorTestingTNewConfigurationInterface) *ConfigurationInterface {
	mock := &ConfigurationInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
