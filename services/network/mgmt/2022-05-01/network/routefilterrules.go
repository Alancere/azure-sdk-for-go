package network

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// RouteFilterRulesClient is the network Client
type RouteFilterRulesClient struct {
	BaseClient
}

// NewRouteFilterRulesClient creates an instance of the RouteFilterRulesClient client.
func NewRouteFilterRulesClient(subscriptionID string) RouteFilterRulesClient {
	return NewRouteFilterRulesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewRouteFilterRulesClientWithBaseURI creates an instance of the RouteFilterRulesClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewRouteFilterRulesClientWithBaseURI(baseURI string, subscriptionID string) RouteFilterRulesClient {
	return RouteFilterRulesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a route in the specified route filter.
// Parameters:
// resourceGroupName - the name of the resource group.
// routeFilterName - the name of the route filter.
// ruleName - the name of the route filter rule.
// routeFilterRuleParameters - parameters supplied to the create or update route filter rule operation.
func (client RouteFilterRulesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, routeFilterName string, ruleName string, routeFilterRuleParameters RouteFilterRule) (result RouteFilterRulesCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RouteFilterRulesClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: routeFilterRuleParameters,
			Constraints: []validation.Constraint{{Target: "routeFilterRuleParameters.RouteFilterRulePropertiesFormat", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "routeFilterRuleParameters.RouteFilterRulePropertiesFormat.RouteFilterRuleType", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "routeFilterRuleParameters.RouteFilterRulePropertiesFormat.Communities", Name: validation.Null, Rule: true, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("network.RouteFilterRulesClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, routeFilterName, ruleName, routeFilterRuleParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.RouteFilterRulesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.RouteFilterRulesClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client RouteFilterRulesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, routeFilterName string, ruleName string, routeFilterRuleParameters RouteFilterRule) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"routeFilterName":   autorest.Encode("path", routeFilterName),
		"ruleName":          autorest.Encode("path", ruleName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	routeFilterRuleParameters.Etag = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}/routeFilterRules/{ruleName}", pathParameters),
		autorest.WithJSON(routeFilterRuleParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client RouteFilterRulesClient) CreateOrUpdateSender(req *http.Request) (future RouteFilterRulesCreateOrUpdateFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client RouteFilterRulesClient) CreateOrUpdateResponder(resp *http.Response) (result RouteFilterRule, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the specified rule from a route filter.
// Parameters:
// resourceGroupName - the name of the resource group.
// routeFilterName - the name of the route filter.
// ruleName - the name of the rule.
func (client RouteFilterRulesClient) Delete(ctx context.Context, resourceGroupName string, routeFilterName string, ruleName string) (result RouteFilterRulesDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RouteFilterRulesClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, routeFilterName, ruleName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.RouteFilterRulesClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.RouteFilterRulesClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client RouteFilterRulesClient) DeletePreparer(ctx context.Context, resourceGroupName string, routeFilterName string, ruleName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"routeFilterName":   autorest.Encode("path", routeFilterName),
		"ruleName":          autorest.Encode("path", ruleName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}/routeFilterRules/{ruleName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client RouteFilterRulesClient) DeleteSender(req *http.Request) (future RouteFilterRulesDeleteFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client RouteFilterRulesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the specified rule from a route filter.
// Parameters:
// resourceGroupName - the name of the resource group.
// routeFilterName - the name of the route filter.
// ruleName - the name of the rule.
func (client RouteFilterRulesClient) Get(ctx context.Context, resourceGroupName string, routeFilterName string, ruleName string) (result RouteFilterRule, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RouteFilterRulesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, routeFilterName, ruleName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.RouteFilterRulesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.RouteFilterRulesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.RouteFilterRulesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client RouteFilterRulesClient) GetPreparer(ctx context.Context, resourceGroupName string, routeFilterName string, ruleName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"routeFilterName":   autorest.Encode("path", routeFilterName),
		"ruleName":          autorest.Encode("path", ruleName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}/routeFilterRules/{ruleName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client RouteFilterRulesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client RouteFilterRulesClient) GetResponder(resp *http.Response) (result RouteFilterRule, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByRouteFilter gets all RouteFilterRules in a route filter.
// Parameters:
// resourceGroupName - the name of the resource group.
// routeFilterName - the name of the route filter.
func (client RouteFilterRulesClient) ListByRouteFilter(ctx context.Context, resourceGroupName string, routeFilterName string) (result RouteFilterRuleListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RouteFilterRulesClient.ListByRouteFilter")
		defer func() {
			sc := -1
			if result.rfrlr.Response.Response != nil {
				sc = result.rfrlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByRouteFilterNextResults
	req, err := client.ListByRouteFilterPreparer(ctx, resourceGroupName, routeFilterName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.RouteFilterRulesClient", "ListByRouteFilter", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByRouteFilterSender(req)
	if err != nil {
		result.rfrlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.RouteFilterRulesClient", "ListByRouteFilter", resp, "Failure sending request")
		return
	}

	result.rfrlr, err = client.ListByRouteFilterResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.RouteFilterRulesClient", "ListByRouteFilter", resp, "Failure responding to request")
		return
	}
	if result.rfrlr.hasNextLink() && result.rfrlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByRouteFilterPreparer prepares the ListByRouteFilter request.
func (client RouteFilterRulesClient) ListByRouteFilterPreparer(ctx context.Context, resourceGroupName string, routeFilterName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"routeFilterName":   autorest.Encode("path", routeFilterName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}/routeFilterRules", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByRouteFilterSender sends the ListByRouteFilter request. The method will close the
// http.Response Body if it receives an error.
func (client RouteFilterRulesClient) ListByRouteFilterSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByRouteFilterResponder handles the response to the ListByRouteFilter request. The method always
// closes the http.Response Body.
func (client RouteFilterRulesClient) ListByRouteFilterResponder(resp *http.Response) (result RouteFilterRuleListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByRouteFilterNextResults retrieves the next set of results, if any.
func (client RouteFilterRulesClient) listByRouteFilterNextResults(ctx context.Context, lastResults RouteFilterRuleListResult) (result RouteFilterRuleListResult, err error) {
	req, err := lastResults.routeFilterRuleListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.RouteFilterRulesClient", "listByRouteFilterNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByRouteFilterSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.RouteFilterRulesClient", "listByRouteFilterNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByRouteFilterResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.RouteFilterRulesClient", "listByRouteFilterNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByRouteFilterComplete enumerates all values, automatically crossing page boundaries as required.
func (client RouteFilterRulesClient) ListByRouteFilterComplete(ctx context.Context, resourceGroupName string, routeFilterName string) (result RouteFilterRuleListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RouteFilterRulesClient.ListByRouteFilter")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByRouteFilter(ctx, resourceGroupName, routeFilterName)
	return
}
