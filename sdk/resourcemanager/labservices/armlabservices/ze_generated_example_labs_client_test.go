//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armlabservices_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/labservices/armlabservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/labservices/resource-manager/Microsoft.LabServices/preview/2021-11-15-preview/examples/Labs/listLabs.json
func ExampleLabsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armlabservices.NewLabsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	pager := client.NewListBySubscriptionPager(&armlabservices.LabsClientListBySubscriptionOptions{Filter: nil})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
			return
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/labservices/resource-manager/Microsoft.LabServices/preview/2021-11-15-preview/examples/Labs/listResourceGroupLabs.json
func ExampleLabsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armlabservices.NewLabsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	pager := client.NewListByResourceGroupPager("<resource-group-name>",
		nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
			return
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/labservices/resource-manager/Microsoft.LabServices/preview/2021-11-15-preview/examples/Labs/getLab.json
func ExampleLabsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armlabservices.NewLabsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<lab-name>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/labservices/resource-manager/Microsoft.LabServices/preview/2021-11-15-preview/examples/Labs/putLab.json
func ExampleLabsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armlabservices.NewLabsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<lab-name>",
		armlabservices.Lab{
			Location: to.Ptr("<location>"),
			Properties: &armlabservices.LabProperties{
				Description: to.Ptr("<description>"),
				AutoShutdownProfile: &armlabservices.AutoShutdownProfile{
					DisconnectDelay:          to.Ptr("<disconnect-delay>"),
					IdleDelay:                to.Ptr("<idle-delay>"),
					NoConnectDelay:           to.Ptr("<no-connect-delay>"),
					ShutdownOnDisconnect:     to.Ptr(armlabservices.EnableStateEnabled),
					ShutdownOnIdle:           to.Ptr(armlabservices.ShutdownOnIdleModeUserAbsence),
					ShutdownWhenNotConnected: to.Ptr(armlabservices.EnableStateEnabled),
				},
				ConnectionProfile: &armlabservices.ConnectionProfile{
					ClientRdpAccess: to.Ptr(armlabservices.ConnectionTypePublic),
					ClientSSHAccess: to.Ptr(armlabservices.ConnectionTypePublic),
					WebRdpAccess:    to.Ptr(armlabservices.ConnectionTypeNone),
					WebSSHAccess:    to.Ptr(armlabservices.ConnectionTypeNone),
				},
				LabPlanID: to.Ptr("<lab-plan-id>"),
				SecurityProfile: &armlabservices.SecurityProfile{
					OpenAccess: to.Ptr(armlabservices.EnableStateDisabled),
				},
				Title: to.Ptr("<title>"),
				VirtualMachineProfile: &armlabservices.VirtualMachineProfile{
					AdditionalCapabilities: &armlabservices.VirtualMachineAdditionalCapabilities{
						InstallGpuDrivers: to.Ptr(armlabservices.EnableStateDisabled),
					},
					AdminUser: &armlabservices.Credentials{
						Username: to.Ptr("<username>"),
					},
					CreateOption: to.Ptr(armlabservices.CreateOptionTemplateVM),
					ImageReference: &armlabservices.ImageReference{
						Offer:     to.Ptr("<offer>"),
						Publisher: to.Ptr("<publisher>"),
						SKU:       to.Ptr("<sku>"),
						Version:   to.Ptr("<version>"),
					},
					SKU: &armlabservices.SKU{
						Name: to.Ptr("<name>"),
					},
					UsageQuota:        to.Ptr("<usage-quota>"),
					UseSharedPassword: to.Ptr(armlabservices.EnableStateDisabled),
				},
				NetworkProfile: &armlabservices.LabNetworkProfile{
					SubnetID: to.Ptr("<subnet-id>"),
				},
				State: to.Ptr(armlabservices.LabStateDraft),
			},
		},
		&armlabservices.LabsClientBeginCreateOrUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/labservices/resource-manager/Microsoft.LabServices/preview/2021-11-15-preview/examples/Labs/patchLab.json
func ExampleLabsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armlabservices.NewLabsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<lab-name>",
		armlabservices.LabUpdate{
			Properties: &armlabservices.LabUpdateProperties{
				SecurityProfile: &armlabservices.SecurityProfile{
					OpenAccess: to.Ptr(armlabservices.EnableStateEnabled),
				},
			},
		},
		&armlabservices.LabsClientBeginUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/labservices/resource-manager/Microsoft.LabServices/preview/2021-11-15-preview/examples/Labs/deleteLab.json
func ExampleLabsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armlabservices.NewLabsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<lab-name>",
		&armlabservices.LabsClientBeginDeleteOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/labservices/resource-manager/Microsoft.LabServices/preview/2021-11-15-preview/examples/Labs/publishLab.json
func ExampleLabsClient_BeginPublish() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armlabservices.NewLabsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginPublish(ctx,
		"<resource-group-name>",
		"<lab-name>",
		&armlabservices.LabsClientBeginPublishOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/labservices/resource-manager/Microsoft.LabServices/preview/2021-11-15-preview/examples/Labs/syncLab.json
func ExampleLabsClient_BeginSyncGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armlabservices.NewLabsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginSyncGroup(ctx,
		"<resource-group-name>",
		"<lab-name>",
		&armlabservices.LabsClientBeginSyncGroupOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
}
