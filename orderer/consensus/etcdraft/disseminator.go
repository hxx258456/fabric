/*
 Copyright IBM Corp All Rights Reserved.

 SPDX-License-Identifier: Apache-2.0
*/

package etcdraft

import (
	"sync"

	"github.com/hxx258456/fabric-protos-go-cc/orderer"
)

// Disseminator piggybacks cluster metadata, if any, to egress messages.
type Disseminator struct {
	RPC

	l        sync.Mutex
	sent     map[uint64]bool
	metadata []byte
}

func (d *Disseminator) SendConsensus(dest uint64, msg *orderer.ConsensusRequest) error {
	d.l.Lock()
	defer d.l.Unlock()

	if !d.sent[dest] && len(d.metadata) != 0 {
		msg.Metadata = d.metadata
		d.sent[dest] = true
	}

	return d.RPC.SendConsensus(dest, msg)
}

func (d *Disseminator) UpdateMetadata(m []byte) {
	d.l.Lock()
	defer d.l.Unlock()

	d.sent = make(map[uint64]bool)
	d.metadata = m
}
