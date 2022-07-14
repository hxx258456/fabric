/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package metadata_test

import (
	"fmt"
	"runtime"
	"testing"

	common "github.com/hxx258456/fabric/common/metadata"
	"github.com/hxx258456/fabric/orderer/common/metadata"
	"github.com/stretchr/testify/require"
)

func TestGetVersionInfo(t *testing.T) {
	expected := fmt.Sprintf(
		"%s:\n Version: %s\n Commit SHA: %s\n Go version: %s\n OS/Arch: %s\n",
		metadata.ProgramName, common.Version,
		common.CommitSHA,
		runtime.Version(),
		fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	)
	require.Equal(t, expected, metadata.GetVersionInfo())
}
