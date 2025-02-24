/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package decorator

import (
	"testing"

	"github.com/hxx258456/fabric-protos-go-cc/peer"
	"github.com/stretchr/testify/require"
)

func TestDecorator(t *testing.T) {
	dec := NewDecorator()
	in := &peer.ChaincodeInput{
		Args: [][]byte{{1, 2, 3}},
	}
	out := dec.Decorate(nil, in)
	require.Equal(t, in, out)
}
