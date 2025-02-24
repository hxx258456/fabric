/*
Copyright IBM Corp, SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package auth

import (
	"context"
	"encoding/binary"
	"testing"

	"github.com/hxx258456/fabric-protos-go-cc/peer"
	"github.com/stretchr/testify/require"
)

func TestChainFilters(t *testing.T) {
	iterations := 15
	filters := createNFilters(iterations)
	endorser := &mockEndorserServer{}
	initialProposal := &peer.SignedProposal{ProposalBytes: make([]byte, 4)}
	binary.BigEndian.PutUint32(initialProposal.ProposalBytes, 0)

	firstFilter := ChainFilters(endorser, filters...)
	firstFilter.ProcessProposal(context.Background(), initialProposal)
	for i := 0; i < iterations; i++ {
		require.Equal(t, uint32(i), filters[i].(*mockAuthFilter).sequence,
			"Expected filters to be invoked in the provided sequence")
	}

	require.Equal(t, uint32(iterations), endorser.sequence,
		"Expected endorser to be invoked after filters")

	// Test with no filters
	binary.BigEndian.PutUint32(initialProposal.ProposalBytes, 0)
	firstFilter = ChainFilters(endorser)
	firstFilter.ProcessProposal(context.Background(), initialProposal)
	require.Equal(t, uint32(0), endorser.sequence,
		"Expected endorser to be invoked first")
}

func createNFilters(n int) []Filter {
	filters := make([]Filter, n)
	for i := 0; i < n; i++ {
		filters[i] = &mockAuthFilter{}
	}
	return filters
}

type mockEndorserServer struct {
	sequence uint32
}

func (es *mockEndorserServer) ProcessProposal(ctx context.Context, prop *peer.SignedProposal) (*peer.ProposalResponse, error) {
	es.sequence = binary.BigEndian.Uint32(prop.ProposalBytes)
	binary.BigEndian.PutUint32(prop.ProposalBytes, es.sequence+1)
	return nil, nil
}

type mockAuthFilter struct {
	sequence uint32
	next     peer.EndorserServer
}

func (f *mockAuthFilter) ProcessProposal(ctx context.Context, prop *peer.SignedProposal) (*peer.ProposalResponse, error) {
	f.sequence = binary.BigEndian.Uint32(prop.ProposalBytes)
	binary.BigEndian.PutUint32(prop.ProposalBytes, f.sequence+1)
	return f.next.ProcessProposal(ctx, prop)
}

func (f *mockAuthFilter) Init(next peer.EndorserServer) {
	f.next = next
}
