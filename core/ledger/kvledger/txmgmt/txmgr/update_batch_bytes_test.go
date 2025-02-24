/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package txmgr

import (
	"testing"

	proto "github.com/golang/protobuf/proto"
	"github.com/hxx258456/fabric/core/ledger/internal/version"
	"github.com/hxx258456/fabric/core/ledger/kvledger/txmgmt/privacyenabledstate"
	"github.com/stretchr/testify/require"
)

func TestUpdateBatchBytesBuilderOnlyPublicWrites(t *testing.T) {
	updateBatch := privacyenabledstate.NewUpdateBatch()
	updateBatch.PubUpdates.Put("ns1", "key1", []byte("value1"), version.NewHeight(1, 1))
	updateBatch.PubUpdates.Put("ns1", "key2", []byte("value2"), version.NewHeight(1, 2))
	updateBatch.PubUpdates.Put("ns2", "key3", []byte("value3"), version.NewHeight(1, 3))
	updateBatch.PubUpdates.Put("ns3", "key4", []byte("value4"), version.NewHeight(1, 4))
	updateBatch.PubUpdates.Put("ns3", "key5", []byte("value5"), version.NewHeight(1, 5))
	updateBatch.PubUpdates.Delete("ns3", "key6", version.NewHeight(1, 6))

	bytes, err := deterministicBytesForPubAndHashUpdates(updateBatch)
	require.NoError(t, err)
	require.True(t, len(bytes) > 0)
	for i := 0; i < 100; i++ {
		b, _ := deterministicBytesForPubAndHashUpdates(updateBatch)
		require.Equal(t, bytes, b)
	}

	expectedUpdates := &Updates{
		Kvwrites: []*KVWrite{
			{
				Namespace:    "ns1",
				Key:          []byte("key1"),
				Value:        []byte("value1"),
				VersionBytes: version.NewHeight(1, 1).ToBytes(),
			},
			{
				Key:          []byte("key2"), // Namespace should not be present, if same as for the previous entry
				Value:        []byte("value2"),
				VersionBytes: version.NewHeight(1, 2).ToBytes(),
			},
			{
				Namespace:    "ns2",
				Key:          []byte("key3"),
				Value:        []byte("value3"),
				VersionBytes: version.NewHeight(1, 3).ToBytes(),
			},
			{
				Namespace:    "ns3",
				Key:          []byte("key4"),
				Value:        []byte("value4"),
				VersionBytes: version.NewHeight(1, 4).ToBytes(),
			},
			{
				Key:          []byte("key5"),
				Value:        []byte("value5"),
				VersionBytes: version.NewHeight(1, 5).ToBytes(),
			},
			{
				Key:          []byte("key6"),
				IsDelete:     true,
				VersionBytes: version.NewHeight(1, 6).ToBytes(),
			},
		},
	}
	expectedBytes, err := proto.Marshal(expectedUpdates)
	require.NoError(t, err)
	require.Equal(t, expectedBytes, bytes)
}

func TestUpdateBatchBytesBuilderPublicWritesAndColls(t *testing.T) {
	updateBatch := privacyenabledstate.NewUpdateBatch()
	updateBatch.PubUpdates.Put("ns1", "key1", []byte("value1"), version.NewHeight(1, 1))
	updateBatch.PubUpdates.Put("ns1", "key2", []byte("value2"), version.NewHeight(1, 2))
	updateBatch.HashUpdates.Put("ns1", "coll1", []byte("key3"), []byte("value3"), version.NewHeight(1, 3))
	updateBatch.HashUpdates.Put("ns1", "coll1", []byte("key4"), []byte("value4"), version.NewHeight(1, 4))
	updateBatch.HashUpdates.Put("ns1", "coll2", []byte("key5"), []byte("value5"), version.NewHeight(1, 5))
	updateBatch.HashUpdates.Delete("ns1", "coll2", []byte("key6"), version.NewHeight(1, 6))
	updateBatch.PubUpdates.Put("ns2", "key7", []byte("value7"), version.NewHeight(1, 7))

	bytes, err := deterministicBytesForPubAndHashUpdates(updateBatch)
	require.NoError(t, err)
	require.True(t, len(bytes) > 0)
	for i := 0; i < 100; i++ {
		b, _ := deterministicBytesForPubAndHashUpdates(updateBatch)
		require.Equal(t, bytes, b)
	}

	expectedUpdates := &Updates{
		Kvwrites: []*KVWrite{
			{
				Namespace:    "ns1",
				Key:          []byte("key1"),
				Value:        []byte("value1"),
				VersionBytes: version.NewHeight(1, 1).ToBytes(),
			},
			{
				Key:          []byte("key2"), // Namespace should not be present, if same as for the previous entry
				Value:        []byte("value2"),
				VersionBytes: version.NewHeight(1, 2).ToBytes(),
			},
			{
				Collection:   "coll1",
				Key:          []byte("key3"),
				Value:        []byte("value3"),
				VersionBytes: version.NewHeight(1, 3).ToBytes(),
			},
			{
				Key:          []byte("key4"), // Collection should not be present, if same as for the previous entry
				Value:        []byte("value4"),
				VersionBytes: version.NewHeight(1, 4).ToBytes(),
			},
			{
				Collection:   "coll2",
				Key:          []byte("key5"),
				Value:        []byte("value5"),
				VersionBytes: version.NewHeight(1, 5).ToBytes(),
			},
			{
				Key:          []byte("key6"),
				IsDelete:     true,
				VersionBytes: version.NewHeight(1, 6).ToBytes(),
			},
			{
				Namespace:    "ns2",
				Key:          []byte("key7"),
				Value:        []byte("value7"),
				VersionBytes: version.NewHeight(1, 7).ToBytes(),
			},
		},
	}
	expectedBytes, err := proto.Marshal(expectedUpdates)
	require.NoError(t, err)
	require.Equal(t, expectedBytes, bytes)
}

func TestUpdateBatchBytesBuilderOnlyChannelConfig(t *testing.T) {
	updateBatch := privacyenabledstate.NewUpdateBatch()
	updateBatch.PubUpdates.Put("", "resourcesconfigtx.CHANNEL_CONFIG_KEY", []byte("value1"), version.NewHeight(1, 1))

	bytes, err := deterministicBytesForPubAndHashUpdates(updateBatch)
	require.NoError(t, err)
	expectedUpdates := &Updates{}
	expectedBytes, err := proto.Marshal(expectedUpdates)
	require.NoError(t, err)
	require.Equal(t, expectedBytes, bytes)
}
