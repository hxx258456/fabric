/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/
package bridge

import (
	"github.com/hxx258456/ccgo/sm2"

	"github.com/golang/protobuf/proto"
	"github.com/hxx258456/fabric/bccsp"
	cryptolib "github.com/hxx258456/fabric/idemix"
	"github.com/hyperledger/fabric-amcl/amcl/FP256BN"
	"github.com/pkg/errors"
)

// Revocation encapsulates the idemix algorithms for revocation
type Revocation struct {
}

// NewKey generate a new revocation key-pair.
func (*Revocation) NewKey() (*sm2.PrivateKey, error) {
	return cryptolib.GenerateLongTermRevocationKey()
}

// Sign generates a new CRI with the respect to the passed unrevoked handles, epoch, and revocation algorithm.
func (*Revocation) Sign(key *sm2.PrivateKey, unrevokedHandles [][]byte, epoch int, alg bccsp.RevocationAlgorithm) (res []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			res = nil
			err = errors.Errorf("failure [%s]", r)
		}
	}()

	handles := make([]*FP256BN.BIG, len(unrevokedHandles))
	for i := 0; i < len(unrevokedHandles); i++ {
		handles[i] = FP256BN.FromBytes(unrevokedHandles[i])
	}
	cri, err := cryptolib.CreateCRI(key, handles, epoch, cryptolib.RevocationAlgorithm(alg), NewRandOrPanic())
	if err != nil {
		return nil, errors.WithMessage(err, "failed creating CRI")
	}

	return proto.Marshal(cri)
}

// Verify checks that the passed serialised CRI (criRaw) is valid with the respect to the passed revocation public key,
// epoch, and revocation algorithm.
func (*Revocation) Verify(pk *sm2.PublicKey, criRaw []byte, epoch int, alg bccsp.RevocationAlgorithm) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.Errorf("failure [%s]", r)
		}
	}()

	cri := &cryptolib.CredentialRevocationInformation{}
	err = proto.Unmarshal(criRaw, cri)
	if err != nil {
		return err
	}

	return cryptolib.VerifyEpochPK(
		pk,
		cri.EpochPk,
		cri.EpochPkSig,
		int(cri.Epoch),
		cryptolib.RevocationAlgorithm(cri.RevocationAlg),
	)
}
