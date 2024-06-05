//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhybridcompute

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

// SettingsClient contains the methods for the Settings group.
// Don't use this type directly, use NewSettingsClient() instead.
type SettingsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSettingsClient creates a new instance of SettingsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSettingsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SettingsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SettingsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Returns the base Settings for the target resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-31-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - baseProvider - The name of the base Resource Provider.
//   - baseResourceType - The name of the base Resource Type.
//   - baseResourceName - The name of the base resource.
//   - settingsResourceName - The name of the settings resource.
//   - options - SettingsClientGetOptions contains the optional parameters for the SettingsClient.Get method.
func (client *SettingsClient) Get(ctx context.Context, resourceGroupName string, baseProvider string, baseResourceType string, baseResourceName string, settingsResourceName string, options *SettingsClientGetOptions) (SettingsClientGetResponse, error) {
	var err error
	const operationName = "SettingsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, baseProvider, baseResourceType, baseResourceName, settingsResourceName, options)
	if err != nil {
		return SettingsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SettingsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SettingsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *SettingsClient) getCreateRequest(ctx context.Context, resourceGroupName string, baseProvider string, baseResourceType string, baseResourceName string, settingsResourceName string, options *SettingsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{baseProvider}/{baseResourceType}/{baseResourceName}/providers/Microsoft.HybridCompute/settings/{settingsResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if baseProvider == "" {
		return nil, errors.New("parameter baseProvider cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{baseProvider}", url.PathEscape(baseProvider))
	if baseResourceType == "" {
		return nil, errors.New("parameter baseResourceType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{baseResourceType}", url.PathEscape(baseResourceType))
	if baseResourceName == "" {
		return nil, errors.New("parameter baseResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{baseResourceName}", url.PathEscape(baseResourceName))
	if settingsResourceName == "" {
		return nil, errors.New("parameter settingsResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{settingsResourceName}", url.PathEscape(settingsResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-31-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SettingsClient) getHandleResponse(resp *http.Response) (SettingsClientGetResponse, error) {
	result := SettingsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Settings); err != nil {
		return SettingsClientGetResponse{}, err
	}
	return result, nil
}

// Patch - Update the base Settings of the target resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-31-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - baseProvider - The name of the base Resource Provider.
//   - baseResourceType - The name of the base Resource Type.
//   - baseResourceName - The name of the base resource.
//   - settingsResourceName - The name of the settings resource.
//   - parameters - Settings details
//   - options - SettingsClientPatchOptions contains the optional parameters for the SettingsClient.Patch method.
func (client *SettingsClient) Patch(ctx context.Context, resourceGroupName string, baseProvider string, baseResourceType string, baseResourceName string, settingsResourceName string, parameters Settings, options *SettingsClientPatchOptions) (SettingsClientPatchResponse, error) {
	var err error
	const operationName = "SettingsClient.Patch"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.patchCreateRequest(ctx, resourceGroupName, baseProvider, baseResourceType, baseResourceName, settingsResourceName, parameters, options)
	if err != nil {
		return SettingsClientPatchResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SettingsClientPatchResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SettingsClientPatchResponse{}, err
	}
	resp, err := client.patchHandleResponse(httpResp)
	return resp, err
}

// patchCreateRequest creates the Patch request.
func (client *SettingsClient) patchCreateRequest(ctx context.Context, resourceGroupName string, baseProvider string, baseResourceType string, baseResourceName string, settingsResourceName string, parameters Settings, options *SettingsClientPatchOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{baseProvider}/{baseResourceType}/{baseResourceName}/providers/Microsoft.HybridCompute/settings/{settingsResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if baseProvider == "" {
		return nil, errors.New("parameter baseProvider cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{baseProvider}", url.PathEscape(baseProvider))
	if baseResourceType == "" {
		return nil, errors.New("parameter baseResourceType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{baseResourceType}", url.PathEscape(baseResourceType))
	if baseResourceName == "" {
		return nil, errors.New("parameter baseResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{baseResourceName}", url.PathEscape(baseResourceName))
	if settingsResourceName == "" {
		return nil, errors.New("parameter settingsResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{settingsResourceName}", url.PathEscape(settingsResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-31-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// patchHandleResponse handles the Patch response.
func (client *SettingsClient) patchHandleResponse(resp *http.Response) (SettingsClientPatchResponse, error) {
	result := SettingsClientPatchResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Settings); err != nil {
		return SettingsClientPatchResponse{}, err
	}
	return result, nil
}

// Update - Updates the base Settings of the target resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-31-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - baseProvider - The name of the base Resource Provider.
//   - baseResourceType - The name of the base Resource Type.
//   - baseResourceName - The name of the base resource.
//   - settingsResourceName - The name of the settings resource.
//   - parameters - Settings details
//   - options - SettingsClientUpdateOptions contains the optional parameters for the SettingsClient.Update method.
func (client *SettingsClient) Update(ctx context.Context, resourceGroupName string, baseProvider string, baseResourceType string, baseResourceName string, settingsResourceName string, parameters Settings, options *SettingsClientUpdateOptions) (SettingsClientUpdateResponse, error) {
	var err error
	const operationName = "SettingsClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, baseProvider, baseResourceType, baseResourceName, settingsResourceName, parameters, options)
	if err != nil {
		return SettingsClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SettingsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return SettingsClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *SettingsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, baseProvider string, baseResourceType string, baseResourceName string, settingsResourceName string, parameters Settings, options *SettingsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{baseProvider}/{baseResourceType}/{baseResourceName}/providers/Microsoft.HybridCompute/settings/{settingsResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if baseProvider == "" {
		return nil, errors.New("parameter baseProvider cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{baseProvider}", url.PathEscape(baseProvider))
	if baseResourceType == "" {
		return nil, errors.New("parameter baseResourceType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{baseResourceType}", url.PathEscape(baseResourceType))
	if baseResourceName == "" {
		return nil, errors.New("parameter baseResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{baseResourceName}", url.PathEscape(baseResourceName))
	if settingsResourceName == "" {
		return nil, errors.New("parameter settingsResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{settingsResourceName}", url.PathEscape(settingsResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-31-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *SettingsClient) updateHandleResponse(resp *http.Response) (SettingsClientUpdateResponse, error) {
	result := SettingsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Settings); err != nil {
		return SettingsClientUpdateResponse{}, err
	}
	return result, nil
}
