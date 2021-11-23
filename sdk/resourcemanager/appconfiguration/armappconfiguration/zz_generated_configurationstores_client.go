//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappconfiguration

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

// ConfigurationStoresClient contains the methods for the ConfigurationStores group.
// Don't use this type directly, use NewConfigurationStoresClient() instead.
type ConfigurationStoresClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewConfigurationStoresClient creates a new instance of ConfigurationStoresClient with the specified values.
func NewConfigurationStoresClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ConfigurationStoresClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &ConfigurationStoresClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// BeginCreate - Creates a configuration store with the specified parameters.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ConfigurationStoresClient) BeginCreate(ctx context.Context, resourceGroupName string, configStoreName string, configStoreCreationParameters ConfigurationStore, options *ConfigurationStoresBeginCreateOptions) (ConfigurationStoresCreatePollerResponse, error) {
	resp, err := client.create(ctx, resourceGroupName, configStoreName, configStoreCreationParameters, options)
	if err != nil {
		return ConfigurationStoresCreatePollerResponse{}, err
	}
	result := ConfigurationStoresCreatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ConfigurationStoresClient.Create", "", resp, client.pl, client.createHandleError)
	if err != nil {
		return ConfigurationStoresCreatePollerResponse{}, err
	}
	result.Poller = &ConfigurationStoresCreatePoller{
		pt: pt,
	}
	return result, nil
}

// Create - Creates a configuration store with the specified parameters.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ConfigurationStoresClient) create(ctx context.Context, resourceGroupName string, configStoreName string, configStoreCreationParameters ConfigurationStore, options *ConfigurationStoresBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, configStoreName, configStoreCreationParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, client.createHandleError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *ConfigurationStoresClient) createCreateRequest(ctx context.Context, resourceGroupName string, configStoreName string, configStoreCreationParameters ConfigurationStore, options *ConfigurationStoresBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppConfiguration/configurationStores/{configStoreName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if configStoreName == "" {
		return nil, errors.New("parameter configStoreName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configStoreName}", url.PathEscape(configStoreName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, configStoreCreationParameters)
}

