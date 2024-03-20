//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armiotfirmwaredefense

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

// SummariesClient contains the methods for the Summaries group.
// Don't use this type directly, use NewSummariesClient() instead.
type SummariesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSummariesClient creates a new instance of SummariesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSummariesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SummariesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SummariesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Get an analysis result summary of a firmware by name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-10
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the firmware analysis workspace.
//   - firmwareID - The id of the firmware.
//   - summaryName - The Firmware analysis summary name describing the type of summary.
//   - options - SummariesClientGetOptions contains the optional parameters for the SummariesClient.Get method.
func (client *SummariesClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, firmwareID string, summaryName SummaryName, options *SummariesClientGetOptions) (SummariesClientGetResponse, error) {
	var err error
	const operationName = "SummariesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, firmwareID, summaryName, options)
	if err != nil {
		return SummariesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SummariesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SummariesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *SummariesClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, firmwareID string, summaryName SummaryName, options *SummariesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.IoTFirmwareDefense/workspaces/{workspaceName}/firmwares/{firmwareId}/summaries/{summaryName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if firmwareID == "" {
		return nil, errors.New("parameter firmwareID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{firmwareId}", url.PathEscape(firmwareID))
	if summaryName == "" {
		return nil, errors.New("parameter summaryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{summaryName}", url.PathEscape(string(summaryName)))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-10")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SummariesClient) getHandleResponse(resp *http.Response) (SummariesClientGetResponse, error) {
	result := SummariesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SummaryResource); err != nil {
		return SummariesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByFirmwarePager - Lists analysis result summary names of a firmware. To fetch the full summary data, get that summary
// by name.
//
// Generated from API version 2024-01-10
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the firmware analysis workspace.
//   - firmwareID - The id of the firmware.
//   - options - SummariesClientListByFirmwareOptions contains the optional parameters for the SummariesClient.NewListByFirmwarePager
//     method.
func (client *SummariesClient) NewListByFirmwarePager(resourceGroupName string, workspaceName string, firmwareID string, options *SummariesClientListByFirmwareOptions) *runtime.Pager[SummariesClientListByFirmwareResponse] {
	return runtime.NewPager(runtime.PagingHandler[SummariesClientListByFirmwareResponse]{
		More: func(page SummariesClientListByFirmwareResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SummariesClientListByFirmwareResponse) (SummariesClientListByFirmwareResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SummariesClient.NewListByFirmwarePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByFirmwareCreateRequest(ctx, resourceGroupName, workspaceName, firmwareID, options)
			}, nil)
			if err != nil {
				return SummariesClientListByFirmwareResponse{}, err
			}
			return client.listByFirmwareHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByFirmwareCreateRequest creates the ListByFirmware request.
func (client *SummariesClient) listByFirmwareCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, firmwareID string, options *SummariesClientListByFirmwareOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.IoTFirmwareDefense/workspaces/{workspaceName}/firmwares/{firmwareId}/summaries"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if firmwareID == "" {
		return nil, errors.New("parameter firmwareID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{firmwareId}", url.PathEscape(firmwareID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-10")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByFirmwareHandleResponse handles the ListByFirmware response.
func (client *SummariesClient) listByFirmwareHandleResponse(resp *http.Response) (SummariesClientListByFirmwareResponse, error) {
	result := SummariesClientListByFirmwareResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SummaryListResult); err != nil {
		return SummariesClientListByFirmwareResponse{}, err
	}
	return result, nil
}
