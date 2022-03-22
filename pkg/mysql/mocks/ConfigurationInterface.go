// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	mysql "cabother-go-helpers/pkg/mysql"

	mock "github.com/stretchr/testify/mock"
)

// ConfigurationInterface is an autogenerated mock type for the ConfigurationInterface type
type ConfigurationInterface struct {
	mock.Mock
}

// GetMySQLConfigs provides a mock function with given fields:
func (_m *ConfigurationInterface) GetMySQLConfigs() *mysql.MySQLConfig {
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
