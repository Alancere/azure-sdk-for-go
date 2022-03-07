//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package publicipaddress_test

import (
	"context"
	"testing"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/internal/recording"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/scenario_test"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
)

var (
	ctx               context.Context
	cred              azcore.TokenCredential
	pathToPackage     = "sdk/resourcemanager/network/armnetwork/scenario_test/publicipaddress/testdata"
	options           *arm.ClientOptions
	resourceGroup     *armresources.ResourceGroup
	resourceName      string
	location          = scenario_test.GetEnv("LOCATION", "eastus")
	resourceGroupName = scenario_test.GetEnv("RESOURCE_GROUP_NAME", "")
	subscriptionId    = scenario_test.GetEnv("SUBSCRIPTION_ID", scenario_test.GetEnv("AZURE_SUBSCRIPTION_ID", ""))
)

func TestPublicipaddress(t *testing.T) {
	// Setup for test
	scenario_test.StartRecording(t, pathToPackage)
	ctx = context.Background()
	options = scenario_test.CreateArmOptions(t)
	cred = scenario_test.CreateCred(t, ctx, options)
	resourceGroup = scenario_test.CreateResourceGroup(t, ctx, cred, subscriptionId, location, options)
	resourceGroupName = *resourceGroup.Name
	// Clenup for test
	t.Cleanup(func() {
		scenario_test.DeleteResourceGroup(t, ctx, cred, subscriptionId, resourceGroupName, options)
		scenario_test.StopRecording(t)
	})
	prepare(t)
	scenarioPublicipaddress(t)
	cleanup(t)
}

func prepare(t *testing.T) {
	// From step Generate_Unique_Name
	{
		template := map[string]interface{}{
			"$schema":        "https://schema.management.azure.com/schemas/2019-04-01/deploymentTemplate.json#",
			"contentVersion": "1.0.0.0",
			"outputs": map[string]interface{}{
				"resourceName": map[string]interface{}{
					"type":  "string",
					"value": "[variables('name').value]",
				},
			},
			"resources": []interface{}{},
			"variables": map[string]interface{}{
				"name": map[string]interface{}{
					"type": "string",
					"metadata": map[string]interface{}{
						"description": "Name of the Resource Group.",
					},
					"value": "[concat('sw',uniqueString(resourceGroup().id))]",
				},
			},
		}
		params := map[string]interface{}{}
		deploymentExtend, err := scenario_test.CreateDeployment(ctx, cred, options, subscriptionId, resourceGroupName, "Generate_Unique_Name", template, params)
		if err != nil {
			t.Fatalf("Deployment error: %v", err)
		}
		resourceName = deploymentExtend.Properties.Outputs["resourceName"].(map[string]interface{})["value"].(string)
	}
}

