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

// VMInsightsClient contains the methods for the VMInsights group.
// Don't use this type directly, use NewVMInsightsClient() instead.
type VMInsightsClient struct {
	ep string
	pl runtime.Pipeline
}

// NewVMInsightsClient creates a new instance of VMInsightsClient with the specified values.
func NewVMInsightsClient(credential azcore.TokenCredential, options *arm.ClientOptions) *VMInsightsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &VMInsightsClient{ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// GetOnboardingStatus - Retrieves the VM Insights onboarding status for the specified resource or resource scope.
// If the operation fails it returns the *ResponseWithError error type.
func (client *VMInsightsClient) GetOnboardingStatus(ctx context.Context, resourceURI string, options *VMInsightsGetOnboardingStatusOptions) (VMInsightsGetOnboardingStatusResponse, error) {
	req, err := client.getOnboardingStatusCreateRequest(ctx, resourceURI, options)
	if err != nil {
		return VMInsightsGetOnboardingStatusResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VMInsightsGetOnboardingStatusResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VMInsightsGetOnboardingStatusResponse{}, client.getOnboardingStatusHandleError(resp)
	}
	return client.getOnboardingStatusHandleResponse(resp)
}

// getOnboardingStatusCreateRequest creates the GetOnboardingStatus request.
func (client *VMInsightsClient) getOnboardingStatusCreateRequest(ctx context.Context, resourceURI string, options *VMInsightsGetOnboardingStatusOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.Insights/vmInsightsOnboardingStatuses/default"
	if resourceURI == "" {
		return nil, errors.New("parameter resourceURI cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-11-27-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getOnboardingStatusHandleResponse handles the GetOnboardingStatus response.
func (client *VMInsightsClient) getOnboardingStatusHandleResponse(resp *http.Response) (VMInsightsGetOnboardingStatusResponse, error) {
	result := VMInsightsGetOnboardingStatusResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VMInsightsOnboardingStatus); err != nil {
		return VMInsightsGetOnboardingStatusResponse{}, err
	}
	return result, nil
}

// getOnboardingStatusHandleError handles the GetOnboardingStatus error response.
func (client *VMInsightsClient) getOnboardingStatusHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ResponseWithError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
