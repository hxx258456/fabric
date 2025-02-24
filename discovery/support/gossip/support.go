/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package gossip

import (
	"github.com/hxx258456/fabric-protos-go-cc/gossip"
	"github.com/hxx258456/fabric/gossip/api"
	"github.com/hxx258456/fabric/gossip/common"
	"github.com/hxx258456/fabric/gossip/discovery"
	"github.com/hxx258456/fabric/gossip/protoext"
)

//go:generate counterfeiter -o mocks/gossip.go -fake-name Gossip . Gossip

type Gossip interface {
	// IdentityInfo returns identity information about peers
	IdentityInfo() api.PeerIdentitySet
	// GetPeers returns the NetworkMembers considered alive
	Peers() []discovery.NetworkMember
	// PeersOfChannel returns the NetworkMembers considered alive
	// and also subscribed to the channel given
	PeersOfChannel(common.ChannelID) []discovery.NetworkMember
	// SelfChannelInfo returns the peer's latest StateInfo message of a given channel
	SelfChannelInfo(common.ChannelID) *protoext.SignedGossipMessage
	// SelfMembershipInfo returns the peer's membership information
	SelfMembershipInfo() discovery.NetworkMember
}

// DiscoverySupport implements support that is used for service discovery
// that is obtained from gossip
type DiscoverySupport struct {
	Gossip
}

// NewDiscoverySupport creates a new DiscoverySupport
func NewDiscoverySupport(g Gossip) *DiscoverySupport {
	return &DiscoverySupport{g}
}

// ChannelExists returns whether a given channel exists or not
func (s *DiscoverySupport) ChannelExists(channel string) bool {
	return s.SelfChannelInfo(common.ChannelID(channel)) != nil
}

// PeersOfChannel returns the NetworkMembers considered alive
// and also subscribed to the channel given
func (s *DiscoverySupport) PeersOfChannel(chain common.ChannelID) discovery.Members {
	msg := s.SelfChannelInfo(chain)
	if msg == nil {
		return nil
	}
	stateInf := msg.GetStateInfo()
	selfMember := discovery.NetworkMember{
		Properties: stateInf.Properties,
		PKIid:      stateInf.PkiId,
		Envelope:   msg.Envelope,
	}
	return append(s.Gossip.PeersOfChannel(chain), selfMember)
}

// Peers returns the NetworkMembers considered alive
func (s *DiscoverySupport) Peers() discovery.Members {
	peers := s.Gossip.Peers()
	peers = append(peers, s.Gossip.SelfMembershipInfo())
	// Return only the peers that have an external endpoint, and sanitizes the envelopes.
	return discovery.Members(peers).Filter(discovery.HasExternalEndpoint).Map(sanitizeEnvelope)
}

func sanitizeEnvelope(member discovery.NetworkMember) discovery.NetworkMember {
	// Make a local copy of the member
	returnedMember := member
	if returnedMember.Envelope == nil {
		return returnedMember
	}
	returnedMember.Envelope = &gossip.Envelope{
		Payload:   member.Envelope.Payload,
		Signature: member.Envelope.Signature,
	}
	return returnedMember
}
