//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/33c4457b1d13f83965f4fe3367dca4a6df898100/specification/sql/resource-manager/Microsoft.Sql/stable/2014-04-01/examples/DataMaskingPolicyCreateOrUpdateMax.json
func ExampleDataMaskingPoliciesClient_CreateOrUpdate_createOrUpdateDataMaskingPolicyMax() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDataMaskingPoliciesClient().CreateOrUpdate(ctx, "sqlcrudtest-6852", "sqlcrudtest-2080", "sqlcrudtest-331", armsql.DataMaskingPolicy{
		Properties: &armsql.DataMaskingPolicyProperties{
			DataMaskingState: to.Ptr(armsql.DataMaskingStateEnabled),
			ExemptPrincipals: to.Ptr("testuser;"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DataMaskingPolicy = armsql.DataMaskingPolicy{
	// 	Name: to.Ptr("Default"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/databases/dataMaskingPolicies"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/sqlcrudtest-6852/providers/Microsoft.Sql/servers/sqlcrudtest-2080/databases/sqlcrudtest-331/dataMaskingPolicies/Default"),
	// 	Location: to.Ptr("Central US"),
	// 	Properties: &armsql.DataMaskingPolicyProperties{
	// 		ApplicationPrincipals: to.Ptr(""),
	// 		DataMaskingState: to.Ptr(armsql.DataMaskingStateEnabled),
	// 		ExemptPrincipals: to.Ptr("testuser;"),
	// 		MaskingLevel: to.Ptr(""),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/33c4457b1d13f83965f4fe3367dca4a6df898100/specification/sql/resource-manager/Microsoft.Sql/stable/2014-04-01/examples/DataMaskingPolicyCreateOrUpdateMin.json
func ExampleDataMaskingPoliciesClient_CreateOrUpdate_createOrUpdateDataMaskingPolicyMin() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDataMaskingPoliciesClient().CreateOrUpdate(ctx, "sqlcrudtest-6852", "sqlcrudtest-2080", "sqlcrudtest-331", armsql.DataMaskingPolicy{
		Properties: &armsql.DataMaskingPolicyProperties{
			DataMaskingState: to.Ptr(armsql.DataMaskingStateEnabled),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DataMaskingPolicy = armsql.DataMaskingPolicy{
	// 	Name: to.Ptr("Default"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/databases/dataMaskingPolicies"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/sqlcrudtest-6852/providers/Microsoft.Sql/servers/sqlcrudtest-2080/databases/sqlcrudtest-331/dataMaskingPolicies/Default"),
	// 	Location: to.Ptr("Central US"),
	// 	Properties: &armsql.DataMaskingPolicyProperties{
	// 		ApplicationPrincipals: to.Ptr(""),
	// 		DataMaskingState: to.Ptr(armsql.DataMaskingStateEnabled),
	// 		ExemptPrincipals: to.Ptr(""),
	// 		MaskingLevel: to.Ptr(""),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/33c4457b1d13f83965f4fe3367dca4a6df898100/specification/sql/resource-manager/Microsoft.Sql/stable/2014-04-01/examples/DataMaskingPolicyGet.json
func ExampleDataMaskingPoliciesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDataMaskingPoliciesClient().Get(ctx, "sqlcrudtest-6852", "sqlcrudtest-2080", "sqlcrudtest-331", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DataMaskingPolicy = armsql.DataMaskingPolicy{
	// 	Name: to.Ptr("Default"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/databases/dataMaskingPolicies"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/sqlcrudtest-6852/providers/Microsoft.Sql/servers/sqlcrudtest-2080/databases/sqlcrudtest-331/dataMaskingPolicies/Default"),
	// 	Location: to.Ptr("Central US"),
	// 	Properties: &armsql.DataMaskingPolicyProperties{
	// 		ApplicationPrincipals: to.Ptr(""),
	// 		DataMaskingState: to.Ptr(armsql.DataMaskingStateEnabled),
	// 		ExemptPrincipals: to.Ptr(""),
	// 		MaskingLevel: to.Ptr(""),
	// 	},
	// }
}
