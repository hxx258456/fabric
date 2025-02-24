// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	"github.com/hxx258456/fabric-protos-go-cc/discovery"
	"github.com/hxx258456/fabric-protos-go-cc/peer"
	"github.com/hxx258456/fabric/gossip/api"
	"github.com/hxx258456/fabric/gossip/common"
	discoverya "github.com/hxx258456/fabric/gossip/discovery"
)

type Discovery struct {
	ConfigStub        func(string) (*discovery.ConfigResult, error)
	configMutex       sync.RWMutex
	configArgsForCall []struct {
		arg1 string
	}
	configReturns struct {
		result1 *discovery.ConfigResult
		result2 error
	}
	configReturnsOnCall map[int]struct {
		result1 *discovery.ConfigResult
		result2 error
	}
	IdentityInfoStub        func() api.PeerIdentitySet
	identityInfoMutex       sync.RWMutex
	identityInfoArgsForCall []struct {
	}
	identityInfoReturns struct {
		result1 api.PeerIdentitySet
	}
	identityInfoReturnsOnCall map[int]struct {
		result1 api.PeerIdentitySet
	}
	PeersForEndorsementStub        func(common.ChannelID, *peer.ChaincodeInterest) (*discovery.EndorsementDescriptor, error)
	peersForEndorsementMutex       sync.RWMutex
	peersForEndorsementArgsForCall []struct {
		arg1 common.ChannelID
		arg2 *peer.ChaincodeInterest
	}
	peersForEndorsementReturns struct {
		result1 *discovery.EndorsementDescriptor
		result2 error
	}
	peersForEndorsementReturnsOnCall map[int]struct {
		result1 *discovery.EndorsementDescriptor
		result2 error
	}
	PeersOfChannelStub        func(common.ChannelID) discoverya.Members
	peersOfChannelMutex       sync.RWMutex
	peersOfChannelArgsForCall []struct {
		arg1 common.ChannelID
	}
	peersOfChannelReturns struct {
		result1 discoverya.Members
	}
	peersOfChannelReturnsOnCall map[int]struct {
		result1 discoverya.Members
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Discovery) Config(arg1 string) (*discovery.ConfigResult, error) {
	fake.configMutex.Lock()
	ret, specificReturn := fake.configReturnsOnCall[len(fake.configArgsForCall)]
	fake.configArgsForCall = append(fake.configArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.ConfigStub
	fakeReturns := fake.configReturns
	fake.recordInvocation("Config", []interface{}{arg1})
	fake.configMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *Discovery) ConfigCallCount() int {
	fake.configMutex.RLock()
	defer fake.configMutex.RUnlock()
	return len(fake.configArgsForCall)
}

func (fake *Discovery) ConfigCalls(stub func(string) (*discovery.ConfigResult, error)) {
	fake.configMutex.Lock()
	defer fake.configMutex.Unlock()
	fake.ConfigStub = stub
}

func (fake *Discovery) ConfigArgsForCall(i int) string {
	fake.configMutex.RLock()
	defer fake.configMutex.RUnlock()
	argsForCall := fake.configArgsForCall[i]
	return argsForCall.arg1
}

func (fake *Discovery) ConfigReturns(result1 *discovery.ConfigResult, result2 error) {
	fake.configMutex.Lock()
	defer fake.configMutex.Unlock()
	fake.ConfigStub = nil
	fake.configReturns = struct {
		result1 *discovery.ConfigResult
		result2 error
	}{result1, result2}
}

func (fake *Discovery) ConfigReturnsOnCall(i int, result1 *discovery.ConfigResult, result2 error) {
	fake.configMutex.Lock()
	defer fake.configMutex.Unlock()
	fake.ConfigStub = nil
	if fake.configReturnsOnCall == nil {
		fake.configReturnsOnCall = make(map[int]struct {
			result1 *discovery.ConfigResult
			result2 error
		})
	}
	fake.configReturnsOnCall[i] = struct {
		result1 *discovery.ConfigResult
		result2 error
	}{result1, result2}
}

func (fake *Discovery) IdentityInfo() api.PeerIdentitySet {
	fake.identityInfoMutex.Lock()
	ret, specificReturn := fake.identityInfoReturnsOnCall[len(fake.identityInfoArgsForCall)]
	fake.identityInfoArgsForCall = append(fake.identityInfoArgsForCall, struct {
	}{})
	stub := fake.IdentityInfoStub
	fakeReturns := fake.identityInfoReturns
	fake.recordInvocation("IdentityInfo", []interface{}{})
	fake.identityInfoMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *Discovery) IdentityInfoCallCount() int {
	fake.identityInfoMutex.RLock()
	defer fake.identityInfoMutex.RUnlock()
	return len(fake.identityInfoArgsForCall)
}

func (fake *Discovery) IdentityInfoCalls(stub func() api.PeerIdentitySet) {
	fake.identityInfoMutex.Lock()
	defer fake.identityInfoMutex.Unlock()
	fake.IdentityInfoStub = stub
}

func (fake *Discovery) IdentityInfoReturns(result1 api.PeerIdentitySet) {
	fake.identityInfoMutex.Lock()
	defer fake.identityInfoMutex.Unlock()
	fake.IdentityInfoStub = nil
	fake.identityInfoReturns = struct {
		result1 api.PeerIdentitySet
	}{result1}
}

func (fake *Discovery) IdentityInfoReturnsOnCall(i int, result1 api.PeerIdentitySet) {
	fake.identityInfoMutex.Lock()
	defer fake.identityInfoMutex.Unlock()
	fake.IdentityInfoStub = nil
	if fake.identityInfoReturnsOnCall == nil {
		fake.identityInfoReturnsOnCall = make(map[int]struct {
			result1 api.PeerIdentitySet
		})
	}
	fake.identityInfoReturnsOnCall[i] = struct {
		result1 api.PeerIdentitySet
	}{result1}
}

func (fake *Discovery) PeersForEndorsement(arg1 common.ChannelID, arg2 *peer.ChaincodeInterest) (*discovery.EndorsementDescriptor, error) {
	fake.peersForEndorsementMutex.Lock()
	ret, specificReturn := fake.peersForEndorsementReturnsOnCall[len(fake.peersForEndorsementArgsForCall)]
	fake.peersForEndorsementArgsForCall = append(fake.peersForEndorsementArgsForCall, struct {
		arg1 common.ChannelID
		arg2 *peer.ChaincodeInterest
	}{arg1, arg2})
	stub := fake.PeersForEndorsementStub
	fakeReturns := fake.peersForEndorsementReturns
	fake.recordInvocation("PeersForEndorsement", []interface{}{arg1, arg2})
	fake.peersForEndorsementMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *Discovery) PeersForEndorsementCallCount() int {
	fake.peersForEndorsementMutex.RLock()
	defer fake.peersForEndorsementMutex.RUnlock()
	return len(fake.peersForEndorsementArgsForCall)
}

func (fake *Discovery) PeersForEndorsementCalls(stub func(common.ChannelID, *peer.ChaincodeInterest) (*discovery.EndorsementDescriptor, error)) {
	fake.peersForEndorsementMutex.Lock()
	defer fake.peersForEndorsementMutex.Unlock()
	fake.PeersForEndorsementStub = stub
}

func (fake *Discovery) PeersForEndorsementArgsForCall(i int) (common.ChannelID, *peer.ChaincodeInterest) {
	fake.peersForEndorsementMutex.RLock()
	defer fake.peersForEndorsementMutex.RUnlock()
	argsForCall := fake.peersForEndorsementArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *Discovery) PeersForEndorsementReturns(result1 *discovery.EndorsementDescriptor, result2 error) {
	fake.peersForEndorsementMutex.Lock()
	defer fake.peersForEndorsementMutex.Unlock()
	fake.PeersForEndorsementStub = nil
	fake.peersForEndorsementReturns = struct {
		result1 *discovery.EndorsementDescriptor
		result2 error
	}{result1, result2}
}

func (fake *Discovery) PeersForEndorsementReturnsOnCall(i int, result1 *discovery.EndorsementDescriptor, result2 error) {
	fake.peersForEndorsementMutex.Lock()
	defer fake.peersForEndorsementMutex.Unlock()
	fake.PeersForEndorsementStub = nil
	if fake.peersForEndorsementReturnsOnCall == nil {
		fake.peersForEndorsementReturnsOnCall = make(map[int]struct {
			result1 *discovery.EndorsementDescriptor
			result2 error
		})
	}
	fake.peersForEndorsementReturnsOnCall[i] = struct {
		result1 *discovery.EndorsementDescriptor
		result2 error
	}{result1, result2}
}

func (fake *Discovery) PeersOfChannel(arg1 common.ChannelID) discoverya.Members {
	fake.peersOfChannelMutex.Lock()
	ret, specificReturn := fake.peersOfChannelReturnsOnCall[len(fake.peersOfChannelArgsForCall)]
	fake.peersOfChannelArgsForCall = append(fake.peersOfChannelArgsForCall, struct {
		arg1 common.ChannelID
	}{arg1})
	stub := fake.PeersOfChannelStub
	fakeReturns := fake.peersOfChannelReturns
	fake.recordInvocation("PeersOfChannel", []interface{}{arg1})
	fake.peersOfChannelMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *Discovery) PeersOfChannelCallCount() int {
	fake.peersOfChannelMutex.RLock()
	defer fake.peersOfChannelMutex.RUnlock()
	return len(fake.peersOfChannelArgsForCall)
}

func (fake *Discovery) PeersOfChannelCalls(stub func(common.ChannelID) discoverya.Members) {
	fake.peersOfChannelMutex.Lock()
	defer fake.peersOfChannelMutex.Unlock()
	fake.PeersOfChannelStub = stub
}

func (fake *Discovery) PeersOfChannelArgsForCall(i int) common.ChannelID {
	fake.peersOfChannelMutex.RLock()
	defer fake.peersOfChannelMutex.RUnlock()
	argsForCall := fake.peersOfChannelArgsForCall[i]
	return argsForCall.arg1
}

func (fake *Discovery) PeersOfChannelReturns(result1 discoverya.Members) {
	fake.peersOfChannelMutex.Lock()
	defer fake.peersOfChannelMutex.Unlock()
	fake.PeersOfChannelStub = nil
	fake.peersOfChannelReturns = struct {
		result1 discoverya.Members
	}{result1}
}

func (fake *Discovery) PeersOfChannelReturnsOnCall(i int, result1 discoverya.Members) {
	fake.peersOfChannelMutex.Lock()
	defer fake.peersOfChannelMutex.Unlock()
	fake.PeersOfChannelStub = nil
	if fake.peersOfChannelReturnsOnCall == nil {
		fake.peersOfChannelReturnsOnCall = make(map[int]struct {
			result1 discoverya.Members
		})
	}
	fake.peersOfChannelReturnsOnCall[i] = struct {
		result1 discoverya.Members
	}{result1}
}

func (fake *Discovery) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.configMutex.RLock()
	defer fake.configMutex.RUnlock()
	fake.identityInfoMutex.RLock()
	defer fake.identityInfoMutex.RUnlock()
	fake.peersForEndorsementMutex.RLock()
	defer fake.peersForEndorsementMutex.RUnlock()
	fake.peersOfChannelMutex.RLock()
	defer fake.peersOfChannelMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Discovery) recordInvocation(key string, args []interface{}) {
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
