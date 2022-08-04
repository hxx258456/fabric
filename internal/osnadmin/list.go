/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package osnadmin

import (
	"fmt"

	http "github.com/hxx258456/ccgo/gmhttp"
	tls "github.com/hxx258456/ccgo/gmtls"
	"github.com/hxx258456/ccgo/x509"
)

// Lists the channels an OSN is a member of.
func ListAllChannels(osnURL string, caCertPool *x509.CertPool, tlsClientCert tls.Certificate) (*http.Response, error) {
	url := fmt.Sprintf("%s/participation/v1/channels", osnURL)

	return httpGet(url, caCertPool, tlsClientCert)
}

// Lists a single channel an OSN is a member of.
func ListSingleChannel(osnURL, channelID string, caCertPool *x509.CertPool, tlsClientCert tls.Certificate) (*http.Response, error) {
	url := fmt.Sprintf("%s/participation/v1/channels/%s", osnURL, channelID)

	return httpGet(url, caCertPool, tlsClientCert)
}
