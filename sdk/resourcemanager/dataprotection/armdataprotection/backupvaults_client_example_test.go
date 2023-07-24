//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdataprotection_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c53808ba54beef57059371708f1fa6949a11a280/specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2023-05-01/examples/VaultCRUD/GetBackupVaultsInSubscription.json
func ExampleBackupVaultsClient_NewGetInSubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataprotection.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewBackupVaultsClient().NewGetInSubscriptionPager(nil)
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
		// page.BackupVaultResourceList = armdataprotection.BackupVaultResourceList{
		// 	Value: []*armdataprotection.BackupVaultResource{
		// 		{
		// 			Name: to.Ptr("ExampleVault1"),
		// 			Type: to.Ptr("Microsoft.DataProtection/BackupVaults"),
		// 			ID: to.Ptr("/subscriptions/0b352192-dcac-4cc7-992e-a96190ccc68c/resourceGroups/SampleResourceGroup2/providers/Microsoft.DataProtection/BackupVaults/ExampleVault1"),
		// 			Location: to.Ptr("WestUS"),
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("val1"),
		// 			},
		// 			Identity: &armdataprotection.DppIdentityDetails{
		// 				Type: to.Ptr("None"),
		// 			},
		// 			Properties: &armdataprotection.BackupVault{
		// 				FeatureSettings: &armdataprotection.FeatureSettings{
		// 					CrossRegionRestoreSettings: &armdataprotection.CrossRegionRestoreSettings{
		// 						State: to.Ptr(armdataprotection.CrossRegionRestoreStateEnabled),
		// 					},
		// 				},
		// 				ProvisioningState: to.Ptr(armdataprotection.ProvisioningStateSucceeded),
		// 				SecureScore: to.Ptr(armdataprotection.SecureScoreLevelAdequate),
		// 				SecuritySettings: &armdataprotection.SecuritySettings{
		// 					SoftDeleteSettings: &armdataprotection.SoftDeleteSettings{
		// 						RetentionDurationInDays: to.Ptr[float64](14),
		// 						State: to.Ptr(armdataprotection.SoftDeleteState("Enabled")),
		// 					},
		// 				},
		// 				StorageSettings: []*armdataprotection.StorageSetting{
		// 					{
		// 						Type: to.Ptr(armdataprotection.StorageSettingTypesLocallyRedundant),
		// 						DatastoreType: to.Ptr(armdataprotection.StorageSettingStoreTypesVaultStore),
		// 				}},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("ExampleVault2"),
		// 			Type: to.Ptr("Microsoft.DataProtection/BackupVaults"),
		// 			ID: to.Ptr("/subscriptions/0b352192-dcac-4cc7-992e-a96190ccc68c/resourceGroups/SampleResourceGroup/providers/Microsoft.DataProtection/BackupVaults/ExampleVault2"),
		// 			Location: to.Ptr("WestUS"),
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("val1"),
		// 			},
		// 			Identity: &armdataprotection.DppIdentityDetails{
		// 				Type: to.Ptr("None"),
		// 			},
		// 			Properties: &armdataprotection.BackupVault{
		// 				FeatureSettings: &armdataprotection.FeatureSettings{
		// 					CrossRegionRestoreSettings: &armdataprotection.CrossRegionRestoreSettings{
		// 						State: to.Ptr(armdataprotection.CrossRegionRestoreStateEnabled),
		// 					},
		// 				},
		// 				ProvisioningState: to.Ptr(armdataprotection.ProvisioningStateSucceeded),
		// 				SecureScore: to.Ptr(armdataprotection.SecureScoreLevelAdequate),
		// 				SecuritySettings: &armdataprotection.SecuritySettings{
		// 					SoftDeleteSettings: &armdataprotection.SoftDeleteSettings{
		// 						RetentionDurationInDays: to.Ptr[float64](14),
		// 						State: to.Ptr(armdataprotection.SoftDeleteState("Enabled")),
		// 					},
		// 				},
		// 				StorageSettings: []*armdataprotection.StorageSetting{
		// 					{
		// 						Type: to.Ptr(armdataprotection.StorageSettingTypesLocallyRedundant),
		// 						DatastoreType: to.Ptr(armdataprotection.StorageSettingStoreTypesVaultStore),
		// 				}},
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c53808ba54beef57059371708f1fa6949a11a280/specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2023-05-01/examples/VaultCRUD/GetBackupVaultsInResourceGroup.json
func ExampleBackupVaultsClient_NewGetInResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataprotection.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewBackupVaultsClient().NewGetInResourceGroupPager("SampleResourceGroup", nil)
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
		// page.BackupVaultResourceList = armdataprotection.BackupVaultResourceList{
		// 	Value: []*armdataprotection.BackupVaultResource{
		// 		{
		// 			Name: to.Ptr("ExampleVault1"),
		// 			Type: to.Ptr("Microsoft.DataProtection/BackupVaults"),
		// 			ID: to.Ptr("/subscriptions/0b352192-dcac-4cc7-992e-a96190ccc68c/resourceGroups/SampleResourceGroup/providers/Microsoft.DataProtection/BackupVaults/ExampleVault1"),
		// 			Location: to.Ptr("WestUS"),
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("val1"),
		// 			},
		// 			Identity: &armdataprotection.DppIdentityDetails{
		// 				Type: to.Ptr("None"),
		// 			},
		// 			Properties: &armdataprotection.BackupVault{
		// 				FeatureSettings: &armdataprotection.FeatureSettings{
		// 					CrossRegionRestoreSettings: &armdataprotection.CrossRegionRestoreSettings{
		// 						State: to.Ptr(armdataprotection.CrossRegionRestoreStateEnabled),
		// 					},
		// 				},
		// 				ProvisioningState: to.Ptr(armdataprotection.ProvisioningStateSucceeded),
		// 				SecureScore: to.Ptr(armdataprotection.SecureScoreLevelAdequate),
		// 				SecuritySettings: &armdataprotection.SecuritySettings{
		// 					SoftDeleteSettings: &armdataprotection.SoftDeleteSettings{
		// 						RetentionDurationInDays: to.Ptr[float64](14),
		// 						State: to.Ptr(armdataprotection.SoftDeleteState("Enabled")),
		// 					},
		// 				},
		// 				StorageSettings: []*armdataprotection.StorageSetting{
		// 					{
		// 						Type: to.Ptr(armdataprotection.StorageSettingTypesLocallyRedundant),
		// 						DatastoreType: to.Ptr(armdataprotection.StorageSettingStoreTypesVaultStore),
		// 				}},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("ExampleVault2"),
		// 			Type: to.Ptr("Microsoft.DataProtection/BackupVaults"),
		// 			ID: to.Ptr("/subscriptions/0b352192-dcac-4cc7-992e-a96190ccc68c/resourceGroups/SampleResourceGroup/providers/Microsoft.DataProtection/BackupVaults/ExampleVault2"),
		// 			Location: to.Ptr("WestUS"),
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("val1"),
		// 			},
		// 			Identity: &armdataprotection.DppIdentityDetails{
		// 				Type: to.Ptr("SystemAssigned"),
		// 				PrincipalID: to.Ptr("c009b9a0-0024-417c-83cd-025d3776045d"),
		// 				TenantID: to.Ptr("83abe5cd-bcc3-441a-bd86-e6a75360cecc"),
		// 			},
		// 			Properties: &armdataprotection.BackupVault{
		// 				FeatureSettings: &armdataprotection.FeatureSettings{
		// 					CrossRegionRestoreSettings: &armdataprotection.CrossRegionRestoreSettings{
		// 						State: to.Ptr(armdataprotection.CrossRegionRestoreStateEnabled),
		// 					},
		// 				},
		// 				MonitoringSettings: &armdataprotection.MonitoringSettings{
		// 					AzureMonitorAlertSettings: &armdataprotection.AzureMonitorAlertSettings{
		// 						AlertsForAllJobFailures: to.Ptr(armdataprotection.AlertsStateEnabled),
		// 					},
		// 				},
		// 				ProvisioningState: to.Ptr(armdataprotection.ProvisioningStateSucceeded),
		// 				SecureScore: to.Ptr(armdataprotection.SecureScoreLevelAdequate),
		// 				SecuritySettings: &armdataprotection.SecuritySettings{
		// 					SoftDeleteSettings: &armdataprotection.SoftDeleteSettings{
		// 						RetentionDurationInDays: to.Ptr[float64](14),
		// 						State: to.Ptr(armdataprotection.SoftDeleteState("Enabled")),
		// 					},
		// 				},
		// 				StorageSettings: []*armdataprotection.StorageSetting{
		// 					{
		// 						Type: to.Ptr(armdataprotection.StorageSettingTypesLocallyRedundant),
		// 						DatastoreType: to.Ptr(armdataprotection.StorageSettingStoreTypesVaultStore),
		// 				}},
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c53808ba54beef57059371708f1fa6949a11a280/specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2023-05-01/examples/VaultCRUD/GetBackupVault.json
func ExampleBackupVaultsClient_Get_getBackupVault() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataprotection.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBackupVaultsClient().Get(ctx, "SampleResourceGroup", "swaggerExample", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.BackupVaultResource = armdataprotection.BackupVaultResource{
	// 	Name: to.Ptr("swaggerExample"),
	// 	Type: to.Ptr("Microsoft.DataProtection/Backupvaults"),
	// 	ID: to.Ptr("/subscriptions/0b352192-dcac-4cc7-992e-a96190ccc68c/resourceGroups/SampleResourceGroup/providers/Microsoft.DataProtection/Backupvaults/swaggerExample"),
	// 	Location: to.Ptr("WestUS"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("val1"),
	// 	},
	// 	Identity: &armdataprotection.DppIdentityDetails{
	// 		Type: to.Ptr("None"),
	// 	},
	// 	Properties: &armdataprotection.BackupVault{
	// 		FeatureSettings: &armdataprotection.FeatureSettings{
	// 			CrossRegionRestoreSettings: &armdataprotection.CrossRegionRestoreSettings{
	// 				State: to.Ptr(armdataprotection.CrossRegionRestoreStateEnabled),
	// 			},
	// 		},
	// 		MonitoringSettings: &armdataprotection.MonitoringSettings{
	// 			AzureMonitorAlertSettings: &armdataprotection.AzureMonitorAlertSettings{
	// 				AlertsForAllJobFailures: to.Ptr(armdataprotection.AlertsStateEnabled),
	// 			},
	// 		},
	// 		ProvisioningState: to.Ptr(armdataprotection.ProvisioningStateSucceeded),
	// 		SecureScore: to.Ptr(armdataprotection.SecureScoreLevelAdequate),
	// 		SecuritySettings: &armdataprotection.SecuritySettings{
	// 			SoftDeleteSettings: &armdataprotection.SoftDeleteSettings{
	// 				RetentionDurationInDays: to.Ptr[float64](14),
	// 				State: to.Ptr(armdataprotection.SoftDeleteState("Enabled")),
	// 			},
	// 		},
	// 		StorageSettings: []*armdataprotection.StorageSetting{
	// 			{
	// 				Type: to.Ptr(armdataprotection.StorageSettingTypesLocallyRedundant),
	// 				DatastoreType: to.Ptr(armdataprotection.StorageSettingStoreTypesVaultStore),
	// 		}},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c53808ba54beef57059371708f1fa6949a11a280/specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2023-05-01/examples/VaultCRUD/GetBackupVaultWithMSI.json
func ExampleBackupVaultsClient_Get_getBackupVaultWithMsi() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataprotection.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBackupVaultsClient().Get(ctx, "SampleResourceGroup", "swaggerExample", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.BackupVaultResource = armdataprotection.BackupVaultResource{
	// 	Name: to.Ptr("swaggerExample"),
	// 	Type: to.Ptr("Microsoft.DataProtection/Backupvaults"),
	// 	ID: to.Ptr("/subscriptions/0b352192-dcac-4cc7-992e-a96190ccc68c/resourceGroups/SampleResourceGroup/providers/Microsoft.DataProtection/Backupvaults/swaggerExample"),
	// 	Location: to.Ptr("WestUS"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("val1"),
	// 	},
	// 	Identity: &armdataprotection.DppIdentityDetails{
	// 		Type: to.Ptr("SystemAssigned"),
	// 		PrincipalID: to.Ptr("c009b9a0-0024-417c-83cd-025d3776045d"),
	// 		TenantID: to.Ptr("83abe5cd-bcc3-441a-bd86-e6a75360cecc"),
	// 	},
	// 	Properties: &armdataprotection.BackupVault{
	// 		FeatureSettings: &armdataprotection.FeatureSettings{
	// 			CrossRegionRestoreSettings: &armdataprotection.CrossRegionRestoreSettings{
	// 				State: to.Ptr(armdataprotection.CrossRegionRestoreStateEnabled),
	// 			},
	// 		},
	// 		ProvisioningState: to.Ptr(armdataprotection.ProvisioningStateSucceeded),
	// 		SecureScore: to.Ptr(armdataprotection.SecureScoreLevelAdequate),
	// 		SecuritySettings: &armdataprotection.SecuritySettings{
	// 			SoftDeleteSettings: &armdataprotection.SoftDeleteSettings{
	// 				RetentionDurationInDays: to.Ptr[float64](14),
	// 				State: to.Ptr(armdataprotection.SoftDeleteState("Enabled")),
	// 			},
	// 		},
	// 		StorageSettings: []*armdataprotection.StorageSetting{
	// 			{
	// 				Type: to.Ptr(armdataprotection.StorageSettingTypesLocallyRedundant),
	// 				DatastoreType: to.Ptr(armdataprotection.StorageSettingStoreTypesVaultStore),
	// 		}},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c53808ba54beef57059371708f1fa6949a11a280/specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2023-05-01/examples/VaultCRUD/PutBackupVault.json
func ExampleBackupVaultsClient_BeginCreateOrUpdate_createBackupVault() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataprotection.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBackupVaultsClient().BeginCreateOrUpdate(ctx, "SampleResourceGroup", "swaggerExample", armdataprotection.BackupVaultResource{
		Location: to.Ptr("WestUS"),
		Tags: map[string]*string{
			"key1": to.Ptr("val1"),
		},
		Identity: &armdataprotection.DppIdentityDetails{
			Type: to.Ptr("None"),
		},
		Properties: &armdataprotection.BackupVault{
			FeatureSettings: &armdataprotection.FeatureSettings{
				CrossRegionRestoreSettings: &armdataprotection.CrossRegionRestoreSettings{
					State: to.Ptr(armdataprotection.CrossRegionRestoreStateEnabled),
				},
			},
			MonitoringSettings: &armdataprotection.MonitoringSettings{
				AzureMonitorAlertSettings: &armdataprotection.AzureMonitorAlertSettings{
					AlertsForAllJobFailures: to.Ptr(armdataprotection.AlertsStateEnabled),
				},
			},
			SecuritySettings: &armdataprotection.SecuritySettings{
				SoftDeleteSettings: &armdataprotection.SoftDeleteSettings{
					RetentionDurationInDays: to.Ptr[float64](14),
					State:                   to.Ptr(armdataprotection.SoftDeleteState("Enabled")),
				},
			},
			StorageSettings: []*armdataprotection.StorageSetting{
				{
					Type:          to.Ptr(armdataprotection.StorageSettingTypesLocallyRedundant),
					DatastoreType: to.Ptr(armdataprotection.StorageSettingStoreTypesVaultStore),
				}},
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
	// res.BackupVaultResource = armdataprotection.BackupVaultResource{
	// 	Name: to.Ptr("swaggerExample"),
	// 	Type: to.Ptr("Microsoft.DataProtection/Backupvaults"),
	// 	ID: to.Ptr("/subscriptions/0b352192-dcac-4cc7-992e-a96190ccc68c/resourceGroups/SampleResourceGroup/providers/Microsoft.DataProtection/Backupvaults/swaggerExample"),
	// 	Location: to.Ptr("WestUS"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("val1"),
	// 	},
	// 	Identity: &armdataprotection.DppIdentityDetails{
	// 		Type: to.Ptr("None"),
	// 	},
	// 	Properties: &armdataprotection.BackupVault{
	// 		FeatureSettings: &armdataprotection.FeatureSettings{
	// 			CrossRegionRestoreSettings: &armdataprotection.CrossRegionRestoreSettings{
	// 				State: to.Ptr(armdataprotection.CrossRegionRestoreStateEnabled),
	// 			},
	// 		},
	// 		MonitoringSettings: &armdataprotection.MonitoringSettings{
	// 			AzureMonitorAlertSettings: &armdataprotection.AzureMonitorAlertSettings{
	// 				AlertsForAllJobFailures: to.Ptr(armdataprotection.AlertsStateEnabled),
	// 			},
	// 		},
	// 		ProvisioningState: to.Ptr(armdataprotection.ProvisioningStateSucceeded),
	// 		SecureScore: to.Ptr(armdataprotection.SecureScoreLevelAdequate),
	// 		SecuritySettings: &armdataprotection.SecuritySettings{
	// 			SoftDeleteSettings: &armdataprotection.SoftDeleteSettings{
	// 				RetentionDurationInDays: to.Ptr[float64](14),
	// 				State: to.Ptr(armdataprotection.SoftDeleteState("Enabled")),
	// 			},
	// 		},
	// 		StorageSettings: []*armdataprotection.StorageSetting{
	// 			{
	// 				Type: to.Ptr(armdataprotection.StorageSettingTypesLocallyRedundant),
	// 				DatastoreType: to.Ptr(armdataprotection.StorageSettingStoreTypesVaultStore),
	// 		}},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c53808ba54beef57059371708f1fa6949a11a280/specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2023-05-01/examples/VaultCRUD/PutBackupVaultWithMSI.json
func ExampleBackupVaultsClient_BeginCreateOrUpdate_createBackupVaultWithMsi() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataprotection.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBackupVaultsClient().BeginCreateOrUpdate(ctx, "SampleResourceGroup", "swaggerExample", armdataprotection.BackupVaultResource{
		Location: to.Ptr("WestUS"),
		Tags: map[string]*string{
			"key1": to.Ptr("val1"),
		},
		Identity: &armdataprotection.DppIdentityDetails{
			Type: to.Ptr("systemAssigned"),
		},
		Properties: &armdataprotection.BackupVault{
			FeatureSettings: &armdataprotection.FeatureSettings{
				CrossRegionRestoreSettings: &armdataprotection.CrossRegionRestoreSettings{
					State: to.Ptr(armdataprotection.CrossRegionRestoreStateEnabled),
				},
			},
			MonitoringSettings: &armdataprotection.MonitoringSettings{
				AzureMonitorAlertSettings: &armdataprotection.AzureMonitorAlertSettings{
					AlertsForAllJobFailures: to.Ptr(armdataprotection.AlertsStateEnabled),
				},
			},
			SecuritySettings: &armdataprotection.SecuritySettings{
				SoftDeleteSettings: &armdataprotection.SoftDeleteSettings{
					RetentionDurationInDays: to.Ptr[float64](14),
					State:                   to.Ptr(armdataprotection.SoftDeleteState("Enabled")),
				},
			},
			StorageSettings: []*armdataprotection.StorageSetting{
				{
					Type:          to.Ptr(armdataprotection.StorageSettingTypesLocallyRedundant),
					DatastoreType: to.Ptr(armdataprotection.StorageSettingStoreTypesVaultStore),
				}},
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
	// res.BackupVaultResource = armdataprotection.BackupVaultResource{
	// 	Name: to.Ptr("swaggerExample"),
	// 	Type: to.Ptr("Microsoft.DataProtection/Backupvaults"),
	// 	ID: to.Ptr("/subscriptions/0b352192-dcac-4cc7-992e-a96190ccc68c/resourceGroups/SampleResourceGroup/providers/Microsoft.DataProtection/Backupvaults/swaggerExample"),
	// 	Location: to.Ptr("WestUS"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("val1"),
	// 	},
	// 	Identity: &armdataprotection.DppIdentityDetails{
	// 		Type: to.Ptr("SystemAssigned"),
	// 		PrincipalID: to.Ptr("c009b9a0-0024-417c-83cd-025d3776045d"),
	// 		TenantID: to.Ptr("83abe5cd-bcc3-441a-bd86-e6a75360cecc"),
	// 	},
	// 	Properties: &armdataprotection.BackupVault{
	// 		FeatureSettings: &armdataprotection.FeatureSettings{
	// 			CrossRegionRestoreSettings: &armdataprotection.CrossRegionRestoreSettings{
	// 				State: to.Ptr(armdataprotection.CrossRegionRestoreStateEnabled),
	// 			},
	// 		},
	// 		MonitoringSettings: &armdataprotection.MonitoringSettings{
	// 			AzureMonitorAlertSettings: &armdataprotection.AzureMonitorAlertSettings{
	// 				AlertsForAllJobFailures: to.Ptr(armdataprotection.AlertsStateEnabled),
	// 			},
	// 		},
	// 		ProvisioningState: to.Ptr(armdataprotection.ProvisioningStateSucceeded),
	// 		SecureScore: to.Ptr(armdataprotection.SecureScoreLevelAdequate),
	// 		SecuritySettings: &armdataprotection.SecuritySettings{
	// 			SoftDeleteSettings: &armdataprotection.SoftDeleteSettings{
	// 				RetentionDurationInDays: to.Ptr[float64](14),
	// 				State: to.Ptr(armdataprotection.SoftDeleteState("Enabled")),
	// 			},
	// 		},
	// 		StorageSettings: []*armdataprotection.StorageSetting{
	// 			{
	// 				Type: to.Ptr(armdataprotection.StorageSettingTypesLocallyRedundant),
	// 				DatastoreType: to.Ptr(armdataprotection.StorageSettingStoreTypesVaultStore),
	// 		}},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c53808ba54beef57059371708f1fa6949a11a280/specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2023-05-01/examples/VaultCRUD/DeleteBackupVault.json
func ExampleBackupVaultsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataprotection.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBackupVaultsClient().BeginDelete(ctx, "SampleResourceGroup", "swaggerExample", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c53808ba54beef57059371708f1fa6949a11a280/specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2023-05-01/examples/VaultCRUD/PatchBackupVault.json
func ExampleBackupVaultsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataprotection.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBackupVaultsClient().BeginUpdate(ctx, "SampleResourceGroup", "swaggerExample", armdataprotection.PatchResourceRequestInput{
		Properties: &armdataprotection.PatchBackupVaultInput{
			MonitoringSettings: &armdataprotection.MonitoringSettings{
				AzureMonitorAlertSettings: &armdataprotection.AzureMonitorAlertSettings{
					AlertsForAllJobFailures: to.Ptr(armdataprotection.AlertsStateEnabled),
				},
			},
		},
		Tags: map[string]*string{
			"newKey": to.Ptr("newVal"),
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
	// res.BackupVaultResource = armdataprotection.BackupVaultResource{
	// 	Name: to.Ptr("swaggerExample"),
	// 	Type: to.Ptr("Microsoft.DataProtection/Backupvaults"),
	// 	ID: to.Ptr("/subscriptions/0b352192-dcac-4cc7-992e-a96190ccc68c/resourceGroups/SampleResourceGroup/providers/Microsoft.DataProtection/Backupvaults/swaggerExample"),
	// 	Location: to.Ptr("WestUS"),
	// 	Tags: map[string]*string{
	// 		"newKey": to.Ptr("newVal"),
	// 	},
	// 	Properties: &armdataprotection.BackupVault{
	// 		MonitoringSettings: &armdataprotection.MonitoringSettings{
	// 			AzureMonitorAlertSettings: &armdataprotection.AzureMonitorAlertSettings{
	// 				AlertsForAllJobFailures: to.Ptr(armdataprotection.AlertsStateEnabled),
	// 			},
	// 		},
	// 		ProvisioningState: to.Ptr(armdataprotection.ProvisioningStateSucceeded),
	// 		StorageSettings: []*armdataprotection.StorageSetting{
	// 			{
	// 				Type: to.Ptr(armdataprotection.StorageSettingTypesLocallyRedundant),
	// 				DatastoreType: to.Ptr(armdataprotection.StorageSettingStoreTypesVaultStore),
	// 		}},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c53808ba54beef57059371708f1fa6949a11a280/specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2023-05-01/examples/VaultCRUD/CheckBackupVaultsNameAvailability.json
func ExampleBackupVaultsClient_CheckNameAvailability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataprotection.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBackupVaultsClient().CheckNameAvailability(ctx, "SampleResourceGroup", "westus", armdataprotection.CheckNameAvailabilityRequest{
		Name: to.Ptr("swaggerExample"),
		Type: to.Ptr("Microsoft.DataProtection/BackupVaults"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CheckNameAvailabilityResult = armdataprotection.CheckNameAvailabilityResult{
	// 	NameAvailable: to.Ptr(true),
	// }
}
