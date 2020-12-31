package storagesync

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

// RegisteredServersClient is the microsoft Storage Sync Service API
type RegisteredServersClient struct {
	BaseClient
}

// NewRegisteredServersClient creates an instance of the RegisteredServersClient client.
func NewRegisteredServersClient(subscriptionID string) RegisteredServersClient {
	return NewRegisteredServersClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewRegisteredServersClientWithBaseURI creates an instance of the RegisteredServersClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewRegisteredServersClientWithBaseURI(baseURI string, subscriptionID string) RegisteredServersClient {
	return RegisteredServersClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create add a new registered server.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// storageSyncServiceName - name of Storage Sync Service resource.
// serverID - GUID identifying the on-premises server.
// parameters - body of Registered Server object.
func (client RegisteredServersClient) Create(ctx context.Context, resourceGroupName string, storageSyncServiceName string, serverID string, parameters RegisteredServerCreateParameters) (result RegisteredServersCreateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RegisteredServersClient.Create")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storagesync.RegisteredServersClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, storageSyncServiceName, serverID, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagesync.RegisteredServersClient", "Create", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagesync.RegisteredServersClient", "Create", nil, "Failure sending request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client RegisteredServersClient) CreatePreparer(ctx context.Context, resourceGroupName string, storageSyncServiceName string, serverID string, parameters RegisteredServerCreateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"serverId":               autorest.Encode("path", serverID),
		"storageSyncServiceName": autorest.Encode("path", storageSyncServiceName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/registeredServers/{serverId}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client RegisteredServersClient) CreateSender(req *http.Request) (future RegisteredServersCreateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client RegisteredServersClient) (rs RegisteredServer, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "storagesync.RegisteredServersCreateFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("storagesync.RegisteredServersCreateFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		if rs.Response.Response, err = future.GetResult(sender); err == nil && rs.Response.Response.StatusCode != http.StatusNoContent {
			rs, err = client.CreateResponder(rs.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "storagesync.RegisteredServersCreateFuture", "Result", rs.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client RegisteredServersClient) CreateResponder(resp *http.Response) (result RegisteredServer, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete the given registered server.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// storageSyncServiceName - name of Storage Sync Service resource.
// serverID - GUID identifying the on-premises server.
func (client RegisteredServersClient) Delete(ctx context.Context, resourceGroupName string, storageSyncServiceName string, serverID string) (result RegisteredServersDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RegisteredServersClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storagesync.RegisteredServersClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, storageSyncServiceName, serverID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagesync.RegisteredServersClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagesync.RegisteredServersClient", "Delete", nil, "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client RegisteredServersClient) DeletePreparer(ctx context.Context, resourceGroupName string, storageSyncServiceName string, serverID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"serverId":               autorest.Encode("path", serverID),
		"storageSyncServiceName": autorest.Encode("path", storageSyncServiceName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/registeredServers/{serverId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client RegisteredServersClient) DeleteSender(req *http.Request) (future RegisteredServersDeleteFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client RegisteredServersClient) (ar autorest.Response, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "storagesync.RegisteredServersDeleteFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("storagesync.RegisteredServersDeleteFuture")
			return
		}
		ar.Response = future.Response()
		return
	}
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client RegisteredServersClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get a given registered server.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// storageSyncServiceName - name of Storage Sync Service resource.
// serverID - GUID identifying the on-premises server.
func (client RegisteredServersClient) Get(ctx context.Context, resourceGroupName string, storageSyncServiceName string, serverID string) (result RegisteredServer, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RegisteredServersClient.Get")
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
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storagesync.RegisteredServersClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, storageSyncServiceName, serverID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagesync.RegisteredServersClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storagesync.RegisteredServersClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagesync.RegisteredServersClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client RegisteredServersClient) GetPreparer(ctx context.Context, resourceGroupName string, storageSyncServiceName string, serverID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"serverId":               autorest.Encode("path", serverID),
		"storageSyncServiceName": autorest.Encode("path", storageSyncServiceName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/registeredServers/{serverId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client RegisteredServersClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client RegisteredServersClient) GetResponder(resp *http.Response) (result RegisteredServer, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByStorageSyncService get a given registered server list.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// storageSyncServiceName - name of Storage Sync Service resource.
func (client RegisteredServersClient) ListByStorageSyncService(ctx context.Context, resourceGroupName string, storageSyncServiceName string) (result RegisteredServerArray, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RegisteredServersClient.ListByStorageSyncService")
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
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storagesync.RegisteredServersClient", "ListByStorageSyncService", err.Error())
	}

	req, err := client.ListByStorageSyncServicePreparer(ctx, resourceGroupName, storageSyncServiceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagesync.RegisteredServersClient", "ListByStorageSyncService", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByStorageSyncServiceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storagesync.RegisteredServersClient", "ListByStorageSyncService", resp, "Failure sending request")
		return
	}

	result, err = client.ListByStorageSyncServiceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagesync.RegisteredServersClient", "ListByStorageSyncService", resp, "Failure responding to request")
		return
	}

	return
}

// ListByStorageSyncServicePreparer prepares the ListByStorageSyncService request.
func (client RegisteredServersClient) ListByStorageSyncServicePreparer(ctx context.Context, resourceGroupName string, storageSyncServiceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"storageSyncServiceName": autorest.Encode("path", storageSyncServiceName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/registeredServers", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByStorageSyncServiceSender sends the ListByStorageSyncService request. The method will close the
// http.Response Body if it receives an error.
func (client RegisteredServersClient) ListByStorageSyncServiceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByStorageSyncServiceResponder handles the response to the ListByStorageSyncService request. The method always
// closes the http.Response Body.
func (client RegisteredServersClient) ListByStorageSyncServiceResponder(resp *http.Response) (result RegisteredServerArray, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// TriggerRollover triggers Server certificate rollover.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// storageSyncServiceName - name of Storage Sync Service resource.
// serverID - server Id
// parameters - body of Trigger Rollover request.
func (client RegisteredServersClient) TriggerRollover(ctx context.Context, resourceGroupName string, storageSyncServiceName string, serverID string, parameters TriggerRolloverRequest) (result RegisteredServersTriggerRolloverFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RegisteredServersClient.TriggerRollover")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storagesync.RegisteredServersClient", "TriggerRollover", err.Error())
	}

	req, err := client.TriggerRolloverPreparer(ctx, resourceGroupName, storageSyncServiceName, serverID, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagesync.RegisteredServersClient", "TriggerRollover", nil, "Failure preparing request")
		return
	}

	result, err = client.TriggerRolloverSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagesync.RegisteredServersClient", "TriggerRollover", nil, "Failure sending request")
		return
	}

	return
}

// TriggerRolloverPreparer prepares the TriggerRollover request.
func (client RegisteredServersClient) TriggerRolloverPreparer(ctx context.Context, resourceGroupName string, storageSyncServiceName string, serverID string, parameters TriggerRolloverRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"serverId":               autorest.Encode("path", serverID),
		"storageSyncServiceName": autorest.Encode("path", storageSyncServiceName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/registeredServers/{serverId}/triggerRollover", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// TriggerRolloverSender sends the TriggerRollover request. The method will close the
// http.Response Body if it receives an error.
func (client RegisteredServersClient) TriggerRolloverSender(req *http.Request) (future RegisteredServersTriggerRolloverFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client RegisteredServersClient) (ar autorest.Response, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "storagesync.RegisteredServersTriggerRolloverFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("storagesync.RegisteredServersTriggerRolloverFuture")
			return
		}
		ar.Response = future.Response()
		return
	}
	return
}

// TriggerRolloverResponder handles the response to the TriggerRollover request. The method always
// closes the http.Response Body.
func (client RegisteredServersClient) TriggerRolloverResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}
