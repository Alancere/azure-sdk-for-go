//go:build go1.16
// +build go1.16

// Code generated by Microsoft (R) AutoRest Code Generator (autorest: 3.4.3, generator: @autorest/go@4.0.0-preview.27)
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmysql

import (
	"context"
	"errors"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// VirtualNetworkRulesClient contains the methods for the VirtualNetworkRules group.
// Don't use this type directly, use NewVirtualNetworkRulesClient() instead.
type VirtualNetworkRulesClient struct {
	con            *connection
	subscriptionID string
}

// NewVirtualNetworkRulesClient creates a new instance of VirtualNetworkRulesClient with the specified values.
func NewVirtualNetworkRulesClient(con *connection, subscriptionID string) *VirtualNetworkRulesClient {
	return &VirtualNetworkRulesClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates or updates an existing virtual network rule.
// If the operation fails it returns a generic error.
func (client *VirtualNetworkRulesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, virtualNetworkRuleName string, parameters VirtualNetworkRule, options *VirtualNetworkRulesBeginCreateOrUpdateOptions) (VirtualNetworkRulesCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, serverName, virtualNetworkRuleName, parameters, options)
	if err != nil {
		return VirtualNetworkRulesCreateOrUpdatePollerResponse{}, err
	}
	result := VirtualNetworkRulesCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("VirtualNetworkRulesClient.CreateOrUpdate", "", resp, client.con.Pipeline(), client.createOrUpdateHandleError)
	if err != nil {
		return VirtualNetworkRulesCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &VirtualNetworkRulesCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates an existing virtual network rule.
// If the operation fails it returns a generic error.
func (client *VirtualNetworkRulesClient) createOrUpdate(ctx context.Context, resourceGroupName string, serverName string, virtualNetworkRuleName string, parameters VirtualNetworkRule, options *VirtualNetworkRulesBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serverName, virtualNetworkRuleName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *VirtualNetworkRulesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, virtualNetworkRuleName string, parameters VirtualNetworkRule, options *VirtualNetworkRulesBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/virtualNetworkRules/{virtualNetworkRuleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if virtualNetworkRuleName == "" {
		return nil, errors.New("parameter virtualNetworkRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkRuleName}", url.PathEscape(virtualNetworkRuleName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *VirtualNetworkRulesClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// BeginDelete - Deletes the virtual network rule with the given name.
// If the operation fails it returns a generic error.
func (client *VirtualNetworkRulesClient) BeginDelete(ctx context.Context, resourceGroupName string, serverName string, virtualNetworkRuleName string, options *VirtualNetworkRulesBeginDeleteOptions) (VirtualNetworkRulesDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, serverName, virtualNetworkRuleName, options)
	if err != nil {
		return VirtualNetworkRulesDeletePollerResponse{}, err
	}
	result := VirtualNetworkRulesDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("VirtualNetworkRulesClient.Delete", "", resp, client.con.Pipeline(), client.deleteHandleError)
	if err != nil {
		return VirtualNetworkRulesDeletePollerResponse{}, err
	}
	result.Poller = &VirtualNetworkRulesDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the virtual network rule with the given name.
// If the operation fails it returns a generic error.
func (client *VirtualNetworkRulesClient) deleteOperation(ctx context.Context, resourceGroupName string, serverName string, virtualNetworkRuleName string, options *VirtualNetworkRulesBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serverName, virtualNetworkRuleName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *VirtualNetworkRulesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serverName string, virtualNetworkRuleName string, options *VirtualNetworkRulesBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/virtualNetworkRules/{virtualNetworkRuleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if virtualNetworkRuleName == "" {
		return nil, errors.New("parameter virtualNetworkRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkRuleName}", url.PathEscape(virtualNetworkRuleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *VirtualNetworkRulesClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Get - Gets a virtual network rule.
// If the operation fails it returns a generic error.
func (client *VirtualNetworkRulesClient) Get(ctx context.Context, resourceGroupName string, serverName string, virtualNetworkRuleName string, options *VirtualNetworkRulesGetOptions) (VirtualNetworkRulesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, virtualNetworkRuleName, options)
	if err != nil {
		return VirtualNetworkRulesGetResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return VirtualNetworkRulesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualNetworkRulesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VirtualNetworkRulesClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, virtualNetworkRuleName string, options *VirtualNetworkRulesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/virtualNetworkRules/{virtualNetworkRuleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if virtualNetworkRuleName == "" {
		return nil, errors.New("parameter virtualNetworkRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkRuleName}", url.PathEscape(virtualNetworkRuleName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VirtualNetworkRulesClient) getHandleResponse(resp *http.Response) (VirtualNetworkRulesGetResponse, error) {
	result := VirtualNetworkRulesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualNetworkRule); err != nil {
		return VirtualNetworkRulesGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *VirtualNetworkRulesClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListByServer - Gets a list of virtual network rules in a server.
// If the operation fails it returns a generic error.
func (client *VirtualNetworkRulesClient) ListByServer(resourceGroupName string, serverName string, options *VirtualNetworkRulesListByServerOptions) *VirtualNetworkRulesListByServerPager {
	return &VirtualNetworkRulesListByServerPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByServerCreateRequest(ctx, resourceGroupName, serverName, options)
		},
		advancer: func(ctx context.Context, resp VirtualNetworkRulesListByServerResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.VirtualNetworkRuleListResult.NextLink)
		},
	}
}

// listByServerCreateRequest creates the ListByServer request.
func (client *VirtualNetworkRulesClient) listByServerCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *VirtualNetworkRulesListByServerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/virtualNetworkRules"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByServerHandleResponse handles the ListByServer response.
func (client *VirtualNetworkRulesClient) listByServerHandleResponse(resp *http.Response) (VirtualNetworkRulesListByServerResponse, error) {
	result := VirtualNetworkRulesListByServerResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualNetworkRuleListResult); err != nil {
		return VirtualNetworkRulesListByServerResponse{}, err
	}
	return result, nil
}

// listByServerHandleError handles the ListByServer error response.
func (client *VirtualNetworkRulesClient) listByServerHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
