/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package osnadmin

import (
	http "github.com/hxx258456/ccgo/gmhttp"
	tls "github.com/hxx258456/ccgo/gmtls"
	"github.com/hxx258456/ccgo/x509"
)

func httpClient(caCertPool *x509.CertPool, tlsClientCert tls.Certificate) *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      caCertPool,
				Certificates: []tls.Certificate{tlsClientCert},
			},
		},
	}
}

func httpDo(req *http.Request, caCertPool *x509.CertPool, tlsClientCert tls.Certificate) (*http.Response, error) {
	client := httpClient(caCertPool, tlsClientCert)
	return client.Do(req)
}

func httpGet(url string, caCertPool *x509.CertPool, tlsClientCert tls.Certificate) (*http.Response, error) {
	client := httpClient(caCertPool, tlsClientCert)
	return client.Get(url)
}
