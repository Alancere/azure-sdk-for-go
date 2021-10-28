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

// APIManagementServiceSKUsClient contains the methods for the APIManagementServiceSKUs group.
// Don't use this type directly, use NewAPIManagementServiceSKUsClient() instead.
type APIManagementServiceSKUsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewAPIManagementServiceSKUsClient creates a new instance of APIManagementServiceSKUsClient with the specified values.
func NewAPIManagementServiceSKUsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *APIManagementServiceSKUsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &APIManagementServiceSKUsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// ListAvailableServiceSKUs - Gets all available SKU for a given API Management service
// If the operation fails it returns the *ErrorResponse error type.
func (client *APIManagementServiceSKUsClient) ListAvailableServiceSKUs(resourceGroupName string, serviceName string, options *APIManagementServiceSKUsListAvailableServiceSKUsOptions) *APIManagementServiceSKUsListAvailableServiceSKUsPager {
	return &APIManagementServiceSKUsListAvailableServiceSKUsPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listAvailableServiceSKUsCreateRequest(ctx, resourceGroupName, serviceName, options)
		},
		advancer: func(ctx context.Context, resp APIManagementServiceSKUsListAvailableServiceSKUsResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ResourceSKUResults.NextLink)
		},
	}
}

// listAvailableServiceSKUsCreateRequest creates the ListAvailableServiceSKUs request.
func (client *APIManagementServiceSKUsClient) listAvailableServiceSKUsCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *APIManagementServiceSKUsListAvailableServiceSKUsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/skus"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
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

// listAvailableServiceSKUsHandleResponse handles the ListAvailableServiceSKUs response.
func (client *APIManagementServiceSKUsClient) listAvailableServiceSKUsHandleResponse(resp *http.Response) (APIManagementServiceSKUsListAvailableServiceSKUsResponse, error) {
	result := APIManagementServiceSKUsListAvailableServiceSKUsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourceSKUResults); err != nil {
		return APIManagementServiceSKUsListAvailableServiceSKUsResponse{}, err
	}
	return result, nil
}

// listAvailableServiceSKUsHandleError handles the ListAvailableServiceSKUs error response.
func (client *APIManagementServiceSKUsClient) listAvailableServiceSKUsHandleError(resp *http.Response) error {
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
