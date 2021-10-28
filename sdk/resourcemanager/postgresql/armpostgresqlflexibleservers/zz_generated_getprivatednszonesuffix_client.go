//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpostgresqlflexibleservers

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// GetPrivateDNSZoneSuffixClient contains the methods for the GetPrivateDNSZoneSuffix group.
// Don't use this type directly, use NewGetPrivateDNSZoneSuffixClient() instead.
type GetPrivateDNSZoneSuffixClient struct {
	ep string
	pl runtime.Pipeline
}

// NewGetPrivateDNSZoneSuffixClient creates a new instance of GetPrivateDNSZoneSuffixClient with the specified values.
func NewGetPrivateDNSZoneSuffixClient(credential azcore.TokenCredential, options *arm.ClientOptions) *GetPrivateDNSZoneSuffixClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &GetPrivateDNSZoneSuffixClient{ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// Execute - Get private DNS zone suffix in the cloud
// If the operation fails it returns the *CloudError error type.
func (client *GetPrivateDNSZoneSuffixClient) Execute(ctx context.Context, options *GetPrivateDNSZoneSuffixExecuteOptions) (GetPrivateDNSZoneSuffixExecuteResponse, error) {
	req, err := client.executeCreateRequest(ctx, options)
	if err != nil {
		return GetPrivateDNSZoneSuffixExecuteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return GetPrivateDNSZoneSuffixExecuteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return GetPrivateDNSZoneSuffixExecuteResponse{}, client.executeHandleError(resp)
	}
	return client.executeHandleResponse(resp)
}

// executeCreateRequest creates the Execute request.
func (client *GetPrivateDNSZoneSuffixClient) executeCreateRequest(ctx context.Context, options *GetPrivateDNSZoneSuffixExecuteOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.DBforPostgreSQL/getPrivateDnsZoneSuffix"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// executeHandleResponse handles the Execute response.
func (client *GetPrivateDNSZoneSuffixClient) executeHandleResponse(resp *http.Response) (GetPrivateDNSZoneSuffixExecuteResponse, error) {
	result := GetPrivateDNSZoneSuffixExecuteResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return GetPrivateDNSZoneSuffixExecuteResponse{}, err
	}
	return result, nil
}

// executeHandleError handles the Execute error response.
func (client *GetPrivateDNSZoneSuffixClient) executeHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
