//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmanagementgroups

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strconv"
)

// EntitiesClient contains the methods for the Entities group.
// Don't use this type directly, use NewEntitiesClient() instead.
type EntitiesClient struct {
	ep string
	pl runtime.Pipeline
}

// NewEntitiesClient creates a new instance of EntitiesClient with the specified values.
func NewEntitiesClient(credential azcore.TokenCredential, options *arm.ClientOptions) *EntitiesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &EntitiesClient{ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// List - List all entities (Management Groups, Subscriptions, etc.) for the authenticated user.
// If the operation fails it returns the *ErrorResponse error type.
func (client *EntitiesClient) List(options *EntitiesListOptions) *EntitiesListPager {
	return &EntitiesListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp EntitiesListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.EntityListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *EntitiesClient) listCreateRequest(ctx context.Context, options *EntitiesListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Management/getEntities"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-01")
	if options != nil && options.Skiptoken != nil {
		reqQP.Set("$skiptoken", *options.Skiptoken)
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Select != nil {
		reqQP.Set("$select", *options.Select)
	}
	if options != nil && options.Search != nil {
		reqQP.Set("$search", string(*options.Search))
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.View != nil {
		reqQP.Set("$view", string(*options.View))
	}
	if options != nil && options.GroupName != nil {
		reqQP.Set("groupName", *options.GroupName)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.CacheControl != nil {
		req.Raw().Header.Set("Cache-Control", *options.CacheControl)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *EntitiesClient) listHandleResponse(resp *http.Response) (EntitiesListResponse, error) {
	result := EntitiesListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.EntityListResult); err != nil {
		return EntitiesListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *EntitiesClient) listHandleError(resp *http.Response) error {
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
