/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package main

// idemixgen is a command line tool that generates the CA's keys and
// generates MSP configs for siging and for verification
// This tool can be used to setup the peers and CA to support
// the Identity Mixer MSP

import (
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/golang/protobuf/proto"
	"github.com/hxx258456/ccgo/sm2"
	"github.com/hxx258456/ccgo/x509"
	"github.com/hxx258456/fabric/bccsp/utils"
	"github.com/hxx258456/fabric/common/tools/idemixgen/idemixca"
	"github.com/hxx258456/fabric/common/tools/idemixgen/metadata"
	"github.com/hxx258456/fabric/idemix"
	"github.com/hxx258456/fabric/msp"
	"github.com/pkg/errors"
	"gopkg.in/alecthomas/kingpin.v2"
)

const (
	IdemixDirIssuer             = "ca"
	IdemixConfigIssuerSecretKey = "IssuerSecretKey"
	IdemixConfigRevocationKey   = "RevocationKey"
)

// command line flags
var (
	app = kingpin.New("idemixgen", "Utility for generating key material to be used with the Identity Mixer MSP in Hyperledger Fabric")

	outputDir = app.Flag("output", "The output directory in which to place artifacts").Default("idemix-config").String()

	genIssuerKey            = app.Command("ca-keygen", "Generate CA key material")
	genSignerConfig         = app.Command("signerconfig", "Generate a default signer for this Idemix MSP")
	genCAInput              = genSignerConfig.Flag("ca-input", "The folder where CA's secrets are stored").String()
	genCredOU               = genSignerConfig.Flag("org-unit", "The Organizational Unit of the default signer").Short('u').String()
	genCredIsAdmin          = genSignerConfig.Flag("admin", "Make the default signer admin").Short('a').Bool()
	genCredEnrollmentId     = genSignerConfig.Flag("enrollmentId", "The enrollment id of the default signer").Short('e').String()
	genCredRevocationHandle = genSignerConfig.Flag("revocationHandle", "The handle used to revoke this signer").Short('r').Int()

	version = app.Command("version", "Show version information")
)

func main() {
	app.HelpFlag.Short('h')

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {

	case genIssuerKey.FullCommand():
		isk, ipk, err := idemixca.GenerateIssuerKey()
		handleError(err)

		revocationKey, err := idemix.GenerateLongTermRevocationKey()
		handleError(err)
		//TODO
		encodedRevocationSK, err := utils.PrivateKeyToDER(revocationKey)
		handleError(err)
		pemEncodedRevocationSK := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: encodedRevocationSK})
		handleError(err)
		encodedRevocationPK, err := x509.MarshalPKIXPublicKey(revocationKey.Public())
		handleError(err)
		pemEncodedRevocationPK := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: encodedRevocationPK})

		// Prevent overwriting the existing key
		path := filepath.Join(*outputDir, IdemixDirIssuer)
		checkDirectoryNotExists(path, fmt.Sprintf("Directory %s already exists", path))

		path = filepath.Join(*outputDir, msp.IdemixConfigDirMsp)
		checkDirectoryNotExists(path, fmt.Sprintf("Directory %s already exists", path))

		// write private and public keys to the file
		handleError(os.MkdirAll(filepath.Join(*outputDir, IdemixDirIssuer), 0770))
		handleError(os.MkdirAll(filepath.Join(*outputDir, msp.IdemixConfigDirMsp), 0770))
		writeFile(filepath.Join(*outputDir, IdemixDirIssuer, IdemixConfigIssuerSecretKey), isk)
		writeFile(filepath.Join(*outputDir, IdemixDirIssuer, IdemixConfigRevocationKey), pemEncodedRevocationSK)
		writeFile(filepath.Join(*outputDir, IdemixDirIssuer, msp.IdemixConfigFileIssuerPublicKey), ipk)
		writeFile(filepath.Join(*outputDir, msp.IdemixConfigDirMsp, msp.IdemixConfigFileRevocationPublicKey), pemEncodedRevocationPK)
		writeFile(filepath.Join(*outputDir, msp.IdemixConfigDirMsp, msp.IdemixConfigFileIssuerPublicKey), ipk)

	case genSignerConfig.FullCommand():
		roleMask := 0
		if *genCredIsAdmin {
			roleMask = msp.GetRoleMaskFromIdemixRole(msp.ADMIN)
		} else {
			roleMask = msp.GetRoleMaskFromIdemixRole(msp.MEMBER)
		}
		if *genCAInput == "" {
			genCAInput = outputDir
		}
		ipk, ipkRaw := readIssuerKey()
		rsk := readRevocationKey()
		rpk := readRevocationPublicKey()

		config, err := idemixca.GenerateSignerConfig(
			roleMask,
			*genCredOU,
			*genCredEnrollmentId,
			*genCredRevocationHandle,
			ipk, rsk,
		)
		handleError(err)

		path := filepath.Join(*outputDir, msp.IdemixConfigDirUser)
		checkDirectoryNotExists(path, fmt.Sprintf("This MSP config already contains a directory \"%s\"", path))

		// Write config to file
		handleError(os.MkdirAll(filepath.Join(*outputDir, msp.IdemixConfigDirUser), 0770))
		writeFile(filepath.Join(*outputDir, msp.IdemixConfigDirUser, msp.IdemixConfigFileSigner), config)

		// Write CA public info in case genCAInput != outputDir
		if *genCAInput != *outputDir {
			handleError(os.MkdirAll(filepath.Join(*outputDir, msp.IdemixConfigDirMsp), 0770))
			writeFile(filepath.Join(*outputDir, msp.IdemixConfigDirMsp, msp.IdemixConfigFileRevocationPublicKey), rpk)
			writeFile(filepath.Join(*outputDir, msp.IdemixConfigDirMsp, msp.IdemixConfigFileIssuerPublicKey), ipkRaw)
		}

	case version.FullCommand():
		printVersion()

	}
}

