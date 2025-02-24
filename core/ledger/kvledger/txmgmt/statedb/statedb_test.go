/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package statedb

import (
	"sort"
	"testing"

	"github.com/hxx258456/fabric/core/ledger/internal/version"
	"github.com/stretchr/testify/require"
)

func TestPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("Nil value to Put() did not panic\n")
		}
	}()

	batch := NewUpdateBatch()
	// The following call to Put() should result in panic
	batch.Put("ns1", "key1", nil, nil)
}

// Test Put(), Get(), and Delete()
func TestPutGetDeleteExistsGetUpdates(t *testing.T) {
	batch := NewUpdateBatch()
	batch.Put("ns1", "key1", []byte("value1"), version.NewHeight(1, 1))

	// Get() should return above inserted <k,v> pair
	actualVersionedValue := batch.Get("ns1", "key1")
	require.Equal(t, &VersionedValue{Value: []byte("value1"), Version: version.NewHeight(1, 1)}, actualVersionedValue)
	// Exists() should return false as key2 does not exist
	actualResult := batch.Exists("ns1", "key2")
	expectedResult := false
	require.Equal(t, expectedResult, actualResult)

	// Exists() should return false as ns3 does not exist
	actualResult = batch.Exists("ns3", "key2")
	expectedResult = false
	require.Equal(t, expectedResult, actualResult)

	// Get() should return nil as key2 does not exist
	actualVersionedValue = batch.Get("ns1", "key2")
	require.Nil(t, actualVersionedValue)
	// Get() should return nil as ns3 does not exist
	actualVersionedValue = batch.Get("ns3", "key2")
	require.Nil(t, actualVersionedValue)

	batch.Put("ns1", "key2", []byte("value2"), version.NewHeight(1, 2))
	// Exists() should return true as key2 exists
	actualResult = batch.Exists("ns1", "key2")
	expectedResult = true
	require.Equal(t, expectedResult, actualResult)

	// GetUpdatedNamespaces should return 3 namespaces
	batch.Put("ns2", "key2", []byte("value2"), version.NewHeight(1, 2))
	batch.Put("ns3", "key2", []byte("value2"), version.NewHeight(1, 2))
	actualNamespaces := batch.GetUpdatedNamespaces()
	sort.Strings(actualNamespaces)
	expectedNamespaces := []string{"ns1", "ns2", "ns3"}
	require.Equal(t, expectedNamespaces, actualNamespaces)

	// GetUpdates should return two VersionedValues for the namespace ns1
	expectedUpdates := make(map[string]*VersionedValue)
	expectedUpdates["key1"] = &VersionedValue{Value: []byte("value1"), Version: version.NewHeight(1, 1)}
	expectedUpdates["key2"] = &VersionedValue{Value: []byte("value2"), Version: version.NewHeight(1, 2)}
	actualUpdates := batch.GetUpdates("ns1")
	require.Equal(t, expectedUpdates, actualUpdates)

	actualUpdates = batch.GetUpdates("ns4")
	require.Nil(t, actualUpdates)

	// Delete the above inserted <k,v> pair
	batch.Delete("ns1", "key2", version.NewHeight(1, 2))
	// Exists() should return true after deleting key2
	// Exists() should return true iff the key has action(Put/Delete) in this batch
	actualResult = batch.Exists("ns1", "key2")
	expectedResult = true
	require.Equal(t, expectedResult, actualResult)
}

