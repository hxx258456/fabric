/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package protoutil_test

import (
	"testing"

	"github.com/hxx258456/fabric-protos-go-cc/common"
	"github.com/hxx258456/fabric/protoutil"
	"github.com/stretchr/testify/require"
)

func TestNewConfigGroup(t *testing.T) {
	require.Equal(t,
		&common.ConfigGroup{
			Groups:   make(map[string]*common.ConfigGroup),
			Values:   make(map[string]*common.ConfigValue),
			Policies: make(map[string]*common.ConfigPolicy),
		},
		protoutil.NewConfigGroup(),
	)
}
