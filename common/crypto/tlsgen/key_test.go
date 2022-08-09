/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package tlsgen

import (
	"encoding/pem"
	"testing"

	tls "github.com/hxx258456/ccgo/gmtls"
	"github.com/hxx258456/ccgo/x509"

	"github.com/stretchr/testify/require"
)

func TestLoadCert(t *testing.T) {
	pair, err := newCertKeyPair(false, false, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, pair)
	tlsCertPair, err := tls.X509KeyPair(pair.Cert, pair.Key)
	require.NoError(t, err)
	require.NotNil(t, tlsCertPair)
	block, _ := pem.Decode(pair.Cert)
	cert, err := x509.ParseCertificate(block.Bytes)
	require.NoError(t, err)
	require.NotNil(t, cert)
}
