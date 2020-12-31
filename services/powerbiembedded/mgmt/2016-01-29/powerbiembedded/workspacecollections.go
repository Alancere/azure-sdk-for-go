package powerbiembedded

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

// WorkspaceCollectionsClient is the client to manage your Power BI Embedded workspace collections and retrieve
// workspaces.
type WorkspaceCollectionsClient struct {
	BaseClient
}

// NewWorkspaceCollectionsClient creates an instance of the WorkspaceCollectionsClient client.
func NewWorkspaceCollectionsClient(subscriptionID string) WorkspaceCollectionsClient {
	return NewWorkspaceCollectionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWorkspaceCollectionsClientWithBaseURI creates an instance of the WorkspaceCollectionsClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewWorkspaceCollectionsClientWithBaseURI(baseURI string, subscriptionID string) WorkspaceCollectionsClient {
	return WorkspaceCollectionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CheckNameAvailability verify the specified Power BI Workspace Collection name is valid and not already in use.
// Parameters:
// location - azure location
// body - check name availability request
func (client WorkspaceCollectionsClient) CheckNameAvailability(ctx context.Context, location string, body CheckNameRequest) (result CheckNameResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WorkspaceCollectionsClient.CheckNameAvailability")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CheckNameAvailabilityPreparer(ctx, location, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "powerbiembedded.WorkspaceCollectionsClient", "CheckNameAvailability", nil, "Failure preparing request")
		return
	}

	resp, err := client.CheckNameAvailabilitySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "powerbiembedded.WorkspaceCollectionsClient", "CheckNameAvailability", resp, "Failure sending request")
		return
	}

	result, err = client.CheckNameAvailabilityResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "powerbiembedded.WorkspaceCollectionsClient", "CheckNameAvailability", resp, "Failure responding to request")
		return
	}

	return
}

// CheckNameAvailabilityPreparer prepares the CheckNameAvailability request.
func (client WorkspaceCollectionsClient) CheckNameAvailabilityPreparer(ctx context.Context, location string, body CheckNameRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-01-29"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.PowerBI/locations/{location}/checkNameAvailability", pathParameters),
		autorest.WithJSON(body),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CheckNameAvailabilitySender sends the CheckNameAvailability request. The method will close the
// http.Response Body if it receives an error.
func (client WorkspaceCollectionsClient) CheckNameAvailabilitySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CheckNameAvailabilityResponder handles the response to the CheckNameAvailability request. The method always
// closes the http.Response Body.
func (client WorkspaceCollectionsClient) CheckNameAvailabilityResponder(resp *http.Response) (result CheckNameResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Create creates a new Power BI Workspace Collection with the specified properties. A Power BI Workspace Collection
// contains one or more workspaces, and can be used to provision keys that provide API access to those workspaces.
// Parameters:
// resourceGroupName - azure resource group
// workspaceCollectionName - power BI Embedded Workspace Collection name
// body - create workspace collection request
func (client WorkspaceCollectionsClient) Create(ctx context.Context, resourceGroupName string, workspaceCollectionName string, body CreateWorkspaceCollectionRequest) (result WorkspaceCollection, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WorkspaceCollectionsClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: body,
			Constraints: []validation.Constraint{{Target: "body.Sku", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "body.Sku.Name", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "body.Sku.Tier", Name: validation.Null, Rule: true, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("powerbiembedded.WorkspaceCollectionsClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, workspaceCollectionName, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "powerbiembedded.WorkspaceCollectionsClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "powerbiembedded.WorkspaceCollectionsClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "powerbiembedded.WorkspaceCollectionsClient", "Create", resp, "Failure responding to request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client WorkspaceCollectionsClient) CreatePreparer(ctx context.Context, resourceGroupName string, workspaceCollectionName string, body CreateWorkspaceCollectionRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
		"workspaceCollectionName": autorest.Encode("path", workspaceCollectionName),
	}

	const APIVersion = "2016-01-29"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBI/workspaceCollections/{workspaceCollectionName}", pathParameters),
		autorest.WithJSON(body),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client WorkspaceCollectionsClient) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client WorkspaceCollectionsClient) CreateResponder(resp *http.Response) (result WorkspaceCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete a Power BI Workspace Collection.
// Parameters:
// resourceGroupName - azure resource group
// workspaceCollectionName - power BI Embedded Workspace Collection name
func (client WorkspaceCollectionsClient) Delete(ctx context.Context, resourceGroupName string, workspaceCollectionName string) (result WorkspaceCollectionsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WorkspaceCollectionsClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, workspaceCollectionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "powerbiembedded.WorkspaceCollectionsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "powerbiembedded.WorkspaceCollectionsClient", "Delete", nil, "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client WorkspaceCollectionsClient) DeletePreparer(ctx context.Context, resourceGroupName string, workspaceCollectionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
		"workspaceCollectionName": autorest.Encode("path", workspaceCollectionName),
	}

	const APIVersion = "2016-01-29"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBI/workspaceCollections/{workspaceCollectionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client WorkspaceCollectionsClient) DeleteSender(req *http.Request) (future WorkspaceCollectionsDeleteFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client WorkspaceCollectionsClient) (ar autorest.Response, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "powerbiembedded.WorkspaceCollectionsDeleteFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("powerbiembedded.WorkspaceCollectionsDeleteFuture")
			return
		}
		ar.Response = future.Response()
		return
	}
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client WorkspaceCollectionsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// GetAccessKeys retrieves the primary and secondary access keys for the specified Power BI Workspace Collection.
// Parameters:
// resourceGroupName - azure resource group
// workspaceCollectionName - power BI Embedded Workspace Collection name
func (client WorkspaceCollectionsClient) GetAccessKeys(ctx context.Context, resourceGroupName string, workspaceCollectionName string) (result WorkspaceCollectionAccessKeys, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WorkspaceCollectionsClient.GetAccessKeys")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetAccessKeysPreparer(ctx, resourceGroupName, workspaceCollectionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "powerbiembedded.WorkspaceCollectionsClient", "GetAccessKeys", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetAccessKeysSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "powerbiembedded.WorkspaceCollectionsClient", "GetAccessKeys", resp, "Failure sending request")
		return
	}

	result, err = client.GetAccessKeysResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "powerbiembedded.WorkspaceCollectionsClient", "GetAccessKeys", resp, "Failure responding to request")
		return
	}

	return
}

