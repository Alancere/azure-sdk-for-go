//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmonitor

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

// DiagnosticSettingsClient contains the methods for the DiagnosticSettings group.
// Don't use this type directly, use NewDiagnosticSettingsClient() instead.
type DiagnosticSettingsClient struct {
	ep string
	pl runtime.Pipeline
}

// NewDiagnosticSettingsClient creates a new instance of DiagnosticSettingsClient with the specified values.
func NewDiagnosticSettingsClient(credential azcore.TokenCredential, options *arm.ClientOptions) *DiagnosticSettingsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &DiagnosticSettingsClient{ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// CreateOrUpdate - Creates or updates diagnostic settings for the specified resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *DiagnosticSettingsClient) CreateOrUpdate(ctx context.Context, resourceURI string, name string, parameters DiagnosticSettingsResource, options *DiagnosticSettingsCreateOrUpdateOptions) (DiagnosticSettingsCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceURI, name, parameters, options)
	if err != nil {
		return DiagnosticSettingsCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DiagnosticSettingsCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DiagnosticSettingsCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DiagnosticSettingsClient) createOrUpdateCreateRequest(ctx context.Context, resourceURI string, name string, parameters DiagnosticSettingsResource, options *DiagnosticSettingsCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.Insights/diagnosticSettings/{name}"
	if resourceURI == "" {
		return nil, errors.New("parameter resourceURI cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *DiagnosticSettingsClient) createOrUpdateHandleResponse(resp *http.Response) (DiagnosticSettingsCreateOrUpdateResponse, error) {
	result := DiagnosticSettingsCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DiagnosticSettingsResource); err != nil {
		return DiagnosticSettingsCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *DiagnosticSettingsClient) createOrUpdateHandleError(resp *http.Response) error {
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

// Delete - Deletes existing diagnostic settings for the specified resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *DiagnosticSettingsClient) Delete(ctx context.Context, resourceURI string, name string, options *DiagnosticSettingsDeleteOptions) (DiagnosticSettingsDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceURI, name, options)
	if err != nil {
		return DiagnosticSettingsDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DiagnosticSettingsDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return DiagnosticSettingsDeleteResponse{}, client.deleteHandleError(resp)
	}
	return DiagnosticSettingsDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DiagnosticSettingsClient) deleteCreateRequest(ctx context.Context, resourceURI string, name string, options *DiagnosticSettingsDeleteOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.Insights/diagnosticSettings/{name}"
	if resourceURI == "" {
		return nil, errors.New("parameter resourceURI cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *DiagnosticSettingsClient) deleteHandleError(resp *http.Response) error {
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

// Get - Gets the active diagnostic settings for the specified resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *DiagnosticSettingsClient) Get(ctx context.Context, resourceURI string, name string, options *DiagnosticSettingsGetOptions) (DiagnosticSettingsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceURI, name, options)
	if err != nil {
		return DiagnosticSettingsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DiagnosticSettingsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DiagnosticSettingsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DiagnosticSettingsClient) getCreateRequest(ctx context.Context, resourceURI string, name string, options *DiagnosticSettingsGetOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.Insights/diagnosticSettings/{name}"
	if resourceURI == "" {
		return nil, errors.New("parameter resourceURI cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DiagnosticSettingsClient) getHandleResponse(resp *http.Response) (DiagnosticSettingsGetResponse, error) {
	result := DiagnosticSettingsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DiagnosticSettingsResource); err != nil {
		return DiagnosticSettingsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *DiagnosticSettingsClient) getHandleError(resp *http.Response) error {
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

// List - Gets the active diagnostic settings list for the specified resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *DiagnosticSettingsClient) List(ctx context.Context, resourceURI string, options *DiagnosticSettingsListOptions) (DiagnosticSettingsListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceURI, options)
	if err != nil {
		return DiagnosticSettingsListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DiagnosticSettingsListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DiagnosticSettingsListResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *DiagnosticSettingsClient) listCreateRequest(ctx context.Context, resourceURI string, options *DiagnosticSettingsListOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.Insights/diagnosticSettings"
	if resourceURI == "" {
		return nil, errors.New("parameter resourceURI cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DiagnosticSettingsClient) listHandleResponse(resp *http.Response) (DiagnosticSettingsListResponse, error) {
	result := DiagnosticSettingsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DiagnosticSettingsResourceCollection); err != nil {
		return DiagnosticSettingsListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *DiagnosticSettingsClient) listHandleError(resp *http.Response) error {
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
