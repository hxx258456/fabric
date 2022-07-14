// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"sync"

	"github.com/hxx258456/fabric/core/ledger/kvledger/txmgmt/statedb"
)

type ChannelInfoProvider struct {
	NamespacesAndCollectionsStub        func(statedb.VersionedDB) (map[string][]string, error)
	namespacesAndCollectionsMutex       sync.RWMutex
	namespacesAndCollectionsArgsForCall []struct {
		arg1 statedb.VersionedDB
	}
	namespacesAndCollectionsReturns struct {
		result1 map[string][]string
		result2 error
	}
	namespacesAndCollectionsReturnsOnCall map[int]struct {
		result1 map[string][]string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ChannelInfoProvider) NamespacesAndCollections(arg1 statedb.VersionedDB) (map[string][]string, error) {
	fake.namespacesAndCollectionsMutex.Lock()
	ret, specificReturn := fake.namespacesAndCollectionsReturnsOnCall[len(fake.namespacesAndCollectionsArgsForCall)]
	fake.namespacesAndCollectionsArgsForCall = append(fake.namespacesAndCollectionsArgsForCall, struct {
		arg1 statedb.VersionedDB
	}{arg1})
	fake.recordInvocation("NamespacesAndCollections", []interface{}{arg1})
	fake.namespacesAndCollectionsMutex.Unlock()
	if fake.NamespacesAndCollectionsStub != nil {
		return fake.NamespacesAndCollectionsStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.namespacesAndCollectionsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ChannelInfoProvider) NamespacesAndCollectionsCallCount() int {
	fake.namespacesAndCollectionsMutex.RLock()
	defer fake.namespacesAndCollectionsMutex.RUnlock()
	return len(fake.namespacesAndCollectionsArgsForCall)
}

func (fake *ChannelInfoProvider) NamespacesAndCollectionsCalls(stub func(statedb.VersionedDB) (map[string][]string, error)) {
	fake.namespacesAndCollectionsMutex.Lock()
	defer fake.namespacesAndCollectionsMutex.Unlock()
	fake.NamespacesAndCollectionsStub = stub
}

func (fake *ChannelInfoProvider) NamespacesAndCollectionsArgsForCall(i int) statedb.VersionedDB {
	fake.namespacesAndCollectionsMutex.RLock()
	defer fake.namespacesAndCollectionsMutex.RUnlock()
	argsForCall := fake.namespacesAndCollectionsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ChannelInfoProvider) NamespacesAndCollectionsReturns(result1 map[string][]string, result2 error) {
	fake.namespacesAndCollectionsMutex.Lock()
	defer fake.namespacesAndCollectionsMutex.Unlock()
	fake.NamespacesAndCollectionsStub = nil
	fake.namespacesAndCollectionsReturns = struct {
		result1 map[string][]string
		result2 error
	}{result1, result2}
}

func (fake *ChannelInfoProvider) NamespacesAndCollectionsReturnsOnCall(i int, result1 map[string][]string, result2 error) {
	fake.namespacesAndCollectionsMutex.Lock()
	defer fake.namespacesAndCollectionsMutex.Unlock()
	fake.NamespacesAndCollectionsStub = nil
	if fake.namespacesAndCollectionsReturnsOnCall == nil {
		fake.namespacesAndCollectionsReturnsOnCall = make(map[int]struct {
			result1 map[string][]string
			result2 error
		})
	}
	fake.namespacesAndCollectionsReturnsOnCall[i] = struct {
		result1 map[string][]string
		result2 error
	}{result1, result2}
}

func (fake *ChannelInfoProvider) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.namespacesAndCollectionsMutex.RLock()
	defer fake.namespacesAndCollectionsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *ChannelInfoProvider) recordInvocation(key string, args []interface{}) {
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
