package vmwarecloudsimple

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
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// VirtualMachinesClient is the description of the new service
type VirtualMachinesClient struct {
	BaseClient
}

// NewVirtualMachinesClient creates an instance of the VirtualMachinesClient client.
func NewVirtualMachinesClient(subscriptionID string, referer string) VirtualMachinesClient {
	return NewVirtualMachinesClientWithBaseURI(DefaultBaseURI, subscriptionID, referer)
}

// NewVirtualMachinesClientWithBaseURI creates an instance of the VirtualMachinesClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewVirtualMachinesClientWithBaseURI(baseURI string, subscriptionID string, referer string) VirtualMachinesClient {
	return VirtualMachinesClient{NewWithBaseURI(baseURI, subscriptionID, referer)}
}

// CreateOrUpdate create Or Update Virtual Machine
// Parameters:
// resourceGroupName - the name of the resource group
// virtualMachineName - virtual machine name
// virtualMachineRequest - create or Update Virtual Machine request
func (client VirtualMachinesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, virtualMachineName string, virtualMachineRequest VirtualMachine) (result VirtualMachinesCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualMachinesClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: virtualMachineName,
			Constraints: []validation.Constraint{{Target: "virtualMachineName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]([-_.a-zA-Z0-9]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: virtualMachineRequest,
			Constraints: []validation.Constraint{{Target: "virtualMachineRequest.Location", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "virtualMachineRequest.Name", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "virtualMachineRequest.Name", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]([-_.a-zA-Z0-9]*[a-zA-Z0-9])?$`, Chain: nil}}},
				{Target: "virtualMachineRequest.VirtualMachineProperties", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "virtualMachineRequest.VirtualMachineProperties.AmountOfRAM", Name: validation.Null, Rule: true, Chain: nil},
						{Target: "virtualMachineRequest.VirtualMachineProperties.NumberOfCores", Name: validation.Null, Rule: true, Chain: nil},
						{Target: "virtualMachineRequest.VirtualMachineProperties.PrivateCloudID", Name: validation.Null, Rule: true, Chain: nil},
						{Target: "virtualMachineRequest.VirtualMachineProperties.ResourcePool", Name: validation.Null, Rule: false,
							Chain: []validation.Constraint{{Target: "virtualMachineRequest.VirtualMachineProperties.ResourcePool.ID", Name: validation.Null, Rule: true, Chain: nil}}},
					}}}}}); err != nil {
		return result, validation.NewError("vmwarecloudsimple.VirtualMachinesClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, virtualMachineName, virtualMachineRequest)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.VirtualMachinesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.VirtualMachinesClient", "CreateOrUpdate", nil, "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client VirtualMachinesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, virtualMachineName string, virtualMachineRequest VirtualMachine) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
		"virtualMachineName": autorest.Encode("path", virtualMachineName),
	}

	const APIVersion = "2019-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	virtualMachineRequest.ID = nil
	virtualMachineRequest.Name = nil
	virtualMachineRequest.Type = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/virtualMachines/{virtualMachineName}", pathParameters),
		autorest.WithJSON(virtualMachineRequest),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("Referer", client.Referer))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachinesClient) CreateOrUpdateSender(req *http.Request) (future VirtualMachinesCreateOrUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client VirtualMachinesClient) (VM VirtualMachine, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "vmwarecloudsimple.VirtualMachinesCreateOrUpdateFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("vmwarecloudsimple.VirtualMachinesCreateOrUpdateFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		if VM.Response.Response, err = future.GetResult(sender); err == nil && VM.Response.Response.StatusCode != http.StatusNoContent {
			VM, err = client.CreateOrUpdateResponder(VM.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "vmwarecloudsimple.VirtualMachinesCreateOrUpdateFuture", "Result", VM.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client VirtualMachinesClient) CreateOrUpdateResponder(resp *http.Response) (result VirtualMachine, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete virtual machine
// Parameters:
// resourceGroupName - the name of the resource group
// virtualMachineName - virtual machine name
func (client VirtualMachinesClient) Delete(ctx context.Context, resourceGroupName string, virtualMachineName string) (result VirtualMachinesDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualMachinesClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, virtualMachineName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.VirtualMachinesClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.VirtualMachinesClient", "Delete", nil, "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client VirtualMachinesClient) DeletePreparer(ctx context.Context, resourceGroupName string, virtualMachineName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
		"virtualMachineName": autorest.Encode("path", virtualMachineName),
	}

	const APIVersion = "2019-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/virtualMachines/{virtualMachineName}", pathParameters),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("Referer", client.Referer))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachinesClient) DeleteSender(req *http.Request) (future VirtualMachinesDeleteFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client VirtualMachinesClient) (ar autorest.Response, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "vmwarecloudsimple.VirtualMachinesDeleteFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("vmwarecloudsimple.VirtualMachinesDeleteFuture")
			return
		}
		ar.Response = future.Response()
		return
	}
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client VirtualMachinesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get virtual machine
// Parameters:
// resourceGroupName - the name of the resource group
// virtualMachineName - virtual machine name
func (client VirtualMachinesClient) Get(ctx context.Context, resourceGroupName string, virtualMachineName string) (result VirtualMachine, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualMachinesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, virtualMachineName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.VirtualMachinesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.VirtualMachinesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.VirtualMachinesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client VirtualMachinesClient) GetPreparer(ctx context.Context, resourceGroupName string, virtualMachineName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
		"virtualMachineName": autorest.Encode("path", virtualMachineName),
	}

	const APIVersion = "2019-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/virtualMachines/{virtualMachineName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachinesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client VirtualMachinesClient) GetResponder(resp *http.Response) (result VirtualMachine, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroup returns list of virtual machine within resource group
// Parameters:
// resourceGroupName - the name of the resource group
// filter - the filter to apply on the list operation
// top - the maximum number of record sets to return
// skipToken - to be used by nextLink implementation
func (client VirtualMachinesClient) ListByResourceGroup(ctx context.Context, resourceGroupName string, filter string, top *int32, skipToken string) (result VirtualMachineListResponsePage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualMachinesClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.vmlr.Response.Response != nil {
				sc = result.vmlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName, filter, top, skipToken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.VirtualMachinesClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.vmlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.VirtualMachinesClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.vmlr, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.VirtualMachinesClient", "ListByResourceGroup", resp, "Failure responding to request")
		return
	}
	if result.vmlr.hasNextLink() && result.vmlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client VirtualMachinesClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string, filter string, top *int32, skipToken string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(skipToken) > 0 {
		queryParameters["$skipToken"] = autorest.Encode("query", skipToken)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/virtualMachines", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachinesClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client VirtualMachinesClient) ListByResourceGroupResponder(resp *http.Response) (result VirtualMachineListResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client VirtualMachinesClient) listByResourceGroupNextResults(ctx context.Context, lastResults VirtualMachineListResponse) (result VirtualMachineListResponse, err error) {
	req, err := lastResults.virtualMachineListResponsePreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "vmwarecloudsimple.VirtualMachinesClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "vmwarecloudsimple.VirtualMachinesClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.VirtualMachinesClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client VirtualMachinesClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, filter string, top *int32, skipToken string) (result VirtualMachineListResponseIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualMachinesClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName, filter, top, skipToken)
	return
}

// ListBySubscription returns list virtual machine within subscription
// Parameters:
// filter - the filter to apply on the list operation
// top - the maximum number of record sets to return
// skipToken - to be used by nextLink implementation
func (client VirtualMachinesClient) ListBySubscription(ctx context.Context, filter string, top *int32, skipToken string) (result VirtualMachineListResponsePage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualMachinesClient.ListBySubscription")
		defer func() {
			sc := -1
			if result.vmlr.Response.Response != nil {
				sc = result.vmlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listBySubscriptionNextResults
	req, err := client.ListBySubscriptionPreparer(ctx, filter, top, skipToken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.VirtualMachinesClient", "ListBySubscription", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.vmlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.VirtualMachinesClient", "ListBySubscription", resp, "Failure sending request")
		return
	}

	result.vmlr, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.VirtualMachinesClient", "ListBySubscription", resp, "Failure responding to request")
		return
	}
	if result.vmlr.hasNextLink() && result.vmlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListBySubscriptionPreparer prepares the ListBySubscription request.
func (client VirtualMachinesClient) ListBySubscriptionPreparer(ctx context.Context, filter string, top *int32, skipToken string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(skipToken) > 0 {
		queryParameters["$skipToken"] = autorest.Encode("query", skipToken)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/virtualMachines", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListBySubscriptionSender sends the ListBySubscription request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachinesClient) ListBySubscriptionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListBySubscriptionResponder handles the response to the ListBySubscription request. The method always
// closes the http.Response Body.
func (client VirtualMachinesClient) ListBySubscriptionResponder(resp *http.Response) (result VirtualMachineListResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listBySubscriptionNextResults retrieves the next set of results, if any.
func (client VirtualMachinesClient) listBySubscriptionNextResults(ctx context.Context, lastResults VirtualMachineListResponse) (result VirtualMachineListResponse, err error) {
	req, err := lastResults.virtualMachineListResponsePreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "vmwarecloudsimple.VirtualMachinesClient", "listBySubscriptionNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "vmwarecloudsimple.VirtualMachinesClient", "listBySubscriptionNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.VirtualMachinesClient", "listBySubscriptionNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListBySubscriptionComplete enumerates all values, automatically crossing page boundaries as required.
func (client VirtualMachinesClient) ListBySubscriptionComplete(ctx context.Context, filter string, top *int32, skipToken string) (result VirtualMachineListResponseIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualMachinesClient.ListBySubscription")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListBySubscription(ctx, filter, top, skipToken)
	return
}

// Start power on virtual machine
// Parameters:
// resourceGroupName - the name of the resource group
// virtualMachineName - virtual machine name
func (client VirtualMachinesClient) Start(ctx context.Context, resourceGroupName string, virtualMachineName string) (result VirtualMachinesStartFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualMachinesClient.Start")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.StartPreparer(ctx, resourceGroupName, virtualMachineName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.VirtualMachinesClient", "Start", nil, "Failure preparing request")
		return
	}

	result, err = client.StartSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.VirtualMachinesClient", "Start", nil, "Failure sending request")
		return
	}

	return
}

// StartPreparer prepares the Start request.
func (client VirtualMachinesClient) StartPreparer(ctx context.Context, resourceGroupName string, virtualMachineName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
		"virtualMachineName": autorest.Encode("path", virtualMachineName),
	}

	const APIVersion = "2019-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/virtualMachines/{virtualMachineName}/start", pathParameters),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("Referer", client.Referer))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// StartSender sends the Start request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachinesClient) StartSender(req *http.Request) (future VirtualMachinesStartFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client VirtualMachinesClient) (ar autorest.Response, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "vmwarecloudsimple.VirtualMachinesStartFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("vmwarecloudsimple.VirtualMachinesStartFuture")
			return
		}
		ar.Response = future.Response()
		return
	}
	return
}

// StartResponder handles the response to the Start request. The method always
// closes the http.Response Body.
func (client VirtualMachinesClient) StartResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Stop power off virtual machine, options: shutdown, poweroff, and suspend
// Parameters:
// resourceGroupName - the name of the resource group
// virtualMachineName - virtual machine name
// mParameter - body stop mode parameter (reboot, shutdown, etc...)
// mode - query stop mode parameter (reboot, shutdown, etc...)
func (client VirtualMachinesClient) Stop(ctx context.Context, resourceGroupName string, virtualMachineName string, mParameter *VirtualMachineStopMode, mode StopMode) (result VirtualMachinesStopFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualMachinesClient.Stop")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.StopPreparer(ctx, resourceGroupName, virtualMachineName, mParameter, mode)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.VirtualMachinesClient", "Stop", nil, "Failure preparing request")
		return
	}

	result, err = client.StopSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.VirtualMachinesClient", "Stop", nil, "Failure sending request")
		return
	}

	return
}

// StopPreparer prepares the Stop request.
func (client VirtualMachinesClient) StopPreparer(ctx context.Context, resourceGroupName string, virtualMachineName string, mParameter *VirtualMachineStopMode, mode StopMode) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
		"virtualMachineName": autorest.Encode("path", virtualMachineName),
	}

	const APIVersion = "2019-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(string(mode)) > 0 {
		queryParameters["mode"] = autorest.Encode("query", mode)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/virtualMachines/{virtualMachineName}/stop", pathParameters),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("Referer", client.Referer))
	if mParameter != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(mParameter))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// StopSender sends the Stop request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachinesClient) StopSender(req *http.Request) (future VirtualMachinesStopFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client VirtualMachinesClient) (ar autorest.Response, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "vmwarecloudsimple.VirtualMachinesStopFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("vmwarecloudsimple.VirtualMachinesStopFuture")
			return
		}
		ar.Response = future.Response()
		return
	}
	return
}

// StopResponder handles the response to the Stop request. The method always
// closes the http.Response Body.
func (client VirtualMachinesClient) StopResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Update patch virtual machine properties
// Parameters:
// resourceGroupName - the name of the resource group
// virtualMachineName - virtual machine name
// virtualMachineRequest - patch virtual machine request
func (client VirtualMachinesClient) Update(ctx context.Context, resourceGroupName string, virtualMachineName string, virtualMachineRequest PatchPayload) (result VirtualMachinesUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualMachinesClient.Update")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, resourceGroupName, virtualMachineName, virtualMachineRequest)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.VirtualMachinesClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.VirtualMachinesClient", "Update", nil, "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client VirtualMachinesClient) UpdatePreparer(ctx context.Context, resourceGroupName string, virtualMachineName string, virtualMachineRequest PatchPayload) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
		"virtualMachineName": autorest.Encode("path", virtualMachineName),
	}

	const APIVersion = "2019-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/virtualMachines/{virtualMachineName}", pathParameters),
		autorest.WithJSON(virtualMachineRequest),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachinesClient) UpdateSender(req *http.Request) (future VirtualMachinesUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client VirtualMachinesClient) (VM VirtualMachine, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "vmwarecloudsimple.VirtualMachinesUpdateFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("vmwarecloudsimple.VirtualMachinesUpdateFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		if VM.Response.Response, err = future.GetResult(sender); err == nil && VM.Response.Response.StatusCode != http.StatusNoContent {
			VM, err = client.UpdateResponder(VM.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "vmwarecloudsimple.VirtualMachinesUpdateFuture", "Result", VM.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client VirtualMachinesClient) UpdateResponder(resp *http.Response) (result VirtualMachine, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
