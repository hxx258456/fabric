/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package msgprocessor

import (
	"time"

	"github.com/hxx258456/fabric-protos-go-cc/common"
	"github.com/hxx258456/fabric/common/channelconfig"
	"github.com/hxx258456/fabric/common/crypto"
	"github.com/hxx258456/fabric/protoutil"
	"github.com/pkg/errors"
)

type resources interface {
	// OrdererConfig returns the config.Orderer for the channel
	// and whether the Orderer config exists
	OrdererConfig() (channelconfig.Orderer, bool)
}

// NewExpirationRejectRule returns a rule that rejects messages signed by identities
// who's identities have expired, given the capability is active
func NewExpirationRejectRule(filterSupport resources) Rule {
	return &expirationRejectRule{filterSupport: filterSupport}
}

type expirationRejectRule struct {
	filterSupport resources
}

// Apply checks whether the identity that created the envelope has expired
func (exp *expirationRejectRule) Apply(message *common.Envelope) error {
	ordererConf, ok := exp.filterSupport.OrdererConfig()
	if !ok {
		logger.Panic("Programming error: orderer config not found")
	}
	if !ordererConf.Capabilities().ExpirationCheck() {
		return nil
	}
	signedData, err := protoutil.EnvelopeAsSignedData(message)
	if err != nil {
		return errors.Errorf("could not convert message to signedData: %s", err)
	}
	expirationTime := crypto.ExpiresAt(signedData[0].Identity)
	// Identity cannot expire, or identity has not expired yet
	if expirationTime.IsZero() || time.Now().Before(expirationTime) {
		return nil
	}
	return errors.New("broadcast client identity expired")
}
