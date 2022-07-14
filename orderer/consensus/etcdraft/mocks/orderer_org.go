// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	"github.com/hxx258456/fabric/msp"
)

type OrdererOrg struct {
	EndpointsStub        func() []string
	endpointsMutex       sync.RWMutex
	endpointsArgsForCall []struct {
	}
	endpointsReturns struct {
		result1 []string
	}
	endpointsReturnsOnCall map[int]struct {
		result1 []string
	}
	MSPStub        func() msp.MSP
	mSPMutex       sync.RWMutex
	mSPArgsForCall []struct {
	}
	mSPReturns struct {
		result1 msp.MSP
	}
	mSPReturnsOnCall map[int]struct {
		result1 msp.MSP
	}
	MSPIDStub        func() string
	mSPIDMutex       sync.RWMutex
	mSPIDArgsForCall []struct {
	}
	mSPIDReturns struct {
		result1 string
	}
	mSPIDReturnsOnCall map[int]struct {
		result1 string
	}
	NameStub        func() string
	nameMutex       sync.RWMutex
	nameArgsForCall []struct {
	}
	nameReturns struct {
		result1 string
	}
	nameReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *OrdererOrg) Endpoints() []string {
	fake.endpointsMutex.Lock()
	ret, specificReturn := fake.endpointsReturnsOnCall[len(fake.endpointsArgsForCall)]
	fake.endpointsArgsForCall = append(fake.endpointsArgsForCall, struct {
	}{})
	fake.recordInvocation("Endpoints", []interface{}{})
	fake.endpointsMutex.Unlock()
	if fake.EndpointsStub != nil {
		return fake.EndpointsStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.endpointsReturns
	return fakeReturns.result1
}

func (fake *OrdererOrg) EndpointsCallCount() int {
	fake.endpointsMutex.RLock()
	defer fake.endpointsMutex.RUnlock()
	return len(fake.endpointsArgsForCall)
}

func (fake *OrdererOrg) EndpointsCalls(stub func() []string) {
	fake.endpointsMutex.Lock()
	defer fake.endpointsMutex.Unlock()
	fake.EndpointsStub = stub
}

func (fake *OrdererOrg) EndpointsReturns(result1 []string) {
	fake.endpointsMutex.Lock()
	defer fake.endpointsMutex.Unlock()
	fake.EndpointsStub = nil
	fake.endpointsReturns = struct {
		result1 []string
	}{result1}
}

func (fake *OrdererOrg) EndpointsReturnsOnCall(i int, result1 []string) {
	fake.endpointsMutex.Lock()
	defer fake.endpointsMutex.Unlock()
	fake.EndpointsStub = nil
	if fake.endpointsReturnsOnCall == nil {
		fake.endpointsReturnsOnCall = make(map[int]struct {
			result1 []string
		})
	}
	fake.endpointsReturnsOnCall[i] = struct {
		result1 []string
	}{result1}
}

func (fake *OrdererOrg) MSP() msp.MSP {
	fake.mSPMutex.Lock()
	ret, specificReturn := fake.mSPReturnsOnCall[len(fake.mSPArgsForCall)]
	fake.mSPArgsForCall = append(fake.mSPArgsForCall, struct {
	}{})
	fake.recordInvocation("MSP", []interface{}{})
	fake.mSPMutex.Unlock()
	if fake.MSPStub != nil {
		return fake.MSPStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.mSPReturns
	return fakeReturns.result1
}

func (fake *OrdererOrg) MSPCallCount() int {
	fake.mSPMutex.RLock()
	defer fake.mSPMutex.RUnlock()
	return len(fake.mSPArgsForCall)
}

func (fake *OrdererOrg) MSPCalls(stub func() msp.MSP) {
	fake.mSPMutex.Lock()
	defer fake.mSPMutex.Unlock()
	fake.MSPStub = stub
}

func (fake *OrdererOrg) MSPReturns(result1 msp.MSP) {
	fake.mSPMutex.Lock()
	defer fake.mSPMutex.Unlock()
	fake.MSPStub = nil
	fake.mSPReturns = struct {
		result1 msp.MSP
	}{result1}
}

func (fake *OrdererOrg) MSPReturnsOnCall(i int, result1 msp.MSP) {
	fake.mSPMutex.Lock()
	defer fake.mSPMutex.Unlock()
	fake.MSPStub = nil
	if fake.mSPReturnsOnCall == nil {
		fake.mSPReturnsOnCall = make(map[int]struct {
			result1 msp.MSP
		})
	}
	fake.mSPReturnsOnCall[i] = struct {
		result1 msp.MSP
	}{result1}
}

func (fake *OrdererOrg) MSPID() string {
	fake.mSPIDMutex.Lock()
	ret, specificReturn := fake.mSPIDReturnsOnCall[len(fake.mSPIDArgsForCall)]
	fake.mSPIDArgsForCall = append(fake.mSPIDArgsForCall, struct {
	}{})
	fake.recordInvocation("MSPID", []interface{}{})
	fake.mSPIDMutex.Unlock()
	if fake.MSPIDStub != nil {
		return fake.MSPIDStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.mSPIDReturns
	return fakeReturns.result1
}

func (fake *OrdererOrg) MSPIDCallCount() int {
	fake.mSPIDMutex.RLock()
	defer fake.mSPIDMutex.RUnlock()
	return len(fake.mSPIDArgsForCall)
}

func (fake *OrdererOrg) MSPIDCalls(stub func() string) {
	fake.mSPIDMutex.Lock()
	defer fake.mSPIDMutex.Unlock()
	fake.MSPIDStub = stub
}

func (fake *OrdererOrg) MSPIDReturns(result1 string) {
	fake.mSPIDMutex.Lock()
	defer fake.mSPIDMutex.Unlock()
	fake.MSPIDStub = nil
	fake.mSPIDReturns = struct {
		result1 string
	}{result1}
}

func (fake *OrdererOrg) MSPIDReturnsOnCall(i int, result1 string) {
	fake.mSPIDMutex.Lock()
	defer fake.mSPIDMutex.Unlock()
	fake.MSPIDStub = nil
	if fake.mSPIDReturnsOnCall == nil {
		fake.mSPIDReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.mSPIDReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *OrdererOrg) Name() string {
	fake.nameMutex.Lock()
	ret, specificReturn := fake.nameReturnsOnCall[len(fake.nameArgsForCall)]
	fake.nameArgsForCall = append(fake.nameArgsForCall, struct {
	}{})
	fake.recordInvocation("Name", []interface{}{})
	fake.nameMutex.Unlock()
	if fake.NameStub != nil {
		return fake.NameStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.nameReturns
	return fakeReturns.result1
}

func (fake *OrdererOrg) NameCallCount() int {
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	return len(fake.nameArgsForCall)
}

func (fake *OrdererOrg) NameCalls(stub func() string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = stub
}

func (fake *OrdererOrg) NameReturns(result1 string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = nil
	fake.nameReturns = struct {
		result1 string
	}{result1}
}

func (fake *OrdererOrg) NameReturnsOnCall(i int, result1 string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = nil
	if fake.nameReturnsOnCall == nil {
		fake.nameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.nameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *OrdererOrg) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.endpointsMutex.RLock()
	defer fake.endpointsMutex.RUnlock()
	fake.mSPMutex.RLock()
	defer fake.mSPMutex.RUnlock()
	fake.mSPIDMutex.RLock()
	defer fake.mSPIDMutex.RUnlock()
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *OrdererOrg) recordInvocation(key string, args []interface{}) {
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
