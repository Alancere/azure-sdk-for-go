//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement"
)

// x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementListTemplates.json
func ExampleEmailTemplateClient_ListByService() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armapimanagement.NewEmailTemplateClient("<subscription-id>", cred, nil)
	pager := client.ListByService("<resource-group-name>",
		"<service-name>",
		&armapimanagement.EmailTemplateClientListByServiceOptions{Filter: nil,
			Top:  nil,
			Skip: nil,
		})
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

// x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementHeadEmailTemplate.json
func ExampleEmailTemplateClient_GetEntityTag() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armapimanagement.NewEmailTemplateClient("<subscription-id>", cred, nil)
	_, err = client.GetEntityTag(ctx,
		"<resource-group-name>",
		"<service-name>",
		armapimanagement.TemplateName("newIssueNotificationMessage"),
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementGetTemplate.json
func ExampleEmailTemplateClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armapimanagement.NewEmailTemplateClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<service-name>",
		armapimanagement.TemplateName("newIssueNotificationMessage"),
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.EmailTemplateClientGetResult)
}

// x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateTemplate.json
func ExampleEmailTemplateClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armapimanagement.NewEmailTemplateClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<service-name>",
		armapimanagement.TemplateName("newIssueNotificationMessage"),
		armapimanagement.EmailTemplateUpdateParameters{
			Properties: &armapimanagement.EmailTemplateUpdateParameterProperties{
				Subject: to.StringPtr("<subject>"),
			},
		},
		&armapimanagement.EmailTemplateClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.EmailTemplateClientCreateOrUpdateResult)
}

// x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementUpdateTemplate.json
func ExampleEmailTemplateClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armapimanagement.NewEmailTemplateClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<service-name>",
		armapimanagement.TemplateName("newIssueNotificationMessage"),
		"<if-match>",
		armapimanagement.EmailTemplateUpdateParameters{
			Properties: &armapimanagement.EmailTemplateUpdateParameterProperties{
				Body:    to.StringPtr("<body>"),
				Subject: to.StringPtr("<subject>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.EmailTemplateClientUpdateResult)
}

// x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementDeleteTemplate.json
func ExampleEmailTemplateClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armapimanagement.NewEmailTemplateClient("<subscription-id>", cred, nil)
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<service-name>",
		armapimanagement.TemplateName("newIssueNotificationMessage"),
		"<if-match>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
