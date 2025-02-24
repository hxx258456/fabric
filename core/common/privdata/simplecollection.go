/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package privdata

import (
	"github.com/hxx258456/fabric-protos-go-cc/peer"
	"github.com/hxx258456/fabric/common/policies"
	"github.com/hxx258456/fabric/msp"
	"github.com/hxx258456/fabric/protoutil"
	"github.com/pkg/errors"
)

// SimpleCollection implements a collection with static properties
// and a public member set
type SimpleCollection struct {
	name         string
	accessPolicy policies.Policy
	memberOrgs   map[string]struct{}
	conf         peer.StaticCollectionConfig
}

type SimpleCollectionPersistenceConfigs struct {
	blockToLive uint64
}

// NewSimpleCollection returns a simple collection object based on a given
// StaticCollectionConfig proto that has all the necessary information
func NewSimpleCollection(collectionConfig *peer.StaticCollectionConfig, deserializer msp.IdentityDeserializer) (*SimpleCollection, error) {
	sc := &SimpleCollection{}
	err := sc.Setup(collectionConfig, deserializer)
	return sc, err
}

// CollectionID returns the collection's ID
func (sc *SimpleCollection) CollectionID() string {
	return sc.name
}

// MemberOrgs returns the MSP IDs that are part of this collection
func (sc *SimpleCollection) MemberOrgs() map[string]struct{} {
	return sc.memberOrgs
}

// RequiredPeerCount returns the minimum number of peers
// required to send private data to
func (sc *SimpleCollection) RequiredPeerCount() int {
	return int(sc.conf.RequiredPeerCount)
}

// MaximumPeerCount returns the maximum number of peers
// to which the private data will be sent
func (sc *SimpleCollection) MaximumPeerCount() int {
	return int(sc.conf.MaximumPeerCount)
}

// AccessFilter returns the member filter function that evaluates signed data
// against the member access policy of this collection
func (sc *SimpleCollection) AccessFilter() Filter {
	return func(sd protoutil.SignedData) bool {
		if err := sc.accessPolicy.EvaluateSignedData([]*protoutil.SignedData{&sd}); err != nil {
			return false
		}
		return true
	}
}

// IsMemberOnlyRead returns whether only collection member
// has the read permission
func (sc *SimpleCollection) IsMemberOnlyRead() bool {
	return sc.conf.MemberOnlyRead
}

// IsMemberOnlyWrite returns whether only collection member
// has the write permission
func (sc *SimpleCollection) IsMemberOnlyWrite() bool {
	return sc.conf.MemberOnlyWrite
}

// Setup configures a simple collection object based on a given
// StaticCollectionConfig proto that has all the necessary information
func (sc *SimpleCollection) Setup(collectionConfig *peer.StaticCollectionConfig, deserializer msp.IdentityDeserializer) error {
	if collectionConfig == nil {
		return errors.New("Nil config passed to collection setup")
	}
	sc.conf = *collectionConfig
	sc.name = collectionConfig.GetName()

	// get the access signature policy envelope
	collectionPolicyConfig := collectionConfig.GetMemberOrgsPolicy()
	if collectionPolicyConfig == nil {
		return errors.New("Collection config policy is nil")
	}
	accessPolicyEnvelope := collectionPolicyConfig.GetSignaturePolicy()
	if accessPolicyEnvelope == nil {
		return errors.New("Collection config access policy is nil")
	}

	err := sc.setupAccessPolicy(collectionPolicyConfig, deserializer)
	if err != nil {
		return err
	}

	// get member org MSP IDs from the envelope, identities that fail to deserialize will not be returned
	sc.memberOrgs = getMemberOrgs(accessPolicyEnvelope.Identities, deserializer)

	return nil
}

// setupAccessPolicy configures a simple collection object based on a given
// StaticCollectionConfig proto that has all the necessary information
func (sc *SimpleCollection) setupAccessPolicy(collectionPolicyConfig *peer.CollectionPolicyConfig, deserializer msp.IdentityDeserializer) error {
	var err error
	sc.accessPolicy, err = getPolicy(collectionPolicyConfig, deserializer)
	return err
}

// BlockToLive return collection's block to live configuration
func (s *SimpleCollectionPersistenceConfigs) BlockToLive() uint64 {
	return s.blockToLive
}
