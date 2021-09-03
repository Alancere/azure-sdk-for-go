//go:build go1.16
// +build go1.16

// Code generated by Microsoft (R) AutoRest Code Generator (autorest: 3.4.3, generator: @autorest/go@4.0.0-preview.27)
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpostgresql

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// VirtualNetworkSubnetUsageClient contains the methods for the VirtualNetworkSubnetUsage group.
// Don't use this type directly, use NewVirtualNetworkSubnetUsageClient() instead.
type VirtualNetworkSubnetUsageClient struct {
	con            *connection
	subscriptionID string
}

// NewVirtualNetworkSubnetUsageClient creates a new instance of VirtualNetworkSubnetUsageClient with the specified values.
func NewVirtualNetworkSubnetUsageClient(con *connection, subscriptionID string) *VirtualNetworkSubnetUsageClient {
	return &VirtualNetworkSubnetUsageClient{con: con, subscriptionID: subscriptionID}
}

// Execute - Get virtual network subnet usage for a given vNet resource id.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualNetworkSubnetUsageClient) Execute(ctx context.Context, locationName string, parameters VirtualNetworkSubnetUsageParameter, options *VirtualNetworkSubnetUsageExecuteOptions) (VirtualNetworkSubnetUsageExecuteResponse, error) {
	req, err := client.executeCreateRequest(ctx, locationName, parameters, options)
	if err != nil {
		return VirtualNetworkSubnetUsageExecuteResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return VirtualNetworkSubnetUsageExecuteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualNetworkSubnetUsageExecuteResponse{}, client.executeHandleError(resp)
	}
	return client.executeHandleResponse(resp)
}

// executeCreateRequest creates the Execute request.
func (client *VirtualNetworkSubnetUsageClient) executeCreateRequest(ctx context.Context, locationName string, parameters VirtualNetworkSubnetUsageParameter, options *VirtualNetworkSubnetUsageExecuteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DBforPostgreSQL/locations/{locationName}/checkVirtualNetworkSubnetUsage"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// executeHandleResponse handles the Execute response.
func (client *VirtualNetworkSubnetUsageClient) executeHandleResponse(resp *http.Response) (VirtualNetworkSubnetUsageExecuteResponse, error) {
	result := VirtualNetworkSubnetUsageExecuteResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualNetworkSubnetUsageResult); err != nil {
		return VirtualNetworkSubnetUsageExecuteResponse{}, err
	}
	return result, nil
}

// executeHandleError handles the Execute error response.
func (client *VirtualNetworkSubnetUsageClient) executeHandleError(resp *http.Response) error {
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
