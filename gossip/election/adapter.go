/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package election

import (
	"bytes"
	"sync"
	"time"

	proto "github.com/hxx258456/fabric-protos-go-cc/gossip"
	"github.com/hxx258456/fabric/gossip/common"
	"github.com/hxx258456/fabric/gossip/discovery"
	"github.com/hxx258456/fabric/gossip/metrics"
	"github.com/hxx258456/fabric/gossip/protoext"
	"github.com/hxx258456/fabric/gossip/util"
)

type msgImpl struct {
	msg *proto.GossipMessage
}

func (mi *msgImpl) SenderID() peerID {
	return mi.msg.GetLeadershipMsg().PkiId
}

func (mi *msgImpl) IsProposal() bool {
	return !mi.IsDeclaration()
}

func (mi *msgImpl) IsDeclaration() bool {
	return mi.msg.GetLeadershipMsg().IsDeclaration
}

type peerImpl struct {
	member discovery.NetworkMember
}

func (pi *peerImpl) ID() peerID {
	return peerID(pi.member.PKIid)
}

type gossip interface {
	// PeersOfChannel returns the NetworkMembers considered alive in a channel
	PeersOfChannel(channel common.ChannelID) []discovery.NetworkMember

	// Accept returns a dedicated read-only channel for messages sent by other nodes that match a certain predicate.
	// If passThrough is false, the messages are processed by the gossip layer beforehand.
	// If passThrough is true, the gossip layer doesn't intervene and the messages
	// can be used to send a reply back to the sender
	Accept(acceptor common.MessageAcceptor, passThrough bool) (<-chan *proto.GossipMessage, <-chan protoext.ReceivedMessage)

	// Gossip sends a message to other peers to the network
	Gossip(msg *proto.GossipMessage)

	// IsInMyOrg checks whether a network member is in this peer's org
	IsInMyOrg(member discovery.NetworkMember) bool
}

type adapterImpl struct {
	gossip    gossip
	selfPKIid common.PKIidType

	incTime uint64
	seqNum  uint64

	channel common.ChannelID

	logger util.Logger

	doneCh   chan struct{}
	stopOnce *sync.Once
	metrics  *metrics.ElectionMetrics
}

// NewAdapter creates new leader election adapter
func NewAdapter(gossip gossip, pkiid common.PKIidType, channel common.ChannelID,
	metrics *metrics.ElectionMetrics) LeaderElectionAdapter {
	return &adapterImpl{
		gossip:    gossip,
		selfPKIid: pkiid,

		incTime: uint64(time.Now().UnixNano()),
		seqNum:  uint64(0),

		channel: channel,

		logger: util.GetLogger(util.ElectionLogger, ""),

		doneCh:   make(chan struct{}),
		stopOnce: &sync.Once{},
		metrics:  metrics,
	}
}

func (ai *adapterImpl) Gossip(msg Msg) {
	ai.gossip.Gossip(msg.(*msgImpl).msg)
}

func (ai *adapterImpl) Accept() <-chan Msg {
	adapterCh, _ := ai.gossip.Accept(func(message interface{}) bool {
		// Get only leadership org and channel messages
		return message.(*proto.GossipMessage).Tag == proto.GossipMessage_CHAN_AND_ORG &&
			protoext.IsLeadershipMsg(message.(*proto.GossipMessage)) &&
			bytes.Equal(message.(*proto.GossipMessage).Channel, ai.channel)
	}, false)

	msgCh := make(chan Msg)

	go func(inCh <-chan *proto.GossipMessage, outCh chan Msg, stopCh chan struct{}) {
		for {
			select {
			case <-stopCh:
				return
			case gossipMsg, ok := <-inCh:
				if ok {
					outCh <- &msgImpl{gossipMsg}
				} else {
					return
				}
			}
		}
	}(adapterCh, msgCh, ai.doneCh)
	return msgCh
}

func (ai *adapterImpl) CreateMessage(isDeclaration bool) Msg {
	ai.seqNum++
	seqNum := ai.seqNum

	leadershipMsg := &proto.LeadershipMessage{
		PkiId:         ai.selfPKIid,
		IsDeclaration: isDeclaration,
		Timestamp: &proto.PeerTime{
			IncNum: ai.incTime,
			SeqNum: seqNum,
		},
	}

	msg := &proto.GossipMessage{
		Nonce:   0,
		Tag:     proto.GossipMessage_CHAN_AND_ORG,
		Content: &proto.GossipMessage_LeadershipMsg{LeadershipMsg: leadershipMsg},
		Channel: ai.channel,
	}
	return &msgImpl{msg}
}

func (ai *adapterImpl) Peers() []Peer {
	peers := ai.gossip.PeersOfChannel(ai.channel)

	var res []Peer
	for _, peer := range peers {
		if ai.gossip.IsInMyOrg(peer) {
			res = append(res, &peerImpl{peer})
		}
	}

	return res
}

func (ai *adapterImpl) ReportMetrics(isLeader bool) {
	var leadershipBit float64
	if isLeader {
		leadershipBit = 1
	}
	ai.metrics.Declaration.With("channel", string(ai.channel)).Set(leadershipBit)
}

func (ai *adapterImpl) Stop() {
	stopFunc := func() {
		close(ai.doneCh)
	}
	ai.stopOnce.Do(stopFunc)
}
