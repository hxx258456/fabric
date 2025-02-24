/*
Copyright IBM Corp. 2017 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package sysccprovider

import (
	"github.com/hxx258456/fabric/common/channelconfig"
	"github.com/hxx258456/fabric/common/policies"
	"github.com/hxx258456/fabric/core/ledger"
)

// SystemChaincodeProvider provides an abstraction layer that is
// used for different packages to interact with code in the
// system chaincode package without importing it; more methods
// should be added below if necessary
type SystemChaincodeProvider interface {
	// GetQueryExecutorForLedger returns a query executor for the
	// ledger of the supplied channel.
	// That's useful for system chaincodes that require unfettered
	// access to the ledger
	GetQueryExecutorForLedger(cid string) (ledger.QueryExecutor, error)

	// GetApplicationConfig returns the configtxapplication.SharedConfig for the channel
	// and whether the Application config exists
	GetApplicationConfig(cid string) (channelconfig.Application, bool)

	// Returns the policy manager associated to the passed channel
	// and whether the policy manager exists
	PolicyManager(channelID string) (policies.Manager, bool)
}

// ChaincodeInstance is unique identifier of chaincode instance
type ChaincodeInstance struct {
	ChannelID        string
	ChaincodeName    string
	ChaincodeVersion string
}

func (ci *ChaincodeInstance) String() string {
	return ci.ChannelID + "." + ci.ChaincodeName + "#" + ci.ChaincodeVersion
}
