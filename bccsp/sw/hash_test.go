/*
Copyright IBM Corp. 2017 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package sw

import (
	"errors"
	"reflect"
	"testing"

	"github.com/hxx258456/ccgo/sm3"
	mocks2 "github.com/hxx258456/fabric/bccsp/mocks"
	"github.com/hxx258456/fabric/bccsp/sw/mocks"
	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	t.Parallel()

	expectetMsg := []byte{1, 2, 3, 4}
	expectedOpts := &mocks2.HashOpts{}
	expectetValue := []byte{1, 2, 3, 4, 5}
	expectedErr := errors.New("Expected Error")

	hashers := make(map[reflect.Type]Hasher)
	hashers[reflect.TypeOf(&mocks2.HashOpts{})] = &mocks.Hasher{
		MsgArg:  expectetMsg,
		OptsArg: expectedOpts,
		Value:   expectetValue,
		Err:     nil,
	}
	csp := CSP{Hashers: hashers}
	value, err := csp.Hash(expectetMsg, expectedOpts)
	assert.Equal(t, expectetValue, value)
	assert.Nil(t, err)

	hashers = make(map[reflect.Type]Hasher)
	hashers[reflect.TypeOf(&mocks2.HashOpts{})] = &mocks.Hasher{
		MsgArg:  expectetMsg,
		OptsArg: expectedOpts,
		Value:   nil,
		Err:     expectedErr,
	}
	csp = CSP{Hashers: hashers}
	value, err = csp.Hash(expectetMsg, expectedOpts)
	assert.Nil(t, value)
	assert.Contains(t, err.Error(), expectedErr.Error())
}

func TestGetHash(t *testing.T) {
	t.Parallel()

	expectedOpts := &mocks2.HashOpts{}
	expectetValue := sm3.New()
	expectedErr := errors.New("Expected Error")

	hashers := make(map[reflect.Type]Hasher)
	hashers[reflect.TypeOf(&mocks2.HashOpts{})] = &mocks.Hasher{
		OptsArg:   expectedOpts,
		ValueHash: expectetValue,
		Err:       nil,
	}
	csp := CSP{Hashers: hashers}
	value, err := csp.GetHash(expectedOpts)
	assert.Equal(t, expectetValue, value)
	assert.Nil(t, err)

	hashers = make(map[reflect.Type]Hasher)
	hashers[reflect.TypeOf(&mocks2.HashOpts{})] = &mocks.Hasher{
		OptsArg:   expectedOpts,
		ValueHash: expectetValue,
		Err:       expectedErr,
	}
	csp = CSP{Hashers: hashers}
	value, err = csp.GetHash(expectedOpts)
	assert.Nil(t, value)
	assert.Contains(t, err.Error(), expectedErr.Error())
}

func TestHasher(t *testing.T) {
	t.Parallel()

	hasher := &hasher{hash: sm3.New}

	msg := []byte("Hello World")
	out, err := hasher.Hash(msg, nil)
	assert.NoError(t, err)
	h := sm3.New()
	h.Write(msg)
	out2 := h.Sum(nil)
	assert.Equal(t, out, out2)

	hf, err := hasher.GetHash(nil)
	assert.NoError(t, err)
	assert.Equal(t, hf, sm3.New())
}
