//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapimanagement

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
	"net/url"
	"strings"
)

// APIManagementSKUsClient contains the methods for the APIManagementSKUs group.
// Don't use this type directly, use NewAPIManagementSKUsClient() instead.
type APIManagementSKUsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewAPIManagementSKUsClient creates a new instance of APIManagementSKUsClient with the specified values.
func NewAPIManagementSKUsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *APIManagementSKUsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &APIManagementSKUsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// List - Gets the list of Microsoft.ApiManagement SKUs available for your Subscription.
// If the operation fails it returns the *ErrorResponse error type.
func (client *APIManagementSKUsClient) List(options *APIManagementSKUsListOptions) *APIManagementSKUsListPager {
	return &APIManagementSKUsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp APIManagementSKUsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.APIManagementSKUsResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *APIManagementSKUsClient) listCreateRequest(ctx context.Context, options *APIManagementSKUsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ApiManagement/skus"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *APIManagementSKUsClient) listHandleResponse(resp *http.Response) (APIManagementSKUsListResponse, error) {
	result := APIManagementSKUsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.APIManagementSKUsResult); err != nil {
		return APIManagementSKUsListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *APIManagementSKUsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
