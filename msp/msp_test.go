/*
Copyright IBM Corp. 2017 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0

*/

package msp

import (
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509/pkix"
	"encoding/asn1"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"reflect"
	"testing"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/hxx258456/ccgo/sm2"
	"github.com/hxx258456/ccgo/sm3"
	"github.com/hxx258456/ccgo/x509"
	"github.com/hxx258456/fabric-protos-go-cc/msp"
	"github.com/hxx258456/fabric/bccsp"
	"github.com/hxx258456/fabric/bccsp/factory"
	"github.com/hxx258456/fabric/bccsp/sw"
	"github.com/hxx258456/fabric/core/config/configtest"
	"github.com/hxx258456/fabric/protoutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var notACert = `-----BEGIN X509 CRL-----
MIIBYzCCAQgCAQEwCgYIKoZIzj0EAwIwfzELMAkGA1UEBhMCVVMxEzARBgNVBAgT
CkNhbGlmb3JuaWExFjAUBgNVBAcTDVNhbiBGcmFuY2lzY28xHzAdBgNVBAoTFklu
dGVybmV0IFdpZGdldHMsIEluYy4xDDAKBgNVBAsTA1dXVzEUMBIGA1UEAxMLZXhh
bXBsZS5jb20XDTE3MDEyMzIwNTYyMFoXDTE3MDEyNjIwNTYyMFowJzAlAhQERXCx
LHROap1vM3CV40EHOghPTBcNMTcwMTIzMjA0NzMxWqAvMC0wHwYDVR0jBBgwFoAU
F2dCPaqegj/ExR2fW8OZ0bWcSBAwCgYDVR0UBAMCAQgwCgYIKoZIzj0EAwIDSQAw
RgIhAOTTpQYkGO+gwVe1LQOcNMD5fzFViOwBUraMrk6dRMlmAiEA8z2dpXKGwHrj
FRBbKkDnSpaVcZgjns+mLdHV2JkF0gk=
-----END X509 CRL-----`

func TestMSPParsers(t *testing.T) {
	_, _, err := localMsp.(*bccspmsp).getIdentityFromConf(nil)
	assert.Error(t, err)
	_, _, err = localMsp.(*bccspmsp).getIdentityFromConf([]byte("barf"))
	assert.Error(t, err)
	_, _, err = localMsp.(*bccspmsp).getIdentityFromConf([]byte(notACert))
	assert.Error(t, err)

	_, err = localMsp.(*bccspmsp).getSigningIdentityFromConf(nil)
	assert.Error(t, err)

	sigid := &msp.SigningIdentityInfo{PublicSigner: []byte("barf"), PrivateSigner: nil}
	_, err = localMsp.(*bccspmsp).getSigningIdentityFromConf(sigid)
	assert.Error(t, err)

	keyinfo := &msp.KeyInfo{KeyIdentifier: "PEER", KeyMaterial: nil}
	sigid = &msp.SigningIdentityInfo{PublicSigner: []byte("barf"), PrivateSigner: keyinfo}
	_, err = localMsp.(*bccspmsp).getSigningIdentityFromConf(sigid)
	assert.Error(t, err)
}

func TestGetSigningIdentityFromConfWithWrongPrivateCert(t *testing.T) {
	// Temporary Replace root certs
	oldRoots := localMsp.(*bccspmsp).opts.Roots
	defer func() {
		// Restore original root certs
		localMsp.(*bccspmsp).opts.Roots = oldRoots
	}()
	_, cert := generateSelfSignedCert(t, time.Now())
	localMsp.(*bccspmsp).opts.Roots = x509.NewCertPool()
	localMsp.(*bccspmsp).opts.Roots.AddCert(cert)

	// Use self signed cert as public key. Convert DER to PEM format
	pem := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: cert.Raw})

	// Use wrong formatted private cert
	keyinfo := &msp.KeyInfo{
		KeyMaterial:   []byte("wrong encoding"),
		KeyIdentifier: "MyPrivateKey",
	}
	sigid := &msp.SigningIdentityInfo{PublicSigner: pem, PrivateSigner: keyinfo}
	_, err := localMsp.(*bccspmsp).getSigningIdentityFromConf(sigid)
	assert.EqualError(t, err, "MyPrivateKey: wrong PEM encoding")
}

func TestMSPSetupNoCryptoConf(t *testing.T) {
	mspDir := configtest.GetDevMspDir()
	conf, err := GetLocalMspConfig(mspDir, nil, "SampleOrg")
	if err != nil {
		fmt.Printf("Setup should have succeeded, got err %s instead", err)
		os.Exit(-1)
	}

	mspconf := &msp.FabricMSPConfig{}
	err = proto.Unmarshal(conf.Config, mspconf)
	assert.NoError(t, err)

	// here we test the case of an MSP configuration
	// where the hash function to be used to obtain
	// the identity identifier is unspecified - a
	// sane default should be picked
	mspconf.CryptoConfig.IdentityIdentifierHashFunction = ""
	b, err := proto.Marshal(mspconf)
	assert.NoError(t, err)
	conf.Config = b
	newmsp, err := newBccspMsp(MSPv1_0, factory.GetDefault())
	assert.NoError(t, err)
	err = newmsp.Setup(conf)
	assert.NoError(t, err)

	// here we test the case of an MSP configuration
	// where the hash function to be used to compute
	// signatures is unspecified - a sane default
	// should be picked
	mspconf.CryptoConfig.SignatureHashFamily = ""
	b, err = proto.Marshal(mspconf)
	assert.NoError(t, err)
	conf.Config = b
	newmsp, err = newBccspMsp(MSPv1_0, factory.GetDefault())
	assert.NoError(t, err)
	err = newmsp.Setup(conf)
	assert.NoError(t, err)

	// here we test the case of an MSP configuration
	// that has NO crypto configuration specified;
	// the code will use appropriate defaults
	mspconf.CryptoConfig = nil
	b, err = proto.Marshal(mspconf)
	assert.NoError(t, err)
	conf.Config = b
	newmsp, err = newBccspMsp(MSPv1_0, factory.GetDefault())
	assert.NoError(t, err)
	err = newmsp.Setup(conf)
	assert.NoError(t, err)
}

func TestGetters(t *testing.T) {
	typ := localMsp.GetType()
	assert.Equal(t, typ, FABRIC)
	assert.NotNil(t, localMsp.GetTLSRootCerts())
	assert.NotNil(t, localMsp.GetTLSIntermediateCerts())
}

func TestMSPSetupBad(t *testing.T) {
	_, err := GetLocalMspConfig("barf", nil, "SampleOrg")
	if err == nil {
		t.Fatalf("Setup should have failed on an invalid config file")
		return
	}

	mgr := NewMSPManager()
	err = mgr.Setup(nil)
	assert.NoError(t, err)
	err = mgr.Setup([]MSP{})
	assert.NoError(t, err)
}

func TestDoubleSetup(t *testing.T) {
	// note that we've already called setup once on this
	err := mspMgr.Setup(nil)
	assert.NoError(t, err)
}

type bccspNoKeyLookupKS struct {
	bccsp.BCCSP
}

func (*bccspNoKeyLookupKS) GetKey(ski []byte) (k bccsp.Key, err error) {
	return nil, errors.New("not found")
}

func TestNotFoundInBCCSP(t *testing.T) {
	cryptoProvider, err := sw.NewDefaultSecurityLevelWithKeystore(sw.NewDummyKeyStore())
	assert.NoError(t, err)

	dir := configtest.GetDevMspDir()
	conf, err := GetLocalMspConfig(dir, nil, "SampleOrg")

	assert.NoError(t, err)

	thisMSP, err := newBccspMsp(MSPv1_0, cryptoProvider)
	assert.NoError(t, err)
	ks, err := sw.NewFileBasedKeyStore(nil, filepath.Join(dir, "keystore"), true)
	assert.NoError(t, err)
	csp, err := sw.NewWithParams(true, 256, "SHA2", ks)
	assert.NoError(t, err)
	thisMSP.(*bccspmsp).bccsp = &bccspNoKeyLookupKS{csp}

	err = thisMSP.Setup(conf)
	assert.Error(t, err)
	assert.Contains(t, "KeyMaterial not found in SigningIdentityInfo", err.Error())
}

func TestGetIdentities(t *testing.T) {
	_, err := localMsp.GetDefaultSigningIdentity()
	if err != nil {
		t.Fatalf("GetDefaultSigningIdentity failed with err %s", err)
		return
	}
}

func TestDeserializeIdentityFails(t *testing.T) {
	_, err := localMsp.DeserializeIdentity([]byte("barf"))
	assert.Error(t, err)

	id := &msp.SerializedIdentity{Mspid: "SampleOrg", IdBytes: []byte("barfr")}
	b, err := proto.Marshal(id)
	assert.NoError(t, err)
	_, err = localMsp.DeserializeIdentity(b)
	assert.Error(t, err)

	id = &msp.SerializedIdentity{Mspid: "SampleOrg", IdBytes: []byte(notACert)}
	b, err = proto.Marshal(id)
	assert.NoError(t, err)
	_, err = localMsp.DeserializeIdentity(b)
	assert.Error(t, err)
}

func TestGetSigningIdentityFromVerifyingMSP(t *testing.T) {
	cryptoProvider, err := sw.NewDefaultSecurityLevelWithKeystore(sw.NewDummyKeyStore())
	assert.NoError(t, err)

	mspDir := configtest.GetDevMspDir()
	conf, err = GetVerifyingMspConfig(mspDir, "SampleOrg", ProviderTypeToString(FABRIC))
	if err != nil {
		fmt.Printf("Setup should have succeeded, got err %s instead", err)
		os.Exit(-1)
	}

	newmsp, err := newBccspMsp(MSPv1_0, cryptoProvider)
	assert.NoError(t, err)
	err = newmsp.Setup(conf)
	assert.NoError(t, err)

	_, err = newmsp.GetDefaultSigningIdentity()
	assert.Error(t, err)
	_, err = newmsp.GetSigningIdentity(nil)
	assert.Error(t, err)
}

func TestValidateDefaultSigningIdentity(t *testing.T) {
	id, err := localMsp.GetDefaultSigningIdentity()
	assert.NoError(t, err)

	err = localMsp.Validate(id.GetPublicVersion())
	assert.NoError(t, err)
}

func TestSerializeIdentities(t *testing.T) {
	id, err := localMsp.GetDefaultSigningIdentity()
	if err != nil {
		t.Fatalf("GetSigningIdentity should have succeeded, got err %s", err)
		return
	}

	serializedID, err := id.Serialize()
	if err != nil {
		t.Fatalf("Serialize should have succeeded, got err %s", err)
		return
	}

	idBack, err := localMsp.DeserializeIdentity(serializedID)
	if err != nil {
		t.Fatalf("DeserializeIdentity should have succeeded, got err %s", err)
		return
	}

	err = localMsp.Validate(idBack)
	if err != nil {
		t.Fatalf("The identity should be valid, got err %s", err)
		return
	}

	if !reflect.DeepEqual(id.GetPublicVersion(), idBack) {
		t.Fatalf("Identities should be equal (%s) (%s)", id, idBack)
		return
	}
}

// 国密改造后散列算法改为SM3
func computeSKI(key *sm2.PublicKey) []byte {
	raw := elliptic.Marshal(key.Curve, key.X, key.Y)
	hash := sm3.Sm3Sum(raw)
	return hash[:]
}

func TestValidHostname(t *testing.T) {
	tests := []struct {
		name  string
		valid bool
	}{
		{"", false},
		{".", false},
		{"example.com", true},
		{"example.com.", true},
		{"*.example.com", true},
		{".example.com", false},
		{"host.*.example.com", false},
		{"localhost", true},
		{"-localhost", false},
		{"Not_Quite.example.com", true},
		{"weird:colon.example.com", true},
		{"1-2-3.example.com", true},
	}
	for _, tt := range tests {
		if tt.valid {
			require.True(t, validHostname(tt.name), "expected %s to be a valid hostname", tt.name)
		} else {
			require.False(t, validHostname(tt.name), "expected %s to be an invalid hostname", tt.name)
		}
	}
}

func TestValidateCANameConstraintsMitigation(t *testing.T) {
	// Prior to Go 1.15, if a signing certificate contains a name constraint, the
	// leaf certificate does not include a SAN, and the leaf common name looks
	// like a valid hostname, the certificate chain would fail to validate.
	// (This behavior may have been introduced with Go 1.10.)
	//
	// In Go 1.15, the behavior has changed and, by default, the same structure
	// will validate. This test asserts on the old behavior.
	caKey, err := sm2.GenerateKey(rand.Reader)
	require.NoError(t, err)
	leafKey, err := sm2.GenerateKey(rand.Reader)
	require.NoError(t, err)

	caKeyUsage := x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment | x509.KeyUsageCertSign | x509.KeyUsageCRLSign
	caTemplate := x509.Certificate{
		Subject:      pkix.Name{CommonName: "TestCA"},
		SerialNumber: big.NewInt(1),
		NotBefore:    time.Now().Add(-1 * time.Hour),
		NotAfter:     time.Now().Add(2 * time.Hour),
		// TODO gmgo的x509没有ExcludedDNSDomains
		// ExcludedDNSDomains:          []string{"example.com"},
		PermittedDNSDomainsCritical: true,
		IsCA:                        true,
		BasicConstraintsValid:       true,
		KeyUsage:                    caKeyUsage,
		SubjectKeyId:                computeSKI(caKey.Public().(*sm2.PublicKey)),
	}
	caCertBytes, err := x509.CreateCertificate(rand.Reader, &caTemplate, &caTemplate, caKey.Public(), caKey)
	require.NoError(t, err)
	ca, err := x509.ParseCertificate(caCertBytes)
	require.NoError(t, err)

	leafTemplate := x509.Certificate{
		Subject:      pkix.Name{CommonName: "localhost"},
		SerialNumber: big.NewInt(2),
		NotBefore:    time.Now().Add(-1 * time.Hour),
		NotAfter:     time.Now().Add(2 * time.Hour),
		KeyUsage:     x509.KeyUsageDigitalSignature,
		SubjectKeyId: computeSKI(leafKey.Public().(*sm2.PublicKey)),
	}
	leafCertBytes, err := x509.CreateCertificate(rand.Reader, &leafTemplate, ca, leafKey.Public(), caKey)
	require.NoError(t, err)

	keyBytes, err := x509.MarshalPKCS8PrivateKey(leafKey)
	require.NoError(t, err)

	caCertPem := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: caCertBytes})
	leafCertPem := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: leafCertBytes})
	leafKeyPem := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: keyBytes})

	t.Run("VerifyNameConstraintsSingleCert", func(t *testing.T) {
		for _, der := range [][]byte{caCertBytes, leafCertBytes} {
			cert, err := x509.ParseCertificate(der)
			require.NoError(t, err, "failed to parse certificate")

			err = verifyLegacyNameConstraints([]*x509.Certificate{cert})
			require.NoError(t, err, "single certificate should not trigger legacy constraints")
		}
	})

	t.Run("VerifyNameConstraints", func(t *testing.T) {
		var certs []*x509.Certificate
		for _, der := range [][]byte{leafCertBytes, caCertBytes} {
			cert, err := x509.ParseCertificate(der)
			require.NoError(t, err, "failed to parse certificate")
			certs = append(certs, cert)
		}

		err = verifyLegacyNameConstraints(certs)
		require.Error(t, err, "certificate chain should trigger legacy constraints")
		var cie x509.CertificateInvalidError
		require.True(t, errors.As(err, &cie))
		require.Equal(t, x509.NameConstraintsWithoutSANs, cie.Reason)
	})

	t.Run("VerifyNameConstraintsWithSAN", func(t *testing.T) {
		caCert, err := x509.ParseCertificate(caCertBytes)
		require.NoError(t, err)

		leafTemplate := leafTemplate
		leafTemplate.DNSNames = []string{"localhost"}

		leafCertBytes, err := x509.CreateCertificate(rand.Reader, &leafTemplate, caCert, leafKey.Public(), caKey)
		require.NoError(t, err)

		leafCert, err := x509.ParseCertificate(leafCertBytes)
		require.NoError(t, err)

		err = verifyLegacyNameConstraints([]*x509.Certificate{leafCert, caCert})
		require.NoError(t, err, "signer with name constraints and leaf with SANs should be valid")
	})

	t.Run("ValidationAtSetup", func(t *testing.T) {
		fabricMSPConfig := &msp.FabricMSPConfig{
			Name:      "ConstraintsMSP",
			RootCerts: [][]byte{caCertPem},
			SigningIdentity: &msp.SigningIdentityInfo{
				PublicSigner: leafCertPem,
				PrivateSigner: &msp.KeyInfo{
					KeyIdentifier: "Certificate Without SAN",
					KeyMaterial:   leafKeyPem,
				},
			},
		}
		mspConfig := &msp.MSPConfig{
			Config: protoutil.MarshalOrPanic(fabricMSPConfig),
		}

		ks, err := sw.NewFileBasedKeyStore(nil, filepath.Join(configtest.GetDevMspDir(), "keystore"), true)
		require.NoError(t, err)
		cryptoProvider, err := sw.NewDefaultSecurityLevelWithKeystore(ks)
		require.NoError(t, err)

		testMSP, err := NewBccspMspWithKeyStore(MSPv1_0, ks, cryptoProvider)
		require.NoError(t, err)

		err = testMSP.Setup(mspConfig)
		require.Error(t, err)
		var cie x509.CertificateInvalidError
		require.True(t, errors.As(err, &cie))
		require.Equal(t, x509.NameConstraintsWithoutSANs, cie.Reason)
	})
}

func TestIsWellFormed(t *testing.T) {
	mspMgr := NewMSPManager()

	id, err := localMsp.GetDefaultSigningIdentity()
	if err != nil {
		t.Fatalf("GetSigningIdentity should have succeeded, got err %s", err)
		return
	}

	serializedID, err := id.Serialize()
	if err != nil {
		t.Fatalf("Serialize should have succeeded, got err %s", err)
		return
	}

	sId := &msp.SerializedIdentity{}
	err = proto.Unmarshal(serializedID, sId)
	assert.NoError(t, err)

	// An MSP Manager without any MSPs should not recognize the identity since
	// not providers are registered
	err = mspMgr.IsWellFormed(sId)
	assert.Error(t, err)
	assert.Equal(t, "no MSP provider recognizes the identity", err.Error())

	// Add the MSP to the MSP Manager
	mspMgr.Setup([]MSP{localMsp})

	err = localMsp.IsWellFormed(sId)
	assert.NoError(t, err)
	err = mspMgr.IsWellFormed(sId)
	assert.NoError(t, err)

	bl, _ := pem.Decode(sId.IdBytes)
	assert.Equal(t, "CERTIFICATE", bl.Type)

	// Now, strip off the type from the PEM block. It should still be valid
	bl.Type = ""
	sId.IdBytes = pem.EncodeToMemory(bl)

	err = localMsp.IsWellFormed(sId)
	assert.NoError(t, err)

	// Now, corrupt the type of the PEM block.
	// make sure it isn't considered well formed by both an MSP and an MSP Manager
	bl.Type = "foo"
	sId.IdBytes = pem.EncodeToMemory(bl)
	err = localMsp.IsWellFormed(sId)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "pem type is")
	assert.Contains(t, err.Error(), "should be 'CERTIFICATE' or missing")

	err = mspMgr.IsWellFormed(sId)
	assert.Error(t, err)
	assert.Equal(t, "no MSP provider recognizes the identity", err.Error())

	// Restore the identity to what it was
	sId = &msp.SerializedIdentity{}
	err = proto.Unmarshal(serializedID, sId)
	assert.NoError(t, err)

	// Append some trailing junk at the end
	sId.IdBytes = append(sId.IdBytes, []byte{1, 2, 3}...)
	// And ensure it is deemed invalid
	err = localMsp.IsWellFormed(sId)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "for MSP SampleOrg has trailing bytes")

	// Parse the certificate of the identity
	cert, err := x509.ParseCertificate(bl.Bytes)
	assert.NoError(t, err)
	// Obtain the ECDSA signature
	// r, s, err := utils.UnmarshalECDSASignature(cert.Signature)
	r, s, err := sw.UnmarshalSM2Signature(cert.Signature)
	assert.NoError(t, err)

	// Modify it by appending some bytes to its end.
	modifiedSig, err := asn1.Marshal(rst{
		R: r,
		S: s,
		T: big.NewInt(100),
	})
	assert.NoError(t, err)
	newCert, err := certFromSM2Cert(cert)
	assert.NoError(t, err)
	newCert.SignatureValue.Bytes = modifiedSig
	newCert.SignatureValue.BitLength = len(newCert.SignatureValue.Bytes) * 8
	newCert.Raw = nil
	// Pour it back into the identity
	rawCert, err := asn1.Marshal(newCert)
	assert.NoError(t, err)
	sId.IdBytes = pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: rawCert})
	// Ensure it is invalid now and the signature modification is detected
	err = localMsp.IsWellFormed(sId)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "for MSP SampleOrg has a non canonical signature")
}

type rst struct {
	R, S, T *big.Int
}

func TestValidateCAIdentity(t *testing.T) {
	caID := getIdentity(t, cacerts)

	err := localMsp.Validate(caID)
	assert.Error(t, err)
}

func TestBadAdminIdentity(t *testing.T) {
	cryptoProvider, err := sw.NewDefaultSecurityLevelWithKeystore(sw.NewDummyKeyStore())
	assert.NoError(t, err)

	conf, err := GetLocalMspConfig("testdata/badadmin", nil, "SampleOrg")
	assert.NoError(t, err)

	thisMSP, err := newBccspMsp(MSPv1_0, cryptoProvider)
	assert.NoError(t, err)
	ks, err := sw.NewFileBasedKeyStore(nil, filepath.Join("testdata/badadmin", "keystore"), true)
	assert.NoError(t, err)
	csp, err := sw.NewWithParams(true, 256, "SHA2", ks)
	assert.NoError(t, err)
	thisMSP.(*bccspmsp).bccsp = csp

	err = thisMSP.Setup(conf)
	assert.Error(t, err)
}

func TestValidateAdminIdentity(t *testing.T) {
	caID := getIdentity(t, admincerts)

	err := localMsp.Validate(caID)
	assert.NoError(t, err)
}

func TestSerializeIdentitiesWithWrongMSP(t *testing.T) {
	id, err := localMsp.GetDefaultSigningIdentity()
	if err != nil {
		t.Fatalf("GetSigningIdentity should have succeeded, got err %s", err)
		return
	}

	serializedID, err := id.Serialize()
	if err != nil {
		t.Fatalf("Serialize should have succeeded, got err %s", err)
		return
	}

	sid := &msp.SerializedIdentity{}
	err = proto.Unmarshal(serializedID, sid)
	assert.NoError(t, err)

	sid.Mspid += "BARF"

	serializedID, err = proto.Marshal(sid)
	assert.NoError(t, err)

	_, err = localMsp.DeserializeIdentity(serializedID)
	assert.Error(t, err)
}

func TestSerializeIdentitiesWithMSPManager(t *testing.T) {
	id, err := localMsp.GetDefaultSigningIdentity()
	if err != nil {
		t.Fatalf("GetSigningIdentity should have succeeded, got err %s", err)
		return
	}

	serializedID, err := id.Serialize()
	if err != nil {
		t.Fatalf("Serialize should have succeeded, got err %s", err)
		return
	}

	_, err = mspMgr.DeserializeIdentity(serializedID)
	assert.NoError(t, err)

	sid := &msp.SerializedIdentity{}
	err = proto.Unmarshal(serializedID, sid)
	assert.NoError(t, err)

	sid.Mspid += "BARF"

	serializedID, err = proto.Marshal(sid)
	assert.NoError(t, err)

	_, err = mspMgr.DeserializeIdentity(serializedID)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), fmt.Sprintf("MSP %s is not defined on channel", sid.Mspid))

	_, err = mspMgr.DeserializeIdentity([]byte("barf"))
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "could not deserialize")
}

func TestIdentitiesGetters(t *testing.T) {
	id, err := localMsp.GetDefaultSigningIdentity()
	if err != nil {
		t.Fatalf("GetSigningIdentity should have succeeded, got err %s", err)
		return
	}

	idid := id.GetIdentifier()
	assert.NotNil(t, idid)
	mspid := id.GetMSPIdentifier()
	assert.NotNil(t, mspid)
	assert.False(t, id.Anonymous())
}

func TestSignAndVerify(t *testing.T) {
	id, err := localMsp.GetDefaultSigningIdentity()
	if err != nil {
		t.Fatalf("GetSigningIdentity should have succeeded")
		return
	}

	serializedID, err := id.Serialize()
	if err != nil {
		t.Fatalf("Serialize should have succeeded")
		return
	}

	idBack, err := localMsp.DeserializeIdentity(serializedID)
	if err != nil {
		t.Fatalf("DeserializeIdentity should have succeeded")
		return
	}

	msg := []byte("foo")
	sig, err := id.Sign(msg)
	if err != nil {
		t.Fatalf("Sign should have succeeded")
		return
	}

	err = id.Verify(msg, sig)
	if err != nil {
		t.Fatalf("The signature should be valid")
		return
	}

	err = idBack.Verify(msg, sig)
	if err != nil {
		t.Fatalf("The signature should be valid")
		return
	}

	err = id.Verify(msg[1:], sig)
	assert.Error(t, err)
	err = id.Verify(msg, sig[1:])
	assert.Error(t, err)
}

func TestSignAndVerifyFailures(t *testing.T) {
	msg := []byte("foo")

	id, err := localMspBad.GetDefaultSigningIdentity()
	if err != nil {
		t.Fatalf("GetSigningIdentity should have succeeded")
		return
	}

	hash := id.(*signingidentity).msp.cryptoConfig.SignatureHashFamily
	id.(*signingidentity).msp.cryptoConfig.SignatureHashFamily = "barf"

	_, err = id.Sign(msg)
	assert.Error(t, err)

	id.(*signingidentity).msp.cryptoConfig.SignatureHashFamily = hash

	sig, err := id.Sign(msg)
	if err != nil {
		t.Fatalf("Sign should have succeeded")
		return
	}

	id.(*signingidentity).msp.cryptoConfig.SignatureHashFamily = "barf"

	err = id.Verify(msg, sig)
	assert.Error(t, err)

	id.(*signingidentity).msp.cryptoConfig.SignatureHashFamily = hash
}

func TestSignAndVerifyOtherHash(t *testing.T) {
	id, err := localMsp.GetDefaultSigningIdentity()
	if err != nil {
		t.Fatalf("GetSigningIdentity should have succeeded")
		return
	}

	hash := id.(*signingidentity).msp.cryptoConfig.SignatureHashFamily
	id.(*signingidentity).msp.cryptoConfig.SignatureHashFamily = bccsp.SM3

	msg := []byte("foo")
	sig, err := id.Sign(msg)
	if err != nil {
		t.Fatalf("Sign should have succeeded")
		return
	}

	err = id.Verify(msg, sig)
	assert.NoError(t, err)

	id.(*signingidentity).msp.cryptoConfig.SignatureHashFamily = hash
}

func TestSignAndVerify_longMessage(t *testing.T) {
	id, err := localMsp.GetDefaultSigningIdentity()
	if err != nil {
		t.Fatalf("GetSigningIdentity should have succeeded")
		return
	}

	serializedID, err := id.Serialize()
	if err != nil {
		t.Fatalf("Serialize should have succeeded")
		return
	}

	idBack, err := localMsp.DeserializeIdentity(serializedID)
	if err != nil {
		t.Fatalf("DeserializeIdentity should have succeeded")
		return
	}

	msg := []byte("ABCDEFGABCDEFGABCDEFGABCDEFGABCDEFGABCDEFGABCDEFGABCDEFGABCDEFGABCDEFGABCDEFGABCDEFGABCDEFGABCDEFG")
	sig, err := id.Sign(msg)
	if err != nil {
		t.Fatalf("Sign should have succeeded")
		return
	}

	err = id.Verify(msg, sig)
	if err != nil {
		t.Fatalf("The signature should be valid")
		return
	}

	err = idBack.Verify(msg, sig)
	if err != nil {
		t.Fatalf("The signature should be valid")
		return
	}
}

func TestGetOU(t *testing.T) {
	id, err := localMsp.GetDefaultSigningIdentity()
	if err != nil {
		t.Fatalf("GetSigningIdentity should have succeeded")
		return
	}

	assert.Equal(t, "COP", id.GetOrganizationalUnits()[0].OrganizationalUnitIdentifier)
}

func TestGetOUFail(t *testing.T) {
	id, err := localMspBad.GetDefaultSigningIdentity()
	if err != nil {
		t.Fatalf("GetSigningIdentity should have succeeded")
		return
	}

	certTmp := id.(*signingidentity).cert
	id.(*signingidentity).cert = nil
	ou := id.GetOrganizationalUnits()
	assert.Nil(t, ou)

	id.(*signingidentity).cert = certTmp

	opts := id.(*signingidentity).msp.opts
	id.(*signingidentity).msp.opts = nil
	ou = id.GetOrganizationalUnits()
	assert.Nil(t, ou)

	id.(*signingidentity).msp.opts = opts
}

func TestCertificationIdentifierComputation(t *testing.T) {
	id, err := localMsp.GetDefaultSigningIdentity()
	assert.NoError(t, err)

	chain, err := localMsp.(*bccspmsp).getCertificationChain(id.GetPublicVersion())
	assert.NoError(t, err)

	// Hash the chain
	// Use the hash of the identity's certificate as id in the IdentityIdentifier
	hashOpt, err := bccsp.GetHashOpt(localMsp.(*bccspmsp).cryptoConfig.IdentityIdentifierHashFunction)
	assert.NoError(t, err)

	hf, err := localMsp.(*bccspmsp).bccsp.GetHash(hashOpt)
	assert.NoError(t, err)
	// Skipping first cert because it belongs to the identity
	for i := 1; i < len(chain); i++ {
		hf.Write(chain[i].Raw)
	}
	sum := hf.Sum(nil)

	assert.Equal(t, sum, id.GetOrganizationalUnits()[0].CertifiersIdentifier)
}

func TestOUPolicyPrincipal(t *testing.T) {
	id, err := localMsp.GetDefaultSigningIdentity()
	assert.NoError(t, err)

	cid, err := localMsp.(*bccspmsp).getCertificationChainIdentifier(id.GetPublicVersion())
	assert.NoError(t, err)

	ou := &msp.OrganizationUnit{
		OrganizationalUnitIdentifier: "COP",
		MspIdentifier:                "SampleOrg",
		CertifiersIdentifier:         cid,
	}
	bytes, err := proto.Marshal(ou)
	assert.NoError(t, err)

	principal := &msp.MSPPrincipal{
		PrincipalClassification: msp.MSPPrincipal_ORGANIZATION_UNIT,
		Principal:               bytes,
	}

	err = id.SatisfiesPrincipal(principal)
	assert.NoError(t, err)
}

func TestOUPolicyPrincipalBadPrincipal(t *testing.T) {
	id, err := localMsp.GetDefaultSigningIdentity()
	assert.NoError(t, err)

	principal := &msp.MSPPrincipal{
		PrincipalClassification: msp.MSPPrincipal_ORGANIZATION_UNIT,
		Principal:               []byte("barf"),
	}

	err = id.SatisfiesPrincipal(principal)
	assert.Error(t, err)
}

func TestOUPolicyPrincipalBadMSPID(t *testing.T) {
	id, err := localMsp.GetDefaultSigningIdentity()
	assert.NoError(t, err)

	cid, err := localMsp.(*bccspmsp).getCertificationChainIdentifier(id.GetPublicVersion())
	assert.NoError(t, err)

	ou := &msp.OrganizationUnit{
		OrganizationalUnitIdentifier: "COP",
		MspIdentifier:                "SampleOrgbarfbarf",
		CertifiersIdentifier:         cid,
	}
	bytes, err := proto.Marshal(ou)
	assert.NoError(t, err)

	principal := &msp.MSPPrincipal{
		PrincipalClassification: msp.MSPPrincipal_ORGANIZATION_UNIT,
		Principal:               bytes,
	}

	err = id.SatisfiesPrincipal(principal)
	assert.Error(t, err)
}

func TestOUPolicyPrincipalBadPath(t *testing.T) {
	id, err := localMsp.GetDefaultSigningIdentity()
	assert.NoError(t, err)

	ou := &msp.OrganizationUnit{
		OrganizationalUnitIdentifier: "COP",
		MspIdentifier:                "SampleOrg",
		CertifiersIdentifier:         nil,
	}
	bytes, err := proto.Marshal(ou)
	assert.NoError(t, err)

	principal := &msp.MSPPrincipal{
		PrincipalClassification: msp.MSPPrincipal_ORGANIZATION_UNIT,
		Principal:               bytes,
	}

	err = id.SatisfiesPrincipal(principal)
	assert.Error(t, err)

	ou = &msp.OrganizationUnit{
		OrganizationalUnitIdentifier: "COP",
		MspIdentifier:                "SampleOrg",
		CertifiersIdentifier:         []byte{0, 1, 2, 3, 4},
	}
	bytes, err = proto.Marshal(ou)
	assert.NoError(t, err)

	principal = &msp.MSPPrincipal{
		PrincipalClassification: msp.MSPPrincipal_ORGANIZATION_UNIT,
		Principal:               bytes,
	}

	err = id.SatisfiesPrincipal(principal)
	assert.Error(t, err)
}

func TestPolicyPrincipalBogusType(t *testing.T) {
	id, err := localMsp.GetDefaultSigningIdentity()
	assert.NoError(t, err)

	principalBytes, err := proto.Marshal(&msp.MSPRole{Role: 35, MspIdentifier: "SampleOrg"})
	assert.NoError(t, err)

	principal := &msp.MSPPrincipal{
		PrincipalClassification: 35,
		Principal:               principalBytes}

	err = id.SatisfiesPrincipal(principal)
	assert.Error(t, err)
}

func TestPolicyPrincipalBogusRole(t *testing.T) {
	id, err := localMsp.GetDefaultSigningIdentity()
	assert.NoError(t, err)

	principalBytes, err := proto.Marshal(&msp.MSPRole{Role: 35, MspIdentifier: "SampleOrg"})
	assert.NoError(t, err)

	principal := &msp.MSPPrincipal{
		PrincipalClassification: msp.MSPPrincipal_ROLE,
		Principal:               principalBytes}

	err = id.SatisfiesPrincipal(principal)
	assert.Error(t, err)
}

func TestPolicyPrincipalWrongMSPID(t *testing.T) {
	id, err := localMsp.GetDefaultSigningIdentity()
	assert.NoError(t, err)

	principalBytes, err := proto.Marshal(&msp.MSPRole{Role: msp.MSPRole_MEMBER, MspIdentifier: "SampleOrgBARFBARF"})
	assert.NoError(t, err)

	principal := &msp.MSPPrincipal{
		PrincipalClassification: msp.MSPPrincipal_ROLE,
		Principal:               principalBytes}

	err = id.SatisfiesPrincipal(principal)
	assert.Error(t, err)
}

func TestMemberPolicyPrincipal(t *testing.T) {
	id, err := localMsp.GetDefaultSigningIdentity()
	assert.NoError(t, err)

	principalBytes, err := proto.Marshal(&msp.MSPRole{Role: msp.MSPRole_MEMBER, MspIdentifier: "SampleOrg"})
	assert.NoError(t, err)

	principal := &msp.MSPPrincipal{
		PrincipalClassification: msp.MSPPrincipal_ROLE,
		Principal:               principalBytes}

	err = id.SatisfiesPrincipal(principal)
	assert.NoError(t, err)
}

func TestAdminPolicyPrincipal(t *testing.T) {
	id, err := localMsp.GetDefaultSigningIdentity()
	assert.NoError(t, err)

	principalBytes, err := proto.Marshal(&msp.MSPRole{Role: msp.MSPRole_ADMIN, MspIdentifier: "SampleOrg"})
	assert.NoError(t, err)

	principal := &msp.MSPPrincipal{
		PrincipalClassification: msp.MSPPrincipal_ROLE,
		Principal:               principalBytes}

	err = id.SatisfiesPrincipal(principal)
	assert.NoError(t, err)
}

// Combine one or more MSPPrincipals into a MSPPrincipal of type
// MSPPrincipal_COMBINED.
func createCombinedPrincipal(principals ...*msp.MSPPrincipal) (*msp.MSPPrincipal, error) {
	if len(principals) == 0 {
		return nil, errors.New("no principals in CombinedPrincipal")
	}
	principalsArray := append([]*msp.MSPPrincipal(nil), principals...)
	combinedPrincipal := &msp.CombinedPrincipal{Principals: principalsArray}
	combinedPrincipalBytes, err := proto.Marshal(combinedPrincipal)
	if err != nil {
		return nil, err
	}
	principalsCombined := &msp.MSPPrincipal{PrincipalClassification: msp.MSPPrincipal_COMBINED, Principal: combinedPrincipalBytes}
	return principalsCombined, nil
}

func TestMultilevelAdminAndMemberPolicyPrincipal(t *testing.T) {
	id, err := localMspV13.GetDefaultSigningIdentity()
	assert.NoError(t, err)

	adminPrincipalBytes, err := proto.Marshal(&msp.MSPRole{Role: msp.MSPRole_ADMIN, MspIdentifier: "SampleOrg"})
	assert.NoError(t, err)

	memberPrincipalBytes, err := proto.Marshal(&msp.MSPRole{Role: msp.MSPRole_MEMBER, MspIdentifier: "SampleOrg"})
	assert.NoError(t, err)

	adminPrincipal := &msp.MSPPrincipal{
		PrincipalClassification: msp.MSPPrincipal_ROLE,
		Principal:               adminPrincipalBytes}

	memberPrincipal := &msp.MSPPrincipal{
		PrincipalClassification: msp.MSPPrincipal_ROLE,
		Principal:               memberPrincipalBytes}

	// CombinedPrincipal with Admin and Member principals
	levelOneCombinedPrincipal, err := createCombinedPrincipal(adminPrincipal, memberPrincipal)
	assert.NoError(t, err)
	err = id.SatisfiesPrincipal(levelOneCombinedPrincipal)
	assert.NoError(t, err)

	// Nested CombinedPrincipal
	levelTwoCombinedPrincipal, err := createCombinedPrincipal(levelOneCombinedPrincipal)
	assert.NoError(t, err)
	err = id.SatisfiesPrincipal(levelTwoCombinedPrincipal)
	assert.NoError(t, err)

	// Double nested CombinedPrincipal
	levelThreeCombinedPrincipal, err := createCombinedPrincipal(levelTwoCombinedPrincipal)
	assert.NoError(t, err)
	err = id.SatisfiesPrincipal(levelThreeCombinedPrincipal)
	assert.NoError(t, err)
}

func TestMultilevelAdminAndMemberPolicyPrincipalPreV12(t *testing.T) {
	id, err := localMspV11.GetDefaultSigningIdentity()
	assert.NoError(t, err)

	adminPrincipalBytes, err := proto.Marshal(&msp.MSPRole{Role: msp.MSPRole_ADMIN, MspIdentifier: "SampleOrg"})
	assert.NoError(t, err)

	memberPrincipalBytes, err := proto.Marshal(&msp.MSPRole{Role: msp.MSPRole_MEMBER, MspIdentifier: "SampleOrg"})
	assert.NoError(t, err)

	adminPrincipal := &msp.MSPPrincipal{
		PrincipalClassification: msp.MSPPrincipal_ROLE,
		Principal:               adminPrincipalBytes}

	memberPrincipal := &msp.MSPPrincipal{
		PrincipalClassification: msp.MSPPrincipal_ROLE,
		Principal:               memberPrincipalBytes}

	// CombinedPrincipal with Admin and Member principals
	levelOneCombinedPrincipal, err := createCombinedPrincipal(adminPrincipal, memberPrincipal)
	assert.NoError(t, err)
	err = id.SatisfiesPrincipal(levelOneCombinedPrincipal)
	assert.Error(t, err)

	// Nested CombinedPrincipal
	levelTwoCombinedPrincipal, err := createCombinedPrincipal(levelOneCombinedPrincipal)
	assert.NoError(t, err)
	err = id.SatisfiesPrincipal(levelTwoCombinedPrincipal)
	assert.Error(t, err)

	// Double nested CombinedPrincipal
	levelThreeCombinedPrincipal, err := createCombinedPrincipal(levelTwoCombinedPrincipal)
	assert.NoError(t, err)
	err = id.SatisfiesPrincipal(levelThreeCombinedPrincipal)
	assert.Error(t, err)
}

func TestAdminPolicyPrincipalFails(t *testing.T) {
	id, err := localMspV13.GetDefaultSigningIdentity()
	assert.NoError(t, err)

	principalBytes, err := proto.Marshal(&msp.MSPRole{Role: msp.MSPRole_ADMIN, MspIdentifier: "SampleOrg"})
	assert.NoError(t, err)

	principal := &msp.MSPPrincipal{
		PrincipalClassification: msp.MSPPrincipal_ROLE,
		Principal:               principalBytes}

	// remove the admin so validation will fail
	localMspV13.(*bccspmsp).admins = make([]Identity, 0)

	err = id.SatisfiesPrincipal(principal)
	assert.Error(t, err)
}

func TestMultilevelAdminAndMemberPolicyPrincipalFails(t *testing.T) {
	id, err := localMspV13.GetDefaultSigningIdentity()
	assert.NoError(t, err)

	adminPrincipalBytes, err := proto.Marshal(&msp.MSPRole{Role: msp.MSPRole_ADMIN, MspIdentifier: "SampleOrg"})
	assert.NoError(t, err)

	memberPrincipalBytes, err := proto.Marshal(&msp.MSPRole{Role: msp.MSPRole_MEMBER, MspIdentifier: "SampleOrg"})
	assert.NoError(t, err)

	adminPrincipal := &msp.MSPPrincipal{
		PrincipalClassification: msp.MSPPrincipal_ROLE,
		Principal:               adminPrincipalBytes}

	memberPrincipal := &msp.MSPPrincipal{
		PrincipalClassification: msp.MSPPrincipal_ROLE,
		Principal:               memberPrincipalBytes}

	// remove the admin so validation will fail
	localMspV13.(*bccspmsp).admins = make([]Identity, 0)

	// CombinedPrincipal with Admin and Member principals
	levelOneCombinedPrincipal, err := createCombinedPrincipal(adminPrincipal, memberPrincipal)
	assert.NoError(t, err)
	err = id.SatisfiesPrincipal(levelOneCombinedPrincipal)
	assert.Error(t, err)

	// Nested CombinedPrincipal
	levelTwoCombinedPrincipal, err := createCombinedPrincipal(levelOneCombinedPrincipal)
	assert.NoError(t, err)
	err = id.SatisfiesPrincipal(levelTwoCombinedPrincipal)
	assert.Error(t, err)

	// Double nested CombinedPrincipal
	levelThreeCombinedPrincipal, err := createCombinedPrincipal(levelTwoCombinedPrincipal)
	assert.NoError(t, err)
	err = id.SatisfiesPrincipal(levelThreeCombinedPrincipal)
	assert.Error(t, err)
}

func TestIdentityExpiresAt(t *testing.T) {
	thisMSP := getLocalMSP(t, "testdata/expiration")
	assert.NotNil(t, thisMSP)
	si, err := thisMSP.GetDefaultSigningIdentity()
	assert.NoError(t, err)
	expirationDate := si.GetPublicVersion().ExpiresAt()
	assert.Equal(t, time.Date(2027, 8, 17, 12, 19, 48, 0, time.UTC), expirationDate)
}

func TestIdentityExpired(t *testing.T) {
	cryptoProvider, err := sw.NewDefaultSecurityLevelWithKeystore(sw.NewDummyKeyStore())
	assert.NoError(t, err)

	expiredCertsDir := "testdata/expired"
	conf, err := GetLocalMspConfig(expiredCertsDir, nil, "SampleOrg")
	assert.NoError(t, err)

	thisMSP, err := newBccspMsp(MSPv1_0, cryptoProvider)
	assert.NoError(t, err)

	ks, err := sw.NewFileBasedKeyStore(nil, filepath.Join(expiredCertsDir, "keystore"), true)
	assert.NoError(t, err)

	csp, err := sw.NewWithParams(true, 256, "SHA2", ks)
	assert.NoError(t, err)
	thisMSP.(*bccspmsp).bccsp = csp

	err = thisMSP.Setup(conf)
	if err != nil {
		assert.Contains(t, err.Error(), "signing identity expired")
	} else {
		t.Fatal("Should have failed when loading expired certs")
	}
}

func TestIdentityPolicyPrincipal(t *testing.T) {
	id, err := localMsp.GetDefaultSigningIdentity()
	assert.NoError(t, err)

	idSerialized, err := id.Serialize()
	assert.NoError(t, err)

	principal := &msp.MSPPrincipal{
		PrincipalClassification: msp.MSPPrincipal_IDENTITY,
		Principal:               idSerialized}

	err = id.SatisfiesPrincipal(principal)
	assert.NoError(t, err)
}

func TestIdentityPolicyPrincipalBadBytes(t *testing.T) {
	id, err := localMsp.GetDefaultSigningIdentity()
	assert.NoError(t, err)

	principal := &msp.MSPPrincipal{
		PrincipalClassification: msp.MSPPrincipal_IDENTITY,
		Principal:               []byte("barf")}

	err = id.SatisfiesPrincipal(principal)
	assert.Error(t, err)
}

func TestMSPOus(t *testing.T) {
	// Set the OUIdentifiers
	backup := localMsp.(*bccspmsp).ouIdentifiers
	defer func() { localMsp.(*bccspmsp).ouIdentifiers = backup }()
	sid, err := localMsp.GetDefaultSigningIdentity()
	assert.NoError(t, err)
	sidBytes, err := sid.Serialize()
	assert.NoError(t, err)
	id, err := localMsp.DeserializeIdentity(sidBytes)
	assert.NoError(t, err)

	localMsp.(*bccspmsp).ouIdentifiers = map[string][][]byte{
		"COP": {id.GetOrganizationalUnits()[0].CertifiersIdentifier},
	}
	assert.NoError(t, localMsp.Validate(id))

	id, err = localMsp.DeserializeIdentity(sidBytes)
	assert.NoError(t, err)

	localMsp.(*bccspmsp).ouIdentifiers = map[string][][]byte{
		"COP2": {id.GetOrganizationalUnits()[0].CertifiersIdentifier},
	}
	assert.Error(t, localMsp.Validate(id))

	id, err = localMsp.DeserializeIdentity(sidBytes)
	assert.NoError(t, err)

	localMsp.(*bccspmsp).ouIdentifiers = map[string][][]byte{
		"COP": {{0, 1, 2, 3, 4}},
	}
	assert.Error(t, localMsp.Validate(id))
}

const othercert = `-----BEGIN CERTIFICATE-----
MIIDAzCCAqigAwIBAgIBAjAKBggqhkjOPQQDAjBsMQswCQYDVQQGEwJHQjEQMA4G
A1UECAwHRW5nbGFuZDEOMAwGA1UECgwFQmFyMTkxDjAMBgNVBAsMBUJhcjE5MQ4w
DAYDVQQDDAVCYXIxOTEbMBkGCSqGSIb3DQEJARYMQmFyMTktY2xpZW50MB4XDTE3
MDIwOTE2MDcxMFoXDTE4MDIxOTE2MDcxMFowfDELMAkGA1UEBhMCR0IxEDAOBgNV
BAgMB0VuZ2xhbmQxEDAOBgNVBAcMB0lwc3dpY2gxDjAMBgNVBAoMBUJhcjE5MQ4w
DAYDVQQLDAVCYXIxOTEOMAwGA1UEAwwFQmFyMTkxGTAXBgkqhkiG9w0BCQEWCkJh
cjE5LXBlZXIwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQlRSnAyD+ND6qmaRV7
AS/BPJKX5dZt3gBe1v/RewOpc1zJeXQNWACAk0ae3mv5u9l0HxI6TXJIAQSwJACu
Rqsyo4IBKTCCASUwCQYDVR0TBAIwADARBglghkgBhvhCAQEEBAMCBkAwMwYJYIZI
AYb4QgENBCYWJE9wZW5TU0wgR2VuZXJhdGVkIFNlcnZlciBDZXJ0aWZpY2F0ZTAd
BgNVHQ4EFgQUwHzbLJQMaWd1cpHdkSaEFxdKB1owgYsGA1UdIwSBgzCBgIAUYxFe
+cXOD5iQ223bZNdOuKCRiTKhZaRjMGExCzAJBgNVBAYTAkdCMRAwDgYDVQQIDAdF
bmdsYW5kMRAwDgYDVQQHDAdJcHN3aWNoMQ4wDAYDVQQKDAVCYXIxOTEOMAwGA1UE
CwwFQmFyMTkxDjAMBgNVBAMMBUJhcjE5ggEBMA4GA1UdDwEB/wQEAwIFoDATBgNV
HSUEDDAKBggrBgEFBQcDATAKBggqhkjOPQQDAgNJADBGAiEAuMq65lOaie4705Ol
Ow52DjbaO2YuIxK2auBCqNIu0gECIQCDoKdUQ/sa+9Ah1mzneE6iz/f/YFVWo4EP
HeamPGiDTQ==
-----END CERTIFICATE-----
`

func TestIdentityPolicyPrincipalFails(t *testing.T) {
	id, err := localMsp.GetDefaultSigningIdentity()
	assert.NoError(t, err)

	sid, err := NewSerializedIdentity("SampleOrg", []byte(othercert))
	assert.NoError(t, err)

	principal := &msp.MSPPrincipal{
		PrincipalClassification: msp.MSPPrincipal_IDENTITY,
		Principal:               sid}

	err = id.SatisfiesPrincipal(principal)
	assert.Error(t, err)
}

var conf *msp.MSPConfig
var localMsp MSP
var localMspV11 MSP
var localMspV13 MSP

// Required because deleting the cert or msp options from localMsp causes parallel tests to fail
var localMspBad MSP
var mspMgr MSPManager

func TestMain(m *testing.M) {
	var err error

	mspDir := configtest.GetDevMspDir()
	conf, err = GetLocalMspConfig(mspDir, nil, "SampleOrg")
	if err != nil {
		fmt.Printf("Setup should have succeeded, got err %s instead", err)
		os.Exit(-1)
	}

	localMsp, err = newBccspMsp(MSPv1_0, factory.GetDefault())
	if err != nil {
		fmt.Printf("Constructor for msp should have succeeded, got err %s instead", err)
		os.Exit(-1)
	}

	localMspBad, err = newBccspMsp(MSPv1_0, factory.GetDefault())
	if err != nil {
		fmt.Printf("Constructor for msp should have succeeded, got err %s instead", err)
		os.Exit(-1)
	}

	localMspV13, err = newBccspMsp(MSPv1_3, factory.GetDefault())
	if err != nil {
		fmt.Printf("Constructor for V1.3 msp should have succeeded, got err %s instead", err)
		os.Exit(-1)
	}

	localMspV11, err = newBccspMsp(MSPv1_1, factory.GetDefault())
	if err != nil {
		fmt.Printf("Constructor for V1.1 msp should have succeeded, got err %s instead", err)
		os.Exit(-1)
	}

	err = localMspV11.Setup(conf)
	if err != nil {
		fmt.Printf("Setup for V1.1 msp should have succeeded, got err %s instead", err)
		os.Exit(-1)
	}

	err = localMspV13.Setup(conf)
	if err != nil {
		fmt.Printf("Setup for V1.3 msp should have succeeded, got err %s instead", err)
		os.Exit(-1)
	}

	err = localMsp.Setup(conf)
	if err != nil {
		fmt.Printf("Setup for msp should have succeeded, got err %s instead", err)
		os.Exit(-1)
	}

	err = localMspBad.Setup(conf)
	if err != nil {
		fmt.Printf("Setup for msp should have succeeded, got err %s instead", err)
		os.Exit(-1)
	}

	mspMgr = NewMSPManager()
	err = mspMgr.Setup([]MSP{localMsp})
	if err != nil {
		fmt.Printf("Setup for msp manager should have succeeded, got err %s instead", err)
		os.Exit(-1)
	}

	id, err := localMsp.GetIdentifier()
	if err != nil {
		fmt.Println("Failed obtaining identifier for localMSP")
		os.Exit(-1)
	}

	msps, err := mspMgr.GetMSPs()
	if err != nil {
		fmt.Println("Failed obtaining MSPs from MSP manager")
		os.Exit(-1)
	}

	if msps[id] == nil {
		fmt.Println("Couldn't find localMSP in MSP manager")
		os.Exit(-1)
	}

	retVal := m.Run()
	os.Exit(retVal)
}

func getIdentity(t *testing.T, path string) Identity {
	mspDir := configtest.GetDevMspDir()
	pems, err := getPemMaterialFromDir(filepath.Join(mspDir, path))
	assert.NoError(t, err)

	id, _, err := localMsp.(*bccspmsp).getIdentityFromConf(pems[0])
	assert.NoError(t, err)

	return id
}

func getLocalMSPWithVersionAndError(t *testing.T, dir string, version MSPVersion) (MSP, error) {
	conf, err := GetLocalMspConfig(dir, nil, "SampleOrg")
	assert.NoError(t, err)

	ks, err := sw.NewFileBasedKeyStore(nil, filepath.Join(dir, "keystore"), true)
	assert.NoError(t, err)
	cryptoProvider, err := sw.NewDefaultSecurityLevelWithKeystore(sw.NewDummyKeyStore())
	assert.NoError(t, err)
	thisMSP, err := NewBccspMspWithKeyStore(version, ks, cryptoProvider)
	assert.NoError(t, err)

	return thisMSP, thisMSP.Setup(conf)
}

func getLocalMSP(t *testing.T, dir string) MSP {
	conf, err := GetLocalMspConfig(dir, nil, "SampleOrg")
	assert.NoError(t, err)

	ks, err := sw.NewFileBasedKeyStore(nil, filepath.Join(dir, "keystore"), true)
	assert.NoError(t, err)

	cryptoProvider, err := sw.NewDefaultSecurityLevelWithKeystore(sw.NewDummyKeyStore())
	assert.NoError(t, err)
	thisMSP, err := NewBccspMspWithKeyStore(MSPv1_0, ks, cryptoProvider)
	assert.NoError(t, err)
	err = thisMSP.Setup(conf)
	assert.NoError(t, err)

	return thisMSP
}

func getLocalMSPWithVersion(t *testing.T, dir string, version MSPVersion) MSP {
	conf, err := GetLocalMspConfig(dir, nil, "SampleOrg")
	assert.NoError(t, err)

	ks, err := sw.NewFileBasedKeyStore(nil, filepath.Join(dir, "keystore"), true)
	assert.NoError(t, err)
	cryptoProvider, err := sw.NewDefaultSecurityLevelWithKeystore(sw.NewDummyKeyStore())
	assert.NoError(t, err)
	thisMSP, err := NewBccspMspWithKeyStore(version, ks, cryptoProvider)
	assert.NoError(t, err)

	err = thisMSP.Setup(conf)
	assert.NoError(t, err)

	return thisMSP
}

func TestCollectEmptyCombinedPrincipal(t *testing.T) {
	var principalsArray []*msp.MSPPrincipal
	combinedPrincipal := &msp.CombinedPrincipal{Principals: principalsArray}
	combinedPrincipalBytes, err := proto.Marshal(combinedPrincipal)
	assert.NoError(t, err, "Error marshalling empty combined principal")
	principalsCombined := &msp.MSPPrincipal{PrincipalClassification: msp.MSPPrincipal_COMBINED, Principal: combinedPrincipalBytes}
	_, err = collectPrincipals(principalsCombined, MSPv1_3)
	assert.Error(t, err)
}

func TestCollectPrincipalContainingEmptyCombinedPrincipal(t *testing.T) {
	var principalsArray []*msp.MSPPrincipal
	combinedPrincipal := &msp.CombinedPrincipal{Principals: principalsArray}
	combinedPrincipalBytes, err := proto.Marshal(combinedPrincipal)
	assert.NoError(t, err, "Error marshalling empty combined principal")
	emptyPrincipal := &msp.MSPPrincipal{PrincipalClassification: msp.MSPPrincipal_COMBINED, Principal: combinedPrincipalBytes}
	levelOneCombinedPrincipal, err := createCombinedPrincipal(emptyPrincipal)
	assert.NoError(t, err)
	_, err = collectPrincipals(levelOneCombinedPrincipal, MSPv1_3)
	assert.Error(t, err)
}

func TestMSPIdentityIdentifier(t *testing.T) {
	// testdata/mspid
	// 1) a key and a signcert (used to populate the default signing identity) with the cert having a HighS signature
	thisMSP := getLocalMSP(t, "testdata/mspid")

	id, err := thisMSP.GetDefaultSigningIdentity()
	assert.NoError(t, err)
	err = id.Validate()
	assert.NoError(t, err)

	// Check that the identity identifier is computed with the respect to the lowS signature

	idid := id.GetIdentifier()
	assert.NotNil(t, idid)

	// Load and parse cacaert and signcert from folder
	pems, err := getPemMaterialFromDir("testdata/mspid/cacerts")
	assert.NoError(t, err)
	bl, _ := pem.Decode(pems[0])
	assert.NotNil(t, bl)
	// caCertFromFile, err := x509.ParseCertificate(bl.Bytes)
	assert.NoError(t, err)

	pems, err = getPemMaterialFromDir("testdata/mspid/signcerts")
	assert.NoError(t, err)
	bl, _ = pem.Decode(pems[0])
	assert.NotNil(t, bl)
	certFromFile, err := x509.ParseCertificate(bl.Bytes)
	assert.NoError(t, err)
	// Check that the certificates' raws are different, meaning that the identity has been sanitised
	assert.NotEqual(t, certFromFile.Raw, id.(*signingidentity).cert)

	// Check that certFromFile is in HighS
	// _, S, err := utils.UnmarshalECDSASignature(certFromFile.Signature)
	_, _, err = sw.UnmarshalSM2Signature(certFromFile.Signature)
	assert.NoError(t, err)
	// lowS, err := utils.IsLowS(caCertFromFile.PublicKey.(*ecdsa.PublicKey), S)
	// assert.NoError(t, err)
	// assert.False(t, lowS)

	// Check that id.(*signingidentity).cert is in LoswS
	// _, S, err = utils.UnmarshalECDSASignature(id.(*signingidentity).cert.Signature)
	_, _, err = sw.UnmarshalSM2Signature(id.(*signingidentity).cert.Signature)
	assert.NoError(t, err)
	// lowS, err = utils.IsLowS(caCertFromFile.PublicKey.(*ecdsa.PublicKey), S)
	// assert.NoError(t, err)
	// assert.True(t, lowS)

	// Compute the digest for certFromFile
	thisBCCSPMsp := thisMSP.(*bccspmsp)
	hashOpt, err := bccsp.GetHashOpt(thisBCCSPMsp.cryptoConfig.IdentityIdentifierHashFunction)
	assert.NoError(t, err)
	digest, err := thisBCCSPMsp.bccsp.Hash(certFromFile.Raw, hashOpt)
	assert.NoError(t, err)

	// Compare with the digest computed from the sanitised cert
	assert.NotEqual(t, idid.Id, hex.EncodeToString(digest))
}

func TestAnonymityIdentity(t *testing.T) {
	id, err := localMspV13.GetDefaultSigningIdentity()
	assert.NoError(t, err)

	principalBytes, err := proto.Marshal(&msp.MSPIdentityAnonymity{AnonymityType: msp.MSPIdentityAnonymity_NOMINAL})
	assert.NoError(t, err)

	principal := &msp.MSPPrincipal{
		PrincipalClassification: msp.MSPPrincipal_ANONYMITY,
		Principal:               principalBytes}

	err = id.SatisfiesPrincipal(principal)
	assert.NoError(t, err)
}

func TestAnonymityIdentityPreV12Fail(t *testing.T) {
	id, err := localMspV11.GetDefaultSigningIdentity()
	assert.NoError(t, err)

	principalBytes, err := proto.Marshal(&msp.MSPIdentityAnonymity{AnonymityType: msp.MSPIdentityAnonymity_NOMINAL})
	assert.NoError(t, err)

	principal := &msp.MSPPrincipal{
		PrincipalClassification: msp.MSPPrincipal_ANONYMITY,
		Principal:               principalBytes}

	err = id.SatisfiesPrincipal(principal)
	assert.Error(t, err)
}

func TestAnonymityIdentityFail(t *testing.T) {
	id, err := localMspV13.GetDefaultSigningIdentity()
	assert.NoError(t, err)

	principalBytes, err := proto.Marshal(&msp.MSPIdentityAnonymity{AnonymityType: msp.MSPIdentityAnonymity_ANONYMOUS})
	assert.NoError(t, err)

	principal := &msp.MSPPrincipal{
		PrincipalClassification: msp.MSPPrincipal_ANONYMITY,
		Principal:               principalBytes}

	err = id.SatisfiesPrincipal(principal)
	assert.Error(t, err)
}

func TestProviderTypeToString(t *testing.T) {
	// Check that the provider type is found for FABRIC
	pt := ProviderTypeToString(FABRIC)
	assert.Equal(t, "bccsp", pt)

	// Check that the provider type is found for IDEMIX
	pt = ProviderTypeToString(IDEMIX)
	assert.Equal(t, "idemix", pt)

	// Check that the provider type is not found
	pt = ProviderTypeToString(OTHER)
	assert.Equal(t, "", pt)
}
