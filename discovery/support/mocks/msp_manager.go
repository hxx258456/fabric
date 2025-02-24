// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	mspa "github.com/hxx258456/fabric-protos-go-cc/msp"
	"github.com/hxx258456/fabric/msp"
)

type MSPManager struct {
	DeserializeIdentityStub        func([]byte) (msp.Identity, error)
	deserializeIdentityMutex       sync.RWMutex
	deserializeIdentityArgsForCall []struct {
		arg1 []byte
	}
	deserializeIdentityReturns struct {
		result1 msp.Identity
		result2 error
	}
	deserializeIdentityReturnsOnCall map[int]struct {
		result1 msp.Identity
		result2 error
	}
	GetMSPsStub        func() (map[string]msp.MSP, error)
	getMSPsMutex       sync.RWMutex
	getMSPsArgsForCall []struct {
	}
	getMSPsReturns struct {
		result1 map[string]msp.MSP
		result2 error
	}
	getMSPsReturnsOnCall map[int]struct {
		result1 map[string]msp.MSP
		result2 error
	}
	IsWellFormedStub        func(*mspa.SerializedIdentity) error
	isWellFormedMutex       sync.RWMutex
	isWellFormedArgsForCall []struct {
		arg1 *mspa.SerializedIdentity
	}
	isWellFormedReturns struct {
		result1 error
	}
	isWellFormedReturnsOnCall map[int]struct {
		result1 error
	}
	SetupStub        func([]msp.MSP) error
	setupMutex       sync.RWMutex
	setupArgsForCall []struct {
		arg1 []msp.MSP
	}
	setupReturns struct {
		result1 error
	}
	setupReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *MSPManager) DeserializeIdentity(arg1 []byte) (msp.Identity, error) {
	var arg1Copy []byte
	if arg1 != nil {
		arg1Copy = make([]byte, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.deserializeIdentityMutex.Lock()
	ret, specificReturn := fake.deserializeIdentityReturnsOnCall[len(fake.deserializeIdentityArgsForCall)]
	fake.deserializeIdentityArgsForCall = append(fake.deserializeIdentityArgsForCall, struct {
		arg1 []byte
	}{arg1Copy})
	fake.recordInvocation("DeserializeIdentity", []interface{}{arg1Copy})
	fake.deserializeIdentityMutex.Unlock()
	if fake.DeserializeIdentityStub != nil {
		return fake.DeserializeIdentityStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.deserializeIdentityReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *MSPManager) DeserializeIdentityCallCount() int {
	fake.deserializeIdentityMutex.RLock()
	defer fake.deserializeIdentityMutex.RUnlock()
	return len(fake.deserializeIdentityArgsForCall)
}

func (fake *MSPManager) DeserializeIdentityCalls(stub func([]byte) (msp.Identity, error)) {
	fake.deserializeIdentityMutex.Lock()
	defer fake.deserializeIdentityMutex.Unlock()
	fake.DeserializeIdentityStub = stub
}

func (fake *MSPManager) DeserializeIdentityArgsForCall(i int) []byte {
	fake.deserializeIdentityMutex.RLock()
	defer fake.deserializeIdentityMutex.RUnlock()
	argsForCall := fake.deserializeIdentityArgsForCall[i]
	return argsForCall.arg1
}

func (fake *MSPManager) DeserializeIdentityReturns(result1 msp.Identity, result2 error) {
	fake.deserializeIdentityMutex.Lock()
	defer fake.deserializeIdentityMutex.Unlock()
	fake.DeserializeIdentityStub = nil
	fake.deserializeIdentityReturns = struct {
		result1 msp.Identity
		result2 error
	}{result1, result2}
}

func (fake *MSPManager) DeserializeIdentityReturnsOnCall(i int, result1 msp.Identity, result2 error) {
	fake.deserializeIdentityMutex.Lock()
	defer fake.deserializeIdentityMutex.Unlock()
	fake.DeserializeIdentityStub = nil
	if fake.deserializeIdentityReturnsOnCall == nil {
		fake.deserializeIdentityReturnsOnCall = make(map[int]struct {
			result1 msp.Identity
			result2 error
		})
	}
	fake.deserializeIdentityReturnsOnCall[i] = struct {
		result1 msp.Identity
		result2 error
	}{result1, result2}
}

func (fake *MSPManager) GetMSPs() (map[string]msp.MSP, error) {
	fake.getMSPsMutex.Lock()
	ret, specificReturn := fake.getMSPsReturnsOnCall[len(fake.getMSPsArgsForCall)]
	fake.getMSPsArgsForCall = append(fake.getMSPsArgsForCall, struct {
	}{})
	fake.recordInvocation("GetMSPs", []interface{}{})
	fake.getMSPsMutex.Unlock()
	if fake.GetMSPsStub != nil {
		return fake.GetMSPsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getMSPsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *MSPManager) GetMSPsCallCount() int {
	fake.getMSPsMutex.RLock()
	defer fake.getMSPsMutex.RUnlock()
	return len(fake.getMSPsArgsForCall)
}

func (fake *MSPManager) GetMSPsCalls(stub func() (map[string]msp.MSP, error)) {
	fake.getMSPsMutex.Lock()
	defer fake.getMSPsMutex.Unlock()
	fake.GetMSPsStub = stub
}

func (fake *MSPManager) GetMSPsReturns(result1 map[string]msp.MSP, result2 error) {
	fake.getMSPsMutex.Lock()
	defer fake.getMSPsMutex.Unlock()
	fake.GetMSPsStub = nil
	fake.getMSPsReturns = struct {
		result1 map[string]msp.MSP
		result2 error
	}{result1, result2}
}

func (fake *MSPManager) GetMSPsReturnsOnCall(i int, result1 map[string]msp.MSP, result2 error) {
	fake.getMSPsMutex.Lock()
	defer fake.getMSPsMutex.Unlock()
	fake.GetMSPsStub = nil
	if fake.getMSPsReturnsOnCall == nil {
		fake.getMSPsReturnsOnCall = make(map[int]struct {
			result1 map[string]msp.MSP
			result2 error
		})
	}
	fake.getMSPsReturnsOnCall[i] = struct {
		result1 map[string]msp.MSP
		result2 error
	}{result1, result2}
}

func (fake *MSPManager) IsWellFormed(arg1 *mspa.SerializedIdentity) error {
	fake.isWellFormedMutex.Lock()
	ret, specificReturn := fake.isWellFormedReturnsOnCall[len(fake.isWellFormedArgsForCall)]
	fake.isWellFormedArgsForCall = append(fake.isWellFormedArgsForCall, struct {
		arg1 *mspa.SerializedIdentity
	}{arg1})
	fake.recordInvocation("IsWellFormed", []interface{}{arg1})
	fake.isWellFormedMutex.Unlock()
	if fake.IsWellFormedStub != nil {
		return fake.IsWellFormedStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.isWellFormedReturns
	return fakeReturns.result1
}

func (fake *MSPManager) IsWellFormedCallCount() int {
	fake.isWellFormedMutex.RLock()
	defer fake.isWellFormedMutex.RUnlock()
	return len(fake.isWellFormedArgsForCall)
}

func (fake *MSPManager) IsWellFormedCalls(stub func(*mspa.SerializedIdentity) error) {
	fake.isWellFormedMutex.Lock()
	defer fake.isWellFormedMutex.Unlock()
	fake.IsWellFormedStub = stub
}

func (fake *MSPManager) IsWellFormedArgsForCall(i int) *mspa.SerializedIdentity {
	fake.isWellFormedMutex.RLock()
	defer fake.isWellFormedMutex.RUnlock()
	argsForCall := fake.isWellFormedArgsForCall[i]
	return argsForCall.arg1
}

func (fake *MSPManager) IsWellFormedReturns(result1 error) {
	fake.isWellFormedMutex.Lock()
	defer fake.isWellFormedMutex.Unlock()
	fake.IsWellFormedStub = nil
	fake.isWellFormedReturns = struct {
		result1 error
	}{result1}
}

func (fake *MSPManager) IsWellFormedReturnsOnCall(i int, result1 error) {
	fake.isWellFormedMutex.Lock()
	defer fake.isWellFormedMutex.Unlock()
	fake.IsWellFormedStub = nil
	if fake.isWellFormedReturnsOnCall == nil {
		fake.isWellFormedReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.isWellFormedReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *MSPManager) Setup(arg1 []msp.MSP) error {
	var arg1Copy []msp.MSP
	if arg1 != nil {
		arg1Copy = make([]msp.MSP, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.setupMutex.Lock()
	ret, specificReturn := fake.setupReturnsOnCall[len(fake.setupArgsForCall)]
	fake.setupArgsForCall = append(fake.setupArgsForCall, struct {
		arg1 []msp.MSP
	}{arg1Copy})
	fake.recordInvocation("Setup", []interface{}{arg1Copy})
	fake.setupMutex.Unlock()
	if fake.SetupStub != nil {
		return fake.SetupStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.setupReturns
	return fakeReturns.result1
}

func (fake *MSPManager) SetupCallCount() int {
	fake.setupMutex.RLock()
	defer fake.setupMutex.RUnlock()
	return len(fake.setupArgsForCall)
}

func (fake *MSPManager) SetupCalls(stub func([]msp.MSP) error) {
	fake.setupMutex.Lock()
	defer fake.setupMutex.Unlock()
	fake.SetupStub = stub
}

func (fake *MSPManager) SetupArgsForCall(i int) []msp.MSP {
	fake.setupMutex.RLock()
	defer fake.setupMutex.RUnlock()
	argsForCall := fake.setupArgsForCall[i]
	return argsForCall.arg1
}

func (fake *MSPManager) SetupReturns(result1 error) {
	fake.setupMutex.Lock()
	defer fake.setupMutex.Unlock()
	fake.SetupStub = nil
	fake.setupReturns = struct {
		result1 error
	}{result1}
}

func (fake *MSPManager) SetupReturnsOnCall(i int, result1 error) {
	fake.setupMutex.Lock()
	defer fake.setupMutex.Unlock()
	fake.SetupStub = nil
	if fake.setupReturnsOnCall == nil {
		fake.setupReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setupReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *MSPManager) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.deserializeIdentityMutex.RLock()
	defer fake.deserializeIdentityMutex.RUnlock()
	fake.getMSPsMutex.RLock()
	defer fake.getMSPsMutex.RUnlock()
	fake.isWellFormedMutex.RLock()
	defer fake.isWellFormedMutex.RUnlock()
	fake.setupMutex.RLock()
	defer fake.setupMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *MSPManager) recordInvocation(key string, args []interface{}) {
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
