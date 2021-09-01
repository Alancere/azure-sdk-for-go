package qnamaker

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

// KnowledgebaseClient is the an API for QnAMaker Service
type KnowledgebaseClient struct {
	BaseClient
}

// NewKnowledgebaseClient creates an instance of the KnowledgebaseClient client.
func NewKnowledgebaseClient(endpoint string) KnowledgebaseClient {
	return KnowledgebaseClient{New(endpoint)}
}

// Create sends the create request.
// Parameters:
// createKbPayload - post body of the request.
func (client KnowledgebaseClient) Create(ctx context.Context, createKbPayload CreateKbDTO) (result Operation, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/KnowledgebaseClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: createKbPayload,
			Constraints: []validation.Constraint{{Target: "createKbPayload.Name", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "createKbPayload.Name", Name: validation.MaxLength, Rule: 100, Chain: nil},
					{Target: "createKbPayload.Name", Name: validation.MinLength, Rule: 1, Chain: nil},
				}},
				{Target: "createKbPayload.DefaultAnswerUsedForExtraction", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "createKbPayload.DefaultAnswerUsedForExtraction", Name: validation.MaxLength, Rule: 300, Chain: nil},
						{Target: "createKbPayload.DefaultAnswerUsedForExtraction", Name: validation.MinLength, Rule: 1, Chain: nil},
					}},
				{Target: "createKbPayload.Language", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "createKbPayload.Language", Name: validation.MaxLength, Rule: 100, Chain: nil},
						{Target: "createKbPayload.Language", Name: validation.MinLength, Rule: 1, Chain: nil},
					}},
				{Target: "createKbPayload.DefaultAnswer", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "createKbPayload.DefaultAnswer", Name: validation.MaxLength, Rule: 300, Chain: nil},
						{Target: "createKbPayload.DefaultAnswer", Name: validation.MinLength, Rule: 1, Chain: nil},
					}}}}}); err != nil {
		return result, validation.NewError("qnamaker.KnowledgebaseClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, createKbPayload)
	if err != nil {
		err = autorest.NewErrorWithError(err, "qnamaker.KnowledgebaseClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "qnamaker.KnowledgebaseClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "qnamaker.KnowledgebaseClient", "Create", resp, "Failure responding to request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client KnowledgebaseClient) CreatePreparer(ctx context.Context, createKbPayload CreateKbDTO) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"":         autorest.Encode("path"),
		"Endpoint": client.Endpoint,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{Endpoint}/qnamaker/v5.0-preview.1", urlParameters),
		autorest.WithPath("/knowledgebases/create"),
		autorest.WithJSON(createKbPayload))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client KnowledgebaseClient) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client KnowledgebaseClient) CreateResponder(resp *http.Response) (result Operation, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete sends the delete request.
// Parameters:
// kbID - knowledgebase id.
func (client KnowledgebaseClient) Delete(ctx context.Context, kbID string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/KnowledgebaseClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, kbID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "qnamaker.KnowledgebaseClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "qnamaker.KnowledgebaseClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "qnamaker.KnowledgebaseClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client KnowledgebaseClient) DeletePreparer(ctx context.Context, kbID string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"":         autorest.Encode("path"),
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"kbId": autorest.Encode("path", kbID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithCustomBaseURL("{Endpoint}/qnamaker/v5.0-preview.1", urlParameters),
		autorest.WithPathParameters("/knowledgebases/{kbId}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client KnowledgebaseClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client KnowledgebaseClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Download sends the download request.
// Parameters:
// kbID - knowledgebase id.
// environment - specifies whether environment is Test or Prod.
// source - the source property filter to apply.
// changedSince - the last changed status property filter to apply.
func (client KnowledgebaseClient) Download(ctx context.Context, kbID string, environment EnvironmentType, source string, changedSince string) (result QnADocumentsDTO, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/KnowledgebaseClient.Download")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DownloadPreparer(ctx, kbID, environment, source, changedSince)
	if err != nil {
		err = autorest.NewErrorWithError(err, "qnamaker.KnowledgebaseClient", "Download", nil, "Failure preparing request")
		return
	}

	resp, err := client.DownloadSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "qnamaker.KnowledgebaseClient", "Download", resp, "Failure sending request")
		return
	}

	result, err = client.DownloadResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "qnamaker.KnowledgebaseClient", "Download", resp, "Failure responding to request")
		return
	}

	return
}

// DownloadPreparer prepares the Download request.
func (client KnowledgebaseClient) DownloadPreparer(ctx context.Context, kbID string, environment EnvironmentType, source string, changedSince string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"":         autorest.Encode("path"),
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"environment": autorest.Encode("path", environment),
		"kbId":        autorest.Encode("path", kbID),
	}

	queryParameters := map[string]interface{}{}
	if len(source) > 0 {
		queryParameters["source"] = autorest.Encode("query", source)
	}
	if len(changedSince) > 0 {
		queryParameters["changedSince"] = autorest.Encode("query", changedSince)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{Endpoint}/qnamaker/v5.0-preview.1", urlParameters),
		autorest.WithPathParameters("/knowledgebases/{kbId}/{environment}/qna", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DownloadSender sends the Download request. The method will close the
// http.Response Body if it receives an error.
func (client KnowledgebaseClient) DownloadSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DownloadResponder handles the response to the Download request. The method always
// closes the http.Response Body.
func (client KnowledgebaseClient) DownloadResponder(resp *http.Response) (result QnADocumentsDTO, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GenerateAnswer sends the generate answer request.
// Parameters:
// kbID - knowledgebase id.
// generateAnswerPayload - post body of the request.
func (client KnowledgebaseClient) GenerateAnswer(ctx context.Context, kbID string, generateAnswerPayload QueryDTO) (result QnASearchResultList, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/KnowledgebaseClient.GenerateAnswer")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GenerateAnswerPreparer(ctx, kbID, generateAnswerPayload)
	if err != nil {
		err = autorest.NewErrorWithError(err, "qnamaker.KnowledgebaseClient", "GenerateAnswer", nil, "Failure preparing request")
		return
	}

	resp, err := client.GenerateAnswerSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "qnamaker.KnowledgebaseClient", "GenerateAnswer", resp, "Failure sending request")
		return
	}

	result, err = client.GenerateAnswerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "qnamaker.KnowledgebaseClient", "GenerateAnswer", resp, "Failure responding to request")
		return
	}

	return
}

