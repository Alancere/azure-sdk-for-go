//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsecurity_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/af3f7994582c0cbd61a48b636907ad2ac95d332c/specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/JitNetworkAccessPolicies/GetJitNetworkAccessPoliciesSubscription_example.json
func ExampleJitNetworkAccessPoliciesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewJitNetworkAccessPoliciesClient().NewListPager(nil)
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
		// page.JitNetworkAccessPoliciesList = armsecurity.JitNetworkAccessPoliciesList{
		// 	Value: []*armsecurity.JitNetworkAccessPolicy{
		// 		{
		// 			Kind: to.Ptr("Basic"),
		// 			Location: to.Ptr("westeurope"),
		// 			Name: to.Ptr("default"),
		// 			Type: to.Ptr("Microsoft.Security/locations/jitNetworkAccessPolicies"),
		// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg1/providers/Microsoft.Security/locations/westeurope/jitNetworkAccessPolicies/default"),
		// 			Properties: &armsecurity.JitNetworkAccessPolicyProperties{
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				Requests: []*armsecurity.JitNetworkAccessRequest{
		// 					{
		// 						Justification: to.Ptr("testing a new version of the product"),
		// 						Requestor: to.Ptr("barbara@contoso.com"),
		// 						StartTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-05-17T08:06:45.5691611Z"); return t}()),
		// 						VirtualMachines: []*armsecurity.JitNetworkAccessRequestVirtualMachine{
		// 							{
		// 								ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg1/providers/Microsoft.Compute/virtualMachines/vm1"),
		// 								Ports: []*armsecurity.JitNetworkAccessRequestPort{
		// 									{
		// 										AllowedSourceAddressPrefix: to.Ptr("192.127.0.2"),
		// 										EndTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-05-17T09:06:45.5691611Z"); return t}()),
		// 										Number: to.Ptr[int32](3389),
		// 										Status: to.Ptr(armsecurity.StatusInitiated),
		// 										StatusReason: to.Ptr(armsecurity.StatusReasonUserRequested),
		// 								}},
		// 						}},
		// 				}},
		// 				VirtualMachines: []*armsecurity.JitNetworkAccessPolicyVirtualMachine{
		// 					{
		// 						ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg1/providers/Microsoft.Compute/virtualMachines/vm1"),
		// 						Ports: []*armsecurity.JitNetworkAccessPortRule{
		// 							{
		// 								AllowedSourceAddressPrefix: to.Ptr("*"),
		// 								MaxRequestAccessDuration: to.Ptr("PT3H"),
		// 								Number: to.Ptr[int32](22),
		// 								Protocol: to.Ptr(armsecurity.ProtocolAll),
		// 							},
		// 							{
		// 								AllowedSourceAddressPrefix: to.Ptr("*"),
		// 								MaxRequestAccessDuration: to.Ptr("PT3H"),
		// 								Number: to.Ptr[int32](3389),
		// 								Protocol: to.Ptr(armsecurity.ProtocolAll),
		// 						}},
		// 				}},
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/af3f7994582c0cbd61a48b636907ad2ac95d332c/specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/JitNetworkAccessPolicies/GetJitNetworkAccessPoliciesSubscriptionLocation_example.json
func ExampleJitNetworkAccessPoliciesClient_NewListByRegionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewJitNetworkAccessPoliciesClient().NewListByRegionPager("westeurope", nil)
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
		// page.JitNetworkAccessPoliciesList = armsecurity.JitNetworkAccessPoliciesList{
		// 	Value: []*armsecurity.JitNetworkAccessPolicy{
		// 		{
		// 			Kind: to.Ptr("Basic"),
		// 			Location: to.Ptr("westeurope"),
		// 			Name: to.Ptr("default"),
		// 			Type: to.Ptr("Microsoft.Security/locations/jitNetworkAccessPolicies"),
		// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg1/providers/Microsoft.Security/locations/westeurope/jitNetworkAccessPolicies/default"),
		// 			Properties: &armsecurity.JitNetworkAccessPolicyProperties{
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				Requests: []*armsecurity.JitNetworkAccessRequest{
		// 					{
		// 						Justification: to.Ptr("testing a new version of the product"),
		// 						Requestor: to.Ptr("barbara@contoso.com"),
		// 						StartTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-05-17T08:06:45.5691611Z"); return t}()),
		// 						VirtualMachines: []*armsecurity.JitNetworkAccessRequestVirtualMachine{
		// 							{
		// 								ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg1/providers/Microsoft.Compute/virtualMachines/vm1"),
		// 								Ports: []*armsecurity.JitNetworkAccessRequestPort{
		// 									{
		// 										AllowedSourceAddressPrefix: to.Ptr("192.127.0.2"),
		// 										EndTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-05-17T09:06:45.5691611Z"); return t}()),
		// 										Number: to.Ptr[int32](3389),
		// 										Status: to.Ptr(armsecurity.StatusInitiated),
		// 										StatusReason: to.Ptr(armsecurity.StatusReasonUserRequested),
		// 								}},
		// 						}},
		// 				}},
		// 				VirtualMachines: []*armsecurity.JitNetworkAccessPolicyVirtualMachine{
		// 					{
		// 						ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg1/providers/Microsoft.Compute/virtualMachines/vm1"),
		// 						Ports: []*armsecurity.JitNetworkAccessPortRule{
		// 							{
		// 								AllowedSourceAddressPrefix: to.Ptr("*"),
		// 								MaxRequestAccessDuration: to.Ptr("PT3H"),
		// 								Number: to.Ptr[int32](22),
		// 								Protocol: to.Ptr(armsecurity.ProtocolAll),
		// 							},
		// 							{
		// 								AllowedSourceAddressPrefix: to.Ptr("*"),
		// 								MaxRequestAccessDuration: to.Ptr("PT3H"),
		// 								Number: to.Ptr[int32](3389),
		// 								Protocol: to.Ptr(armsecurity.ProtocolAll),
		// 						}},
		// 				}},
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/af3f7994582c0cbd61a48b636907ad2ac95d332c/specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/JitNetworkAccessPolicies/GetJitNetworkAccessPoliciesResourceGroup_example.json
func ExampleJitNetworkAccessPoliciesClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewJitNetworkAccessPoliciesClient().NewListByResourceGroupPager("myRg1", nil)
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
		// page.JitNetworkAccessPoliciesList = armsecurity.JitNetworkAccessPoliciesList{
		// 	Value: []*armsecurity.JitNetworkAccessPolicy{
		// 		{
		// 			Kind: to.Ptr("Basic"),
		// 			Location: to.Ptr("westeurope"),
		// 			Name: to.Ptr("default"),
		// 			Type: to.Ptr("Microsoft.Security/locations/jitNetworkAccessPolicies"),
		// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg1/providers/Microsoft.Security/locations/westeurope/jitNetworkAccessPolicies/default"),
		// 			Properties: &armsecurity.JitNetworkAccessPolicyProperties{
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				Requests: []*armsecurity.JitNetworkAccessRequest{
		// 					{
		// 						Justification: to.Ptr("testing a new version of the product"),
		// 						Requestor: to.Ptr("barbara@contoso.com"),
		// 						StartTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-05-17T08:06:45.5691611Z"); return t}()),
		// 						VirtualMachines: []*armsecurity.JitNetworkAccessRequestVirtualMachine{
		// 							{
		// 								ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg1/providers/Microsoft.Compute/virtualMachines/vm1"),
		// 								Ports: []*armsecurity.JitNetworkAccessRequestPort{
		// 									{
		// 										AllowedSourceAddressPrefix: to.Ptr("192.127.0.2"),
		// 										EndTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-05-17T09:06:45.5691611Z"); return t}()),
		// 										Number: to.Ptr[int32](3389),
		// 										Status: to.Ptr(armsecurity.StatusInitiated),
		// 										StatusReason: to.Ptr(armsecurity.StatusReasonUserRequested),
		// 								}},
		// 						}},
		// 				}},
		// 				VirtualMachines: []*armsecurity.JitNetworkAccessPolicyVirtualMachine{
		// 					{
		// 						ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg1/providers/Microsoft.Compute/virtualMachines/vm1"),
		// 						Ports: []*armsecurity.JitNetworkAccessPortRule{
		// 							{
		// 								AllowedSourceAddressPrefix: to.Ptr("*"),
		// 								MaxRequestAccessDuration: to.Ptr("PT3H"),
		// 								Number: to.Ptr[int32](22),
		// 								Protocol: to.Ptr(armsecurity.ProtocolAll),
		// 							},
		// 							{
		// 								AllowedSourceAddressPrefix: to.Ptr("*"),
		// 								MaxRequestAccessDuration: to.Ptr("PT3H"),
		// 								Number: to.Ptr[int32](3389),
		// 								Protocol: to.Ptr(armsecurity.ProtocolAll),
		// 						}},
		// 				}},
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/af3f7994582c0cbd61a48b636907ad2ac95d332c/specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/JitNetworkAccessPolicies/GetJitNetworkAccessPoliciesResourceGroupLocation_example.json
func ExampleJitNetworkAccessPoliciesClient_NewListByResourceGroupAndRegionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewJitNetworkAccessPoliciesClient().NewListByResourceGroupAndRegionPager("myRg1", "westeurope", nil)
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
		// page.JitNetworkAccessPoliciesList = armsecurity.JitNetworkAccessPoliciesList{
		// 	Value: []*armsecurity.JitNetworkAccessPolicy{
		// 		{
		// 			Kind: to.Ptr("Basic"),
		// 			Location: to.Ptr("westeurope"),
		// 			Name: to.Ptr("default"),
		// 			Type: to.Ptr("Microsoft.Security/locations/jitNetworkAccessPolicies"),
		// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg1/providers/Microsoft.Security/locations/westeurope/jitNetworkAccessPolicies/default"),
		// 			Properties: &armsecurity.JitNetworkAccessPolicyProperties{
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				Requests: []*armsecurity.JitNetworkAccessRequest{
		// 					{
		// 						Justification: to.Ptr("testing a new version of the product"),
		// 						Requestor: to.Ptr("barbara@contoso.com"),
		// 						StartTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-05-17T08:06:45.5691611Z"); return t}()),
		// 						VirtualMachines: []*armsecurity.JitNetworkAccessRequestVirtualMachine{
		// 							{
		// 								ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg1/providers/Microsoft.Compute/virtualMachines/vm1"),
		// 								Ports: []*armsecurity.JitNetworkAccessRequestPort{
		// 									{
		// 										AllowedSourceAddressPrefix: to.Ptr("192.127.0.2"),
		// 										EndTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-05-17T09:06:45.5691611Z"); return t}()),
		// 										Number: to.Ptr[int32](3389),
		// 										Status: to.Ptr(armsecurity.StatusInitiated),
		// 										StatusReason: to.Ptr(armsecurity.StatusReasonUserRequested),
		// 								}},
		// 						}},
		// 				}},
		// 				VirtualMachines: []*armsecurity.JitNetworkAccessPolicyVirtualMachine{
		// 					{
		// 						ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg1/providers/Microsoft.Compute/virtualMachines/vm1"),
		// 						Ports: []*armsecurity.JitNetworkAccessPortRule{
		// 							{
		// 								AllowedSourceAddressPrefix: to.Ptr("*"),
		// 								MaxRequestAccessDuration: to.Ptr("PT3H"),
		// 								Number: to.Ptr[int32](22),
		// 								Protocol: to.Ptr(armsecurity.ProtocolAll),
		// 							},
		// 							{
		// 								AllowedSourceAddressPrefix: to.Ptr("*"),
		// 								MaxRequestAccessDuration: to.Ptr("PT3H"),
		// 								Number: to.Ptr[int32](3389),
		// 								Protocol: to.Ptr(armsecurity.ProtocolAll),
		// 						}},
		// 				}},
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/af3f7994582c0cbd61a48b636907ad2ac95d332c/specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/JitNetworkAccessPolicies/GetJitNetworkAccessPolicy_example.json
func ExampleJitNetworkAccessPoliciesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewJitNetworkAccessPoliciesClient().Get(ctx, "myRg1", "westeurope", "default", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.JitNetworkAccessPolicy = armsecurity.JitNetworkAccessPolicy{
	// 	Kind: to.Ptr("Basic"),
	// 	Location: to.Ptr("westeurope"),
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Security/locations/jitNetworkAccessPolicies"),
	// 	ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg1/providers/Microsoft.Security/locations/westeurope/jitNetworkAccessPolicies/default"),
	// 	Properties: &armsecurity.JitNetworkAccessPolicyProperties{
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		Requests: []*armsecurity.JitNetworkAccessRequest{
	// 			{
	// 				Justification: to.Ptr("testing a new version of the product"),
	// 				Requestor: to.Ptr("barbara@contoso.com"),
	// 				StartTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-05-17T08:06:45.5691611Z"); return t}()),
	// 				VirtualMachines: []*armsecurity.JitNetworkAccessRequestVirtualMachine{
	// 					{
	// 						ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg1/providers/Microsoft.Compute/virtualMachines/vm1"),
	// 						Ports: []*armsecurity.JitNetworkAccessRequestPort{
	// 							{
	// 								AllowedSourceAddressPrefix: to.Ptr("192.127.0.2"),
	// 								EndTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-05-17T09:06:45.5691611Z"); return t}()),
	// 								Number: to.Ptr[int32](3389),
	// 								Status: to.Ptr(armsecurity.StatusInitiated),
	// 								StatusReason: to.Ptr(armsecurity.StatusReasonUserRequested),
	// 						}},
	// 				}},
	// 		}},
	// 		VirtualMachines: []*armsecurity.JitNetworkAccessPolicyVirtualMachine{
	// 			{
	// 				ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg1/providers/Microsoft.Compute/virtualMachines/vm1"),
	// 				Ports: []*armsecurity.JitNetworkAccessPortRule{
	// 					{
	// 						AllowedSourceAddressPrefix: to.Ptr("*"),
	// 						MaxRequestAccessDuration: to.Ptr("PT3H"),
	// 						Number: to.Ptr[int32](22),
	// 						Protocol: to.Ptr(armsecurity.ProtocolAll),
	// 					},
	// 					{
	// 						AllowedSourceAddressPrefix: to.Ptr("*"),
	// 						MaxRequestAccessDuration: to.Ptr("PT3H"),
	// 						Number: to.Ptr[int32](3389),
	// 						Protocol: to.Ptr(armsecurity.ProtocolAll),
	// 				}},
	// 		}},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/af3f7994582c0cbd61a48b636907ad2ac95d332c/specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/JitNetworkAccessPolicies/CreateJitNetworkAccessPolicy_example.json
func ExampleJitNetworkAccessPoliciesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewJitNetworkAccessPoliciesClient().CreateOrUpdate(ctx, "myRg1", "westeurope", "default", armsecurity.JitNetworkAccessPolicy{
		Kind:     to.Ptr("Basic"),
		Location: to.Ptr("westeurope"),
		Name:     to.Ptr("default"),
		Type:     to.Ptr("Microsoft.Security/locations/jitNetworkAccessPolicies"),
		ID:       to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg1/providers/Microsoft.Security/locations/westeurope/jitNetworkAccessPolicies/default"),
		Properties: &armsecurity.JitNetworkAccessPolicyProperties{
			ProvisioningState: to.Ptr("Succeeded"),
			Requests: []*armsecurity.JitNetworkAccessRequest{
				{
					Requestor:    to.Ptr("barbara@contoso.com"),
					StartTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-05-17T08:06:45.5691611Z"); return t }()),
					VirtualMachines: []*armsecurity.JitNetworkAccessRequestVirtualMachine{
						{
							ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg1/providers/Microsoft.Compute/virtualMachines/vm1"),
							Ports: []*armsecurity.JitNetworkAccessRequestPort{
								{
									AllowedSourceAddressPrefix: to.Ptr("192.127.0.2"),
									EndTimeUTC:                 to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-05-17T09:06:45.5691611Z"); return t }()),
									Number:                     to.Ptr[int32](3389),
									Status:                     to.Ptr(armsecurity.StatusInitiated),
									StatusReason:               to.Ptr(armsecurity.StatusReasonUserRequested),
								}},
						}},
				}},
			VirtualMachines: []*armsecurity.JitNetworkAccessPolicyVirtualMachine{
				{
					ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg1/providers/Microsoft.Compute/virtualMachines/vm1"),
					Ports: []*armsecurity.JitNetworkAccessPortRule{
						{
							AllowedSourceAddressPrefix: to.Ptr("*"),
							MaxRequestAccessDuration:   to.Ptr("PT3H"),
							Number:                     to.Ptr[int32](22),
							Protocol:                   to.Ptr(armsecurity.ProtocolAll),
						},
						{
							AllowedSourceAddressPrefix: to.Ptr("*"),
							MaxRequestAccessDuration:   to.Ptr("PT3H"),
							Number:                     to.Ptr[int32](3389),
							Protocol:                   to.Ptr(armsecurity.ProtocolAll),
						}},
				}},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.JitNetworkAccessPolicy = armsecurity.JitNetworkAccessPolicy{
	// 	Kind: to.Ptr("Basic"),
	// 	Location: to.Ptr("westeurope"),
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Security/locations/jitNetworkAccessPolicies"),
	// 	ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg1/providers/Microsoft.Security/locations/westeurope/jitNetworkAccessPolicies/default"),
	// 	Properties: &armsecurity.JitNetworkAccessPolicyProperties{
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		Requests: []*armsecurity.JitNetworkAccessRequest{
	// 			{
	// 				Requestor: to.Ptr("barbara@contoso.com"),
	// 				StartTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-05-17T08:06:45.5691611Z"); return t}()),
	// 				VirtualMachines: []*armsecurity.JitNetworkAccessRequestVirtualMachine{
	// 					{
	// 						ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg1/providers/Microsoft.Compute/virtualMachines/vm1"),
	// 						Ports: []*armsecurity.JitNetworkAccessRequestPort{
	// 							{
	// 								AllowedSourceAddressPrefix: to.Ptr("192.127.0.2"),
	// 								EndTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-05-17T09:06:45.5691611Z"); return t}()),
	// 								Number: to.Ptr[int32](3389),
	// 								Status: to.Ptr(armsecurity.StatusInitiated),
	// 								StatusReason: to.Ptr(armsecurity.StatusReasonUserRequested),
	// 						}},
	// 				}},
	// 		}},
	// 		VirtualMachines: []*armsecurity.JitNetworkAccessPolicyVirtualMachine{
	// 			{
	// 				ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg1/providers/Microsoft.Compute/virtualMachines/vm1"),
	// 				Ports: []*armsecurity.JitNetworkAccessPortRule{
	// 					{
	// 						AllowedSourceAddressPrefix: to.Ptr("*"),
	// 						MaxRequestAccessDuration: to.Ptr("PT3H"),
	// 						Number: to.Ptr[int32](22),
	// 						Protocol: to.Ptr(armsecurity.ProtocolAll),
	// 					},
	// 					{
	// 						AllowedSourceAddressPrefix: to.Ptr("*"),
	// 						MaxRequestAccessDuration: to.Ptr("PT3H"),
	// 						Number: to.Ptr[int32](3389),
	// 						Protocol: to.Ptr(armsecurity.ProtocolAll),
	// 				}},
	// 		}},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/af3f7994582c0cbd61a48b636907ad2ac95d332c/specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/JitNetworkAccessPolicies/DeleteJitNetworkAccessPolicy_example.json
func ExampleJitNetworkAccessPoliciesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewJitNetworkAccessPoliciesClient().Delete(ctx, "myRg1", "westeurope", "default", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/af3f7994582c0cbd61a48b636907ad2ac95d332c/specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/JitNetworkAccessPolicies/InitiateJitNetworkAccessPolicy_example.json
func ExampleJitNetworkAccessPoliciesClient_Initiate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewJitNetworkAccessPoliciesClient().Initiate(ctx, "myRg1", "westeurope", "default", armsecurity.JitNetworkAccessPolicyInitiateRequest{
		Justification: to.Ptr("testing a new version of the product"),
		VirtualMachines: []*armsecurity.JitNetworkAccessPolicyInitiateVirtualMachine{
			{
				ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg1/providers/Microsoft.Compute/virtualMachines/vm1"),
				Ports: []*armsecurity.JitNetworkAccessPolicyInitiatePort{
					{
						AllowedSourceAddressPrefix: to.Ptr("192.127.0.2"),
						Number:                     to.Ptr[int32](3389),
					}},
			}},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
