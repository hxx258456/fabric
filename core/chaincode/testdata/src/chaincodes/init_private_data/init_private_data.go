/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"fmt"

	pb "github.com/hxx258456/fabric-protos-go-cc/peer"
	"github.com/hyperledger/fabric-chaincode-go/shim"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

// Init initializes a private state
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	if err := stub.PutPrivateData("dummyColl", "dummyKey", []byte("dummyValue")); err != nil {
		return shim.Error(fmt.Sprintf("put operation failed. Error storing state: %s", err))
	}
	return shim.Success(nil)
}

// Invoke is a no-op
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s", err)
	}
}
