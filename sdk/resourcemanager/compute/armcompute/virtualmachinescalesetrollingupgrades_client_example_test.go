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

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-07-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSetRollingUpgrade_Cancel_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetRollingUpgradesClient_BeginCancel_virtualMachineScaleSetRollingUpgradeCancelMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachineScaleSetRollingUpgradesClient().BeginCancel(ctx, "rgcompute", "aaaaa", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-07-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSetRollingUpgrade_Cancel_MinimumSet_Gen.json
func ExampleVirtualMachineScaleSetRollingUpgradesClient_BeginCancel_virtualMachineScaleSetRollingUpgradeCancelMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachineScaleSetRollingUpgradesClient().BeginCancel(ctx, "rgcompute", "aaaaaa", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-07-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSetRollingUpgrade_StartOSUpgrade_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetRollingUpgradesClient_BeginStartOSUpgrade_virtualMachineScaleSetRollingUpgradeStartOsUpgradeMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachineScaleSetRollingUpgradesClient().BeginStartOSUpgrade(ctx, "rgcompute", "aaaa", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-07-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSetRollingUpgrade_StartOSUpgrade_MinimumSet_Gen.json
func ExampleVirtualMachineScaleSetRollingUpgradesClient_BeginStartOSUpgrade_virtualMachineScaleSetRollingUpgradeStartOsUpgradeMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachineScaleSetRollingUpgradesClient().BeginStartOSUpgrade(ctx, "rgcompute", "aaaaaaaaaaaaaaaaaa", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-07-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSetExtension_RollingUpgrade.json
func ExampleVirtualMachineScaleSetRollingUpgradesClient_BeginStartExtensionUpgrade() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachineScaleSetRollingUpgradesClient().BeginStartExtensionUpgrade(ctx, "myResourceGroup", "{vmss-name}", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-07-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSetRollingUpgrade_GetLatest_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetRollingUpgradesClient_GetLatest_virtualMachineScaleSetRollingUpgradeGetLatestMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineScaleSetRollingUpgradesClient().GetLatest(ctx, "rgcompute", "aaaaaaaaaaaaaaaaaaaaaaaaa", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RollingUpgradeStatusInfo = armcompute.RollingUpgradeStatusInfo{
	// 	Name: to.Ptr("aaaaaaaaaaaaaaaaaaa"),
	// 	Type: to.Ptr("aaaaaaaaaaaaaaaaaaaaa"),
	// 	ID: to.Ptr("aaaaaaaaaa"),
	// 	Location: to.Ptr("aaaaaa"),
	// 	Tags: map[string]*string{
	// 		"key8533": to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaa"),
	// 	},
	// 	Properties: &armcompute.RollingUpgradeStatusInfoProperties{
	// 		Error: &armcompute.APIError{
	// 			Code: to.Ptr("aaaaaaa"),
	// 			Innererror: &armcompute.InnerError{
	// 				Errordetail: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaa"),
	// 				Exceptiontype: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 			},
	// 			Message: to.Ptr("aaaaaaaaa"),
	// 			Target: to.Ptr("aaaaaaa"),
	// 			Details: []*armcompute.APIErrorBase{
	// 				{
	// 					Code: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaa"),
	// 					Message: to.Ptr("aa"),
	// 					Target: to.Ptr("aaaa"),
	// 			}},
	// 		},
	// 		Policy: &armcompute.RollingUpgradePolicy{
	// 			EnableCrossZoneUpgrade: to.Ptr(true),
	// 			MaxBatchInstancePercent: to.Ptr[int32](49),
	// 			MaxSurge: to.Ptr(true),
	// 			MaxUnhealthyInstancePercent: to.Ptr[int32](81),
	// 			MaxUnhealthyUpgradedInstancePercent: to.Ptr[int32](98),
	// 			PauseTimeBetweenBatches: to.Ptr("aaaaaaaaaaaaaaa"),
	// 			PrioritizeUnhealthyInstances: to.Ptr(true),
	// 			RollbackFailedInstancesOnPolicyBreach: to.Ptr(true),
	// 		},
	// 		Progress: &armcompute.RollingUpgradeProgressInfo{
	// 			FailedInstanceCount: to.Ptr[int32](25),
	// 			InProgressInstanceCount: to.Ptr[int32](20),
	// 			PendingInstanceCount: to.Ptr[int32](27),
	// 			SuccessfulInstanceCount: to.Ptr[int32](6),
	// 		},
	// 		RunningStatus: &armcompute.RollingUpgradeRunningStatus{
	// 			Code: to.Ptr(armcompute.RollingUpgradeStatusCodeRollingForward),
	// 			LastAction: to.Ptr(armcompute.RollingUpgradeActionTypeStart),
	// 			LastActionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T13:06:23.362Z"); return t}()),
	// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T13:06:23.362Z"); return t}()),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-07-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSetRollingUpgrade_GetLatest_MinimumSet_Gen.json
func ExampleVirtualMachineScaleSetRollingUpgradesClient_GetLatest_virtualMachineScaleSetRollingUpgradeGetLatestMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineScaleSetRollingUpgradesClient().GetLatest(ctx, "rgcompute", "aaaaaaaaaaaaaaaaa", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RollingUpgradeStatusInfo = armcompute.RollingUpgradeStatusInfo{
	// 	ID: to.Ptr("aaaaaaaaaa"),
	// 	Location: to.Ptr("aaaaaa"),
	// }
}
