/*
Copyright IBM Corp. 2016 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package validation

import (
	"testing"

	cb "github.com/hxx258456/fabric-protos-go-cc/common"
	peer "github.com/hxx258456/fabric-protos-go-cc/peer"
	"github.com/hxx258456/fabric/bccsp/sw"
	"github.com/hxx258456/fabric/core/config/configtest"
	"github.com/hxx258456/fabric/internal/configtxgen/encoder"
	"github.com/hxx258456/fabric/internal/configtxgen/genesisconfig"
	"github.com/hxx258456/fabric/protoutil"
	"github.com/stretchr/testify/require"
)

func TestValidateConfigTx(t *testing.T) {
	channelID := "testchannelid"
	profile := genesisconfig.Load(genesisconfig.SampleSingleMSPChannelProfile, configtest.GetDevConfigDir())
	chCrtEnv, err := encoder.MakeChannelCreationTransaction(genesisconfig.SampleConsortiumName, nil, profile)
	if err != nil {
		t.Fatalf("MakeChannelCreationTransaction failed, err %s", err)
		return
	}

	updateResult := &cb.Envelope{
		Payload: protoutil.MarshalOrPanic(&cb.Payload{
			Header: &cb.Header{
				ChannelHeader: protoutil.MarshalOrPanic(&cb.ChannelHeader{
					Type:      int32(cb.HeaderType_CONFIG),
					ChannelId: channelID,
				}),
				SignatureHeader: protoutil.MarshalOrPanic(&cb.SignatureHeader{
					Creator: signerSerialized,
					Nonce:   protoutil.CreateNonceOrPanic(),
				}),
			},
			Data: protoutil.MarshalOrPanic(&cb.ConfigEnvelope{
				LastUpdate: chCrtEnv,
			}),
		}),
	}
	cryptoProvider, err := sw.NewDefaultSecurityLevelWithKeystore(sw.NewDummyKeyStore())
	require.NoError(t, err)
	updateResult.Signature, _ = signer.Sign(updateResult.Payload)
	_, txResult := ValidateTransaction(updateResult, cryptoProvider)
	if txResult != peer.TxValidationCode_VALID {
		t.Fatalf("ValidateTransaction failed, err %s", err)
		return
	}
}
