package reservations

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

// Client is the client for the Reservation methods of the Reservations service.
type Client struct {
	BaseClient
}

// NewClient creates an instance of the Client client.
func NewClient() Client {
	return NewClientWithBaseURI(DefaultBaseURI)
}

// NewClientWithBaseURI creates an instance of the Client client using a custom endpoint.  Use this when interacting
// with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewClientWithBaseURI(baseURI string) Client {
	return Client{NewWithBaseURI(baseURI)}
}

// Archive archiving a `Reservation` moves it to `Archived` state.
// Parameters:
// reservationOrderID - order Id of the reservation
// reservationID - id of the Reservation Item
func (client Client) Archive(ctx context.Context, reservationOrderID string, reservationID string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.Archive")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ArchivePreparer(ctx, reservationOrderID, reservationID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservations.Client", "Archive", nil, "Failure preparing request")
		return
	}

	resp, err := client.ArchiveSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "reservations.Client", "Archive", resp, "Failure sending request")
		return
	}

	result, err = client.ArchiveResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservations.Client", "Archive", resp, "Failure responding to request")
		return
	}

	return
}

// ArchivePreparer prepares the Archive request.
func (client Client) ArchivePreparer(ctx context.Context, reservationOrderID string, reservationID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"reservationId":      autorest.Encode("path", reservationID),
		"reservationOrderId": autorest.Encode("path", reservationOrderID),
	}

	const APIVersion = "2019-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Capacity/reservationOrders/{reservationOrderId}/reservations/{reservationId}/archive", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ArchiveSender sends the Archive request. The method will close the
// http.Response Body if it receives an error.
func (client Client) ArchiveSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ArchiveResponder handles the response to the Archive request. The method always
// closes the http.Response Body.
func (client Client) ArchiveResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// AvailableScopes get Available Scopes for `Reservation`.
// Parameters:
// reservationOrderID - order Id of the reservation
// reservationID - id of the Reservation Item
// body - parameter for listing the available scopes
func (client Client) AvailableScopes(ctx context.Context, reservationOrderID string, reservationID string, body AvailableScopeRequest) (result ReservationAvailableScopesFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.AvailableScopes")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.AvailableScopesPreparer(ctx, reservationOrderID, reservationID, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservations.Client", "AvailableScopes", nil, "Failure preparing request")
		return
	}

	result, err = client.AvailableScopesSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservations.Client", "AvailableScopes", nil, "Failure sending request")
		return
	}

	return
}

// AvailableScopesPreparer prepares the AvailableScopes request.
func (client Client) AvailableScopesPreparer(ctx context.Context, reservationOrderID string, reservationID string, body AvailableScopeRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"reservationId":      autorest.Encode("path", reservationID),
		"reservationOrderId": autorest.Encode("path", reservationOrderID),
	}

	const APIVersion = "2019-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Capacity/reservationOrders/{reservationOrderId}/reservations/{reservationId}/availableScopes", pathParameters),
		autorest.WithJSON(body),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// AvailableScopesSender sends the AvailableScopes request. The method will close the
