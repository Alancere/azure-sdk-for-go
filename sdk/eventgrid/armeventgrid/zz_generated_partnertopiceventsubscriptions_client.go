//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armeventgrid

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// PartnerTopicEventSubscriptionsClient contains the methods for the PartnerTopicEventSubscriptions group.
// Don't use this type directly, use NewPartnerTopicEventSubscriptionsClient() instead.
type PartnerTopicEventSubscriptionsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewPartnerTopicEventSubscriptionsClient creates a new instance of PartnerTopicEventSubscriptionsClient with the specified values.
func NewPartnerTopicEventSubscriptionsClient(con *arm.Connection, subscriptionID string) *PartnerTopicEventSubscriptionsClient {
	return &PartnerTopicEventSubscriptionsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Asynchronously creates or updates an event subscription of a partner topic with the specified parameters. Existing event subscriptions
// will be updated with this API.
// If the operation fails it returns a generic error.
func (client *PartnerTopicEventSubscriptionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, partnerTopicName string, eventSubscriptionName string, eventSubscriptionInfo EventSubscription, options *PartnerTopicEventSubscriptionsBeginCreateOrUpdateOptions) (PartnerTopicEventSubscriptionsCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, partnerTopicName, eventSubscriptionName, eventSubscriptionInfo, options)
	if err != nil {
		return PartnerTopicEventSubscriptionsCreateOrUpdatePollerResponse{}, err
	}
	result := PartnerTopicEventSubscriptionsCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("PartnerTopicEventSubscriptionsClient.CreateOrUpdate", "", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return PartnerTopicEventSubscriptionsCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &PartnerTopicEventSubscriptionsCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Asynchronously creates or updates an event subscription of a partner topic with the specified parameters. Existing event subscriptions
// will be updated with this API.
// If the operation fails it returns a generic error.
func (client *PartnerTopicEventSubscriptionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, partnerTopicName string, eventSubscriptionName string, eventSubscriptionInfo EventSubscription, options *PartnerTopicEventSubscriptionsBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, partnerTopicName, eventSubscriptionName, eventSubscriptionInfo, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *PartnerTopicEventSubscriptionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, partnerTopicName string, eventSubscriptionName string, eventSubscriptionInfo EventSubscription, options *PartnerTopicEventSubscriptionsBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerTopics/{partnerTopicName}/eventSubscriptions/{eventSubscriptionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if partnerTopicName == "" {
		return nil, errors.New("parameter partnerTopicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{partnerTopicName}", url.PathEscape(partnerTopicName))
	if eventSubscriptionName == "" {
		return nil, errors.New("parameter eventSubscriptionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventSubscriptionName}", url.PathEscape(eventSubscriptionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, eventSubscriptionInfo)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *PartnerTopicEventSubscriptionsClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// BeginDelete - Delete an event subscription of a partner topic.
// If the operation fails it returns a generic error.
func (client *PartnerTopicEventSubscriptionsClient) BeginDelete(ctx context.Context, resourceGroupName string, partnerTopicName string, eventSubscriptionName string, options *PartnerTopicEventSubscriptionsBeginDeleteOptions) (PartnerTopicEventSubscriptionsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, partnerTopicName, eventSubscriptionName, options)
	if err != nil {
		return PartnerTopicEventSubscriptionsDeletePollerResponse{}, err
	}
	result := PartnerTopicEventSubscriptionsDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("PartnerTopicEventSubscriptionsClient.Delete", "", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return PartnerTopicEventSubscriptionsDeletePollerResponse{}, err
	}
	result.Poller = &PartnerTopicEventSubscriptionsDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Delete an event subscription of a partner topic.
// If the operation fails it returns a generic error.
func (client *PartnerTopicEventSubscriptionsClient) deleteOperation(ctx context.Context, resourceGroupName string, partnerTopicName string, eventSubscriptionName string, options *PartnerTopicEventSubscriptionsBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, partnerTopicName, eventSubscriptionName, options)
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
func (client *PartnerTopicEventSubscriptionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, partnerTopicName string, eventSubscriptionName string, options *PartnerTopicEventSubscriptionsBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerTopics/{partnerTopicName}/eventSubscriptions/{eventSubscriptionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if partnerTopicName == "" {
		return nil, errors.New("parameter partnerTopicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{partnerTopicName}", url.PathEscape(partnerTopicName))
	if eventSubscriptionName == "" {
		return nil, errors.New("parameter eventSubscriptionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventSubscriptionName}", url.PathEscape(eventSubscriptionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *PartnerTopicEventSubscriptionsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Get - Get an event subscription of a partner topic.
// If the operation fails it returns a generic error.
func (client *PartnerTopicEventSubscriptionsClient) Get(ctx context.Context, resourceGroupName string, partnerTopicName string, eventSubscriptionName string, options *PartnerTopicEventSubscriptionsGetOptions) (PartnerTopicEventSubscriptionsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, partnerTopicName, eventSubscriptionName, options)
	if err != nil {
		return PartnerTopicEventSubscriptionsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PartnerTopicEventSubscriptionsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PartnerTopicEventSubscriptionsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PartnerTopicEventSubscriptionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, partnerTopicName string, eventSubscriptionName string, options *PartnerTopicEventSubscriptionsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerTopics/{partnerTopicName}/eventSubscriptions/{eventSubscriptionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if partnerTopicName == "" {
		return nil, errors.New("parameter partnerTopicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{partnerTopicName}", url.PathEscape(partnerTopicName))
	if eventSubscriptionName == "" {
		return nil, errors.New("parameter eventSubscriptionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventSubscriptionName}", url.PathEscape(eventSubscriptionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PartnerTopicEventSubscriptionsClient) getHandleResponse(resp *http.Response) (PartnerTopicEventSubscriptionsGetResponse, error) {
	result := PartnerTopicEventSubscriptionsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.EventSubscription); err != nil {
		return PartnerTopicEventSubscriptionsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *PartnerTopicEventSubscriptionsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// GetDeliveryAttributes - Get all delivery attributes for an event subscription of a partner topic.
// If the operation fails it returns a generic error.
func (client *PartnerTopicEventSubscriptionsClient) GetDeliveryAttributes(ctx context.Context, resourceGroupName string, partnerTopicName string, eventSubscriptionName string, options *PartnerTopicEventSubscriptionsGetDeliveryAttributesOptions) (PartnerTopicEventSubscriptionsGetDeliveryAttributesResponse, error) {
	req, err := client.getDeliveryAttributesCreateRequest(ctx, resourceGroupName, partnerTopicName, eventSubscriptionName, options)
	if err != nil {
		return PartnerTopicEventSubscriptionsGetDeliveryAttributesResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PartnerTopicEventSubscriptionsGetDeliveryAttributesResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PartnerTopicEventSubscriptionsGetDeliveryAttributesResponse{}, client.getDeliveryAttributesHandleError(resp)
	}
	return client.getDeliveryAttributesHandleResponse(resp)
}

// getDeliveryAttributesCreateRequest creates the GetDeliveryAttributes request.
func (client *PartnerTopicEventSubscriptionsClient) getDeliveryAttributesCreateRequest(ctx context.Context, resourceGroupName string, partnerTopicName string, eventSubscriptionName string, options *PartnerTopicEventSubscriptionsGetDeliveryAttributesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerTopics/{partnerTopicName}/eventSubscriptions/{eventSubscriptionName}/getDeliveryAttributes"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if partnerTopicName == "" {
		return nil, errors.New("parameter partnerTopicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{partnerTopicName}", url.PathEscape(partnerTopicName))
	if eventSubscriptionName == "" {
		return nil, errors.New("parameter eventSubscriptionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventSubscriptionName}", url.PathEscape(eventSubscriptionName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getDeliveryAttributesHandleResponse handles the GetDeliveryAttributes response.
func (client *PartnerTopicEventSubscriptionsClient) getDeliveryAttributesHandleResponse(resp *http.Response) (PartnerTopicEventSubscriptionsGetDeliveryAttributesResponse, error) {
	result := PartnerTopicEventSubscriptionsGetDeliveryAttributesResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeliveryAttributeListResult); err != nil {
		return PartnerTopicEventSubscriptionsGetDeliveryAttributesResponse{}, err
	}
	return result, nil
}

// getDeliveryAttributesHandleError handles the GetDeliveryAttributes error response.
func (client *PartnerTopicEventSubscriptionsClient) getDeliveryAttributesHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// GetFullURL - Get the full endpoint URL for an event subscription of a partner topic.
// If the operation fails it returns a generic error.
func (client *PartnerTopicEventSubscriptionsClient) GetFullURL(ctx context.Context, resourceGroupName string, partnerTopicName string, eventSubscriptionName string, options *PartnerTopicEventSubscriptionsGetFullURLOptions) (PartnerTopicEventSubscriptionsGetFullURLResponse, error) {
	req, err := client.getFullURLCreateRequest(ctx, resourceGroupName, partnerTopicName, eventSubscriptionName, options)
	if err != nil {
		return PartnerTopicEventSubscriptionsGetFullURLResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PartnerTopicEventSubscriptionsGetFullURLResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PartnerTopicEventSubscriptionsGetFullURLResponse{}, client.getFullURLHandleError(resp)
	}
	return client.getFullURLHandleResponse(resp)
}

// getFullURLCreateRequest creates the GetFullURL request.
func (client *PartnerTopicEventSubscriptionsClient) getFullURLCreateRequest(ctx context.Context, resourceGroupName string, partnerTopicName string, eventSubscriptionName string, options *PartnerTopicEventSubscriptionsGetFullURLOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerTopics/{partnerTopicName}/eventSubscriptions/{eventSubscriptionName}/getFullUrl"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if partnerTopicName == "" {
		return nil, errors.New("parameter partnerTopicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{partnerTopicName}", url.PathEscape(partnerTopicName))
	if eventSubscriptionName == "" {
		return nil, errors.New("parameter eventSubscriptionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventSubscriptionName}", url.PathEscape(eventSubscriptionName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getFullURLHandleResponse handles the GetFullURL response.
func (client *PartnerTopicEventSubscriptionsClient) getFullURLHandleResponse(resp *http.Response) (PartnerTopicEventSubscriptionsGetFullURLResponse, error) {
	result := PartnerTopicEventSubscriptionsGetFullURLResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.EventSubscriptionFullURL); err != nil {
		return PartnerTopicEventSubscriptionsGetFullURLResponse{}, err
	}
	return result, nil
}

// getFullURLHandleError handles the GetFullURL error response.
func (client *PartnerTopicEventSubscriptionsClient) getFullURLHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListByPartnerTopic - List event subscriptions that belong to a specific partner topic.
// If the operation fails it returns a generic error.
func (client *PartnerTopicEventSubscriptionsClient) ListByPartnerTopic(resourceGroupName string, partnerTopicName string, options *PartnerTopicEventSubscriptionsListByPartnerTopicOptions) *PartnerTopicEventSubscriptionsListByPartnerTopicPager {
	return &PartnerTopicEventSubscriptionsListByPartnerTopicPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByPartnerTopicCreateRequest(ctx, resourceGroupName, partnerTopicName, options)
		},
		advancer: func(ctx context.Context, resp PartnerTopicEventSubscriptionsListByPartnerTopicResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.EventSubscriptionsListResult.NextLink)
		},
	}
}

// listByPartnerTopicCreateRequest creates the ListByPartnerTopic request.
func (client *PartnerTopicEventSubscriptionsClient) listByPartnerTopicCreateRequest(ctx context.Context, resourceGroupName string, partnerTopicName string, options *PartnerTopicEventSubscriptionsListByPartnerTopicOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerTopics/{partnerTopicName}/eventSubscriptions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if partnerTopicName == "" {
		return nil, errors.New("parameter partnerTopicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{partnerTopicName}", url.PathEscape(partnerTopicName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByPartnerTopicHandleResponse handles the ListByPartnerTopic response.
func (client *PartnerTopicEventSubscriptionsClient) listByPartnerTopicHandleResponse(resp *http.Response) (PartnerTopicEventSubscriptionsListByPartnerTopicResponse, error) {
	result := PartnerTopicEventSubscriptionsListByPartnerTopicResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.EventSubscriptionsListResult); err != nil {
		return PartnerTopicEventSubscriptionsListByPartnerTopicResponse{}, err
	}
	return result, nil
}

// listByPartnerTopicHandleError handles the ListByPartnerTopic error response.
func (client *PartnerTopicEventSubscriptionsClient) listByPartnerTopicHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// BeginUpdate - Update event subscription of a partner topic.
// If the operation fails it returns a generic error.
func (client *PartnerTopicEventSubscriptionsClient) BeginUpdate(ctx context.Context, resourceGroupName string, partnerTopicName string, eventSubscriptionName string, eventSubscriptionUpdateParameters EventSubscriptionUpdateParameters, options *PartnerTopicEventSubscriptionsBeginUpdateOptions) (PartnerTopicEventSubscriptionsUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, partnerTopicName, eventSubscriptionName, eventSubscriptionUpdateParameters, options)
	if err != nil {
		return PartnerTopicEventSubscriptionsUpdatePollerResponse{}, err
	}
	result := PartnerTopicEventSubscriptionsUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("PartnerTopicEventSubscriptionsClient.Update", "", resp, client.pl, client.updateHandleError)
	if err != nil {
		return PartnerTopicEventSubscriptionsUpdatePollerResponse{}, err
	}
	result.Poller = &PartnerTopicEventSubscriptionsUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Update event subscription of a partner topic.
// If the operation fails it returns a generic error.
func (client *PartnerTopicEventSubscriptionsClient) update(ctx context.Context, resourceGroupName string, partnerTopicName string, eventSubscriptionName string, eventSubscriptionUpdateParameters EventSubscriptionUpdateParameters, options *PartnerTopicEventSubscriptionsBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, partnerTopicName, eventSubscriptionName, eventSubscriptionUpdateParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated) {
		return nil, client.updateHandleError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *PartnerTopicEventSubscriptionsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, partnerTopicName string, eventSubscriptionName string, eventSubscriptionUpdateParameters EventSubscriptionUpdateParameters, options *PartnerTopicEventSubscriptionsBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerTopics/{partnerTopicName}/eventSubscriptions/{eventSubscriptionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if partnerTopicName == "" {
		return nil, errors.New("parameter partnerTopicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{partnerTopicName}", url.PathEscape(partnerTopicName))
	if eventSubscriptionName == "" {
		return nil, errors.New("parameter eventSubscriptionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventSubscriptionName}", url.PathEscape(eventSubscriptionName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, eventSubscriptionUpdateParameters)
}

// updateHandleError handles the Update error response.
func (client *PartnerTopicEventSubscriptionsClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
