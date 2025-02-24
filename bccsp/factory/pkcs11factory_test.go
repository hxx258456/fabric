//go:build pkcs11
// +build pkcs11

/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package factory

/*
bccsp/factory/pkcs11factory_test.go 国密对应后废弃
*/

// import (
// 	"crypto/sha256"
// 	"encoding/hex"
// 	"testing"

// 	"github.com/hxx258456/fabric/bccsp/pkcs11"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/require"
// )

// func TestPKCS11FactoryName(t *testing.T) {
// 	f := &PKCS11Factory{}
// 	assert.Equal(t, f.Name(), PKCS11BasedFactoryName)
// }

// func TestPKCS11FactoryGetInvalidArgs(t *testing.T) {
// 	f := &PKCS11Factory{}

// 	_, err := f.Get(nil)
// 	assert.Error(t, err, "Invalid config. It must not be nil.")

// 	_, err = f.Get(&FactoryOpts{})
// 	assert.Error(t, err, "Invalid config. It must not be nil.")

// 	opts := &FactoryOpts{
// 		Pkcs11Opts: &pkcs11.PKCS11Opts{},
// 	}
// 	_, err = f.Get(opts)
// 	assert.Error(t, err, "CSP:500 - Failed initializing configuration at [0,]")
// }

// func TestPKCS11FactoryGet(t *testing.T) {
// 	f := &PKCS11Factory{}
// 	lib, pin, label := pkcs11.FindPKCS11Lib()

// 	opts := &FactoryOpts{
// 		Pkcs11Opts: &pkcs11.PKCS11Opts{
// 			SecLevel:   256,
// 			HashFamily: "SHA2",
// 			Library:    lib,
// 			Pin:        pin,
// 			Label:      label,
// 		},
// 	}
// 	csp, err := f.Get(opts)
// 	assert.NoError(t, err)
// 	assert.NotNil(t, csp)

// 	opts = &FactoryOpts{
// 		Pkcs11Opts: &pkcs11.PKCS11Opts{
// 			SecLevel:   256,
// 			HashFamily: "SHA2",
// 			Library:    lib,
// 			Pin:        pin,
// 			Label:      label,
// 		},
// 	}
// 	csp, err = f.Get(opts)
// 	assert.NoError(t, err)
// 	assert.NotNil(t, csp)

// 	opts = &FactoryOpts{
// 		Pkcs11Opts: &pkcs11.PKCS11Opts{
// 			SecLevel:   256,
// 			HashFamily: "SHA2",
// 			Ephemeral:  true,
// 			Library:    lib,
// 			Pin:        pin,
// 			Label:      label,
// 		},
// 	}
// 	csp, err = f.Get(opts)
// 	assert.NoError(t, err)
// 	assert.NotNil(t, csp)
// }

// func TestPKCS11FactoryGetEmptyKeyStorePath(t *testing.T) {
// 	f := &PKCS11Factory{}
// 	lib, pin, label := pkcs11.FindPKCS11Lib()

// 	opts := &FactoryOpts{
// 		Pkcs11Opts: &pkcs11.PKCS11Opts{
// 			SecLevel:   256,
// 			HashFamily: "SHA2",
// 			Library:    lib,
// 			Pin:        pin,
// 			Label:      label,
// 		},
// 	}
// 	csp, err := f.Get(opts)
// 	assert.NoError(t, err)
// 	assert.NotNil(t, csp)

// 	opts = &FactoryOpts{
// 		Pkcs11Opts: &pkcs11.PKCS11Opts{
// 			SecLevel:   256,
// 			HashFamily: "SHA2",
// 			Library:    lib,
// 			Pin:        pin,
// 			Label:      label,
// 		},
// 	}
// 	csp, err = f.Get(opts)
// 	assert.NoError(t, err)
// 	assert.NotNil(t, csp)
// }

// func TestSKIMapper(t *testing.T) {
// 	inputSKI := sha256.New().Sum([]byte("some-ski"))
// 	tests := []struct {
// 		name     string
// 		altID    string
// 		keyIDs   map[string]string
// 		expected []byte
// 	}{
// 		{name: "DefaultBehavior", expected: inputSKI},
// 		{name: "AltIDOnly", altID: "alternate-ID", expected: []byte("alternate-ID")},
// 		{name: "MapEntry", keyIDs: map[string]string{hex.EncodeToString(inputSKI): "mapped-id"}, expected: []byte("mapped-id")},
// 		{name: "AltIDAsDefault", altID: "alternate-ID", keyIDs: map[string]string{"another-ski": "another-id"}, expected: []byte("alternate-ID")},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			options := defaultOptions()
// 			options.AltID = tt.altID
// 			for k, v := range tt.keyIDs {
// 				options.KeyIDs = append(options.KeyIDs, pkcs11.KeyIDMapping{SKI: k, ID: v})
// 			}

// 			mapper := skiMapper(*options)
// 			result := mapper(inputSKI)
// 			require.Equal(t, tt.expected, result, "got %x, want %x", result, tt.expected)
// 		})
// 	}
// }

// func defaultOptions() *pkcs11.PKCS11Opts {
// 	lib, pin, label := pkcs11.FindPKCS11Lib()
// 	return &pkcs11.PKCS11Opts{
// 		SecLevel:   256,
// 		HashFamily: "SHA2",
// 		Library:    lib,
// 		Pin:        pin,
// 		Label:      label,
// 	}
// }
