package web

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

// BillingMetersClient is the webSite Management Client
type BillingMetersClient struct {
	BaseClient
}

// NewBillingMetersClient creates an instance of the BillingMetersClient client.
func NewBillingMetersClient(subscriptionID string) BillingMetersClient {
	return NewBillingMetersClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewBillingMetersClientWithBaseURI creates an instance of the BillingMetersClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewBillingMetersClientWithBaseURI(baseURI string, subscriptionID string) BillingMetersClient {
	return BillingMetersClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List gets a list of meters for a given location.
// Parameters:
// billingLocation - azure Location of billable resource
func (client BillingMetersClient) List(ctx context.Context, billingLocation string) (result BillingMeterCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BillingMetersClient.List")
		defer func() {
			sc := -1
			if result.bmc.Response.Response != nil {
				sc = result.bmc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, billingLocation)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.BillingMetersClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.bmc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "web.BillingMetersClient", "List", resp, "Failure sending request")
		return
	}

	result.bmc, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.BillingMetersClient", "List", resp, "Failure responding to request")
		return
	}
	if result.bmc.hasNextLink() && result.bmc.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client BillingMetersClient) ListPreparer(ctx context.Context, billingLocation string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(billingLocation) > 0 {
		queryParameters["billingLocation"] = autorest.Encode("query", billingLocation)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Web/billingMeters", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client BillingMetersClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client BillingMetersClient) ListResponder(resp *http.Response) (result BillingMeterCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client BillingMetersClient) listNextResults(ctx context.Context, lastResults BillingMeterCollection) (result BillingMeterCollection, err error) {
	req, err := lastResults.billingMeterCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web.BillingMetersClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web.BillingMetersClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.BillingMetersClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client BillingMetersClient) ListComplete(ctx context.Context, billingLocation string) (result BillingMeterCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BillingMetersClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, billingLocation)
	return
}
