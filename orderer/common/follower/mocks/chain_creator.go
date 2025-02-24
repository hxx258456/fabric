// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	"github.com/hxx258456/fabric/orderer/common/follower"
)

type ChainCreator struct {
	SwitchFollowerToChainStub        func(string)
	switchFollowerToChainMutex       sync.RWMutex
	switchFollowerToChainArgsForCall []struct {
		arg1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ChainCreator) SwitchFollowerToChain(arg1 string) {
	fake.switchFollowerToChainMutex.Lock()
	fake.switchFollowerToChainArgsForCall = append(fake.switchFollowerToChainArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("SwitchFollowerToChain", []interface{}{arg1})
	fake.switchFollowerToChainMutex.Unlock()
	if fake.SwitchFollowerToChainStub != nil {
		fake.SwitchFollowerToChainStub(arg1)
	}
}

func (fake *ChainCreator) SwitchFollowerToChainCallCount() int {
	fake.switchFollowerToChainMutex.RLock()
	defer fake.switchFollowerToChainMutex.RUnlock()
	return len(fake.switchFollowerToChainArgsForCall)
}

func (fake *ChainCreator) SwitchFollowerToChainCalls(stub func(string)) {
	fake.switchFollowerToChainMutex.Lock()
	defer fake.switchFollowerToChainMutex.Unlock()
	fake.SwitchFollowerToChainStub = stub
}

func (fake *ChainCreator) SwitchFollowerToChainArgsForCall(i int) string {
	fake.switchFollowerToChainMutex.RLock()
	defer fake.switchFollowerToChainMutex.RUnlock()
	argsForCall := fake.switchFollowerToChainArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ChainCreator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.switchFollowerToChainMutex.RLock()
	defer fake.switchFollowerToChainMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *ChainCreator) recordInvocation(key string, args []interface{}) {
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

var _ follower.ChainCreator = new(ChainCreator)