func printVersion() {
	fmt.Println(metadata.GetVersionInfo())
}

// writeFile writes bytes to a file and panics in case of an error
func writeFile(path string, contents []byte) {
	handleError(ioutil.WriteFile(path, contents, 0640))
}

// readIssuerKey reads the issuer key from the current directory
func readIssuerKey() (*idemix.IssuerKey, []byte) {
	path := filepath.Join(*genCAInput, IdemixDirIssuer, IdemixConfigIssuerSecretKey)
	isk, err := ioutil.ReadFile(path)
	if err != nil {
		handleError(errors.Wrapf(err, "failed to open issuer secret key file: %s", path))
	}
	path = filepath.Join(*genCAInput, IdemixDirIssuer, msp.IdemixConfigFileIssuerPublicKey)
	ipkBytes, err := ioutil.ReadFile(path)
	if err != nil {
		handleError(errors.Wrapf(err, "failed to open issuer public key file: %s", path))
	}
	ipk := &idemix.IssuerPublicKey{}
	handleError(proto.Unmarshal(ipkBytes, ipk))
	key := &idemix.IssuerKey{Isk: isk, Ipk: ipk}

	return key, ipkBytes
}

func readRevocationKey() *sm2.PrivateKey {
	path := filepath.Join(*genCAInput, IdemixDirIssuer, IdemixConfigRevocationKey)
	keyBytes, err := ioutil.ReadFile(path)
	if err != nil {
		handleError(errors.Wrapf(err, "failed to open revocation secret key file: %s", path))
	}

	// block, _ := pem.Decode(keyBytes)
	// if block == nil {
	// 	handleError(errors.Errorf("failed to decode SM2 private key"))
	// }
	// key, err := x509.ParseSm2PrivateKey(block.Bytes)
	key, err := utils.PEMToSm2PrivateKey(keyBytes, nil)
	handleError(err)

	return key
}

func readRevocationPublicKey() []byte {
	path := filepath.Join(*genCAInput, msp.IdemixConfigDirMsp, msp.IdemixConfigFileRevocationPublicKey)
	keyBytes, err := ioutil.ReadFile(path)
	if err != nil {
		handleError(errors.Wrapf(err, "failed to open revocation secret key file: %s", path))
	}

	return keyBytes
}

// checkDirectoryNotExists checks whether a directory with the given path already exists and exits if this is the case
func checkDirectoryNotExists(path string, errorMessage string) {
	_, err := os.Stat(path)
	if err == nil {
		handleError(errors.New(errorMessage))
	}
}

func handleError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
