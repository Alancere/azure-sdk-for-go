//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

// x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/PrivateEndpointDelete.json
func ExamplePrivateEndpointsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armnetwork.NewPrivateEndpointsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<private-endpoint-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/PrivateEndpointGet.json
func ExamplePrivateEndpointsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armnetwork.NewPrivateEndpointsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<private-endpoint-name>",
		&armnetwork.PrivateEndpointsClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.PrivateEndpointsClientGetResult)
}

// x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/PrivateEndpointCreate.json
func ExamplePrivateEndpointsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armnetwork.NewPrivateEndpointsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<private-endpoint-name>",
		armnetwork.PrivateEndpoint{
			Location: to.StringPtr("<location>"),
			Properties: &armnetwork.PrivateEndpointProperties{
				CustomNetworkInterfaceName: to.StringPtr("<custom-network-interface-name>"),
				IPConfigurations: []*armnetwork.PrivateEndpointIPConfiguration{
					{
						Name: to.StringPtr("<name>"),
						Properties: &armnetwork.PrivateEndpointIPConfigurationProperties{
							GroupID:          to.StringPtr("<group-id>"),
							MemberName:       to.StringPtr("<member-name>"),
							PrivateIPAddress: to.StringPtr("<private-ipaddress>"),
						},
					}},
				PrivateLinkServiceConnections: []*armnetwork.PrivateLinkServiceConnection{
					{
						Properties: &armnetwork.PrivateLinkServiceConnectionProperties{
							GroupIDs: []*string{
								to.StringPtr("groupIdFromResource")},
							PrivateLinkServiceID: to.StringPtr("<private-link-service-id>"),
							RequestMessage:       to.StringPtr("<request-message>"),
						},
					}},
				Subnet: &armnetwork.Subnet{
					ID: to.StringPtr("<id>"),
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
	log.Printf("Response result: %#v\n", res.PrivateEndpointsClientCreateOrUpdateResult)
}

// x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/PrivateEndpointList.json
func ExamplePrivateEndpointsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armnetwork.NewPrivateEndpointsClient("<subscription-id>", cred, nil)
	pager := client.List("<resource-group-name>",
		nil)
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}

// x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/PrivateEndpointListAll.json
func ExamplePrivateEndpointsClient_ListBySubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armnetwork.NewPrivateEndpointsClient("<subscription-id>", cred, nil)
	pager := client.ListBySubscription(nil)
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}
