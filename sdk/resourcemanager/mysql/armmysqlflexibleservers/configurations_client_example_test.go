//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmysqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysqlflexibleservers/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2021-05-01/examples/ConfigurationUpdate.json
func ExampleConfigurationsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmysqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewConfigurationsClient().BeginUpdate(ctx, "testrg", "testserver", "event_scheduler", armmysqlflexibleservers.Configuration{
		Properties: &armmysqlflexibleservers.ConfigurationProperties{
			Source: to.Ptr(armmysqlflexibleservers.ConfigurationSourceUserOverride),
			Value:  to.Ptr("on"),
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
	// res.Configuration = armmysqlflexibleservers.Configuration{
	// 	Name: to.Ptr("event_scheduler"),
	// 	Type: to.Ptr("Microsoft.DBforMySQL/flexibleServers/configurations"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforMySQL/flexibleServers/testserver/configurations/event_scheduler"),
	// 	Properties: &armmysqlflexibleservers.ConfigurationProperties{
	// 		Description: to.Ptr("Indicates the status of the Event Scheduler. It is always OFF for a replica server to keep the replication consistency."),
	// 		AllowedValues: to.Ptr("ON,OFF"),
	// 		DataType: to.Ptr("Enumeration"),
	// 		DefaultValue: to.Ptr("OFF"),
	// 		IsConfigPendingRestart: to.Ptr(armmysqlflexibleservers.IsConfigPendingRestartFalse),
	// 		IsDynamicConfig: to.Ptr(armmysqlflexibleservers.IsDynamicConfigTrue),
	// 		IsReadOnly: to.Ptr(armmysqlflexibleservers.IsReadOnlyFalse),
	// 		Source: to.Ptr(armmysqlflexibleservers.ConfigurationSourceUserOverride),
	// 		Value: to.Ptr("ON"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2021-05-01/examples/ConfigurationGet.json
func ExampleConfigurationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmysqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewConfigurationsClient().Get(ctx, "TestGroup", "testserver", "event_scheduler", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Configuration = armmysqlflexibleservers.Configuration{
	// 	Name: to.Ptr("event_scheduler"),
	// 	Type: to.Ptr("Microsoft.DBforMySQL/flexibleServers/configurations"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforMySQL/flexibleServers/testserver/configurations/event_scheduler"),
	// 	Properties: &armmysqlflexibleservers.ConfigurationProperties{
	// 		Description: to.Ptr("Indicates the status of the Event Scheduler. It is always OFF for a replica server to keep the replication consistency."),
	// 		AllowedValues: to.Ptr("ON,OFF"),
	// 		DataType: to.Ptr("Enumeration"),
	// 		DefaultValue: to.Ptr("OFF"),
	// 		IsConfigPendingRestart: to.Ptr(armmysqlflexibleservers.IsConfigPendingRestartFalse),
	// 		IsDynamicConfig: to.Ptr(armmysqlflexibleservers.IsDynamicConfigTrue),
	// 		IsReadOnly: to.Ptr(armmysqlflexibleservers.IsReadOnlyFalse),
	// 		Source: to.Ptr(armmysqlflexibleservers.ConfigurationSourceSystemDefault),
	// 		Value: to.Ptr("OFF"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2021-05-01/examples/ConfigurationsBatchUpdate.json
func ExampleConfigurationsClient_BeginBatchUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmysqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewConfigurationsClient().BeginBatchUpdate(ctx, "testrg", "mysqltestserver", armmysqlflexibleservers.ConfigurationListForBatchUpdate{
		Value: []*armmysqlflexibleservers.ConfigurationForBatchUpdate{
			{
				Name: to.Ptr("event_scheduler"),
				Properties: &armmysqlflexibleservers.ConfigurationForBatchUpdateProperties{
					Value: to.Ptr("OFF"),
				},
			},
			{
				Name: to.Ptr("div_precision_increment"),
				Properties: &armmysqlflexibleservers.ConfigurationForBatchUpdateProperties{
					Value: to.Ptr("8"),
				},
			}},
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
	// res.ConfigurationListResult = armmysqlflexibleservers.ConfigurationListResult{
	// 	Value: []*armmysqlflexibleservers.Configuration{
	// 		{
	// 			Name: to.Ptr("event_scheduler"),
	// 			Type: to.Ptr("Microsoft.DBforMySQL/flexibleServers/configurations"),
	// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/flexibleServers/mysqltestserver/configurations/event_scheduler"),
	// 			Properties: &armmysqlflexibleservers.ConfigurationProperties{
	// 				Description: to.Ptr("Indicates the status of the Event Scheduler. It is always OFF for a replica server to keep the replication consistency."),
	// 				AllowedValues: to.Ptr("ON,OFF"),
	// 				DataType: to.Ptr("Enumeration"),
	// 				DefaultValue: to.Ptr("OFF"),
	// 				IsConfigPendingRestart: to.Ptr(armmysqlflexibleservers.IsConfigPendingRestartFalse),
	// 				IsDynamicConfig: to.Ptr(armmysqlflexibleservers.IsDynamicConfigTrue),
	// 				IsReadOnly: to.Ptr(armmysqlflexibleservers.IsReadOnlyFalse),
	// 				Source: to.Ptr(armmysqlflexibleservers.ConfigurationSourceUserOverride),
	// 				Value: to.Ptr("ON"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("div_precision_increment"),
	// 			Type: to.Ptr("Microsoft.DBforMySQL/flexibleServers/configurations"),
	// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/flexibleServers/mysqltestserver/configurations/div_precision_increment"),
	// 			Properties: &armmysqlflexibleservers.ConfigurationProperties{
	// 				Description: to.Ptr("Number of digits by which to increase the scale of the result of division operations."),
	// 				AllowedValues: to.Ptr("0-30"),
	// 				DataType: to.Ptr("Integer"),
	// 				DefaultValue: to.Ptr("4"),
	// 				IsConfigPendingRestart: to.Ptr(armmysqlflexibleservers.IsConfigPendingRestartFalse),
	// 				IsDynamicConfig: to.Ptr(armmysqlflexibleservers.IsDynamicConfigTrue),
	// 				IsReadOnly: to.Ptr(armmysqlflexibleservers.IsReadOnlyFalse),
	// 				Source: to.Ptr(armmysqlflexibleservers.ConfigurationSourceUserOverride),
	// 				Value: to.Ptr("8"),
	// 			},
	// 	}},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2021-05-01/examples/ConfigurationsListByServer.json
func ExampleConfigurationsClient_NewListByServerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmysqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewConfigurationsClient().NewListByServerPager("testrg", "mysqltestserver", nil)
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
		// page.ConfigurationListResult = armmysqlflexibleservers.ConfigurationListResult{
		// 	Value: []*armmysqlflexibleservers.Configuration{
		// 		{
		// 			Name: to.Ptr("archive"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/flexibleServers/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/flexibleServers/mysqltestserver/configurations/archive"),
		// 			Properties: &armmysqlflexibleservers.ConfigurationProperties{
		// 				Description: to.Ptr("Tell the server to enable or disable archive engine."),
		// 				AllowedValues: to.Ptr("ON,OFF"),
		// 				DataType: to.Ptr("Enumeration"),
		// 				DefaultValue: to.Ptr("OFF"),
		// 				IsConfigPendingRestart: to.Ptr(armmysqlflexibleservers.IsConfigPendingRestartFalse),
		// 				IsDynamicConfig: to.Ptr(armmysqlflexibleservers.IsDynamicConfigFalse),
		// 				IsReadOnly: to.Ptr(armmysqlflexibleservers.IsReadOnlyTrue),
		// 				Source: to.Ptr(armmysqlflexibleservers.ConfigurationSourceSystemDefault),
		// 				Value: to.Ptr("OFF"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("audit_log_enabled"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/flexibleServers/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/flexibleServers/mysqltestserver/configurations/audit_log_enabled"),
		// 			Properties: &armmysqlflexibleservers.ConfigurationProperties{
		// 				Description: to.Ptr("Allow to audit the log."),
		// 				AllowedValues: to.Ptr("ON,OFF"),
		// 				DataType: to.Ptr("Enumeration"),
		// 				DefaultValue: to.Ptr("OFF"),
		// 				IsConfigPendingRestart: to.Ptr(armmysqlflexibleservers.IsConfigPendingRestartFalse),
		// 				IsDynamicConfig: to.Ptr(armmysqlflexibleservers.IsDynamicConfigTrue),
		// 				IsReadOnly: to.Ptr(armmysqlflexibleservers.IsReadOnlyFalse),
		// 				Source: to.Ptr(armmysqlflexibleservers.ConfigurationSourceSystemDefault),
		// 				Value: to.Ptr("OFF"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("audit_log_events"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/flexibleServers/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/flexibleServers/mysqltestserver/configurations/audit_log_events"),
		// 			Properties: &armmysqlflexibleservers.ConfigurationProperties{
		// 				Description: to.Ptr("Select the events to audit logs."),
		// 				AllowedValues: to.Ptr("DDL,DML_SELECT,DML_NONSELECT,DCL,ADMIN,DML,GENERAL,CONNECTION,TABLE_ACCESS"),
		// 				DataType: to.Ptr("Set"),
		// 				DefaultValue: to.Ptr("CONNECTION"),
		// 				IsConfigPendingRestart: to.Ptr(armmysqlflexibleservers.IsConfigPendingRestartFalse),
		// 				IsDynamicConfig: to.Ptr(armmysqlflexibleservers.IsDynamicConfigTrue),
		// 				IsReadOnly: to.Ptr(armmysqlflexibleservers.IsReadOnlyFalse),
		// 				Source: to.Ptr(armmysqlflexibleservers.ConfigurationSourceSystemDefault),
		// 				Value: to.Ptr("CONNECTION"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("audit_log_exclude_users"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/flexibleServers/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/flexibleServers/mysqltestserver/configurations/audit_log_exclude_users"),
		// 			Properties: &armmysqlflexibleservers.ConfigurationProperties{
		// 				Description: to.Ptr("The comma-separated user list whose commands will not be in the audit logs."),
		// 				AllowedValues: to.Ptr(""),
		// 				DataType: to.Ptr("String"),
		// 				DefaultValue: to.Ptr("azure_superuser"),
		// 				IsConfigPendingRestart: to.Ptr(armmysqlflexibleservers.IsConfigPendingRestartFalse),
		// 				IsDynamicConfig: to.Ptr(armmysqlflexibleservers.IsDynamicConfigTrue),
		// 				IsReadOnly: to.Ptr(armmysqlflexibleservers.IsReadOnlyFalse),
		// 				Source: to.Ptr(armmysqlflexibleservers.ConfigurationSourceSystemDefault),
		// 				Value: to.Ptr("azure_superuser"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("audit_log_include_users"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/flexibleServers/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/flexibleServers/mysqltestserver/configurations/audit_log_include_users"),
		// 			Properties: &armmysqlflexibleservers.ConfigurationProperties{
		// 				Description: to.Ptr("The comma-separated user list whose commands will be in the audit logs. It takes higher priority if the same user name is found in audit_log_exclude_users."),
		// 				AllowedValues: to.Ptr(""),
		// 				DataType: to.Ptr("String"),
		// 				DefaultValue: to.Ptr(""),
		// 				IsConfigPendingRestart: to.Ptr(armmysqlflexibleservers.IsConfigPendingRestartFalse),
		// 				IsDynamicConfig: to.Ptr(armmysqlflexibleservers.IsDynamicConfigTrue),
		// 				IsReadOnly: to.Ptr(armmysqlflexibleservers.IsReadOnlyFalse),
		// 				Source: to.Ptr(armmysqlflexibleservers.ConfigurationSourceSystemDefault),
		// 				Value: to.Ptr(""),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("audit_slow_log_enabled"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/flexibleServers/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/flexibleServers/mysqltestserver/configurations/audit_slow_log_enabled"),
		// 			Properties: &armmysqlflexibleservers.ConfigurationProperties{
		// 				Description: to.Ptr("Allow to audit the slow log."),
		// 				AllowedValues: to.Ptr("ON,OFF"),
		// 				DataType: to.Ptr("Enumeration"),
		// 				DefaultValue: to.Ptr("ON"),
		// 				IsConfigPendingRestart: to.Ptr(armmysqlflexibleservers.IsConfigPendingRestartFalse),
		// 				IsDynamicConfig: to.Ptr(armmysqlflexibleservers.IsDynamicConfigTrue),
		// 				IsReadOnly: to.Ptr(armmysqlflexibleservers.IsReadOnlyTrue),
		// 				Source: to.Ptr(armmysqlflexibleservers.ConfigurationSourceSystemDefault),
		// 				Value: to.Ptr("ON"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("auto_generate_certs"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/flexibleServers/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/flexibleServers/mysqltestserver/configurations/auto_generate_certs"),
		// 			Properties: &armmysqlflexibleservers.ConfigurationProperties{
		// 				Description: to.Ptr("Controls whether the server autogenerates SSL key and certificate files in the data directory, if they do not already exist."),
		// 				AllowedValues: to.Ptr("ON,OFF"),
		// 				DataType: to.Ptr("Enumeration"),
		// 				DefaultValue: to.Ptr("OFF"),
		// 				IsConfigPendingRestart: to.Ptr(armmysqlflexibleservers.IsConfigPendingRestartFalse),
		// 				IsDynamicConfig: to.Ptr(armmysqlflexibleservers.IsDynamicConfigFalse),
		// 				IsReadOnly: to.Ptr(armmysqlflexibleservers.IsReadOnlyTrue),
		// 				Source: to.Ptr(armmysqlflexibleservers.ConfigurationSourceSystemDefault),
		// 				Value: to.Ptr("OFF"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("auto_increment_increment"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/flexibleServers/configurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/flexibleServers/mysqltestserver/configurations/auto_increment_increment"),
		// 			Properties: &armmysqlflexibleservers.ConfigurationProperties{
		// 				Description: to.Ptr("The auto_increment_increment is intended for use with source-to-source replication, and can be used to control the operation of AUTO_INCREMENT columns."),
		// 				AllowedValues: to.Ptr("1-65535"),
		// 				DataType: to.Ptr("Integer"),
		// 				DefaultValue: to.Ptr("1"),
		// 				IsConfigPendingRestart: to.Ptr(armmysqlflexibleservers.IsConfigPendingRestartFalse),
		// 				IsDynamicConfig: to.Ptr(armmysqlflexibleservers.IsDynamicConfigTrue),
		// 				IsReadOnly: to.Ptr(armmysqlflexibleservers.IsReadOnlyFalse),
		// 				Source: to.Ptr(armmysqlflexibleservers.ConfigurationSourceSystemDefault),
		// 				Value: to.Ptr("1"),
		// 			},
		// 	}},
		// }
	}
}
