//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapimanagement_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement"
)

// x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementListPrivateEndpointConnections.json
func ExamplePrivateEndpointConnectionClient_ListByService() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armapimanagement.NewPrivateEndpointConnectionClient("<subscription-id>", cred, nil)
	res, err := client.ListByService(ctx,
		"<resource-group-name>",
		"<service-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.PrivateEndpointConnectionClientListByServiceResult)
}

// x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementGetPrivateEndpointConnection.json
func ExamplePrivateEndpointConnectionClient_GetByName() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armapimanagement.NewPrivateEndpointConnectionClient("<subscription-id>", cred, nil)
	res, err := client.GetByName(ctx,
		"<resource-group-name>",
		"<service-name>",
		"<private-endpoint-connection-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.PrivateEndpointConnectionClientGetByNameResult)
}

// x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementApproveOrRejectPrivateEndpointConnection.json
func ExamplePrivateEndpointConnectionClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armapimanagement.NewPrivateEndpointConnectionClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<service-name>",
		"<private-endpoint-connection-name>",
		armapimanagement.PrivateEndpointConnectionRequest{
			ID: to.StringPtr("<id>"),
			Properties: &armapimanagement.PrivateEndpointConnectionRequestProperties{
				PrivateLinkServiceConnectionState: &armapimanagement.PrivateLinkServiceConnectionState{
					Description: to.StringPtr("<description>"),
					Status:      armapimanagement.PrivateEndpointServiceConnectionStatus("Approved").ToPtr(),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.PrivateEndpointConnectionClientCreateOrUpdateResult)
}

// x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementDeletePrivateEndpointConnection.json
func ExamplePrivateEndpointConnectionClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armapimanagement.NewPrivateEndpointConnectionClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<service-name>",
		"<private-endpoint-connection-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementListPrivateLinkGroupResources.json
func ExamplePrivateEndpointConnectionClient_ListPrivateLinkResources() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armapimanagement.NewPrivateEndpointConnectionClient("<subscription-id>", cred, nil)
	res, err := client.ListPrivateLinkResources(ctx,
		"<resource-group-name>",
		"<service-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.PrivateEndpointConnectionClientListPrivateLinkResourcesResult)
}

// x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementGetPrivateLinkGroupResource.json
func ExamplePrivateEndpointConnectionClient_GetPrivateLinkResource() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armapimanagement.NewPrivateEndpointConnectionClient("<subscription-id>", cred, nil)
	res, err := client.GetPrivateLinkResource(ctx,
		"<resource-group-name>",
		"<service-name>",
		"<private-link-sub-resource-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.PrivateEndpointConnectionClientGetPrivateLinkResourceResult)
}
