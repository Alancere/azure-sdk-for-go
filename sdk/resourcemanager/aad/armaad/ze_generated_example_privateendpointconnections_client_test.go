//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armaad_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/aad/armaad"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/azureactivedirectory/resource-manager/Microsoft.Aadiam/stable/2020-03-01/examples/AzureADPrivateEndpointConnectionsList.json
func ExamplePrivateEndpointConnectionsClient_NewListByPolicyNamePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armaad.NewPrivateEndpointConnectionsClient("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListByPolicyNamePager("myResourceGroup",
		"example-policy-5849",
		nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/azureactivedirectory/resource-manager/Microsoft.Aadiam/stable/2020-03-01/examples/AzureADPrivateEndpointConnectionsCreate.json
func ExamplePrivateEndpointConnectionsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armaad.NewPrivateEndpointConnectionsClient("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx,
		"resourcegroup",
		"example-policy-5849",
		"{privateEndpointConnection name}",
		armaad.PrivateEndpointConnection{
			Properties: &armaad.PrivateEndpointConnectionProperties{
				PrivateEndpoint: &armaad.PrivateEndpoint{
					ID: to.Ptr("subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/microsoft.aadiam/privateLinkForAzureAD/ddb1/privateLinkConnections/{privateEndpointConnection name}"),
				},
				PrivateLinkServiceConnectionState: &armaad.PrivateLinkServiceConnectionState{
					Description:     to.Ptr("You may pass"),
					ActionsRequired: to.Ptr("None"),
					Status:          to.Ptr(armaad.PrivateEndpointServiceConnectionStatusApproved),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/azureactivedirectory/resource-manager/Microsoft.Aadiam/stable/2020-03-01/examples/AzureADPrivateEndpointConnectionsDelete.json
func ExamplePrivateEndpointConnectionsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armaad.NewPrivateEndpointConnectionsClient("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDelete(ctx,
		"myResourceGroup",
		"example-policy-5849",
		"{privateEndpointConnection name}",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/azureactivedirectory/resource-manager/Microsoft.Aadiam/stable/2020-03-01/examples/AzureADPrivateEndpointConnectionsGet.json
func ExamplePrivateEndpointConnectionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armaad.NewPrivateEndpointConnectionsClient("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx,
		"myResourceGroup",
		"example-policy-5849",
		"{privateEndpointConnection name}",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
