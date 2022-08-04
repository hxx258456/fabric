/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package nwo

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/pem"
	"fmt"
	"io/ioutil"

	"github.com/hxx258456/ccgo/sm2"
	"github.com/hxx258456/ccgo/x509"

	"github.com/golang/protobuf/proto"
	"github.com/hxx258456/fabric-protos-go-cc/msp"
	"github.com/hxx258456/fabric/bccsp/sw"
)

// A SigningIdentity represents an MSP signing identity.
type SigningIdentity struct {
	CertPath string
	KeyPath  string
	MSPID    string
}

// Serialize returns the probobuf encoding of an msp.SerializedIdenity.
func (s *SigningIdentity) Serialize() ([]byte, error) {
	cert, err := ioutil.ReadFile(s.CertPath)
	if err != nil {
		return nil, err
	}
	return proto.Marshal(&msp.SerializedIdentity{
		Mspid:   s.MSPID,
		IdBytes: cert,
	})
}

// Sign computes a SHA256 message digest, signs it with the associated private
// key, and returns the signature after low-S normlization.
func (s *SigningIdentity) Sign(msg []byte) ([]byte, error) {
	digest := sha256.Sum256(msg)
	pemKey, err := ioutil.ReadFile(s.KeyPath)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(pemKey)
	if block.Type != "EC PRIVATE KEY" && block.Type != "PRIVATE KEY" {
		return nil, fmt.Errorf("file %s does not contain a private key", s.KeyPath)
	}
	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	eckey, ok := key.(*sm2.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("unexpected key type: %T", key)
	}
	r, _s, err := sm2.Sign(rand.Reader, eckey, digest[:])
	if err != nil {
		return nil, err
	}
	// 可能有bug
	return sw.MarshalSM2Signature(r, _s)
}
