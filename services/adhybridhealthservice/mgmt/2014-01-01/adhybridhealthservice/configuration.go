package adhybridhealthservice

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
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// ConfigurationClient is the REST APIs for Azure Active Directory Connect Health
type ConfigurationClient struct {
	BaseClient
}

// NewConfigurationClient creates an instance of the ConfigurationClient client.
func NewConfigurationClient() ConfigurationClient {
	return NewConfigurationClientWithBaseURI(DefaultBaseURI)
}

// NewConfigurationClientWithBaseURI creates an instance of the ConfigurationClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewConfigurationClientWithBaseURI(baseURI string) ConfigurationClient {
	return ConfigurationClient{NewWithBaseURI(baseURI)}
}

// Add onboards a tenant in Azure Active Directory Connect Health.
func (client ConfigurationClient) Add(ctx context.Context) (result Tenant, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConfigurationClient.Add")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.AddPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.ConfigurationClient", "Add", nil, "Failure preparing request")
		return
	}

	resp, err := client.AddSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.ConfigurationClient", "Add", resp, "Failure sending request")
		return
	}

	result, err = client.AddResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.ConfigurationClient", "Add", resp, "Failure responding to request")
		return
	}

	return
}

// AddPreparer prepares the Add request.
func (client ConfigurationClient) AddPreparer(ctx context.Context) (*http.Request, error) {
	const APIVersion = "2014-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.ADHybridHealthService/configuration"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// AddSender sends the Add request. The method will close the
// http.Response Body if it receives an error.
func (client ConfigurationClient) AddSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// AddResponder handles the response to the Add request. The method always
// closes the http.Response Body.
func (client ConfigurationClient) AddResponder(resp *http.Response) (result Tenant, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get gets the details of a tenant onboarded to Azure Active Directory Connect Health.
func (client ConfigurationClient) Get(ctx context.Context) (result Tenant, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConfigurationClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.ConfigurationClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.ConfigurationClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.ConfigurationClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ConfigurationClient) GetPreparer(ctx context.Context) (*http.Request, error) {
	const APIVersion = "2014-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.ADHybridHealthService/configuration"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ConfigurationClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ConfigurationClient) GetResponder(resp *http.Response) (result Tenant, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListAddsConfigurations gets the service configurations.
// Parameters:
// serviceName - the name of the service.
// grouping - the grouping for configurations.
func (client ConfigurationClient) ListAddsConfigurations(ctx context.Context, serviceName string, grouping string) (result AddsConfigurationPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConfigurationClient.ListAddsConfigurations")
		defer func() {
			sc := -1
			if result.ac.Response.Response != nil {
				sc = result.ac.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listAddsConfigurationsNextResults
	req, err := client.ListAddsConfigurationsPreparer(ctx, serviceName, grouping)
	if err != nil {
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.ConfigurationClient", "ListAddsConfigurations", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListAddsConfigurationsSender(req)
	if err != nil {
		result.ac.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.ConfigurationClient", "ListAddsConfigurations", resp, "Failure sending request")
		return
	}

	result.ac, err = client.ListAddsConfigurationsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.ConfigurationClient", "ListAddsConfigurations", resp, "Failure responding to request")
		return
	}
	if result.ac.hasNextLink() && result.ac.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListAddsConfigurationsPreparer prepares the ListAddsConfigurations request.
func (client ConfigurationClient) ListAddsConfigurationsPreparer(ctx context.Context, serviceName string, grouping string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"serviceName": autorest.Encode("path", serviceName),
	}

	queryParameters := map[string]interface{}{}
	if len(grouping) > 0 {
		queryParameters["grouping"] = autorest.Encode("query", grouping)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/configuration", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListAddsConfigurationsSender sends the ListAddsConfigurations request. The method will close the
// http.Response Body if it receives an error.
func (client ConfigurationClient) ListAddsConfigurationsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListAddsConfigurationsResponder handles the response to the ListAddsConfigurations request. The method always
// closes the http.Response Body.
func (client ConfigurationClient) ListAddsConfigurationsResponder(resp *http.Response) (result AddsConfiguration, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listAddsConfigurationsNextResults retrieves the next set of results, if any.
func (client ConfigurationClient) listAddsConfigurationsNextResults(ctx context.Context, lastResults AddsConfiguration) (result AddsConfiguration, err error) {
	req, err := lastResults.addsConfigurationPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "adhybridhealthservice.ConfigurationClient", "listAddsConfigurationsNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListAddsConfigurationsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "adhybridhealthservice.ConfigurationClient", "listAddsConfigurationsNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListAddsConfigurationsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.ConfigurationClient", "listAddsConfigurationsNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListAddsConfigurationsComplete enumerates all values, automatically crossing page boundaries as required.
func (client ConfigurationClient) ListAddsConfigurationsComplete(ctx context.Context, serviceName string, grouping string) (result AddsConfigurationIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConfigurationClient.ListAddsConfigurations")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListAddsConfigurations(ctx, serviceName, grouping)
	return
}

// Update updates tenant properties for tenants onboarded to Azure Active Directory Connect Health.
// Parameters:
// tenant - the tenant object with the properties set to the updated value.
func (client ConfigurationClient) Update(ctx context.Context, tenant Tenant) (result Tenant, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConfigurationClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, tenant)
	if err != nil {
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.ConfigurationClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.ConfigurationClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.ConfigurationClient", "Update", resp, "Failure responding to request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client ConfigurationClient) UpdatePreparer(ctx context.Context, tenant Tenant) (*http.Request, error) {
	const APIVersion = "2014-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.ADHybridHealthService/configuration"),
		autorest.WithJSON(tenant),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client ConfigurationClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client ConfigurationClient) UpdateResponder(resp *http.Response) (result Tenant, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
