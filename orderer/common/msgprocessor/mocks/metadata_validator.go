// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	"github.com/hxx258456/fabric/common/channelconfig"
)

type MetadataValidator struct {
	ValidateConsensusMetadataStub        func(channelconfig.Orderer, channelconfig.Orderer, bool) error
	validateConsensusMetadataMutex       sync.RWMutex
	validateConsensusMetadataArgsForCall []struct {
		arg1 channelconfig.Orderer
		arg2 channelconfig.Orderer
		arg3 bool
	}
	validateConsensusMetadataReturns struct {
		result1 error
	}
	validateConsensusMetadataReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *MetadataValidator) ValidateConsensusMetadata(arg1 channelconfig.Orderer, arg2 channelconfig.Orderer, arg3 bool) error {
	fake.validateConsensusMetadataMutex.Lock()
	ret, specificReturn := fake.validateConsensusMetadataReturnsOnCall[len(fake.validateConsensusMetadataArgsForCall)]
	fake.validateConsensusMetadataArgsForCall = append(fake.validateConsensusMetadataArgsForCall, struct {
		arg1 channelconfig.Orderer
		arg2 channelconfig.Orderer
		arg3 bool
	}{arg1, arg2, arg3})
	fake.recordInvocation("ValidateConsensusMetadata", []interface{}{arg1, arg2, arg3})
	fake.validateConsensusMetadataMutex.Unlock()
	if fake.ValidateConsensusMetadataStub != nil {
		return fake.ValidateConsensusMetadataStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.validateConsensusMetadataReturns
	return fakeReturns.result1
}

func (fake *MetadataValidator) ValidateConsensusMetadataCallCount() int {
	fake.validateConsensusMetadataMutex.RLock()
	defer fake.validateConsensusMetadataMutex.RUnlock()
	return len(fake.validateConsensusMetadataArgsForCall)
}

func (fake *MetadataValidator) ValidateConsensusMetadataCalls(stub func(channelconfig.Orderer, channelconfig.Orderer, bool) error) {
	fake.validateConsensusMetadataMutex.Lock()
	defer fake.validateConsensusMetadataMutex.Unlock()
	fake.ValidateConsensusMetadataStub = stub
}

func (fake *MetadataValidator) ValidateConsensusMetadataArgsForCall(i int) (channelconfig.Orderer, channelconfig.Orderer, bool) {
	fake.validateConsensusMetadataMutex.RLock()
	defer fake.validateConsensusMetadataMutex.RUnlock()
	argsForCall := fake.validateConsensusMetadataArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *MetadataValidator) ValidateConsensusMetadataReturns(result1 error) {
	fake.validateConsensusMetadataMutex.Lock()
	defer fake.validateConsensusMetadataMutex.Unlock()
	fake.ValidateConsensusMetadataStub = nil
	fake.validateConsensusMetadataReturns = struct {
		result1 error
	}{result1}
}

func (fake *MetadataValidator) ValidateConsensusMetadataReturnsOnCall(i int, result1 error) {
	fake.validateConsensusMetadataMutex.Lock()
	defer fake.validateConsensusMetadataMutex.Unlock()
	fake.ValidateConsensusMetadataStub = nil
	if fake.validateConsensusMetadataReturnsOnCall == nil {
		fake.validateConsensusMetadataReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.validateConsensusMetadataReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *MetadataValidator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.validateConsensusMetadataMutex.RLock()
	defer fake.validateConsensusMetadataMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *MetadataValidator) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}
