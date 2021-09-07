//go:build go1.16
// +build go1.16

// Code generated by Microsoft (R) AutoRest Code Generator (autorest: 3.4.3, generator: @autorest/go@4.0.0-preview.27)
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmysql

import (
	"context"
	"errors"
	"fmt"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ServerParametersClient contains the methods for the ServerParameters group.
// Don't use this type directly, use NewServerParametersClient() instead.
type ServerParametersClient struct {
	con            *connection
	subscriptionID string
}

// NewServerParametersClient creates a new instance of ServerParametersClient with the specified values.
func NewServerParametersClient(con *connection, subscriptionID string) *ServerParametersClient {
	return &ServerParametersClient{con: con, subscriptionID: subscriptionID}
}

// BeginListUpdateConfigurations - Update a list of configurations in a given server.
// If the operation fails it returns the *CloudError error type.
func (client *ServerParametersClient) BeginListUpdateConfigurations(ctx context.Context, resourceGroupName string, serverName string, value ConfigurationListResult, options *ServerParametersBeginListUpdateConfigurationsOptions) (ServerParametersListUpdateConfigurationsPollerResponse, error) {
	resp, err := client.listUpdateConfigurations(ctx, resourceGroupName, serverName, value, options)
	if err != nil {
		return ServerParametersListUpdateConfigurationsPollerResponse{}, err
	}
	result := ServerParametersListUpdateConfigurationsPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ServerParametersClient.ListUpdateConfigurations", "azure-async-operation", resp, client.con.Pipeline(), client.listUpdateConfigurationsHandleError)
	if err != nil {
		return ServerParametersListUpdateConfigurationsPollerResponse{}, err
	}
	result.Poller = &ServerParametersListUpdateConfigurationsPoller{
		pt: pt,
	}
	return result, nil
}

// ListUpdateConfigurations - Update a list of configurations in a given server.
// If the operation fails it returns the *CloudError error type.
func (client *ServerParametersClient) listUpdateConfigurations(ctx context.Context, resourceGroupName string, serverName string, value ConfigurationListResult, options *ServerParametersBeginListUpdateConfigurationsOptions) (*http.Response, error) {
	req, err := client.listUpdateConfigurationsCreateRequest(ctx, resourceGroupName, serverName, value, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.listUpdateConfigurationsHandleError(resp)
	}
	return resp, nil
}

// listUpdateConfigurationsCreateRequest creates the ListUpdateConfigurations request.
func (client *ServerParametersClient) listUpdateConfigurationsCreateRequest(ctx context.Context, resourceGroupName string, serverName string, value ConfigurationListResult, options *ServerParametersBeginListUpdateConfigurationsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/updateConfigurations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, value)
}

// listUpdateConfigurationsHandleError handles the ListUpdateConfigurations error response.
func (client *ServerParametersClient) listUpdateConfigurationsHandleError(resp *http.Response) error {
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
