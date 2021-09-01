package hybridcompute

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

// MachineExtensionsClient is the the Hybrid Compute Management Client.
type MachineExtensionsClient struct {
	BaseClient
}

// NewMachineExtensionsClient creates an instance of the MachineExtensionsClient client.
func NewMachineExtensionsClient(subscriptionID string) MachineExtensionsClient {
	return NewMachineExtensionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewMachineExtensionsClientWithBaseURI creates an instance of the MachineExtensionsClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewMachineExtensionsClientWithBaseURI(baseURI string, subscriptionID string) MachineExtensionsClient {
	return MachineExtensionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate the operation to create or update the extension.
// Parameters:
// resourceGroupName - the name of the resource group.
// name - the name of the machine where the extension should be created or updated.
// extensionName - the name of the machine extension.
// extensionParameters - parameters supplied to the Create Machine Extension operation.
func (client MachineExtensionsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, name string, extensionName string, extensionParameters MachineExtension) (result MachineExtensionsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MachineExtensionsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, name, extensionName, extensionParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybridcompute.MachineExtensionsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybridcompute.MachineExtensionsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client MachineExtensionsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, name string, extensionName string, extensionParameters MachineExtension) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"extensionName":     autorest.Encode("path", extensionName),
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-08-02"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/machines/{name}/extensions/{extensionName}", pathParameters),
		autorest.WithJSON(extensionParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client MachineExtensionsClient) CreateOrUpdateSender(req *http.Request) (future MachineExtensionsCreateOrUpdateFuture, err error) {
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
func (client MachineExtensionsClient) CreateOrUpdateResponder(resp *http.Response) (result MachineExtension, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete the operation to delete the extension.
// Parameters:
// resourceGroupName - the name of the resource group.
// name - the name of the machine where the extension should be deleted.
// extensionName - the name of the machine extension.
func (client MachineExtensionsClient) Delete(ctx context.Context, resourceGroupName string, name string, extensionName string) (result MachineExtensionsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MachineExtensionsClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, name, extensionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybridcompute.MachineExtensionsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybridcompute.MachineExtensionsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client MachineExtensionsClient) DeletePreparer(ctx context.Context, resourceGroupName string, name string, extensionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"extensionName":     autorest.Encode("path", extensionName),
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-08-02"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/machines/{name}/extensions/{extensionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client MachineExtensionsClient) DeleteSender(req *http.Request) (future MachineExtensionsDeleteFuture, err error) {
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
func (client MachineExtensionsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get the operation to get the extension.
// Parameters:
// resourceGroupName - the name of the resource group.
// name - the name of the machine containing the extension.
// extensionName - the name of the machine extension.
func (client MachineExtensionsClient) Get(ctx context.Context, resourceGroupName string, name string, extensionName string) (result MachineExtension, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MachineExtensionsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, name, extensionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybridcompute.MachineExtensionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "hybridcompute.MachineExtensionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybridcompute.MachineExtensionsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client MachineExtensionsClient) GetPreparer(ctx context.Context, resourceGroupName string, name string, extensionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"extensionName":     autorest.Encode("path", extensionName),
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-08-02"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/machines/{name}/extensions/{extensionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client MachineExtensionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client MachineExtensionsClient) GetResponder(resp *http.Response) (result MachineExtension, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List the operation to get all extensions of a non-Azure machine
// Parameters:
// resourceGroupName - the name of the resource group.
// name - the name of the machine containing the extension.
// expand - the expand expression to apply on the operation.
func (client MachineExtensionsClient) List(ctx context.Context, resourceGroupName string, name string, expand string) (result MachineExtensionsListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MachineExtensionsClient.List")
		defer func() {
			sc := -1
			if result.melr.Response.Response != nil {
				sc = result.melr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, name, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybridcompute.MachineExtensionsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.melr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "hybridcompute.MachineExtensionsClient", "List", resp, "Failure sending request")
		return
	}

	result.melr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybridcompute.MachineExtensionsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.melr.hasNextLink() && result.melr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client MachineExtensionsClient) ListPreparer(ctx context.Context, resourceGroupName string, name string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-08-02"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/machines/{name}/extensions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client MachineExtensionsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client MachineExtensionsClient) ListResponder(resp *http.Response) (result MachineExtensionsListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client MachineExtensionsClient) listNextResults(ctx context.Context, lastResults MachineExtensionsListResult) (result MachineExtensionsListResult, err error) {
	req, err := lastResults.machineExtensionsListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "hybridcompute.MachineExtensionsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "hybridcompute.MachineExtensionsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybridcompute.MachineExtensionsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client MachineExtensionsClient) ListComplete(ctx context.Context, resourceGroupName string, name string, expand string) (result MachineExtensionsListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MachineExtensionsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, name, expand)
	return
}

// Update the operation to update the extension.
// Parameters:
// resourceGroupName - the name of the resource group.
// name - the name of the machine where the extension should be created or updated.
// extensionName - the name of the machine extension.
// extensionParameters - parameters supplied to the Create Machine Extension operation.
func (client MachineExtensionsClient) Update(ctx context.Context, resourceGroupName string, name string, extensionName string, extensionParameters MachineExtensionUpdate) (result MachineExtensionsUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MachineExtensionsClient.Update")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, resourceGroupName, name, extensionName, extensionParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybridcompute.MachineExtensionsClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybridcompute.MachineExtensionsClient", "Update", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client MachineExtensionsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, name string, extensionName string, extensionParameters MachineExtensionUpdate) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"extensionName":     autorest.Encode("path", extensionName),
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-08-02"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/machines/{name}/extensions/{extensionName}", pathParameters),
		autorest.WithJSON(extensionParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client MachineExtensionsClient) UpdateSender(req *http.Request) (future MachineExtensionsUpdateFuture, err error) {
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

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client MachineExtensionsClient) UpdateResponder(resp *http.Response) (result MachineExtension, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
