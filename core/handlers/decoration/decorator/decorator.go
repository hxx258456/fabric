/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package decorator

import (
	"github.com/hxx258456/fabric-protos-go-cc/peer"
	"github.com/hxx258456/fabric/core/handlers/decoration"
)

// NewDecorator creates a new decorator
func NewDecorator() decoration.Decorator {
	return &decorator{}
}

type decorator struct{}

// Decorate decorates a chaincode input by changing it
func (d *decorator) Decorate(proposal *peer.Proposal, input *peer.ChaincodeInput) *peer.ChaincodeInput {
	return input
}
