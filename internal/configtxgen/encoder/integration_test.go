/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package encoder_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	cb "github.com/hxx258456/fabric-protos-go-cc/common"
	"github.com/hxx258456/fabric/bccsp/sw"
	"github.com/hxx258456/fabric/common/channelconfig"
	"github.com/hxx258456/fabric/core/config/configtest"
	"github.com/hxx258456/fabric/internal/configtxgen/encoder"
	"github.com/hxx258456/fabric/internal/configtxgen/genesisconfig"

	"github.com/pkg/errors"
)

func hasModPolicySet(groupName string, cg *cb.ConfigGroup) error {
	if cg.ModPolicy == "" {
		return errors.Errorf("group %s has empty mod_policy", groupName)
	}

	for valueName, value := range cg.Values {
		if value.ModPolicy == "" {
			return errors.Errorf("group %s has value %s with empty mod_policy", groupName, valueName)
		}
	}

	for policyName, policy := range cg.Policies {
		if policy.ModPolicy == "" {
			return errors.Errorf("group %s has policy %s with empty mod_policy", groupName, policyName)
		}
	}

	for groupName, group := range cg.Groups {
		err := hasModPolicySet(groupName, group)
		if err != nil {
			return errors.WithMessagef(err, "missing sub-mod_policy for group %s", groupName)
		}
	}

	return nil
}

var _ = Describe("Integration", func() {
	DescribeTable("successfully parses the profile",
		func(profile string) {
			config := genesisconfig.Load(profile, configtest.GetDevConfigDir())
			group, err := encoder.NewChannelGroup(config)
			Expect(err).NotTo(HaveOccurred())

			cryptoProvider, err := sw.NewDefaultSecurityLevelWithKeystore(sw.NewDummyKeyStore())
			Expect(err).NotTo(HaveOccurred())
			_, err = channelconfig.NewBundle("test", &cb.Config{
				ChannelGroup: group,
			}, cryptoProvider)
			Expect(err).NotTo(HaveOccurred())

			err = hasModPolicySet("Channel", group)
			Expect(err).NotTo(HaveOccurred())
		},
		Entry("Sample Insecure Solo Profile", genesisconfig.SampleInsecureSoloProfile),
		Entry("Sample Single MSP Solo Profile", genesisconfig.SampleSingleMSPSoloProfile),
		Entry("Sample DevMode Solo Profile", genesisconfig.SampleDevModeSoloProfile),
		Entry("Sample Insecure Kafka Profile", genesisconfig.SampleInsecureKafkaProfile),
		Entry("Sample Single MSP Kafka Profile", genesisconfig.SampleSingleMSPKafkaProfile),
		Entry("Sample DevMode Kafka Profile", genesisconfig.SampleDevModeKafkaProfile),
	)
})
