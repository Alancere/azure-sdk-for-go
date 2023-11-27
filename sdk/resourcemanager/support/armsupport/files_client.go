//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsupport

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

// FilesClient contains the methods for the Files group.
// Don't use this type directly, use NewFilesClient() instead.
type FilesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewFilesClient creates a new instance of FilesClient with the specified values.
//   - subscriptionID - Azure subscription Id.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewFilesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*FilesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &FilesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Create - Creates a new file under a workspace for the specified subscription.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01-preview
//   - fileWorkspaceName - File workspace name.
//   - fileName - File name.
//   - createFileParameters - Create file object
//   - options - FilesClientCreateOptions contains the optional parameters for the FilesClient.Create method.
func (client *FilesClient) Create(ctx context.Context, fileWorkspaceName string, fileName string, createFileParameters FileDetails, options *FilesClientCreateOptions) (FilesClientCreateResponse, error) {
	var err error
	const operationName = "FilesClient.Create"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, fileWorkspaceName, fileName, createFileParameters, options)
	if err != nil {
		return FilesClientCreateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return FilesClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return FilesClientCreateResponse{}, err
	}
	resp, err := client.createHandleResponse(httpResp)
	return resp, err
}

// createCreateRequest creates the Create request.
func (client *FilesClient) createCreateRequest(ctx context.Context, fileWorkspaceName string, fileName string, createFileParameters FileDetails, options *FilesClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Support/fileWorkspaces/{fileWorkspaceName}/files/{fileName}"
	if fileWorkspaceName == "" {
		return nil, errors.New("parameter fileWorkspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fileWorkspaceName}", url.PathEscape(fileWorkspaceName))
	if fileName == "" {
		return nil, errors.New("parameter fileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fileName}", url.PathEscape(fileName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, createFileParameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *FilesClient) createHandleResponse(resp *http.Response) (FilesClientCreateResponse, error) {
	result := FilesClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FileDetails); err != nil {
		return FilesClientCreateResponse{}, err
	}
	return result, nil
}

// Get - Returns details of a specific file in a work space.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01-preview
//   - fileWorkspaceName - File Workspace Name
//   - fileName - File Name
//   - options - FilesClientGetOptions contains the optional parameters for the FilesClient.Get method.
func (client *FilesClient) Get(ctx context.Context, fileWorkspaceName string, fileName string, options *FilesClientGetOptions) (FilesClientGetResponse, error) {
	var err error
	const operationName = "FilesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, fileWorkspaceName, fileName, options)
	if err != nil {
		return FilesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return FilesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return FilesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *FilesClient) getCreateRequest(ctx context.Context, fileWorkspaceName string, fileName string, options *FilesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Support/fileWorkspaces/{fileWorkspaceName}/files/{fileName}"
	if fileWorkspaceName == "" {
		return nil, errors.New("parameter fileWorkspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fileWorkspaceName}", url.PathEscape(fileWorkspaceName))
	if fileName == "" {
		return nil, errors.New("parameter fileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fileName}", url.PathEscape(fileName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *FilesClient) getHandleResponse(resp *http.Response) (FilesClientGetResponse, error) {
	result := FilesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FileDetails); err != nil {
		return FilesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists all the Files information under a workspace for an Azure subscription.
//
// Generated from API version 2022-09-01-preview
//   - fileWorkspaceName - File Workspace Name
//   - options - FilesClientListOptions contains the optional parameters for the FilesClient.NewListPager method.
func (client *FilesClient) NewListPager(fileWorkspaceName string, options *FilesClientListOptions) *runtime.Pager[FilesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[FilesClientListResponse]{
		More: func(page FilesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *FilesClientListResponse) (FilesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "FilesClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, fileWorkspaceName, options)
			}, nil)
			if err != nil {
				return FilesClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *FilesClient) listCreateRequest(ctx context.Context, fileWorkspaceName string, options *FilesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Support/fileWorkspaces/{fileWorkspaceName}/files"
	if fileWorkspaceName == "" {
		return nil, errors.New("parameter fileWorkspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fileWorkspaceName}", url.PathEscape(fileWorkspaceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *FilesClient) listHandleResponse(resp *http.Response) (FilesClientListResponse, error) {
	result := FilesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FilesListResult); err != nil {
		return FilesClientListResponse{}, err
	}
	return result, nil
}

// Upload - This API allows you to upload content to a file
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01-preview
//   - fileWorkspaceName - File WorkspaceName
//   - fileName - File Name
//   - uploadFile - UploadFile object
//   - options - FilesClientUploadOptions contains the optional parameters for the FilesClient.Upload method.
func (client *FilesClient) Upload(ctx context.Context, fileWorkspaceName string, fileName string, uploadFile UploadFile, options *FilesClientUploadOptions) (FilesClientUploadResponse, error) {
	var err error
	const operationName = "FilesClient.Upload"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.uploadCreateRequest(ctx, fileWorkspaceName, fileName, uploadFile, options)
	if err != nil {
		return FilesClientUploadResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return FilesClientUploadResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return FilesClientUploadResponse{}, err
	}
	return FilesClientUploadResponse{}, nil
}

// uploadCreateRequest creates the Upload request.
func (client *FilesClient) uploadCreateRequest(ctx context.Context, fileWorkspaceName string, fileName string, uploadFile UploadFile, options *FilesClientUploadOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Support/fileWorkspaces/{fileWorkspaceName}/files/{fileName}/upload"
	if fileWorkspaceName == "" {
		return nil, errors.New("parameter fileWorkspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fileWorkspaceName}", url.PathEscape(fileWorkspaceName))
	if fileName == "" {
		return nil, errors.New("parameter fileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fileName}", url.PathEscape(fileName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, uploadFile); err != nil {
		return nil, err
	}
	return req, nil
}
