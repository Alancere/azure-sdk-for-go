package batch

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"github.com/gofrs/uuid"
	"net/http"
)

// AccountClient is the a client for issuing REST requests to the Azure Batch service.
type AccountClient struct {
	BaseClient
}

// NewAccountClient creates an instance of the AccountClient client.
func NewAccountClient(batchURL string) AccountClient {
	return AccountClient{New(batchURL)}
}

// ListNodeAgentSkus sends the list node agent skus request.
// Parameters:
// filter - an OData $filter clause. For more information on constructing this filter, see
// https://docs.microsoft.com/en-us/rest/api/batchservice/odata-filters-in-batch#list-node-agent-skus.
// maxResults - the maximum number of items to return in the response. A maximum of 1000 results will be
// returned.
// timeout - the maximum time that the server can spend processing the request, in seconds. The default is 30
// seconds.
// clientRequestID - the caller-generated request identity, in the form of a GUID with no decoration such as
// curly braces, e.g. 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0.
// returnClientRequestID - whether the server should return the client-request-id in the response.
// ocpDate - the time the request was issued. Client libraries typically set this to the current system clock
// time; set it explicitly if you are calling the REST API directly.
func (client AccountClient) ListNodeAgentSkus(ctx context.Context, filter string, maxResults *int32, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result AccountListNodeAgentSkusResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AccountClient.ListNodeAgentSkus")
		defer func() {
			sc := -1
			if result.alnasr.Response.Response != nil {
				sc = result.alnasr.Response.Response.StatusCode
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
		return result, validation.NewError("batch.AccountClient", "ListNodeAgentSkus", err.Error())
	}

	result.fn = client.listNodeAgentSkusNextResults
	req, err := client.ListNodeAgentSkusPreparer(ctx, filter, maxResults, timeout, clientRequestID, returnClientRequestID, ocpDate)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.AccountClient", "ListNodeAgentSkus", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListNodeAgentSkusSender(req)
	if err != nil {
		result.alnasr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "batch.AccountClient", "ListNodeAgentSkus", resp, "Failure sending request")
		return
	}

	result.alnasr, err = client.ListNodeAgentSkusResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.AccountClient", "ListNodeAgentSkus", resp, "Failure responding to request")
		return
	}
	if result.alnasr.hasNextLink() && result.alnasr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListNodeAgentSkusPreparer prepares the ListNodeAgentSkus request.
