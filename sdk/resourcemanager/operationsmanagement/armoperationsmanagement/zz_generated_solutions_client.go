//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armoperationsmanagement

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// SolutionsClient contains the methods for the Solutions group.
// Don't use this type directly, use NewSolutionsClient() instead.
type SolutionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewSolutionsClient creates a new instance of SolutionsClient with the specified values.
// subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
// forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewSolutionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *SolutionsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &SolutionsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// BeginCreateOrUpdate - Creates or updates the Solution.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group to get. The name is case insensitive.
// solutionName - User Solution Name.
// parameters - The parameters required to create OMS Solution.
// options - SolutionsClientBeginCreateOrUpdateOptions contains the optional parameters for the SolutionsClient.BeginCreateOrUpdate
// method.
func (client *SolutionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, solutionName string, parameters Solution, options *SolutionsClientBeginCreateOrUpdateOptions) (SolutionsClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, solutionName, parameters, options)
	if err != nil {
		return SolutionsClientCreateOrUpdatePollerResponse{}, err
	}
	result := SolutionsClientCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("SolutionsClient.CreateOrUpdate", "", resp, client.pl)
	if err != nil {
		return SolutionsClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &SolutionsClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates the Solution.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *SolutionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, solutionName string, parameters Solution, options *SolutionsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, solutionName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SolutionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, solutionName string, parameters Solution, options *SolutionsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationsManagement/solutions/{solutionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if solutionName == "" {
		return nil, errors.New("parameter solutionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{solutionName}", url.PathEscape(solutionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes the solution in the subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group to get. The name is case insensitive.
// solutionName - User Solution Name.
// options - SolutionsClientBeginDeleteOptions contains the optional parameters for the SolutionsClient.BeginDelete method.
func (client *SolutionsClient) BeginDelete(ctx context.Context, resourceGroupName string, solutionName string, options *SolutionsClientBeginDeleteOptions) (SolutionsClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, solutionName, options)
	if err != nil {
		return SolutionsClientDeletePollerResponse{}, err
	}
	result := SolutionsClientDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("SolutionsClient.Delete", "", resp, client.pl)
	if err != nil {
		return SolutionsClientDeletePollerResponse{}, err
	}
	result.Poller = &SolutionsClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the solution in the subscription.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *SolutionsClient) deleteOperation(ctx context.Context, resourceGroupName string, solutionName string, options *SolutionsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, solutionName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SolutionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, solutionName string, options *SolutionsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationsManagement/solutions/{solutionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if solutionName == "" {
		return nil, errors.New("parameter solutionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{solutionName}", url.PathEscape(solutionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Retrieves the user solution.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group to get. The name is case insensitive.
// solutionName - User Solution Name.
// options - SolutionsClientGetOptions contains the optional parameters for the SolutionsClient.Get method.
func (client *SolutionsClient) Get(ctx context.Context, resourceGroupName string, solutionName string, options *SolutionsClientGetOptions) (SolutionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, solutionName, options)
	if err != nil {
		return SolutionsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SolutionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SolutionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SolutionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, solutionName string, options *SolutionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationsManagement/solutions/{solutionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if solutionName == "" {
		return nil, errors.New("parameter solutionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{solutionName}", url.PathEscape(solutionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SolutionsClient) getHandleResponse(resp *http.Response) (SolutionsClientGetResponse, error) {
	result := SolutionsClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Solution); err != nil {
		return SolutionsClientGetResponse{}, err
	}
	return result, nil
}

// ListByResourceGroup - Retrieves the solution list. It will retrieve both first party and third party solutions
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group to get. The name is case insensitive.
// options - SolutionsClientListByResourceGroupOptions contains the optional parameters for the SolutionsClient.ListByResourceGroup
// method.
func (client *SolutionsClient) ListByResourceGroup(ctx context.Context, resourceGroupName string, options *SolutionsClientListByResourceGroupOptions) (SolutionsClientListByResourceGroupResponse, error) {
	req, err := client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
	if err != nil {
		return SolutionsClientListByResourceGroupResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SolutionsClientListByResourceGroupResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SolutionsClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
	}
	return client.listByResourceGroupHandleResponse(resp)
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *SolutionsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *SolutionsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationsManagement/solutions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *SolutionsClient) listByResourceGroupHandleResponse(resp *http.Response) (SolutionsClientListByResourceGroupResponse, error) {
	result := SolutionsClientListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SolutionPropertiesList); err != nil {
		return SolutionsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// ListBySubscription - Retrieves the solution list. It will retrieve both first party and third party solutions
// If the operation fails it returns an *azcore.ResponseError type.
// options - SolutionsClientListBySubscriptionOptions contains the optional parameters for the SolutionsClient.ListBySubscription
// method.
func (client *SolutionsClient) ListBySubscription(ctx context.Context, options *SolutionsClientListBySubscriptionOptions) (SolutionsClientListBySubscriptionResponse, error) {
	req, err := client.listBySubscriptionCreateRequest(ctx, options)
	if err != nil {
		return SolutionsClientListBySubscriptionResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SolutionsClientListBySubscriptionResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SolutionsClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
	}
	return client.listBySubscriptionHandleResponse(resp)
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *SolutionsClient) listBySubscriptionCreateRequest(ctx context.Context, options *SolutionsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.OperationsManagement/solutions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *SolutionsClient) listBySubscriptionHandleResponse(resp *http.Response) (SolutionsClientListBySubscriptionResponse, error) {
	result := SolutionsClientListBySubscriptionResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SolutionPropertiesList); err != nil {
		return SolutionsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Patch a Solution. Only updating tags supported.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group to get. The name is case insensitive.
// solutionName - User Solution Name.
// parameters - The parameters required to patch a Solution.
// options - SolutionsClientBeginUpdateOptions contains the optional parameters for the SolutionsClient.BeginUpdate method.
func (client *SolutionsClient) BeginUpdate(ctx context.Context, resourceGroupName string, solutionName string, parameters SolutionPatch, options *SolutionsClientBeginUpdateOptions) (SolutionsClientUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, solutionName, parameters, options)
	if err != nil {
		return SolutionsClientUpdatePollerResponse{}, err
	}
	result := SolutionsClientUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("SolutionsClient.Update", "", resp, client.pl)
	if err != nil {
		return SolutionsClientUpdatePollerResponse{}, err
	}
	result.Poller = &SolutionsClientUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Patch a Solution. Only updating tags supported.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *SolutionsClient) update(ctx context.Context, resourceGroupName string, solutionName string, parameters SolutionPatch, options *SolutionsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, solutionName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *SolutionsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, solutionName string, parameters SolutionPatch, options *SolutionsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationsManagement/solutions/{solutionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if solutionName == "" {
		return nil, errors.New("parameter solutionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{solutionName}", url.PathEscape(solutionName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}
