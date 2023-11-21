//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapplicationinsights

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// WorkbooksClient contains the methods for the Workbooks group.
// Don't use this type directly, use NewWorkbooksClient() instead.
type WorkbooksClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewWorkbooksClient creates a new instance of WorkbooksClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewWorkbooksClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*WorkbooksClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &WorkbooksClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Create a new workbook.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2015-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - The name of the Application Insights component resource.
//   - workbookProperties - Properties that need to be specified to create a new workbook.
//   - options - WorkbooksClientCreateOrUpdateOptions contains the optional parameters for the WorkbooksClient.CreateOrUpdate
//     method.
func (client *WorkbooksClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, workbookProperties Workbook, options *WorkbooksClientCreateOrUpdateOptions) (WorkbooksClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "WorkbooksClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, resourceName, workbookProperties, options)
	if err != nil {
		return WorkbooksClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkbooksClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return WorkbooksClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *WorkbooksClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, workbookProperties Workbook, options *WorkbooksClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/workbooks/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, workbookProperties); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *WorkbooksClient) createOrUpdateHandleResponse(resp *http.Response) (WorkbooksClientCreateOrUpdateResponse, error) {
	result := WorkbooksClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Workbook); err != nil {
		return WorkbooksClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete a workbook.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2015-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - The name of the Application Insights component resource.
//   - options - WorkbooksClientDeleteOptions contains the optional parameters for the WorkbooksClient.Delete method.
func (client *WorkbooksClient) Delete(ctx context.Context, resourceGroupName string, resourceName string, options *WorkbooksClientDeleteOptions) (WorkbooksClientDeleteResponse, error) {
	var err error
	const operationName = "WorkbooksClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return WorkbooksClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkbooksClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusCreated, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return WorkbooksClientDeleteResponse{}, err
	}
	return WorkbooksClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *WorkbooksClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *WorkbooksClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/workbooks/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a single workbook by its resourceName.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2015-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - The name of the Application Insights component resource.
//   - options - WorkbooksClientGetOptions contains the optional parameters for the WorkbooksClient.Get method.
func (client *WorkbooksClient) Get(ctx context.Context, resourceGroupName string, resourceName string, options *WorkbooksClientGetOptions) (WorkbooksClientGetResponse, error) {
	var err error
	const operationName = "WorkbooksClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return WorkbooksClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkbooksClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return WorkbooksClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *WorkbooksClient) getCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *WorkbooksClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/workbooks/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *WorkbooksClient) getHandleResponse(resp *http.Response) (WorkbooksClientGetResponse, error) {
	result := WorkbooksClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Workbook); err != nil {
		return WorkbooksClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Get all Workbooks defined within a specified resource group and category.
//
// Generated from API version 2015-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - category - Category of workbook to return.
//   - options - WorkbooksClientListByResourceGroupOptions contains the optional parameters for the WorkbooksClient.NewListByResourceGroupPager
//     method.
func (client *WorkbooksClient) NewListByResourceGroupPager(resourceGroupName string, category CategoryType, options *WorkbooksClientListByResourceGroupOptions) *runtime.Pager[WorkbooksClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[WorkbooksClientListByResourceGroupResponse]{
		More: func(page WorkbooksClientListByResourceGroupResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *WorkbooksClientListByResourceGroupResponse) (WorkbooksClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "WorkbooksClient.NewListByResourceGroupPager")
			req, err := client.listByResourceGroupCreateRequest(ctx, resourceGroupName, category, options)
			if err != nil {
				return WorkbooksClientListByResourceGroupResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return WorkbooksClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return WorkbooksClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *WorkbooksClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, category CategoryType, options *WorkbooksClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/workbooks"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("category", string(category))
	if options != nil && options.Tags != nil {
		reqQP.Set("tags", strings.Join(options.Tags, ","))
	}
	if options != nil && options.CanFetchContent != nil {
		reqQP.Set("canFetchContent", strconv.FormatBool(*options.CanFetchContent))
	}
	reqQP.Set("api-version", "2015-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *WorkbooksClient) listByResourceGroupHandleResponse(resp *http.Response) (WorkbooksClientListByResourceGroupResponse, error) {
	result := WorkbooksClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkbooksListResult); err != nil {
		return WorkbooksClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// Update - Updates a workbook that has already been added.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2015-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - The name of the Application Insights component resource.
//   - workbookProperties - Properties that need to be specified to create a new workbook.
//   - options - WorkbooksClientUpdateOptions contains the optional parameters for the WorkbooksClient.Update method.
func (client *WorkbooksClient) Update(ctx context.Context, resourceGroupName string, resourceName string, workbookProperties Workbook, options *WorkbooksClientUpdateOptions) (WorkbooksClientUpdateResponse, error) {
	var err error
	const operationName = "WorkbooksClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, resourceName, workbookProperties, options)
	if err != nil {
		return WorkbooksClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkbooksClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return WorkbooksClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *WorkbooksClient) updateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, workbookProperties Workbook, options *WorkbooksClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/workbooks/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, workbookProperties); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *WorkbooksClient) updateHandleResponse(resp *http.Response) (WorkbooksClientUpdateResponse, error) {
	result := WorkbooksClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Workbook); err != nil {
		return WorkbooksClientUpdateResponse{}, err
	}
	return result, nil
}
