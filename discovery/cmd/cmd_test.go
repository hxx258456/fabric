/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package discovery_test

import (
	"testing"

	discovery "github.com/hxx258456/fabric/discovery/cmd"
	"github.com/hxx258456/fabric/discovery/cmd/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"gopkg.in/alecthomas/kingpin.v2"
)

func TestAddCommands(t *testing.T) {
	app := kingpin.New("foo", "bar")
	cli := &mocks.CommandRegistrar{}
	configFunc := mock.AnythingOfType("common.CLICommand")
	cli.On("Command", discovery.PeersCommand, mock.Anything, configFunc).Return(app.Command(discovery.PeersCommand, ""))
	cli.On("Command", discovery.ConfigCommand, mock.Anything, configFunc).Return(app.Command(discovery.ConfigCommand, ""))
	cli.On("Command", discovery.EndorsersCommand, mock.Anything, configFunc).Return(app.Command(discovery.EndorsersCommand, ""))
	discovery.AddCommands(cli)
	// Ensure that serve and channel flags are were configured for the sub-commands
	for _, cmd := range []string{discovery.PeersCommand, discovery.ConfigCommand, discovery.EndorsersCommand} {
		require.NotNil(t, app.GetCommand(cmd).GetFlag("server"))
		require.NotNil(t, app.GetCommand(cmd).GetFlag("channel"))
	}
	// Ensure that chaincode and collection flags were called for the endorsers
	require.NotNil(t, app.GetCommand(discovery.EndorsersCommand).GetFlag("chaincode"))
	require.NotNil(t, app.GetCommand(discovery.EndorsersCommand).GetFlag("collection"))
}
