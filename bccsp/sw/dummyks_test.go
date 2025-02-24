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
	"testing"

	"github.com/hxx258456/fabric/bccsp/mocks"
	"github.com/stretchr/testify/assert"
)

func TestNewDummyKeyStore(t *testing.T) {
	t.Parallel()

	ks := NewDummyKeyStore()
	assert.NotNil(t, ks)
}

func TestDummyKeyStore_GetKey(t *testing.T) {
	t.Parallel()

	ks := NewDummyKeyStore()
	_, err := ks.GetKey([]byte{0, 1, 2, 3, 4})
	assert.Error(t, err)
}

func TestDummyKeyStore_ReadOnly(t *testing.T) {
	t.Parallel()

	ks := NewDummyKeyStore()
	assert.True(t, ks.ReadOnly())
}

func TestDummyKeyStore_StoreKey(t *testing.T) {
	t.Parallel()

	ks := NewDummyKeyStore()
	err := ks.StoreKey(&mocks.MockKey{})
	assert.Error(t, err)
}
