//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcdn

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// LogAnalyticsClient contains the methods for the LogAnalytics group.
// Don't use this type directly, use NewLogAnalyticsClient() instead.
type LogAnalyticsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewLogAnalyticsClient creates a new instance of LogAnalyticsClient with the specified values.
func NewLogAnalyticsClient(con *arm.Connection, subscriptionID string) *LogAnalyticsClient {
	return &LogAnalyticsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// GetLogAnalyticsLocations - Get all available location names for AFD log analytics report.
// If the operation fails it returns the *AfdErrorResponse error type.
func (client *LogAnalyticsClient) GetLogAnalyticsLocations(ctx context.Context, resourceGroupName string, profileName string, options *LogAnalyticsGetLogAnalyticsLocationsOptions) (LogAnalyticsGetLogAnalyticsLocationsResponse, error) {
	req, err := client.getLogAnalyticsLocationsCreateRequest(ctx, resourceGroupName, profileName, options)
	if err != nil {
		return LogAnalyticsGetLogAnalyticsLocationsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LogAnalyticsGetLogAnalyticsLocationsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LogAnalyticsGetLogAnalyticsLocationsResponse{}, client.getLogAnalyticsLocationsHandleError(resp)
	}
	return client.getLogAnalyticsLocationsHandleResponse(resp)
}

// getLogAnalyticsLocationsCreateRequest creates the GetLogAnalyticsLocations request.
func (client *LogAnalyticsClient) getLogAnalyticsLocationsCreateRequest(ctx context.Context, resourceGroupName string, profileName string, options *LogAnalyticsGetLogAnalyticsLocationsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/getLogAnalyticsLocations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getLogAnalyticsLocationsHandleResponse handles the GetLogAnalyticsLocations response.
func (client *LogAnalyticsClient) getLogAnalyticsLocationsHandleResponse(resp *http.Response) (LogAnalyticsGetLogAnalyticsLocationsResponse, error) {
	result := LogAnalyticsGetLogAnalyticsLocationsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ContinentsResponse); err != nil {
		return LogAnalyticsGetLogAnalyticsLocationsResponse{}, err
	}
	return result, nil
}

// getLogAnalyticsLocationsHandleError handles the GetLogAnalyticsLocations error response.
func (client *LogAnalyticsClient) getLogAnalyticsLocationsHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := AfdErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// GetLogAnalyticsMetrics - Get log report for AFD profile
// If the operation fails it returns the *AfdErrorResponse error type.
func (client *LogAnalyticsClient) GetLogAnalyticsMetrics(ctx context.Context, resourceGroupName string, profileName string, metrics []LogMetric, dateTimeBegin time.Time, dateTimeEnd time.Time, granularity LogMetricsGranularity, customDomains []string, protocols []string, options *LogAnalyticsGetLogAnalyticsMetricsOptions) (LogAnalyticsGetLogAnalyticsMetricsResponse, error) {
	req, err := client.getLogAnalyticsMetricsCreateRequest(ctx, resourceGroupName, profileName, metrics, dateTimeBegin, dateTimeEnd, granularity, customDomains, protocols, options)
	if err != nil {
		return LogAnalyticsGetLogAnalyticsMetricsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LogAnalyticsGetLogAnalyticsMetricsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LogAnalyticsGetLogAnalyticsMetricsResponse{}, client.getLogAnalyticsMetricsHandleError(resp)
	}
	return client.getLogAnalyticsMetricsHandleResponse(resp)
}

// getLogAnalyticsMetricsCreateRequest creates the GetLogAnalyticsMetrics request.
func (client *LogAnalyticsClient) getLogAnalyticsMetricsCreateRequest(ctx context.Context, resourceGroupName string, profileName string, metrics []LogMetric, dateTimeBegin time.Time, dateTimeEnd time.Time, granularity LogMetricsGranularity, customDomains []string, protocols []string, options *LogAnalyticsGetLogAnalyticsMetricsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/getLogAnalyticsMetrics"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	for _, qv := range metrics {
		reqQP.Add("metrics", qv)
	}
	reqQP.Set("dateTimeBegin", dateTimeBegin.Format(time.RFC3339Nano))
	reqQP.Set("dateTimeEnd", dateTimeEnd.Format(time.RFC3339Nano))
	reqQP.Set("granularity", string(granularity))
	if options != nil && options.GroupBy != nil {
		for _, qv := range options.GroupBy {
			reqQP.Add("groupBy", qv)
		}
	}
	if options != nil && options.Continents != nil {
		for _, qv := range options.Continents {
			reqQP.Add("continents", qv)
		}
	}
	if options != nil && options.CountryOrRegions != nil {
		for _, qv := range options.CountryOrRegions {
			reqQP.Add("countryOrRegions", qv)
		}
	}
	for _, qv := range customDomains {
		reqQP.Add("customDomains", qv)
	}
	for _, qv := range protocols {
		reqQP.Add("protocols", qv)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getLogAnalyticsMetricsHandleResponse handles the GetLogAnalyticsMetrics response.
func (client *LogAnalyticsClient) getLogAnalyticsMetricsHandleResponse(resp *http.Response) (LogAnalyticsGetLogAnalyticsMetricsResponse, error) {
	result := LogAnalyticsGetLogAnalyticsMetricsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.MetricsResponse); err != nil {
		return LogAnalyticsGetLogAnalyticsMetricsResponse{}, err
	}
	return result, nil
}

// getLogAnalyticsMetricsHandleError handles the GetLogAnalyticsMetrics error response.
func (client *LogAnalyticsClient) getLogAnalyticsMetricsHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := AfdErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// GetLogAnalyticsRankings - Get log analytics ranking report for AFD profile
// If the operation fails it returns the *AfdErrorResponse error type.
func (client *LogAnalyticsClient) GetLogAnalyticsRankings(ctx context.Context, resourceGroupName string, profileName string, rankings []LogRanking, metrics []LogRankingMetric, maxRanking int32, dateTimeBegin time.Time, dateTimeEnd time.Time, options *LogAnalyticsGetLogAnalyticsRankingsOptions) (LogAnalyticsGetLogAnalyticsRankingsResponse, error) {
	req, err := client.getLogAnalyticsRankingsCreateRequest(ctx, resourceGroupName, profileName, rankings, metrics, maxRanking, dateTimeBegin, dateTimeEnd, options)
	if err != nil {
		return LogAnalyticsGetLogAnalyticsRankingsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LogAnalyticsGetLogAnalyticsRankingsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LogAnalyticsGetLogAnalyticsRankingsResponse{}, client.getLogAnalyticsRankingsHandleError(resp)
	}
	return client.getLogAnalyticsRankingsHandleResponse(resp)
}

// getLogAnalyticsRankingsCreateRequest creates the GetLogAnalyticsRankings request.
func (client *LogAnalyticsClient) getLogAnalyticsRankingsCreateRequest(ctx context.Context, resourceGroupName string, profileName string, rankings []LogRanking, metrics []LogRankingMetric, maxRanking int32, dateTimeBegin time.Time, dateTimeEnd time.Time, options *LogAnalyticsGetLogAnalyticsRankingsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/getLogAnalyticsRankings"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	for _, qv := range rankings {
		reqQP.Add("rankings", qv)
	}
	for _, qv := range metrics {
		reqQP.Add("metrics", qv)
	}
	reqQP.Set("maxRanking", strconv.FormatInt(int64(maxRanking), 10))
	reqQP.Set("dateTimeBegin", dateTimeBegin.Format(time.RFC3339Nano))
	reqQP.Set("dateTimeEnd", dateTimeEnd.Format(time.RFC3339Nano))
	if options != nil && options.CustomDomains != nil {
		for _, qv := range options.CustomDomains {
			reqQP.Add("customDomains", qv)
		}
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getLogAnalyticsRankingsHandleResponse handles the GetLogAnalyticsRankings response.
func (client *LogAnalyticsClient) getLogAnalyticsRankingsHandleResponse(resp *http.Response) (LogAnalyticsGetLogAnalyticsRankingsResponse, error) {
	result := LogAnalyticsGetLogAnalyticsRankingsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RankingsResponse); err != nil {
		return LogAnalyticsGetLogAnalyticsRankingsResponse{}, err
	}
	return result, nil
}

// getLogAnalyticsRankingsHandleError handles the GetLogAnalyticsRankings error response.
func (client *LogAnalyticsClient) getLogAnalyticsRankingsHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := AfdErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// GetLogAnalyticsResources - Get all endpoints and custom domains available for AFD log report
// If the operation fails it returns the *AfdErrorResponse error type.
func (client *LogAnalyticsClient) GetLogAnalyticsResources(ctx context.Context, resourceGroupName string, profileName string, options *LogAnalyticsGetLogAnalyticsResourcesOptions) (LogAnalyticsGetLogAnalyticsResourcesResponse, error) {
	req, err := client.getLogAnalyticsResourcesCreateRequest(ctx, resourceGroupName, profileName, options)
	if err != nil {
		return LogAnalyticsGetLogAnalyticsResourcesResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LogAnalyticsGetLogAnalyticsResourcesResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LogAnalyticsGetLogAnalyticsResourcesResponse{}, client.getLogAnalyticsResourcesHandleError(resp)
	}
	return client.getLogAnalyticsResourcesHandleResponse(resp)
}

// getLogAnalyticsResourcesCreateRequest creates the GetLogAnalyticsResources request.
func (client *LogAnalyticsClient) getLogAnalyticsResourcesCreateRequest(ctx context.Context, resourceGroupName string, profileName string, options *LogAnalyticsGetLogAnalyticsResourcesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/getLogAnalyticsResources"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getLogAnalyticsResourcesHandleResponse handles the GetLogAnalyticsResources response.
func (client *LogAnalyticsClient) getLogAnalyticsResourcesHandleResponse(resp *http.Response) (LogAnalyticsGetLogAnalyticsResourcesResponse, error) {
	result := LogAnalyticsGetLogAnalyticsResourcesResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourcesResponse); err != nil {
		return LogAnalyticsGetLogAnalyticsResourcesResponse{}, err
	}
	return result, nil
}

// getLogAnalyticsResourcesHandleError handles the GetLogAnalyticsResources error response.
func (client *LogAnalyticsClient) getLogAnalyticsResourcesHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := AfdErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// GetWafLogAnalyticsMetrics - Get Waf related log analytics report for AFD profile.
// If the operation fails it returns the *AfdErrorResponse error type.
func (client *LogAnalyticsClient) GetWafLogAnalyticsMetrics(ctx context.Context, resourceGroupName string, profileName string, metrics []WafMetric, dateTimeBegin time.Time, dateTimeEnd time.Time, granularity WafGranularity, options *LogAnalyticsGetWafLogAnalyticsMetricsOptions) (LogAnalyticsGetWafLogAnalyticsMetricsResponse, error) {
	req, err := client.getWafLogAnalyticsMetricsCreateRequest(ctx, resourceGroupName, profileName, metrics, dateTimeBegin, dateTimeEnd, granularity, options)
	if err != nil {
		return LogAnalyticsGetWafLogAnalyticsMetricsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LogAnalyticsGetWafLogAnalyticsMetricsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LogAnalyticsGetWafLogAnalyticsMetricsResponse{}, client.getWafLogAnalyticsMetricsHandleError(resp)
	}
	return client.getWafLogAnalyticsMetricsHandleResponse(resp)
}

// getWafLogAnalyticsMetricsCreateRequest creates the GetWafLogAnalyticsMetrics request.
func (client *LogAnalyticsClient) getWafLogAnalyticsMetricsCreateRequest(ctx context.Context, resourceGroupName string, profileName string, metrics []WafMetric, dateTimeBegin time.Time, dateTimeEnd time.Time, granularity WafGranularity, options *LogAnalyticsGetWafLogAnalyticsMetricsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/getWafLogAnalyticsMetrics"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	for _, qv := range metrics {
		reqQP.Add("metrics", qv)
	}
	reqQP.Set("dateTimeBegin", dateTimeBegin.Format(time.RFC3339Nano))
	reqQP.Set("dateTimeEnd", dateTimeEnd.Format(time.RFC3339Nano))
	reqQP.Set("granularity", string(granularity))
	if options != nil && options.Actions != nil {
		for _, qv := range options.Actions {
			reqQP.Add("actions", qv)
		}
	}
	if options != nil && options.GroupBy != nil {
		for _, qv := range options.GroupBy {
			reqQP.Add("groupBy", qv)
		}
	}
	if options != nil && options.RuleTypes != nil {
		for _, qv := range options.RuleTypes {
			reqQP.Add("ruleTypes", qv)
		}
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getWafLogAnalyticsMetricsHandleResponse handles the GetWafLogAnalyticsMetrics response.
func (client *LogAnalyticsClient) getWafLogAnalyticsMetricsHandleResponse(resp *http.Response) (LogAnalyticsGetWafLogAnalyticsMetricsResponse, error) {
	result := LogAnalyticsGetWafLogAnalyticsMetricsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.WafMetricsResponse); err != nil {
		return LogAnalyticsGetWafLogAnalyticsMetricsResponse{}, err
	}
	return result, nil
}

// getWafLogAnalyticsMetricsHandleError handles the GetWafLogAnalyticsMetrics error response.
func (client *LogAnalyticsClient) getWafLogAnalyticsMetricsHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := AfdErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// GetWafLogAnalyticsRankings - Get WAF log analytics charts for AFD profile
// If the operation fails it returns the *AfdErrorResponse error type.
func (client *LogAnalyticsClient) GetWafLogAnalyticsRankings(ctx context.Context, resourceGroupName string, profileName string, metrics []WafMetric, dateTimeBegin time.Time, dateTimeEnd time.Time, maxRanking int32, rankings []WafRankingType, options *LogAnalyticsGetWafLogAnalyticsRankingsOptions) (LogAnalyticsGetWafLogAnalyticsRankingsResponse, error) {
	req, err := client.getWafLogAnalyticsRankingsCreateRequest(ctx, resourceGroupName, profileName, metrics, dateTimeBegin, dateTimeEnd, maxRanking, rankings, options)
	if err != nil {
		return LogAnalyticsGetWafLogAnalyticsRankingsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LogAnalyticsGetWafLogAnalyticsRankingsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LogAnalyticsGetWafLogAnalyticsRankingsResponse{}, client.getWafLogAnalyticsRankingsHandleError(resp)
	}
	return client.getWafLogAnalyticsRankingsHandleResponse(resp)
}

// getWafLogAnalyticsRankingsCreateRequest creates the GetWafLogAnalyticsRankings request.
func (client *LogAnalyticsClient) getWafLogAnalyticsRankingsCreateRequest(ctx context.Context, resourceGroupName string, profileName string, metrics []WafMetric, dateTimeBegin time.Time, dateTimeEnd time.Time, maxRanking int32, rankings []WafRankingType, options *LogAnalyticsGetWafLogAnalyticsRankingsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/getWafLogAnalyticsRankings"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	for _, qv := range metrics {
		reqQP.Add("metrics", qv)
	}
	reqQP.Set("dateTimeBegin", dateTimeBegin.Format(time.RFC3339Nano))
	reqQP.Set("dateTimeEnd", dateTimeEnd.Format(time.RFC3339Nano))
	reqQP.Set("maxRanking", strconv.FormatInt(int64(maxRanking), 10))
	for _, qv := range rankings {
		reqQP.Add("rankings", qv)
	}
	if options != nil && options.Actions != nil {
		for _, qv := range options.Actions {
			reqQP.Add("actions", qv)
		}
	}
	if options != nil && options.RuleTypes != nil {
		for _, qv := range options.RuleTypes {
			reqQP.Add("ruleTypes", qv)
		}
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getWafLogAnalyticsRankingsHandleResponse handles the GetWafLogAnalyticsRankings response.
func (client *LogAnalyticsClient) getWafLogAnalyticsRankingsHandleResponse(resp *http.Response) (LogAnalyticsGetWafLogAnalyticsRankingsResponse, error) {
	result := LogAnalyticsGetWafLogAnalyticsRankingsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.WafRankingsResponse); err != nil {
		return LogAnalyticsGetWafLogAnalyticsRankingsResponse{}, err
	}
	return result, nil
}

// getWafLogAnalyticsRankingsHandleError handles the GetWafLogAnalyticsRankings error response.
func (client *LogAnalyticsClient) getWafLogAnalyticsRankingsHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := AfdErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
