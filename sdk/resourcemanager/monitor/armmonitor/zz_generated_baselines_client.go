//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmonitor

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
)

// BaselinesClient contains the methods for the Baselines group.
// Don't use this type directly, use NewBaselinesClient() instead.
type BaselinesClient struct {
	ep string
	pl runtime.Pipeline
}

// NewBaselinesClient creates a new instance of BaselinesClient with the specified values.
func NewBaselinesClient(credential azcore.TokenCredential, options *arm.ClientOptions) *BaselinesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &BaselinesClient{ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// List - Lists the metric baseline values for a resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *BaselinesClient) List(ctx context.Context, resourceURI string, options *BaselinesListOptions) (BaselinesListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceURI, options)
	if err != nil {
		return BaselinesListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BaselinesListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BaselinesListResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *BaselinesClient) listCreateRequest(ctx context.Context, resourceURI string, options *BaselinesListOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.Insights/metricBaselines"
	if resourceURI == "" {
		return nil, errors.New("parameter resourceURI cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Metricnames != nil {
		reqQP.Set("metricnames", *options.Metricnames)
	}
	if options != nil && options.Metricnamespace != nil {
		reqQP.Set("metricnamespace", *options.Metricnamespace)
	}
	if options != nil && options.Timespan != nil {
		reqQP.Set("timespan", *options.Timespan)
	}
	if options != nil && options.Interval != nil {
		reqQP.Set("interval", *options.Interval)
	}
	if options != nil && options.Aggregation != nil {
		reqQP.Set("aggregation", *options.Aggregation)
	}
	if options != nil && options.Sensitivities != nil {
		reqQP.Set("sensitivities", *options.Sensitivities)
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.ResultType != nil {
		reqQP.Set("resultType", string(*options.ResultType))
	}
	reqQP.Set("api-version", "2019-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *BaselinesClient) listHandleResponse(resp *http.Response) (BaselinesListResponse, error) {
	result := BaselinesListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.MetricBaselinesResponse); err != nil {
		return BaselinesListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *BaselinesClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
