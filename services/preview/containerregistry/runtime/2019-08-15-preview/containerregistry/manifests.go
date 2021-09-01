package containerregistry

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

// ManifestsClient is the metadata API definition for the Azure Container Registry runtime
type ManifestsClient struct {
	BaseClient
}

// NewManifestsClient creates an instance of the ManifestsClient client.
func NewManifestsClient(loginURI string) ManifestsClient {
	return ManifestsClient{New(loginURI)}
}

// Create put the manifest identified by `name` and `reference` where `reference` can be a tag or digest.
// Parameters:
// name - name of the image (including the namespace)
// reference - a tag or a digest, pointing to a specific image
// payload - manifest body, can take v1 or v2 values depending on accept header
func (client ManifestsClient) Create(ctx context.Context, name string, reference string, payload Manifest) (result SetObject, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ManifestsClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreatePreparer(ctx, name, reference, payload)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.ManifestsClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerregistry.ManifestsClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.ManifestsClient", "Create", resp, "Failure responding to request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client ManifestsClient) CreatePreparer(ctx context.Context, name string, reference string, payload Manifest) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"":    autorest.Encode("path"),
		"url": client.LoginURI,
	}

	pathParameters := map[string]interface{}{
		"name":      autorest.Encode("path", name),
		"reference": autorest.Encode("path", reference),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/vnd.docker.distribution.manifest.v2+json"),
		autorest.AsPut(),
		autorest.WithCustomBaseURL("{url}", urlParameters),
		autorest.WithPathParameters("/v2/{name}/manifests/{reference}", pathParameters),
		autorest.WithJSON(payload))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client ManifestsClient) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client ManifestsClient) CreateResponder(resp *http.Response) (result SetObject, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete the manifest identified by `name` and `reference`. Note that a manifest can _only_ be deleted by
// `digest`.
// Parameters:
// name - name of the image (including the namespace)
// reference - a tag or a digest, pointing to a specific image
func (client ManifestsClient) Delete(ctx context.Context, name string, reference string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ManifestsClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, name, reference)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.ManifestsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "containerregistry.ManifestsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.ManifestsClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ManifestsClient) DeletePreparer(ctx context.Context, name string, reference string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"":    autorest.Encode("path"),
		"url": client.LoginURI,
	}

	pathParameters := map[string]interface{}{
		"name":      autorest.Encode("path", name),
		"reference": autorest.Encode("path", reference),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithCustomBaseURL("{url}", urlParameters),
		autorest.WithPathParameters("/v2/{name}/manifests/{reference}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ManifestsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ManifestsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get the manifest identified by `name` and `reference` where `reference` can be a tag or digest.
// Parameters:
// name - name of the image (including the namespace)
// reference - a tag or a digest, pointing to a specific image
// accept - accept header string delimited by comma. For example,
// application/vnd.docker.distribution.manifest.v2+json
func (client ManifestsClient) Get(ctx context.Context, name string, reference string, accept string) (result ManifestWrapper, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ManifestsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, name, reference, accept)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.ManifestsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerregistry.ManifestsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.ManifestsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ManifestsClient) GetPreparer(ctx context.Context, name string, reference string, accept string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"":    autorest.Encode("path"),
		"url": client.LoginURI,
	}

	pathParameters := map[string]interface{}{
		"name":      autorest.Encode("path", name),
		"reference": autorest.Encode("path", reference),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{url}", urlParameters),
		autorest.WithPathParameters("/v2/{name}/manifests/{reference}", pathParameters))
	if len(accept) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("accept", autorest.String(accept)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ManifestsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ManifestsClient) GetResponder(resp *http.Response) (result ManifestWrapper, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetAttributes get manifest attributes
// Parameters:
// name - name of the image (including the namespace)
// reference - a tag or a digest, pointing to a specific image
func (client ManifestsClient) GetAttributes(ctx context.Context, name string, reference string) (result ManifestAttributes, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ManifestsClient.GetAttributes")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetAttributesPreparer(ctx, name, reference)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.ManifestsClient", "GetAttributes", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetAttributesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerregistry.ManifestsClient", "GetAttributes", resp, "Failure sending request")
		return
	}

	result, err = client.GetAttributesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.ManifestsClient", "GetAttributes", resp, "Failure responding to request")
		return
	}

	return
}