func TestUpdateBatchIterator(t *testing.T) {
	batch := NewUpdateBatch()
	batch.Put("ns1", "key1", []byte("value1"), version.NewHeight(1, 1))
	batch.Put("ns1", "key2", []byte("value2"), version.NewHeight(1, 2))
	batch.Put("ns1", "key3", []byte("value3"), version.NewHeight(1, 3))

	batch.Put("ns2", "key6", []byte("value6"), version.NewHeight(2, 3))
	batch.Put("ns2", "key5", []byte("value5"), version.NewHeight(2, 2))
	batch.Put("ns2", "key4", []byte("value4"), version.NewHeight(2, 1))

	checkItrResults(t, batch.GetRangeScanIterator("ns1", "key2", "key3"), []*VersionedKV{
		{
			&CompositeKey{"ns1", "key2"},
			&VersionedValue{[]byte("value2"), nil, version.NewHeight(1, 2)},
		},
	})

	checkItrResults(t, batch.GetRangeScanIterator("ns2", "key0", "key8"), []*VersionedKV{
		{
			&CompositeKey{"ns2", "key4"},
			&VersionedValue{[]byte("value4"), nil, version.NewHeight(2, 1)},
		},
		{
			&CompositeKey{"ns2", "key5"},
			&VersionedValue{[]byte("value5"), nil, version.NewHeight(2, 2)},
		},
		{
			&CompositeKey{"ns2", "key6"},
			&VersionedValue{[]byte("value6"), nil, version.NewHeight(2, 3)},
		},
	})

	checkItrResults(t, batch.GetRangeScanIterator("ns2", "", ""), []*VersionedKV{
		{
			&CompositeKey{"ns2", "key4"},
			&VersionedValue{[]byte("value4"), nil, version.NewHeight(2, 1)},
		},
		{
			&CompositeKey{"ns2", "key5"},
			&VersionedValue{[]byte("value5"), nil, version.NewHeight(2, 2)},
		},
		{
			&CompositeKey{"ns2", "key6"},
			&VersionedValue{[]byte("value6"), nil, version.NewHeight(2, 3)},
		},
	})

	checkItrResults(t, batch.GetRangeScanIterator("non-existing-ns", "", ""), nil)
}

func checkItrResults(t *testing.T, itr QueryResultsIterator, expectedResults []*VersionedKV) {
	for i := 0; i < len(expectedResults); i++ {
		res, _ := itr.Next()
		require.Equal(t, expectedResults[i], res)
	}
	lastRes, err := itr.Next()
	require.NoError(t, err)
	require.Nil(t, lastRes)
	itr.Close()
}

func TestMergeUpdateBatch(t *testing.T) {
	batch1 := NewUpdateBatch()
	batch1.Put("ns1", "key1", []byte("batch1_value1"), version.NewHeight(1, 1))
	batch1.Put("ns1", "key2", []byte("batch1_value2"), version.NewHeight(2, 2))
	batch1.Put("ns1", "key3", []byte("batch1_value3"), version.NewHeight(3, 3))

	batch2 := NewUpdateBatch()
	batch2.ContainsPostOrderWrites = true
	batch2.Put("ns1", "key1", []byte("batch2_value1"), version.NewHeight(4, 4)) // overwrite key1 with new value
	batch2.Delete("ns1", "key2", version.NewHeight(5, 5))                       // overwrite key2 with deletion
	batch2.Put("ns1", "key4", []byte("batch2_value4"), version.NewHeight(6, 6)) // new key only in batch2
	batch2.Delete("ns1", "key5", version.NewHeight(7, 7))                       // delete key only in batch2
	batch2.Put("ns2", "key6", []byte("batch2_value6"), version.NewHeight(8, 8)) // namespace only in batch2

	batch1.Merge(batch2)

	// prepare final expected batch by writing all updates in the above order
	expectedBatch := NewUpdateBatch()
	expectedBatch.ContainsPostOrderWrites = true
	expectedBatch.Put("ns1", "key1", []byte("batch1_value1"), version.NewHeight(1, 1))
	expectedBatch.Put("ns1", "key2", []byte("batch1_value2"), version.NewHeight(2, 2))
	expectedBatch.Put("ns1", "key3", []byte("batch1_value3"), version.NewHeight(3, 3))
	expectedBatch.Put("ns1", "key1", []byte("batch2_value1"), version.NewHeight(4, 4))
	expectedBatch.Delete("ns1", "key2", version.NewHeight(5, 5))
	expectedBatch.Put("ns1", "key4", []byte("batch2_value4"), version.NewHeight(6, 6))
	expectedBatch.Delete("ns1", "key5", version.NewHeight(7, 7))
	expectedBatch.Put("ns2", "key6", []byte("batch2_value6"), version.NewHeight(8, 8))
	require.Equal(t, expectedBatch, batch1)
}
