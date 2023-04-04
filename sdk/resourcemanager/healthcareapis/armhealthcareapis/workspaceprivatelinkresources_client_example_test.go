//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armhealthcareapis_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/preview/2022-01-31-preview/examples/privatelink/PrivateLinkResourcesListByWorkspace.json
func ExampleWorkspacePrivateLinkResourcesClient_NewListByWorkspacePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhealthcareapis.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkspacePrivateLinkResourcesClient().NewListByWorkspacePager("testRG", "workspace1", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.PrivateLinkResourceListResultDescription = armhealthcareapis.PrivateLinkResourceListResultDescription{
		// 	Value: []*armhealthcareapis.PrivateLinkResourceDescription{
		// 		{
		// 			Name: to.Ptr("healthcareworkspace"),
		// 			Type: to.Ptr("Microsoft.HealthcareApis/workspaces/privateLinkResources"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.HealthcareApis/workspaces/workspace1/privateLinkResources/healthcareworkspace"),
		// 			Properties: &armhealthcareapis.PrivateLinkResourceProperties{
		// 				GroupID: to.Ptr("healthcareworkspace"),
		// 				RequiredMembers: []*string{
		// 					to.Ptr("fhirservices"),
		// 					to.Ptr("dicomservices")},
		// 					RequiredZoneNames: []*string{
		// 						to.Ptr("privatelink.fhir.azurehealthcareapis.com"),
		// 						to.Ptr("privatelink.dicom.azurehealthcareapis.com")},
		// 					},
		// 			}},
		// 		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/preview/2022-01-31-preview/examples/privatelink/WorkspacePrivateLinkResourceGet.json
func ExampleWorkspacePrivateLinkResourcesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhealthcareapis.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkspacePrivateLinkResourcesClient().Get(ctx, "testRG", "workspace1", "healthcareworkspace", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateLinkResourceDescription = armhealthcareapis.PrivateLinkResourceDescription{
	// 	Name: to.Ptr("fhir"),
	// 	Type: to.Ptr("Microsoft.HealthcareApis/workspaces/privateLinkResources"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.HealthcareApis/workspaces/workspace1/privateLinkResources/fhir"),
	// 	Properties: &armhealthcareapis.PrivateLinkResourceProperties{
	// 		GroupID: to.Ptr("healthcareworkspace"),
	// 		RequiredMembers: []*string{
	// 			to.Ptr("fhirservices"),
	// 			to.Ptr("dicomservices")},
	// 			RequiredZoneNames: []*string{
	// 				to.Ptr("privatelink.fhir.azurehealthcareapis.com"),
	// 				to.Ptr("privatelink.dicom.azurehealthcareapis.com")},
	// 			},
	// 		}
}
