//go:build go1.16
// +build go1.16

// Code generated by Microsoft (R) AutoRest Code Generator (autorest: 3.4.3, generator: @autorest/go@4.0.0-preview.27)
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmysql

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// LocationBasedRecommendedActionSessionsOperationStatusClient contains the methods for the LocationBasedRecommendedActionSessionsOperationStatus group.
// Don't use this type directly, use NewLocationBasedRecommendedActionSessionsOperationStatusClient() instead.
type LocationBasedRecommendedActionSessionsOperationStatusClient struct {
	con            *connection
	subscriptionID string
}

// NewLocationBasedRecommendedActionSessionsOperationStatusClient creates a new instance of LocationBasedRecommendedActionSessionsOperationStatusClient with the specified values.
func NewLocationBasedRecommendedActionSessionsOperationStatusClient(con *connection, subscriptionID string) *LocationBasedRecommendedActionSessionsOperationStatusClient {
	return &LocationBasedRecommendedActionSessionsOperationStatusClient{con: con, subscriptionID: subscriptionID}
}

// Get - Recommendation action session operation status.
// If the operation fails it returns a generic error.
func (client *LocationBasedRecommendedActionSessionsOperationStatusClient) Get(ctx context.Context, locationName string, operationID string, options *LocationBasedRecommendedActionSessionsOperationStatusGetOptions) (LocationBasedRecommendedActionSessionsOperationStatusGetResponse, error) {
	req, err := client.getCreateRequest(ctx, locationName, operationID, options)
	if err != nil {
		return LocationBasedRecommendedActionSessionsOperationStatusGetResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return LocationBasedRecommendedActionSessionsOperationStatusGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LocationBasedRecommendedActionSessionsOperationStatusGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *LocationBasedRecommendedActionSessionsOperationStatusClient) getCreateRequest(ctx context.Context, locationName string, operationID string, options *LocationBasedRecommendedActionSessionsOperationStatusGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DBforMySQL/locations/{locationName}/recommendedActionSessionsAzureAsyncOperation/{operationId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *LocationBasedRecommendedActionSessionsOperationStatusClient) getHandleResponse(resp *http.Response) (LocationBasedRecommendedActionSessionsOperationStatusGetResponse, error) {
	result := LocationBasedRecommendedActionSessionsOperationStatusGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RecommendedActionSessionsOperationStatus); err != nil {
		return LocationBasedRecommendedActionSessionsOperationStatusGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *LocationBasedRecommendedActionSessionsOperationStatusClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
