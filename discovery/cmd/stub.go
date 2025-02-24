/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package discovery

import (
	"context"

	"github.com/hxx258456/fabric-protos-go-cc/discovery"
	"github.com/hxx258456/fabric/cmd/common"
	"github.com/hxx258456/fabric/cmd/common/comm"
	"github.com/hxx258456/fabric/cmd/common/signer"
	discoveryclient "github.com/hxx258456/fabric/discovery/client"
	"github.com/hxx258456/fabric/protoutil"
	"github.com/pkg/errors"
)

//go:generate mockery -dir . -name LocalResponse -case underscore -output mocks/

// LocalResponse is the local interface used to generate mocks for foreign interface.
type LocalResponse interface {
	discoveryclient.LocalResponse
}

//go:generate mockery -dir . -name ChannelResponse -case underscore -output mocks/

// ChannelResponse is the local interface used to generate mocks for foreign interface.
type ChannelResponse interface {
	discoveryclient.ChannelResponse
}

//go:generate mockery -dir . -name ServiceResponse -case underscore -output mocks/

// ServiceResponse represents a response sent from the discovery service
type ServiceResponse interface {
	// ForChannel returns a ChannelResponse in the context of a given channel
	ForChannel(string) discoveryclient.ChannelResponse

	// ForLocal returns a LocalResponse in the context of no channel
	ForLocal() discoveryclient.LocalResponse

	// Raw returns the raw response from the server
	Raw() *discovery.Response
}

type response struct {
	raw *discovery.Response
	discoveryclient.Response
}

func (r *response) Raw() *discovery.Response {
	return r.raw
}

// ClientStub is a stub that communicates with the discovery service
// using the discovery client implementation
type ClientStub struct{}

// Send sends the request, and receives a response
func (stub *ClientStub) Send(server string, conf common.Config, req *discoveryclient.Request) (ServiceResponse, error) {
	comm, err := comm.NewClient(conf.TLSConfig)
	if err != nil {
		return nil, err
	}
	signer, err := signer.NewSigner(conf.SignerConfig)
	if err != nil {
		return nil, err
	}
	timeout, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	disc := discoveryclient.NewClient(comm.NewDialer(server), signer.Sign, 0)

	resp, err := disc.Send(timeout, req, &discovery.AuthInfo{
		ClientIdentity:    signer.Creator,
		ClientTlsCertHash: comm.TLSCertHash,
	})
	if err != nil {
		return nil, errors.Errorf("failed connecting to %s: %v", server, err)
	}
	return &response{
		Response: resp,
	}, nil
}

// RawStub is a stub that communicates with the discovery service
// without any intermediary.
type RawStub struct{}

// Send sends the request, and receives a response
func (stub *RawStub) Send(server string, conf common.Config, req *discoveryclient.Request) (ServiceResponse, error) {
	comm, err := comm.NewClient(conf.TLSConfig)
	if err != nil {
		return nil, err
	}
	signer, err := signer.NewSigner(conf.SignerConfig)
	if err != nil {
		return nil, err
	}
	timeout, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	req.Authentication = &discovery.AuthInfo{
		ClientIdentity:    signer.Creator,
		ClientTlsCertHash: comm.TLSCertHash,
	}

	payload := protoutil.MarshalOrPanic(req.Request)
	sig, err := signer.Sign(payload)
	if err != nil {
		return nil, err
	}

	cc, err := comm.NewDialer(server)()
	if err != nil {
		return nil, err
	}
	resp, err := discovery.NewDiscoveryClient(cc).Discover(timeout, &discovery.SignedRequest{
		Payload:   payload,
		Signature: sig,
	})
	if err != nil {
		return nil, err
	}

	return &response{
		raw: resp,
	}, nil
}
