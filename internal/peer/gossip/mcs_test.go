/*
Copyright IBM Corp. 2017 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package gossip

import (
	"crypto/sha256"
	"errors"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/hxx258456/fabric-protos-go-cc/common"
	pmsp "github.com/hxx258456/fabric-protos-go-cc/msp"
	protospeer "github.com/hxx258456/fabric-protos-go-cc/peer"
	"github.com/hxx258456/fabric/bccsp/sw"
	"github.com/hxx258456/fabric/common/policies"
	"github.com/hxx258456/fabric/common/util"
	"github.com/hxx258456/fabric/gossip/api"
	"github.com/hxx258456/fabric/internal/peer/gossip/mocks"
	"github.com/hxx258456/fabric/internal/pkg/identity"
	"github.com/hxx258456/fabric/msp"
	"github.com/hxx258456/fabric/msp/mgmt"
	"github.com/hxx258456/fabric/protoutil"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

//go:generate counterfeiter -o mocks/policy_manager.go -fake-name PolicyManager . policyManager

type policyManager interface {
	policies.Manager
}

//go:generate counterfeiter -o mocks/signer_serializer.go --fake-name SignerSerializer . signerSerializer

type signerSerializer interface {
	identity.SignerSerializer
}

func TestPKIidOfCert(t *testing.T) {
	deserializersManager := &mocks.DeserializersManager{
		LocalDeserializer: &mocks.IdentityDeserializer{Identity: []byte("Alice"), Msg: []byte("msg1"), Mock: mock.Mock{}},
	}
	signer := &mocks.SignerSerializer{}
	signer.SerializeReturns([]byte("Alice"), nil)
	cryptoProvider, err := sw.NewDefaultSecurityLevelWithKeystore(sw.NewDummyKeyStore())
	require.NoError(t, err)
	msgCryptoService := NewMCS(
		&mocks.ChannelPolicyManagerGetterWithManager{},
		signer,
		deserializersManager,
		cryptoProvider,
	)

	peerIdentity := []byte("Alice")
	pkid := msgCryptoService.GetPKIidOfCert(peerIdentity)

	// Check pkid is not nil
	require.NotNil(t, pkid, "PKID must be different from nil")
	// Check that pkid is correctly computed
	id, err := deserializersManager.Deserialize(peerIdentity)
	require.NoError(t, err, "Failed getting validated identity from [% x]", []byte(peerIdentity))
	idRaw := append([]byte(id.Mspid), id.IdBytes...)
	require.NoError(t, err, "Failed marshalling identity identifier [% x]: [%s]", peerIdentity, err)
	h := sha256.New()
	h.Write(idRaw)
	digest := h.Sum(nil)
	require.Equal(t, digest, []byte(pkid), "PKID must be the SHA2-256 of peerIdentity")

	//  The PKI-ID is calculated by concatenating the MspId with IdBytes.
	// Ensure that additional fields haven't been introduced in the code
	v := reflect.Indirect(reflect.ValueOf(id)).Type()
	fieldsThatStartWithXXX := 0
	for i := 0; i < v.NumField(); i++ {
		if strings.Index(v.Field(i).Name, "XXX_") == 0 {
			fieldsThatStartWithXXX++
		}
	}
	require.Equal(t, 2+fieldsThatStartWithXXX, v.NumField())
}

func TestPKIidOfNil(t *testing.T) {
	signer := &mocks.SignerSerializer{}
	cryptoProvider, err := sw.NewDefaultSecurityLevelWithKeystore(sw.NewDummyKeyStore())
	require.NoError(t, err)
	localMSP := mgmt.GetLocalMSP(cryptoProvider)
	msgCryptoService := NewMCS(
		&mocks.ChannelPolicyManagerGetter{},
		signer,
		NewDeserializersManager(localMSP),
		cryptoProvider,
	)

	pkid := msgCryptoService.GetPKIidOfCert(nil)
	// Check pkid is not nil
	require.Nil(t, pkid, "PKID must be nil")
}

func TestValidateIdentity(t *testing.T) {
	deserializersManager := &mocks.DeserializersManager{
		LocalDeserializer: &mocks.IdentityDeserializer{Identity: []byte("Alice"), Msg: []byte("msg1"), Mock: mock.Mock{}},
		ChannelDeserializers: map[string]msp.IdentityDeserializer{
			"A": &mocks.IdentityDeserializer{Identity: []byte("Bob"), Msg: []byte("msg2"), Mock: mock.Mock{}},
		},
	}
	signer := &mocks.SignerSerializer{}
	signer.SerializeReturns([]byte("Charlie"), nil)
	cryptoProvider, err := sw.NewDefaultSecurityLevelWithKeystore(sw.NewDummyKeyStore())
	require.NoError(t, err)
	msgCryptoService := NewMCS(
		&mocks.ChannelPolicyManagerGetterWithManager{},
		signer,
		deserializersManager,
		cryptoProvider,
	)

	err = msgCryptoService.ValidateIdentity([]byte("Alice"))
	require.NoError(t, err)

	err = msgCryptoService.ValidateIdentity([]byte("Bob"))
	require.NoError(t, err)

	err = msgCryptoService.ValidateIdentity([]byte("Charlie"))
	require.Error(t, err)

	err = msgCryptoService.ValidateIdentity(nil)
	require.Error(t, err)

	// Now, pretend the identities are not well formed
	deserializersManager.ChannelDeserializers["A"].(*mocks.IdentityDeserializer).On("IsWellFormed", mock.Anything).Return(errors.New("invalid form"))
	err = msgCryptoService.ValidateIdentity([]byte("Bob"))
	require.Error(t, err)
	require.Equal(t, "identity is not well formed: invalid form", err.Error())

	deserializersManager.LocalDeserializer.(*mocks.IdentityDeserializer).On("IsWellFormed", mock.Anything).Return(errors.New("invalid form"))
	err = msgCryptoService.ValidateIdentity([]byte("Alice"))
	require.Error(t, err)
	require.Equal(t, "identity is not well formed: invalid form", err.Error())
}

func TestSign(t *testing.T) {
	signer := &mocks.SignerSerializer{}
	signer.SignReturns([]byte("signature"), nil)
	cryptoProvider, err := sw.NewDefaultSecurityLevelWithKeystore(sw.NewDummyKeyStore())
	require.NoError(t, err)

	localMSP := mgmt.GetLocalMSP(cryptoProvider)
	msgCryptoService := NewMCS(
		&mocks.ChannelPolicyManagerGetter{},
		signer,
		NewDeserializersManager(localMSP),
		cryptoProvider,
	)

	msg := []byte("Hello World!!!")
	sigma, err := msgCryptoService.Sign(msg)
	require.NoError(t, err, "Failed generating signature")
	require.NotNil(t, sigma, "Signature must be different from nil")
}

func TestVerify(t *testing.T) {
	signer := &mocks.SignerSerializer{}
	signer.SerializeReturns([]byte("Alice"), nil)
	signer.SignReturns([]byte("msg1"), nil)
	cryptoProvider, err := sw.NewDefaultSecurityLevelWithKeystore(sw.NewDummyKeyStore())
	require.NoError(t, err)
	msgCryptoService := NewMCS(
		&mocks.ChannelPolicyManagerGetterWithManager{
			Managers: map[string]policies.Manager{
				"A": &mocks.ChannelPolicyManager{
					Policy: &mocks.Policy{Deserializer: &mocks.IdentityDeserializer{Identity: []byte("Bob"), Msg: []byte("msg2"), Mock: mock.Mock{}}},
				},
				"B": &mocks.ChannelPolicyManager{
					Policy: &mocks.Policy{Deserializer: &mocks.IdentityDeserializer{Identity: []byte("Charlie"), Msg: []byte("msg3"), Mock: mock.Mock{}}},
				},
				"C": nil,
			},
		},
		signer,
		&mocks.DeserializersManager{
			LocalDeserializer: &mocks.IdentityDeserializer{Identity: []byte("Alice"), Msg: []byte("msg1"), Mock: mock.Mock{}},
			ChannelDeserializers: map[string]msp.IdentityDeserializer{
				"A": &mocks.IdentityDeserializer{Identity: []byte("Bob"), Msg: []byte("msg2"), Mock: mock.Mock{}},
				"B": &mocks.IdentityDeserializer{Identity: []byte("Charlie"), Msg: []byte("msg3"), Mock: mock.Mock{}},
				"C": &mocks.IdentityDeserializer{Identity: []byte("Dave"), Msg: []byte("msg4"), Mock: mock.Mock{}},
			},
		},
		cryptoProvider,
	)

	msg := []byte("msg1")
	sigma, err := msgCryptoService.Sign(msg)
	require.NoError(t, err, "Failed generating signature")

	err = msgCryptoService.Verify(api.PeerIdentityType("Alice"), sigma, msg)
	require.NoError(t, err, "Alice should verify the signature")

	err = msgCryptoService.Verify(api.PeerIdentityType("Bob"), sigma, msg)
	require.Error(t, err, "Bob should not verify the signature")

	err = msgCryptoService.Verify(api.PeerIdentityType("Charlie"), sigma, msg)
	require.Error(t, err, "Charlie should not verify the signature")

	sigma, err = msgCryptoService.Sign(msg)
	require.NoError(t, err)
	err = msgCryptoService.Verify(api.PeerIdentityType("Dave"), sigma, msg)
	require.Error(t, err)
	require.Contains(t, err.Error(), "Could not acquire policy manager")

	// Check invalid args
	require.Error(t, msgCryptoService.Verify(nil, sigma, msg))
}

func TestVerifyBlock(t *testing.T) {
	aliceSigner := &mocks.SignerSerializer{}
	aliceSigner.SerializeReturns([]byte("Alice"), nil)
	policyManagerGetter := &mocks.ChannelPolicyManagerGetterWithManager{
		Managers: map[string]policies.Manager{
			"A": &mocks.ChannelPolicyManager{
				Policy: &mocks.Policy{Deserializer: &mocks.IdentityDeserializer{Identity: []byte("Bob"), Msg: []byte("msg2"), Mock: mock.Mock{}}},
			},
			"B": &mocks.ChannelPolicyManager{
				Policy: &mocks.Policy{Deserializer: &mocks.IdentityDeserializer{Identity: []byte("Charlie"), Msg: []byte("msg3"), Mock: mock.Mock{}}},
			},
			"C": &mocks.ChannelPolicyManager{
				Policy: &mocks.Policy{Deserializer: &mocks.IdentityDeserializer{Identity: []byte("Alice"), Msg: []byte("msg1"), Mock: mock.Mock{}}},
			},
			"D": &mocks.ChannelPolicyManager{
				Policy: &mocks.Policy{Deserializer: &mocks.IdentityDeserializer{Identity: []byte("Alice"), Msg: []byte("msg1"), Mock: mock.Mock{}}},
			},
		},
	}

	cryptoProvider, err := sw.NewDefaultSecurityLevelWithKeystore(sw.NewDummyKeyStore())
	require.NoError(t, err)
	msgCryptoService := NewMCS(
		policyManagerGetter,
		aliceSigner,
		&mocks.DeserializersManager{
			LocalDeserializer: &mocks.IdentityDeserializer{Identity: []byte("Alice"), Msg: []byte("msg1"), Mock: mock.Mock{}},
			ChannelDeserializers: map[string]msp.IdentityDeserializer{
				"A": &mocks.IdentityDeserializer{Identity: []byte("Bob"), Msg: []byte("msg2"), Mock: mock.Mock{}},
				"B": &mocks.IdentityDeserializer{Identity: []byte("Charlie"), Msg: []byte("msg3"), Mock: mock.Mock{}},
			},
		},
		cryptoProvider,
	)

	// - Prepare testing valid block, Alice signs it.
	blockRaw, msg := mockBlock(t, "C", 42, aliceSigner, nil)
	policyManagerGetter.Managers["C"].(*mocks.ChannelPolicyManager).Policy.(*mocks.Policy).Deserializer.(*mocks.IdentityDeserializer).Msg = msg
	blockRaw2, msg2 := mockBlock(t, "D", 42, aliceSigner, nil)
	policyManagerGetter.Managers["D"].(*mocks.ChannelPolicyManager).Policy.(*mocks.Policy).Deserializer.(*mocks.IdentityDeserializer).Msg = msg2

	// - Verify block
	require.NoError(t, msgCryptoService.VerifyBlock([]byte("C"), 42, blockRaw))
	// Wrong sequence number claimed
	err = msgCryptoService.VerifyBlock([]byte("C"), 43, blockRaw)
	require.Error(t, err)
	require.Contains(t, err.Error(), "but actual seqNum inside block is")
	delete(policyManagerGetter.Managers, "D")
	nilPolMgrErr := msgCryptoService.VerifyBlock([]byte("D"), 42, blockRaw2)
	require.Contains(t, nilPolMgrErr.Error(), "Could not acquire policy manager")
	require.Error(t, nilPolMgrErr)
	require.Error(t, msgCryptoService.VerifyBlock([]byte("A"), 42, blockRaw))
	require.Error(t, msgCryptoService.VerifyBlock([]byte("B"), 42, blockRaw))

	// - Prepare testing invalid block (wrong data has), Alice signs it.
	blockRaw, msg = mockBlock(t, "C", 42, aliceSigner, []byte{0})
	policyManagerGetter.Managers["C"].(*mocks.ChannelPolicyManager).Policy.(*mocks.Policy).Deserializer.(*mocks.IdentityDeserializer).Msg = msg

	// - Verify block
	require.Error(t, msgCryptoService.VerifyBlock([]byte("C"), 42, blockRaw))

	// Check invalid args
	require.Error(t, msgCryptoService.VerifyBlock([]byte("C"), 42, &common.Block{}))
}

func mockBlock(t *testing.T, channel string, seqNum uint64, localSigner *mocks.SignerSerializer, dataHash []byte) (*common.Block, []byte) {
	block := protoutil.NewBlock(seqNum, nil)

	// Add a fake transaction to the block referring channel "C"
	sProp, _ := protoutil.MockSignedEndorserProposalOrPanic(channel, &protospeer.ChaincodeSpec{}, []byte("transactor"), []byte("transactor's signature"))
	sPropRaw, err := protoutil.Marshal(sProp)
	require.NoError(t, err, "Failed marshalling signed proposal")
	block.Data.Data = [][]byte{sPropRaw}

	// Compute hash of block.Data and put into the Header
	if len(dataHash) != 0 {
		block.Header.DataHash = dataHash
	} else {
		block.Header.DataHash = protoutil.BlockDataHash(block.Data)
	}

	// Add signer's signature to the block
	shdr, err := protoutil.NewSignatureHeader(localSigner)
	require.NoError(t, err, "Failed generating signature header")

	blockSignature := &common.MetadataSignature{
		SignatureHeader: protoutil.MarshalOrPanic(shdr),
	}

	// Note, this value is intentionally nil, as this metadata is only about the signature, there is no additional metadata
	// information required beyond the fact that the metadata item is signed.
	blockSignatureValue := []byte(nil)

	msg := util.ConcatenateBytes(blockSignatureValue, blockSignature.SignatureHeader, protoutil.BlockHeaderBytes(block.Header))
	localSigner.SignReturns(msg, nil)
	blockSignature.Signature, err = localSigner.Sign(msg)
	require.NoError(t, err, "Failed signing block")

	block.Metadata.Metadata[common.BlockMetadataIndex_SIGNATURES] = protoutil.MarshalOrPanic(&common.Metadata{
		Value: blockSignatureValue,
		Signatures: []*common.MetadataSignature{
			blockSignature,
		},
	})

	return block, msg
}

func TestExpiration(t *testing.T) {
	expirationDate := time.Now().Add(time.Minute)
	id1 := &pmsp.SerializedIdentity{
		Mspid:   "X509BasedMSP",
		IdBytes: []byte("X509BasedIdentity"),
	}

	x509IdentityBytes, _ := proto.Marshal(id1)

	id2 := &pmsp.SerializedIdentity{
		Mspid:   "nonX509BasedMSP",
		IdBytes: []byte("nonX509RawIdentity"),
	}

	nonX509IdentityBytes, _ := proto.Marshal(id2)

	deserializersManager := &mocks.DeserializersManager{
		LocalDeserializer: &mocks.IdentityDeserializer{
			Identity: []byte{1, 2, 3},
			Msg:      []byte{1, 2, 3},
		},
		ChannelDeserializers: map[string]msp.IdentityDeserializer{
			"X509BasedMSP": &mocks.IdentityDeserializerWithExpiration{
				Expiration: expirationDate,
				IdentityDeserializer: &mocks.IdentityDeserializer{
					Identity: x509IdentityBytes,
					Msg:      []byte("x509IdentityBytes"),
				},
			},
			"nonX509BasedMSP": &mocks.IdentityDeserializer{
				Identity: nonX509IdentityBytes,
				Msg:      []byte("nonX509IdentityBytes"),
			},
		},
	}
	cryptoProvider, err := sw.NewDefaultSecurityLevelWithKeystore(sw.NewDummyKeyStore())
	require.NoError(t, err)
	msgCryptoService := NewMCS(
		&mocks.ChannelPolicyManagerGetterWithManager{},
		&mocks.SignerSerializer{},
		deserializersManager,
		cryptoProvider,
	)

	// Green path I check the expiration date is as expected
	exp, err := msgCryptoService.Expiration(x509IdentityBytes)
	require.NoError(t, err)
	require.Equal(t, expirationDate.Second(), exp.Second())

	// Green path II - a non-x509 identity has a zero expiration time
	exp, err = msgCryptoService.Expiration(nonX509IdentityBytes)
	require.NoError(t, err)
	require.Zero(t, exp)

	// Bad path I - corrupt the x509 identity and make sure error is returned
	x509IdentityBytes = append(x509IdentityBytes, 0, 0, 0, 0, 0, 0)
	exp, err = msgCryptoService.Expiration(x509IdentityBytes)
	require.Error(t, err)
	require.Contains(t, err.Error(), "No MSP found able to do that")
	require.Zero(t, exp)
}
