//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armappplatform_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2022-05-01-preview/examples/BuildpackBinding_Get.json
func ExampleBuildpackBindingClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappplatform.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBuildpackBindingClient().Get(ctx, "myResourceGroup", "myservice", "default", "default", "myBuildpackBinding", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.BuildpackBindingResource = armappplatform.BuildpackBindingResource{
	// 	Name: to.Ptr("myBuildpackBinding"),
	// 	Type: to.Ptr("Microsoft.AppPlatform/Spring/buildServices/builders/buildpackBindings"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AppPlatform/Spring/myservice/buildServices/default/builders/default/buildpackBindings/myBuildpackBinding"),
	// 	SystemData: &armappplatform.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-11T03:16:03.944Z"); return t}()),
	// 		CreatedBy: to.Ptr("sample-user"),
	// 		CreatedByType: to.Ptr(armappplatform.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-11T03:17:03.944Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("sample-user"),
	// 		LastModifiedByType: to.Ptr(armappplatform.LastModifiedByTypeUser),
	// 	},
	// 	Properties: &armappplatform.BuildpackBindingProperties{
	// 		BindingType: to.Ptr(armappplatform.BindingTypeApplicationInsights),
	// 		LaunchProperties: &armappplatform.BuildpackBindingLaunchProperties{
	// 			Properties: map[string]*string{
	// 				"abc": to.Ptr("def"),
	// 				"any-string": to.Ptr("any-string"),
	// 				"sampling-rate": to.Ptr("12.0"),
	// 			},
	// 			Secrets: map[string]*string{
	// 				"connection-string": to.Ptr("*"),
	// 			},
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2022-05-01-preview/examples/BuildpackBinding_CreateOrUpdate.json
func ExampleBuildpackBindingClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappplatform.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBuildpackBindingClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "myservice", "default", "default", "myBuildpackBinding", armappplatform.BuildpackBindingResource{
		Properties: &armappplatform.BuildpackBindingProperties{
			BindingType: to.Ptr(armappplatform.BindingTypeApplicationInsights),
			LaunchProperties: &armappplatform.BuildpackBindingLaunchProperties{
				Properties: map[string]*string{
					"abc":           to.Ptr("def"),
					"any-string":    to.Ptr("any-string"),
					"sampling-rate": to.Ptr("12.0"),
				},
				Secrets: map[string]*string{
					"connection-string": to.Ptr("XXXXXXXXXXXXXXXXX=XXXXXXXXXXXXX-XXXXXXXXXXXXXXXXXXX;XXXXXXXXXXXXXXXXX=XXXXXXXXXXXXXXXXXXX"),
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.BuildpackBindingResource = armappplatform.BuildpackBindingResource{
	// 	Name: to.Ptr("myBuildpackBinding"),
	// 	Type: to.Ptr("Microsoft.AppPlatform/Spring/buildServices/builders/buildpackBindings"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AppPlatform/Spring/myservice/buildServices/default/builders/default/buildpackBindings/myBuildpackBinding"),
	// 	SystemData: &armappplatform.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-11T03:16:03.944Z"); return t}()),
	// 		CreatedBy: to.Ptr("sample-user"),
	// 		CreatedByType: to.Ptr(armappplatform.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-11T03:17:03.944Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("sample-user"),
	// 		LastModifiedByType: to.Ptr(armappplatform.LastModifiedByTypeUser),
	// 	},
	// 	Properties: &armappplatform.BuildpackBindingProperties{
	// 		BindingType: to.Ptr(armappplatform.BindingTypeApplicationInsights),
	// 		LaunchProperties: &armappplatform.BuildpackBindingLaunchProperties{
	// 			Properties: map[string]*string{
	// 				"abc": to.Ptr("def"),
	// 				"any-string": to.Ptr("any-string"),
	// 				"sampling-rate": to.Ptr("12.0"),
	// 			},
	// 			Secrets: map[string]*string{
	// 				"connection-string": to.Ptr("*"),
	// 			},
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2022-05-01-preview/examples/BuildpackBinding_Delete.json
func ExampleBuildpackBindingClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappplatform.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBuildpackBindingClient().BeginDelete(ctx, "myResourceGroup", "myservice", "default", "default", "myBuildpackBinding", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2022-05-01-preview/examples/BuildpackBinding_List.json
func ExampleBuildpackBindingClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappplatform.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewBuildpackBindingClient().NewListPager("myResourceGroup", "myservice", "default", "default", nil)
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
		// page.BuildpackBindingResourceCollection = armappplatform.BuildpackBindingResourceCollection{
		// 	Value: []*armappplatform.BuildpackBindingResource{
		// 		{
		// 			Name: to.Ptr("myBuildpackBinding"),
		// 			Type: to.Ptr("Microsoft.AppPlatform/Spring/buildServices/builders/buildpackBindings"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AppPlatform/Spring/myservice/buildServices/default/builders/default/buildpackBindings/myBuildpackBinding"),
		// 			SystemData: &armappplatform.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-11T03:16:03.944Z"); return t}()),
		// 				CreatedBy: to.Ptr("sample-user"),
		// 				CreatedByType: to.Ptr(armappplatform.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-11T03:17:03.944Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("sample-user"),
		// 				LastModifiedByType: to.Ptr(armappplatform.LastModifiedByTypeUser),
		// 			},
		// 			Properties: &armappplatform.BuildpackBindingProperties{
		// 				BindingType: to.Ptr(armappplatform.BindingTypeApplicationInsights),
		// 				LaunchProperties: &armappplatform.BuildpackBindingLaunchProperties{
		// 					Properties: map[string]*string{
		// 						"abc": to.Ptr("def"),
		// 						"any-string": to.Ptr("any-string"),
		// 						"sampling-rate": to.Ptr("12.0"),
		// 					},
		// 					Secrets: map[string]*string{
		// 						"connection-string": to.Ptr("*"),
		// 					},
		// 				},
		// 			},
		// 	}},
		// }
	}
}
