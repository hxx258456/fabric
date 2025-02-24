// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	"github.com/hxx258456/fabric/common/channelconfig"
	"github.com/hxx258456/fabric/common/policies"
)

type SigFilterSupport struct {
	OrdererConfigStub        func() (channelconfig.Orderer, bool)
	ordererConfigMutex       sync.RWMutex
	ordererConfigArgsForCall []struct {
	}
	ordererConfigReturns struct {
		result1 channelconfig.Orderer
		result2 bool
	}
	ordererConfigReturnsOnCall map[int]struct {
		result1 channelconfig.Orderer
		result2 bool
	}
	PolicyManagerStub        func() policies.Manager
	policyManagerMutex       sync.RWMutex
	policyManagerArgsForCall []struct {
	}
	policyManagerReturns struct {
		result1 policies.Manager
	}
	policyManagerReturnsOnCall map[int]struct {
		result1 policies.Manager
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *SigFilterSupport) OrdererConfig() (channelconfig.Orderer, bool) {
	fake.ordererConfigMutex.Lock()
	ret, specificReturn := fake.ordererConfigReturnsOnCall[len(fake.ordererConfigArgsForCall)]
	fake.ordererConfigArgsForCall = append(fake.ordererConfigArgsForCall, struct {
	}{})
	fake.recordInvocation("OrdererConfig", []interface{}{})
	fake.ordererConfigMutex.Unlock()
	if fake.OrdererConfigStub != nil {
		return fake.OrdererConfigStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.ordererConfigReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *SigFilterSupport) OrdererConfigCallCount() int {
	fake.ordererConfigMutex.RLock()
	defer fake.ordererConfigMutex.RUnlock()
	return len(fake.ordererConfigArgsForCall)
}

func (fake *SigFilterSupport) OrdererConfigCalls(stub func() (channelconfig.Orderer, bool)) {
	fake.ordererConfigMutex.Lock()
	defer fake.ordererConfigMutex.Unlock()
	fake.OrdererConfigStub = stub
}

func (fake *SigFilterSupport) OrdererConfigReturns(result1 channelconfig.Orderer, result2 bool) {
	fake.ordererConfigMutex.Lock()
	defer fake.ordererConfigMutex.Unlock()
	fake.OrdererConfigStub = nil
	fake.ordererConfigReturns = struct {
		result1 channelconfig.Orderer
		result2 bool
	}{result1, result2}
}

func (fake *SigFilterSupport) OrdererConfigReturnsOnCall(i int, result1 channelconfig.Orderer, result2 bool) {
	fake.ordererConfigMutex.Lock()
	defer fake.ordererConfigMutex.Unlock()
	fake.OrdererConfigStub = nil
	if fake.ordererConfigReturnsOnCall == nil {
		fake.ordererConfigReturnsOnCall = make(map[int]struct {
			result1 channelconfig.Orderer
			result2 bool
		})
	}
	fake.ordererConfigReturnsOnCall[i] = struct {
		result1 channelconfig.Orderer
		result2 bool
	}{result1, result2}
}

func (fake *SigFilterSupport) PolicyManager() policies.Manager {
	fake.policyManagerMutex.Lock()
	ret, specificReturn := fake.policyManagerReturnsOnCall[len(fake.policyManagerArgsForCall)]
	fake.policyManagerArgsForCall = append(fake.policyManagerArgsForCall, struct {
	}{})
	fake.recordInvocation("PolicyManager", []interface{}{})
	fake.policyManagerMutex.Unlock()
	if fake.PolicyManagerStub != nil {
		return fake.PolicyManagerStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.policyManagerReturns
	return fakeReturns.result1
}

func (fake *SigFilterSupport) PolicyManagerCallCount() int {
	fake.policyManagerMutex.RLock()
	defer fake.policyManagerMutex.RUnlock()
	return len(fake.policyManagerArgsForCall)
}

func (fake *SigFilterSupport) PolicyManagerCalls(stub func() policies.Manager) {
	fake.policyManagerMutex.Lock()
	defer fake.policyManagerMutex.Unlock()
	fake.PolicyManagerStub = stub
}

func (fake *SigFilterSupport) PolicyManagerReturns(result1 policies.Manager) {
	fake.policyManagerMutex.Lock()
	defer fake.policyManagerMutex.Unlock()
	fake.PolicyManagerStub = nil
	fake.policyManagerReturns = struct {
		result1 policies.Manager
	}{result1}
}

func (fake *SigFilterSupport) PolicyManagerReturnsOnCall(i int, result1 policies.Manager) {
	fake.policyManagerMutex.Lock()
	defer fake.policyManagerMutex.Unlock()
	fake.PolicyManagerStub = nil
	if fake.policyManagerReturnsOnCall == nil {
		fake.policyManagerReturnsOnCall = make(map[int]struct {
			result1 policies.Manager
		})
	}
	fake.policyManagerReturnsOnCall[i] = struct {
		result1 policies.Manager
	}{result1}
}

func (fake *SigFilterSupport) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.ordererConfigMutex.RLock()
	defer fake.ordererConfigMutex.RUnlock()
	fake.policyManagerMutex.RLock()
	defer fake.policyManagerMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *SigFilterSupport) recordInvocation(key string, args []interface{}) {
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
