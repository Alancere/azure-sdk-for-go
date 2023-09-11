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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5"
	"net/http"
	"net/url"
	"regexp"
)

// LogAnalyticsServer is a fake server for instances of the armcompute.LogAnalyticsClient type.
type LogAnalyticsServer struct{
	// BeginExportRequestRateByInterval is the fake for method LogAnalyticsClient.BeginExportRequestRateByInterval
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginExportRequestRateByInterval func(ctx context.Context, location string, parameters armcompute.RequestRateByIntervalInput, options *armcompute.LogAnalyticsClientBeginExportRequestRateByIntervalOptions) (resp azfake.PollerResponder[armcompute.LogAnalyticsClientExportRequestRateByIntervalResponse], errResp azfake.ErrorResponder)

	// BeginExportThrottledRequests is the fake for method LogAnalyticsClient.BeginExportThrottledRequests
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginExportThrottledRequests func(ctx context.Context, location string, parameters armcompute.ThrottledRequestsInput, options *armcompute.LogAnalyticsClientBeginExportThrottledRequestsOptions) (resp azfake.PollerResponder[armcompute.LogAnalyticsClientExportThrottledRequestsResponse], errResp azfake.ErrorResponder)

}

// NewLogAnalyticsServerTransport creates a new instance of LogAnalyticsServerTransport with the provided implementation.
// The returned LogAnalyticsServerTransport instance is connected to an instance of armcompute.LogAnalyticsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewLogAnalyticsServerTransport(srv *LogAnalyticsServer) *LogAnalyticsServerTransport {
	return &LogAnalyticsServerTransport{
		srv: srv,
		beginExportRequestRateByInterval: newTracker[azfake.PollerResponder[armcompute.LogAnalyticsClientExportRequestRateByIntervalResponse]](),
		beginExportThrottledRequests: newTracker[azfake.PollerResponder[armcompute.LogAnalyticsClientExportThrottledRequestsResponse]](),
	}
}

// LogAnalyticsServerTransport connects instances of armcompute.LogAnalyticsClient to instances of LogAnalyticsServer.
// Don't use this type directly, use NewLogAnalyticsServerTransport instead.
type LogAnalyticsServerTransport struct {
	srv *LogAnalyticsServer
	beginExportRequestRateByInterval *tracker[azfake.PollerResponder[armcompute.LogAnalyticsClientExportRequestRateByIntervalResponse]]
	beginExportThrottledRequests *tracker[azfake.PollerResponder[armcompute.LogAnalyticsClientExportThrottledRequestsResponse]]
}

// Do implements the policy.Transporter interface for LogAnalyticsServerTransport.
func (l *LogAnalyticsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "LogAnalyticsClient.BeginExportRequestRateByInterval":
		resp, err = l.dispatchBeginExportRequestRateByInterval(req)
	case "LogAnalyticsClient.BeginExportThrottledRequests":
		resp, err = l.dispatchBeginExportThrottledRequests(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (l *LogAnalyticsServerTransport) dispatchBeginExportRequestRateByInterval(req *http.Request) (*http.Response, error) {
	if l.srv.BeginExportRequestRateByInterval == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginExportRequestRateByInterval not implemented")}
	}
	beginExportRequestRateByInterval := l.beginExportRequestRateByInterval.get(req)
	if beginExportRequestRateByInterval == nil {
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/logAnalytics/apiAccess/getRequestRateByInterval`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armcompute.RequestRateByIntervalInput](req)
	if err != nil {
		return nil, err
	}
	locationUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := l.srv.BeginExportRequestRateByInterval(req.Context(), locationUnescaped, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
		beginExportRequestRateByInterval = &respr
		l.beginExportRequestRateByInterval.add(req, beginExportRequestRateByInterval)
	}

	resp, err := server.PollerResponderNext(beginExportRequestRateByInterval, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		l.beginExportRequestRateByInterval.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginExportRequestRateByInterval) {
		l.beginExportRequestRateByInterval.remove(req)
	}

	return resp, nil
}

func (l *LogAnalyticsServerTransport) dispatchBeginExportThrottledRequests(req *http.Request) (*http.Response, error) {
	if l.srv.BeginExportThrottledRequests == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginExportThrottledRequests not implemented")}
	}
	beginExportThrottledRequests := l.beginExportThrottledRequests.get(req)
	if beginExportThrottledRequests == nil {
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/logAnalytics/apiAccess/getThrottledRequests`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armcompute.ThrottledRequestsInput](req)
	if err != nil {
		return nil, err
	}
	locationUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := l.srv.BeginExportThrottledRequests(req.Context(), locationUnescaped, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
		beginExportThrottledRequests = &respr
		l.beginExportThrottledRequests.add(req, beginExportThrottledRequests)
	}

	resp, err := server.PollerResponderNext(beginExportThrottledRequests, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		l.beginExportThrottledRequests.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginExportThrottledRequests) {
		l.beginExportThrottledRequests.remove(req)
	}

	return resp, nil
}

