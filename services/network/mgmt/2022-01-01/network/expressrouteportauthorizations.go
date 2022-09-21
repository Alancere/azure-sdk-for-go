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
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// ExpressRoutePortAuthorizationsClient is the network Client
type ExpressRoutePortAuthorizationsClient struct {
	BaseClient
}

// NewExpressRoutePortAuthorizationsClient creates an instance of the ExpressRoutePortAuthorizationsClient client.
func NewExpressRoutePortAuthorizationsClient(subscriptionID string) ExpressRoutePortAuthorizationsClient {
	return NewExpressRoutePortAuthorizationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewExpressRoutePortAuthorizationsClientWithBaseURI creates an instance of the ExpressRoutePortAuthorizationsClient
// client using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI
// (sovereign clouds, Azure stack).
func NewExpressRoutePortAuthorizationsClientWithBaseURI(baseURI string, subscriptionID string) ExpressRoutePortAuthorizationsClient {
	return ExpressRoutePortAuthorizationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates an authorization in the specified express route port.
// Parameters:
// resourceGroupName - the name of the resource group.
// expressRoutePortName - the name of the express route port.
// authorizationName - the name of the authorization.
// authorizationParameters - parameters supplied to the create or update express route port authorization
// operation.
func (client ExpressRoutePortAuthorizationsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, expressRoutePortName string, authorizationName string, authorizationParameters ExpressRoutePortAuthorization) (result ExpressRoutePortAuthorizationsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExpressRoutePortAuthorizationsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, expressRoutePortName, authorizationName, authorizationParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRoutePortAuthorizationsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRoutePortAuthorizationsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ExpressRoutePortAuthorizationsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, expressRoutePortName string, authorizationName string, authorizationParameters ExpressRoutePortAuthorization) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"authorizationName":    autorest.Encode("path", authorizationName),
		"expressRoutePortName": autorest.Encode("path", expressRoutePortName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	authorizationParameters.Etag = nil
	authorizationParameters.Type = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRoutePorts/{expressRoutePortName}/authorizations/{authorizationName}", pathParameters),
		autorest.WithJSON(authorizationParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRoutePortAuthorizationsClient) CreateOrUpdateSender(req *http.Request) (future ExpressRoutePortAuthorizationsCreateOrUpdateFuture, err error) {
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
func (client ExpressRoutePortAuthorizationsClient) CreateOrUpdateResponder(resp *http.Response) (result ExpressRoutePortAuthorization, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the specified authorization from the specified express route port.
// Parameters:
// resourceGroupName - the name of the resource group.
// expressRoutePortName - the name of the express route port.
// authorizationName - the name of the authorization.
func (client ExpressRoutePortAuthorizationsClient) Delete(ctx context.Context, resourceGroupName string, expressRoutePortName string, authorizationName string) (result ExpressRoutePortAuthorizationsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExpressRoutePortAuthorizationsClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, expressRoutePortName, authorizationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRoutePortAuthorizationsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRoutePortAuthorizationsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ExpressRoutePortAuthorizationsClient) DeletePreparer(ctx context.Context, resourceGroupName string, expressRoutePortName string, authorizationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"authorizationName":    autorest.Encode("path", authorizationName),
		"expressRoutePortName": autorest.Encode("path", expressRoutePortName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRoutePorts/{expressRoutePortName}/authorizations/{authorizationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRoutePortAuthorizationsClient) DeleteSender(req *http.Request) (future ExpressRoutePortAuthorizationsDeleteFuture, err error) {
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
func (client ExpressRoutePortAuthorizationsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the specified authorization from the specified express route port.
// Parameters:
// resourceGroupName - the name of the resource group.
// expressRoutePortName - the name of the express route port.
// authorizationName - the name of the authorization.
func (client ExpressRoutePortAuthorizationsClient) Get(ctx context.Context, resourceGroupName string, expressRoutePortName string, authorizationName string) (result ExpressRoutePortAuthorization, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExpressRoutePortAuthorizationsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, expressRoutePortName, authorizationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRoutePortAuthorizationsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.ExpressRoutePortAuthorizationsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRoutePortAuthorizationsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ExpressRoutePortAuthorizationsClient) GetPreparer(ctx context.Context, resourceGroupName string, expressRoutePortName string, authorizationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"authorizationName":    autorest.Encode("path", authorizationName),
		"expressRoutePortName": autorest.Encode("path", expressRoutePortName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRoutePorts/{expressRoutePortName}/authorizations/{authorizationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRoutePortAuthorizationsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ExpressRoutePortAuthorizationsClient) GetResponder(resp *http.Response) (result ExpressRoutePortAuthorization, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets all authorizations in an express route port.
// Parameters:
// resourceGroupName - the name of the resource group.
// expressRoutePortName - the name of the express route port.
func (client ExpressRoutePortAuthorizationsClient) List(ctx context.Context, resourceGroupName string, expressRoutePortName string) (result ExpressRoutePortAuthorizationListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExpressRoutePortAuthorizationsClient.List")
		defer func() {
			sc := -1
			if result.erpalr.Response.Response != nil {
				sc = result.erpalr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, expressRoutePortName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRoutePortAuthorizationsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.erpalr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.ExpressRoutePortAuthorizationsClient", "List", resp, "Failure sending request")
		return
	}

	result.erpalr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRoutePortAuthorizationsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.erpalr.hasNextLink() && result.erpalr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client ExpressRoutePortAuthorizationsClient) ListPreparer(ctx context.Context, resourceGroupName string, expressRoutePortName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"expressRoutePortName": autorest.Encode("path", expressRoutePortName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRoutePorts/{expressRoutePortName}/authorizations", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRoutePortAuthorizationsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ExpressRoutePortAuthorizationsClient) ListResponder(resp *http.Response) (result ExpressRoutePortAuthorizationListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client ExpressRoutePortAuthorizationsClient) listNextResults(ctx context.Context, lastResults ExpressRoutePortAuthorizationListResult) (result ExpressRoutePortAuthorizationListResult, err error) {
	req, err := lastResults.expressRoutePortAuthorizationListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.ExpressRoutePortAuthorizationsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.ExpressRoutePortAuthorizationsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRoutePortAuthorizationsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ExpressRoutePortAuthorizationsClient) ListComplete(ctx context.Context, resourceGroupName string, expressRoutePortName string) (result ExpressRoutePortAuthorizationListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExpressRoutePortAuthorizationsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, expressRoutePortName)
	return
}
