// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import fabric_protos_gomsp "github.com/hxx258456/fabric-protos-go-cc/msp"
import mock "github.com/stretchr/testify/mock"
import msp "github.com/hxx258456/fabric/msp"

import time "time"

// Identity is an autogenerated mock type for the Identity type
type Identity struct {
	mock.Mock
}

// Anonymous provides a mock function with given fields:
func (_m *Identity) Anonymous() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ExpiresAt provides a mock function with given fields:
func (_m *Identity) ExpiresAt() time.Time {
	ret := _m.Called()

	var r0 time.Time
	if rf, ok := ret.Get(0).(func() time.Time); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	return r0
}

// GetIdentifier provides a mock function with given fields:
func (_m *Identity) GetIdentifier() *msp.IdentityIdentifier {
	ret := _m.Called()

	var r0 *msp.IdentityIdentifier
	if rf, ok := ret.Get(0).(func() *msp.IdentityIdentifier); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*msp.IdentityIdentifier)
		}
	}

	return r0
}

// GetMSPIdentifier provides a mock function with given fields:
func (_m *Identity) GetMSPIdentifier() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetOrganizationalUnits provides a mock function with given fields:
func (_m *Identity) GetOrganizationalUnits() []*msp.OUIdentifier {
	ret := _m.Called()

	var r0 []*msp.OUIdentifier
	if rf, ok := ret.Get(0).(func() []*msp.OUIdentifier); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*msp.OUIdentifier)
		}
	}

	return r0
}

// SatisfiesPrincipal provides a mock function with given fields: principal
func (_m *Identity) SatisfiesPrincipal(principal *fabric_protos_gomsp.MSPPrincipal) error {
	ret := _m.Called(principal)

	var r0 error
	if rf, ok := ret.Get(0).(func(*fabric_protos_gomsp.MSPPrincipal) error); ok {
		r0 = rf(principal)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Serialize provides a mock function with given fields:
func (_m *Identity) Serialize() ([]byte, error) {
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

// Validate provides a mock function with given fields:
func (_m *Identity) Validate() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Verify provides a mock function with given fields: msg, sig
func (_m *Identity) Verify(msg []byte, sig []byte) error {
	ret := _m.Called(msg, sig)

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte, []byte) error); ok {
		r0 = rf(msg, sig)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
