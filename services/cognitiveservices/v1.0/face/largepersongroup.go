package face

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

// LargePersonGroupClient is the an API for face detection, verification, and identification.
type LargePersonGroupClient struct {
	BaseClient
}

// NewLargePersonGroupClient creates an instance of the LargePersonGroupClient client.
func NewLargePersonGroupClient(endpoint string) LargePersonGroupClient {
	return LargePersonGroupClient{New(endpoint)}
}

// Create create a new large person group with user-specified largePersonGroupId, name, an optional userData and
// recognitionModel.
// <br /> A large person group is the container of the uploaded person data, including face recognition feature, and up
// to 1,000,000
// people.
// <br /> After creation, use [LargePersonGroup Person -
// Create](https://docs.microsoft.com/rest/api/cognitiveservices/face/largepersongroupperson/create) to add person into
// the group, and call [LargePersonGroup -
// Train](https://docs.microsoft.com/rest/api/cognitiveservices/face/largepersongroup/train) to get this group ready
// for [Face - Identify](https://docs.microsoft.com/rest/api/cognitiveservices/face/face/identify).
// <br /> No image will be stored. Only the person's extracted face features and userData will be stored on server
// until [LargePersonGroup Person -
// Delete](https://docs.microsoft.com/rest/api/cognitiveservices/face/largepersongroupperson/delete) or
// [LargePersonGroup - Delete](https://docs.microsoft.com/rest/api/cognitiveservices/face/largepersongroup/delete) is
// called.
// <br/>'recognitionModel' should be specified to associate with this large person group. The default value for
// 'recognitionModel' is 'recognition_01', if the latest model needed, please explicitly specify the model you need in
// this parameter. New faces that are added to an existing large person group will use the recognition model that's
// already associated with the collection. Existing face features in a large person group can't be updated to features
// extracted by another version of recognition model.
// * 'recognition_01': The default recognition model for [LargePersonGroup -
// Create](https://docs.microsoft.com/rest/api/cognitiveservices/face/largepersongroup/create). All those large person
// groups created before 2019 March are bonded with this recognition model.
// * 'recognition_02': Recognition model released in 2019 March.
// * 'recognition_03': Recognition model released in 2020 May. 'recognition_03' is recommended since its overall
// accuracy is improved compared with 'recognition_01' and 'recognition_02'.
//
// Large person group quota:
// * Free-tier subscription quota: 1,000 large person groups.
// * S0-tier subscription quota: 1,000,000 large person groups.
// Parameters:
// largePersonGroupID - id referencing a particular large person group.
// body - request body for creating new large person group.
func (client LargePersonGroupClient) Create(ctx context.Context, largePersonGroupID string, body MetaDataContract) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LargePersonGroupClient.Create")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: largePersonGroupID,
			Constraints: []validation.Constraint{{Target: "largePersonGroupID", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "largePersonGroupID", Name: validation.Pattern, Rule: `^[a-z0-9-_]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("face.LargePersonGroupClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, largePersonGroupID, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.LargePersonGroupClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "face.LargePersonGroupClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.LargePersonGroupClient", "Create", resp, "Failure responding to request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client LargePersonGroupClient) CreatePreparer(ctx context.Context, largePersonGroupID string, body MetaDataContract) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"":         autorest.Encode("path"),
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"largePersonGroupId": autorest.Encode("path", largePersonGroupID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithCustomBaseURL("{Endpoint}/face/v1.0", urlParameters),
		autorest.WithPathParameters("/largepersongroups/{largePersonGroupId}", pathParameters),
		autorest.WithJSON(body))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client LargePersonGroupClient) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client LargePersonGroupClient) CreateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Delete delete an existing large person group. Persisted face features of all people in the large person group will
// also be deleted.
// Parameters:
// largePersonGroupID - id referencing a particular large person group.
func (client LargePersonGroupClient) Delete(ctx context.Context, largePersonGroupID string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LargePersonGroupClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: largePersonGroupID,
			Constraints: []validation.Constraint{{Target: "largePersonGroupID", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "largePersonGroupID", Name: validation.Pattern, Rule: `^[a-z0-9-_]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("face.LargePersonGroupClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, largePersonGroupID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.LargePersonGroupClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "face.LargePersonGroupClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.LargePersonGroupClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client LargePersonGroupClient) DeletePreparer(ctx context.Context, largePersonGroupID string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"":         autorest.Encode("path"),
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"largePersonGroupId": autorest.Encode("path", largePersonGroupID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithCustomBaseURL("{Endpoint}/face/v1.0", urlParameters),
		autorest.WithPathParameters("/largepersongroups/{largePersonGroupId}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client LargePersonGroupClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client LargePersonGroupClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get retrieve the information of a large person group, including its name, userData and recognitionModel. This API
// returns large person group information only, use [LargePersonGroup Person -
// List](https://docs.microsoft.com/rest/api/cognitiveservices/face/largepersongroupperson/list) instead to retrieve
// person information under the large person group.
// Parameters:
// largePersonGroupID - id referencing a particular large person group.
// returnRecognitionModel - a value indicating whether the operation should return 'recognitionModel' in
// response.
func (client LargePersonGroupClient) Get(ctx context.Context, largePersonGroupID string, returnRecognitionModel *bool) (result LargePersonGroup, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LargePersonGroupClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: largePersonGroupID,
			Constraints: []validation.Constraint{{Target: "largePersonGroupID", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "largePersonGroupID", Name: validation.Pattern, Rule: `^[a-z0-9-_]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("face.LargePersonGroupClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, largePersonGroupID, returnRecognitionModel)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.LargePersonGroupClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "face.LargePersonGroupClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.LargePersonGroupClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client LargePersonGroupClient) GetPreparer(ctx context.Context, largePersonGroupID string, returnRecognitionModel *bool) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"":         autorest.Encode("path"),
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"largePersonGroupId": autorest.Encode("path", largePersonGroupID),
	}

	queryParameters := map[string]interface{}{}
	if returnRecognitionModel != nil {
		queryParameters["returnRecognitionModel"] = autorest.Encode("query", *returnRecognitionModel)
	} else {
		queryParameters["returnRecognitionModel"] = autorest.Encode("query", false)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{Endpoint}/face/v1.0", urlParameters),
		autorest.WithPathParameters("/largepersongroups/{largePersonGroupId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client LargePersonGroupClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client LargePersonGroupClient) GetResponder(resp *http.Response) (result LargePersonGroup, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetTrainingStatus retrieve the training status of a large person group (completed or ongoing).
// Parameters:
// largePersonGroupID - id referencing a particular large person group.
func (client LargePersonGroupClient) GetTrainingStatus(ctx context.Context, largePersonGroupID string) (result TrainingStatus, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LargePersonGroupClient.GetTrainingStatus")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: largePersonGroupID,
			Constraints: []validation.Constraint{{Target: "largePersonGroupID", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "largePersonGroupID", Name: validation.Pattern, Rule: `^[a-z0-9-_]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("face.LargePersonGroupClient", "GetTrainingStatus", err.Error())
	}

	req, err := client.GetTrainingStatusPreparer(ctx, largePersonGroupID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.LargePersonGroupClient", "GetTrainingStatus", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetTrainingStatusSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "face.LargePersonGroupClient", "GetTrainingStatus", resp, "Failure sending request")
		return
	}

	result, err = client.GetTrainingStatusResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.LargePersonGroupClient", "GetTrainingStatus", resp, "Failure responding to request")
		return
	}

	return
}

// GetTrainingStatusPreparer prepares the GetTrainingStatus request.
func (client LargePersonGroupClient) GetTrainingStatusPreparer(ctx context.Context, largePersonGroupID string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"":         autorest.Encode("path"),
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"largePersonGroupId": autorest.Encode("path", largePersonGroupID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{Endpoint}/face/v1.0", urlParameters),
		autorest.WithPathParameters("/largepersongroups/{largePersonGroupId}/training", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetTrainingStatusSender sends the GetTrainingStatus request. The method will close the
// http.Response Body if it receives an error.
func (client LargePersonGroupClient) GetTrainingStatusSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetTrainingStatusResponder handles the response to the GetTrainingStatus request. The method always
// closes the http.Response Body.
func (client LargePersonGroupClient) GetTrainingStatusResponder(resp *http.Response) (result TrainingStatus, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list all existing large person groups’ largePersonGroupId, name, userData and recognitionModel.<br />
// * Large person groups are stored in alphabetical order of largePersonGroupId.
// * "start" parameter (string, optional) is a user-provided largePersonGroupId value that returned entries have larger
// ids by string comparison. "start" set to empty to indicate return from the first item.
// * "top" parameter (int, optional) specifies the number of entries to return. A maximal of 1000 entries can be
// returned in one call. To fetch more, you can specify "start" with the last returned entry’s Id of the current call.
// <br />
// For example, total 5 large person groups: "group1", ..., "group5".
// <br /> "start=&top=" will return all 5 groups.
// <br /> "start=&top=2" will return "group1", "group2".
// <br /> "start=group2&top=3" will return "group3", "group4", "group5".
// Parameters:
// start - list large person groups from the least largePersonGroupId greater than the "start".
// top - the number of large person groups to list.
// returnRecognitionModel - a value indicating whether the operation should return 'recognitionModel' in
// response.
func (client LargePersonGroupClient) List(ctx context.Context, start string, top *int32, returnRecognitionModel *bool) (result ListLargePersonGroup, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LargePersonGroupClient.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: start,
			Constraints: []validation.Constraint{{Target: "start", Name: validation.Empty, Rule: false,
				Chain: []validation.Constraint{{Target: "start", Name: validation.MaxLength, Rule: 64, Chain: nil}}}}},
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMaximum, Rule: int64(1000), Chain: nil},
					{Target: "top", Name: validation.InclusiveMinimum, Rule: int64(1), Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("face.LargePersonGroupClient", "List", err.Error())
	}

	req, err := client.ListPreparer(ctx, start, top, returnRecognitionModel)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.LargePersonGroupClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "face.LargePersonGroupClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.LargePersonGroupClient", "List", resp, "Failure responding to request")
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client LargePersonGroupClient) ListPreparer(ctx context.Context, start string, top *int32, returnRecognitionModel *bool) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"":         autorest.Encode("path"),
		"Endpoint": client.Endpoint,
	}

	queryParameters := map[string]interface{}{}
	if len(start) > 0 {
		queryParameters["start"] = autorest.Encode("query", start)
	}
	if top != nil {
		queryParameters["top"] = autorest.Encode("query", *top)
	} else {
		queryParameters["top"] = autorest.Encode("query", 1000)
	}
	if returnRecognitionModel != nil {
		queryParameters["returnRecognitionModel"] = autorest.Encode("query", *returnRecognitionModel)
	} else {
		queryParameters["returnRecognitionModel"] = autorest.Encode("query", false)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{Endpoint}/face/v1.0", urlParameters),
		autorest.WithPath("/largepersongroups"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client LargePersonGroupClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client LargePersonGroupClient) ListResponder(resp *http.Response) (result ListLargePersonGroup, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Train queue a large person group training task, the training task may not be started immediately.
// Parameters:
// largePersonGroupID - id referencing a particular large person group.
func (client LargePersonGroupClient) Train(ctx context.Context, largePersonGroupID string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LargePersonGroupClient.Train")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: largePersonGroupID,
			Constraints: []validation.Constraint{{Target: "largePersonGroupID", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "largePersonGroupID", Name: validation.Pattern, Rule: `^[a-z0-9-_]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("face.LargePersonGroupClient", "Train", err.Error())
	}

	req, err := client.TrainPreparer(ctx, largePersonGroupID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.LargePersonGroupClient", "Train", nil, "Failure preparing request")
		return
	}

	resp, err := client.TrainSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "face.LargePersonGroupClient", "Train", resp, "Failure sending request")
		return
	}

	result, err = client.TrainResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.LargePersonGroupClient", "Train", resp, "Failure responding to request")
		return
	}

	return
}

// TrainPreparer prepares the Train request.
func (client LargePersonGroupClient) TrainPreparer(ctx context.Context, largePersonGroupID string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"":         autorest.Encode("path"),
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"largePersonGroupId": autorest.Encode("path", largePersonGroupID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{Endpoint}/face/v1.0", urlParameters),
		autorest.WithPathParameters("/largepersongroups/{largePersonGroupId}/train", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// TrainSender sends the Train request. The method will close the
// http.Response Body if it receives an error.
func (client LargePersonGroupClient) TrainSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// TrainResponder handles the response to the Train request. The method always
// closes the http.Response Body.
func (client LargePersonGroupClient) TrainResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Update update an existing large person group's display name and userData. The properties which does not appear in
// request body will not be updated.
// Parameters:
// largePersonGroupID - id referencing a particular large person group.
// body - request body for updating large person group.
func (client LargePersonGroupClient) Update(ctx context.Context, largePersonGroupID string, body NameAndUserDataContract) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LargePersonGroupClient.Update")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: largePersonGroupID,
			Constraints: []validation.Constraint{{Target: "largePersonGroupID", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "largePersonGroupID", Name: validation.Pattern, Rule: `^[a-z0-9-_]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("face.LargePersonGroupClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, largePersonGroupID, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.LargePersonGroupClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "face.LargePersonGroupClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.LargePersonGroupClient", "Update", resp, "Failure responding to request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client LargePersonGroupClient) UpdatePreparer(ctx context.Context, largePersonGroupID string, body NameAndUserDataContract) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"":         autorest.Encode("path"),
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"largePersonGroupId": autorest.Encode("path", largePersonGroupID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithCustomBaseURL("{Endpoint}/face/v1.0", urlParameters),
		autorest.WithPathParameters("/largepersongroups/{largePersonGroupId}", pathParameters),
		autorest.WithJSON(body))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client LargePersonGroupClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client LargePersonGroupClient) UpdateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}
