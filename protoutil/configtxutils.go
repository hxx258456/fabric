/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package protoutil

import "github.com/hxx258456/fabric-protos-go-cc/common"

func NewConfigGroup() *common.ConfigGroup {
	return &common.ConfigGroup{
		Groups:   make(map[string]*common.ConfigGroup),
		Values:   make(map[string]*common.ConfigValue),
		Policies: make(map[string]*common.ConfigPolicy),
	}
}
