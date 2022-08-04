/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package signer

import (
	"crypto/rand"
	"encoding/pem"
	"io/ioutil"
	"strings"

	"github.com/hxx258456/ccgo/x509"

	"github.com/hxx258456/ccgo/sm2"
	"github.com/hxx258456/fabric-protos-go-cc/msp"
	"github.com/hxx258456/fabric/bccsp/sw"
	"github.com/hxx258456/fabric/bccsp/utils"
	"github.com/hxx258456/fabric/common/util"
	"github.com/hxx258456/fabric/protoutil"
	"github.com/pkg/errors"
)

/*
cmd/common/signer/signer.go 貌似是MSP的签名器
*/

// Config holds the configuration for
// creation of a Signer
type Config struct {
	MSPID string
	// 证书, 当前msp的证书?
	IdentityPath string
	// 私钥 当前msp的私钥?
	KeyPath string
}

// Signer signs messages.
// TODO: Ideally we'd use an MSP to be agnostic, but since it's impossible to
// initialize an MSP without a CA cert that signs the signing identity,
// this will do for now.
type Signer struct {
	key     *sm2.PrivateKey
	Creator []byte
}

func (si *Signer) Serialize() ([]byte, error) {
	return si.Creator, nil
}

// NewSigner creates a new Signer out of the given configuration
func NewSigner(conf Config) (*Signer, error) {
	// 序列化客户端mspid
	sId, err := serializeIdentity(conf.IdentityPath, conf.MSPID)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// 读取私钥
	key, err := loadPrivateKey(conf.KeyPath)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// 创建 Signer TODO: 在 Signer埋入 序列化客户端mspid 的目的是啥？
	return &Signer{
		Creator: sId,
		key:     key,
	}, nil
}

// 将客户端证书及MSPID组装为 protof格式的字节流
func serializeIdentity(clientCert string, mspID string) ([]byte, error) {
	b, err := ioutil.ReadFile(clientCert)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if err := validateEnrollmentCertificate(b); err != nil {
		return nil, err
	}
	sId := &msp.SerializedIdentity{
		Mspid:   mspID,
		IdBytes: b,
	}
	return protoutil.MarshalOrPanic(sId), nil
}

func validateEnrollmentCertificate(b []byte) error {
	bl, _ := pem.Decode(b)
	if bl == nil {
		return errors.Errorf("enrollment certificate isn't a valid PEM block")
	}

	if bl.Type != "CERTIFICATE" {
		return errors.Errorf("enrollment certificate should be a certificate, got a %s instead", strings.ToLower(bl.Type))
	}

	if _, err := x509.ParseCertificate(bl.Bytes); err != nil {
		return errors.Errorf("enrollment certificate is not a valid x509 certificate: %v", err)
	}
	return nil
}

func (si *Signer) Sign(msg []byte) ([]byte, error) {
	digest := util.ComputeSM3(msg)
	return signSM2(si.key, digest)
}

// 将未加密的pem文件转为sm2私钥
func loadPrivateKey(file string) (*sm2.PrivateKey, error) {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// bl, _ := pem.Decode(b)
	// if bl == nil {
	// 	return nil, errors.Errorf("failed to decode PEM block from %s", file)
	// }
	key, err := utils.PEMToSm2PrivateKey(b, nil)
	if err != nil {
		return nil, err
	}
	return key, nil
}

func signSM2(k *sm2.PrivateKey, digest []byte) (signature []byte, err error) {
	r, s, err := sm2.Sm2Sign(k, digest, nil, rand.Reader)
	if err != nil {
		return nil, err
	}
	//s, err = utils.ToLowS(&k.PublicKey, s)
	if err != nil {
		return nil, err
	}

	return sw.MarshalSM2Signature(r, s)
}

// // Based on crypto/tls/tls.go but modified for Fabric:
// func parsePrivateKey(der []byte) (crypto.PrivateKey, error) {
// 	// OpenSSL 1.0.0 generates PKCS#8 keys.
// 	if key, err := x509.ParsePKCS8PrivateKey(der); err == nil {
// 		switch key := key.(type) {
// 		// Fabric only supports ECDSA at the moment.
// 		case *ecdsa.PrivateKey:
// 			return key, nil
// 		default:
// 			return nil, errors.Errorf("found unknown private key type (%T) in PKCS#8 wrapping", key)
// 		}
// 	}

// 	// OpenSSL ecparam generates SEC1 EC private keys for ECDSA.
// 	key, err := x509.ParseECPrivateKey(der)
// 	if err != nil {
// 		return nil, errors.Errorf("failed to parse private key: %v", err)
// 	}

// 	return key, nil
// }

// func signECDSA(k *ecdsa.PrivateKey, digest []byte) (signature []byte, err error) {
// 	r, s, err := ecdsa.Sign(rand.Reader, k, digest)
// 	if err != nil {
// 		return nil, err
// 	}

// 	s, err = utils.ToLowS(&k.PublicKey, s)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return marshalECDSASignature(r, s)
// }

// func marshalECDSASignature(r, s *big.Int) ([]byte, error) {
// 	return asn1.Marshal(ECDSASignature{r, s})
// }

// type ECDSASignature struct {
// 	R, S *big.Int
// }
