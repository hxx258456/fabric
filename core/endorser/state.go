/*
Copyright IBM Corp. 2018 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package endorser

import (
	"github.com/hxx258456/fabric-protos-go-cc/ledger/rwset"
	endorsement "github.com/hxx258456/fabric/core/handlers/endorsement/api/state"
	"github.com/hxx258456/fabric/core/ledger"
	"github.com/hxx258456/fabric/core/transientstore"
	"github.com/pkg/errors"
)

//go:generate mockery -dir . -name QueryCreator -case underscore -output mocks/
// QueryCreator creates new QueryExecutors
type QueryCreator interface {
	NewQueryExecutor() (ledger.QueryExecutor, error)
}

// ChannelState defines state operations
type ChannelState struct {
	*transientstore.Store
	QueryCreator
}

// FetchState fetches state
func (cs *ChannelState) FetchState() (endorsement.State, error) {
	qe, err := cs.NewQueryExecutor()
	if err != nil {
		return nil, err
	}

	return &StateContext{
		QueryExecutor: qe,
		Store:         cs.Store,
	}, nil
}

// StateContext defines an execution context that interacts with the state
type StateContext struct {
	*transientstore.Store
	ledger.QueryExecutor
}

// GetTransientByTXID returns the private data associated with this transaction ID.
func (sc *StateContext) GetTransientByTXID(txID string) ([]*rwset.TxPvtReadWriteSet, error) {
	scanner, err := sc.Store.GetTxPvtRWSetByTxid(txID, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer scanner.Close()
	var data []*rwset.TxPvtReadWriteSet
	for {
		res, err := scanner.Next()
		if err != nil {
			return nil, errors.WithStack(err)
		}
		if res == nil {
			break
		}
		if res.PvtSimulationResultsWithConfig == nil {
			continue
		}
		data = append(data, res.PvtSimulationResultsWithConfig.PvtRwset)
	}
	return data, nil
}
