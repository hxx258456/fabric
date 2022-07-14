// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"sync"

	"github.com/hxx258456/fabric/common/chaincode"
)

type MetadataHandler struct {
	InitializeMetadataStub        func(string, chaincode.MetadataSet)
	initializeMetadataMutex       sync.RWMutex
	initializeMetadataArgsForCall []struct {
		arg1 string
		arg2 chaincode.MetadataSet
	}
	UpdateMetadataStub        func(string, chaincode.MetadataSet)
	updateMetadataMutex       sync.RWMutex
	updateMetadataArgsForCall []struct {
		arg1 string
		arg2 chaincode.MetadataSet
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *MetadataHandler) InitializeMetadata(arg1 string, arg2 chaincode.MetadataSet) {
	fake.initializeMetadataMutex.Lock()
	fake.initializeMetadataArgsForCall = append(fake.initializeMetadataArgsForCall, struct {
		arg1 string
		arg2 chaincode.MetadataSet
	}{arg1, arg2})
	fake.recordInvocation("InitializeMetadata", []interface{}{arg1, arg2})
	fake.initializeMetadataMutex.Unlock()
	if fake.InitializeMetadataStub != nil {
		fake.InitializeMetadataStub(arg1, arg2)
	}
}

func (fake *MetadataHandler) InitializeMetadataCallCount() int {
	fake.initializeMetadataMutex.RLock()
	defer fake.initializeMetadataMutex.RUnlock()
	return len(fake.initializeMetadataArgsForCall)
}

func (fake *MetadataHandler) InitializeMetadataCalls(stub func(string, chaincode.MetadataSet)) {
	fake.initializeMetadataMutex.Lock()
	defer fake.initializeMetadataMutex.Unlock()
	fake.InitializeMetadataStub = stub
}

func (fake *MetadataHandler) InitializeMetadataArgsForCall(i int) (string, chaincode.MetadataSet) {
	fake.initializeMetadataMutex.RLock()
	defer fake.initializeMetadataMutex.RUnlock()
	argsForCall := fake.initializeMetadataArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *MetadataHandler) UpdateMetadata(arg1 string, arg2 chaincode.MetadataSet) {
	fake.updateMetadataMutex.Lock()
	fake.updateMetadataArgsForCall = append(fake.updateMetadataArgsForCall, struct {
		arg1 string
		arg2 chaincode.MetadataSet
	}{arg1, arg2})
	fake.recordInvocation("UpdateMetadata", []interface{}{arg1, arg2})
	fake.updateMetadataMutex.Unlock()
	if fake.UpdateMetadataStub != nil {
		fake.UpdateMetadataStub(arg1, arg2)
	}
}

func (fake *MetadataHandler) UpdateMetadataCallCount() int {
	fake.updateMetadataMutex.RLock()
	defer fake.updateMetadataMutex.RUnlock()
	return len(fake.updateMetadataArgsForCall)
}

func (fake *MetadataHandler) UpdateMetadataCalls(stub func(string, chaincode.MetadataSet)) {
	fake.updateMetadataMutex.Lock()
	defer fake.updateMetadataMutex.Unlock()
	fake.UpdateMetadataStub = stub
}

func (fake *MetadataHandler) UpdateMetadataArgsForCall(i int) (string, chaincode.MetadataSet) {
	fake.updateMetadataMutex.RLock()
	defer fake.updateMetadataMutex.RUnlock()
	argsForCall := fake.updateMetadataArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *MetadataHandler) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.initializeMetadataMutex.RLock()
	defer fake.initializeMetadataMutex.RUnlock()
	fake.updateMetadataMutex.RLock()
	defer fake.updateMetadataMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *MetadataHandler) recordInvocation(key string, args []interface{}) {
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
