//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armoperationsmanagement

import (
	"context"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"net/http"
	"time"
)

// ManagementAssociationsClientCreateOrUpdateResponse contains the response from method ManagementAssociationsClient.CreateOrUpdate.
type ManagementAssociationsClientCreateOrUpdateResponse struct {
	ManagementAssociationsClientCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ManagementAssociationsClientCreateOrUpdateResult contains the result from method ManagementAssociationsClient.CreateOrUpdate.
type ManagementAssociationsClientCreateOrUpdateResult struct {
	ManagementAssociation
}

// ManagementAssociationsClientDeleteResponse contains the response from method ManagementAssociationsClient.Delete.
type ManagementAssociationsClientDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ManagementAssociationsClientGetResponse contains the response from method ManagementAssociationsClient.Get.
type ManagementAssociationsClientGetResponse struct {
	ManagementAssociationsClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ManagementAssociationsClientGetResult contains the result from method ManagementAssociationsClient.Get.
type ManagementAssociationsClientGetResult struct {
	ManagementAssociation
}

// ManagementAssociationsClientListBySubscriptionResponse contains the response from method ManagementAssociationsClient.ListBySubscription.
type ManagementAssociationsClientListBySubscriptionResponse struct {
	ManagementAssociationsClientListBySubscriptionResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ManagementAssociationsClientListBySubscriptionResult contains the result from method ManagementAssociationsClient.ListBySubscription.
type ManagementAssociationsClientListBySubscriptionResult struct {
	ManagementAssociationPropertiesList
}

// ManagementConfigurationsClientCreateOrUpdateResponse contains the response from method ManagementConfigurationsClient.CreateOrUpdate.
type ManagementConfigurationsClientCreateOrUpdateResponse struct {
	ManagementConfigurationsClientCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ManagementConfigurationsClientCreateOrUpdateResult contains the result from method ManagementConfigurationsClient.CreateOrUpdate.
type ManagementConfigurationsClientCreateOrUpdateResult struct {
	ManagementConfiguration
}

// ManagementConfigurationsClientDeleteResponse contains the response from method ManagementConfigurationsClient.Delete.
type ManagementConfigurationsClientDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ManagementConfigurationsClientGetResponse contains the response from method ManagementConfigurationsClient.Get.
type ManagementConfigurationsClientGetResponse struct {
	ManagementConfigurationsClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ManagementConfigurationsClientGetResult contains the result from method ManagementConfigurationsClient.Get.
type ManagementConfigurationsClientGetResult struct {
	ManagementConfiguration
}

// ManagementConfigurationsClientListBySubscriptionResponse contains the response from method ManagementConfigurationsClient.ListBySubscription.
type ManagementConfigurationsClientListBySubscriptionResponse struct {
	ManagementConfigurationsClientListBySubscriptionResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ManagementConfigurationsClientListBySubscriptionResult contains the result from method ManagementConfigurationsClient.ListBySubscription.
type ManagementConfigurationsClientListBySubscriptionResult struct {
	ManagementConfigurationPropertiesList
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationsClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// OperationsClientListResult contains the result from method OperationsClient.List.
type OperationsClientListResult struct {
	OperationListResult
}

// SolutionsClientCreateOrUpdatePollerResponse contains the response from method SolutionsClient.CreateOrUpdate.
type SolutionsClientCreateOrUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *SolutionsClientCreateOrUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l SolutionsClientCreateOrUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (SolutionsClientCreateOrUpdateResponse, error) {
	respType := SolutionsClientCreateOrUpdateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.Solution)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a SolutionsClientCreateOrUpdatePollerResponse from the provided client and resume token.
func (l *SolutionsClientCreateOrUpdatePollerResponse) Resume(ctx context.Context, client *SolutionsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("SolutionsClient.CreateOrUpdate", token, client.pl)
	if err != nil {
		return err
	}
	poller := &SolutionsClientCreateOrUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// SolutionsClientCreateOrUpdateResponse contains the response from method SolutionsClient.CreateOrUpdate.
type SolutionsClientCreateOrUpdateResponse struct {
	SolutionsClientCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SolutionsClientCreateOrUpdateResult contains the result from method SolutionsClient.CreateOrUpdate.
type SolutionsClientCreateOrUpdateResult struct {
	Solution
}

// SolutionsClientDeletePollerResponse contains the response from method SolutionsClient.Delete.
type SolutionsClientDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *SolutionsClientDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l SolutionsClientDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (SolutionsClientDeleteResponse, error) {
	respType := SolutionsClientDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a SolutionsClientDeletePollerResponse from the provided client and resume token.
func (l *SolutionsClientDeletePollerResponse) Resume(ctx context.Context, client *SolutionsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("SolutionsClient.Delete", token, client.pl)
	if err != nil {
		return err
	}
	poller := &SolutionsClientDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// SolutionsClientDeleteResponse contains the response from method SolutionsClient.Delete.
type SolutionsClientDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SolutionsClientGetResponse contains the response from method SolutionsClient.Get.
type SolutionsClientGetResponse struct {
	SolutionsClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SolutionsClientGetResult contains the result from method SolutionsClient.Get.
type SolutionsClientGetResult struct {
	Solution
}

// SolutionsClientListByResourceGroupResponse contains the response from method SolutionsClient.ListByResourceGroup.
type SolutionsClientListByResourceGroupResponse struct {
	SolutionsClientListByResourceGroupResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SolutionsClientListByResourceGroupResult contains the result from method SolutionsClient.ListByResourceGroup.
type SolutionsClientListByResourceGroupResult struct {
	SolutionPropertiesList
}

// SolutionsClientListBySubscriptionResponse contains the response from method SolutionsClient.ListBySubscription.
type SolutionsClientListBySubscriptionResponse struct {
	SolutionsClientListBySubscriptionResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SolutionsClientListBySubscriptionResult contains the result from method SolutionsClient.ListBySubscription.
type SolutionsClientListBySubscriptionResult struct {
	SolutionPropertiesList
}

// SolutionsClientUpdatePollerResponse contains the response from method SolutionsClient.Update.
type SolutionsClientUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *SolutionsClientUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l SolutionsClientUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (SolutionsClientUpdateResponse, error) {
	respType := SolutionsClientUpdateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.Solution)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a SolutionsClientUpdatePollerResponse from the provided client and resume token.
func (l *SolutionsClientUpdatePollerResponse) Resume(ctx context.Context, client *SolutionsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("SolutionsClient.Update", token, client.pl)
	if err != nil {
		return err
	}
	poller := &SolutionsClientUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// SolutionsClientUpdateResponse contains the response from method SolutionsClient.Update.
type SolutionsClientUpdateResponse struct {
	SolutionsClientUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SolutionsClientUpdateResult contains the result from method SolutionsClient.Update.
type SolutionsClientUpdateResult struct {
	Solution
}
