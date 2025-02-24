/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package tests

import (
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/hxx258456/fabric-protos-go-cc/ledger/rwset"
	"github.com/hxx258456/fabric-protos-go-cc/ledger/rwset/kvrwset"
	peer "github.com/hxx258456/fabric-protos-go-cc/peer"
	"github.com/hxx258456/fabric/common/util"
	"github.com/stretchr/testify/require"
)

// TestNilValNoDeleteMarker tests for a special writeset which carries a nil value and yet the delete marker is set to false.
// This kind of write-set gets produced in previous versions. See FAB-18386 for more details.
func TestNilValNoDeleteMarker(t *testing.T) {
	env := newEnv(t)
	defer env.cleanup()
	env.initLedgerMgmt()

	testLedger := env.createTestLedgerFromGenesisBlk("test-ledger")
	testLedger.simulateDeployTx("cc1", []*collConf{
		{
			name: "coll1",
		},
	})
	testLedger.cutBlockAndCommitLegacy()

	testLedger.simulateDataTx("txid1", func(s *simulator) {
		s.setState("cc1", "pubKey1", "pubValue1")
		s.setPvtdata("cc1", "coll1", "pvtKey1", "pvtValue1")
		s.setPvtdata("cc1", "coll1", "pvtKey2", "pvtValue2")
	})
	testLedger.cutBlockAndCommitLegacy()

	testLedger.verifyPubState("cc1", "pubKey1", "pubValue1")
	testLedger.verifyPvtdataHashState("cc1", "coll1", "pvtKey1", util.ComputeSHA256ButSm3([]byte("pvtValue1")))
	testLedger.verifyPvtdataHashState("cc1", "coll1", "pvtKey2", util.ComputeSHA256ButSm3([]byte("pvtValue2")))
	testLedger.verifyPvtState("cc1", "coll1", "pvtKey1", "pvtValue1")
	testLedger.verifyPvtState("cc1", "coll1", "pvtKey2", "pvtValue2")

	// Handcraft writeset that includes a delete for each of the above keys. We handcraft here because the one generated by the simulator
	// will always have IsDelete flag true and we want to test when this flag is set to false and the actual value is nil. See FAB-18386 for
	// more details.
	pubWrites := &kvrwset.KVRWSet{
		Writes: []*kvrwset.KVWrite{
			{
				Key:      "pubKey1",
				IsDelete: false,
				Value:    nil,
			},
		},
	}

	hashedWrites := &kvrwset.HashedRWSet{
		HashedWrites: []*kvrwset.KVWriteHash{
			{
				KeyHash:   util.ComputeSHA256ButSm3([]byte("pvtKey1")),
				IsDelete:  false,
				ValueHash: nil,
			},
			{
				KeyHash:   util.ComputeSHA256ButSm3([]byte("pvtKey2")),
				IsDelete:  false,
				ValueHash: util.ComputeSHA256ButSm3([]byte{}),
			},
		},
	}

	pvtWrites := &kvrwset.KVRWSet{
		Writes: []*kvrwset.KVWrite{
			{
				Key:      "pvtKey1",
				IsDelete: false,
			},
			{
				Key:      "pvtKey2",
				IsDelete: false,
			},
		},
	}

	pubWritesBytes, err := proto.Marshal(pubWrites)
	require.NoError(t, err)

	hashedWritesBytes, err := proto.Marshal(hashedWrites)
	require.NoError(t, err)

	pvtWritesBytes, err := proto.Marshal(pvtWrites)
	require.NoError(t, err)

	pubRwset := &rwset.TxReadWriteSet{
		DataModel: rwset.TxReadWriteSet_KV,
		NsRwset: []*rwset.NsReadWriteSet{
			{
				Namespace: "cc1",
				Rwset:     pubWritesBytes,
				CollectionHashedRwset: []*rwset.CollectionHashedReadWriteSet{
					{
						CollectionName: "coll1",
						HashedRwset:    hashedWritesBytes,
						PvtRwsetHash:   util.ComputeSHA256ButSm3(pvtWritesBytes),
					},
				},
			},
		},
	}
	pubRwsetBytes, err := proto.Marshal(pubRwset)
	require.NoError(t, err)
	envelope, err := constructTransaction("txid2", pubRwsetBytes)
	require.NoError(t, err)

	txAndPvtdata := &txAndPvtdata{
		Txid:     "txid2",
		Envelope: envelope,
		Pvtws: &rwset.TxPvtReadWriteSet{
			DataModel: rwset.TxReadWriteSet_KV,
			NsPvtRwset: []*rwset.NsPvtReadWriteSet{
				{
					Namespace: "cc1",
					CollectionPvtRwset: []*rwset.CollectionPvtReadWriteSet{
						{
							CollectionName: "coll1",
							Rwset:          pvtWritesBytes,
						},
					},
				},
			},
		},
	}

	testLedger.submitHandCraftedTx(txAndPvtdata)
	testLedger.cutBlockAndCommitLegacy()

	testLedger.verifyTxValidationCode("txid2", peer.TxValidationCode_VALID)
	testLedger.verifyPubState("cc1", "pubKey1", "")
	testLedger.verifyPvtdataHashState("cc1", "coll1", "pvtKey1", nil)
	testLedger.verifyPvtdataHashState("cc1", "coll1", "pvtKey2", nil)
	testLedger.verifyPvtState("cc1", "coll1", "pvtKey1", "")
	testLedger.verifyPvtState("cc1", "coll1", "pvtKey2", "")
	testLedger.verifyHistory("cc1", "pubKey1", []string{"", "pubValue1"})
}
