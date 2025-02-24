// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"sync"

	"github.com/hxx258456/fabric/core/ledger"
)

type StateListener struct {
	HandleStateUpdatesStub        func(*ledger.StateUpdateTrigger) error
	handleStateUpdatesMutex       sync.RWMutex
	handleStateUpdatesArgsForCall []struct {
		arg1 *ledger.StateUpdateTrigger
	}
	handleStateUpdatesReturns struct {
		result1 error
	}
	handleStateUpdatesReturnsOnCall map[int]struct {
		result1 error
	}
	InitializeStub        func(string, ledger.SimpleQueryExecutor) error
	initializeMutex       sync.RWMutex
	initializeArgsForCall []struct {
		arg1 string
		arg2 ledger.SimpleQueryExecutor
	}
	initializeReturns struct {
		result1 error
	}
	initializeReturnsOnCall map[int]struct {
		result1 error
	}
	InterestedInNamespacesStub        func() []string
	interestedInNamespacesMutex       sync.RWMutex
	interestedInNamespacesArgsForCall []struct {
	}
	interestedInNamespacesReturns struct {
		result1 []string
	}
	interestedInNamespacesReturnsOnCall map[int]struct {
		result1 []string
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
	StateCommitDoneStub        func(string)
	stateCommitDoneMutex       sync.RWMutex
	stateCommitDoneArgsForCall []struct {
		arg1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *StateListener) HandleStateUpdates(arg1 *ledger.StateUpdateTrigger) error {
	fake.handleStateUpdatesMutex.Lock()
	ret, specificReturn := fake.handleStateUpdatesReturnsOnCall[len(fake.handleStateUpdatesArgsForCall)]
	fake.handleStateUpdatesArgsForCall = append(fake.handleStateUpdatesArgsForCall, struct {
		arg1 *ledger.StateUpdateTrigger
	}{arg1})
	fake.recordInvocation("HandleStateUpdates", []interface{}{arg1})
	fake.handleStateUpdatesMutex.Unlock()
	if fake.HandleStateUpdatesStub != nil {
		return fake.HandleStateUpdatesStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.handleStateUpdatesReturns
	return fakeReturns.result1
}

func (fake *StateListener) HandleStateUpdatesCallCount() int {
	fake.handleStateUpdatesMutex.RLock()
	defer fake.handleStateUpdatesMutex.RUnlock()
	return len(fake.handleStateUpdatesArgsForCall)
}

func (fake *StateListener) HandleStateUpdatesCalls(stub func(*ledger.StateUpdateTrigger) error) {
	fake.handleStateUpdatesMutex.Lock()
	defer fake.handleStateUpdatesMutex.Unlock()
	fake.HandleStateUpdatesStub = stub
}

func (fake *StateListener) HandleStateUpdatesArgsForCall(i int) *ledger.StateUpdateTrigger {
	fake.handleStateUpdatesMutex.RLock()
	defer fake.handleStateUpdatesMutex.RUnlock()
	argsForCall := fake.handleStateUpdatesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *StateListener) HandleStateUpdatesReturns(result1 error) {
	fake.handleStateUpdatesMutex.Lock()
	defer fake.handleStateUpdatesMutex.Unlock()
	fake.HandleStateUpdatesStub = nil
	fake.handleStateUpdatesReturns = struct {
		result1 error
	}{result1}
}

func (fake *StateListener) HandleStateUpdatesReturnsOnCall(i int, result1 error) {
	fake.handleStateUpdatesMutex.Lock()
	defer fake.handleStateUpdatesMutex.Unlock()
	fake.HandleStateUpdatesStub = nil
	if fake.handleStateUpdatesReturnsOnCall == nil {
		fake.handleStateUpdatesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.handleStateUpdatesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *StateListener) Initialize(arg1 string, arg2 ledger.SimpleQueryExecutor) error {
	fake.initializeMutex.Lock()
	ret, specificReturn := fake.initializeReturnsOnCall[len(fake.initializeArgsForCall)]
	fake.initializeArgsForCall = append(fake.initializeArgsForCall, struct {
		arg1 string
		arg2 ledger.SimpleQueryExecutor
	}{arg1, arg2})
	fake.recordInvocation("Initialize", []interface{}{arg1, arg2})
	fake.initializeMutex.Unlock()
	if fake.InitializeStub != nil {
		return fake.InitializeStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.initializeReturns
	return fakeReturns.result1
}

func (fake *StateListener) InitializeCallCount() int {
	fake.initializeMutex.RLock()
	defer fake.initializeMutex.RUnlock()
	return len(fake.initializeArgsForCall)
}

func (fake *StateListener) InitializeCalls(stub func(string, ledger.SimpleQueryExecutor) error) {
	fake.initializeMutex.Lock()
	defer fake.initializeMutex.Unlock()
	fake.InitializeStub = stub
}

func (fake *StateListener) InitializeArgsForCall(i int) (string, ledger.SimpleQueryExecutor) {
	fake.initializeMutex.RLock()
	defer fake.initializeMutex.RUnlock()
	argsForCall := fake.initializeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *StateListener) InitializeReturns(result1 error) {
	fake.initializeMutex.Lock()
	defer fake.initializeMutex.Unlock()
	fake.InitializeStub = nil
	fake.initializeReturns = struct {
		result1 error
	}{result1}
}

func (fake *StateListener) InitializeReturnsOnCall(i int, result1 error) {
	fake.initializeMutex.Lock()
	defer fake.initializeMutex.Unlock()
	fake.InitializeStub = nil
	if fake.initializeReturnsOnCall == nil {
		fake.initializeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.initializeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *StateListener) InterestedInNamespaces() []string {
	fake.interestedInNamespacesMutex.Lock()
	ret, specificReturn := fake.interestedInNamespacesReturnsOnCall[len(fake.interestedInNamespacesArgsForCall)]
	fake.interestedInNamespacesArgsForCall = append(fake.interestedInNamespacesArgsForCall, struct {
	}{})
	fake.recordInvocation("InterestedInNamespaces", []interface{}{})
	fake.interestedInNamespacesMutex.Unlock()
	if fake.InterestedInNamespacesStub != nil {
		return fake.InterestedInNamespacesStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.interestedInNamespacesReturns
	return fakeReturns.result1
}

func (fake *StateListener) InterestedInNamespacesCallCount() int {
	fake.interestedInNamespacesMutex.RLock()
	defer fake.interestedInNamespacesMutex.RUnlock()
	return len(fake.interestedInNamespacesArgsForCall)
}

func (fake *StateListener) InterestedInNamespacesCalls(stub func() []string) {
	fake.interestedInNamespacesMutex.Lock()
	defer fake.interestedInNamespacesMutex.Unlock()
	fake.InterestedInNamespacesStub = stub
}

func (fake *StateListener) InterestedInNamespacesReturns(result1 []string) {
	fake.interestedInNamespacesMutex.Lock()
	defer fake.interestedInNamespacesMutex.Unlock()
	fake.InterestedInNamespacesStub = nil
	fake.interestedInNamespacesReturns = struct {
		result1 []string
	}{result1}
}

func (fake *StateListener) InterestedInNamespacesReturnsOnCall(i int, result1 []string) {
	fake.interestedInNamespacesMutex.Lock()
	defer fake.interestedInNamespacesMutex.Unlock()
	fake.InterestedInNamespacesStub = nil
	if fake.interestedInNamespacesReturnsOnCall == nil {
		fake.interestedInNamespacesReturnsOnCall = make(map[int]struct {
			result1 []string
		})
	}
	fake.interestedInNamespacesReturnsOnCall[i] = struct {
		result1 []string
	}{result1}
}

func (fake *StateListener) Name() string {
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

func (fake *StateListener) NameCallCount() int {
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	return len(fake.nameArgsForCall)
}

func (fake *StateListener) NameCalls(stub func() string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = stub
}

func (fake *StateListener) NameReturns(result1 string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = nil
	fake.nameReturns = struct {
		result1 string
	}{result1}
}

func (fake *StateListener) NameReturnsOnCall(i int, result1 string) {
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

func (fake *StateListener) StateCommitDone(arg1 string) {
	fake.stateCommitDoneMutex.Lock()
	fake.stateCommitDoneArgsForCall = append(fake.stateCommitDoneArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("StateCommitDone", []interface{}{arg1})
	fake.stateCommitDoneMutex.Unlock()
	if fake.StateCommitDoneStub != nil {
		fake.StateCommitDoneStub(arg1)
	}
}

func (fake *StateListener) StateCommitDoneCallCount() int {
	fake.stateCommitDoneMutex.RLock()
	defer fake.stateCommitDoneMutex.RUnlock()
	return len(fake.stateCommitDoneArgsForCall)
}

func (fake *StateListener) StateCommitDoneCalls(stub func(string)) {
	fake.stateCommitDoneMutex.Lock()
	defer fake.stateCommitDoneMutex.Unlock()
	fake.StateCommitDoneStub = stub
}

func (fake *StateListener) StateCommitDoneArgsForCall(i int) string {
	fake.stateCommitDoneMutex.RLock()
	defer fake.stateCommitDoneMutex.RUnlock()
	argsForCall := fake.stateCommitDoneArgsForCall[i]
	return argsForCall.arg1
}

func (fake *StateListener) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.handleStateUpdatesMutex.RLock()
	defer fake.handleStateUpdatesMutex.RUnlock()
	fake.initializeMutex.RLock()
	defer fake.initializeMutex.RUnlock()
	fake.interestedInNamespacesMutex.RLock()
	defer fake.interestedInNamespacesMutex.RUnlock()
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	fake.stateCommitDoneMutex.RLock()
	defer fake.stateCommitDoneMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *StateListener) recordInvocation(key string, args []interface{}) {
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

var _ ledger.StateListener = new(StateListener)