func scenarioPublicipaddress(t *testing.T) {
	publicIPName := "test-ip"
	// From step PublicIpAddress_CreateDefaults
	publicIPAddressesClient := armnetwork.NewPublicIPAddressesClient(subscriptionId, cred, options)
	{
		publicIPAddressesClientCreateOrUpdatePollerResponse, err := publicIPAddressesClient.BeginCreateOrUpdate(ctx,
			resourceGroupName,
			"test-ip",
			armnetwork.PublicIPAddress{
				Location: to.StringPtr(location),
			},
			nil)
		if err != nil {
			t.Fatalf("Request error: %v", err)
		}
		var response armnetwork.PublicIPAddressesClientCreateOrUpdateResponse
		if recording.GetRecordMode() == recording.PlaybackMode {
			for {
				_, err = publicIPAddressesClientCreateOrUpdatePollerResponse.Poller.Poll(ctx)
				if err != nil {
					t.Fatalf("Request error: %v", err)
				}
				if publicIPAddressesClientCreateOrUpdatePollerResponse.Poller.Done() {
					response, err = publicIPAddressesClientCreateOrUpdatePollerResponse.Poller.FinalResponse(ctx)
					if err != nil {
						t.Fatalf("Request error: %v", err)
					}
					break
				}
			}
		} else {
			response, err = publicIPAddressesClientCreateOrUpdatePollerResponse.PollUntilDone(ctx, 10*time.Second)
			if err != nil {
				t.Fatalf("Request error: %v", err)
			}
		}
		t.Logf("Response result: %#v\n", response.PublicIPAddressesClientCreateOrUpdateResult)
	}

	// From step PublicIpAddress_CreateDns
	{
		publicIPAddressesClientCreateOrUpdatePollerResponse, err := publicIPAddressesClient.BeginCreateOrUpdate(ctx,
			resourceGroupName,
			"test-ip",
			armnetwork.PublicIPAddress{
				Location: to.StringPtr(location),
				Properties: &armnetwork.PublicIPAddressPropertiesFormat{
					DNSSettings: &armnetwork.PublicIPAddressDNSSettings{
						DomainNameLabel: to.StringPtr("dnslbl"),
					},
				},
			},
			nil)
		if err != nil {
			t.Fatalf("Request error: %v", err)
		}
		var response armnetwork.PublicIPAddressesClientCreateOrUpdateResponse
		if recording.GetRecordMode() == recording.PlaybackMode {
			for {
				_, err = publicIPAddressesClientCreateOrUpdatePollerResponse.Poller.Poll(ctx)
				if err != nil {
					t.Fatalf("Request error: %v", err)
				}
				if publicIPAddressesClientCreateOrUpdatePollerResponse.Poller.Done() {
					response, err = publicIPAddressesClientCreateOrUpdatePollerResponse.Poller.FinalResponse(ctx)
					if err != nil {
						t.Fatalf("Request error: %v", err)
					}
					break
				}
			}
		} else {
			response, err = publicIPAddressesClientCreateOrUpdatePollerResponse.PollUntilDone(ctx, 10*time.Second)
			if err != nil {
				t.Fatalf("Request error: %v", err)
			}
		}
		t.Logf("Response result: %#v\n", response.PublicIPAddressesClientCreateOrUpdateResult)
	}

	// From step PublicIpAddress_Get
	{
		publicIPAddressesClientGetResponse, err := publicIPAddressesClient.Get(ctx,
			resourceGroupName,
			publicIPName,
			&armnetwork.PublicIPAddressesClientGetOptions{Expand: nil})
		if err != nil {
			t.Fatalf("Request error: %v", err)
		}
		t.Logf("Response result: %#v\n", publicIPAddressesClientGetResponse.PublicIPAddressesClientGetResult)
	}

	// From step PublicIpAddress_List
	{
		publicIPAddressesClientListPager := publicIPAddressesClient.List(resourceGroupName,
			nil)
		for publicIPAddressesClientListPager.NextPage(ctx) {
			if err := publicIPAddressesClientListPager.Err(); err != nil {
				t.Fatalf("Failed to advance page: %v", err)
			}
			for _, v := range publicIPAddressesClientListPager.PageResponse().Value {
				t.Logf("Pager result: %#v\n", v)
			}
		}
	}

	// From step PublicIpAddress_ListAll
	{
		publicIPAddressesClientListAllPager := publicIPAddressesClient.ListAll(nil)
		for publicIPAddressesClientListAllPager.NextPage(ctx) {
			if err := publicIPAddressesClientListAllPager.Err(); err != nil {
				t.Fatalf("Failed to advance page: %v", err)
			}
			for _, v := range publicIPAddressesClientListAllPager.PageResponse().Value {
				t.Logf("Pager result: %#v\n", v)
			}
		}
	}

	// From step PublicIpAddress_UpdateTags
	{
		publicIPAddressesClientUpdateTagsResponse, err := publicIPAddressesClient.UpdateTags(ctx,
			resourceGroupName,
			"test-ip",
			armnetwork.TagsObject{
				Tags: map[string]*string{
					"tag1": to.StringPtr("value1"),
					"tag2": to.StringPtr("value2"),
				},
			},
			nil)
		if err != nil {
			t.Fatalf("Request error: %v", err)
		}
		t.Logf("Response result: %#v\n", publicIPAddressesClientUpdateTagsResponse.PublicIPAddressesClientUpdateTagsResult)
	}

	// From step PublicIpAddress_Delete
	{
		publicIPAddressesClientDeletePollerResponse, err := publicIPAddressesClient.BeginDelete(ctx,
			resourceGroupName,
			"test-ip",
			nil)
		if err != nil {
			t.Fatalf("Request error: %v", err)
		}
		if recording.GetRecordMode() == recording.PlaybackMode {
			for {
				_, err = publicIPAddressesClientDeletePollerResponse.Poller.Poll(ctx)
				if err != nil {
					t.Fatalf("Request error: %v", err)
				}
				if publicIPAddressesClientDeletePollerResponse.Poller.Done() {
					_, err = publicIPAddressesClientDeletePollerResponse.Poller.FinalResponse(ctx)
					if err != nil {
						t.Fatalf("Request error: %v", err)
					}
					break
				}
			}
		} else {
			_, err = publicIPAddressesClientDeletePollerResponse.PollUntilDone(ctx, 10*time.Second)
			if err != nil {
				t.Fatalf("Request error: %v", err)
			}
		}
	}
}

func cleanup(t *testing.T) {
}
