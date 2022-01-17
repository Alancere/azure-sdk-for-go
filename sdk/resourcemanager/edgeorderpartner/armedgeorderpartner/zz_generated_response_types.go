//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armedgeorderpartner

import (
	"context"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"net/http"
	"time"
)

// APISClientListOperationsPartnerResponse contains the response from method APISClient.ListOperationsPartner.
type APISClientListOperationsPartnerResponse struct {
	APISClientListOperationsPartnerResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// APISClientListOperationsPartnerResult contains the result from method APISClient.ListOperationsPartner.
type APISClientListOperationsPartnerResult struct {
	OperationListResult
}

// APISClientManageInventoryMetadataPollerResponse contains the response from method APISClient.ManageInventoryMetadata.
type APISClientManageInventoryMetadataPollerResponse struct {
	// Poller contains an initialized poller.
	Poller *APISClientManageInventoryMetadataPoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l APISClientManageInventoryMetadataPollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (APISClientManageInventoryMetadataResponse, error) {
	respType := APISClientManageInventoryMetadataResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a APISClientManageInventoryMetadataPollerResponse from the provided client and resume token.
func (l *APISClientManageInventoryMetadataPollerResponse) Resume(ctx context.Context, client *APISClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("APISClient.ManageInventoryMetadata", token, client.pl)
	if err != nil {
		return err
	}
	poller := &APISClientManageInventoryMetadataPoller{
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

// APISClientManageInventoryMetadataResponse contains the response from method APISClient.ManageInventoryMetadata.
type APISClientManageInventoryMetadataResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// APISClientManageLinkResponse contains the response from method APISClient.ManageLink.
type APISClientManageLinkResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// APISClientSearchInventoriesResponse contains the response from method APISClient.SearchInventories.
type APISClientSearchInventoriesResponse struct {
	APISClientSearchInventoriesResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// APISClientSearchInventoriesResult contains the result from method APISClient.SearchInventories.
type APISClientSearchInventoriesResult struct {
	PartnerInventoryList
}
