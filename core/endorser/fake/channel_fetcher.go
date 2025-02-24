// Code generated by counterfeiter. DO NOT EDIT.
package fake

import (
	"sync"

	"github.com/hxx258456/fabric/core/endorser"
)

type ChannelFetcher struct {
	ChannelStub        func(string) *endorser.Channel
	channelMutex       sync.RWMutex
	channelArgsForCall []struct {
		arg1 string
	}
	channelReturns struct {
		result1 *endorser.Channel
	}
	channelReturnsOnCall map[int]struct {
		result1 *endorser.Channel
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ChannelFetcher) Channel(arg1 string) *endorser.Channel {
	fake.channelMutex.Lock()
	ret, specificReturn := fake.channelReturnsOnCall[len(fake.channelArgsForCall)]
	fake.channelArgsForCall = append(fake.channelArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Channel", []interface{}{arg1})
	fake.channelMutex.Unlock()
	if fake.ChannelStub != nil {
		return fake.ChannelStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.channelReturns
	return fakeReturns.result1
}

func (fake *ChannelFetcher) ChannelCallCount() int {
	fake.channelMutex.RLock()
	defer fake.channelMutex.RUnlock()
	return len(fake.channelArgsForCall)
}

func (fake *ChannelFetcher) ChannelCalls(stub func(string) *endorser.Channel) {
	fake.channelMutex.Lock()
	defer fake.channelMutex.Unlock()
	fake.ChannelStub = stub
}

func (fake *ChannelFetcher) ChannelArgsForCall(i int) string {
	fake.channelMutex.RLock()
	defer fake.channelMutex.RUnlock()
	argsForCall := fake.channelArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ChannelFetcher) ChannelReturns(result1 *endorser.Channel) {
	fake.channelMutex.Lock()
	defer fake.channelMutex.Unlock()
	fake.ChannelStub = nil
	fake.channelReturns = struct {
		result1 *endorser.Channel
	}{result1}
}

func (fake *ChannelFetcher) ChannelReturnsOnCall(i int, result1 *endorser.Channel) {
	fake.channelMutex.Lock()
	defer fake.channelMutex.Unlock()
	fake.ChannelStub = nil
	if fake.channelReturnsOnCall == nil {
		fake.channelReturnsOnCall = make(map[int]struct {
			result1 *endorser.Channel
		})
	}
	fake.channelReturnsOnCall[i] = struct {
		result1 *endorser.Channel
	}{result1}
}

func (fake *ChannelFetcher) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.channelMutex.RLock()
	defer fake.channelMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *ChannelFetcher) recordInvocation(key string, args []interface{}) {
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

var _ endorser.ChannelFetcher = new(ChannelFetcher)
