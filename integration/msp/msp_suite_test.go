/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package msp

import (
	"encoding/json"
	"testing"

	"github.com/hxx258456/fabric/integration"
	"github.com/hxx258456/fabric/integration/nwo"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestMSP(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MSP Suite")
}

var (
	buildServer *nwo.BuildServer
	components  *nwo.Components
)

var _ = SynchronizedBeforeSuite(func() []byte {
	buildServer = nwo.NewBuildServer()
	buildServer.Serve()

	components = buildServer.Components()
	payload, err := json.Marshal(components)
	Expect(err).NotTo(HaveOccurred())
	return payload
}, func(payload []byte) {
	err := json.Unmarshal(payload, &components)
	Expect(err).NotTo(HaveOccurred())
})

var _ = SynchronizedAfterSuite(func() {
}, func() {
	buildServer.Shutdown()
})

func StartPort() int {
	return integration.MSPPort.StartPortForNode()
}
