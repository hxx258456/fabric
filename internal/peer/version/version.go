/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package version

import (
	"fmt"
	"runtime"

	"github.com/hxx258456/fabric/common/metadata"
	"github.com/spf13/cobra"
)

// Program name
const ProgramName = "peer"

// Cmd returns the Cobra Command for Version
func Cmd() *cobra.Command {
	return cobraCommand
}

var cobraCommand = &cobra.Command{
	Use:   "version",
	Short: "Print fabric peer version.",
	Long:  `Print current version of the fabric peer server.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 {
			return fmt.Errorf("trailing args detected")
		}
		// Parsing of the command line is done so silence cmd usage
		cmd.SilenceUsage = true
		fmt.Print(GetInfo())
		return nil
	},
}

// GetInfo returns version information for the peer
func GetInfo() string {
	ccinfo := fmt.Sprintf("  Base Docker Label: %s\n"+
		"  Docker Namespace: %s\n",
		metadata.BaseDockerLabel,
		metadata.DockerNamespace)

	return fmt.Sprintf("%s:\n Version: %s\n Commit SHA: %s\n Go version: %s\n"+
		" OS/Arch: %s\n"+
		" Chaincode:\n%s\n",
		ProgramName, metadata.Version, metadata.CommitSHA, runtime.Version(),
		fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH), ccinfo)
}
