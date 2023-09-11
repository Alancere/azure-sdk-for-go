//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5"
	"net/http"
	"net/url"
	"regexp"
)

// CommunityGalleryImageVersionsServer is a fake server for instances of the armcompute.CommunityGalleryImageVersionsClient type.
type CommunityGalleryImageVersionsServer struct{
	// Get is the fake for method CommunityGalleryImageVersionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, location string, publicGalleryName string, galleryImageName string, galleryImageVersionName string, options *armcompute.CommunityGalleryImageVersionsClientGetOptions) (resp azfake.Responder[armcompute.CommunityGalleryImageVersionsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method CommunityGalleryImageVersionsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(location string, publicGalleryName string, galleryImageName string, options *armcompute.CommunityGalleryImageVersionsClientListOptions) (resp azfake.PagerResponder[armcompute.CommunityGalleryImageVersionsClientListResponse])

}

// NewCommunityGalleryImageVersionsServerTransport creates a new instance of CommunityGalleryImageVersionsServerTransport with the provided implementation.
// The returned CommunityGalleryImageVersionsServerTransport instance is connected to an instance of armcompute.CommunityGalleryImageVersionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewCommunityGalleryImageVersionsServerTransport(srv *CommunityGalleryImageVersionsServer) *CommunityGalleryImageVersionsServerTransport {
	return &CommunityGalleryImageVersionsServerTransport{
		srv: srv,
		newListPager: newTracker[azfake.PagerResponder[armcompute.CommunityGalleryImageVersionsClientListResponse]](),
	}
}

// CommunityGalleryImageVersionsServerTransport connects instances of armcompute.CommunityGalleryImageVersionsClient to instances of CommunityGalleryImageVersionsServer.
// Don't use this type directly, use NewCommunityGalleryImageVersionsServerTransport instead.
type CommunityGalleryImageVersionsServerTransport struct {
	srv *CommunityGalleryImageVersionsServer
	newListPager *tracker[azfake.PagerResponder[armcompute.CommunityGalleryImageVersionsClientListResponse]]
}

// Do implements the policy.Transporter interface for CommunityGalleryImageVersionsServerTransport.
func (c *CommunityGalleryImageVersionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "CommunityGalleryImageVersionsClient.Get":
		resp, err = c.dispatchGet(req)
	case "CommunityGalleryImageVersionsClient.NewListPager":
		resp, err = c.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *CommunityGalleryImageVersionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if c.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/communityGalleries/(?P<publicGalleryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/images/(?P<galleryImageName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions/(?P<galleryImageVersionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	locationUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
	if err != nil {
		return nil, err
	}
	publicGalleryNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("publicGalleryName")])
	if err != nil {
		return nil, err
	}
	galleryImageNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("galleryImageName")])
	if err != nil {
		return nil, err
	}
	galleryImageVersionNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("galleryImageVersionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Get(req.Context(), locationUnescaped, publicGalleryNameUnescaped, galleryImageNameUnescaped, galleryImageVersionNameUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).CommunityGalleryImageVersion, req)
	if err != nil {		return nil, err
	}
	return resp, nil
}

func (c *CommunityGalleryImageVersionsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := c.newListPager.get(req)
	if newListPager == nil {
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/communityGalleries/(?P<publicGalleryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/images/(?P<galleryImageName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	locationUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
	if err != nil {
		return nil, err
	}
	publicGalleryNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("publicGalleryName")])
	if err != nil {
		return nil, err
	}
	galleryImageNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("galleryImageName")])
	if err != nil {
		return nil, err
	}
resp := c.srv.NewListPager(locationUnescaped, publicGalleryNameUnescaped, galleryImageNameUnescaped, nil)
		newListPager = &resp
		c.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armcompute.CommunityGalleryImageVersionsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		c.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		c.newListPager.remove(req)
	}
	return resp, nil
}