// createHandleError handles the Create error response.
func (client *ConfigurationStoresClient) createHandleError(resp *http.Response) error {
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

// BeginDelete - Deletes a configuration store.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ConfigurationStoresClient) BeginDelete(ctx context.Context, resourceGroupName string, configStoreName string, options *ConfigurationStoresBeginDeleteOptions) (ConfigurationStoresDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, configStoreName, options)
	if err != nil {
		return ConfigurationStoresDeletePollerResponse{}, err
	}
	result := ConfigurationStoresDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ConfigurationStoresClient.Delete", "", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return ConfigurationStoresDeletePollerResponse{}, err
	}
	result.Poller = &ConfigurationStoresDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes a configuration store.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ConfigurationStoresClient) deleteOperation(ctx context.Context, resourceGroupName string, configStoreName string, options *ConfigurationStoresBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, configStoreName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ConfigurationStoresClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, configStoreName string, options *ConfigurationStoresBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppConfiguration/configurationStores/{configStoreName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if configStoreName == "" {
		return nil, errors.New("parameter configStoreName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configStoreName}", url.PathEscape(configStoreName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *ConfigurationStoresClient) deleteHandleError(resp *http.Response) error {
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

// Get - Gets the properties of the specified configuration store.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ConfigurationStoresClient) Get(ctx context.Context, resourceGroupName string, configStoreName string, options *ConfigurationStoresGetOptions) (ConfigurationStoresGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, configStoreName, options)
	if err != nil {
		return ConfigurationStoresGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ConfigurationStoresGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ConfigurationStoresGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ConfigurationStoresClient) getCreateRequest(ctx context.Context, resourceGroupName string, configStoreName string, options *ConfigurationStoresGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppConfiguration/configurationStores/{configStoreName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if configStoreName == "" {
		return nil, errors.New("parameter configStoreName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configStoreName}", url.PathEscape(configStoreName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ConfigurationStoresClient) getHandleResponse(resp *http.Response) (ConfigurationStoresGetResponse, error) {
	result := ConfigurationStoresGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConfigurationStore); err != nil {
		return ConfigurationStoresGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ConfigurationStoresClient) getHandleError(resp *http.Response) error {
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

// List - Lists the configuration stores for a given subscription.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ConfigurationStoresClient) List(options *ConfigurationStoresListOptions) *ConfigurationStoresListPager {
	return &ConfigurationStoresListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp ConfigurationStoresListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ConfigurationStoreListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *ConfigurationStoresClient) listCreateRequest(ctx context.Context, options *ConfigurationStoresListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.AppConfiguration/configurationStores"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01-preview")
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ConfigurationStoresClient) listHandleResponse(resp *http.Response) (ConfigurationStoresListResponse, error) {
	result := ConfigurationStoresListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConfigurationStoreListResult); err != nil {
		return ConfigurationStoresListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *ConfigurationStoresClient) listHandleError(resp *http.Response) error {
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

// ListByResourceGroup - Lists the configuration stores for a given resource group.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ConfigurationStoresClient) ListByResourceGroup(resourceGroupName string, options *ConfigurationStoresListByResourceGroupOptions) *ConfigurationStoresListByResourceGroupPager {
	return &ConfigurationStoresListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp ConfigurationStoresListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ConfigurationStoreListResult.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ConfigurationStoresClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ConfigurationStoresListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppConfiguration/configurationStores"
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
	reqQP.Set("api-version", "2021-03-01-preview")
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ConfigurationStoresClient) listByResourceGroupHandleResponse(resp *http.Response) (ConfigurationStoresListByResourceGroupResponse, error) {
	result := ConfigurationStoresListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConfigurationStoreListResult); err != nil {
		return ConfigurationStoresListByResourceGroupResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *ConfigurationStoresClient) listByResourceGroupHandleError(resp *http.Response) error {
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

// ListKeys - Lists the access key for the specified configuration store.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ConfigurationStoresClient) ListKeys(resourceGroupName string, configStoreName string, options *ConfigurationStoresListKeysOptions) *ConfigurationStoresListKeysPager {
	return &ConfigurationStoresListKeysPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listKeysCreateRequest(ctx, resourceGroupName, configStoreName, options)
		},
		advancer: func(ctx context.Context, resp ConfigurationStoresListKeysResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.APIKeyListResult.NextLink)
		},
	}
}

// listKeysCreateRequest creates the ListKeys request.
func (client *ConfigurationStoresClient) listKeysCreateRequest(ctx context.Context, resourceGroupName string, configStoreName string, options *ConfigurationStoresListKeysOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppConfiguration/configurationStores/{configStoreName}/listKeys"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if configStoreName == "" {
		return nil, errors.New("parameter configStoreName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configStoreName}", url.PathEscape(configStoreName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01-preview")
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listKeysHandleResponse handles the ListKeys response.
func (client *ConfigurationStoresClient) listKeysHandleResponse(resp *http.Response) (ConfigurationStoresListKeysResponse, error) {
	result := ConfigurationStoresListKeysResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.APIKeyListResult); err != nil {
		return ConfigurationStoresListKeysResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listKeysHandleError handles the ListKeys error response.
func (client *ConfigurationStoresClient) listKeysHandleError(resp *http.Response) error {
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

// RegenerateKey - Regenerates an access key for the specified configuration store.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ConfigurationStoresClient) RegenerateKey(ctx context.Context, resourceGroupName string, configStoreName string, regenerateKeyParameters RegenerateKeyParameters, options *ConfigurationStoresRegenerateKeyOptions) (ConfigurationStoresRegenerateKeyResponse, error) {
	req, err := client.regenerateKeyCreateRequest(ctx, resourceGroupName, configStoreName, regenerateKeyParameters, options)
	if err != nil {
		return ConfigurationStoresRegenerateKeyResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ConfigurationStoresRegenerateKeyResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ConfigurationStoresRegenerateKeyResponse{}, client.regenerateKeyHandleError(resp)
	}
	return client.regenerateKeyHandleResponse(resp)
}

// regenerateKeyCreateRequest creates the RegenerateKey request.
func (client *ConfigurationStoresClient) regenerateKeyCreateRequest(ctx context.Context, resourceGroupName string, configStoreName string, regenerateKeyParameters RegenerateKeyParameters, options *ConfigurationStoresRegenerateKeyOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppConfiguration/configurationStores/{configStoreName}/regenerateKey"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if configStoreName == "" {
		return nil, errors.New("parameter configStoreName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configStoreName}", url.PathEscape(configStoreName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, regenerateKeyParameters)
}

// regenerateKeyHandleResponse handles the RegenerateKey response.
func (client *ConfigurationStoresClient) regenerateKeyHandleResponse(resp *http.Response) (ConfigurationStoresRegenerateKeyResponse, error) {
	result := ConfigurationStoresRegenerateKeyResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.APIKey); err != nil {
		return ConfigurationStoresRegenerateKeyResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// regenerateKeyHandleError handles the RegenerateKey error response.
func (client *ConfigurationStoresClient) regenerateKeyHandleError(resp *http.Response) error {
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

// BeginUpdate - Updates a configuration store with the specified parameters.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ConfigurationStoresClient) BeginUpdate(ctx context.Context, resourceGroupName string, configStoreName string, configStoreUpdateParameters ConfigurationStoreUpdateParameters, options *ConfigurationStoresBeginUpdateOptions) (ConfigurationStoresUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, configStoreName, configStoreUpdateParameters, options)
	if err != nil {
		return ConfigurationStoresUpdatePollerResponse{}, err
	}
	result := ConfigurationStoresUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ConfigurationStoresClient.Update", "", resp, client.pl, client.updateHandleError)
	if err != nil {
		return ConfigurationStoresUpdatePollerResponse{}, err
	}
	result.Poller = &ConfigurationStoresUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Updates a configuration store with the specified parameters.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ConfigurationStoresClient) update(ctx context.Context, resourceGroupName string, configStoreName string, configStoreUpdateParameters ConfigurationStoreUpdateParameters, options *ConfigurationStoresBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, configStoreName, configStoreUpdateParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, client.updateHandleError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *ConfigurationStoresClient) updateCreateRequest(ctx context.Context, resourceGroupName string, configStoreName string, configStoreUpdateParameters ConfigurationStoreUpdateParameters, options *ConfigurationStoresBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppConfiguration/configurationStores/{configStoreName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if configStoreName == "" {
		return nil, errors.New("parameter configStoreName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configStoreName}", url.PathEscape(configStoreName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, configStoreUpdateParameters)
}

// updateHandleError handles the Update error response.
func (client *ConfigurationStoresClient) updateHandleError(resp *http.Response) error {
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
