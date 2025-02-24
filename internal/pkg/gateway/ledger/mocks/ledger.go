// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	"github.com/hxx258456/fabric-protos-go-cc/common"
	"github.com/hxx258456/fabric-protos-go-cc/peer"
	ledgerb "github.com/hxx258456/fabric/common/ledger"
	ledgera "github.com/hxx258456/fabric/core/ledger"
	"github.com/hxx258456/fabric/internal/pkg/gateway/ledger"
)

type Ledger struct {
	CommitNotificationsChannelStub        func(<-chan struct{}) (<-chan *ledgera.CommitNotification, error)
	commitNotificationsChannelMutex       sync.RWMutex
	commitNotificationsChannelArgsForCall []struct {
		arg1 <-chan struct{}
	}
	commitNotificationsChannelReturns struct {
		result1 <-chan *ledgera.CommitNotification
		result2 error
	}
	commitNotificationsChannelReturnsOnCall map[int]struct {
		result1 <-chan *ledgera.CommitNotification
		result2 error
	}
	GetBlockByTxIDStub        func(string) (*common.Block, error)
	getBlockByTxIDMutex       sync.RWMutex
	getBlockByTxIDArgsForCall []struct {
		arg1 string
	}
	getBlockByTxIDReturns struct {
		result1 *common.Block
		result2 error
	}
	getBlockByTxIDReturnsOnCall map[int]struct {
		result1 *common.Block
		result2 error
	}
	GetBlockchainInfoStub        func() (*common.BlockchainInfo, error)
	getBlockchainInfoMutex       sync.RWMutex
	getBlockchainInfoArgsForCall []struct {
	}
	getBlockchainInfoReturns struct {
		result1 *common.BlockchainInfo
		result2 error
	}
	getBlockchainInfoReturnsOnCall map[int]struct {
		result1 *common.BlockchainInfo
		result2 error
	}
	GetBlocksIteratorStub        func(uint64) (ledgerb.ResultsIterator, error)
	getBlocksIteratorMutex       sync.RWMutex
	getBlocksIteratorArgsForCall []struct {
		arg1 uint64
	}
	getBlocksIteratorReturns struct {
		result1 ledgerb.ResultsIterator
		result2 error
	}
	getBlocksIteratorReturnsOnCall map[int]struct {
		result1 ledgerb.ResultsIterator
		result2 error
	}
	GetTxValidationCodeByTxIDStub        func(string) (peer.TxValidationCode, uint64, error)
	getTxValidationCodeByTxIDMutex       sync.RWMutex
	getTxValidationCodeByTxIDArgsForCall []struct {
		arg1 string
	}
	getTxValidationCodeByTxIDReturns struct {
		result1 peer.TxValidationCode
		result2 uint64
		result3 error
	}
	getTxValidationCodeByTxIDReturnsOnCall map[int]struct {
		result1 peer.TxValidationCode
		result2 uint64
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Ledger) CommitNotificationsChannel(arg1 <-chan struct{}) (<-chan *ledgera.CommitNotification, error) {
	fake.commitNotificationsChannelMutex.Lock()
	ret, specificReturn := fake.commitNotificationsChannelReturnsOnCall[len(fake.commitNotificationsChannelArgsForCall)]
	fake.commitNotificationsChannelArgsForCall = append(fake.commitNotificationsChannelArgsForCall, struct {
		arg1 <-chan struct{}
	}{arg1})
	stub := fake.CommitNotificationsChannelStub
	fakeReturns := fake.commitNotificationsChannelReturns
	fake.recordInvocation("CommitNotificationsChannel", []interface{}{arg1})
	fake.commitNotificationsChannelMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *Ledger) CommitNotificationsChannelCallCount() int {
	fake.commitNotificationsChannelMutex.RLock()
	defer fake.commitNotificationsChannelMutex.RUnlock()
	return len(fake.commitNotificationsChannelArgsForCall)
}

func (fake *Ledger) CommitNotificationsChannelCalls(stub func(<-chan struct{}) (<-chan *ledgera.CommitNotification, error)) {
	fake.commitNotificationsChannelMutex.Lock()
	defer fake.commitNotificationsChannelMutex.Unlock()
	fake.CommitNotificationsChannelStub = stub
}

func (fake *Ledger) CommitNotificationsChannelArgsForCall(i int) <-chan struct{} {
	fake.commitNotificationsChannelMutex.RLock()
	defer fake.commitNotificationsChannelMutex.RUnlock()
	argsForCall := fake.commitNotificationsChannelArgsForCall[i]
	return argsForCall.arg1
}

func (fake *Ledger) CommitNotificationsChannelReturns(result1 <-chan *ledgera.CommitNotification, result2 error) {
	fake.commitNotificationsChannelMutex.Lock()
	defer fake.commitNotificationsChannelMutex.Unlock()
	fake.CommitNotificationsChannelStub = nil
	fake.commitNotificationsChannelReturns = struct {
		result1 <-chan *ledgera.CommitNotification
		result2 error
	}{result1, result2}
}

func (fake *Ledger) CommitNotificationsChannelReturnsOnCall(i int, result1 <-chan *ledgera.CommitNotification, result2 error) {
	fake.commitNotificationsChannelMutex.Lock()
	defer fake.commitNotificationsChannelMutex.Unlock()
	fake.CommitNotificationsChannelStub = nil
	if fake.commitNotificationsChannelReturnsOnCall == nil {
		fake.commitNotificationsChannelReturnsOnCall = make(map[int]struct {
			result1 <-chan *ledgera.CommitNotification
			result2 error
		})
	}
	fake.commitNotificationsChannelReturnsOnCall[i] = struct {
		result1 <-chan *ledgera.CommitNotification
		result2 error
	}{result1, result2}
}

func (fake *Ledger) GetBlockByTxID(arg1 string) (*common.Block, error) {
	fake.getBlockByTxIDMutex.Lock()
	ret, specificReturn := fake.getBlockByTxIDReturnsOnCall[len(fake.getBlockByTxIDArgsForCall)]
	fake.getBlockByTxIDArgsForCall = append(fake.getBlockByTxIDArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetBlockByTxIDStub
	fakeReturns := fake.getBlockByTxIDReturns
	fake.recordInvocation("GetBlockByTxID", []interface{}{arg1})
	fake.getBlockByTxIDMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *Ledger) GetBlockByTxIDCallCount() int {
	fake.getBlockByTxIDMutex.RLock()
	defer fake.getBlockByTxIDMutex.RUnlock()
	return len(fake.getBlockByTxIDArgsForCall)
}

func (fake *Ledger) GetBlockByTxIDCalls(stub func(string) (*common.Block, error)) {
	fake.getBlockByTxIDMutex.Lock()
	defer fake.getBlockByTxIDMutex.Unlock()
	fake.GetBlockByTxIDStub = stub
}

func (fake *Ledger) GetBlockByTxIDArgsForCall(i int) string {
	fake.getBlockByTxIDMutex.RLock()
	defer fake.getBlockByTxIDMutex.RUnlock()
	argsForCall := fake.getBlockByTxIDArgsForCall[i]
	return argsForCall.arg1
}

func (fake *Ledger) GetBlockByTxIDReturns(result1 *common.Block, result2 error) {
	fake.getBlockByTxIDMutex.Lock()
	defer fake.getBlockByTxIDMutex.Unlock()
	fake.GetBlockByTxIDStub = nil
	fake.getBlockByTxIDReturns = struct {
		result1 *common.Block
		result2 error
	}{result1, result2}
}

func (fake *Ledger) GetBlockByTxIDReturnsOnCall(i int, result1 *common.Block, result2 error) {
	fake.getBlockByTxIDMutex.Lock()
	defer fake.getBlockByTxIDMutex.Unlock()
	fake.GetBlockByTxIDStub = nil
	if fake.getBlockByTxIDReturnsOnCall == nil {
		fake.getBlockByTxIDReturnsOnCall = make(map[int]struct {
			result1 *common.Block
			result2 error
		})
	}
	fake.getBlockByTxIDReturnsOnCall[i] = struct {
		result1 *common.Block
		result2 error
	}{result1, result2}
}

func (fake *Ledger) GetBlockchainInfo() (*common.BlockchainInfo, error) {
	fake.getBlockchainInfoMutex.Lock()
	ret, specificReturn := fake.getBlockchainInfoReturnsOnCall[len(fake.getBlockchainInfoArgsForCall)]
	fake.getBlockchainInfoArgsForCall = append(fake.getBlockchainInfoArgsForCall, struct {
	}{})
	stub := fake.GetBlockchainInfoStub
	fakeReturns := fake.getBlockchainInfoReturns
	fake.recordInvocation("GetBlockchainInfo", []interface{}{})
	fake.getBlockchainInfoMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *Ledger) GetBlockchainInfoCallCount() int {
	fake.getBlockchainInfoMutex.RLock()
	defer fake.getBlockchainInfoMutex.RUnlock()
	return len(fake.getBlockchainInfoArgsForCall)
}

func (fake *Ledger) GetBlockchainInfoCalls(stub func() (*common.BlockchainInfo, error)) {
	fake.getBlockchainInfoMutex.Lock()
	defer fake.getBlockchainInfoMutex.Unlock()
	fake.GetBlockchainInfoStub = stub
}

func (fake *Ledger) GetBlockchainInfoReturns(result1 *common.BlockchainInfo, result2 error) {
	fake.getBlockchainInfoMutex.Lock()
	defer fake.getBlockchainInfoMutex.Unlock()
	fake.GetBlockchainInfoStub = nil
	fake.getBlockchainInfoReturns = struct {
		result1 *common.BlockchainInfo
		result2 error
	}{result1, result2}
}

func (fake *Ledger) GetBlockchainInfoReturnsOnCall(i int, result1 *common.BlockchainInfo, result2 error) {
	fake.getBlockchainInfoMutex.Lock()
	defer fake.getBlockchainInfoMutex.Unlock()
	fake.GetBlockchainInfoStub = nil
	if fake.getBlockchainInfoReturnsOnCall == nil {
		fake.getBlockchainInfoReturnsOnCall = make(map[int]struct {
			result1 *common.BlockchainInfo
			result2 error
		})
	}
	fake.getBlockchainInfoReturnsOnCall[i] = struct {
		result1 *common.BlockchainInfo
		result2 error
	}{result1, result2}
}

func (fake *Ledger) GetBlocksIterator(arg1 uint64) (ledgerb.ResultsIterator, error) {
	fake.getBlocksIteratorMutex.Lock()
	ret, specificReturn := fake.getBlocksIteratorReturnsOnCall[len(fake.getBlocksIteratorArgsForCall)]
	fake.getBlocksIteratorArgsForCall = append(fake.getBlocksIteratorArgsForCall, struct {
		arg1 uint64
	}{arg1})
	stub := fake.GetBlocksIteratorStub
	fakeReturns := fake.getBlocksIteratorReturns
	fake.recordInvocation("GetBlocksIterator", []interface{}{arg1})
	fake.getBlocksIteratorMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *Ledger) GetBlocksIteratorCallCount() int {
	fake.getBlocksIteratorMutex.RLock()
	defer fake.getBlocksIteratorMutex.RUnlock()
	return len(fake.getBlocksIteratorArgsForCall)
}

func (fake *Ledger) GetBlocksIteratorCalls(stub func(uint64) (ledgerb.ResultsIterator, error)) {
	fake.getBlocksIteratorMutex.Lock()
	defer fake.getBlocksIteratorMutex.Unlock()
	fake.GetBlocksIteratorStub = stub
}

func (fake *Ledger) GetBlocksIteratorArgsForCall(i int) uint64 {
	fake.getBlocksIteratorMutex.RLock()
	defer fake.getBlocksIteratorMutex.RUnlock()
	argsForCall := fake.getBlocksIteratorArgsForCall[i]
	return argsForCall.arg1
}

func (fake *Ledger) GetBlocksIteratorReturns(result1 ledgerb.ResultsIterator, result2 error) {
	fake.getBlocksIteratorMutex.Lock()
	defer fake.getBlocksIteratorMutex.Unlock()
	fake.GetBlocksIteratorStub = nil
	fake.getBlocksIteratorReturns = struct {
		result1 ledgerb.ResultsIterator
		result2 error
	}{result1, result2}
}

func (fake *Ledger) GetBlocksIteratorReturnsOnCall(i int, result1 ledgerb.ResultsIterator, result2 error) {
	fake.getBlocksIteratorMutex.Lock()
	defer fake.getBlocksIteratorMutex.Unlock()
	fake.GetBlocksIteratorStub = nil
	if fake.getBlocksIteratorReturnsOnCall == nil {
		fake.getBlocksIteratorReturnsOnCall = make(map[int]struct {
			result1 ledgerb.ResultsIterator
			result2 error
		})
	}
	fake.getBlocksIteratorReturnsOnCall[i] = struct {
		result1 ledgerb.ResultsIterator
		result2 error
	}{result1, result2}
}

func (fake *Ledger) GetTxValidationCodeByTxID(arg1 string) (peer.TxValidationCode, uint64, error) {
	fake.getTxValidationCodeByTxIDMutex.Lock()
	ret, specificReturn := fake.getTxValidationCodeByTxIDReturnsOnCall[len(fake.getTxValidationCodeByTxIDArgsForCall)]
	fake.getTxValidationCodeByTxIDArgsForCall = append(fake.getTxValidationCodeByTxIDArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetTxValidationCodeByTxIDStub
	fakeReturns := fake.getTxValidationCodeByTxIDReturns
	fake.recordInvocation("GetTxValidationCodeByTxID", []interface{}{arg1})
	fake.getTxValidationCodeByTxIDMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *Ledger) GetTxValidationCodeByTxIDCallCount() int {
	fake.getTxValidationCodeByTxIDMutex.RLock()
	defer fake.getTxValidationCodeByTxIDMutex.RUnlock()
	return len(fake.getTxValidationCodeByTxIDArgsForCall)
}

func (fake *Ledger) GetTxValidationCodeByTxIDCalls(stub func(string) (peer.TxValidationCode, uint64, error)) {
	fake.getTxValidationCodeByTxIDMutex.Lock()
	defer fake.getTxValidationCodeByTxIDMutex.Unlock()
	fake.GetTxValidationCodeByTxIDStub = stub
}

func (fake *Ledger) GetTxValidationCodeByTxIDArgsForCall(i int) string {
	fake.getTxValidationCodeByTxIDMutex.RLock()
	defer fake.getTxValidationCodeByTxIDMutex.RUnlock()
	argsForCall := fake.getTxValidationCodeByTxIDArgsForCall[i]
	return argsForCall.arg1
}

func (fake *Ledger) GetTxValidationCodeByTxIDReturns(result1 peer.TxValidationCode, result2 uint64, result3 error) {
	fake.getTxValidationCodeByTxIDMutex.Lock()
	defer fake.getTxValidationCodeByTxIDMutex.Unlock()
	fake.GetTxValidationCodeByTxIDStub = nil
	fake.getTxValidationCodeByTxIDReturns = struct {
		result1 peer.TxValidationCode
		result2 uint64
		result3 error
	}{result1, result2, result3}
}

func (fake *Ledger) GetTxValidationCodeByTxIDReturnsOnCall(i int, result1 peer.TxValidationCode, result2 uint64, result3 error) {
	fake.getTxValidationCodeByTxIDMutex.Lock()
	defer fake.getTxValidationCodeByTxIDMutex.Unlock()
	fake.GetTxValidationCodeByTxIDStub = nil
	if fake.getTxValidationCodeByTxIDReturnsOnCall == nil {
		fake.getTxValidationCodeByTxIDReturnsOnCall = make(map[int]struct {
			result1 peer.TxValidationCode
			result2 uint64
			result3 error
		})
	}
	fake.getTxValidationCodeByTxIDReturnsOnCall[i] = struct {
		result1 peer.TxValidationCode
		result2 uint64
		result3 error
	}{result1, result2, result3}
}

func (fake *Ledger) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.commitNotificationsChannelMutex.RLock()
	defer fake.commitNotificationsChannelMutex.RUnlock()
	fake.getBlockByTxIDMutex.RLock()
	defer fake.getBlockByTxIDMutex.RUnlock()
	fake.getBlockchainInfoMutex.RLock()
	defer fake.getBlockchainInfoMutex.RUnlock()
	fake.getBlocksIteratorMutex.RLock()
	defer fake.getBlocksIteratorMutex.RUnlock()
	fake.getTxValidationCodeByTxIDMutex.RLock()
	defer fake.getTxValidationCodeByTxIDMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Ledger) recordInvocation(key string, args []interface{}) {
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

var _ ledger.Ledger = new(Ledger)