// GetAccessKeysPreparer prepares the GetAccessKeys request.
func (client WorkspaceCollectionsClient) GetAccessKeysPreparer(ctx context.Context, resourceGroupName string, workspaceCollectionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
		"workspaceCollectionName": autorest.Encode("path", workspaceCollectionName),
	}

	const APIVersion = "2016-01-29"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBI/workspaceCollections/{workspaceCollectionName}/listKeys", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetAccessKeysSender sends the GetAccessKeys request. The method will close the
// http.Response Body if it receives an error.
func (client WorkspaceCollectionsClient) GetAccessKeysSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetAccessKeysResponder handles the response to the GetAccessKeys request. The method always
// closes the http.Response Body.
func (client WorkspaceCollectionsClient) GetAccessKeysResponder(resp *http.Response) (result WorkspaceCollectionAccessKeys, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetByName retrieves an existing Power BI Workspace Collection.
// Parameters:
// resourceGroupName - azure resource group
// workspaceCollectionName - power BI Embedded Workspace Collection name
func (client WorkspaceCollectionsClient) GetByName(ctx context.Context, resourceGroupName string, workspaceCollectionName string) (result WorkspaceCollection, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WorkspaceCollectionsClient.GetByName")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetByNamePreparer(ctx, resourceGroupName, workspaceCollectionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "powerbiembedded.WorkspaceCollectionsClient", "GetByName", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetByNameSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "powerbiembedded.WorkspaceCollectionsClient", "GetByName", resp, "Failure sending request")
		return
	}

	result, err = client.GetByNameResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "powerbiembedded.WorkspaceCollectionsClient", "GetByName", resp, "Failure responding to request")
		return
	}

	return
}

// GetByNamePreparer prepares the GetByName request.
func (client WorkspaceCollectionsClient) GetByNamePreparer(ctx context.Context, resourceGroupName string, workspaceCollectionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
		"workspaceCollectionName": autorest.Encode("path", workspaceCollectionName),
	}

	const APIVersion = "2016-01-29"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBI/workspaceCollections/{workspaceCollectionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetByNameSender sends the GetByName request. The method will close the
// http.Response Body if it receives an error.
func (client WorkspaceCollectionsClient) GetByNameSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetByNameResponder handles the response to the GetByName request. The method always
// closes the http.Response Body.
func (client WorkspaceCollectionsClient) GetByNameResponder(resp *http.Response) (result WorkspaceCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroup retrieves all existing Power BI workspace collections in the specified resource group.
// Parameters:
// resourceGroupName - azure resource group
func (client WorkspaceCollectionsClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (result WorkspaceCollectionList, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WorkspaceCollectionsClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "powerbiembedded.WorkspaceCollectionsClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "powerbiembedded.WorkspaceCollectionsClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "powerbiembedded.WorkspaceCollectionsClient", "ListByResourceGroup", resp, "Failure responding to request")
		return
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client WorkspaceCollectionsClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-01-29"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBI/workspaceCollections", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client WorkspaceCollectionsClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client WorkspaceCollectionsClient) ListByResourceGroupResponder(resp *http.Response) (result WorkspaceCollectionList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListBySubscription retrieves all existing Power BI workspace collections in the specified subscription.
func (client WorkspaceCollectionsClient) ListBySubscription(ctx context.Context) (result WorkspaceCollectionList, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WorkspaceCollectionsClient.ListBySubscription")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListBySubscriptionPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "powerbiembedded.WorkspaceCollectionsClient", "ListBySubscription", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "powerbiembedded.WorkspaceCollectionsClient", "ListBySubscription", resp, "Failure sending request")
		return
	}

	result, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "powerbiembedded.WorkspaceCollectionsClient", "ListBySubscription", resp, "Failure responding to request")
		return
	}

	return
}

