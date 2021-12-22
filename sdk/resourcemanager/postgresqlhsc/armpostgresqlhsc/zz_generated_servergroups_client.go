//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpostgresqlhsc

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

// ServerGroupsClient contains the methods for the ServerGroups group.
// Don't use this type directly, use NewServerGroupsClient() instead.
type ServerGroupsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewServerGroupsClient creates a new instance of ServerGroupsClient with the specified values.
func NewServerGroupsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ServerGroupsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &ServerGroupsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// CheckNameAvailability - Check the availability of name for resource
// If the operation fails it returns the *CloudError error type.
func (client *ServerGroupsClient) CheckNameAvailability(ctx context.Context, nameAvailabilityRequest NameAvailabilityRequest, options *ServerGroupsCheckNameAvailabilityOptions) (ServerGroupsCheckNameAvailabilityResponse, error) {
	req, err := client.checkNameAvailabilityCreateRequest(ctx, nameAvailabilityRequest, options)
	if err != nil {
		return ServerGroupsCheckNameAvailabilityResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServerGroupsCheckNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServerGroupsCheckNameAvailabilityResponse{}, client.checkNameAvailabilityHandleError(resp)
	}
	return client.checkNameAvailabilityHandleResponse(resp)
}

// checkNameAvailabilityCreateRequest creates the CheckNameAvailability request.
func (client *ServerGroupsClient) checkNameAvailabilityCreateRequest(ctx context.Context, nameAvailabilityRequest NameAvailabilityRequest, options *ServerGroupsCheckNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DBForPostgreSql/checkNameAvailability"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-05-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, nameAvailabilityRequest)
}

