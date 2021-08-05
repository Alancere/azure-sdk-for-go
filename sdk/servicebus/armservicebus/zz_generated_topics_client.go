// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armservicebus

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// TopicsClient contains the methods for the Topics group.
// Don't use this type directly, use NewTopicsClient() instead.
type TopicsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewTopicsClient creates a new instance of TopicsClient with the specified values.
func NewTopicsClient(con *armcore.Connection, subscriptionID string) *TopicsClient {
	return &TopicsClient{con: con, subscriptionID: subscriptionID}
}

// CreateOrUpdate - Creates a topic in the specified namespace.
// If the operation fails it returns the *ErrorResponse error type.
func (client *TopicsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, parameters SBTopic, options *TopicsCreateOrUpdateOptions) (TopicsCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, namespaceName, topicName, parameters, options)
	if err != nil {
		return TopicsCreateOrUpdateResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return TopicsCreateOrUpdateResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return TopicsCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *TopicsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, parameters SBTopic, options *TopicsCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if topicName == "" {
		return nil, errors.New("parameter topicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{topicName}", url.PathEscape(topicName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2018-01-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *TopicsClient) createOrUpdateHandleResponse(resp *azcore.Response) (TopicsCreateOrUpdateResponse, error) {
	result := TopicsCreateOrUpdateResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.SBTopic); err != nil {
		return TopicsCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *TopicsClient) createOrUpdateHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// CreateOrUpdateAuthorizationRule - Creates an authorization rule for the specified topic.
// If the operation fails it returns the *ErrorResponse error type.
func (client *TopicsClient) CreateOrUpdateAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, authorizationRuleName string, parameters SBAuthorizationRule, options *TopicsCreateOrUpdateAuthorizationRuleOptions) (TopicsCreateOrUpdateAuthorizationRuleResponse, error) {
	req, err := client.createOrUpdateAuthorizationRuleCreateRequest(ctx, resourceGroupName, namespaceName, topicName, authorizationRuleName, parameters, options)
	if err != nil {
		return TopicsCreateOrUpdateAuthorizationRuleResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return TopicsCreateOrUpdateAuthorizationRuleResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return TopicsCreateOrUpdateAuthorizationRuleResponse{}, client.createOrUpdateAuthorizationRuleHandleError(resp)
	}
	return client.createOrUpdateAuthorizationRuleHandleResponse(resp)
}

// createOrUpdateAuthorizationRuleCreateRequest creates the CreateOrUpdateAuthorizationRule request.
func (client *TopicsClient) createOrUpdateAuthorizationRuleCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, authorizationRuleName string, parameters SBAuthorizationRule, options *TopicsCreateOrUpdateAuthorizationRuleOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}/authorizationRules/{authorizationRuleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if topicName == "" {
		return nil, errors.New("parameter topicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{topicName}", url.PathEscape(topicName))
	if authorizationRuleName == "" {
		return nil, errors.New("parameter authorizationRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorizationRuleName}", url.PathEscape(authorizationRuleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2018-01-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// createOrUpdateAuthorizationRuleHandleResponse handles the CreateOrUpdateAuthorizationRule response.
func (client *TopicsClient) createOrUpdateAuthorizationRuleHandleResponse(resp *azcore.Response) (TopicsCreateOrUpdateAuthorizationRuleResponse, error) {
	result := TopicsCreateOrUpdateAuthorizationRuleResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.SBAuthorizationRule); err != nil {
		return TopicsCreateOrUpdateAuthorizationRuleResponse{}, err
	}
	return result, nil
}

// createOrUpdateAuthorizationRuleHandleError handles the CreateOrUpdateAuthorizationRule error response.
func (client *TopicsClient) createOrUpdateAuthorizationRuleHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Delete - Deletes a topic from the specified namespace and resource group.
// If the operation fails it returns the *ErrorResponse error type.
func (client *TopicsClient) Delete(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, options *TopicsDeleteOptions) (TopicsDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, namespaceName, topicName, options)
	if err != nil {
		return TopicsDeleteResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return TopicsDeleteResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusNoContent) {
		return TopicsDeleteResponse{}, client.deleteHandleError(resp)
	}
	return TopicsDeleteResponse{RawResponse: resp.Response}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *TopicsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, options *TopicsDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if topicName == "" {
		return nil, errors.New("parameter topicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{topicName}", url.PathEscape(topicName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2018-01-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *TopicsClient) deleteHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// DeleteAuthorizationRule - Deletes a topic authorization rule.
// If the operation fails it returns the *ErrorResponse error type.
func (client *TopicsClient) DeleteAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, authorizationRuleName string, options *TopicsDeleteAuthorizationRuleOptions) (TopicsDeleteAuthorizationRuleResponse, error) {
	req, err := client.deleteAuthorizationRuleCreateRequest(ctx, resourceGroupName, namespaceName, topicName, authorizationRuleName, options)
	if err != nil {
		return TopicsDeleteAuthorizationRuleResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return TopicsDeleteAuthorizationRuleResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusNoContent) {
		return TopicsDeleteAuthorizationRuleResponse{}, client.deleteAuthorizationRuleHandleError(resp)
	}
	return TopicsDeleteAuthorizationRuleResponse{RawResponse: resp.Response}, nil
}

// deleteAuthorizationRuleCreateRequest creates the DeleteAuthorizationRule request.
func (client *TopicsClient) deleteAuthorizationRuleCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, authorizationRuleName string, options *TopicsDeleteAuthorizationRuleOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}/authorizationRules/{authorizationRuleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if topicName == "" {
		return nil, errors.New("parameter topicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{topicName}", url.PathEscape(topicName))
	if authorizationRuleName == "" {
		return nil, errors.New("parameter authorizationRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorizationRuleName}", url.PathEscape(authorizationRuleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2018-01-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteAuthorizationRuleHandleError handles the DeleteAuthorizationRule error response.
func (client *TopicsClient) deleteAuthorizationRuleHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Get - Returns a description for the specified topic.
// If the operation fails it returns the *ErrorResponse error type.
func (client *TopicsClient) Get(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, options *TopicsGetOptions) (TopicsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, namespaceName, topicName, options)
	if err != nil {
		return TopicsGetResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return TopicsGetResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return TopicsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *TopicsClient) getCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, options *TopicsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if topicName == "" {
		return nil, errors.New("parameter topicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{topicName}", url.PathEscape(topicName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2018-01-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *TopicsClient) getHandleResponse(resp *azcore.Response) (TopicsGetResponse, error) {
	result := TopicsGetResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.SBTopic); err != nil {
		return TopicsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *TopicsClient) getHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// GetAuthorizationRule - Returns the specified authorization rule.
// If the operation fails it returns the *ErrorResponse error type.
func (client *TopicsClient) GetAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, authorizationRuleName string, options *TopicsGetAuthorizationRuleOptions) (TopicsGetAuthorizationRuleResponse, error) {
	req, err := client.getAuthorizationRuleCreateRequest(ctx, resourceGroupName, namespaceName, topicName, authorizationRuleName, options)
	if err != nil {
		return TopicsGetAuthorizationRuleResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return TopicsGetAuthorizationRuleResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return TopicsGetAuthorizationRuleResponse{}, client.getAuthorizationRuleHandleError(resp)
	}
	return client.getAuthorizationRuleHandleResponse(resp)
}

// getAuthorizationRuleCreateRequest creates the GetAuthorizationRule request.
func (client *TopicsClient) getAuthorizationRuleCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, authorizationRuleName string, options *TopicsGetAuthorizationRuleOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}/authorizationRules/{authorizationRuleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if topicName == "" {
		return nil, errors.New("parameter topicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{topicName}", url.PathEscape(topicName))
	if authorizationRuleName == "" {
		return nil, errors.New("parameter authorizationRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorizationRuleName}", url.PathEscape(authorizationRuleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2018-01-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getAuthorizationRuleHandleResponse handles the GetAuthorizationRule response.
func (client *TopicsClient) getAuthorizationRuleHandleResponse(resp *azcore.Response) (TopicsGetAuthorizationRuleResponse, error) {
	result := TopicsGetAuthorizationRuleResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.SBAuthorizationRule); err != nil {
		return TopicsGetAuthorizationRuleResponse{}, err
	}
	return result, nil
}

// getAuthorizationRuleHandleError handles the GetAuthorizationRule error response.
func (client *TopicsClient) getAuthorizationRuleHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// ListAuthorizationRules - Gets authorization rules for a topic.
// If the operation fails it returns the *ErrorResponse error type.
func (client *TopicsClient) ListAuthorizationRules(resourceGroupName string, namespaceName string, topicName string, options *TopicsListAuthorizationRulesOptions) TopicsListAuthorizationRulesPager {
	return &topicsListAuthorizationRulesPager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listAuthorizationRulesCreateRequest(ctx, resourceGroupName, namespaceName, topicName, options)
		},
		advancer: func(ctx context.Context, resp TopicsListAuthorizationRulesResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.SBAuthorizationRuleListResult.NextLink)
		},
	}
}

// listAuthorizationRulesCreateRequest creates the ListAuthorizationRules request.
func (client *TopicsClient) listAuthorizationRulesCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, options *TopicsListAuthorizationRulesOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}/authorizationRules"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if topicName == "" {
		return nil, errors.New("parameter topicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{topicName}", url.PathEscape(topicName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2018-01-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listAuthorizationRulesHandleResponse handles the ListAuthorizationRules response.
func (client *TopicsClient) listAuthorizationRulesHandleResponse(resp *azcore.Response) (TopicsListAuthorizationRulesResponse, error) {
	result := TopicsListAuthorizationRulesResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.SBAuthorizationRuleListResult); err != nil {
		return TopicsListAuthorizationRulesResponse{}, err
	}
	return result, nil
}

// listAuthorizationRulesHandleError handles the ListAuthorizationRules error response.
func (client *TopicsClient) listAuthorizationRulesHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// ListByNamespace - Gets all the topics in a namespace.
// If the operation fails it returns the *ErrorResponse error type.
func (client *TopicsClient) ListByNamespace(resourceGroupName string, namespaceName string, options *TopicsListByNamespaceOptions) TopicsListByNamespacePager {
	return &topicsListByNamespacePager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByNamespaceCreateRequest(ctx, resourceGroupName, namespaceName, options)
		},
		advancer: func(ctx context.Context, resp TopicsListByNamespaceResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.SBTopicListResult.NextLink)
		},
	}
}

// listByNamespaceCreateRequest creates the ListByNamespace request.
func (client *TopicsClient) listByNamespaceCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, options *TopicsListByNamespaceOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2018-01-01-preview")
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listByNamespaceHandleResponse handles the ListByNamespace response.
func (client *TopicsClient) listByNamespaceHandleResponse(resp *azcore.Response) (TopicsListByNamespaceResponse, error) {
	result := TopicsListByNamespaceResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.SBTopicListResult); err != nil {
		return TopicsListByNamespaceResponse{}, err
	}
	return result, nil
}

// listByNamespaceHandleError handles the ListByNamespace error response.
func (client *TopicsClient) listByNamespaceHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// ListKeys - Gets the primary and secondary connection strings for the topic.
// If the operation fails it returns the *ErrorResponse error type.
func (client *TopicsClient) ListKeys(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, authorizationRuleName string, options *TopicsListKeysOptions) (TopicsListKeysResponse, error) {
	req, err := client.listKeysCreateRequest(ctx, resourceGroupName, namespaceName, topicName, authorizationRuleName, options)
	if err != nil {
		return TopicsListKeysResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return TopicsListKeysResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return TopicsListKeysResponse{}, client.listKeysHandleError(resp)
	}
	return client.listKeysHandleResponse(resp)
}

// listKeysCreateRequest creates the ListKeys request.
func (client *TopicsClient) listKeysCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, authorizationRuleName string, options *TopicsListKeysOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}/authorizationRules/{authorizationRuleName}/ListKeys"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if topicName == "" {
		return nil, errors.New("parameter topicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{topicName}", url.PathEscape(topicName))
	if authorizationRuleName == "" {
		return nil, errors.New("parameter authorizationRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorizationRuleName}", url.PathEscape(authorizationRuleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2018-01-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listKeysHandleResponse handles the ListKeys response.
func (client *TopicsClient) listKeysHandleResponse(resp *azcore.Response) (TopicsListKeysResponse, error) {
	result := TopicsListKeysResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.AccessKeys); err != nil {
		return TopicsListKeysResponse{}, err
	}
	return result, nil
}

// listKeysHandleError handles the ListKeys error response.
func (client *TopicsClient) listKeysHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// RegenerateKeys - Regenerates primary or secondary connection strings for the topic.
// If the operation fails it returns the *ErrorResponse error type.
func (client *TopicsClient) RegenerateKeys(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, authorizationRuleName string, parameters RegenerateAccessKeyParameters, options *TopicsRegenerateKeysOptions) (TopicsRegenerateKeysResponse, error) {
	req, err := client.regenerateKeysCreateRequest(ctx, resourceGroupName, namespaceName, topicName, authorizationRuleName, parameters, options)
	if err != nil {
		return TopicsRegenerateKeysResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return TopicsRegenerateKeysResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return TopicsRegenerateKeysResponse{}, client.regenerateKeysHandleError(resp)
	}
	return client.regenerateKeysHandleResponse(resp)
}

// regenerateKeysCreateRequest creates the RegenerateKeys request.
func (client *TopicsClient) regenerateKeysCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, authorizationRuleName string, parameters RegenerateAccessKeyParameters, options *TopicsRegenerateKeysOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}/authorizationRules/{authorizationRuleName}/regenerateKeys"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if topicName == "" {
		return nil, errors.New("parameter topicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{topicName}", url.PathEscape(topicName))
	if authorizationRuleName == "" {
		return nil, errors.New("parameter authorizationRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorizationRuleName}", url.PathEscape(authorizationRuleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2018-01-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// regenerateKeysHandleResponse handles the RegenerateKeys response.
func (client *TopicsClient) regenerateKeysHandleResponse(resp *azcore.Response) (TopicsRegenerateKeysResponse, error) {
	result := TopicsRegenerateKeysResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.AccessKeys); err != nil {
		return TopicsRegenerateKeysResponse{}, err
	}
	return result, nil
}

// regenerateKeysHandleError handles the RegenerateKeys error response.
func (client *TopicsClient) regenerateKeysHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}
