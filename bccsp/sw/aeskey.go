/*
Copyright IBM Corp. 2016 All Rights Reserved.

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

/*
bccsp/sw/aeskey.go 国密对应后废弃
*/

// import (
// 	"crypto/sha256"
// 	"errors"

// 	"github.com/hxx258456/fabric/bccsp"
// )

// /*
// bccsp/sw/aeskey.go 定义AES密钥结构体，并实现`bccsp.Key`(bccsp/bccsp.go)接口
// */

// type AESPrivateKey struct {
// 	privKey    []byte
// 	exportable bool
// }

// // Bytes converts this key to its byte representation,
// // if this operation is allowed.
// func (k *AESPrivateKey) Bytes() (raw []byte, err error) {
// 	if k.exportable {
// 		return k.privKey, nil
// 	}

// 	return nil, errors.New("not supported")
// }

// // SKI returns the subject key identifier of this key.
// func (k *AESPrivateKey) SKI() (ski []byte) {
// 	hash := sha256.New()
// 	hash.Write([]byte{0x01})
// 	hash.Write(k.privKey)
// 	return hash.Sum(nil)
// }

// // Symmetric returns true if this key is a symmetric key,
// // false if this key is asymmetric
// func (k *AESPrivateKey) Symmetric() bool {
// 	return true
// }

// // Private returns true if this key is a private key,
// // false otherwise.
// func (k *AESPrivateKey) Private() bool {
// 	return true
// }

// // PublicKey returns the corresponding public key part of an asymmetric public/private key pair.
// // This method returns an error in symmetric key schemes.
// func (k *AESPrivateKey) PublicKey() (bccsp.Key, error) {
// 	return nil, errors.New("cannot call this method on a symmetric key")
// }

// func (k *AESPrivateKey) InsideKey() interface{} {
// 	return k.privKey
// }