// http.Response Body if it receives an error.
func (client Client) AvailableScopesSender(req *http.Request) (future ReservationAvailableScopesFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client Client) (p Properties, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "reservations.ReservationAvailableScopesFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("reservations.ReservationAvailableScopesFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		if p.Response.Response, err = future.GetResult(sender); err == nil && p.Response.Response.StatusCode != http.StatusNoContent {
			p, err = client.AvailableScopesResponder(p.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "reservations.ReservationAvailableScopesFuture", "Result", p.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// AvailableScopesResponder handles the response to the AvailableScopes request. The method always
// closes the http.Response Body.
func (client Client) AvailableScopesResponder(resp *http.Response) (result Properties, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get get specific `Reservation` details.
// Parameters:
// reservationID - id of the Reservation Item
// reservationOrderID - order Id of the reservation
// expand - supported value of this query is renewProperties
func (client Client) Get(ctx context.Context, reservationID string, reservationOrderID string, expand string) (result Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, reservationID, reservationOrderID, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservations.Client", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "reservations.Client", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservations.Client", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client Client) GetPreparer(ctx context.Context, reservationID string, reservationOrderID string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"reservationId":      autorest.Encode("path", reservationID),
		"reservationOrderId": autorest.Encode("path", reservationOrderID),
	}

	const APIVersion = "2019-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Capacity/reservationOrders/{reservationOrderId}/reservations/{reservationId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client Client) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client Client) GetResponder(resp *http.Response) (result Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list `Reservation`s within a single `ReservationOrder`.
// Parameters:
// reservationOrderID - order Id of the reservation
func (client Client) List(ctx context.Context, reservationOrderID string) (result ListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.List")
		defer func() {
			sc := -1
			if result.l.Response.Response != nil {
				sc = result.l.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, reservationOrderID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservations.Client", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.l.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "reservations.Client", "List", resp, "Failure sending request")
		return
	}

	result.l, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservations.Client", "List", resp, "Failure responding to request")
		return
	}
	if result.l.hasNextLink() && result.l.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client Client) ListPreparer(ctx context.Context, reservationOrderID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"reservationOrderId": autorest.Encode("path", reservationOrderID),
	}

	const APIVersion = "2019-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Capacity/reservationOrders/{reservationOrderId}/reservations", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client Client) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client Client) ListResponder(resp *http.Response) (result List, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client Client) listNextResults(ctx context.Context, lastResults List) (result List, err error) {
	req, err := lastResults.listPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "reservations.Client", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "reservations.Client", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservations.Client", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client Client) ListComplete(ctx context.Context, reservationOrderID string) (result ListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, reservationOrderID)
	return
}

// ListRevisions list of all the revisions for the `Reservation`.
// Parameters:
// reservationID - id of the Reservation Item
// reservationOrderID - order Id of the reservation
func (client Client) ListRevisions(ctx context.Context, reservationID string, reservationOrderID string) (result ListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.ListRevisions")
		defer func() {
			sc := -1
			if result.l.Response.Response != nil {
				sc = result.l.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listRevisionsNextResults
	req, err := client.ListRevisionsPreparer(ctx, reservationID, reservationOrderID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservations.Client", "ListRevisions", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListRevisionsSender(req)
	if err != nil {
		result.l.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "reservations.Client", "ListRevisions", resp, "Failure sending request")
		return
	}

	result.l, err = client.ListRevisionsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservations.Client", "ListRevisions", resp, "Failure responding to request")
		return
	}
	if result.l.hasNextLink() && result.l.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListRevisionsPreparer prepares the ListRevisions request.
func (client Client) ListRevisionsPreparer(ctx context.Context, reservationID string, reservationOrderID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"reservationId":      autorest.Encode("path", reservationID),
		"reservationOrderId": autorest.Encode("path", reservationOrderID),
	}

	const APIVersion = "2019-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Capacity/reservationOrders/{reservationOrderId}/reservations/{reservationId}/revisions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListRevisionsSender sends the ListRevisions request. The method will close the
// http.Response Body if it receives an error.
func (client Client) ListRevisionsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListRevisionsResponder handles the response to the ListRevisions request. The method always
// closes the http.Response Body.
func (client Client) ListRevisionsResponder(resp *http.Response) (result List, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listRevisionsNextResults retrieves the next set of results, if any.
func (client Client) listRevisionsNextResults(ctx context.Context, lastResults List) (result List, err error) {
	req, err := lastResults.listPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "reservations.Client", "listRevisionsNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListRevisionsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "reservations.Client", "listRevisionsNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListRevisionsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservations.Client", "listRevisionsNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListRevisionsComplete enumerates all values, automatically crossing page boundaries as required.
func (client Client) ListRevisionsComplete(ctx context.Context, reservationID string, reservationOrderID string) (result ListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.ListRevisions")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListRevisions(ctx, reservationID, reservationOrderID)
	return
}

// Merge merge the specified `Reservation`s into a new `Reservation`. The two `Reservation`s being merged must have
// same properties.
// Parameters:
// reservationOrderID - order Id of the reservation
// body - information needed for commercial request for a reservation
func (client Client) Merge(ctx context.Context, reservationOrderID string, body MergeRequest) (result ReservationMergeFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.Merge")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.MergePreparer(ctx, reservationOrderID, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservations.Client", "Merge", nil, "Failure preparing request")
		return
	}

	result, err = client.MergeSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservations.Client", "Merge", nil, "Failure sending request")
		return
	}

	return
}

// MergePreparer prepares the Merge request.
func (client Client) MergePreparer(ctx context.Context, reservationOrderID string, body MergeRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"reservationOrderId": autorest.Encode("path", reservationOrderID),
	}

	const APIVersion = "2019-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Capacity/reservationOrders/{reservationOrderId}/merge", pathParameters),
		autorest.WithJSON(body),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// MergeSender sends the Merge request. The method will close the
// http.Response Body if it receives an error.
func (client Client) MergeSender(req *http.Request) (future ReservationMergeFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client Client) (lr ListResponse, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "reservations.ReservationMergeFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("reservations.ReservationMergeFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		if lr.Response.Response, err = future.GetResult(sender); err == nil && lr.Response.Response.StatusCode != http.StatusNoContent {
			lr, err = client.MergeResponder(lr.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "reservations.ReservationMergeFuture", "Result", lr.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// MergeResponder handles the response to the Merge request. The method always
// closes the http.Response Body.
func (client Client) MergeResponder(resp *http.Response) (result ListResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Split split a `Reservation` into two `Reservation`s with specified quantity distribution.
// Parameters:
// reservationOrderID - order Id of the reservation
// body - information needed to Split a reservation item
func (client Client) Split(ctx context.Context, reservationOrderID string, body SplitRequest) (result SplitFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.Split")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.SplitPreparer(ctx, reservationOrderID, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservations.Client", "Split", nil, "Failure preparing request")
		return
	}

	result, err = client.SplitSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservations.Client", "Split", nil, "Failure sending request")
		return
	}

	return
}

// SplitPreparer prepares the Split request.
func (client Client) SplitPreparer(ctx context.Context, reservationOrderID string, body SplitRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"reservationOrderId": autorest.Encode("path", reservationOrderID),
	}

	const APIVersion = "2019-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Capacity/reservationOrders/{reservationOrderId}/split", pathParameters),
		autorest.WithJSON(body),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// SplitSender sends the Split request. The method will close the
// http.Response Body if it receives an error.
func (client Client) SplitSender(req *http.Request) (future SplitFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client Client) (lr ListResponse, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "reservations.SplitFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("reservations.SplitFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		if lr.Response.Response, err = future.GetResult(sender); err == nil && lr.Response.Response.StatusCode != http.StatusNoContent {
			lr, err = client.SplitResponder(lr.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "reservations.SplitFuture", "Result", lr.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// SplitResponder handles the response to the Split request. The method always
// closes the http.Response Body.
func (client Client) SplitResponder(resp *http.Response) (result ListResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Unarchive unarchiving a `Reservation` moves it to the state it was before archiving.
// Parameters:
// reservationOrderID - order Id of the reservation
// reservationID - id of the Reservation Item
func (client Client) Unarchive(ctx context.Context, reservationOrderID string, reservationID string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.Unarchive")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UnarchivePreparer(ctx, reservationOrderID, reservationID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservations.Client", "Unarchive", nil, "Failure preparing request")
		return
	}

	resp, err := client.UnarchiveSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "reservations.Client", "Unarchive", resp, "Failure sending request")
		return
	}

	result, err = client.UnarchiveResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservations.Client", "Unarchive", resp, "Failure responding to request")
		return
	}

	return
}

// UnarchivePreparer prepares the Unarchive request.
func (client Client) UnarchivePreparer(ctx context.Context, reservationOrderID string, reservationID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"reservationId":      autorest.Encode("path", reservationID),
		"reservationOrderId": autorest.Encode("path", reservationOrderID),
	}

	const APIVersion = "2019-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Capacity/reservationOrders/{reservationOrderId}/reservations/{reservationId}/unarchive", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UnarchiveSender sends the Unarchive request. The method will close the
// http.Response Body if it receives an error.
func (client Client) UnarchiveSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// UnarchiveResponder handles the response to the Unarchive request. The method always
// closes the http.Response Body.
func (client Client) UnarchiveResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Update updates the applied scopes of the `Reservation`.
// Parameters:
// reservationOrderID - order Id of the reservation
// reservationID - id of the Reservation Item
// parameters - information needed to patch a reservation item
func (client Client) Update(ctx context.Context, reservationOrderID string, reservationID string, parameters Patch) (result ReservationUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.Update")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, reservationOrderID, reservationID, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservations.Client", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservations.Client", "Update", nil, "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client Client) UpdatePreparer(ctx context.Context, reservationOrderID string, reservationID string, parameters Patch) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"reservationId":      autorest.Encode("path", reservationID),
		"reservationOrderId": autorest.Encode("path", reservationOrderID),
	}

	const APIVersion = "2019-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Capacity/reservationOrders/{reservationOrderId}/reservations/{reservationId}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client Client) UpdateSender(req *http.Request) (future ReservationUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client Client) (r Response, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "reservations.ReservationUpdateFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("reservations.ReservationUpdateFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		if r.Response.Response, err = future.GetResult(sender); err == nil && r.Response.Response.StatusCode != http.StatusNoContent {
			r, err = client.UpdateResponder(r.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "reservations.ReservationUpdateFuture", "Result", r.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client Client) UpdateResponder(resp *http.Response) (result Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
