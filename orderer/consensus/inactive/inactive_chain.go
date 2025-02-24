/*
Copyright IBM Corp. 2017 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package inactive

import (
	"github.com/hxx258456/fabric-protos-go-cc/common"
	"github.com/hxx258456/fabric/orderer/common/types"
)

// Chain implements an inactive consenter.Chain
// which is used to denote that the current orderer node
// does not service a specific channel.
type Chain struct {
	Err error
}

func (c *Chain) Order(_ *common.Envelope, _ uint64) error {
	return c.Err
}

func (c *Chain) Configure(_ *common.Envelope, _ uint64) error {
	return c.Err
}

func (c *Chain) WaitReady() error {
	return c.Err
}

func (*Chain) Errored() <-chan struct{} {
	closedChannel := make(chan struct{})
	close(closedChannel)
	return closedChannel
}

func (c *Chain) Start() {
}

func (c *Chain) Halt() {
}

// StatusReport returns the ConsensusRelation & Status
func (c *Chain) StatusReport() (types.ConsensusRelation, types.Status) {
	return types.ConsensusRelationConfigTracker, types.StatusInactive
}
