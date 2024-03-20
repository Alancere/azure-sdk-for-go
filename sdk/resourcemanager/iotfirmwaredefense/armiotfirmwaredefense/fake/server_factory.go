//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
	"sync"
)

// ServerFactory is a fake server for instances of the armiotfirmwaredefense.ClientFactory type.
type ServerFactory struct {
	BinaryHardeningServer    BinaryHardeningServer
	CryptoCertificatesServer CryptoCertificatesServer
	CryptoKeysServer         CryptoKeysServer
	CvesServer               CvesServer
	FirmwaresServer          FirmwaresServer
	OperationsServer         OperationsServer
	PasswordHashesServer     PasswordHashesServer
	SbomComponentsServer     SbomComponentsServer
	SummariesServer          SummariesServer
	WorkspacesServer         WorkspacesServer
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of armiotfirmwaredefense.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of armiotfirmwaredefense.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                        *ServerFactory
	trMu                       sync.Mutex
	trBinaryHardeningServer    *BinaryHardeningServerTransport
	trCryptoCertificatesServer *CryptoCertificatesServerTransport
	trCryptoKeysServer         *CryptoKeysServerTransport
	trCvesServer               *CvesServerTransport
	trFirmwaresServer          *FirmwaresServerTransport
	trOperationsServer         *OperationsServerTransport
	trPasswordHashesServer     *PasswordHashesServerTransport
	trSbomComponentsServer     *SbomComponentsServerTransport
	trSummariesServer          *SummariesServerTransport
	trWorkspacesServer         *WorkspacesServerTransport
}

// Do implements the policy.Transporter interface for ServerFactoryTransport.
func (s *ServerFactoryTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	client := method[:strings.Index(method, ".")]
	var resp *http.Response
	var err error

	switch client {
	case "BinaryHardeningClient":
		initServer(s, &s.trBinaryHardeningServer, func() *BinaryHardeningServerTransport {
			return NewBinaryHardeningServerTransport(&s.srv.BinaryHardeningServer)
		})
		resp, err = s.trBinaryHardeningServer.Do(req)
	case "CryptoCertificatesClient":
		initServer(s, &s.trCryptoCertificatesServer, func() *CryptoCertificatesServerTransport {
			return NewCryptoCertificatesServerTransport(&s.srv.CryptoCertificatesServer)
		})
		resp, err = s.trCryptoCertificatesServer.Do(req)
	case "CryptoKeysClient":
		initServer(s, &s.trCryptoKeysServer, func() *CryptoKeysServerTransport { return NewCryptoKeysServerTransport(&s.srv.CryptoKeysServer) })
		resp, err = s.trCryptoKeysServer.Do(req)
	case "CvesClient":
		initServer(s, &s.trCvesServer, func() *CvesServerTransport { return NewCvesServerTransport(&s.srv.CvesServer) })
		resp, err = s.trCvesServer.Do(req)
	case "FirmwaresClient":
		initServer(s, &s.trFirmwaresServer, func() *FirmwaresServerTransport { return NewFirmwaresServerTransport(&s.srv.FirmwaresServer) })
		resp, err = s.trFirmwaresServer.Do(req)
	case "OperationsClient":
		initServer(s, &s.trOperationsServer, func() *OperationsServerTransport { return NewOperationsServerTransport(&s.srv.OperationsServer) })
		resp, err = s.trOperationsServer.Do(req)
	case "PasswordHashesClient":
		initServer(s, &s.trPasswordHashesServer, func() *PasswordHashesServerTransport {
			return NewPasswordHashesServerTransport(&s.srv.PasswordHashesServer)
		})
		resp, err = s.trPasswordHashesServer.Do(req)
	case "SbomComponentsClient":
		initServer(s, &s.trSbomComponentsServer, func() *SbomComponentsServerTransport {
			return NewSbomComponentsServerTransport(&s.srv.SbomComponentsServer)
		})
		resp, err = s.trSbomComponentsServer.Do(req)
	case "SummariesClient":
		initServer(s, &s.trSummariesServer, func() *SummariesServerTransport { return NewSummariesServerTransport(&s.srv.SummariesServer) })
		resp, err = s.trSummariesServer.Do(req)
	case "WorkspacesClient":
		initServer(s, &s.trWorkspacesServer, func() *WorkspacesServerTransport { return NewWorkspacesServerTransport(&s.srv.WorkspacesServer) })
		resp, err = s.trWorkspacesServer.Do(req)
	default:
		err = fmt.Errorf("unhandled client %s", client)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func initServer[T any](s *ServerFactoryTransport, dst **T, src func() *T) {
	s.trMu.Lock()
	if *dst == nil {
		*dst = src()
	}
	s.trMu.Unlock()
}