func (client AccountClient) ListNodeAgentSkusPreparer(ctx context.Context, filter string, maxResults *int32, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"":         autorest.Encode("path"),
		"batchUrl": client.BatchURL,
	}

	const APIVersion = "2018-12-01.8.0"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if maxResults != nil {
		queryParameters["maxresults"] = autorest.Encode("query", *maxResults)
	} else {
		queryParameters["maxresults"] = autorest.Encode("query", 1000)
	}
	if timeout != nil {
		queryParameters["timeout"] = autorest.Encode("query", *timeout)
	} else {
		queryParameters["timeout"] = autorest.Encode("query", 30)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{batchUrl}", urlParameters),
		autorest.WithPath("/nodeagentskus"),
		autorest.WithQueryParameters(queryParameters))
	if clientRequestID != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("client-request-id", autorest.String(clientRequestID)))
	}
	if returnClientRequestID != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("return-client-request-id", autorest.String(returnClientRequestID)))
	} else {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("return-client-request-id", autorest.String(false)))
	}
	if ocpDate != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("ocp-date", autorest.String(ocpDate)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListNodeAgentSkusSender sends the ListNodeAgentSkus request. The method will close the
// http.Response Body if it receives an error.
func (client AccountClient) ListNodeAgentSkusSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListNodeAgentSkusResponder handles the response to the ListNodeAgentSkus request. The method always
// closes the http.Response Body.
func (client AccountClient) ListNodeAgentSkusResponder(resp *http.Response) (result AccountListNodeAgentSkusResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNodeAgentSkusNextResults retrieves the next set of results, if any.
func (client AccountClient) listNodeAgentSkusNextResults(ctx context.Context, lastResults AccountListNodeAgentSkusResult) (result AccountListNodeAgentSkusResult, err error) {
	req, err := lastResults.accountListNodeAgentSkusResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "batch.AccountClient", "listNodeAgentSkusNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListNodeAgentSkusSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "batch.AccountClient", "listNodeAgentSkusNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListNodeAgentSkusResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.AccountClient", "listNodeAgentSkusNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListNodeAgentSkusComplete enumerates all values, automatically crossing page boundaries as required.
func (client AccountClient) ListNodeAgentSkusComplete(ctx context.Context, filter string, maxResults *int32, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result AccountListNodeAgentSkusResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AccountClient.ListNodeAgentSkus")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListNodeAgentSkus(ctx, filter, maxResults, timeout, clientRequestID, returnClientRequestID, ocpDate)
	return
}

// ListPoolNodeCounts gets the number of nodes in each state, grouped by pool.
// Parameters:
// filter - an OData $filter clause. For more information on constructing this filter, see
// https://docs.microsoft.com/en-us/rest/api/batchservice/odata-filters-in-batch.
// maxResults - the maximum number of items to return in the response.
// timeout - the maximum time that the server can spend processing the request, in seconds. The default is 30
// seconds.
// clientRequestID - the caller-generated request identity, in the form of a GUID with no decoration such as
// curly braces, e.g. 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0.
// returnClientRequestID - whether the server should return the client-request-id in the response.
// ocpDate - the time the request was issued. Client libraries typically set this to the current system clock
// time; set it explicitly if you are calling the REST API directly.
func (client AccountClient) ListPoolNodeCounts(ctx context.Context, filter string, maxResults *int32, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result PoolNodeCountsListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AccountClient.ListPoolNodeCounts")
		defer func() {
			sc := -1
			if result.pnclr.Response.Response != nil {
				sc = result.pnclr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: maxResults,
			Constraints: []validation.Constraint{{Target: "maxResults", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "maxResults", Name: validation.InclusiveMaximum, Rule: int64(10), Chain: nil},
					{Target: "maxResults", Name: validation.InclusiveMinimum, Rule: int64(1), Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("batch.AccountClient", "ListPoolNodeCounts", err.Error())
	}

	result.fn = client.listPoolNodeCountsNextResults
	req, err := client.ListPoolNodeCountsPreparer(ctx, filter, maxResults, timeout, clientRequestID, returnClientRequestID, ocpDate)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.AccountClient", "ListPoolNodeCounts", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListPoolNodeCountsSender(req)
	if err != nil {
		result.pnclr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "batch.AccountClient", "ListPoolNodeCounts", resp, "Failure sending request")
		return
	}

	result.pnclr, err = client.ListPoolNodeCountsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.AccountClient", "ListPoolNodeCounts", resp, "Failure responding to request")
		return
	}
	if result.pnclr.hasNextLink() && result.pnclr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPoolNodeCountsPreparer prepares the ListPoolNodeCounts request.
func (client AccountClient) ListPoolNodeCountsPreparer(ctx context.Context, filter string, maxResults *int32, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"":         autorest.Encode("path"),
		"batchUrl": client.BatchURL,
	}

	const APIVersion = "2018-12-01.8.0"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if maxResults != nil {
		queryParameters["maxresults"] = autorest.Encode("query", *maxResults)
	} else {
		queryParameters["maxresults"] = autorest.Encode("query", 10)
	}
	if timeout != nil {
		queryParameters["timeout"] = autorest.Encode("query", *timeout)
	} else {
		queryParameters["timeout"] = autorest.Encode("query", 30)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{batchUrl}", urlParameters),
		autorest.WithPath("/nodecounts"),
		autorest.WithQueryParameters(queryParameters))
	if clientRequestID != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("client-request-id", autorest.String(clientRequestID)))
	}
	if returnClientRequestID != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("return-client-request-id", autorest.String(returnClientRequestID)))
	} else {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("return-client-request-id", autorest.String(false)))
	}
	if ocpDate != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("ocp-date", autorest.String(ocpDate)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListPoolNodeCountsSender sends the ListPoolNodeCounts request. The method will close the
// http.Response Body if it receives an error.
func (client AccountClient) ListPoolNodeCountsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListPoolNodeCountsResponder handles the response to the ListPoolNodeCounts request. The method always
// closes the http.Response Body.
func (client AccountClient) ListPoolNodeCountsResponder(resp *http.Response) (result PoolNodeCountsListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listPoolNodeCountsNextResults retrieves the next set of results, if any.
func (client AccountClient) listPoolNodeCountsNextResults(ctx context.Context, lastResults PoolNodeCountsListResult) (result PoolNodeCountsListResult, err error) {
	req, err := lastResults.poolNodeCountsListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "batch.AccountClient", "listPoolNodeCountsNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListPoolNodeCountsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "batch.AccountClient", "listPoolNodeCountsNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListPoolNodeCountsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.AccountClient", "listPoolNodeCountsNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListPoolNodeCountsComplete enumerates all values, automatically crossing page boundaries as required.
func (client AccountClient) ListPoolNodeCountsComplete(ctx context.Context, filter string, maxResults *int32, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result PoolNodeCountsListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AccountClient.ListPoolNodeCounts")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListPoolNodeCounts(ctx, filter, maxResults, timeout, clientRequestID, returnClientRequestID, ocpDate)
	return
}
