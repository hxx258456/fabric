/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package privdata

import "github.com/hxx258456/fabric/common/channelconfig"

//go:generate mockery -dir ./ -name AppCapabilities -case underscore -output mocks/
// appCapabilities local interface used to generate mock for foreign interface.
type AppCapabilities interface {
	channelconfig.ApplicationCapabilities
}
