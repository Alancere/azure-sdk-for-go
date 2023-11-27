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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloads/armworkloads"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
)

// Server is a fake server for instances of the armworkloads.Client type.
type Server struct {
	// SAPAvailabilityZoneDetails is the fake for method Client.SAPAvailabilityZoneDetails
	// HTTP status codes to indicate success: http.StatusOK
	SAPAvailabilityZoneDetails func(ctx context.Context, location string, options *armworkloads.ClientSAPAvailabilityZoneDetailsOptions) (resp azfake.Responder[armworkloads.ClientSAPAvailabilityZoneDetailsResponse], errResp azfake.ErrorResponder)

	// SAPDiskConfigurations is the fake for method Client.SAPDiskConfigurations
	// HTTP status codes to indicate success: http.StatusOK
	SAPDiskConfigurations func(ctx context.Context, location string, options *armworkloads.ClientSAPDiskConfigurationsOptions) (resp azfake.Responder[armworkloads.ClientSAPDiskConfigurationsResponse], errResp azfake.ErrorResponder)

	// SAPSizingRecommendations is the fake for method Client.SAPSizingRecommendations
	// HTTP status codes to indicate success: http.StatusOK
	SAPSizingRecommendations func(ctx context.Context, location string, options *armworkloads.ClientSAPSizingRecommendationsOptions) (resp azfake.Responder[armworkloads.ClientSAPSizingRecommendationsResponse], errResp azfake.ErrorResponder)

	// SAPSupportedSKU is the fake for method Client.SAPSupportedSKU
	// HTTP status codes to indicate success: http.StatusOK
	SAPSupportedSKU func(ctx context.Context, location string, options *armworkloads.ClientSAPSupportedSKUOptions) (resp azfake.Responder[armworkloads.ClientSAPSupportedSKUResponse], errResp azfake.ErrorResponder)
}

// NewServerTransport creates a new instance of ServerTransport with the provided implementation.
// The returned ServerTransport instance is connected to an instance of armworkloads.Client via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerTransport(srv *Server) *ServerTransport {
	return &ServerTransport{srv: srv}
}

// ServerTransport connects instances of armworkloads.Client to instances of Server.
// Don't use this type directly, use NewServerTransport instead.
type ServerTransport struct {
	srv *Server
}

// Do implements the policy.Transporter interface for ServerTransport.
func (s *ServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "Client.SAPAvailabilityZoneDetails":
		resp, err = s.dispatchSAPAvailabilityZoneDetails(req)
	case "Client.SAPDiskConfigurations":
		resp, err = s.dispatchSAPDiskConfigurations(req)
	case "Client.SAPSizingRecommendations":
		resp, err = s.dispatchSAPSizingRecommendations(req)
	case "Client.SAPSupportedSKU":
		resp, err = s.dispatchSAPSupportedSKU(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *ServerTransport) dispatchSAPAvailabilityZoneDetails(req *http.Request) (*http.Response, error) {
	if s.srv.SAPAvailabilityZoneDetails == nil {
		return nil, &nonRetriableError{errors.New("fake for method SAPAvailabilityZoneDetails not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Workloads/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sapVirtualInstanceMetadata/default/getAvailabilityZoneDetails`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armworkloads.SAPAvailabilityZoneDetailsRequest](req)
	if err != nil {
		return nil, err
	}
	locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
	if err != nil {
		return nil, err
	}
	var options *armworkloads.ClientSAPAvailabilityZoneDetailsOptions
	if !reflect.ValueOf(body).IsZero() {
		options = &armworkloads.ClientSAPAvailabilityZoneDetailsOptions{
			SAPAvailabilityZoneDetails: &body,
		}
	}
	respr, errRespr := s.srv.SAPAvailabilityZoneDetails(req.Context(), locationParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SAPAvailabilityZoneDetailsResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ServerTransport) dispatchSAPDiskConfigurations(req *http.Request) (*http.Response, error) {
	if s.srv.SAPDiskConfigurations == nil {
		return nil, &nonRetriableError{errors.New("fake for method SAPDiskConfigurations not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Workloads/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sapVirtualInstanceMetadata/default/getDiskConfigurations`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armworkloads.SAPDiskConfigurationsRequest](req)
	if err != nil {
		return nil, err
	}
	locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
	if err != nil {
		return nil, err
	}
	var options *armworkloads.ClientSAPDiskConfigurationsOptions
	if !reflect.ValueOf(body).IsZero() {
		options = &armworkloads.ClientSAPDiskConfigurationsOptions{
			SAPDiskConfigurations: &body,
		}
	}
	respr, errRespr := s.srv.SAPDiskConfigurations(req.Context(), locationParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SAPDiskConfigurationsResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ServerTransport) dispatchSAPSizingRecommendations(req *http.Request) (*http.Response, error) {
	if s.srv.SAPSizingRecommendations == nil {
		return nil, &nonRetriableError{errors.New("fake for method SAPSizingRecommendations not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Workloads/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sapVirtualInstanceMetadata/default/getSizingRecommendations`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armworkloads.SAPSizingRecommendationRequest](req)
	if err != nil {
		return nil, err
	}
	locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
	if err != nil {
		return nil, err
	}
	var options *armworkloads.ClientSAPSizingRecommendationsOptions
	if !reflect.ValueOf(body).IsZero() {
		options = &armworkloads.ClientSAPSizingRecommendationsOptions{
			SAPSizingRecommendation: &body,
		}
	}
	respr, errRespr := s.srv.SAPSizingRecommendations(req.Context(), locationParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SAPSizingRecommendationResultClassification, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ServerTransport) dispatchSAPSupportedSKU(req *http.Request) (*http.Response, error) {
	if s.srv.SAPSupportedSKU == nil {
		return nil, &nonRetriableError{errors.New("fake for method SAPSupportedSKU not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Workloads/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sapVirtualInstanceMetadata/default/getSapSupportedSku`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armworkloads.SAPSupportedSKUsRequest](req)
	if err != nil {
		return nil, err
	}
	locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
	if err != nil {
		return nil, err
	}
	var options *armworkloads.ClientSAPSupportedSKUOptions
	if !reflect.ValueOf(body).IsZero() {
		options = &armworkloads.ClientSAPSupportedSKUOptions{
			SAPSupportedSKU: &body,
		}
	}
	respr, errRespr := s.srv.SAPSupportedSKU(req.Context(), locationParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SAPSupportedResourceSKUsResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