// GenerateAnswerPreparer prepares the GenerateAnswer request.
func (client KnowledgebaseClient) GenerateAnswerPreparer(ctx context.Context, kbID string, generateAnswerPayload QueryDTO) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"":         autorest.Encode("path"),
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"kbId": autorest.Encode("path", kbID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{Endpoint}/qnamaker/v5.0-preview.1", urlParameters),
		autorest.WithPathParameters("/knowledgebases/{kbId}/generateAnswer", pathParameters),
		autorest.WithJSON(generateAnswerPayload))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GenerateAnswerSender sends the GenerateAnswer request. The method will close the
// http.Response Body if it receives an error.
func (client KnowledgebaseClient) GenerateAnswerSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GenerateAnswerResponder handles the response to the GenerateAnswer request. The method always
// closes the http.Response Body.
func (client KnowledgebaseClient) GenerateAnswerResponder(resp *http.Response) (result QnASearchResultList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetDetails sends the get details request.
// Parameters:
// kbID - knowledgebase id.
func (client KnowledgebaseClient) GetDetails(ctx context.Context, kbID string) (result KnowledgebaseDTO, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/KnowledgebaseClient.GetDetails")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetDetailsPreparer(ctx, kbID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "qnamaker.KnowledgebaseClient", "GetDetails", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetDetailsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "qnamaker.KnowledgebaseClient", "GetDetails", resp, "Failure sending request")
		return
	}

	result, err = client.GetDetailsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "qnamaker.KnowledgebaseClient", "GetDetails", resp, "Failure responding to request")
		return
	}

	return
}

