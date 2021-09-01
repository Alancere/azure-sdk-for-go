package authoring

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
	"github.com/gofrs/uuid"
	"net/http"
)

// ExamplesClient is the client for the Examples methods of the Authoring service.
type ExamplesClient struct {
	BaseClient
}

// NewExamplesClient creates an instance of the ExamplesClient client.
func NewExamplesClient(endpoint string) ExamplesClient {
	return ExamplesClient{New(endpoint)}
}

// Add adds a labeled example utterance in a version of the application.
// Parameters:
// appID - the application ID.
// versionID - the version ID.
// exampleLabelObject - a labeled example utterance with the expected intent and entities.
// enableNestedChildren - toggles nested/flat format
func (client ExamplesClient) Add(ctx context.Context, appID uuid.UUID, versionID string, exampleLabelObject ExampleLabelObject, enableNestedChildren *bool) (result LabelExampleResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExamplesClient.Add")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.AddPreparer(ctx, appID, versionID, exampleLabelObject, enableNestedChildren)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authoring.ExamplesClient", "Add", nil, "Failure preparing request")
		return
	}

	resp, err := client.AddSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "authoring.ExamplesClient", "Add", resp, "Failure sending request")
		return
	}

	result, err = client.AddResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authoring.ExamplesClient", "Add", resp, "Failure responding to request")
		return
	}

	return
}

// AddPreparer prepares the Add request.
func (client ExamplesClient) AddPreparer(ctx context.Context, appID uuid.UUID, versionID string, exampleLabelObject ExampleLabelObject, enableNestedChildren *bool) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"":         autorest.Encode("path"),
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"appId":     autorest.Encode("path", appID),
		"versionId": autorest.Encode("path", versionID),
	}

	queryParameters := map[string]interface{}{}
	if enableNestedChildren != nil {
		queryParameters["enableNestedChildren"] = autorest.Encode("query", *enableNestedChildren)
	} else {
		queryParameters["enableNestedChildren"] = autorest.Encode("query", false)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{Endpoint}/luis/authoring/v3.0-preview", urlParameters),
		autorest.WithPathParameters("/apps/{appId}/versions/{versionId}/example", pathParameters),
		autorest.WithJSON(exampleLabelObject),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// AddSender sends the Add request. The method will close the
// http.Response Body if it receives an error.
func (client ExamplesClient) AddSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// AddResponder handles the response to the Add request. The method always
// closes the http.Response Body.
func (client ExamplesClient) AddResponder(resp *http.Response) (result LabelExampleResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Batch adds a batch of labeled example utterances to a version of the application.
// Parameters:
// appID - the application ID.
// versionID - the version ID.
// exampleLabelObjectArray - array of example utterances.
// enableNestedChildren - toggles nested/flat format
func (client ExamplesClient) Batch(ctx context.Context, appID uuid.UUID, versionID string, exampleLabelObjectArray []ExampleLabelObject, enableNestedChildren *bool) (result ListBatchLabelExample, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExamplesClient.Batch")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: exampleLabelObjectArray,
			Constraints: []validation.Constraint{{Target: "exampleLabelObjectArray", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("authoring.ExamplesClient", "Batch", err.Error())
	}

	req, err := client.BatchPreparer(ctx, appID, versionID, exampleLabelObjectArray, enableNestedChildren)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authoring.ExamplesClient", "Batch", nil, "Failure preparing request")
		return
	}

	resp, err := client.BatchSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "authoring.ExamplesClient", "Batch", resp, "Failure sending request")
		return
	}

	result, err = client.BatchResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authoring.ExamplesClient", "Batch", resp, "Failure responding to request")
		return
	}

	return
}

