// Code generated by counterfeiter. DO NOT EDIT.
package fake

import (
	"sync"

	"github.com/hxx258456/fabric-protos-go-cc/peer"
	"github.com/hxx258456/fabric/common/ledger"
	"github.com/hxx258456/fabric/core/chaincode"
)

type QueryResponseBuilder struct {
	BuildQueryResponseStub        func(*chaincode.TransactionContext, ledger.ResultsIterator, string, bool, int32) (*peer.QueryResponse, error)
	buildQueryResponseMutex       sync.RWMutex
	buildQueryResponseArgsForCall []struct {
		arg1 *chaincode.TransactionContext
		arg2 ledger.ResultsIterator
		arg3 string
		arg4 bool
		arg5 int32
	}
	buildQueryResponseReturns struct {
		result1 *peer.QueryResponse
		result2 error
	}
	buildQueryResponseReturnsOnCall map[int]struct {
		result1 *peer.QueryResponse
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *QueryResponseBuilder) BuildQueryResponse(arg1 *chaincode.TransactionContext, arg2 ledger.ResultsIterator, arg3 string, arg4 bool, arg5 int32) (*peer.QueryResponse, error) {
	fake.buildQueryResponseMutex.Lock()
	ret, specificReturn := fake.buildQueryResponseReturnsOnCall[len(fake.buildQueryResponseArgsForCall)]
	fake.buildQueryResponseArgsForCall = append(fake.buildQueryResponseArgsForCall, struct {
		arg1 *chaincode.TransactionContext
		arg2 ledger.ResultsIterator
		arg3 string
		arg4 bool
		arg5 int32
	}{arg1, arg2, arg3, arg4, arg5})
	fake.recordInvocation("BuildQueryResponse", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.buildQueryResponseMutex.Unlock()
	if fake.BuildQueryResponseStub != nil {
		return fake.BuildQueryResponseStub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.buildQueryResponseReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *QueryResponseBuilder) BuildQueryResponseCallCount() int {
	fake.buildQueryResponseMutex.RLock()
	defer fake.buildQueryResponseMutex.RUnlock()
	return len(fake.buildQueryResponseArgsForCall)
}

func (fake *QueryResponseBuilder) BuildQueryResponseCalls(stub func(*chaincode.TransactionContext, ledger.ResultsIterator, string, bool, int32) (*peer.QueryResponse, error)) {
	fake.buildQueryResponseMutex.Lock()
	defer fake.buildQueryResponseMutex.Unlock()
	fake.BuildQueryResponseStub = stub
}

func (fake *QueryResponseBuilder) BuildQueryResponseArgsForCall(i int) (*chaincode.TransactionContext, ledger.ResultsIterator, string, bool, int32) {
	fake.buildQueryResponseMutex.RLock()
	defer fake.buildQueryResponseMutex.RUnlock()
	argsForCall := fake.buildQueryResponseArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *QueryResponseBuilder) BuildQueryResponseReturns(result1 *peer.QueryResponse, result2 error) {
	fake.buildQueryResponseMutex.Lock()
	defer fake.buildQueryResponseMutex.Unlock()
	fake.BuildQueryResponseStub = nil
	fake.buildQueryResponseReturns = struct {
		result1 *peer.QueryResponse
		result2 error
	}{result1, result2}
}

func (fake *QueryResponseBuilder) BuildQueryResponseReturnsOnCall(i int, result1 *peer.QueryResponse, result2 error) {
	fake.buildQueryResponseMutex.Lock()
	defer fake.buildQueryResponseMutex.Unlock()
	fake.BuildQueryResponseStub = nil
	if fake.buildQueryResponseReturnsOnCall == nil {
		fake.buildQueryResponseReturnsOnCall = make(map[int]struct {
			result1 *peer.QueryResponse
			result2 error
		})
	}
	fake.buildQueryResponseReturnsOnCall[i] = struct {
		result1 *peer.QueryResponse
		result2 error
	}{result1, result2}
}

func (fake *QueryResponseBuilder) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.buildQueryResponseMutex.RLock()
	defer fake.buildQueryResponseMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *QueryResponseBuilder) recordInvocation(key string, args []interface{}) {
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
