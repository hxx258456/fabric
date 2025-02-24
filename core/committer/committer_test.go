/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package committer

import (
	"testing"

	"github.com/hxx258456/fabric-protos-go-cc/common"
	peer "github.com/hxx258456/fabric-protos-go-cc/peer"
	"github.com/hxx258456/fabric/common/configtx/test"
	"github.com/hxx258456/fabric/common/ledger"
	"github.com/hxx258456/fabric/common/ledger/testutil"
	ledger2 "github.com/hxx258456/fabric/core/ledger"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type mockLedger struct {
	height       uint64
	currentHash  []byte
	previousHash []byte
	mock.Mock
}

func (m *mockLedger) GetConfigHistoryRetriever() (ledger2.ConfigHistoryRetriever, error) {
	args := m.Called()
	return args.Get(0).(ledger2.ConfigHistoryRetriever), args.Error(1)
}

func (m *mockLedger) GetBlockchainInfo() (*common.BlockchainInfo, error) {
	info := &common.BlockchainInfo{
		Height:            m.height,
		CurrentBlockHash:  m.currentHash,
		PreviousBlockHash: m.previousHash,
	}
	return info, nil
}

func (m *mockLedger) DoesPvtDataInfoExist(blkNum uint64) (bool, error) {
	args := m.Called()
	return args.Get(0).(bool), args.Error(1)
}

func (m *mockLedger) GetBlockByNumber(blockNumber uint64) (*common.Block, error) {
	args := m.Called(blockNumber)
	return args.Get(0).(*common.Block), args.Error(1)
}

func (m *mockLedger) GetBlocksIterator(startBlockNumber uint64) (ledger.ResultsIterator, error) {
	args := m.Called(startBlockNumber)
	return args.Get(0).(ledger.ResultsIterator), args.Error(1)
}

func (m *mockLedger) Close() {
}

// TxIDExists returns true if the specified txID is already present in one of the already committed blocks
func (m *mockLedger) TxIDExists(txID string) (bool, error) {
	args := m.Called(txID)
	return args.Get(0).(bool), args.Error(1)
}

func (m *mockLedger) GetTransactionByID(txID string) (*peer.ProcessedTransaction, error) {
	args := m.Called(txID)
	return args.Get(0).(*peer.ProcessedTransaction), args.Error(1)
}

func (m *mockLedger) GetBlockByHash(blockHash []byte) (*common.Block, error) {
	args := m.Called(blockHash)
	return args.Get(0).(*common.Block), args.Error(1)
}

func (m *mockLedger) GetBlockByTxID(txID string) (*common.Block, error) {
	args := m.Called(txID)
	return args.Get(0).(*common.Block), args.Error(1)
}

func (m *mockLedger) GetTxValidationCodeByTxID(txID string) (peer.TxValidationCode, error) {
	args := m.Called(txID)
	return args.Get(0).(peer.TxValidationCode), args.Error(1)
}

func (m *mockLedger) NewTxSimulator(txid string) (ledger2.TxSimulator, error) {
	args := m.Called(txid)
	return args.Get(0).(ledger2.TxSimulator), args.Error(1)
}

func (m *mockLedger) NewQueryExecutor() (ledger2.QueryExecutor, error) {
	args := m.Called()
	return args.Get(0).(ledger2.QueryExecutor), args.Error(1)
}

func (m *mockLedger) NewHistoryQueryExecutor() (ledger2.HistoryQueryExecutor, error) {
	args := m.Called()
	return args.Get(0).(ledger2.HistoryQueryExecutor), args.Error(1)
}

func (m *mockLedger) GetPvtDataAndBlockByNum(blockNum uint64, filter ledger2.PvtNsCollFilter) (*ledger2.BlockAndPvtData, error) {
	args := m.Called(blockNum, filter)
	return args.Get(0).(*ledger2.BlockAndPvtData), args.Error(1)
}

func (m *mockLedger) GetPvtDataByNum(blockNum uint64, filter ledger2.PvtNsCollFilter) ([]*ledger2.TxPvtData, error) {
	args := m.Called(blockNum, filter)
	return args.Get(0).([]*ledger2.TxPvtData), args.Error(1)
}

func (m *mockLedger) CommitLegacy(blockAndPvtdata *ledger2.BlockAndPvtData, commitOpts *ledger2.CommitOptions) error {
	m.height += 1
	m.previousHash = m.currentHash
	m.currentHash = blockAndPvtdata.Block.Header.DataHash
	args := m.Called(blockAndPvtdata)
	return args.Error(0)
}

func (m *mockLedger) CommitPvtDataOfOldBlocks(reconciledPvtdata []*ledger2.ReconciledPvtdata, unreconciled ledger2.MissingPvtDataInfo) ([]*ledger2.PvtdataHashMismatch, error) {
	panic("implement me")
}

func (m *mockLedger) GetMissingPvtDataTracker() (ledger2.MissingPvtDataTracker, error) {
	panic("implement me")
}

func createLedger(channelID string) (*common.Block, *mockLedger) {
	gb, _ := test.MakeGenesisBlock(channelID)
	ledger := &mockLedger{
		height:       1,
		previousHash: []byte{},
		currentHash:  gb.Header.DataHash,
	}
	return gb, ledger
}

func TestKVLedgerBlockStorage(t *testing.T) {
	t.Parallel()
	gb, ledger := createLedger("TestLedger")
	block1 := testutil.ConstructBlock(t, 1, gb.Header.DataHash, [][]byte{{1, 2, 3, 4}, {5, 6, 7, 8}}, true)

	ledger.On("CommitLegacy", mock.Anything).Run(func(args mock.Arguments) {
		b := args.Get(0).(*ledger2.BlockAndPvtData)
		require.Equal(t, uint64(1), b.Block.Header.GetNumber())
		require.Equal(t, gb.Header.DataHash, b.Block.Header.PreviousHash)
		require.Equal(t, block1.Header.DataHash, b.Block.Header.DataHash)
	}).Return(nil)

	ledger.On("GetBlockByNumber", uint64(0)).Return(gb, nil)

	committer := NewLedgerCommitter(ledger)
	height, err := committer.LedgerHeight()
	require.Equal(t, uint64(1), height)
	require.NoError(t, err)

	err = committer.CommitLegacy(&ledger2.BlockAndPvtData{Block: block1}, &ledger2.CommitOptions{})
	require.NoError(t, err)

	height, err = committer.LedgerHeight()
	require.Equal(t, uint64(2), height)
	require.NoError(t, err)

	blocks := committer.GetBlocks([]uint64{0})
	require.Equal(t, 1, len(blocks))
	require.NoError(t, err)
}
