//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmixedreality_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mixedreality/armmixedreality"
)

// x-ms-original-file: specification/mixedreality/resource-manager/Microsoft.MixedReality/preview/2021-03-01-preview/examples/remote-rendering/GetBySubscription.json
func ExampleRemoteRenderingAccountsClient_ListBySubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmixedreality.NewRemoteRenderingAccountsClient("<subscription-id>", cred, nil)
	pager := client.ListBySubscription(nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("RemoteRenderingAccount.ID: %s\n", *v.ID)
		}
	}
}

// x-ms-original-file: specification/mixedreality/resource-manager/Microsoft.MixedReality/preview/2021-03-01-preview/examples/remote-rendering/GetByResourceGroup.json
func ExampleRemoteRenderingAccountsClient_ListByResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmixedreality.NewRemoteRenderingAccountsClient("<subscription-id>", cred, nil)
	pager := client.ListByResourceGroup("<resource-group-name>",
		nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("RemoteRenderingAccount.ID: %s\n", *v.ID)
		}
	}
}

// x-ms-original-file: specification/mixedreality/resource-manager/Microsoft.MixedReality/preview/2021-03-01-preview/examples/remote-rendering/Delete.json
func ExampleRemoteRenderingAccountsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmixedreality.NewRemoteRenderingAccountsClient("<subscription-id>", cred, nil)
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<account-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/mixedreality/resource-manager/Microsoft.MixedReality/preview/2021-03-01-preview/examples/remote-rendering/Get.json
func ExampleRemoteRenderingAccountsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmixedreality.NewRemoteRenderingAccountsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<account-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("RemoteRenderingAccount.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/mixedreality/resource-manager/Microsoft.MixedReality/preview/2021-03-01-preview/examples/remote-rendering/Patch.json
func ExampleRemoteRenderingAccountsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmixedreality.NewRemoteRenderingAccountsClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<account-name>",
		armmixedreality.RemoteRenderingAccount{
			TrackedResource: armmixedreality.TrackedResource{
				Location: to.StringPtr("<location>"),
				Tags: map[string]*string{
					"hero":    to.StringPtr("romeo"),
					"heroine": to.StringPtr("juliet"),
				},
			},
			Identity: &armmixedreality.Identity{
				Type: to.StringPtr("<type>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("RemoteRenderingAccount.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/mixedreality/resource-manager/Microsoft.MixedReality/preview/2021-03-01-preview/examples/remote-rendering/Put.json
func ExampleRemoteRenderingAccountsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmixedreality.NewRemoteRenderingAccountsClient("<subscription-id>", cred, nil)
	res, err := client.Create(ctx,
		"<resource-group-name>",
		"<account-name>",
		armmixedreality.RemoteRenderingAccount{
			TrackedResource: armmixedreality.TrackedResource{
				Location: to.StringPtr("<location>"),
			},
			Identity: &armmixedreality.Identity{
				Type: to.StringPtr("<type>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("RemoteRenderingAccount.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/mixedreality/resource-manager/Microsoft.MixedReality/preview/2021-03-01-preview/examples/remote-rendering/ListKeys.json
func ExampleRemoteRenderingAccountsClient_ListKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmixedreality.NewRemoteRenderingAccountsClient("<subscription-id>", cred, nil)
	_, err = client.ListKeys(ctx,
		"<resource-group-name>",
		"<account-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/mixedreality/resource-manager/Microsoft.MixedReality/preview/2021-03-01-preview/examples/remote-rendering/RegenerateKey.json
func ExampleRemoteRenderingAccountsClient_RegenerateKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmixedreality.NewRemoteRenderingAccountsClient("<subscription-id>", cred, nil)
	_, err = client.RegenerateKeys(ctx,
		"<resource-group-name>",
		"<account-name>",
		armmixedreality.AccountKeyRegenerateRequest{
			Serial: armmixedreality.SerialPrimary.ToPtr(),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
