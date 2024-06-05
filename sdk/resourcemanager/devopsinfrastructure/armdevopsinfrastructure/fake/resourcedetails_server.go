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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devopsinfrastructure/armdevopsinfrastructure"
	"net/http"
	"net/url"
	"regexp"
)

// ResourceDetailsServer is a fake server for instances of the armdevopsinfrastructure.ResourceDetailsClient type.
type ResourceDetailsServer struct {
	// NewListByPoolPager is the fake for method ResourceDetailsClient.NewListByPoolPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByPoolPager func(resourceGroupName string, poolName string, options *armdevopsinfrastructure.ResourceDetailsClientListByPoolOptions) (resp azfake.PagerResponder[armdevopsinfrastructure.ResourceDetailsClientListByPoolResponse])
}

// NewResourceDetailsServerTransport creates a new instance of ResourceDetailsServerTransport with the provided implementation.
// The returned ResourceDetailsServerTransport instance is connected to an instance of armdevopsinfrastructure.ResourceDetailsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewResourceDetailsServerTransport(srv *ResourceDetailsServer) *ResourceDetailsServerTransport {
	return &ResourceDetailsServerTransport{
		srv:                srv,
		newListByPoolPager: newTracker[azfake.PagerResponder[armdevopsinfrastructure.ResourceDetailsClientListByPoolResponse]](),
	}
}

// ResourceDetailsServerTransport connects instances of armdevopsinfrastructure.ResourceDetailsClient to instances of ResourceDetailsServer.
// Don't use this type directly, use NewResourceDetailsServerTransport instead.
type ResourceDetailsServerTransport struct {
	srv                *ResourceDetailsServer
	newListByPoolPager *tracker[azfake.PagerResponder[armdevopsinfrastructure.ResourceDetailsClientListByPoolResponse]]
}

// Do implements the policy.Transporter interface for ResourceDetailsServerTransport.
func (r *ResourceDetailsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ResourceDetailsClient.NewListByPoolPager":
		resp, err = r.dispatchNewListByPoolPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *ResourceDetailsServerTransport) dispatchNewListByPoolPager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListByPoolPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByPoolPager not implemented")}
	}
	newListByPoolPager := r.newListByPoolPager.get(req)
	if newListByPoolPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DevOpsInfrastructure/pools/(?P<poolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resources`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		poolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("poolName")])
		if err != nil {
			return nil, err
		}
		resp := r.srv.NewListByPoolPager(resourceGroupNameParam, poolNameParam, nil)
		newListByPoolPager = &resp
		r.newListByPoolPager.add(req, newListByPoolPager)
		server.PagerResponderInjectNextLinks(newListByPoolPager, req, func(page *armdevopsinfrastructure.ResourceDetailsClientListByPoolResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByPoolPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		r.newListByPoolPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByPoolPager) {
		r.newListByPoolPager.remove(req)
	}
	return resp, nil
}
