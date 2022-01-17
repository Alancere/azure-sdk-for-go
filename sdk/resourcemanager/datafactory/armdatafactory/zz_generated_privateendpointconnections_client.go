//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatafactory

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// PrivateEndPointConnectionsClient contains the methods for the PrivateEndPointConnections group.
// Don't use this type directly, use NewPrivateEndPointConnectionsClient() instead.
type PrivateEndPointConnectionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewPrivateEndPointConnectionsClient creates a new instance of PrivateEndPointConnectionsClient with the specified values.
// subscriptionID - The subscription identifier.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewPrivateEndPointConnectionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *PrivateEndPointConnectionsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &PrivateEndPointConnectionsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// ListByFactory - Lists Private endpoint connections
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name.
// factoryName - The factory name.
// options - PrivateEndPointConnectionsClientListByFactoryOptions contains the optional parameters for the PrivateEndPointConnectionsClient.ListByFactory
// method.
func (client *PrivateEndPointConnectionsClient) ListByFactory(resourceGroupName string, factoryName string, options *PrivateEndPointConnectionsClientListByFactoryOptions) *PrivateEndPointConnectionsClientListByFactoryPager {
	return &PrivateEndPointConnectionsClientListByFactoryPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByFactoryCreateRequest(ctx, resourceGroupName, factoryName, options)
		},
		advancer: func(ctx context.Context, resp PrivateEndPointConnectionsClientListByFactoryResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.PrivateEndpointConnectionListResponse.NextLink)
		},
	}
}

// listByFactoryCreateRequest creates the ListByFactory request.
func (client *PrivateEndPointConnectionsClient) listByFactoryCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, options *PrivateEndPointConnectionsClientListByFactoryOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/privateEndPointConnections"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByFactoryHandleResponse handles the ListByFactory response.
func (client *PrivateEndPointConnectionsClient) listByFactoryHandleResponse(resp *http.Response) (PrivateEndPointConnectionsClientListByFactoryResponse, error) {
	result := PrivateEndPointConnectionsClientListByFactoryResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateEndpointConnectionListResponse); err != nil {
		return PrivateEndPointConnectionsClientListByFactoryResponse{}, err
	}
	return result, nil
}
