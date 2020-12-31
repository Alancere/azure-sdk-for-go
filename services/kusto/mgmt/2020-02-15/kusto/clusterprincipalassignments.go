package kusto

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

// ClusterPrincipalAssignmentsClient is the the Azure Kusto management API provides a RESTful set of web services that
// interact with Azure Kusto services to manage your clusters and databases. The API enables you to create, update, and
// delete clusters and databases.
type ClusterPrincipalAssignmentsClient struct {
	BaseClient
}

// NewClusterPrincipalAssignmentsClient creates an instance of the ClusterPrincipalAssignmentsClient client.
func NewClusterPrincipalAssignmentsClient(subscriptionID string) ClusterPrincipalAssignmentsClient {
	return NewClusterPrincipalAssignmentsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewClusterPrincipalAssignmentsClientWithBaseURI creates an instance of the ClusterPrincipalAssignmentsClient client
// using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewClusterPrincipalAssignmentsClientWithBaseURI(baseURI string, subscriptionID string) ClusterPrincipalAssignmentsClient {
	return ClusterPrincipalAssignmentsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CheckNameAvailability checks that the principal assignment name is valid and is not already in use.
// Parameters:
// resourceGroupName - the name of the resource group containing the Kusto cluster.
// clusterName - the name of the Kusto cluster.
// principalAssignmentName - the name of the principal assignment.
func (client ClusterPrincipalAssignmentsClient) CheckNameAvailability(ctx context.Context, resourceGroupName string, clusterName string, principalAssignmentName ClusterPrincipalAssignmentCheckNameRequest) (result CheckNameResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ClusterPrincipalAssignmentsClient.CheckNameAvailability")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: principalAssignmentName,
			Constraints: []validation.Constraint{{Target: "principalAssignmentName.Name", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "principalAssignmentName.Type", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("kusto.ClusterPrincipalAssignmentsClient", "CheckNameAvailability", err.Error())
	}

	req, err := client.CheckNameAvailabilityPreparer(ctx, resourceGroupName, clusterName, principalAssignmentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.ClusterPrincipalAssignmentsClient", "CheckNameAvailability", nil, "Failure preparing request")
		return
	}

	resp, err := client.CheckNameAvailabilitySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "kusto.ClusterPrincipalAssignmentsClient", "CheckNameAvailability", resp, "Failure sending request")
		return
	}

	result, err = client.CheckNameAvailabilityResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.ClusterPrincipalAssignmentsClient", "CheckNameAvailability", resp, "Failure responding to request")
		return
	}

	return
}

