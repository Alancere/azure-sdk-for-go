//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-07-01/examples/virtualMachineImageExamples/VirtualMachineImagesEdgeZone_Get_MaximumSet_Gen.json
func ExampleVirtualMachineImagesEdgeZoneClient_Get_virtualMachineImagesEdgeZoneGetMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineImagesEdgeZoneClient().Get(ctx, "aaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaa", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualMachineImage = armcompute.VirtualMachineImage{
	// 	ID: to.Ptr("aaaaaaaaaaa"),
	// 	Name: to.Ptr("aaaaaaaaa"),
	// 	ExtendedLocation: &armcompute.ExtendedLocation{
	// 		Name: to.Ptr("aaaaaaaaaaaaaaaaaaaaa"),
	// 		Type: to.Ptr(armcompute.ExtendedLocationTypesEdgeZone),
	// 	},
	// 	Location: to.Ptr("aaaaa"),
	// 	Tags: map[string]*string{
	// 		"key6817": to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 	},
	// 	Properties: &armcompute.VirtualMachineImageProperties{
	// 		AutomaticOSUpgradeProperties: &armcompute.AutomaticOSUpgradeProperties{
	// 			AutomaticOSUpgradeSupported: to.Ptr(true),
	// 		},
	// 		DataDiskImages: []*armcompute.DataDiskImage{
	// 			{
	// 				Lun: to.Ptr[int32](17),
	// 		}},
	// 		Disallowed: &armcompute.DisallowedConfiguration{
	// 			VMDiskType: to.Ptr(armcompute.VMDiskTypesNone),
	// 		},
	// 		Features: []*armcompute.VirtualMachineImageFeature{
	// 			{
	// 				Name: to.Ptr("aaaaaaaaaaaaaaaaaaaaa"),
	// 				Value: to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
	// 		}},
	// 		HyperVGeneration: to.Ptr(armcompute.HyperVGenerationTypesV1),
	// 		OSDiskImage: &armcompute.OSDiskImage{
	// 			OperatingSystem: to.Ptr(armcompute.OperatingSystemTypesWindows),
	// 		},
	// 		Plan: &armcompute.PurchasePlan{
	// 			Name: to.Ptr("aaaaaaaaa"),
	// 			Product: to.Ptr("aaaaaaaaaaaaaa"),
	// 			Publisher: to.Ptr("aaaaaaaaaaaaaaaaaaa"),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-07-01/examples/virtualMachineImageExamples/VirtualMachineImagesEdgeZone_Get_MinimumSet_Gen.json
func ExampleVirtualMachineImagesEdgeZoneClient_Get_virtualMachineImagesEdgeZoneGetMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineImagesEdgeZoneClient().Get(ctx, "aaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaa", "aaaaaaaaaaaaaaaaaa", "aa", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualMachineImage = armcompute.VirtualMachineImage{
	// 	ID: to.Ptr("aaaaaaaaaaa"),
	// 	Name: to.Ptr("aaaaaaaaa"),
	// 	Location: to.Ptr("aaaaa"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-07-01/examples/virtualMachineImageExamples/VirtualMachineImagesEdgeZone_List_MaximumSet_Gen.json
func ExampleVirtualMachineImagesEdgeZoneClient_List_virtualMachineImagesEdgeZoneListMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineImagesEdgeZoneClient().List(ctx, "aaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaa", "aaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaa", &armcompute.VirtualMachineImagesEdgeZoneClientListOptions{Expand: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaa"),
		Top:     to.Ptr[int32](12),
		Orderby: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualMachineImageResourceArray = []*armcompute.VirtualMachineImageResource{
	// 	{
	// 		ID: to.Ptr("aaaaaaaaaaa"),
	// 		Name: to.Ptr("aaaaaaaa"),
	// 		ExtendedLocation: &armcompute.ExtendedLocation{
	// 			Name: to.Ptr("aaaaaaaaaaaaaaaaaaaaa"),
	// 			Type: to.Ptr(armcompute.ExtendedLocationTypesEdgeZone),
	// 		},
	// 		Location: to.Ptr("aaaaaaaaaaaaaaaaaa"),
	// 		Tags: map[string]*string{
	// 			"key7868": to.Ptr("aaaaa"),
	// 		},
	// }}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-07-01/examples/virtualMachineImageExamples/VirtualMachineImagesEdgeZone_List_MinimumSet_Gen.json
func ExampleVirtualMachineImagesEdgeZoneClient_List_virtualMachineImagesEdgeZoneListMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineImagesEdgeZoneClient().List(ctx, "aaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaa", "aaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaa", &armcompute.VirtualMachineImagesEdgeZoneClientListOptions{Expand: nil,
		Top:     nil,
		Orderby: nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualMachineImageResourceArray = []*armcompute.VirtualMachineImageResource{
	// 	{
	// 		ID: to.Ptr("aaaaaaaaaaa"),
	// 		Name: to.Ptr("aaaaaaaa"),
	// 		Location: to.Ptr("aaaaaaaaaaaaaaaaaa"),
	// }}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-07-01/examples/virtualMachineImageExamples/VirtualMachineImagesEdgeZone_ListOffers_MaximumSet_Gen.json
func ExampleVirtualMachineImagesEdgeZoneClient_ListOffers_virtualMachineImagesEdgeZoneListOffersMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineImagesEdgeZoneClient().ListOffers(ctx, "aaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaa", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualMachineImageResourceArray = []*armcompute.VirtualMachineImageResource{
	// 	{
	// 		ID: to.Ptr("aaaaaaaaaaa"),
	// 		Name: to.Ptr("aaaaaaaa"),
	// 		ExtendedLocation: &armcompute.ExtendedLocation{
	// 			Name: to.Ptr("aaaaaaaaaaaaaaaaaaaaa"),
	// 			Type: to.Ptr(armcompute.ExtendedLocationTypesEdgeZone),
	// 		},
	// 		Location: to.Ptr("aaaaaaaaaaaaaaaaaa"),
	// 		Tags: map[string]*string{
	// 			"key7868": to.Ptr("aaaaa"),
	// 		},
	// }}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-07-01/examples/virtualMachineImageExamples/VirtualMachineImagesEdgeZone_ListOffers_MinimumSet_Gen.json
func ExampleVirtualMachineImagesEdgeZoneClient_ListOffers_virtualMachineImagesEdgeZoneListOffersMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineImagesEdgeZoneClient().ListOffers(ctx, "aaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaa", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualMachineImageResourceArray = []*armcompute.VirtualMachineImageResource{
	// 	{
	// 		ID: to.Ptr("aaaaaaaaaaa"),
	// 		Name: to.Ptr("aaaaaaaa"),
	// 		Location: to.Ptr("aaaaaaaaaaaaaaaaaa"),
	// }}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-07-01/examples/virtualMachineImageExamples/VirtualMachineImagesEdgeZone_ListPublishers_MaximumSet_Gen.json
func ExampleVirtualMachineImagesEdgeZoneClient_ListPublishers_virtualMachineImagesEdgeZoneListPublishersMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineImagesEdgeZoneClient().ListPublishers(ctx, "aaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaa", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualMachineImageResourceArray = []*armcompute.VirtualMachineImageResource{
	// 	{
	// 		ID: to.Ptr("aaaaaaaaaaa"),
	// 		Name: to.Ptr("aaaaaaaa"),
	// 		ExtendedLocation: &armcompute.ExtendedLocation{
	// 			Name: to.Ptr("aaaaaaaaaaaaaaaaaaaaa"),
	// 			Type: to.Ptr(armcompute.ExtendedLocationTypesEdgeZone),
	// 		},
	// 		Location: to.Ptr("aaaaaaaaaaaaaaaaaa"),
	// 		Tags: map[string]*string{
	// 			"key7868": to.Ptr("aaaaa"),
	// 		},
	// }}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-07-01/examples/virtualMachineImageExamples/VirtualMachineImagesEdgeZone_ListPublishers_MinimumSet_Gen.json
func ExampleVirtualMachineImagesEdgeZoneClient_ListPublishers_virtualMachineImagesEdgeZoneListPublishersMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineImagesEdgeZoneClient().ListPublishers(ctx, "aaaa", "aaaaaaaaaaaaaaaaaaaaaaa", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualMachineImageResourceArray = []*armcompute.VirtualMachineImageResource{
	// 	{
	// 		ID: to.Ptr("aaaaaaaaaaa"),
	// 		Name: to.Ptr("aaaaaaaa"),
	// 		Location: to.Ptr("aaaaaaaaaaaaaaaaaa"),
	// }}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-07-01/examples/virtualMachineImageExamples/VirtualMachineImagesEdgeZone_ListSkus_MaximumSet_Gen.json
func ExampleVirtualMachineImagesEdgeZoneClient_ListSKUs_virtualMachineImagesEdgeZoneListSkusMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineImagesEdgeZoneClient().ListSKUs(ctx, "aaaaaaaaaaaa", "aaaaa", "aaaaaaaaaaaa", "aaaaaaaaaaaa", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualMachineImageResourceArray = []*armcompute.VirtualMachineImageResource{
	// 	{
	// 		ID: to.Ptr("aaaaaaaaaaa"),
	// 		Name: to.Ptr("aaaaaaaa"),
	// 		ExtendedLocation: &armcompute.ExtendedLocation{
	// 			Name: to.Ptr("aaaaaaaaaaaaaaaaaaaaa"),
	// 			Type: to.Ptr(armcompute.ExtendedLocationTypesEdgeZone),
	// 		},
	// 		Location: to.Ptr("aaaaaaaaaaaaaaaaaa"),
	// 		Tags: map[string]*string{
	// 			"key7868": to.Ptr("aaaaa"),
	// 		},
	// }}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-07-01/examples/virtualMachineImageExamples/VirtualMachineImagesEdgeZone_ListSkus_MinimumSet_Gen.json
func ExampleVirtualMachineImagesEdgeZoneClient_ListSKUs_virtualMachineImagesEdgeZoneListSkusMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineImagesEdgeZoneClient().ListSKUs(ctx, "aaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaa", "aaaaaaaaaaaa", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualMachineImageResourceArray = []*armcompute.VirtualMachineImageResource{
	// 	{
	// 		ID: to.Ptr("aaaaaaaaaaa"),
	// 		Name: to.Ptr("aaaaaaaa"),
	// 		Location: to.Ptr("aaaaaaaaaaaaaaaaaa"),
	// }}
}
