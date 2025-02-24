/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package chaincode

import (
	"testing"

	"github.com/hxx258456/fabric/bccsp/sw"
	"github.com/stretchr/testify/require"
)

func TestInstantiateCmd(t *testing.T) {
	mockCF, err := getMockChaincodeCmdFactory()
	require.NoError(t, err, "Error getting mock chaincode command factory")
	cryptoProvider, err := sw.NewDefaultSecurityLevelWithKeystore(sw.NewDummyKeyStore())
	require.NoError(t, err)

	// basic function tests
	tests := []struct {
		name          string
		args          []string
		errorExpected bool
		errMsg        string
	}{
		{
			name:          "successful",
			args:          []string{"-n", "example02", "-v", "anotherversion", "-C", "mychannel", "-c", "{\"Args\": [\"init\",\"a\",\"100\",\"b\",\"200\"]}"},
			errorExpected: false,
			errMsg:        "Run chaincode instantiate cmd error",
		},
		{
			name:          "no option",
			args:          []string{},
			errorExpected: true,
			errMsg:        "Expected error executing instantiate command without required options",
		},
		{
			name:          "missing version",
			args:          []string{"-n", "example02", "-C", "mychannel", "-c", "{\"Args\": [\"init\",\"a\",\"100\",\"b\",\"200\"]}"},
			errorExpected: true,
			errMsg:        "Expected error executing instantiate command without the -v option",
		},
		{
			name:          "missing name",
			args:          []string{"-v", "anotherversion", "-C", "mychannel", "-c", "{\"Args\": [\"init\",\"a\",\"100\",\"b\",\"200\"]}"},
			errorExpected: true,
			errMsg:        "Expected error executing instantiate command without the -n option",
		},
		{
			name:          "missing channelID",
			args:          []string{"-n", "example02", "-v", "anotherversion", "-c", "{\"Args\": [\"init\",\"a\",\"100\",\"b\",\"200\"]}"},
			errorExpected: true,
			errMsg:        "Expected error executing instantiate command without the -C option",
		},
		{
			name:          "missing ctor",
			args:          []string{"-n", "example02", "-C", "mychannel", "-v", "anotherversion"},
			errorExpected: true,
			errMsg:        "Expected error executing instantiate command without the -c option",
		},
		{
			name:          "successful with policy",
			args:          []string{"-P", "OR('MSP.member', 'MSP.WITH.DOTS.member', 'MSP-WITH-DASHES.member')", "-n", "example02", "-v", "anotherversion", "-C", "mychannel", "-c", "{\"Args\": [\"init\",\"a\",\"100\",\"b\",\"200\"]}"},
			errorExpected: false,
			errMsg:        "Run chaincode instantiate cmd error",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			resetFlags()
			cmd := instantiateCmd(mockCF, cryptoProvider)
			addFlags(cmd)
			cmd.SetArgs(test.args)
			err = cmd.Execute()
			checkError(t, err, test.errorExpected, test.errMsg)
		})
	}
}

func checkError(t *testing.T, err error, expectedError bool, msg string) {
	if expectedError {
		require.Error(t, err, msg)
	} else {
		require.NoError(t, err, msg)
	}
}