// CheckNameAvailabilityPreparer prepares the CheckNameAvailability request.
func (client ClusterPrincipalAssignmentsClient) CheckNameAvailabilityPreparer(ctx context.Context, resourceGroupName string, clusterName string, principalAssignmentName ClusterPrincipalAssignmentCheckNameRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-02-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/checkPrincipalAssignmentNameAvailability", pathParameters),
		autorest.WithJSON(principalAssignmentName),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CheckNameAvailabilitySender sends the CheckNameAvailability request. The method will close the
// http.Response Body if it receives an error.
func (client ClusterPrincipalAssignmentsClient) CheckNameAvailabilitySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CheckNameAvailabilityResponder handles the response to the CheckNameAvailability request. The method always
// closes the http.Response Body.
func (client ClusterPrincipalAssignmentsClient) CheckNameAvailabilityResponder(resp *http.Response) (result CheckNameResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CreateOrUpdate create a Kusto cluster principalAssignment.
// Parameters:
// resourceGroupName - the name of the resource group containing the Kusto cluster.
// clusterName - the name of the Kusto cluster.
// principalAssignmentName - the name of the Kusto principalAssignment.
// parameters - the Kusto cluster principalAssignment's parameters supplied for the operation.
func (client ClusterPrincipalAssignmentsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, principalAssignmentName string, parameters ClusterPrincipalAssignment) (result ClusterPrincipalAssignmentsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ClusterPrincipalAssignmentsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.ClusterPrincipalProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.ClusterPrincipalProperties.PrincipalID", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("kusto.ClusterPrincipalAssignmentsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, clusterName, principalAssignmentName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.ClusterPrincipalAssignmentsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.ClusterPrincipalAssignmentsClient", "CreateOrUpdate", nil, "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ClusterPrincipalAssignmentsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, clusterName string, principalAssignmentName string, parameters ClusterPrincipalAssignment) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":             autorest.Encode("path", clusterName),
		"principalAssignmentName": autorest.Encode("path", principalAssignmentName),
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-02-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/principalAssignments/{principalAssignmentName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ClusterPrincipalAssignmentsClient) CreateOrUpdateSender(req *http.Request) (future ClusterPrincipalAssignmentsCreateOrUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client ClusterPrincipalAssignmentsClient) (cpa ClusterPrincipalAssignment, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "kusto.ClusterPrincipalAssignmentsCreateOrUpdateFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("kusto.ClusterPrincipalAssignmentsCreateOrUpdateFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		if cpa.Response.Response, err = future.GetResult(sender); err == nil && cpa.Response.Response.StatusCode != http.StatusNoContent {
			cpa, err = client.CreateOrUpdateResponder(cpa.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "kusto.ClusterPrincipalAssignmentsCreateOrUpdateFuture", "Result", cpa.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ClusterPrincipalAssignmentsClient) CreateOrUpdateResponder(resp *http.Response) (result ClusterPrincipalAssignment, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a Kusto cluster principalAssignment.
// Parameters:
// resourceGroupName - the name of the resource group containing the Kusto cluster.
// clusterName - the name of the Kusto cluster.
// principalAssignmentName - the name of the Kusto principalAssignment.
func (client ClusterPrincipalAssignmentsClient) Delete(ctx context.Context, resourceGroupName string, clusterName string, principalAssignmentName string) (result ClusterPrincipalAssignmentsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ClusterPrincipalAssignmentsClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, clusterName, principalAssignmentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.ClusterPrincipalAssignmentsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.ClusterPrincipalAssignmentsClient", "Delete", nil, "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ClusterPrincipalAssignmentsClient) DeletePreparer(ctx context.Context, resourceGroupName string, clusterName string, principalAssignmentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":             autorest.Encode("path", clusterName),
		"principalAssignmentName": autorest.Encode("path", principalAssignmentName),
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-02-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/principalAssignments/{principalAssignmentName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ClusterPrincipalAssignmentsClient) DeleteSender(req *http.Request) (future ClusterPrincipalAssignmentsDeleteFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client ClusterPrincipalAssignmentsClient) (ar autorest.Response, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "kusto.ClusterPrincipalAssignmentsDeleteFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("kusto.ClusterPrincipalAssignmentsDeleteFuture")
			return
		}
		ar.Response = future.Response()
		return
	}
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ClusterPrincipalAssignmentsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets a Kusto cluster principalAssignment.
// Parameters:
// resourceGroupName - the name of the resource group containing the Kusto cluster.
// clusterName - the name of the Kusto cluster.
// principalAssignmentName - the name of the Kusto principalAssignment.
func (client ClusterPrincipalAssignmentsClient) Get(ctx context.Context, resourceGroupName string, clusterName string, principalAssignmentName string) (result ClusterPrincipalAssignment, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ClusterPrincipalAssignmentsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, clusterName, principalAssignmentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.ClusterPrincipalAssignmentsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "kusto.ClusterPrincipalAssignmentsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.ClusterPrincipalAssignmentsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ClusterPrincipalAssignmentsClient) GetPreparer(ctx context.Context, resourceGroupName string, clusterName string, principalAssignmentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":             autorest.Encode("path", clusterName),
		"principalAssignmentName": autorest.Encode("path", principalAssignmentName),
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-02-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/principalAssignments/{principalAssignmentName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ClusterPrincipalAssignmentsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ClusterPrincipalAssignmentsClient) GetResponder(resp *http.Response) (result ClusterPrincipalAssignment, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists all Kusto cluster principalAssignments.
// Parameters:
// resourceGroupName - the name of the resource group containing the Kusto cluster.
// clusterName - the name of the Kusto cluster.
func (client ClusterPrincipalAssignmentsClient) List(ctx context.Context, resourceGroupName string, clusterName string) (result ClusterPrincipalAssignmentListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ClusterPrincipalAssignmentsClient.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListPreparer(ctx, resourceGroupName, clusterName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.ClusterPrincipalAssignmentsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "kusto.ClusterPrincipalAssignmentsClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.ClusterPrincipalAssignmentsClient", "List", resp, "Failure responding to request")
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client ClusterPrincipalAssignmentsClient) ListPreparer(ctx context.Context, resourceGroupName string, clusterName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-02-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/principalAssignments", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ClusterPrincipalAssignmentsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ClusterPrincipalAssignmentsClient) ListResponder(resp *http.Response) (result ClusterPrincipalAssignmentListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