// GetAttributesPreparer prepares the GetAttributes request.
func (client ManifestsClient) GetAttributesPreparer(ctx context.Context, name string, reference string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"":    autorest.Encode("path"),
		"url": client.LoginURI,
	}

	pathParameters := map[string]interface{}{
		"name":      autorest.Encode("path", name),
		"reference": autorest.Encode("path", reference),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{url}", urlParameters),
		autorest.WithPathParameters("/acr/v1/{name}/_manifests/{reference}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetAttributesSender sends the GetAttributes request. The method will close the
// http.Response Body if it receives an error.
func (client ManifestsClient) GetAttributesSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetAttributesResponder handles the response to the GetAttributes request. The method always
// closes the http.Response Body.
func (client ManifestsClient) GetAttributesResponder(resp *http.Response) (result ManifestAttributes, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetList list manifests of a repository
// Parameters:
// name - name of the image (including the namespace)
// last - query parameter for the last item in previous query. Result set will include values lexically after
// last.
// n - query parameter for max number of items
// orderby - orderby query parameter
func (client ManifestsClient) GetList(ctx context.Context, name string, last string, n *int32, orderby string) (result AcrManifests, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ManifestsClient.GetList")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetListPreparer(ctx, name, last, n, orderby)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.ManifestsClient", "GetList", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerregistry.ManifestsClient", "GetList", resp, "Failure sending request")
		return
	}

	result, err = client.GetListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.ManifestsClient", "GetList", resp, "Failure responding to request")
		return
	}

	return
}

// GetListPreparer prepares the GetList request.
func (client ManifestsClient) GetListPreparer(ctx context.Context, name string, last string, n *int32, orderby string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"":    autorest.Encode("path"),
		"url": client.LoginURI,
	}

	pathParameters := map[string]interface{}{
		"name": autorest.Encode("path", name),
	}

	queryParameters := map[string]interface{}{}
	if len(last) > 0 {
		queryParameters["last"] = autorest.Encode("query", last)
	}
	if n != nil {
		queryParameters["n"] = autorest.Encode("query", *n)
	}
	if len(orderby) > 0 {
		queryParameters["orderby"] = autorest.Encode("query", orderby)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{url}", urlParameters),
		autorest.WithPathParameters("/acr/v1/{name}/_manifests", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetListSender sends the GetList request. The method will close the
// http.Response Body if it receives an error.
func (client ManifestsClient) GetListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetListResponder handles the response to the GetList request. The method always
// closes the http.Response Body.
func (client ManifestsClient) GetListResponder(resp *http.Response) (result AcrManifests, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// UpdateAttributes update attributes of a manifest
// Parameters:
// name - name of the image (including the namespace)
// reference - a tag or a digest, pointing to a specific image
// value - repository attribute value
func (client ManifestsClient) UpdateAttributes(ctx context.Context, name string, reference string, value *ChangeableAttributes) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ManifestsClient.UpdateAttributes")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdateAttributesPreparer(ctx, name, reference, value)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.ManifestsClient", "UpdateAttributes", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateAttributesSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "containerregistry.ManifestsClient", "UpdateAttributes", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateAttributesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.ManifestsClient", "UpdateAttributes", resp, "Failure responding to request")
		return
	}

	return
}

// UpdateAttributesPreparer prepares the UpdateAttributes request.
func (client ManifestsClient) UpdateAttributesPreparer(ctx context.Context, name string, reference string, value *ChangeableAttributes) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"":    autorest.Encode("path"),
		"url": client.LoginURI,
	}

	pathParameters := map[string]interface{}{
		"name":      autorest.Encode("path", name),
		"reference": autorest.Encode("path", reference),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithCustomBaseURL("{url}", urlParameters),
		autorest.WithPathParameters("/acr/v1/{name}/_manifests/{reference}", pathParameters))
	if value != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(value))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateAttributesSender sends the UpdateAttributes request. The method will close the
// http.Response Body if it receives an error.
func (client ManifestsClient) UpdateAttributesSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// UpdateAttributesResponder handles the response to the UpdateAttributes request. The method always
// closes the http.Response Body.
func (client ManifestsClient) UpdateAttributesResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}
