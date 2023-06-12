//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armredis

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// LinkedServerClient contains the methods for the LinkedServer group.
// Don't use this type directly, use NewLinkedServerClient() instead.
type LinkedServerClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewLinkedServerClient creates a new instance of LinkedServerClient with the specified values.
//   - subscriptionID - Gets subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewLinkedServerClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*LinkedServerClient, error) {
	cl, err := arm.NewClient(moduleName+".LinkedServerClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &LinkedServerClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Adds a linked server to the Redis cache (requires Premium SKU).
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
//   - resourceGroupName - The name of the resource group.
//   - name - The name of the Redis cache.
//   - linkedServerName - The name of the linked server that is being added to the Redis cache.
//   - parameters - Parameters supplied to the Create Linked server operation.
//   - options - LinkedServerClientBeginCreateOptions contains the optional parameters for the LinkedServerClient.BeginCreate
//     method.
func (client *LinkedServerClient) BeginCreate(ctx context.Context, resourceGroupName string, name string, linkedServerName string, parameters LinkedServerCreateParameters, options *LinkedServerClientBeginCreateOptions) (*runtime.Poller[LinkedServerClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, name, linkedServerName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[LinkedServerClientCreateResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[LinkedServerClientCreateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Create - Adds a linked server to the Redis cache (requires Premium SKU).
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
func (client *LinkedServerClient) create(ctx context.Context, resourceGroupName string, name string, linkedServerName string, parameters LinkedServerCreateParameters, options *LinkedServerClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, name, linkedServerName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *LinkedServerClient) createCreateRequest(ctx context.Context, resourceGroupName string, name string, linkedServerName string, parameters LinkedServerCreateParameters, options *LinkedServerClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redis/{name}/linkedServers/{linkedServerName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if linkedServerName == "" {
		return nil, errors.New("parameter linkedServerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{linkedServerName}", url.PathEscape(linkedServerName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes the linked server from a redis cache (requires Premium SKU).
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
//   - resourceGroupName - The name of the resource group.
//   - name - The name of the redis cache.
//   - linkedServerName - The name of the linked server that is being added to the Redis cache.
//   - options - LinkedServerClientBeginDeleteOptions contains the optional parameters for the LinkedServerClient.BeginDelete
//     method.
func (client *LinkedServerClient) BeginDelete(ctx context.Context, resourceGroupName string, name string, linkedServerName string, options *LinkedServerClientBeginDeleteOptions) (*runtime.Poller[LinkedServerClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, name, linkedServerName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[LinkedServerClientDeleteResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[LinkedServerClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes the linked server from a redis cache (requires Premium SKU).
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
func (client *LinkedServerClient) deleteOperation(ctx context.Context, resourceGroupName string, name string, linkedServerName string, options *LinkedServerClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, name, linkedServerName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *LinkedServerClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, name string, linkedServerName string, options *LinkedServerClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redis/{name}/linkedServers/{linkedServerName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if linkedServerName == "" {
		return nil, errors.New("parameter linkedServerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{linkedServerName}", url.PathEscape(linkedServerName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the detailed information about a linked server of a redis cache (requires Premium SKU).
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
//   - resourceGroupName - The name of the resource group.
//   - name - The name of the redis cache.
//   - linkedServerName - The name of the linked server.
//   - options - LinkedServerClientGetOptions contains the optional parameters for the LinkedServerClient.Get method.
func (client *LinkedServerClient) Get(ctx context.Context, resourceGroupName string, name string, linkedServerName string, options *LinkedServerClientGetOptions) (LinkedServerClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, name, linkedServerName, options)
	if err != nil {
		return LinkedServerClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LinkedServerClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LinkedServerClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *LinkedServerClient) getCreateRequest(ctx context.Context, resourceGroupName string, name string, linkedServerName string, options *LinkedServerClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redis/{name}/linkedServers/{linkedServerName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if linkedServerName == "" {
		return nil, errors.New("parameter linkedServerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{linkedServerName}", url.PathEscape(linkedServerName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *LinkedServerClient) getHandleResponse(resp *http.Response) (LinkedServerClientGetResponse, error) {
	result := LinkedServerClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LinkedServerWithProperties); err != nil {
		return LinkedServerClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets the list of linked servers associated with this redis cache (requires Premium SKU).
//
// Generated from API version 2023-04-01
//   - resourceGroupName - The name of the resource group.
//   - name - The name of the redis cache.
//   - options - LinkedServerClientListOptions contains the optional parameters for the LinkedServerClient.NewListPager method.
func (client *LinkedServerClient) NewListPager(resourceGroupName string, name string, options *LinkedServerClientListOptions) *runtime.Pager[LinkedServerClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[LinkedServerClientListResponse]{
		More: func(page LinkedServerClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *LinkedServerClientListResponse) (LinkedServerClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, name, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return LinkedServerClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return LinkedServerClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return LinkedServerClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *LinkedServerClient) listCreateRequest(ctx context.Context, resourceGroupName string, name string, options *LinkedServerClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redis/{name}/linkedServers"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *LinkedServerClient) listHandleResponse(resp *http.Response) (LinkedServerClientListResponse, error) {
	result := LinkedServerClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LinkedServerWithPropertiesList); err != nil {
		return LinkedServerClientListResponse{}, err
	}
	return result, nil
}