// BatchPreparer prepares the Batch request.
func (client ExamplesClient) BatchPreparer(ctx context.Context, appID uuid.UUID, versionID string, exampleLabelObjectArray []ExampleLabelObject, enableNestedChildren *bool) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"":         autorest.Encode("path"),
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"appId":     autorest.Encode("path", appID),
		"versionId": autorest.Encode("path", versionID),
	}

	queryParameters := map[string]interface{}{}
	if enableNestedChildren != nil {
		queryParameters["enableNestedChildren"] = autorest.Encode("query", *enableNestedChildren)
	} else {
		queryParameters["enableNestedChildren"] = autorest.Encode("query", false)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{Endpoint}/luis/authoring/v3.0-preview", urlParameters),
		autorest.WithPathParameters("/apps/{appId}/versions/{versionId}/examples", pathParameters),
		autorest.WithJSON(exampleLabelObjectArray),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// BatchSender sends the Batch request. The method will close the
// http.Response Body if it receives an error.
func (client ExamplesClient) BatchSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// BatchResponder handles the response to the Batch request. The method always
// closes the http.Response Body.
func (client ExamplesClient) BatchResponder(resp *http.Response) (result ListBatchLabelExample, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusMultiStatus),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the labeled example utterances with the specified ID from a version of the application.
// Parameters:
// appID - the application ID.
// versionID - the version ID.
// exampleID - the example ID.
func (client ExamplesClient) Delete(ctx context.Context, appID uuid.UUID, versionID string, exampleID int32) (result OperationStatus, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExamplesClient.Delete")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, appID, versionID, exampleID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authoring.ExamplesClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "authoring.ExamplesClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authoring.ExamplesClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ExamplesClient) DeletePreparer(ctx context.Context, appID uuid.UUID, versionID string, exampleID int32) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"":         autorest.Encode("path"),
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"appId":     autorest.Encode("path", appID),
		"exampleId": autorest.Encode("path", exampleID),
		"versionId": autorest.Encode("path", versionID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithCustomBaseURL("{Endpoint}/luis/authoring/v3.0-preview", urlParameters),
		autorest.WithPathParameters("/apps/{appId}/versions/{versionId}/examples/{exampleId}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ExamplesClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ExamplesClient) DeleteResponder(resp *http.Response) (result OperationStatus, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List returns example utterances to be reviewed from a version of the application.
// Parameters:
// appID - the application ID.
// versionID - the version ID.
// skip - the number of entries to skip. Default value is 0.
// take - the number of entries to return. Maximum page size is 500. Default is 100.
// enableNestedChildren - toggles nested/flat format
func (client ExamplesClient) List(ctx context.Context, appID uuid.UUID, versionID string, skip *int32, take *int32, enableNestedChildren *bool) (result ListLabeledUtterance, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExamplesClient.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: skip,
			Constraints: []validation.Constraint{{Target: "skip", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "skip", Name: validation.InclusiveMinimum, Rule: int64(0), Chain: nil}}}}},
		{TargetValue: take,
			Constraints: []validation.Constraint{{Target: "take", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "take", Name: validation.InclusiveMaximum, Rule: int64(500), Chain: nil},
					{Target: "take", Name: validation.InclusiveMinimum, Rule: int64(0), Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("authoring.ExamplesClient", "List", err.Error())
	}

	req, err := client.ListPreparer(ctx, appID, versionID, skip, take, enableNestedChildren)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authoring.ExamplesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "authoring.ExamplesClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authoring.ExamplesClient", "List", resp, "Failure responding to request")
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client ExamplesClient) ListPreparer(ctx context.Context, appID uuid.UUID, versionID string, skip *int32, take *int32, enableNestedChildren *bool) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"":         autorest.Encode("path"),
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"appId":     autorest.Encode("path", appID),
		"versionId": autorest.Encode("path", versionID),
	}

	queryParameters := map[string]interface{}{}
	if skip != nil {
		queryParameters["skip"] = autorest.Encode("query", *skip)
	} else {
		queryParameters["skip"] = autorest.Encode("query", 0)
	}
	if take != nil {
		queryParameters["take"] = autorest.Encode("query", *take)
	} else {
		queryParameters["take"] = autorest.Encode("query", 100)
	}
	if enableNestedChildren != nil {
		queryParameters["enableNestedChildren"] = autorest.Encode("query", *enableNestedChildren)
	} else {
		queryParameters["enableNestedChildren"] = autorest.Encode("query", false)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{Endpoint}/luis/authoring/v3.0-preview", urlParameters),
		autorest.WithPathParameters("/apps/{appId}/versions/{versionId}/examples", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ExamplesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ExamplesClient) ListResponder(resp *http.Response) (result ListLabeledUtterance, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
