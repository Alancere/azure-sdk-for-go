package postgresqlhsc

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

// ConfigurationsClient is the the Microsoft Azure management API provides create, read, update, and delete
// functionality for Azure PostgreSQL Hyperscale (Citus) resources including server groups, servers, databases,
// firewall rules, VNET rules, security alert policies, log files and configurations.
type ConfigurationsClient struct {
	BaseClient
}

// NewConfigurationsClient creates an instance of the ConfigurationsClient client.
func NewConfigurationsClient(subscriptionID string) ConfigurationsClient {
	return NewConfigurationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewConfigurationsClientWithBaseURI creates an instance of the ConfigurationsClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewConfigurationsClientWithBaseURI(baseURI string, subscriptionID string) ConfigurationsClient {
	return ConfigurationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get gets information about single server group configuration.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// serverGroupName - the name of the server group.
// configurationName - the name of the server group configuration.
func (client ConfigurationsClient) Get(ctx context.Context, resourceGroupName string, serverGroupName string, configurationName string) (result ServerGroupConfiguration, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConfigurationsClient.Get")
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
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: serverGroupName,
			Constraints: []validation.Constraint{{Target: "serverGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "serverGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: configurationName,
			Constraints: []validation.Constraint{{Target: "configurationName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "configurationName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("postgresqlhsc.ConfigurationsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, serverGroupName, configurationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "postgresqlhsc.ConfigurationsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "postgresqlhsc.ConfigurationsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "postgresqlhsc.ConfigurationsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ConfigurationsClient) GetPreparer(ctx context.Context, resourceGroupName string, serverGroupName string, configurationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"configurationName": autorest.Encode("path", configurationName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverGroupName":   autorest.Encode("path", serverGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-10-05-privatepreview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBForPostgreSql/serverGroupsv2/{serverGroupName}/configurations/{configurationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ConfigurationsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ConfigurationsClient) GetResponder(resp *http.Response) (result ServerGroupConfiguration, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByServer list all the configurations of a server in server group.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// serverGroupName - the name of the server group.
// serverName - the name of the server.
func (client ConfigurationsClient) ListByServer(ctx context.Context, resourceGroupName string, serverGroupName string, serverName string) (result ServerConfigurationListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConfigurationsClient.ListByServer")
		defer func() {
			sc := -1
			if result.sclr.Response.Response != nil {
				sc = result.sclr.Response.Response.StatusCode
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
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: serverGroupName,
			Constraints: []validation.Constraint{{Target: "serverGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "serverGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: serverName,
			Constraints: []validation.Constraint{{Target: "serverName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "serverName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("postgresqlhsc.ConfigurationsClient", "ListByServer", err.Error())
	}

	result.fn = client.listByServerNextResults
	req, err := client.ListByServerPreparer(ctx, resourceGroupName, serverGroupName, serverName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "postgresqlhsc.ConfigurationsClient", "ListByServer", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByServerSender(req)
	if err != nil {
		result.sclr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "postgresqlhsc.ConfigurationsClient", "ListByServer", resp, "Failure sending request")
		return
	}

	result.sclr, err = client.ListByServerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "postgresqlhsc.ConfigurationsClient", "ListByServer", resp, "Failure responding to request")
		return
	}
	if result.sclr.hasNextLink() && result.sclr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByServerPreparer prepares the ListByServer request.
func (client ConfigurationsClient) ListByServerPreparer(ctx context.Context, resourceGroupName string, serverGroupName string, serverName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverGroupName":   autorest.Encode("path", serverGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-10-05-privatepreview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBForPostgreSql/serverGroupsv2/{serverGroupName}/servers/{serverName}/configurations", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByServerSender sends the ListByServer request. The method will close the
// http.Response Body if it receives an error.
func (client ConfigurationsClient) ListByServerSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByServerResponder handles the response to the ListByServer request. The method always
// closes the http.Response Body.
func (client ConfigurationsClient) ListByServerResponder(resp *http.Response) (result ServerConfigurationListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByServerNextResults retrieves the next set of results, if any.
func (client ConfigurationsClient) listByServerNextResults(ctx context.Context, lastResults ServerConfigurationListResult) (result ServerConfigurationListResult, err error) {
	req, err := lastResults.serverConfigurationListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "postgresqlhsc.ConfigurationsClient", "listByServerNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByServerSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "postgresqlhsc.ConfigurationsClient", "listByServerNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByServerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "postgresqlhsc.ConfigurationsClient", "listByServerNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByServerComplete enumerates all values, automatically crossing page boundaries as required.
func (client ConfigurationsClient) ListByServerComplete(ctx context.Context, resourceGroupName string, serverGroupName string, serverName string) (result ServerConfigurationListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConfigurationsClient.ListByServer")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByServer(ctx, resourceGroupName, serverGroupName, serverName)
	return
}

// ListByServerGroup list all the configurations of a server group.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// serverGroupName - the name of the server group.
func (client ConfigurationsClient) ListByServerGroup(ctx context.Context, resourceGroupName string, serverGroupName string) (result ServerGroupConfigurationListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConfigurationsClient.ListByServerGroup")
		defer func() {
			sc := -1
			if result.sgclr.Response.Response != nil {
				sc = result.sgclr.Response.Response.StatusCode
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
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: serverGroupName,
			Constraints: []validation.Constraint{{Target: "serverGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "serverGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("postgresqlhsc.ConfigurationsClient", "ListByServerGroup", err.Error())
	}

	result.fn = client.listByServerGroupNextResults
	req, err := client.ListByServerGroupPreparer(ctx, resourceGroupName, serverGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "postgresqlhsc.ConfigurationsClient", "ListByServerGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByServerGroupSender(req)
	if err != nil {
		result.sgclr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "postgresqlhsc.ConfigurationsClient", "ListByServerGroup", resp, "Failure sending request")
		return
	}

	result.sgclr, err = client.ListByServerGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "postgresqlhsc.ConfigurationsClient", "ListByServerGroup", resp, "Failure responding to request")
		return
	}
	if result.sgclr.hasNextLink() && result.sgclr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByServerGroupPreparer prepares the ListByServerGroup request.
func (client ConfigurationsClient) ListByServerGroupPreparer(ctx context.Context, resourceGroupName string, serverGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverGroupName":   autorest.Encode("path", serverGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-10-05-privatepreview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBForPostgreSql/serverGroupsv2/{serverGroupName}/configurations", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByServerGroupSender sends the ListByServerGroup request. The method will close the
// http.Response Body if it receives an error.
func (client ConfigurationsClient) ListByServerGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByServerGroupResponder handles the response to the ListByServerGroup request. The method always
// closes the http.Response Body.
func (client ConfigurationsClient) ListByServerGroupResponder(resp *http.Response) (result ServerGroupConfigurationListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByServerGroupNextResults retrieves the next set of results, if any.
func (client ConfigurationsClient) listByServerGroupNextResults(ctx context.Context, lastResults ServerGroupConfigurationListResult) (result ServerGroupConfigurationListResult, err error) {
	req, err := lastResults.serverGroupConfigurationListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "postgresqlhsc.ConfigurationsClient", "listByServerGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByServerGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "postgresqlhsc.ConfigurationsClient", "listByServerGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByServerGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "postgresqlhsc.ConfigurationsClient", "listByServerGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByServerGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client ConfigurationsClient) ListByServerGroupComplete(ctx context.Context, resourceGroupName string, serverGroupName string) (result ServerGroupConfigurationListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConfigurationsClient.ListByServerGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByServerGroup(ctx, resourceGroupName, serverGroupName)
	return
}

// Update updates configuration of server role groups in a server group
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// serverGroupName - the name of the server group.
// configurationName - the name of the server group configuration.
// parameters - the required parameters for updating a server group configuration.
func (client ConfigurationsClient) Update(ctx context.Context, resourceGroupName string, serverGroupName string, configurationName string, parameters ServerGroupConfiguration) (result ConfigurationsUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConfigurationsClient.Update")
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
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: serverGroupName,
			Constraints: []validation.Constraint{{Target: "serverGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "serverGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: configurationName,
			Constraints: []validation.Constraint{{Target: "configurationName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "configurationName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("postgresqlhsc.ConfigurationsClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, serverGroupName, configurationName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "postgresqlhsc.ConfigurationsClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "postgresqlhsc.ConfigurationsClient", "Update", nil, "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client ConfigurationsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, serverGroupName string, configurationName string, parameters ServerGroupConfiguration) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"configurationName": autorest.Encode("path", configurationName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverGroupName":   autorest.Encode("path", serverGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-10-05-privatepreview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBForPostgreSql/serverGroupsv2/{serverGroupName}/configurations/{configurationName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client ConfigurationsClient) UpdateSender(req *http.Request) (future ConfigurationsUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client ConfigurationsClient) (sgc ServerGroupConfiguration, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "postgresqlhsc.ConfigurationsUpdateFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("postgresqlhsc.ConfigurationsUpdateFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		if sgc.Response.Response, err = future.GetResult(sender); err == nil && sgc.Response.Response.StatusCode != http.StatusNoContent {
			sgc, err = client.UpdateResponder(sgc.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "postgresqlhsc.ConfigurationsUpdateFuture", "Result", sgc.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client ConfigurationsClient) UpdateResponder(resp *http.Response) (result ServerGroupConfiguration, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
