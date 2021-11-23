//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armchangeanalysis

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
	"time"
)

// ChangesClient contains the methods for the Changes group.
// Don't use this type directly, use NewChangesClient() instead.
type ChangesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewChangesClient creates a new instance of ChangesClient with the specified values.
func NewChangesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ChangesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &ChangesClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// ListChangesByResourceGroup - List the changes of a resource group within the specified time range. Customer data will always be masked.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ChangesClient) ListChangesByResourceGroup(resourceGroupName string, startTime time.Time, endTime time.Time, options *ChangesListChangesByResourceGroupOptions) *ChangesListChangesByResourceGroupPager {
	return &ChangesListChangesByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listChangesByResourceGroupCreateRequest(ctx, resourceGroupName, startTime, endTime, options)
		},
		advancer: func(ctx context.Context, resp ChangesListChangesByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ChangeList.NextLink)
		},
	}
}

// listChangesByResourceGroupCreateRequest creates the ListChangesByResourceGroup request.
func (client *ChangesClient) listChangesByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, startTime time.Time, endTime time.Time, options *ChangesListChangesByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ChangeAnalysis/changes"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-01")
	reqQP.Set("$startTime", startTime.Format(time.RFC3339Nano))
	reqQP.Set("$endTime", endTime.Format(time.RFC3339Nano))
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listChangesByResourceGroupHandleResponse handles the ListChangesByResourceGroup response.
func (client *ChangesClient) listChangesByResourceGroupHandleResponse(resp *http.Response) (ChangesListChangesByResourceGroupResponse, error) {
	result := ChangesListChangesByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ChangeList); err != nil {
		return ChangesListChangesByResourceGroupResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listChangesByResourceGroupHandleError handles the ListChangesByResourceGroup error response.
func (client *ChangesClient) listChangesByResourceGroupHandleError(resp *http.Response) error {
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

// ListChangesBySubscription - List the changes of a subscription within the specified time range. Customer data will always be masked.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ChangesClient) ListChangesBySubscription(startTime time.Time, endTime time.Time, options *ChangesListChangesBySubscriptionOptions) *ChangesListChangesBySubscriptionPager {
	return &ChangesListChangesBySubscriptionPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listChangesBySubscriptionCreateRequest(ctx, startTime, endTime, options)
		},
		advancer: func(ctx context.Context, resp ChangesListChangesBySubscriptionResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ChangeList.NextLink)
		},
	}
}

// listChangesBySubscriptionCreateRequest creates the ListChangesBySubscription request.
func (client *ChangesClient) listChangesBySubscriptionCreateRequest(ctx context.Context, startTime time.Time, endTime time.Time, options *ChangesListChangesBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ChangeAnalysis/changes"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-01")
	reqQP.Set("$startTime", startTime.Format(time.RFC3339Nano))
	reqQP.Set("$endTime", endTime.Format(time.RFC3339Nano))
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listChangesBySubscriptionHandleResponse handles the ListChangesBySubscription response.
func (client *ChangesClient) listChangesBySubscriptionHandleResponse(resp *http.Response) (ChangesListChangesBySubscriptionResponse, error) {
	result := ChangesListChangesBySubscriptionResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ChangeList); err != nil {
		return ChangesListChangesBySubscriptionResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listChangesBySubscriptionHandleError handles the ListChangesBySubscription error response.
func (client *ChangesClient) listChangesBySubscriptionHandleError(resp *http.Response) error {
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
