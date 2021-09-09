package batchai

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

// FileServersClient is the the Azure BatchAI Management API.
type FileServersClient struct {
	BaseClient
}

// NewFileServersClient creates an instance of the FileServersClient client.
func NewFileServersClient(subscriptionID string) FileServersClient {
	return NewFileServersClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewFileServersClientWithBaseURI creates an instance of the FileServersClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewFileServersClientWithBaseURI(baseURI string, subscriptionID string) FileServersClient {
	return FileServersClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create creates a file server.
// Parameters:
// resourceGroupName - name of the resource group to which the resource belongs.
// fileServerName - the name of the file server within the specified resource group. File server names can only
// contain a combination of alphanumeric characters along with dash (-) and underscore (_). The name must be
// from 1 through 64 characters long.
// parameters - the parameters to provide for file server creation.
func (client FileServersClient) Create(ctx context.Context, resourceGroupName string, fileServerName string, parameters FileServerCreateParameters) (result FileServersCreateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FileServersClient.Create")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: fileServerName,
			Constraints: []validation.Constraint{{Target: "fileServerName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "fileServerName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "fileServerName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Location", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.FileServerBaseProperties", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameters.FileServerBaseProperties.VMSize", Name: validation.Null, Rule: true, Chain: nil},
						{Target: "parameters.FileServerBaseProperties.SSHConfiguration", Name: validation.Null, Rule: true,
							Chain: []validation.Constraint{{Target: "parameters.FileServerBaseProperties.SSHConfiguration.UserAccountSettings", Name: validation.Null, Rule: true,
								Chain: []validation.Constraint{{Target: "parameters.FileServerBaseProperties.SSHConfiguration.UserAccountSettings.AdminUserName", Name: validation.Null, Rule: true, Chain: nil}}},
							}},
						{Target: "parameters.FileServerBaseProperties.DataDisks", Name: validation.Null, Rule: true,
							Chain: []validation.Constraint{{Target: "parameters.FileServerBaseProperties.DataDisks.DiskSizeInGB", Name: validation.Null, Rule: true, Chain: nil},
								{Target: "parameters.FileServerBaseProperties.DataDisks.DiskCount", Name: validation.Null, Rule: true, Chain: nil},
							}},
						{Target: "parameters.FileServerBaseProperties.Subnet", Name: validation.Null, Rule: false,
							Chain: []validation.Constraint{{Target: "parameters.FileServerBaseProperties.Subnet.ID", Name: validation.Null, Rule: true, Chain: nil}}},
					}}}}}); err != nil {
		return result, validation.NewError("batchai.FileServersClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, fileServerName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.FileServersClient", "Create", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.FileServersClient", "Create", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client FileServersClient) CreatePreparer(ctx context.Context, resourceGroupName string, fileServerName string, parameters FileServerCreateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fileServerName":    autorest.Encode("path", fileServerName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/fileServers/{fileServerName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client FileServersClient) CreateSender(req *http.Request) (future FileServersCreateFuture, err error) {
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

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client FileServersClient) CreateResponder(resp *http.Response) (result FileServer, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete a file Server.
// Parameters:
// resourceGroupName - name of the resource group to which the resource belongs.
// fileServerName - the name of the file server within the specified resource group. File server names can only
// contain a combination of alphanumeric characters along with dash (-) and underscore (_). The name must be
// from 1 through 64 characters long.
func (client FileServersClient) Delete(ctx context.Context, resourceGroupName string, fileServerName string) (result FileServersDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FileServersClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: fileServerName,
			Constraints: []validation.Constraint{{Target: "fileServerName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "fileServerName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "fileServerName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("batchai.FileServersClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, fileServerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.FileServersClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.FileServersClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client FileServersClient) DeletePreparer(ctx context.Context, resourceGroupName string, fileServerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fileServerName":    autorest.Encode("path", fileServerName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/fileServers/{fileServerName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client FileServersClient) DeleteSender(req *http.Request) (future FileServersDeleteFuture, err error) {
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
func (client FileServersClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets information about the specified Cluster.
// Parameters:
// resourceGroupName - name of the resource group to which the resource belongs.
// fileServerName - the name of the file server within the specified resource group. File server names can only
// contain a combination of alphanumeric characters along with dash (-) and underscore (_). The name must be
// from 1 through 64 characters long.
func (client FileServersClient) Get(ctx context.Context, resourceGroupName string, fileServerName string) (result FileServer, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FileServersClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: fileServerName,
			Constraints: []validation.Constraint{{Target: "fileServerName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "fileServerName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "fileServerName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("batchai.FileServersClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, fileServerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.FileServersClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "batchai.FileServersClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.FileServersClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client FileServersClient) GetPreparer(ctx context.Context, resourceGroupName string, fileServerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fileServerName":    autorest.Encode("path", fileServerName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/fileServers/{fileServerName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client FileServersClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client FileServersClient) GetResponder(resp *http.Response) (result FileServer, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List to list all the file servers available under the given subscription (and across all resource groups within that
// subscription)
// Parameters:
// filter - an OData $filter clause. Used to filter results that are returned in the GET response.
// selectParameter - an OData $select clause. Used to select the properties to be returned in the GET response.
// maxResults - the maximum number of items to return in the response. A maximum of 1000 files can be returned.
func (client FileServersClient) List(ctx context.Context, filter string, selectParameter string, maxResults *int32) (result FileServerListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FileServersClient.List")
		defer func() {
			sc := -1
			if result.fslr.Response.Response != nil {
				sc = result.fslr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: maxResults,
			Constraints: []validation.Constraint{{Target: "maxResults", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "maxResults", Name: validation.InclusiveMaximum, Rule: int64(1000), Chain: nil},
					{Target: "maxResults", Name: validation.InclusiveMinimum, Rule: int64(1), Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("batchai.FileServersClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, filter, selectParameter, maxResults)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.FileServersClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.fslr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "batchai.FileServersClient", "List", resp, "Failure sending request")
		return
	}

	result.fslr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.FileServersClient", "List", resp, "Failure responding to request")
		return
	}
	if result.fslr.hasNextLink() && result.fslr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client FileServersClient) ListPreparer(ctx context.Context, filter string, selectParameter string, maxResults *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if len(selectParameter) > 0 {
		queryParameters["$select"] = autorest.Encode("query", selectParameter)
	}
	if maxResults != nil {
		queryParameters["maxresults"] = autorest.Encode("query", *maxResults)
	} else {
		queryParameters["maxresults"] = autorest.Encode("query", 1000)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.BatchAI/fileServers", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client FileServersClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client FileServersClient) ListResponder(resp *http.Response) (result FileServerListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client FileServersClient) listNextResults(ctx context.Context, lastResults FileServerListResult) (result FileServerListResult, err error) {
	req, err := lastResults.fileServerListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "batchai.FileServersClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "batchai.FileServersClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.FileServersClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client FileServersClient) ListComplete(ctx context.Context, filter string, selectParameter string, maxResults *int32) (result FileServerListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FileServersClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, filter, selectParameter, maxResults)
	return
}

// ListByResourceGroup gets a formatted list of file servers and their properties associated within the specified
// resource group.
// Parameters:
// resourceGroupName - name of the resource group to which the resource belongs.
// filter - an OData $filter clause. Used to filter results that are returned in the GET response.
// selectParameter - an OData $select clause. Used to select the properties to be returned in the GET response.
// maxResults - the maximum number of items to return in the response. A maximum of 1000 files can be returned.
func (client FileServersClient) ListByResourceGroup(ctx context.Context, resourceGroupName string, filter string, selectParameter string, maxResults *int32) (result FileServerListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FileServersClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.fslr.Response.Response != nil {
				sc = result.fslr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: maxResults,
			Constraints: []validation.Constraint{{Target: "maxResults", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "maxResults", Name: validation.InclusiveMaximum, Rule: int64(1000), Chain: nil},
					{Target: "maxResults", Name: validation.InclusiveMinimum, Rule: int64(1), Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("batchai.FileServersClient", "ListByResourceGroup", err.Error())
	}

	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName, filter, selectParameter, maxResults)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.FileServersClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.fslr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "batchai.FileServersClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.fslr, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.FileServersClient", "ListByResourceGroup", resp, "Failure responding to request")
		return
	}
	if result.fslr.hasNextLink() && result.fslr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client FileServersClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string, filter string, selectParameter string, maxResults *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if len(selectParameter) > 0 {
		queryParameters["$select"] = autorest.Encode("query", selectParameter)
	}
	if maxResults != nil {
		queryParameters["maxresults"] = autorest.Encode("query", *maxResults)
	} else {
		queryParameters["maxresults"] = autorest.Encode("query", 1000)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/fileServers", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client FileServersClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client FileServersClient) ListByResourceGroupResponder(resp *http.Response) (result FileServerListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client FileServersClient) listByResourceGroupNextResults(ctx context.Context, lastResults FileServerListResult) (result FileServerListResult, err error) {
	req, err := lastResults.fileServerListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "batchai.FileServersClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "batchai.FileServersClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.FileServersClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client FileServersClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, filter string, selectParameter string, maxResults *int32) (result FileServerListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FileServersClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName, filter, selectParameter, maxResults)
	return
}
