// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	common "github.com/hxx258456/fabric-protos-go-cc/common"
	mock "github.com/stretchr/testify/mock"

	protoutil "github.com/hxx258456/fabric/protoutil"
)

// BlockVerifier is an autogenerated mock type for the BlockVerifier type
type BlockVerifier struct {
	mock.Mock
}

// VerifyBlockSignature provides a mock function with given fields: sd, config
func (_m *BlockVerifier) VerifyBlockSignature(sd []*protoutil.SignedData, config *common.ConfigEnvelope) error {
	ret := _m.Called(sd, config)

	var r0 error
	if rf, ok := ret.Get(0).(func([]*protoutil.SignedData, *common.ConfigEnvelope) error); ok {
		r0 = rf(sd, config)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
