// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armoperationalinsights

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// SavedSearchesClient contains the methods for the SavedSearches group.
// Don't use this type directly, use NewSavedSearchesClient() instead.
type SavedSearchesClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewSavedSearchesClient creates a new instance of SavedSearchesClient with the specified values.
func NewSavedSearchesClient(con *armcore.Connection, subscriptionID string) *SavedSearchesClient {
	return &SavedSearchesClient{con: con, subscriptionID: subscriptionID}
}

// CreateOrUpdate - Creates or updates a saved search for a given workspace.
// If the operation fails it returns a generic error.
func (client *SavedSearchesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, savedSearchID string, parameters SavedSearch, options *SavedSearchesCreateOrUpdateOptions) (SavedSearchesCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, savedSearchID, parameters, options)
	if err != nil {
		return SavedSearchesCreateOrUpdateResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return SavedSearchesCreateOrUpdateResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return SavedSearchesCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SavedSearchesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, savedSearchID string, parameters SavedSearch, options *SavedSearchesCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/savedSearches/{savedSearchId}"
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
	if savedSearchID == "" {
		return nil, errors.New("parameter savedSearchID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{savedSearchId}", url.PathEscape(savedSearchID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-08-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *SavedSearchesClient) createOrUpdateHandleResponse(resp *azcore.Response) (SavedSearchesCreateOrUpdateResponse, error) {
	result := SavedSearchesCreateOrUpdateResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.SavedSearch); err != nil {
		return SavedSearchesCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *SavedSearchesClient) createOrUpdateHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// Delete - Deletes the specified saved search in a given workspace.
// If the operation fails it returns a generic error.
func (client *SavedSearchesClient) Delete(ctx context.Context, resourceGroupName string, workspaceName string, savedSearchID string, options *SavedSearchesDeleteOptions) (SavedSearchesDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, savedSearchID, options)
	if err != nil {
		return SavedSearchesDeleteResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return SavedSearchesDeleteResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return SavedSearchesDeleteResponse{}, client.deleteHandleError(resp)
	}
	return SavedSearchesDeleteResponse{RawResponse: resp.Response}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SavedSearchesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, savedSearchID string, options *SavedSearchesDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/savedSearches/{savedSearchId}"
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
	if savedSearchID == "" {
		return nil, errors.New("parameter savedSearchID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{savedSearchId}", url.PathEscape(savedSearchID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-08-01")
	req.URL.RawQuery = reqQP.Encode()
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *SavedSearchesClient) deleteHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// Get - Gets the specified saved search for a given workspace.
// If the operation fails it returns a generic error.
func (client *SavedSearchesClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, savedSearchID string, options *SavedSearchesGetOptions) (SavedSearchesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, savedSearchID, options)
	if err != nil {
		return SavedSearchesGetResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return SavedSearchesGetResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return SavedSearchesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SavedSearchesClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, savedSearchID string, options *SavedSearchesGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/savedSearches/{savedSearchId}"
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
	if savedSearchID == "" {
		return nil, errors.New("parameter savedSearchID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{savedSearchId}", url.PathEscape(savedSearchID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-08-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SavedSearchesClient) getHandleResponse(resp *azcore.Response) (SavedSearchesGetResponse, error) {
	result := SavedSearchesGetResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.SavedSearch); err != nil {
		return SavedSearchesGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *SavedSearchesClient) getHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// ListByWorkspace - Gets the saved searches for a given Log Analytics Workspace
// If the operation fails it returns a generic error.
func (client *SavedSearchesClient) ListByWorkspace(ctx context.Context, resourceGroupName string, workspaceName string, options *SavedSearchesListByWorkspaceOptions) (SavedSearchesListByWorkspaceResponse, error) {
	req, err := client.listByWorkspaceCreateRequest(ctx, resourceGroupName, workspaceName, options)
	if err != nil {
		return SavedSearchesListByWorkspaceResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return SavedSearchesListByWorkspaceResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return SavedSearchesListByWorkspaceResponse{}, client.listByWorkspaceHandleError(resp)
	}
	return client.listByWorkspaceHandleResponse(resp)
}

// listByWorkspaceCreateRequest creates the ListByWorkspace request.
func (client *SavedSearchesClient) listByWorkspaceCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *SavedSearchesListByWorkspaceOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/savedSearches"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-08-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listByWorkspaceHandleResponse handles the ListByWorkspace response.
func (client *SavedSearchesClient) listByWorkspaceHandleResponse(resp *azcore.Response) (SavedSearchesListByWorkspaceResponse, error) {
	result := SavedSearchesListByWorkspaceResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.SavedSearchesListResult); err != nil {
		return SavedSearchesListByWorkspaceResponse{}, err
	}
	return result, nil
}

// listByWorkspaceHandleError handles the ListByWorkspace error response.
func (client *SavedSearchesClient) listByWorkspaceHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}
