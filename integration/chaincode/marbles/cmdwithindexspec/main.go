/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"fmt"
	"os"

	"github.com/hxx258456/fabric-chaincode-go-cc/shim"
	"github.com/hxx258456/fabric/integration/chaincode/marbles"
)

func main() {
	err := shim.Start(&marbles.SimpleChaincode{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Exiting marbles.SimpleChaincode: %s", err)
		os.Exit(2)
	}
}
