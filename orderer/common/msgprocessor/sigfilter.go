/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package msgprocessor

import (
	"fmt"

	cb "github.com/hxx258456/fabric-protos-go-cc/common"
	"github.com/hxx258456/fabric-protos-go-cc/orderer"
	"github.com/hxx258456/fabric/common/channelconfig"
	"github.com/hxx258456/fabric/common/policies"
	"github.com/hxx258456/fabric/protoutil"
	"github.com/pkg/errors"
)

// SigFilterSupport provides the resources required for the signature filter
type SigFilterSupport interface {
	// PolicyManager returns a reference to the current policy manager
	PolicyManager() policies.Manager
	// OrdererConfig returns the config.Orderer for the channel and whether the Orderer config exists
	OrdererConfig() (channelconfig.Orderer, bool)
}

// SigFilter stores the name of the policy to apply to deliver requests to
// determine whether a client is authorized
type SigFilter struct {
	normalPolicyName      string
	maintenancePolicyName string
	support               SigFilterSupport
}

// NewSigFilter creates a new signature filter, at every evaluation, the policy manager is called
// to retrieve the latest version of the policy.
//
// normalPolicyName is applied when Orderer/ConsensusType.State = NORMAL
// maintenancePolicyName is applied when Orderer/ConsensusType.State = MAINTENANCE
func NewSigFilter(normalPolicyName, maintenancePolicyName string, support SigFilterSupport) *SigFilter {
	return &SigFilter{
		normalPolicyName:      normalPolicyName,
		maintenancePolicyName: maintenancePolicyName,
		support:               support,
	}
}

// Apply applies the policy given, resulting in Reject or Forward, never Accept
func (sf *SigFilter) Apply(message *cb.Envelope) error {
	ordererConf, ok := sf.support.OrdererConfig()
	if !ok {
		logger.Panic("Programming error: orderer config not found")
	}

	signedData, err := protoutil.EnvelopeAsSignedData(message)
	if err != nil {
		return fmt.Errorf("could not convert message to signedData: %s", err)
	}

	// In maintenance mode, we typically require the signature of /Channel/Orderer/Writers.
	// This will filter out configuration changes that are not related to consensus-type migration
	// (e.g on /Channel/Application), and will block Deliver requests from peers (which are normally /Channel/Readers).
	policyName := sf.normalPolicyName
	if ordererConf.ConsensusState() == orderer.ConsensusType_STATE_MAINTENANCE {
		policyName = sf.maintenancePolicyName
	}

	policy, ok := sf.support.PolicyManager().GetPolicy(policyName)
	if !ok {
		return fmt.Errorf("could not find policy %s", policyName)
	}

	err = policy.EvaluateSignedData(signedData)
	if err != nil {
		logger.Warnw("SigFilter evaluation failed", "error", err.Error(), "ConsensusState", ordererConf.ConsensusState(), "policyName", policyName, "signingIdentity", protoutil.LogMessageForSerializedIdentities(signedData))
		return errors.Wrap(errors.WithStack(ErrPermissionDenied), err.Error())
	}
	return nil
}
