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

// FirewallPolicyIdpsSignaturesFilterValuesClient is the network Client
type FirewallPolicyIdpsSignaturesFilterValuesClient struct {
	BaseClient
}

// NewFirewallPolicyIdpsSignaturesFilterValuesClient creates an instance of the
// FirewallPolicyIdpsSignaturesFilterValuesClient client.
func NewFirewallPolicyIdpsSignaturesFilterValuesClient(subscriptionID string) FirewallPolicyIdpsSignaturesFilterValuesClient {
	return NewFirewallPolicyIdpsSignaturesFilterValuesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewFirewallPolicyIdpsSignaturesFilterValuesClientWithBaseURI creates an instance of the
// FirewallPolicyIdpsSignaturesFilterValuesClient client using a custom endpoint.  Use this when interacting with an
// Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewFirewallPolicyIdpsSignaturesFilterValuesClientWithBaseURI(baseURI string, subscriptionID string) FirewallPolicyIdpsSignaturesFilterValuesClient {
	return FirewallPolicyIdpsSignaturesFilterValuesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List retrieves the current filter values for the signatures overrides
// Parameters:
// resourceGroupName - the name of the resource group.
// firewallPolicyName - the name of the Firewall Policy.
func (client FirewallPolicyIdpsSignaturesFilterValuesClient) List(ctx context.Context, parameters SignatureOverridesFilterValuesQuery, resourceGroupName string, firewallPolicyName string) (result SignatureOverridesFilterValuesResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FirewallPolicyIdpsSignaturesFilterValuesClient.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListPreparer(ctx, parameters, resourceGroupName, firewallPolicyName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.FirewallPolicyIdpsSignaturesFilterValuesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.FirewallPolicyIdpsSignaturesFilterValuesClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.FirewallPolicyIdpsSignaturesFilterValuesClient", "List", resp, "Failure responding to request")
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client FirewallPolicyIdpsSignaturesFilterValuesClient) ListPreparer(ctx context.Context, parameters SignatureOverridesFilterValuesQuery, resourceGroupName string, firewallPolicyName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"firewallPolicyName": autorest.Encode("path", firewallPolicyName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}/listIdpsFilterOptions", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client FirewallPolicyIdpsSignaturesFilterValuesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client FirewallPolicyIdpsSignaturesFilterValuesClient) ListResponder(resp *http.Response) (result SignatureOverridesFilterValuesResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
