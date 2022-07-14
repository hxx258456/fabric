// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import discovery "github.com/hxx258456/fabric/discovery/client"
import mock "github.com/stretchr/testify/mock"

// LocalResponse is an autogenerated mock type for the LocalResponse type
type LocalResponse struct {
	mock.Mock
}

// Peers provides a mock function with given fields:
func (_m *LocalResponse) Peers() ([]*discovery.Peer, error) {
	ret := _m.Called()

	var r0 []*discovery.Peer
	if rf, ok := ret.Get(0).(func() []*discovery.Peer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*discovery.Peer)
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