// checkNameAvailabilityHandleResponse handles the CheckNameAvailability response.
func (client *ServerGroupsClient) checkNameAvailabilityHandleResponse(resp *http.Response) (ServerGroupsCheckNameAvailabilityResponse, error) {
	result := ServerGroupsCheckNameAvailabilityResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.NameAvailability); err != nil {
		return ServerGroupsCheckNameAvailabilityResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// checkNameAvailabilityHandleError handles the CheckNameAvailability error response.
func (client *ServerGroupsClient) checkNameAvailabilityHandleError(resp *http.Response) error {
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

// BeginCreateOrUpdate - Creates a new server group with servers.
// If the operation fails it returns the *CloudError error type.
func (client *ServerGroupsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serverGroupName string, parameters ServerGroup, options *ServerGroupsBeginCreateOrUpdateOptions) (ServerGroupsCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, serverGroupName, parameters, options)
	if err != nil {
		return ServerGroupsCreateOrUpdatePollerResponse{}, err
	}
	result := ServerGroupsCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ServerGroupsClient.CreateOrUpdate", "", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return ServerGroupsCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &ServerGroupsCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates a new server group with servers.
// If the operation fails it returns the *CloudError error type.
func (client *ServerGroupsClient) createOrUpdate(ctx context.Context, resourceGroupName string, serverGroupName string, parameters ServerGroup, options *ServerGroupsBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serverGroupName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ServerGroupsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serverGroupName string, parameters ServerGroup, options *ServerGroupsBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBForPostgreSql/serverGroupsv2/{serverGroupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverGroupName == "" {
		return nil, errors.New("parameter serverGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverGroupName}", url.PathEscape(serverGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-05-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *ServerGroupsClient) createOrUpdateHandleError(resp *http.Response) error {
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

// BeginDelete - Deletes a server group together with servers in it.
// If the operation fails it returns the *CloudError error type.
func (client *ServerGroupsClient) BeginDelete(ctx context.Context, resourceGroupName string, serverGroupName string, options *ServerGroupsBeginDeleteOptions) (ServerGroupsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, serverGroupName, options)
	if err != nil {
		return ServerGroupsDeletePollerResponse{}, err
	}
	result := ServerGroupsDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ServerGroupsClient.Delete", "", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return ServerGroupsDeletePollerResponse{}, err
	}
	result.Poller = &ServerGroupsDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes a server group together with servers in it.
// If the operation fails it returns the *CloudError error type.
func (client *ServerGroupsClient) deleteOperation(ctx context.Context, resourceGroupName string, serverGroupName string, options *ServerGroupsBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serverGroupName, options)
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
func (client *ServerGroupsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serverGroupName string, options *ServerGroupsBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBForPostgreSql/serverGroupsv2/{serverGroupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverGroupName == "" {
		return nil, errors.New("parameter serverGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverGroupName}", url.PathEscape(serverGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-05-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *ServerGroupsClient) deleteHandleError(resp *http.Response) error {
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

// Get - Gets information about a server group.
// If the operation fails it returns the *CloudError error type.
func (client *ServerGroupsClient) Get(ctx context.Context, resourceGroupName string, serverGroupName string, options *ServerGroupsGetOptions) (ServerGroupsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverGroupName, options)
	if err != nil {
		return ServerGroupsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServerGroupsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServerGroupsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ServerGroupsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverGroupName string, options *ServerGroupsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBForPostgreSql/serverGroupsv2/{serverGroupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverGroupName == "" {
		return nil, errors.New("parameter serverGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverGroupName}", url.PathEscape(serverGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-05-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ServerGroupsClient) getHandleResponse(resp *http.Response) (ServerGroupsGetResponse, error) {
	result := ServerGroupsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerGroup); err != nil {
		return ServerGroupsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ServerGroupsClient) getHandleError(resp *http.Response) error {
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

// List - List all the server groups in a given subscription.
// If the operation fails it returns the *CloudError error type.
func (client *ServerGroupsClient) List(options *ServerGroupsListOptions) *ServerGroupsListPager {
	return &ServerGroupsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp ServerGroupsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ServerGroupListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *ServerGroupsClient) listCreateRequest(ctx context.Context, options *ServerGroupsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DBForPostgreSql/serverGroupsv2"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-05-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ServerGroupsClient) listHandleResponse(resp *http.Response) (ServerGroupsListResponse, error) {
	result := ServerGroupsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerGroupListResult); err != nil {
		return ServerGroupsListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *ServerGroupsClient) listHandleError(resp *http.Response) error {
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

// ListByResourceGroup - List all the server groups in a given resource group.
// If the operation fails it returns the *CloudError error type.
func (client *ServerGroupsClient) ListByResourceGroup(resourceGroupName string, options *ServerGroupsListByResourceGroupOptions) *ServerGroupsListByResourceGroupPager {
	return &ServerGroupsListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp ServerGroupsListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ServerGroupListResult.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ServerGroupsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ServerGroupsListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBForPostgreSql/serverGroupsv2"
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
	reqQP.Set("api-version", "2020-10-05-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ServerGroupsClient) listByResourceGroupHandleResponse(resp *http.Response) (ServerGroupsListByResourceGroupResponse, error) {
	result := ServerGroupsListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerGroupListResult); err != nil {
		return ServerGroupsListByResourceGroupResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *ServerGroupsClient) listByResourceGroupHandleError(resp *http.Response) error {
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

// BeginRestart - Restarts the server group.
// If the operation fails it returns the *CloudError error type.
func (client *ServerGroupsClient) BeginRestart(ctx context.Context, resourceGroupName string, serverGroupName string, options *ServerGroupsBeginRestartOptions) (ServerGroupsRestartPollerResponse, error) {
	resp, err := client.restart(ctx, resourceGroupName, serverGroupName, options)
	if err != nil {
		return ServerGroupsRestartPollerResponse{}, err
	}
	result := ServerGroupsRestartPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ServerGroupsClient.Restart", "", resp, client.pl, client.restartHandleError)
	if err != nil {
		return ServerGroupsRestartPollerResponse{}, err
	}
	result.Poller = &ServerGroupsRestartPoller{
		pt: pt,
	}
	return result, nil
}

// Restart - Restarts the server group.
// If the operation fails it returns the *CloudError error type.
func (client *ServerGroupsClient) restart(ctx context.Context, resourceGroupName string, serverGroupName string, options *ServerGroupsBeginRestartOptions) (*http.Response, error) {
	req, err := client.restartCreateRequest(ctx, resourceGroupName, serverGroupName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.restartHandleError(resp)
	}
	return resp, nil
}

// restartCreateRequest creates the Restart request.
func (client *ServerGroupsClient) restartCreateRequest(ctx context.Context, resourceGroupName string, serverGroupName string, options *ServerGroupsBeginRestartOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBForPostgreSql/serverGroupsv2/{serverGroupName}/restart"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverGroupName == "" {
		return nil, errors.New("parameter serverGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverGroupName}", url.PathEscape(serverGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-05-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// restartHandleError handles the Restart error response.
func (client *ServerGroupsClient) restartHandleError(resp *http.Response) error {
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

// BeginStart - Starts the server group.
// If the operation fails it returns the *CloudError error type.
func (client *ServerGroupsClient) BeginStart(ctx context.Context, resourceGroupName string, serverGroupName string, options *ServerGroupsBeginStartOptions) (ServerGroupsStartPollerResponse, error) {
	resp, err := client.start(ctx, resourceGroupName, serverGroupName, options)
	if err != nil {
		return ServerGroupsStartPollerResponse{}, err
	}
	result := ServerGroupsStartPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ServerGroupsClient.Start", "", resp, client.pl, client.startHandleError)
	if err != nil {
		return ServerGroupsStartPollerResponse{}, err
	}
	result.Poller = &ServerGroupsStartPoller{
		pt: pt,
	}
	return result, nil
}

// Start - Starts the server group.
// If the operation fails it returns the *CloudError error type.
func (client *ServerGroupsClient) start(ctx context.Context, resourceGroupName string, serverGroupName string, options *ServerGroupsBeginStartOptions) (*http.Response, error) {
	req, err := client.startCreateRequest(ctx, resourceGroupName, serverGroupName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.startHandleError(resp)
	}
	return resp, nil
}

// startCreateRequest creates the Start request.
func (client *ServerGroupsClient) startCreateRequest(ctx context.Context, resourceGroupName string, serverGroupName string, options *ServerGroupsBeginStartOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBForPostgreSql/serverGroupsv2/{serverGroupName}/start"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverGroupName == "" {
		return nil, errors.New("parameter serverGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverGroupName}", url.PathEscape(serverGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-05-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// startHandleError handles the Start error response.
func (client *ServerGroupsClient) startHandleError(resp *http.Response) error {
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

// BeginStop - Stops the server group.
// If the operation fails it returns the *CloudError error type.
func (client *ServerGroupsClient) BeginStop(ctx context.Context, resourceGroupName string, serverGroupName string, options *ServerGroupsBeginStopOptions) (ServerGroupsStopPollerResponse, error) {
	resp, err := client.stop(ctx, resourceGroupName, serverGroupName, options)
	if err != nil {
		return ServerGroupsStopPollerResponse{}, err
	}
	result := ServerGroupsStopPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ServerGroupsClient.Stop", "", resp, client.pl, client.stopHandleError)
	if err != nil {
		return ServerGroupsStopPollerResponse{}, err
	}
	result.Poller = &ServerGroupsStopPoller{
		pt: pt,
	}
	return result, nil
}

// Stop - Stops the server group.
// If the operation fails it returns the *CloudError error type.
func (client *ServerGroupsClient) stop(ctx context.Context, resourceGroupName string, serverGroupName string, options *ServerGroupsBeginStopOptions) (*http.Response, error) {
	req, err := client.stopCreateRequest(ctx, resourceGroupName, serverGroupName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.stopHandleError(resp)
	}
	return resp, nil
}

// stopCreateRequest creates the Stop request.
func (client *ServerGroupsClient) stopCreateRequest(ctx context.Context, resourceGroupName string, serverGroupName string, options *ServerGroupsBeginStopOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBForPostgreSql/serverGroupsv2/{serverGroupName}/stop"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverGroupName == "" {
		return nil, errors.New("parameter serverGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverGroupName}", url.PathEscape(serverGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-05-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// stopHandleError handles the Stop error response.
func (client *ServerGroupsClient) stopHandleError(resp *http.Response) error {
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

// BeginUpdate - Updates an existing server group. The request body can contain one to many of the properties present in the normal server group definition.
// If the operation fails it returns the *CloudError error type.
func (client *ServerGroupsClient) BeginUpdate(ctx context.Context, resourceGroupName string, serverGroupName string, parameters ServerGroupForUpdate, options *ServerGroupsBeginUpdateOptions) (ServerGroupsUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, serverGroupName, parameters, options)
	if err != nil {
		return ServerGroupsUpdatePollerResponse{}, err
	}
	result := ServerGroupsUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ServerGroupsClient.Update", "", resp, client.pl, client.updateHandleError)
	if err != nil {
		return ServerGroupsUpdatePollerResponse{}, err
	}
	result.Poller = &ServerGroupsUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Updates an existing server group. The request body can contain one to many of the properties present in the normal server group definition.
// If the operation fails it returns the *CloudError error type.
func (client *ServerGroupsClient) update(ctx context.Context, resourceGroupName string, serverGroupName string, parameters ServerGroupForUpdate, options *ServerGroupsBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, serverGroupName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.updateHandleError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *ServerGroupsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, serverGroupName string, parameters ServerGroupForUpdate, options *ServerGroupsBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBForPostgreSql/serverGroupsv2/{serverGroupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverGroupName == "" {
		return nil, errors.New("parameter serverGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverGroupName}", url.PathEscape(serverGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-05-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleError handles the Update error response.
func (client *ServerGroupsClient) updateHandleError(resp *http.Response) error {
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
