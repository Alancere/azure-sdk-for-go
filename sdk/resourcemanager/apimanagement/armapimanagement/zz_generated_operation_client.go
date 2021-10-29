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
	"strconv"
	"strings"
)

// OperationClient contains the methods for the Operation group.
// Don't use this type directly, use NewOperationClient() instead.
type OperationClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewOperationClient creates a new instance of OperationClient with the specified values.
func NewOperationClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *OperationClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &OperationClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// ListByTags - Lists a collection of operations associated with tags.
// If the operation fails it returns the *ErrorResponse error type.
func (client *OperationClient) ListByTags(resourceGroupName string, serviceName string, apiID string, options *OperationListByTagsOptions) *OperationListByTagsPager {
	return &OperationListByTagsPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByTagsCreateRequest(ctx, resourceGroupName, serviceName, apiID, options)
		},
		advancer: func(ctx context.Context, resp OperationListByTagsResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.TagResourceCollection.NextLink)
		},
	}
}

// listByTagsCreateRequest creates the ListByTags request.
func (client *OperationClient) listByTagsCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, apiID string, options *OperationListByTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/operationsByTags"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if apiID == "" {
		return nil, errors.New("parameter apiID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiId}", url.PathEscape(apiID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	if options != nil && options.IncludeNotTaggedOperations != nil {
		reqQP.Set("includeNotTaggedOperations", strconv.FormatBool(*options.IncludeNotTaggedOperations))
	}
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByTagsHandleResponse handles the ListByTags response.
func (client *OperationClient) listByTagsHandleResponse(resp *http.Response) (OperationListByTagsResponse, error) {
	result := OperationListByTagsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.TagResourceCollection); err != nil {
		return OperationListByTagsResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByTagsHandleError handles the ListByTags error response.
func (client *OperationClient) listByTagsHandleError(resp *http.Response) error {
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
