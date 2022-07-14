/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package channel

import (
	"testing"

	"github.com/golang/protobuf/proto"
	cb "github.com/hxx258456/fabric-protos-go-cc/common"
	pb "github.com/hxx258456/fabric-protos-go-cc/peer"
	"github.com/hxx258456/fabric/internal/peer/common"
	"github.com/stretchr/testify/require"
)

func TestGetChannelInfo(t *testing.T) {
	InitMSP()
	resetFlags()

	mockBlockchainInfo := &cb.BlockchainInfo{
		Height:            1,
		CurrentBlockHash:  []byte("CurrentBlockHash"),
		PreviousBlockHash: []byte("PreviousBlockHash"),
	}
	mockPayload, err := proto.Marshal(mockBlockchainInfo)
	require.NoError(t, err)

	mockResponse := &pb.ProposalResponse{
		Response: &pb.Response{
			Status:  200,
			Payload: mockPayload,
		},
		Endorsement: &pb.Endorsement{},
	}

	signer, err := common.GetDefaultSigner()
	require.NoError(t, err)

	mockCF := &ChannelCmdFactory{
		EndorserClient:   common.GetMockEndorserClient(mockResponse, nil),
		BroadcastFactory: mockBroadcastClientFactory,
		Signer:           signer,
	}

	cmd := getinfoCmd(mockCF)
	AddFlags(cmd)

	args := []string{"-c", mockChannel}
	cmd.SetArgs(args)

	require.NoError(t, cmd.Execute())
}

func TestGetChannelInfoMissingChannelID(t *testing.T) {
	InitMSP()
	resetFlags()

	signer, err := common.GetDefaultSigner()
	if err != nil {
		t.Fatalf("Get default signer error: %v", err)
	}

	mockCF := &ChannelCmdFactory{
		Signer: signer,
	}

	cmd := getinfoCmd(mockCF)

	AddFlags(cmd)

	cmd.SetArgs([]string{})

	require.Error(t, cmd.Execute())
}
