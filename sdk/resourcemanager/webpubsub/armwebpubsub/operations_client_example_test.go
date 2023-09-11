//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armwebpubsub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/webpubsub/armwebpubsub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/webpubsub/resource-manager/Microsoft.SignalRService/preview/2023-06-01-preview/examples/Operations_List.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armwebpubsub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOperationsClient().NewListPager(nil)
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
		// page.OperationList = armwebpubsub.OperationList{
		// 	Value: []*armwebpubsub.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.SignalRService/WebPubSub/read"),
		// 			Display: &armwebpubsub.OperationDisplay{
		// 				Description: to.Ptr("View the resource settings and configurations in the management portal or through API"),
		// 				Operation: to.Ptr("Manage WebPubSub (read-only)"),
		// 				Provider: to.Ptr("Microsoft.SignalRService"),
		// 				Resource: to.Ptr("WebPubSub"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Properties: &armwebpubsub.OperationProperties{
		// 			},
		// 	}},
		// }
	}
}
