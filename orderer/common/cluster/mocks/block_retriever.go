// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	common "github.com/hxx258456/fabric-protos-go-cc/common"
	mock "github.com/stretchr/testify/mock"
)

// BlockRetriever is an autogenerated mock type for the BlockRetriever type
type BlockRetriever struct {
	mock.Mock
}

// Block provides a mock function with given fields: number
func (_m *BlockRetriever) Block(number uint64) *common.Block {
	ret := _m.Called(number)

	var r0 *common.Block
	if rf, ok := ret.Get(0).(func(uint64) *common.Block); ok {
		r0 = rf(number)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*common.Block)
		}
	}

	return r0
}
