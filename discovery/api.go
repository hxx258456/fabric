/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package discovery

import (
	discprotos "github.com/hxx258456/fabric-protos-go-cc/discovery"
	peer "github.com/hxx258456/fabric-protos-go-cc/peer"
	"github.com/hxx258456/fabric/gossip/api"
	"github.com/hxx258456/fabric/gossip/common"
	"github.com/hxx258456/fabric/gossip/discovery"
	"github.com/hxx258456/fabric/protoutil"
)

// AccessControlSupport checks if clients are eligible of being serviced
type AccessControlSupport interface {
	// Eligible returns whether the given peer is eligible for receiving
	// service from the discovery service for a given channel
	EligibleForService(channel string, data protoutil.SignedData) error
}

// ConfigSequenceSupport returns the config sequence of the given channel
type ConfigSequenceSupport interface {
	// ConfigSequence returns the configuration sequence of the a given channel
	ConfigSequence(channel string) uint64
}

// GossipSupport aggregates abilities that the gossip module
// provides to the discovery service, such as knowing information about peers
type GossipSupport interface {
	// ChannelExists returns whether a given channel exists or not
	ChannelExists(channel string) bool

	// PeersOfChannel returns the NetworkMembers considered alive
	// and also subscribed to the channel given
	PeersOfChannel(common.ChannelID) discovery.Members

	// Peers returns the NetworkMembers considered alive
	Peers() discovery.Members

	// IdentityInfo returns identity information about peers
	IdentityInfo() api.PeerIdentitySet
}

// EndorsementSupport provides knowledge of endorsement policy selection
// for chaincodes
type EndorsementSupport interface {
	// PeersForEndorsement returns an EndorsementDescriptor for a given set of peers, channel, and chaincode
	PeersForEndorsement(channel common.ChannelID, interest *peer.ChaincodeInterest) (*discprotos.EndorsementDescriptor, error)

	// PeersAuthorizedByCriteria returns the peers of the channel that are authorized by the given chaincode interest
	// That is - taking in account if the chaincode(s) in the interest are installed on the peers, and also
	// taking in account whether the peers are part of the collections of the chaincodes.
	// If a nil interest, or an empty interest is passed - no filtering is done.
	PeersAuthorizedByCriteria(chainID common.ChannelID, interest *peer.ChaincodeInterest) (discovery.Members, error)
}

// ConfigSupport provides access to channel configuration
type ConfigSupport interface {
	// Config returns the channel's configuration
	Config(channel string) (*discprotos.ConfigResult, error)
}

// Support defines an interface that allows the discovery service
// to obtain information that other peer components have
type Support interface {
	AccessControlSupport
	GossipSupport
	EndorsementSupport
	ConfigSupport
	ConfigSequenceSupport
}
