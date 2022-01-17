//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmediaservices

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
	"strconv"
	"strings"
)

// LiveEventsClient contains the methods for the LiveEvents group.
// Don't use this type directly, use NewLiveEventsClient() instead.
type LiveEventsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewLiveEventsClient creates a new instance of LiveEventsClient with the specified values.
// subscriptionID - The unique identifier for a Microsoft Azure subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewLiveEventsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *LiveEventsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &LiveEventsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// BeginAllocate - A live event is in StandBy state after allocation completes, and is ready to start.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the Azure subscription.
// accountName - The Media Services account name.
// liveEventName - The name of the live event, maximum length is 32.
// options - LiveEventsClientBeginAllocateOptions contains the optional parameters for the LiveEventsClient.BeginAllocate
// method.
func (client *LiveEventsClient) BeginAllocate(ctx context.Context, resourceGroupName string, accountName string, liveEventName string, options *LiveEventsClientBeginAllocateOptions) (LiveEventsClientAllocatePollerResponse, error) {
	resp, err := client.allocate(ctx, resourceGroupName, accountName, liveEventName, options)
	if err != nil {
		return LiveEventsClientAllocatePollerResponse{}, err
	}
	result := LiveEventsClientAllocatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("LiveEventsClient.Allocate", "", resp, client.pl)
	if err != nil {
		return LiveEventsClientAllocatePollerResponse{}, err
	}
	result.Poller = &LiveEventsClientAllocatePoller{
		pt: pt,
	}
	return result, nil
}

