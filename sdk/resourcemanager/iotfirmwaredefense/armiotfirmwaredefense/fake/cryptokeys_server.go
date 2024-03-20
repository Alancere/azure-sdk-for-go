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
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotfirmwaredefense/armiotfirmwaredefense"
	"net/http"
	"net/url"
	"regexp"
)

// CryptoKeysServer is a fake server for instances of the armiotfirmwaredefense.CryptoKeysClient type.
type CryptoKeysServer struct {
	// NewListByFirmwarePager is the fake for method CryptoKeysClient.NewListByFirmwarePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByFirmwarePager func(resourceGroupName string, workspaceName string, firmwareID string, options *armiotfirmwaredefense.CryptoKeysClientListByFirmwareOptions) (resp azfake.PagerResponder[armiotfirmwaredefense.CryptoKeysClientListByFirmwareResponse])
}

// NewCryptoKeysServerTransport creates a new instance of CryptoKeysServerTransport with the provided implementation.
// The returned CryptoKeysServerTransport instance is connected to an instance of armiotfirmwaredefense.CryptoKeysClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewCryptoKeysServerTransport(srv *CryptoKeysServer) *CryptoKeysServerTransport {
	return &CryptoKeysServerTransport{
		srv:                    srv,
		newListByFirmwarePager: newTracker[azfake.PagerResponder[armiotfirmwaredefense.CryptoKeysClientListByFirmwareResponse]](),
	}
}

// CryptoKeysServerTransport connects instances of armiotfirmwaredefense.CryptoKeysClient to instances of CryptoKeysServer.
// Don't use this type directly, use NewCryptoKeysServerTransport instead.
type CryptoKeysServerTransport struct {
	srv                    *CryptoKeysServer
	newListByFirmwarePager *tracker[azfake.PagerResponder[armiotfirmwaredefense.CryptoKeysClientListByFirmwareResponse]]
}

// Do implements the policy.Transporter interface for CryptoKeysServerTransport.
func (c *CryptoKeysServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "CryptoKeysClient.NewListByFirmwarePager":
		resp, err = c.dispatchNewListByFirmwarePager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *CryptoKeysServerTransport) dispatchNewListByFirmwarePager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListByFirmwarePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByFirmwarePager not implemented")}
	}
	newListByFirmwarePager := c.newListByFirmwarePager.get(req)
	if newListByFirmwarePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.IoTFirmwareDefense/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/firmwares/(?P<firmwareId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/cryptoKeys`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
		if err != nil {
			return nil, err
		}
		firmwareIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("firmwareId")])
		if err != nil {
			return nil, err
		}
		resp := c.srv.NewListByFirmwarePager(resourceGroupNameParam, workspaceNameParam, firmwareIDParam, nil)
		newListByFirmwarePager = &resp
		c.newListByFirmwarePager.add(req, newListByFirmwarePager)
		server.PagerResponderInjectNextLinks(newListByFirmwarePager, req, func(page *armiotfirmwaredefense.CryptoKeysClientListByFirmwareResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByFirmwarePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		c.newListByFirmwarePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByFirmwarePager) {
		c.newListByFirmwarePager.remove(req)
	}
	return resp, nil
}