// ListBySubscriptionPreparer prepares the ListBySubscription request.
func (client WorkspaceCollectionsClient) ListBySubscriptionPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-01-29"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.PowerBI/workspaceCollections", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListBySubscriptionSender sends the ListBySubscription request. The method will close the
// http.Response Body if it receives an error.
func (client WorkspaceCollectionsClient) ListBySubscriptionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListBySubscriptionResponder handles the response to the ListBySubscription request. The method always
// closes the http.Response Body.
func (client WorkspaceCollectionsClient) ListBySubscriptionResponder(resp *http.Response) (result WorkspaceCollectionList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Migrate migrates an existing Power BI Workspace Collection to a different resource group and/or subscription.
// Parameters:
// resourceGroupName - azure resource group
// body - workspace migration request
func (client WorkspaceCollectionsClient) Migrate(ctx context.Context, resourceGroupName string, body MigrateWorkspaceCollectionRequest) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WorkspaceCollectionsClient.Migrate")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.MigratePreparer(ctx, resourceGroupName, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "powerbiembedded.WorkspaceCollectionsClient", "Migrate", nil, "Failure preparing request")
		return
	}

	resp, err := client.MigrateSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "powerbiembedded.WorkspaceCollectionsClient", "Migrate", resp, "Failure sending request")
		return
	}

	result, err = client.MigrateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "powerbiembedded.WorkspaceCollectionsClient", "Migrate", resp, "Failure responding to request")
		return
	}

	return
}

// MigratePreparer prepares the Migrate request.
func (client WorkspaceCollectionsClient) MigratePreparer(ctx context.Context, resourceGroupName string, body MigrateWorkspaceCollectionRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-01-29"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/moveResources", pathParameters),
		autorest.WithJSON(body),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// MigrateSender sends the Migrate request. The method will close the
// http.Response Body if it receives an error.
func (client WorkspaceCollectionsClient) MigrateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// MigrateResponder handles the response to the Migrate request. The method always
// closes the http.Response Body.
func (client WorkspaceCollectionsClient) MigrateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// RegenerateKey regenerates the primary or secondary access key for the specified Power BI Workspace Collection.
// Parameters:
// resourceGroupName - azure resource group
// workspaceCollectionName - power BI Embedded Workspace Collection name
// body - access key to regenerate
func (client WorkspaceCollectionsClient) RegenerateKey(ctx context.Context, resourceGroupName string, workspaceCollectionName string, body WorkspaceCollectionAccessKey) (result WorkspaceCollectionAccessKeys, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WorkspaceCollectionsClient.RegenerateKey")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.RegenerateKeyPreparer(ctx, resourceGroupName, workspaceCollectionName, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "powerbiembedded.WorkspaceCollectionsClient", "RegenerateKey", nil, "Failure preparing request")
		return
	}

	resp, err := client.RegenerateKeySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "powerbiembedded.WorkspaceCollectionsClient", "RegenerateKey", resp, "Failure sending request")
		return
	}

	result, err = client.RegenerateKeyResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "powerbiembedded.WorkspaceCollectionsClient", "RegenerateKey", resp, "Failure responding to request")
		return
	}

	return
}

// RegenerateKeyPreparer prepares the RegenerateKey request.
func (client WorkspaceCollectionsClient) RegenerateKeyPreparer(ctx context.Context, resourceGroupName string, workspaceCollectionName string, body WorkspaceCollectionAccessKey) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
		"workspaceCollectionName": autorest.Encode("path", workspaceCollectionName),
	}

	const APIVersion = "2016-01-29"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBI/workspaceCollections/{workspaceCollectionName}/regenerateKey", pathParameters),
		autorest.WithJSON(body),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RegenerateKeySender sends the RegenerateKey request. The method will close the
// http.Response Body if it receives an error.
func (client WorkspaceCollectionsClient) RegenerateKeySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// RegenerateKeyResponder handles the response to the RegenerateKey request. The method always
// closes the http.Response Body.
func (client WorkspaceCollectionsClient) RegenerateKeyResponder(resp *http.Response) (result WorkspaceCollectionAccessKeys, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update update an existing Power BI Workspace Collection with the specified properties.
// Parameters:
// resourceGroupName - azure resource group
// workspaceCollectionName - power BI Embedded Workspace Collection name
// body - update workspace collection request
func (client WorkspaceCollectionsClient) Update(ctx context.Context, resourceGroupName string, workspaceCollectionName string, body UpdateWorkspaceCollectionRequest) (result WorkspaceCollection, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WorkspaceCollectionsClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, resourceGroupName, workspaceCollectionName, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "powerbiembedded.WorkspaceCollectionsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "powerbiembedded.WorkspaceCollectionsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "powerbiembedded.WorkspaceCollectionsClient", "Update", resp, "Failure responding to request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client WorkspaceCollectionsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, workspaceCollectionName string, body UpdateWorkspaceCollectionRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
		"workspaceCollectionName": autorest.Encode("path", workspaceCollectionName),
	}

	const APIVersion = "2016-01-29"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBI/workspaceCollections/{workspaceCollectionName}", pathParameters),
		autorest.WithJSON(body),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client WorkspaceCollectionsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client WorkspaceCollectionsClient) UpdateResponder(resp *http.Response) (result WorkspaceCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
