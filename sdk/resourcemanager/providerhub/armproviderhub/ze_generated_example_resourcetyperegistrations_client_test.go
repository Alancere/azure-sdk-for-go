//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armproviderhub_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/providerhub/armproviderhub"
)

// x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/ResourceTypeRegistrations_Get.json
func ExampleResourceTypeRegistrationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armproviderhub.NewResourceTypeRegistrationsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<provider-namespace>",
		"<resource-type>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ResourceTypeRegistrationsClientGetResult)
}

// x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/ResourceTypeRegistrations_CreateOrUpdate.json
func ExampleResourceTypeRegistrationsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armproviderhub.NewResourceTypeRegistrationsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<provider-namespace>",
		"<resource-type>",
		armproviderhub.ResourceTypeRegistration{
			Properties: &armproviderhub.ResourceTypeRegistrationProperties{
				Endpoints: []*armproviderhub.ResourceTypeEndpoint{
					{
						APIVersions: []*string{
							to.StringPtr("2020-06-01-preview")},
						Locations: []*string{
							to.StringPtr("West US"),
							to.StringPtr("East US"),
							to.StringPtr("North Europe")},
						RequiredFeatures: []*string{
							to.StringPtr("<feature flag>")},
					}},
				Regionality: armproviderhub.Regionality("Regional").ToPtr(),
				RoutingType: armproviderhub.RoutingType("Default").ToPtr(),
				SwaggerSpecifications: []*armproviderhub.SwaggerSpecification{
					{
						APIVersions: []*string{
							to.StringPtr("2020-06-01-preview")},
						SwaggerSpecFolderURI: to.StringPtr("<swagger-spec-folder-uri>"),
					}},
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
	log.Printf("Response result: %#v\n", res.ResourceTypeRegistrationsClientCreateOrUpdateResult)
}

// x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/ResourceTypeRegistrations_Delete.json
func ExampleResourceTypeRegistrationsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armproviderhub.NewResourceTypeRegistrationsClient("<subscription-id>", cred, nil)
	_, err = client.Delete(ctx,
		"<provider-namespace>",
		"<resource-type>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/ResourceTypeRegistrations_ListByProviderRegistration.json
func ExampleResourceTypeRegistrationsClient_ListByProviderRegistration() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armproviderhub.NewResourceTypeRegistrationsClient("<subscription-id>", cred, nil)
	pager := client.ListByProviderRegistration("<provider-namespace>",
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
