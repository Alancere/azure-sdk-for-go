package security

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
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

// AllowedConnectionsClient is the API spec for Microsoft.Security (Azure Security Center) resource provider
type AllowedConnectionsClient struct {
	BaseClient
}

// NewAllowedConnectionsClient creates an instance of the AllowedConnectionsClient client.
func NewAllowedConnectionsClient(subscriptionID string, ascLocation string) AllowedConnectionsClient {
	return NewAllowedConnectionsClientWithBaseURI(DefaultBaseURI, subscriptionID, ascLocation)
}

// NewAllowedConnectionsClientWithBaseURI creates an instance of the AllowedConnectionsClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewAllowedConnectionsClientWithBaseURI(baseURI string, subscriptionID string, ascLocation string) AllowedConnectionsClient {
	return AllowedConnectionsClient{NewWithBaseURI(baseURI, subscriptionID, ascLocation)}
}

// Get gets the list of all possible traffic between resources for the subscription and location, based on connection
// type.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// connectionType - the type of allowed connections (Internal, External)
func (client AllowedConnectionsClient) Get(ctx context.Context, resourceGroupName string, connectionType ConnectionType) (result AllowedConnectionsResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AllowedConnectionsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("security.AllowedConnectionsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, connectionType)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.AllowedConnectionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "security.AllowedConnectionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.AllowedConnectionsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client AllowedConnectionsClient) GetPreparer(ctx context.Context, resourceGroupName string, connectionType ConnectionType) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"ascLocation":       autorest.Encode("path", client.AscLocation),
		"connectionType":    autorest.Encode("path", connectionType),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/locations/{ascLocation}/allowedConnections/{connectionType}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client AllowedConnectionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client AllowedConnectionsClient) GetResponder(resp *http.Response) (result AllowedConnectionsResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets the list of all possible traffic between resources for the subscription
func (client AllowedConnectionsClient) List(ctx context.Context) (result AllowedConnectionsListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AllowedConnectionsClient.List")
		defer func() {
			sc := -1
			if result.ACL.Response.Response != nil {
				sc = result.ACL.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("security.AllowedConnectionsClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.AllowedConnectionsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.ACL.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "security.AllowedConnectionsClient", "List", resp, "Failure sending request")
		return
	}

	result.ACL, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.AllowedConnectionsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.ACL.hasNextLink() && result.ACL.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client AllowedConnectionsClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Security/allowedConnections", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client AllowedConnectionsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client AllowedConnectionsClient) ListResponder(resp *http.Response) (result AllowedConnectionsList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client AllowedConnectionsClient) listNextResults(ctx context.Context, lastResults AllowedConnectionsList) (result AllowedConnectionsList, err error) {
	req, err := lastResults.allowedConnectionsListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "security.AllowedConnectionsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "security.AllowedConnectionsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.AllowedConnectionsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client AllowedConnectionsClient) ListComplete(ctx context.Context) (result AllowedConnectionsListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AllowedConnectionsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx)
	return
}

// ListByHomeRegion gets the list of all possible traffic between resources for the subscription and location.
func (client AllowedConnectionsClient) ListByHomeRegion(ctx context.Context) (result AllowedConnectionsListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AllowedConnectionsClient.ListByHomeRegion")
		defer func() {
			sc := -1
			if result.ACL.Response.Response != nil {
				sc = result.ACL.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("security.AllowedConnectionsClient", "ListByHomeRegion", err.Error())
	}

	result.fn = client.listByHomeRegionNextResults
	req, err := client.ListByHomeRegionPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.AllowedConnectionsClient", "ListByHomeRegion", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByHomeRegionSender(req)
	if err != nil {
		result.ACL.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "security.AllowedConnectionsClient", "ListByHomeRegion", resp, "Failure sending request")
		return
	}

	result.ACL, err = client.ListByHomeRegionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.AllowedConnectionsClient", "ListByHomeRegion", resp, "Failure responding to request")
		return
	}
	if result.ACL.hasNextLink() && result.ACL.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByHomeRegionPreparer prepares the ListByHomeRegion request.
func (client AllowedConnectionsClient) ListByHomeRegionPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"ascLocation":    autorest.Encode("path", client.AscLocation),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Security/locations/{ascLocation}/allowedConnections", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByHomeRegionSender sends the ListByHomeRegion request. The method will close the
// http.Response Body if it receives an error.
func (client AllowedConnectionsClient) ListByHomeRegionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByHomeRegionResponder handles the response to the ListByHomeRegion request. The method always
// closes the http.Response Body.
func (client AllowedConnectionsClient) ListByHomeRegionResponder(resp *http.Response) (result AllowedConnectionsList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByHomeRegionNextResults retrieves the next set of results, if any.
func (client AllowedConnectionsClient) listByHomeRegionNextResults(ctx context.Context, lastResults AllowedConnectionsList) (result AllowedConnectionsList, err error) {
	req, err := lastResults.allowedConnectionsListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "security.AllowedConnectionsClient", "listByHomeRegionNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByHomeRegionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "security.AllowedConnectionsClient", "listByHomeRegionNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByHomeRegionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.AllowedConnectionsClient", "listByHomeRegionNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByHomeRegionComplete enumerates all values, automatically crossing page boundaries as required.
func (client AllowedConnectionsClient) ListByHomeRegionComplete(ctx context.Context) (result AllowedConnectionsListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AllowedConnectionsClient.ListByHomeRegion")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByHomeRegion(ctx)
	return
}
