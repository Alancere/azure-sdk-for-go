package frontdoor

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

// EndpointsClient is the frontDoor Client
type EndpointsClient struct {
	BaseClient
}

// NewEndpointsClient creates an instance of the EndpointsClient client.
func NewEndpointsClient(subscriptionID string) EndpointsClient {
	return NewEndpointsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewEndpointsClientWithBaseURI creates an instance of the EndpointsClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewEndpointsClientWithBaseURI(baseURI string, subscriptionID string) EndpointsClient {
	return EndpointsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// PurgeContent removes a content from Front Door.
// Parameters:
// resourceGroupName - name of the Resource group within the Azure subscription.
// frontDoorName - name of the Front Door which is globally unique.
// contentFilePaths - the path to the content to be purged. Path can be a full URL, e.g. '/pictures/city.png'
// which removes a single file, or a directory with a wildcard, e.g. '/pictures/*' which removes all folders
// and files in the directory.
func (client EndpointsClient) PurgeContent(ctx context.Context, resourceGroupName string, frontDoorName string, contentFilePaths PurgeParameters) (result EndpointsPurgeContentFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EndpointsClient.PurgeContent")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 80, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_\-\(\)\.]*[^\.]$`, Chain: nil}}},
		{TargetValue: frontDoorName,
			Constraints: []validation.Constraint{{Target: "frontDoorName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "frontDoorName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "frontDoorName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]+([-a-zA-Z0-9]?[a-zA-Z0-9])*$`, Chain: nil}}},
		{TargetValue: contentFilePaths,
			Constraints: []validation.Constraint{{Target: "contentFilePaths.ContentPaths", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("frontdoor.EndpointsClient", "PurgeContent", err.Error())
	}

	req, err := client.PurgeContentPreparer(ctx, resourceGroupName, frontDoorName, contentFilePaths)
	if err != nil {
		err = autorest.NewErrorWithError(err, "frontdoor.EndpointsClient", "PurgeContent", nil, "Failure preparing request")
		return
	}

	result, err = client.PurgeContentSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "frontdoor.EndpointsClient", "PurgeContent", result.Response(), "Failure sending request")
		return
	}

	return
}

// PurgeContentPreparer prepares the PurgeContent request.
func (client EndpointsClient) PurgeContentPreparer(ctx context.Context, resourceGroupName string, frontDoorName string, contentFilePaths PurgeParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"frontDoorName":     autorest.Encode("path", frontDoorName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoors/{frontDoorName}/purge", pathParameters),
		autorest.WithJSON(contentFilePaths),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PurgeContentSender sends the PurgeContent request. The method will close the
// http.Response Body if it receives an error.
func (client EndpointsClient) PurgeContentSender(req *http.Request) (future EndpointsPurgeContentFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// PurgeContentResponder handles the response to the PurgeContent request. The method always
// closes the http.Response Body.
func (client EndpointsClient) PurgeContentResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}
