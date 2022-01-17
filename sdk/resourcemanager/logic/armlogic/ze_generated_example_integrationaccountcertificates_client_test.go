//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armlogic_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic"
)

// x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccountCertificates_List.json
func ExampleIntegrationAccountCertificatesClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armlogic.NewIntegrationAccountCertificatesClient("<subscription-id>", cred, nil)
	pager := client.List("<resource-group-name>",
		"<integration-account-name>",
		&armlogic.IntegrationAccountCertificatesClientListOptions{Top: nil})
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}

// x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccountCertificates_Get.json
func ExampleIntegrationAccountCertificatesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armlogic.NewIntegrationAccountCertificatesClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<integration-account-name>",
		"<certificate-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.IntegrationAccountCertificatesClientGetResult)
}

// x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccountCertificates_CreateOrUpdate.json
func ExampleIntegrationAccountCertificatesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armlogic.NewIntegrationAccountCertificatesClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<integration-account-name>",
		"<certificate-name>",
		armlogic.IntegrationAccountCertificate{
			Location: to.StringPtr("<location>"),
			Properties: &armlogic.IntegrationAccountCertificateProperties{
				Key: &armlogic.KeyVaultKeyReference{
					KeyName: to.StringPtr("<key-name>"),
					KeyVault: &armlogic.KeyVaultKeyReferenceKeyVault{
						ID: to.StringPtr("<id>"),
					},
					KeyVersion: to.StringPtr("<key-version>"),
				},
				PublicCertificate: to.StringPtr("<public-certificate>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.IntegrationAccountCertificatesClientCreateOrUpdateResult)
}

// x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccountCertificates_Delete.json
func ExampleIntegrationAccountCertificatesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armlogic.NewIntegrationAccountCertificatesClient("<subscription-id>", cred, nil)
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<integration-account-name>",
		"<certificate-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