// GetDetailsPreparer prepares the GetDetails request.
func (client KnowledgebaseClient) GetDetailsPreparer(ctx context.Context, kbID string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"":         autorest.Encode("path"),
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"kbId": autorest.Encode("path", kbID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{Endpoint}/qnamaker/v5.0-preview.1", urlParameters),
		autorest.WithPathParameters("/knowledgebases/{kbId}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetDetailsSender sends the GetDetails request. The method will close the
// http.Response Body if it receives an error.
func (client KnowledgebaseClient) GetDetailsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetDetailsResponder handles the response to the GetDetails request. The method always
// closes the http.Response Body.
func (client KnowledgebaseClient) GetDetailsResponder(resp *http.Response) (result KnowledgebaseDTO, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListAll sends the list all request.
func (client KnowledgebaseClient) ListAll(ctx context.Context) (result KnowledgebasesDTO, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/KnowledgebaseClient.ListAll")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListAllPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "qnamaker.KnowledgebaseClient", "ListAll", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListAllSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "qnamaker.KnowledgebaseClient", "ListAll", resp, "Failure sending request")
		return
	}

	result, err = client.ListAllResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "qnamaker.KnowledgebaseClient", "ListAll", resp, "Failure responding to request")
		return
	}

	return
}

// ListAllPreparer prepares the ListAll request.
func (client KnowledgebaseClient) ListAllPreparer(ctx context.Context) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"":         autorest.Encode("path"),
		"Endpoint": client.Endpoint,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{Endpoint}/qnamaker/v5.0-preview.1", urlParameters),
		autorest.WithPath("/knowledgebases"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListAllSender sends the ListAll request. The method will close the
// http.Response Body if it receives an error.
func (client KnowledgebaseClient) ListAllSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListAllResponder handles the response to the ListAll request. The method always
// closes the http.Response Body.
func (client KnowledgebaseClient) ListAllResponder(resp *http.Response) (result KnowledgebasesDTO, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Publish sends the publish request.
// Parameters:
// kbID - knowledgebase id.
func (client KnowledgebaseClient) Publish(ctx context.Context, kbID string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/KnowledgebaseClient.Publish")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.PublishPreparer(ctx, kbID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "qnamaker.KnowledgebaseClient", "Publish", nil, "Failure preparing request")
		return
	}

	resp, err := client.PublishSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "qnamaker.KnowledgebaseClient", "Publish", resp, "Failure sending request")
		return
	}

	result, err = client.PublishResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "qnamaker.KnowledgebaseClient", "Publish", resp, "Failure responding to request")
		return
	}

	return
}

// PublishPreparer prepares the Publish request.
func (client KnowledgebaseClient) PublishPreparer(ctx context.Context, kbID string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"":         autorest.Encode("path"),
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"kbId": autorest.Encode("path", kbID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{Endpoint}/qnamaker/v5.0-preview.1", urlParameters),
		autorest.WithPathParameters("/knowledgebases/{kbId}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PublishSender sends the Publish request. The method will close the
// http.Response Body if it receives an error.
func (client KnowledgebaseClient) PublishSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// PublishResponder handles the response to the Publish request. The method always
// closes the http.Response Body.
func (client KnowledgebaseClient) PublishResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Replace sends the replace request.
// Parameters:
// kbID - knowledgebase id.
// replaceKb - an instance of ReplaceKbDTO which contains list of qnas to be uploaded
func (client KnowledgebaseClient) Replace(ctx context.Context, kbID string, replaceKb ReplaceKbDTO) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/KnowledgebaseClient.Replace")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: replaceKb,
			Constraints: []validation.Constraint{{Target: "replaceKb.QnAList", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("qnamaker.KnowledgebaseClient", "Replace", err.Error())
	}

	req, err := client.ReplacePreparer(ctx, kbID, replaceKb)
	if err != nil {
		err = autorest.NewErrorWithError(err, "qnamaker.KnowledgebaseClient", "Replace", nil, "Failure preparing request")
		return
	}

	resp, err := client.ReplaceSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "qnamaker.KnowledgebaseClient", "Replace", resp, "Failure sending request")
		return
	}

	result, err = client.ReplaceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "qnamaker.KnowledgebaseClient", "Replace", resp, "Failure responding to request")
		return
	}

	return
}

