//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpolicy_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy"
)

// x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/preview/2020-07-01-preview/examples/deletePolicyExemption.json
func ExampleExemptionsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armpolicy.NewExemptionsClient("<subscription-id>", cred, nil)
	_, err = client.Delete(ctx,
		"<scope>",
		"<policy-exemption-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/preview/2020-07-01-preview/examples/createOrUpdatePolicyExemption.json
func ExampleExemptionsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armpolicy.NewExemptionsClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<scope>",
		"<policy-exemption-name>",
		armpolicy.Exemption{
			Properties: &armpolicy.ExemptionProperties{
				Description:       to.StringPtr("<description>"),
				DisplayName:       to.StringPtr("<display-name>"),
				ExemptionCategory: armpolicy.ExemptionCategory("Waiver").ToPtr(),
				Metadata: map[string]interface{}{
					"reason": "Temporary exemption for a expensive VM demo",
				},
				PolicyAssignmentID: to.StringPtr("<policy-assignment-id>"),
				PolicyDefinitionReferenceIDs: []*string{
					to.StringPtr("Limit_Skus")},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ExemptionsClientCreateOrUpdateResult)
}

// x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/preview/2020-07-01-preview/examples/getPolicyExemption.json
func ExampleExemptionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armpolicy.NewExemptionsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<scope>",
		"<policy-exemption-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ExemptionsClientGetResult)
}

// x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/preview/2020-07-01-preview/examples/listPolicyExemptionsForSubscription.json
func ExampleExemptionsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armpolicy.NewExemptionsClient("<subscription-id>", cred, nil)
	pager := client.List(&armpolicy.ExemptionsClientListOptions{Filter: to.StringPtr("<filter>")})
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

// x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/preview/2020-07-01-preview/examples/listPolicyExemptionsForResourceGroup.json
func ExampleExemptionsClient_ListForResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armpolicy.NewExemptionsClient("<subscription-id>", cred, nil)
	pager := client.ListForResourceGroup("<resource-group-name>",
		&armpolicy.ExemptionsClientListForResourceGroupOptions{Filter: to.StringPtr("<filter>")})
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

// x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/preview/2020-07-01-preview/examples/listPolicyExemptionsForResource.json
func ExampleExemptionsClient_ListForResource() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armpolicy.NewExemptionsClient("<subscription-id>", cred, nil)
	pager := client.ListForResource("<resource-group-name>",
		"<resource-provider-namespace>",
		"<parent-resource-path>",
		"<resource-type>",
		"<resource-name>",
		&armpolicy.ExemptionsClientListForResourceOptions{Filter: nil})
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

// x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/preview/2020-07-01-preview/examples/listPolicyExemptionsForManagementGroup.json
func ExampleExemptionsClient_ListForManagementGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armpolicy.NewExemptionsClient("<subscription-id>", cred, nil)
	pager := client.ListForManagementGroup("<management-group-id>",
		&armpolicy.ExemptionsClientListForManagementGroupOptions{Filter: to.StringPtr("<filter>")})
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
