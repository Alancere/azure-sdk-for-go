//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdataboxedge

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

// ContainersClient contains the methods for the Containers group.
// Don't use this type directly, use NewContainersClient() instead.
type ContainersClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewContainersClient creates a new instance of ContainersClient with the specified values.
func NewContainersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ContainersClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &ContainersClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// BeginCreateOrUpdate - Creates a new container or updates an existing container on the device.
// If the operation fails it returns the *CloudError error type.
func (client *ContainersClient) BeginCreateOrUpdate(ctx context.Context, deviceName string, storageAccountName string, containerName string, resourceGroupName string, containerParam Container, options *ContainersBeginCreateOrUpdateOptions) (ContainersCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, deviceName, storageAccountName, containerName, resourceGroupName, containerParam, options)
	if err != nil {
		return ContainersCreateOrUpdatePollerResponse{}, err
	}
	result := ContainersCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ContainersClient.CreateOrUpdate", "", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return ContainersCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &ContainersCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates a new container or updates an existing container on the device.
// If the operation fails it returns the *CloudError error type.
func (client *ContainersClient) createOrUpdate(ctx context.Context, deviceName string, storageAccountName string, containerName string, resourceGroupName string, containerParam Container, options *ContainersBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, deviceName, storageAccountName, containerName, resourceGroupName, containerParam, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ContainersClient) createOrUpdateCreateRequest(ctx context.Context, deviceName string, storageAccountName string, containerName string, resourceGroupName string, containerParam Container, options *ContainersBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/storageAccounts/{storageAccountName}/containers/{containerName}"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if storageAccountName == "" {
		return nil, errors.New("parameter storageAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageAccountName}", url.PathEscape(storageAccountName))
	if containerName == "" {
		return nil, errors.New("parameter containerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerName}", url.PathEscape(containerName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, containerParam)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *ContainersClient) createOrUpdateHandleError(resp *http.Response) error {
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

// BeginDelete - Deletes the container on the Data Box Edge/Data Box Gateway device.
// If the operation fails it returns the *CloudError error type.
func (client *ContainersClient) BeginDelete(ctx context.Context, deviceName string, storageAccountName string, containerName string, resourceGroupName string, options *ContainersBeginDeleteOptions) (ContainersDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, deviceName, storageAccountName, containerName, resourceGroupName, options)
	if err != nil {
		return ContainersDeletePollerResponse{}, err
	}
	result := ContainersDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ContainersClient.Delete", "", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return ContainersDeletePollerResponse{}, err
	}
	result.Poller = &ContainersDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the container on the Data Box Edge/Data Box Gateway device.
// If the operation fails it returns the *CloudError error type.
func (client *ContainersClient) deleteOperation(ctx context.Context, deviceName string, storageAccountName string, containerName string, resourceGroupName string, options *ContainersBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, deviceName, storageAccountName, containerName, resourceGroupName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ContainersClient) deleteCreateRequest(ctx context.Context, deviceName string, storageAccountName string, containerName string, resourceGroupName string, options *ContainersBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/storageAccounts/{storageAccountName}/containers/{containerName}"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if storageAccountName == "" {
		return nil, errors.New("parameter storageAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageAccountName}", url.PathEscape(storageAccountName))
	if containerName == "" {
		return nil, errors.New("parameter containerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerName}", url.PathEscape(containerName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *ContainersClient) deleteHandleError(resp *http.Response) error {
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

// Get - Gets a container by name.
// If the operation fails it returns the *CloudError error type.
func (client *ContainersClient) Get(ctx context.Context, deviceName string, storageAccountName string, containerName string, resourceGroupName string, options *ContainersGetOptions) (ContainersGetResponse, error) {
	req, err := client.getCreateRequest(ctx, deviceName, storageAccountName, containerName, resourceGroupName, options)
	if err != nil {
		return ContainersGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ContainersGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ContainersGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ContainersClient) getCreateRequest(ctx context.Context, deviceName string, storageAccountName string, containerName string, resourceGroupName string, options *ContainersGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/storageAccounts/{storageAccountName}/containers/{containerName}"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if storageAccountName == "" {
		return nil, errors.New("parameter storageAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageAccountName}", url.PathEscape(storageAccountName))
	if containerName == "" {
		return nil, errors.New("parameter containerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerName}", url.PathEscape(containerName))
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
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ContainersClient) getHandleResponse(resp *http.Response) (ContainersGetResponse, error) {
	result := ContainersGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Container); err != nil {
		return ContainersGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ContainersClient) getHandleError(resp *http.Response) error {
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

// ListByStorageAccount - Lists all the containers of a storage Account in a Data Box Edge/Data Box Gateway device.
// If the operation fails it returns the *CloudError error type.
func (client *ContainersClient) ListByStorageAccount(deviceName string, storageAccountName string, resourceGroupName string, options *ContainersListByStorageAccountOptions) *ContainersListByStorageAccountPager {
	return &ContainersListByStorageAccountPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByStorageAccountCreateRequest(ctx, deviceName, storageAccountName, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp ContainersListByStorageAccountResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ContainerList.NextLink)
		},
	}
}

// listByStorageAccountCreateRequest creates the ListByStorageAccount request.
func (client *ContainersClient) listByStorageAccountCreateRequest(ctx context.Context, deviceName string, storageAccountName string, resourceGroupName string, options *ContainersListByStorageAccountOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/storageAccounts/{storageAccountName}/containers"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if storageAccountName == "" {
		return nil, errors.New("parameter storageAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageAccountName}", url.PathEscape(storageAccountName))
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
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByStorageAccountHandleResponse handles the ListByStorageAccount response.
func (client *ContainersClient) listByStorageAccountHandleResponse(resp *http.Response) (ContainersListByStorageAccountResponse, error) {
	result := ContainersListByStorageAccountResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ContainerList); err != nil {
		return ContainersListByStorageAccountResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByStorageAccountHandleError handles the ListByStorageAccount error response.
func (client *ContainersClient) listByStorageAccountHandleError(resp *http.Response) error {
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

// BeginRefresh - Refreshes the container metadata with the data from the cloud.
// If the operation fails it returns the *CloudError error type.
func (client *ContainersClient) BeginRefresh(ctx context.Context, deviceName string, storageAccountName string, containerName string, resourceGroupName string, options *ContainersBeginRefreshOptions) (ContainersRefreshPollerResponse, error) {
	resp, err := client.refresh(ctx, deviceName, storageAccountName, containerName, resourceGroupName, options)
	if err != nil {
		return ContainersRefreshPollerResponse{}, err
	}
	result := ContainersRefreshPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ContainersClient.Refresh", "", resp, client.pl, client.refreshHandleError)
	if err != nil {
		return ContainersRefreshPollerResponse{}, err
	}
	result.Poller = &ContainersRefreshPoller{
		pt: pt,
	}
	return result, nil
}

// Refresh - Refreshes the container metadata with the data from the cloud.
// If the operation fails it returns the *CloudError error type.
func (client *ContainersClient) refresh(ctx context.Context, deviceName string, storageAccountName string, containerName string, resourceGroupName string, options *ContainersBeginRefreshOptions) (*http.Response, error) {
	req, err := client.refreshCreateRequest(ctx, deviceName, storageAccountName, containerName, resourceGroupName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.refreshHandleError(resp)
	}
	return resp, nil
}

// refreshCreateRequest creates the Refresh request.
func (client *ContainersClient) refreshCreateRequest(ctx context.Context, deviceName string, storageAccountName string, containerName string, resourceGroupName string, options *ContainersBeginRefreshOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/storageAccounts/{storageAccountName}/containers/{containerName}/refresh"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if storageAccountName == "" {
		return nil, errors.New("parameter storageAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageAccountName}", url.PathEscape(storageAccountName))
	if containerName == "" {
		return nil, errors.New("parameter containerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerName}", url.PathEscape(containerName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
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

// refreshHandleError handles the Refresh error response.
func (client *ContainersClient) refreshHandleError(resp *http.Response) error {
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
