// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"sync"

	"github.com/hxx258456/fabric-protos-go-cc/peer"
	"github.com/hxx258456/fabric/core/chaincode/persistence"
)

type LegacyCCPackageLocator struct {
	GetChaincodeDepSpecStub        func(string) (*peer.ChaincodeDeploymentSpec, error)
	getChaincodeDepSpecMutex       sync.RWMutex
	getChaincodeDepSpecArgsForCall []struct {
		arg1 string
	}
	getChaincodeDepSpecReturns struct {
		result1 *peer.ChaincodeDeploymentSpec
		result2 error
	}
	getChaincodeDepSpecReturnsOnCall map[int]struct {
		result1 *peer.ChaincodeDeploymentSpec
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *LegacyCCPackageLocator) GetChaincodeDepSpec(arg1 string) (*peer.ChaincodeDeploymentSpec, error) {
	fake.getChaincodeDepSpecMutex.Lock()
	ret, specificReturn := fake.getChaincodeDepSpecReturnsOnCall[len(fake.getChaincodeDepSpecArgsForCall)]
	fake.getChaincodeDepSpecArgsForCall = append(fake.getChaincodeDepSpecArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetChaincodeDepSpec", []interface{}{arg1})
	fake.getChaincodeDepSpecMutex.Unlock()
	if fake.GetChaincodeDepSpecStub != nil {
		return fake.GetChaincodeDepSpecStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getChaincodeDepSpecReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *LegacyCCPackageLocator) GetChaincodeDepSpecCallCount() int {
	fake.getChaincodeDepSpecMutex.RLock()
	defer fake.getChaincodeDepSpecMutex.RUnlock()
	return len(fake.getChaincodeDepSpecArgsForCall)
}

func (fake *LegacyCCPackageLocator) GetChaincodeDepSpecCalls(stub func(string) (*peer.ChaincodeDeploymentSpec, error)) {
	fake.getChaincodeDepSpecMutex.Lock()
	defer fake.getChaincodeDepSpecMutex.Unlock()
	fake.GetChaincodeDepSpecStub = stub
}

func (fake *LegacyCCPackageLocator) GetChaincodeDepSpecArgsForCall(i int) string {
	fake.getChaincodeDepSpecMutex.RLock()
	defer fake.getChaincodeDepSpecMutex.RUnlock()
	argsForCall := fake.getChaincodeDepSpecArgsForCall[i]
	return argsForCall.arg1
}

func (fake *LegacyCCPackageLocator) GetChaincodeDepSpecReturns(result1 *peer.ChaincodeDeploymentSpec, result2 error) {
	fake.getChaincodeDepSpecMutex.Lock()
	defer fake.getChaincodeDepSpecMutex.Unlock()
	fake.GetChaincodeDepSpecStub = nil
	fake.getChaincodeDepSpecReturns = struct {
		result1 *peer.ChaincodeDeploymentSpec
		result2 error
	}{result1, result2}
}

func (fake *LegacyCCPackageLocator) GetChaincodeDepSpecReturnsOnCall(i int, result1 *peer.ChaincodeDeploymentSpec, result2 error) {
	fake.getChaincodeDepSpecMutex.Lock()
	defer fake.getChaincodeDepSpecMutex.Unlock()
	fake.GetChaincodeDepSpecStub = nil
	if fake.getChaincodeDepSpecReturnsOnCall == nil {
		fake.getChaincodeDepSpecReturnsOnCall = make(map[int]struct {
			result1 *peer.ChaincodeDeploymentSpec
			result2 error
		})
	}
	fake.getChaincodeDepSpecReturnsOnCall[i] = struct {
		result1 *peer.ChaincodeDeploymentSpec
		result2 error
	}{result1, result2}
}

func (fake *LegacyCCPackageLocator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getChaincodeDepSpecMutex.RLock()
	defer fake.getChaincodeDepSpecMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *LegacyCCPackageLocator) recordInvocation(key string, args []interface{}) {
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

var _ persistence.LegacyCCPackageLocator = new(LegacyCCPackageLocator)
