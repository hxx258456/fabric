// Code generated by counterfeiter. DO NOT EDIT.
package fake

import (
	"sync"

	"github.com/hxx258456/fabric-protos-go-cc/transientstore"
	"github.com/hxx258456/fabric/core/endorser"
)

type PrivateDataDistributor struct {
	DistributePrivateDataStub        func(string, string, *transientstore.TxPvtReadWriteSetWithConfigInfo, uint64) error
	distributePrivateDataMutex       sync.RWMutex
	distributePrivateDataArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 *transientstore.TxPvtReadWriteSetWithConfigInfo
		arg4 uint64
	}
	distributePrivateDataReturns struct {
		result1 error
	}
	distributePrivateDataReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *PrivateDataDistributor) DistributePrivateData(arg1 string, arg2 string, arg3 *transientstore.TxPvtReadWriteSetWithConfigInfo, arg4 uint64) error {
	fake.distributePrivateDataMutex.Lock()
	ret, specificReturn := fake.distributePrivateDataReturnsOnCall[len(fake.distributePrivateDataArgsForCall)]
	fake.distributePrivateDataArgsForCall = append(fake.distributePrivateDataArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 *transientstore.TxPvtReadWriteSetWithConfigInfo
		arg4 uint64
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("DistributePrivateData", []interface{}{arg1, arg2, arg3, arg4})
	fake.distributePrivateDataMutex.Unlock()
	if fake.DistributePrivateDataStub != nil {
		return fake.DistributePrivateDataStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.distributePrivateDataReturns
	return fakeReturns.result1
}

func (fake *PrivateDataDistributor) DistributePrivateDataCallCount() int {
	fake.distributePrivateDataMutex.RLock()
	defer fake.distributePrivateDataMutex.RUnlock()
	return len(fake.distributePrivateDataArgsForCall)
}

func (fake *PrivateDataDistributor) DistributePrivateDataCalls(stub func(string, string, *transientstore.TxPvtReadWriteSetWithConfigInfo, uint64) error) {
	fake.distributePrivateDataMutex.Lock()
	defer fake.distributePrivateDataMutex.Unlock()
	fake.DistributePrivateDataStub = stub
}

func (fake *PrivateDataDistributor) DistributePrivateDataArgsForCall(i int) (string, string, *transientstore.TxPvtReadWriteSetWithConfigInfo, uint64) {
	fake.distributePrivateDataMutex.RLock()
	defer fake.distributePrivateDataMutex.RUnlock()
	argsForCall := fake.distributePrivateDataArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *PrivateDataDistributor) DistributePrivateDataReturns(result1 error) {
	fake.distributePrivateDataMutex.Lock()
	defer fake.distributePrivateDataMutex.Unlock()
	fake.DistributePrivateDataStub = nil
	fake.distributePrivateDataReturns = struct {
		result1 error
	}{result1}
}

func (fake *PrivateDataDistributor) DistributePrivateDataReturnsOnCall(i int, result1 error) {
	fake.distributePrivateDataMutex.Lock()
	defer fake.distributePrivateDataMutex.Unlock()
	fake.DistributePrivateDataStub = nil
	if fake.distributePrivateDataReturnsOnCall == nil {
		fake.distributePrivateDataReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.distributePrivateDataReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *PrivateDataDistributor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.distributePrivateDataMutex.RLock()
	defer fake.distributePrivateDataMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *PrivateDataDistributor) recordInvocation(key string, args []interface{}) {
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

var _ endorser.PrivateDataDistributor = new(PrivateDataDistributor)
