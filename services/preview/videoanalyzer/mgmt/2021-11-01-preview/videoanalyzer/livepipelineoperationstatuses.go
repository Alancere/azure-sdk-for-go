//Deprecated: We’re retiring the Azure Video Analyzer preview service, you're advised to transition your applications off of Video Analyzer by 01 December 2022. This SDK is no longer maintained and won’t work after the service is retired

package videoanalyzer

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

// LivePipelineOperationStatusesClient is the azure Video Analyzer provides a platform for you to build intelligent
// video applications that span the edge and the cloud
type LivePipelineOperationStatusesClient struct {
	BaseClient
}

// NewLivePipelineOperationStatusesClient creates an instance of the LivePipelineOperationStatusesClient client.
func NewLivePipelineOperationStatusesClient(subscriptionID string) LivePipelineOperationStatusesClient {
	return NewLivePipelineOperationStatusesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewLivePipelineOperationStatusesClientWithBaseURI creates an instance of the LivePipelineOperationStatusesClient
// client using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI
// (sovereign clouds, Azure stack).
func NewLivePipelineOperationStatusesClientWithBaseURI(baseURI string, subscriptionID string) LivePipelineOperationStatusesClient {
	return LivePipelineOperationStatusesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get get the operation status of a live pipeline.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// accountName - the Azure Video Analyzer account name.
// livePipelineName - live pipeline unique identifier.
// operationID - the operation ID.
func (client LivePipelineOperationStatusesClient) Get(ctx context.Context, resourceGroupName string, accountName string, livePipelineName string, operationID string) (result LivePipelineOperationStatus, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LivePipelineOperationStatusesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("videoanalyzer.LivePipelineOperationStatusesClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, accountName, livePipelineName, operationID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "videoanalyzer.LivePipelineOperationStatusesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "videoanalyzer.LivePipelineOperationStatusesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "videoanalyzer.LivePipelineOperationStatusesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client LivePipelineOperationStatusesClient) GetPreparer(ctx context.Context, resourceGroupName string, accountName string, livePipelineName string, operationID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"livePipelineName":  autorest.Encode("path", livePipelineName),
		"operationId":       autorest.Encode("path", operationID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/videoAnalyzers/{accountName}/livePipelines/{livePipelineName}/operationStatuses/{operationId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client LivePipelineOperationStatusesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client LivePipelineOperationStatusesClient) GetResponder(resp *http.Response) (result LivePipelineOperationStatus, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
