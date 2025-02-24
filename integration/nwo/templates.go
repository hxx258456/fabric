/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package nwo

import "github.com/hxx258456/fabric/integration/nwo/template"

// Templates can be used to provide custom templates to GenerateConfigTree.
type Templates struct {
	ConfigTx string `yaml:"configtx,omitempty"`
	Core     string `yaml:"core,omitempty"`
	Crypto   string `yaml:"crypto,omitempty"`
	Orderer  string `yaml:"orderer,omitempty"`
}

func (t *Templates) ConfigTxTemplate() string {
	if t.ConfigTx != "" {
		return t.ConfigTx
	}
	return template.DefaultConfigTx
}

func (t *Templates) CoreTemplate() string {
	if t.Core != "" {
		return t.Core
	}
	return template.DefaultCore
}

func (t *Templates) CryptoTemplate() string {
	if t.Crypto != "" {
		return t.Crypto
	}
	return template.DefaultCrypto
}

func (t *Templates) OrdererTemplate() string {
	if t.Orderer != "" {
		return t.Orderer
	}
	return template.DefaultOrderer
}
