package eventgrid

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// PrivateLinkResourcesClient is the azure EventGrid Management Client
type PrivateLinkResourcesClient struct {
	BaseClient
}

// NewPrivateLinkResourcesClient creates an instance of the PrivateLinkResourcesClient client.
func NewPrivateLinkResourcesClient(subscriptionID string) PrivateLinkResourcesClient {
	return NewPrivateLinkResourcesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPrivateLinkResourcesClientWithBaseURI creates an instance of the PrivateLinkResourcesClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewPrivateLinkResourcesClientWithBaseURI(baseURI string, subscriptionID string) PrivateLinkResourcesClient {
	return PrivateLinkResourcesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get get properties of a private link resource.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription.
// parentType - the type of the parent resource. This can be either \'topics\', \'domains\', or
// \'partnerNamespaces\'.
// parentName - the name of the parent resource (namely, either, the topic name, domain name, or partner
// namespace name).
// privateLinkResourceName - the name of private link resource.
func (client PrivateLinkResourcesClient) Get(ctx context.Context, resourceGroupName string, parentType string, parentName string, privateLinkResourceName string) (result PrivateLinkResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateLinkResourcesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, parentType, parentName, privateLinkResourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.PrivateLinkResourcesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "eventgrid.PrivateLinkResourcesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.PrivateLinkResourcesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client PrivateLinkResourcesClient) GetPreparer(ctx context.Context, resourceGroupName string, parentType string, parentName string, privateLinkResourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"parentName":              autorest.Encode("path", parentName),
		"parentType":              autorest.Encode("path", parentType),
		"privateLinkResourceName": autorest.Encode("path", privateLinkResourceName),
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-10-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/{parentType}/{parentName}/privateLinkResources/{privateLinkResourceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateLinkResourcesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client PrivateLinkResourcesClient) GetResponder(resp *http.Response) (result PrivateLinkResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResource list all the private link resources under a topic, domain, or partner namespace.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription.
// parentType - the type of the parent resource. This can be either \'topics\', \'domains\', or
// \'partnerNamespaces\'.
// parentName - the name of the parent resource (namely, either, the topic name, domain name, or partner
// namespace name).
// filter - the query used to filter the search results using OData syntax. Filtering is permitted on the
// 'name' property only and with limited number of OData operations. These operations are: the 'contains'
// function as well as the following logical operations: not, and, or, eq (for equal), and ne (for not equal).
// No arithmetic operations are supported. The following is a valid filter example: $filter=contains(namE,
// 'PATTERN') and name ne 'PATTERN-1'. The following is not a valid filter example: $filter=location eq
// 'westus'.
// top - the number of results to return per page for the list operation. Valid range for top parameter is 1 to
// 100. If not specified, the default number of results to be returned is 20 items per page.
func (client PrivateLinkResourcesClient) ListByResource(ctx context.Context, resourceGroupName string, parentType string, parentName string, filter string, top *int32) (result PrivateLinkResourcesListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateLinkResourcesClient.ListByResource")
		defer func() {
			sc := -1
			if result.plrlr.Response.Response != nil {
				sc = result.plrlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByResourceNextResults
	req, err := client.ListByResourcePreparer(ctx, resourceGroupName, parentType, parentName, filter, top)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.PrivateLinkResourcesClient", "ListByResource", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceSender(req)
	if err != nil {
		result.plrlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "eventgrid.PrivateLinkResourcesClient", "ListByResource", resp, "Failure sending request")
		return
	}

	result.plrlr, err = client.ListByResourceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.PrivateLinkResourcesClient", "ListByResource", resp, "Failure responding to request")
		return
	}
	if result.plrlr.hasNextLink() && result.plrlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByResourcePreparer prepares the ListByResource request.
func (client PrivateLinkResourcesClient) ListByResourcePreparer(ctx context.Context, resourceGroupName string, parentType string, parentName string, filter string, top *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"parentName":        autorest.Encode("path", parentName),
		"parentType":        autorest.Encode("path", parentType),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-10-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/{parentType}/{parentName}/privateLinkResources", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceSender sends the ListByResource request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateLinkResourcesClient) ListByResourceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceResponder handles the response to the ListByResource request. The method always
// closes the http.Response Body.
func (client PrivateLinkResourcesClient) ListByResourceResponder(resp *http.Response) (result PrivateLinkResourcesListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceNextResults retrieves the next set of results, if any.
func (client PrivateLinkResourcesClient) listByResourceNextResults(ctx context.Context, lastResults PrivateLinkResourcesListResult) (result PrivateLinkResourcesListResult, err error) {
	req, err := lastResults.privateLinkResourcesListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "eventgrid.PrivateLinkResourcesClient", "listByResourceNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "eventgrid.PrivateLinkResourcesClient", "listByResourceNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.PrivateLinkResourcesClient", "listByResourceNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceComplete enumerates all values, automatically crossing page boundaries as required.
func (client PrivateLinkResourcesClient) ListByResourceComplete(ctx context.Context, resourceGroupName string, parentType string, parentName string, filter string, top *int32) (result PrivateLinkResourcesListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateLinkResourcesClient.ListByResource")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByResource(ctx, resourceGroupName, parentType, parentName, filter, top)
	return
}
