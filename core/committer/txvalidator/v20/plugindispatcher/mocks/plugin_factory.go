// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	validation "github.com/hxx258456/fabric/core/handlers/validation/api"
)

// PluginFactory is an autogenerated mock type for the PluginFactory type
type PluginFactory struct {
	mock.Mock
}

// New provides a mock function with given fields:
func (_m *PluginFactory) New() validation.Plugin {
	ret := _m.Called()

	var r0 validation.Plugin
	if rf, ok := ret.Get(0).(func() validation.Plugin); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(validation.Plugin)
		}
	}

	return r0
}
