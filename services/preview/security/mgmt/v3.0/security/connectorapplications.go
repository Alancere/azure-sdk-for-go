package security

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

// ConnectorApplicationsClient is the API spec for Microsoft.Security (Azure Security Center) resource provider
type ConnectorApplicationsClient struct {
	BaseClient
}

// NewConnectorApplicationsClient creates an instance of the ConnectorApplicationsClient client.
func NewConnectorApplicationsClient(subscriptionID string) ConnectorApplicationsClient {
	return NewConnectorApplicationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewConnectorApplicationsClientWithBaseURI creates an instance of the ConnectorApplicationsClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewConnectorApplicationsClientWithBaseURI(baseURI string, subscriptionID string) ConnectorApplicationsClient {
	return ConnectorApplicationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List get a list of all relevant applications over a security connector level scope
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// securityConnectorName - the security connector name.
func (client ConnectorApplicationsClient) List(ctx context.Context, resourceGroupName string, securityConnectorName string) (result ApplicationsListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConnectorApplicationsClient.List")
		defer func() {
			sc := -1
			if result.al.Response.Response != nil {
				sc = result.al.Response.Response.StatusCode
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
		return result, validation.NewError("security.ConnectorApplicationsClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, securityConnectorName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.ConnectorApplicationsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.al.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "security.ConnectorApplicationsClient", "List", resp, "Failure sending request")
		return
	}

	result.al, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.ConnectorApplicationsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.al.hasNextLink() && result.al.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client ConnectorApplicationsClient) ListPreparer(ctx context.Context, resourceGroupName string, securityConnectorName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"securityConnectorName": autorest.Encode("path", securityConnectorName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/securityConnectors/{securityConnectorName}/providers/Microsoft.Security/applications", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ConnectorApplicationsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ConnectorApplicationsClient) ListResponder(resp *http.Response) (result ApplicationsList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client ConnectorApplicationsClient) listNextResults(ctx context.Context, lastResults ApplicationsList) (result ApplicationsList, err error) {
	req, err := lastResults.applicationsListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "security.ConnectorApplicationsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "security.ConnectorApplicationsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.ConnectorApplicationsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ConnectorApplicationsClient) ListComplete(ctx context.Context, resourceGroupName string, securityConnectorName string) (result ApplicationsListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConnectorApplicationsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, securityConnectorName)
	return
}
