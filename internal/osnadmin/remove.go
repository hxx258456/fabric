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

// Removes an OSN from an existing channel.
func Remove(osnURL, channelID string, caCertPool *x509.CertPool, tlsClientCert tls.Certificate) (*http.Response, error) {
	url := fmt.Sprintf("%s/participation/v1/channels/%s", osnURL, channelID)

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}

	return httpDo(req, caCertPool, tlsClientCert)
}
