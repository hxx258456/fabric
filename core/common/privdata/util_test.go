/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package privdata

import (
	"testing"

	"github.com/hxx258456/fabric-protos-go-cc/peer"
	"github.com/stretchr/testify/require"
)

func TestGetPolicy(t *testing.T) {
	accessPolicy, err := getPolicy(getAccessPolicy([]string{"signer0", "signer1"}), &mockDeserializer{})
	require.NotNil(t, accessPolicy)
	require.Nil(t, err)
}

func TestGetPolicyFailed(t *testing.T) {
	// nil policy config
	_, err := getPolicy(nil, &mockDeserializer{})
	require.EqualError(t, err, "collection policy config is nil")

	// nil collectionPolicyConfig.GetSignaturePolicy()
	_, err = getPolicy(&peer.CollectionPolicyConfig{}, &mockDeserializer{})
	require.EqualError(t, err, "collection config access policy is nil")

	// faulty policy config: index out of range
	_, err = getPolicy(getBadAccessPolicy([]string{"signer0", "signer1"}, 3), &mockDeserializer{})
	require.EqualError(t, err, "failed constructing policy object out of collection policy config: identity index out of range, requested 3, but identities length is 2")
}
