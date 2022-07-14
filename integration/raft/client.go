/*
Copyright IBM Corp All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package raft

import (
	"github.com/hxx258456/fabric-protos-go-cc/common"
	orderer "github.com/hxx258456/fabric-protos-go-cc/orderer"
	"github.com/hxx258456/fabric/integration/nwo"
	"github.com/hxx258456/fabric/integration/ordererclient"
	"github.com/hxx258456/fabric/protoutil"
	. "github.com/onsi/gomega"
)

func FetchBlock(n *nwo.Network, o *nwo.Orderer, seq uint64, channel string) *common.Block {
	denv := CreateDeliverEnvelope(n, o, seq, channel)
	Expect(denv).NotTo(BeNil())

	var blk *common.Block
	Eventually(func() error {
		var err error
		blk, err = ordererclient.Deliver(n, o, denv)
		return err
	}, n.EventuallyTimeout).ShouldNot(HaveOccurred())

	return blk
}

func CreateBroadcastEnvelope(n *nwo.Network, entity interface{}, channel string, data []byte) *common.Envelope {
	var signer *nwo.SigningIdentity
	switch creator := entity.(type) {
	case *nwo.Peer:
		signer = n.PeerUserSigner(creator, "Admin")
	case *nwo.Orderer:
		signer = n.OrdererUserSigner(creator, "Admin")
	}
	Expect(signer).NotTo(BeNil())

	env, err := protoutil.CreateSignedEnvelope(
		common.HeaderType_MESSAGE,
		channel,
		signer,
		&common.Envelope{Payload: data},
		0,
		0,
	)
	Expect(err).NotTo(HaveOccurred())

	return env
}

// CreateDeliverEnvelope creates a deliver env to seek for specified block.
func CreateDeliverEnvelope(n *nwo.Network, o *nwo.Orderer, blkNum uint64, channel string) *common.Envelope {
	specified := &orderer.SeekPosition{
		Type: &orderer.SeekPosition_Specified{
			Specified: &orderer.SeekSpecified{Number: blkNum},
		},
	}
	env, err := protoutil.CreateSignedEnvelope(
		common.HeaderType_DELIVER_SEEK_INFO,
		channel,
		n.OrdererUserSigner(o, "Admin"),
		&orderer.SeekInfo{
			Start:    specified,
			Stop:     specified,
			Behavior: orderer.SeekInfo_BLOCK_UNTIL_READY,
		},
		0,
		0,
	)
	Expect(err).NotTo(HaveOccurred())

	return env
}
