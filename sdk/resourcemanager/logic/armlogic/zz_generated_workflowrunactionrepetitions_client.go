//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armlogic

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// WorkflowRunActionRepetitionsClient contains the methods for the WorkflowRunActionRepetitions group.
// Don't use this type directly, use NewWorkflowRunActionRepetitionsClient() instead.
type WorkflowRunActionRepetitionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewWorkflowRunActionRepetitionsClient creates a new instance of WorkflowRunActionRepetitionsClient with the specified values.
// subscriptionID - The subscription id.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewWorkflowRunActionRepetitionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*WorkflowRunActionRepetitionsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &WorkflowRunActionRepetitionsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Get a workflow run action repetition.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-05-01
// resourceGroupName - The resource group name.
// workflowName - The workflow name.
// runName - The workflow run name.
// actionName - The workflow action name.
// repetitionName - The workflow repetition.
// options - WorkflowRunActionRepetitionsClientGetOptions contains the optional parameters for the WorkflowRunActionRepetitionsClient.Get
// method.
func (client *WorkflowRunActionRepetitionsClient) Get(ctx context.Context, resourceGroupName string, workflowName string, runName string, actionName string, repetitionName string, options *WorkflowRunActionRepetitionsClientGetOptions) (WorkflowRunActionRepetitionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workflowName, runName, actionName, repetitionName, options)
	if err != nil {
		return WorkflowRunActionRepetitionsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkflowRunActionRepetitionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WorkflowRunActionRepetitionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *WorkflowRunActionRepetitionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, workflowName string, runName string, actionName string, repetitionName string, options *WorkflowRunActionRepetitionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs/{runName}/actions/{actionName}/repetitions/{repetitionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workflowName == "" {
		return nil, errors.New("parameter workflowName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workflowName}", url.PathEscape(workflowName))
	if runName == "" {
		return nil, errors.New("parameter runName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runName}", url.PathEscape(runName))
	if actionName == "" {
		return nil, errors.New("parameter actionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{actionName}", url.PathEscape(actionName))
	if repetitionName == "" {
		return nil, errors.New("parameter repetitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{repetitionName}", url.PathEscape(repetitionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *WorkflowRunActionRepetitionsClient) getHandleResponse(resp *http.Response) (WorkflowRunActionRepetitionsClientGetResponse, error) {
	result := WorkflowRunActionRepetitionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkflowRunActionRepetitionDefinition); err != nil {
		return WorkflowRunActionRepetitionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Get all of a workflow run action repetitions.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-05-01
// resourceGroupName - The resource group name.
// workflowName - The workflow name.
// runName - The workflow run name.
// actionName - The workflow action name.
// options - WorkflowRunActionRepetitionsClientListOptions contains the optional parameters for the WorkflowRunActionRepetitionsClient.List
// method.
func (client *WorkflowRunActionRepetitionsClient) NewListPager(resourceGroupName string, workflowName string, runName string, actionName string, options *WorkflowRunActionRepetitionsClientListOptions) *runtime.Pager[WorkflowRunActionRepetitionsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[WorkflowRunActionRepetitionsClientListResponse]{
		More: func(page WorkflowRunActionRepetitionsClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *WorkflowRunActionRepetitionsClientListResponse) (WorkflowRunActionRepetitionsClientListResponse, error) {
			req, err := client.listCreateRequest(ctx, resourceGroupName, workflowName, runName, actionName, options)
			if err != nil {
				return WorkflowRunActionRepetitionsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return WorkflowRunActionRepetitionsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return WorkflowRunActionRepetitionsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *WorkflowRunActionRepetitionsClient) listCreateRequest(ctx context.Context, resourceGroupName string, workflowName string, runName string, actionName string, options *WorkflowRunActionRepetitionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs/{runName}/actions/{actionName}/repetitions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workflowName == "" {
		return nil, errors.New("parameter workflowName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workflowName}", url.PathEscape(workflowName))
	if runName == "" {
		return nil, errors.New("parameter runName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runName}", url.PathEscape(runName))
	if actionName == "" {
		return nil, errors.New("parameter actionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{actionName}", url.PathEscape(actionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *WorkflowRunActionRepetitionsClient) listHandleResponse(resp *http.Response) (WorkflowRunActionRepetitionsClientListResponse, error) {
	result := WorkflowRunActionRepetitionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkflowRunActionRepetitionDefinitionCollection); err != nil {
		return WorkflowRunActionRepetitionsClientListResponse{}, err
	}
	return result, nil
}

// NewListExpressionTracesPager - Lists a workflow run expression trace.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-05-01
// resourceGroupName - The resource group name.
// workflowName - The workflow name.
// runName - The workflow run name.
// actionName - The workflow action name.
// repetitionName - The workflow repetition.
// options - WorkflowRunActionRepetitionsClientListExpressionTracesOptions contains the optional parameters for the WorkflowRunActionRepetitionsClient.ListExpressionTraces
// method.
func (client *WorkflowRunActionRepetitionsClient) NewListExpressionTracesPager(resourceGroupName string, workflowName string, runName string, actionName string, repetitionName string, options *WorkflowRunActionRepetitionsClientListExpressionTracesOptions) *runtime.Pager[WorkflowRunActionRepetitionsClientListExpressionTracesResponse] {
	return runtime.NewPager(runtime.PagingHandler[WorkflowRunActionRepetitionsClientListExpressionTracesResponse]{
		More: func(page WorkflowRunActionRepetitionsClientListExpressionTracesResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *WorkflowRunActionRepetitionsClientListExpressionTracesResponse) (WorkflowRunActionRepetitionsClientListExpressionTracesResponse, error) {
			req, err := client.listExpressionTracesCreateRequest(ctx, resourceGroupName, workflowName, runName, actionName, repetitionName, options)
			if err != nil {
				return WorkflowRunActionRepetitionsClientListExpressionTracesResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return WorkflowRunActionRepetitionsClientListExpressionTracesResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return WorkflowRunActionRepetitionsClientListExpressionTracesResponse{}, runtime.NewResponseError(resp)
			}
			return client.listExpressionTracesHandleResponse(resp)
		},
	})
}

// listExpressionTracesCreateRequest creates the ListExpressionTraces request.
func (client *WorkflowRunActionRepetitionsClient) listExpressionTracesCreateRequest(ctx context.Context, resourceGroupName string, workflowName string, runName string, actionName string, repetitionName string, options *WorkflowRunActionRepetitionsClientListExpressionTracesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs/{runName}/actions/{actionName}/repetitions/{repetitionName}/listExpressionTraces"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workflowName == "" {
		return nil, errors.New("parameter workflowName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workflowName}", url.PathEscape(workflowName))
	if runName == "" {
		return nil, errors.New("parameter runName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runName}", url.PathEscape(runName))
	if actionName == "" {
		return nil, errors.New("parameter actionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{actionName}", url.PathEscape(actionName))
	if repetitionName == "" {
		return nil, errors.New("parameter repetitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{repetitionName}", url.PathEscape(repetitionName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listExpressionTracesHandleResponse handles the ListExpressionTraces response.
func (client *WorkflowRunActionRepetitionsClient) listExpressionTracesHandleResponse(resp *http.Response) (WorkflowRunActionRepetitionsClientListExpressionTracesResponse, error) {
	result := WorkflowRunActionRepetitionsClientListExpressionTracesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExpressionTraces); err != nil {
		return WorkflowRunActionRepetitionsClientListExpressionTracesResponse{}, err
	}
	return result, nil
}
