// Code generated by counterfeiter. DO NOT EDIT.
package fake

import (
	"sync"

	"github.com/hxx258456/fabric/core/chaincode"
)

type Registry struct {
	DeregisterStub        func(string) error
	deregisterMutex       sync.RWMutex
	deregisterArgsForCall []struct {
		arg1 string
	}
	deregisterReturns struct {
		result1 error
	}
	deregisterReturnsOnCall map[int]struct {
		result1 error
	}
	FailedStub        func(string, error)
	failedMutex       sync.RWMutex
	failedArgsForCall []struct {
		arg1 string
		arg2 error
	}
	ReadyStub        func(string)
	readyMutex       sync.RWMutex
	readyArgsForCall []struct {
		arg1 string
	}
	RegisterStub        func(*chaincode.Handler) error
	registerMutex       sync.RWMutex
	registerArgsForCall []struct {
		arg1 *chaincode.Handler
	}
	registerReturns struct {
		result1 error
	}
	registerReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Registry) Deregister(arg1 string) error {
	fake.deregisterMutex.Lock()
	ret, specificReturn := fake.deregisterReturnsOnCall[len(fake.deregisterArgsForCall)]
	fake.deregisterArgsForCall = append(fake.deregisterArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Deregister", []interface{}{arg1})
	fake.deregisterMutex.Unlock()
	if fake.DeregisterStub != nil {
		return fake.DeregisterStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.deregisterReturns
	return fakeReturns.result1
}

func (fake *Registry) DeregisterCallCount() int {
	fake.deregisterMutex.RLock()
	defer fake.deregisterMutex.RUnlock()
	return len(fake.deregisterArgsForCall)
}

func (fake *Registry) DeregisterCalls(stub func(string) error) {
	fake.deregisterMutex.Lock()
	defer fake.deregisterMutex.Unlock()
	fake.DeregisterStub = stub
}

func (fake *Registry) DeregisterArgsForCall(i int) string {
	fake.deregisterMutex.RLock()
	defer fake.deregisterMutex.RUnlock()
	argsForCall := fake.deregisterArgsForCall[i]
	return argsForCall.arg1
}

func (fake *Registry) DeregisterReturns(result1 error) {
	fake.deregisterMutex.Lock()
	defer fake.deregisterMutex.Unlock()
	fake.DeregisterStub = nil
	fake.deregisterReturns = struct {
		result1 error
	}{result1}
}

func (fake *Registry) DeregisterReturnsOnCall(i int, result1 error) {
	fake.deregisterMutex.Lock()
	defer fake.deregisterMutex.Unlock()
	fake.DeregisterStub = nil
	if fake.deregisterReturnsOnCall == nil {
		fake.deregisterReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deregisterReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Registry) Failed(arg1 string, arg2 error) {
	fake.failedMutex.Lock()
	fake.failedArgsForCall = append(fake.failedArgsForCall, struct {
		arg1 string
		arg2 error
	}{arg1, arg2})
	fake.recordInvocation("Failed", []interface{}{arg1, arg2})
	fake.failedMutex.Unlock()
	if fake.FailedStub != nil {
		fake.FailedStub(arg1, arg2)
	}
}

func (fake *Registry) FailedCallCount() int {
	fake.failedMutex.RLock()
	defer fake.failedMutex.RUnlock()
	return len(fake.failedArgsForCall)
}

func (fake *Registry) FailedCalls(stub func(string, error)) {
	fake.failedMutex.Lock()
	defer fake.failedMutex.Unlock()
	fake.FailedStub = stub
}

func (fake *Registry) FailedArgsForCall(i int) (string, error) {
	fake.failedMutex.RLock()
	defer fake.failedMutex.RUnlock()
	argsForCall := fake.failedArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *Registry) Ready(arg1 string) {
	fake.readyMutex.Lock()
	fake.readyArgsForCall = append(fake.readyArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Ready", []interface{}{arg1})
	fake.readyMutex.Unlock()
	if fake.ReadyStub != nil {
		fake.ReadyStub(arg1)
	}
}

func (fake *Registry) ReadyCallCount() int {
	fake.readyMutex.RLock()
	defer fake.readyMutex.RUnlock()
	return len(fake.readyArgsForCall)
}

func (fake *Registry) ReadyCalls(stub func(string)) {
	fake.readyMutex.Lock()
	defer fake.readyMutex.Unlock()
	fake.ReadyStub = stub
}

func (fake *Registry) ReadyArgsForCall(i int) string {
	fake.readyMutex.RLock()
	defer fake.readyMutex.RUnlock()
	argsForCall := fake.readyArgsForCall[i]
	return argsForCall.arg1
}

func (fake *Registry) Register(arg1 *chaincode.Handler) error {
	fake.registerMutex.Lock()
	ret, specificReturn := fake.registerReturnsOnCall[len(fake.registerArgsForCall)]
	fake.registerArgsForCall = append(fake.registerArgsForCall, struct {
		arg1 *chaincode.Handler
	}{arg1})
	fake.recordInvocation("Register", []interface{}{arg1})
	fake.registerMutex.Unlock()
	if fake.RegisterStub != nil {
		return fake.RegisterStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.registerReturns
	return fakeReturns.result1
}

func (fake *Registry) RegisterCallCount() int {
	fake.registerMutex.RLock()
	defer fake.registerMutex.RUnlock()
	return len(fake.registerArgsForCall)
}

func (fake *Registry) RegisterCalls(stub func(*chaincode.Handler) error) {
	fake.registerMutex.Lock()
	defer fake.registerMutex.Unlock()
	fake.RegisterStub = stub
}

func (fake *Registry) RegisterArgsForCall(i int) *chaincode.Handler {
	fake.registerMutex.RLock()
	defer fake.registerMutex.RUnlock()
	argsForCall := fake.registerArgsForCall[i]
	return argsForCall.arg1
}

func (fake *Registry) RegisterReturns(result1 error) {
	fake.registerMutex.Lock()
	defer fake.registerMutex.Unlock()
	fake.RegisterStub = nil
	fake.registerReturns = struct {
		result1 error
	}{result1}
}

func (fake *Registry) RegisterReturnsOnCall(i int, result1 error) {
	fake.registerMutex.Lock()
	defer fake.registerMutex.Unlock()
	fake.RegisterStub = nil
	if fake.registerReturnsOnCall == nil {
		fake.registerReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.registerReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Registry) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.deregisterMutex.RLock()
	defer fake.deregisterMutex.RUnlock()
	fake.failedMutex.RLock()
	defer fake.failedMutex.RUnlock()
	fake.readyMutex.RLock()
	defer fake.readyMutex.RUnlock()
	fake.registerMutex.RLock()
	defer fake.registerMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Registry) recordInvocation(key string, args []interface{}) {
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