// ReplacePreparer prepares the Replace request.
func (client KnowledgebaseClient) ReplacePreparer(ctx context.Context, kbID string, replaceKb ReplaceKbDTO) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"":         autorest.Encode("path"),
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"kbId": autorest.Encode("path", kbID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithCustomBaseURL("{Endpoint}/qnamaker/v5.0-preview.1", urlParameters),
		autorest.WithPathParameters("/knowledgebases/{kbId}", pathParameters),
		autorest.WithJSON(replaceKb))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ReplaceSender sends the Replace request. The method will close the
// http.Response Body if it receives an error.
func (client KnowledgebaseClient) ReplaceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ReplaceResponder handles the response to the Replace request. The method always
// closes the http.Response Body.
func (client KnowledgebaseClient) ReplaceResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Train sends the train request.
// Parameters:
// kbID - knowledgebase id.
// trainPayload - post body of the request.
func (client KnowledgebaseClient) Train(ctx context.Context, kbID string, trainPayload FeedbackRecordsDTO) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/KnowledgebaseClient.Train")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.TrainPreparer(ctx, kbID, trainPayload)
	if err != nil {
		err = autorest.NewErrorWithError(err, "qnamaker.KnowledgebaseClient", "Train", nil, "Failure preparing request")
		return
	}

	resp, err := client.TrainSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "qnamaker.KnowledgebaseClient", "Train", resp, "Failure sending request")
		return
	}

	result, err = client.TrainResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "qnamaker.KnowledgebaseClient", "Train", resp, "Failure responding to request")
		return
	}

	return
}

// TrainPreparer prepares the Train request.
func (client KnowledgebaseClient) TrainPreparer(ctx context.Context, kbID string, trainPayload FeedbackRecordsDTO) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"":         autorest.Encode("path"),
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"kbId": autorest.Encode("path", kbID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{Endpoint}/qnamaker/v5.0-preview.1", urlParameters),
		autorest.WithPathParameters("/knowledgebases/{kbId}/train", pathParameters),
		autorest.WithJSON(trainPayload))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// TrainSender sends the Train request. The method will close the
// http.Response Body if it receives an error.
func (client KnowledgebaseClient) TrainSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// TrainResponder handles the response to the Train request. The method always
// closes the http.Response Body.
func (client KnowledgebaseClient) TrainResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Update sends the update request.
// Parameters:
// kbID - knowledgebase id.
// updateKb - post body of the request.
func (client KnowledgebaseClient) Update(ctx context.Context, kbID string, updateKb UpdateKbOperationDTO) (result Operation, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/KnowledgebaseClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, kbID, updateKb)
	if err != nil {
		err = autorest.NewErrorWithError(err, "qnamaker.KnowledgebaseClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "qnamaker.KnowledgebaseClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "qnamaker.KnowledgebaseClient", "Update", resp, "Failure responding to request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client KnowledgebaseClient) UpdatePreparer(ctx context.Context, kbID string, updateKb UpdateKbOperationDTO) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"":         autorest.Encode("path"),
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"kbId": autorest.Encode("path", kbID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithCustomBaseURL("{Endpoint}/qnamaker/v5.0-preview.1", urlParameters),
		autorest.WithPathParameters("/knowledgebases/{kbId}", pathParameters),
		autorest.WithJSON(updateKb))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client KnowledgebaseClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client KnowledgebaseClient) UpdateResponder(resp *http.Response) (result Operation, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