// Allocate - A live event is in StandBy state after allocation completes, and is ready to start.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *LiveEventsClient) allocate(ctx context.Context, resourceGroupName string, accountName string, liveEventName string, options *LiveEventsClientBeginAllocateOptions) (*http.Response, error) {
	req, err := client.allocateCreateRequest(ctx, resourceGroupName, accountName, liveEventName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// allocateCreateRequest creates the Allocate request.
func (client *LiveEventsClient) allocateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, liveEventName string, options *LiveEventsClientBeginAllocateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/liveEvents/{liveEventName}/allocate"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if liveEventName == "" {
		return nil, errors.New("parameter liveEventName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{liveEventName}", url.PathEscape(liveEventName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// BeginCreate - Creates a new live event.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the Azure subscription.
// accountName - The Media Services account name.
// liveEventName - The name of the live event, maximum length is 32.
// parameters - Live event properties needed for creation.
// options - LiveEventsClientBeginCreateOptions contains the optional parameters for the LiveEventsClient.BeginCreate method.
func (client *LiveEventsClient) BeginCreate(ctx context.Context, resourceGroupName string, accountName string, liveEventName string, parameters LiveEvent, options *LiveEventsClientBeginCreateOptions) (LiveEventsClientCreatePollerResponse, error) {
	resp, err := client.create(ctx, resourceGroupName, accountName, liveEventName, parameters, options)
	if err != nil {
		return LiveEventsClientCreatePollerResponse{}, err
	}
	result := LiveEventsClientCreatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("LiveEventsClient.Create", "", resp, client.pl)
	if err != nil {
		return LiveEventsClientCreatePollerResponse{}, err
	}
	result.Poller = &LiveEventsClientCreatePoller{
		pt: pt,
	}
	return result, nil
}

// Create - Creates a new live event.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *LiveEventsClient) create(ctx context.Context, resourceGroupName string, accountName string, liveEventName string, parameters LiveEvent, options *LiveEventsClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, accountName, liveEventName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *LiveEventsClient) createCreateRequest(ctx context.Context, resourceGroupName string, accountName string, liveEventName string, parameters LiveEvent, options *LiveEventsClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/liveEvents/{liveEventName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if liveEventName == "" {
		return nil, errors.New("parameter liveEventName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{liveEventName}", url.PathEscape(liveEventName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	if options != nil && options.AutoStart != nil {
		reqQP.Set("autoStart", strconv.FormatBool(*options.AutoStart))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes a live event.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the Azure subscription.
// accountName - The Media Services account name.
// liveEventName - The name of the live event, maximum length is 32.
// options - LiveEventsClientBeginDeleteOptions contains the optional parameters for the LiveEventsClient.BeginDelete method.
func (client *LiveEventsClient) BeginDelete(ctx context.Context, resourceGroupName string, accountName string, liveEventName string, options *LiveEventsClientBeginDeleteOptions) (LiveEventsClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, accountName, liveEventName, options)
	if err != nil {
		return LiveEventsClientDeletePollerResponse{}, err
	}
	result := LiveEventsClientDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("LiveEventsClient.Delete", "", resp, client.pl)
	if err != nil {
		return LiveEventsClientDeletePollerResponse{}, err
	}
	result.Poller = &LiveEventsClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes a live event.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *LiveEventsClient) deleteOperation(ctx context.Context, resourceGroupName string, accountName string, liveEventName string, options *LiveEventsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, liveEventName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *LiveEventsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, liveEventName string, options *LiveEventsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/liveEvents/{liveEventName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if liveEventName == "" {
		return nil, errors.New("parameter liveEventName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{liveEventName}", url.PathEscape(liveEventName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets properties of a live event.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the Azure subscription.
// accountName - The Media Services account name.
// liveEventName - The name of the live event, maximum length is 32.
// options - LiveEventsClientGetOptions contains the optional parameters for the LiveEventsClient.Get method.
func (client *LiveEventsClient) Get(ctx context.Context, resourceGroupName string, accountName string, liveEventName string, options *LiveEventsClientGetOptions) (LiveEventsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, liveEventName, options)
	if err != nil {
		return LiveEventsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LiveEventsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LiveEventsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *LiveEventsClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, liveEventName string, options *LiveEventsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/liveEvents/{liveEventName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if liveEventName == "" {
		return nil, errors.New("parameter liveEventName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{liveEventName}", url.PathEscape(liveEventName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
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
func (client *LiveEventsClient) getHandleResponse(resp *http.Response) (LiveEventsClientGetResponse, error) {
	result := LiveEventsClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.LiveEvent); err != nil {
		return LiveEventsClientGetResponse{}, err
	}
	return result, nil
}

// List - Lists all the live events in the account.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the Azure subscription.
// accountName - The Media Services account name.
// options - LiveEventsClientListOptions contains the optional parameters for the LiveEventsClient.List method.
func (client *LiveEventsClient) List(resourceGroupName string, accountName string, options *LiveEventsClientListOptions) *LiveEventsClientListPager {
	return &LiveEventsClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, accountName, options)
		},
		advancer: func(ctx context.Context, resp LiveEventsClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.LiveEventListResult.ODataNextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *LiveEventsClient) listCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *LiveEventsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/liveEvents"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *LiveEventsClient) listHandleResponse(resp *http.Response) (LiveEventsClientListResponse, error) {
	result := LiveEventsClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.LiveEventListResult); err != nil {
		return LiveEventsClientListResponse{}, err
	}
	return result, nil
}

// BeginReset - Resets an existing live event. All live outputs for the live event are deleted and the live event is stopped
// and will be started again. All assets used by the live outputs and streaming locators
// created on these assets are unaffected.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the Azure subscription.
// accountName - The Media Services account name.
// liveEventName - The name of the live event, maximum length is 32.
// options - LiveEventsClientBeginResetOptions contains the optional parameters for the LiveEventsClient.BeginReset method.
func (client *LiveEventsClient) BeginReset(ctx context.Context, resourceGroupName string, accountName string, liveEventName string, options *LiveEventsClientBeginResetOptions) (LiveEventsClientResetPollerResponse, error) {
	resp, err := client.reset(ctx, resourceGroupName, accountName, liveEventName, options)
	if err != nil {
		return LiveEventsClientResetPollerResponse{}, err
	}
	result := LiveEventsClientResetPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("LiveEventsClient.Reset", "", resp, client.pl)
	if err != nil {
		return LiveEventsClientResetPollerResponse{}, err
	}
	result.Poller = &LiveEventsClientResetPoller{
		pt: pt,
	}
	return result, nil
}

// Reset - Resets an existing live event. All live outputs for the live event are deleted and the live event is stopped and
// will be started again. All assets used by the live outputs and streaming locators
// created on these assets are unaffected.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *LiveEventsClient) reset(ctx context.Context, resourceGroupName string, accountName string, liveEventName string, options *LiveEventsClientBeginResetOptions) (*http.Response, error) {
	req, err := client.resetCreateRequest(ctx, resourceGroupName, accountName, liveEventName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// resetCreateRequest creates the Reset request.
func (client *LiveEventsClient) resetCreateRequest(ctx context.Context, resourceGroupName string, accountName string, liveEventName string, options *LiveEventsClientBeginResetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/liveEvents/{liveEventName}/reset"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if liveEventName == "" {
		return nil, errors.New("parameter liveEventName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{liveEventName}", url.PathEscape(liveEventName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// BeginStart - A live event in Stopped or StandBy state will be in Running state after the start operation completes.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the Azure subscription.
// accountName - The Media Services account name.
// liveEventName - The name of the live event, maximum length is 32.
// options - LiveEventsClientBeginStartOptions contains the optional parameters for the LiveEventsClient.BeginStart method.
func (client *LiveEventsClient) BeginStart(ctx context.Context, resourceGroupName string, accountName string, liveEventName string, options *LiveEventsClientBeginStartOptions) (LiveEventsClientStartPollerResponse, error) {
	resp, err := client.start(ctx, resourceGroupName, accountName, liveEventName, options)
	if err != nil {
		return LiveEventsClientStartPollerResponse{}, err
	}
	result := LiveEventsClientStartPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("LiveEventsClient.Start", "", resp, client.pl)
	if err != nil {
		return LiveEventsClientStartPollerResponse{}, err
	}
	result.Poller = &LiveEventsClientStartPoller{
		pt: pt,
	}
	return result, nil
}

// Start - A live event in Stopped or StandBy state will be in Running state after the start operation completes.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *LiveEventsClient) start(ctx context.Context, resourceGroupName string, accountName string, liveEventName string, options *LiveEventsClientBeginStartOptions) (*http.Response, error) {
	req, err := client.startCreateRequest(ctx, resourceGroupName, accountName, liveEventName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// startCreateRequest creates the Start request.
func (client *LiveEventsClient) startCreateRequest(ctx context.Context, resourceGroupName string, accountName string, liveEventName string, options *LiveEventsClientBeginStartOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/liveEvents/{liveEventName}/start"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if liveEventName == "" {
		return nil, errors.New("parameter liveEventName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{liveEventName}", url.PathEscape(liveEventName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// BeginStop - Stops a running live event.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the Azure subscription.
// accountName - The Media Services account name.
// liveEventName - The name of the live event, maximum length is 32.
// parameters - LiveEvent stop parameters
// options - LiveEventsClientBeginStopOptions contains the optional parameters for the LiveEventsClient.BeginStop method.
func (client *LiveEventsClient) BeginStop(ctx context.Context, resourceGroupName string, accountName string, liveEventName string, parameters LiveEventActionInput, options *LiveEventsClientBeginStopOptions) (LiveEventsClientStopPollerResponse, error) {
	resp, err := client.stop(ctx, resourceGroupName, accountName, liveEventName, parameters, options)
	if err != nil {
		return LiveEventsClientStopPollerResponse{}, err
	}
	result := LiveEventsClientStopPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("LiveEventsClient.Stop", "", resp, client.pl)
	if err != nil {
		return LiveEventsClientStopPollerResponse{}, err
	}
	result.Poller = &LiveEventsClientStopPoller{
		pt: pt,
	}
	return result, nil
}

// Stop - Stops a running live event.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *LiveEventsClient) stop(ctx context.Context, resourceGroupName string, accountName string, liveEventName string, parameters LiveEventActionInput, options *LiveEventsClientBeginStopOptions) (*http.Response, error) {
	req, err := client.stopCreateRequest(ctx, resourceGroupName, accountName, liveEventName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// stopCreateRequest creates the Stop request.
func (client *LiveEventsClient) stopCreateRequest(ctx context.Context, resourceGroupName string, accountName string, liveEventName string, parameters LiveEventActionInput, options *LiveEventsClientBeginStopOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/liveEvents/{liveEventName}/stop"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if liveEventName == "" {
		return nil, errors.New("parameter liveEventName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{liveEventName}", url.PathEscape(liveEventName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginUpdate - Updates settings on an existing live event.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the Azure subscription.
// accountName - The Media Services account name.
// liveEventName - The name of the live event, maximum length is 32.
// parameters - Live event properties needed for patch.
// options - LiveEventsClientBeginUpdateOptions contains the optional parameters for the LiveEventsClient.BeginUpdate method.
func (client *LiveEventsClient) BeginUpdate(ctx context.Context, resourceGroupName string, accountName string, liveEventName string, parameters LiveEvent, options *LiveEventsClientBeginUpdateOptions) (LiveEventsClientUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, accountName, liveEventName, parameters, options)
	if err != nil {
		return LiveEventsClientUpdatePollerResponse{}, err
	}
	result := LiveEventsClientUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("LiveEventsClient.Update", "", resp, client.pl)
	if err != nil {
		return LiveEventsClientUpdatePollerResponse{}, err
	}
	result.Poller = &LiveEventsClientUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Updates settings on an existing live event.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *LiveEventsClient) update(ctx context.Context, resourceGroupName string, accountName string, liveEventName string, parameters LiveEvent, options *LiveEventsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, accountName, liveEventName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *LiveEventsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, liveEventName string, parameters LiveEvent, options *LiveEventsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/liveEvents/{liveEventName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if liveEventName == "" {
		return nil, errors.New("parameter liveEventName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{liveEventName}", url.PathEscape(liveEventName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}
