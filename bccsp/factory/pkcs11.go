//go:build pkcs11
// +build pkcs11

/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package factory

/*
bccsp/factory/pkcs11.go 国密对应后废弃
*/

// import (
// 	"github.com/hxx258456/fabric/bccsp"
// 	"github.com/hxx258456/fabric/bccsp/pkcs11"
// 	"github.com/pkg/errors"
// )

// /*
// bccsp/factory/pkcs11.go 为 PKCS11Factory 提供 FactoryOpts 及相关的函数
// 需要添加编译条件: `pkcs11`
// */

// const pkcs11Enabled = false

// // FactoryOpts holds configuration information used to initialize factory implementations
// type FactoryOpts struct {
// 	ProviderName string             `mapstructure:"default" json:"default" yaml:"Default"`
// 	SwOpts       *SwOpts            `mapstructure:"SW,omitempty" json:"SW,omitempty" yaml:"SwOpts"`
// 	Pkcs11Opts   *pkcs11.PKCS11Opts `mapstructure:"PKCS11,omitempty" json:"PKCS11,omitempty" yaml:"PKCS11"`
// }

// // InitFactories must be called before using factory interfaces
// // It is acceptable to call with config = nil, in which case
// // some defaults will get used
// // Error is returned only if defaultBCCSP cannot be found
// func InitFactories(config *FactoryOpts) error {
// 	factoriesInitOnce.Do(func() {
// 		factoriesInitError = initFactories(config)
// 	})

// 	return factoriesInitError
// }

// func initFactories(config *FactoryOpts) error {
// 	// Take some precautions on default opts
// 	if config == nil {
// 		config = GetDefaultOpts()
// 	}

// 	if config.ProviderName == "" {
// 		config.ProviderName = "SW"
// 	}

// 	if config.SwOpts == nil {
// 		config.SwOpts = GetDefaultOpts().SwOpts
// 	}

// 	// // Software-Based BCCSP
// 	// if config.ProviderName == "GM" && config.SwOpts != nil {
// 	// 	f := &GMFactory{}
// 	// 	var err error
// 	// 	defaultBCCSP, err = initBCCSP(f, config)
// 	// 	if err != nil {
// 	// 		return errors.Wrap(err, "Failed initializing SW.BCCSP")
// 	// 	}
// 	// }

// 	// Software-Based BCCSP
// 	if config.ProviderName == "SW" && config.SwOpts != nil {
// 		f := &SWFactory{}
// 		var err error
// 		defaultBCCSP, err = initBCCSP(f, config)
// 		if err != nil {
// 			return errors.Wrap(err, "Failed initializing SW.BCCSP")
// 		}
// 	}

// 	// PKCS11-Based BCCSP
// 	if config.ProviderName == "PKCS11" && config.Pkcs11Opts != nil {
// 		f := &PKCS11Factory{}
// 		var err error
// 		defaultBCCSP, err = initBCCSP(f, config)
// 		if err != nil {
// 			return errors.Wrapf(err, "Failed initializing PKCS11.BCCSP")
// 		}
// 	}

// 	if defaultBCCSP == nil {
// 		return errors.Errorf("Could not find default `%s` BCCSP", config.ProviderName)
// 	}

// 	return nil
// }

// // GetBCCSPFromOpts returns a BCCSP created according to the options passed in input.
// func GetBCCSPFromOpts(config *FactoryOpts) (bccsp.BCCSP, error) {
// 	var f BCCSPFactory
// 	switch config.ProviderName {
// 	// case "GM":
// 	// 	f = &GMFactory{}
// 	case "SW":
// 		f = &SWFactory{}
// 	case "PKCS11":
// 		f = &PKCS11Factory{}
// 	default:
// 		return nil, errors.Errorf("Could not find BCCSP, no '%s' provider", config.ProviderName)
// 	}

// 	csp, err := f.Get(config)
// 	if err != nil {
// 		return nil, errors.Wrapf(err, "Could not initialize BCCSP %s", f.Name())
// 	}
// 	return csp, nil
// }
