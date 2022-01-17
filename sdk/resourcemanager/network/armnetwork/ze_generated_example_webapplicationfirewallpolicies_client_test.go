//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

// x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/WafListPolicies.json
func ExampleWebApplicationFirewallPoliciesClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armnetwork.NewWebApplicationFirewallPoliciesClient("<subscription-id>", cred, nil)
	pager := client.List("<resource-group-name>",
		nil)
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

// x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/WafListAllPolicies.json
func ExampleWebApplicationFirewallPoliciesClient_ListAll() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armnetwork.NewWebApplicationFirewallPoliciesClient("<subscription-id>", cred, nil)
	pager := client.ListAll(nil)
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

// x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/WafPolicyGet.json
func ExampleWebApplicationFirewallPoliciesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armnetwork.NewWebApplicationFirewallPoliciesClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<policy-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.WebApplicationFirewallPoliciesClientGetResult)
}

// x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/WafPolicyCreateOrUpdate.json
func ExampleWebApplicationFirewallPoliciesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armnetwork.NewWebApplicationFirewallPoliciesClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<policy-name>",
		armnetwork.WebApplicationFirewallPolicy{
			Location: to.StringPtr("<location>"),
			Properties: &armnetwork.WebApplicationFirewallPolicyPropertiesFormat{
				CustomRules: []*armnetwork.WebApplicationFirewallCustomRule{
					{
						Name:   to.StringPtr("<name>"),
						Action: armnetwork.WebApplicationFirewallAction("Block").ToPtr(),
						MatchConditions: []*armnetwork.MatchCondition{
							{
								MatchValues: []*string{
									to.StringPtr("192.168.1.0/24"),
									to.StringPtr("10.0.0.0/24")},
								MatchVariables: []*armnetwork.MatchVariable{
									{
										VariableName: armnetwork.WebApplicationFirewallMatchVariable("RemoteAddr").ToPtr(),
									}},
								Operator: armnetwork.WebApplicationFirewallOperator("IPMatch").ToPtr(),
							}},
						Priority: to.Int32Ptr(1),
						RuleType: armnetwork.WebApplicationFirewallRuleType("MatchRule").ToPtr(),
					},
					{
						Name:   to.StringPtr("<name>"),
						Action: armnetwork.WebApplicationFirewallAction("Block").ToPtr(),
						MatchConditions: []*armnetwork.MatchCondition{
							{
								MatchValues: []*string{
									to.StringPtr("192.168.1.0/24")},
								MatchVariables: []*armnetwork.MatchVariable{
									{
										VariableName: armnetwork.WebApplicationFirewallMatchVariable("RemoteAddr").ToPtr(),
									}},
								Operator: armnetwork.WebApplicationFirewallOperator("IPMatch").ToPtr(),
							},
							{
								MatchValues: []*string{
									to.StringPtr("Windows")},
								MatchVariables: []*armnetwork.MatchVariable{
									{
										Selector:     to.StringPtr("<selector>"),
										VariableName: armnetwork.WebApplicationFirewallMatchVariable("RequestHeaders").ToPtr(),
									}},
								Operator: armnetwork.WebApplicationFirewallOperator("Contains").ToPtr(),
							}},
						Priority: to.Int32Ptr(2),
						RuleType: armnetwork.WebApplicationFirewallRuleType("MatchRule").ToPtr(),
					}},
				ManagedRules: &armnetwork.ManagedRulesDefinition{
					Exclusions: []*armnetwork.OwaspCrsExclusionEntry{
						{
							ExclusionManagedRuleSets: []*armnetwork.ExclusionManagedRuleSet{
								{
									RuleGroups: []*armnetwork.ExclusionManagedRuleGroup{
										{
											RuleGroupName: to.StringPtr("<rule-group-name>"),
											Rules: []*armnetwork.ExclusionManagedRule{
												{
													RuleID: to.StringPtr("<rule-id>"),
												}},
										},
										{
											RuleGroupName: to.StringPtr("<rule-group-name>"),
										}},
									RuleSetType:    to.StringPtr("<rule-set-type>"),
									RuleSetVersion: to.StringPtr("<rule-set-version>"),
								}},
							MatchVariable:         armnetwork.OwaspCrsExclusionEntryMatchVariable("RequestArgNames").ToPtr(),
							Selector:              to.StringPtr("<selector>"),
							SelectorMatchOperator: armnetwork.OwaspCrsExclusionEntrySelectorMatchOperator("StartsWith").ToPtr(),
						},
						{
							ExclusionManagedRuleSets: []*armnetwork.ExclusionManagedRuleSet{
								{
									RuleGroups:     []*armnetwork.ExclusionManagedRuleGroup{},
									RuleSetType:    to.StringPtr("<rule-set-type>"),
									RuleSetVersion: to.StringPtr("<rule-set-version>"),
								}},
							MatchVariable:         armnetwork.OwaspCrsExclusionEntryMatchVariable("RequestArgNames").ToPtr(),
							Selector:              to.StringPtr("<selector>"),
							SelectorMatchOperator: armnetwork.OwaspCrsExclusionEntrySelectorMatchOperator("EndsWith").ToPtr(),
						},
						{
							MatchVariable:         armnetwork.OwaspCrsExclusionEntryMatchVariable("RequestArgNames").ToPtr(),
							Selector:              to.StringPtr("<selector>"),
							SelectorMatchOperator: armnetwork.OwaspCrsExclusionEntrySelectorMatchOperator("StartsWith").ToPtr(),
						},
						{
							MatchVariable:         armnetwork.OwaspCrsExclusionEntryMatchVariable("RequestArgValues").ToPtr(),
							Selector:              to.StringPtr("<selector>"),
							SelectorMatchOperator: armnetwork.OwaspCrsExclusionEntrySelectorMatchOperator("StartsWith").ToPtr(),
						}},
					ManagedRuleSets: []*armnetwork.ManagedRuleSet{
						{
							RuleSetType:    to.StringPtr("<rule-set-type>"),
							RuleSetVersion: to.StringPtr("<rule-set-version>"),
						}},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.WebApplicationFirewallPoliciesClientCreateOrUpdateResult)
}

// x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/WafPolicyDelete.json
func ExampleWebApplicationFirewallPoliciesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armnetwork.NewWebApplicationFirewallPoliciesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<policy-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}
