//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armeventhub

import (
	"context"
	"net/http"
	"time"

	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
)

// ClustersCreateOrUpdatePollerResponse contains the response from method Clusters.CreateOrUpdate.
type ClustersCreateOrUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ClustersCreateOrUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
func (l ClustersCreateOrUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ClustersCreateOrUpdateResponse, error) {
	respType := ClustersCreateOrUpdateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.Cluster)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a ClustersCreateOrUpdatePollerResponse from the provided client and resume token.
func (l *ClustersCreateOrUpdatePollerResponse) Resume(ctx context.Context, client *ClustersClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("ClustersClient.CreateOrUpdate", token, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return err
	}
	poller := &ClustersCreateOrUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// ClustersCreateOrUpdateResponse contains the response from method Clusters.CreateOrUpdate.
type ClustersCreateOrUpdateResponse struct {
	ClustersCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ClustersCreateOrUpdateResult contains the result from method Clusters.CreateOrUpdate.
type ClustersCreateOrUpdateResult struct {
	Cluster
}

// ClustersDeletePollerResponse contains the response from method Clusters.Delete.
type ClustersDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ClustersDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
func (l ClustersDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ClustersDeleteResponse, error) {
	respType := ClustersDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a ClustersDeletePollerResponse from the provided client and resume token.
func (l *ClustersDeletePollerResponse) Resume(ctx context.Context, client *ClustersClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("ClustersClient.Delete", token, client.pl, client.deleteHandleError)
	if err != nil {
		return err
	}
	poller := &ClustersDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// ClustersDeleteResponse contains the response from method Clusters.Delete.
type ClustersDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ClustersGetResponse contains the response from method Clusters.Get.
type ClustersGetResponse struct {
	ClustersGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ClustersGetResult contains the result from method Clusters.Get.
type ClustersGetResult struct {
	Cluster
}

// ClustersListAvailableClusterRegionResponse contains the response from method Clusters.ListAvailableClusterRegion.
type ClustersListAvailableClusterRegionResponse struct {
	ClustersListAvailableClusterRegionResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ClustersListAvailableClusterRegionResult contains the result from method Clusters.ListAvailableClusterRegion.
type ClustersListAvailableClusterRegionResult struct {
	AvailableClustersList
}

// ClustersListByResourceGroupResponse contains the response from method Clusters.ListByResourceGroup.
type ClustersListByResourceGroupResponse struct {
	ClustersListByResourceGroupResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ClustersListByResourceGroupResult contains the result from method Clusters.ListByResourceGroup.
type ClustersListByResourceGroupResult struct {
	ClusterListResult
}

// ClustersListNamespacesResponse contains the response from method Clusters.ListNamespaces.
type ClustersListNamespacesResponse struct {
	ClustersListNamespacesResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ClustersListNamespacesResult contains the result from method Clusters.ListNamespaces.
type ClustersListNamespacesResult struct {
	EHNamespaceIDListResult
}

// ClustersUpdatePollerResponse contains the response from method Clusters.Update.
type ClustersUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ClustersUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
func (l ClustersUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ClustersUpdateResponse, error) {
	respType := ClustersUpdateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.Cluster)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a ClustersUpdatePollerResponse from the provided client and resume token.
func (l *ClustersUpdatePollerResponse) Resume(ctx context.Context, client *ClustersClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("ClustersClient.Update", token, client.pl, client.updateHandleError)
	if err != nil {
		return err
	}
	poller := &ClustersUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// ClustersUpdateResponse contains the response from method Clusters.Update.
type ClustersUpdateResponse struct {
	ClustersUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ClustersUpdateResult contains the result from method Clusters.Update.
type ClustersUpdateResult struct {
	Cluster
}

// ConfigurationGetResponse contains the response from method Configuration.Get.
type ConfigurationGetResponse struct {
	ConfigurationGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ConfigurationGetResult contains the result from method Configuration.Get.
type ConfigurationGetResult struct {
	ClusterQuotaConfigurationProperties
}

// ConfigurationPatchResponse contains the response from method Configuration.Patch.
type ConfigurationPatchResponse struct {
	ConfigurationPatchResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ConfigurationPatchResult contains the result from method Configuration.Patch.
type ConfigurationPatchResult struct {
	ClusterQuotaConfigurationProperties
}

// ConsumerGroupsCreateOrUpdateResponse contains the response from method ConsumerGroups.CreateOrUpdate.
type ConsumerGroupsCreateOrUpdateResponse struct {
	ConsumerGroupsCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ConsumerGroupsCreateOrUpdateResult contains the result from method ConsumerGroups.CreateOrUpdate.
type ConsumerGroupsCreateOrUpdateResult struct {
	ConsumerGroup
}

// ConsumerGroupsDeleteResponse contains the response from method ConsumerGroups.Delete.
type ConsumerGroupsDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ConsumerGroupsGetResponse contains the response from method ConsumerGroups.Get.
type ConsumerGroupsGetResponse struct {
	ConsumerGroupsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ConsumerGroupsGetResult contains the result from method ConsumerGroups.Get.
type ConsumerGroupsGetResult struct {
	ConsumerGroup
}

// ConsumerGroupsListByEventHubResponse contains the response from method ConsumerGroups.ListByEventHub.
type ConsumerGroupsListByEventHubResponse struct {
	ConsumerGroupsListByEventHubResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ConsumerGroupsListByEventHubResult contains the result from method ConsumerGroups.ListByEventHub.
type ConsumerGroupsListByEventHubResult struct {
	ConsumerGroupListResult
}

// DisasterRecoveryConfigsBreakPairingResponse contains the response from method DisasterRecoveryConfigs.BreakPairing.
type DisasterRecoveryConfigsBreakPairingResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DisasterRecoveryConfigsCheckNameAvailabilityResponse contains the response from method DisasterRecoveryConfigs.CheckNameAvailability.
type DisasterRecoveryConfigsCheckNameAvailabilityResponse struct {
	DisasterRecoveryConfigsCheckNameAvailabilityResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DisasterRecoveryConfigsCheckNameAvailabilityResult contains the result from method DisasterRecoveryConfigs.CheckNameAvailability.
type DisasterRecoveryConfigsCheckNameAvailabilityResult struct {
	CheckNameAvailabilityResult
}

// DisasterRecoveryConfigsCreateOrUpdateResponse contains the response from method DisasterRecoveryConfigs.CreateOrUpdate.
type DisasterRecoveryConfigsCreateOrUpdateResponse struct {
	DisasterRecoveryConfigsCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DisasterRecoveryConfigsCreateOrUpdateResult contains the result from method DisasterRecoveryConfigs.CreateOrUpdate.
type DisasterRecoveryConfigsCreateOrUpdateResult struct {
	ArmDisasterRecovery
}

// DisasterRecoveryConfigsDeleteResponse contains the response from method DisasterRecoveryConfigs.Delete.
type DisasterRecoveryConfigsDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DisasterRecoveryConfigsFailOverResponse contains the response from method DisasterRecoveryConfigs.FailOver.
type DisasterRecoveryConfigsFailOverResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DisasterRecoveryConfigsGetAuthorizationRuleResponse contains the response from method DisasterRecoveryConfigs.GetAuthorizationRule.
type DisasterRecoveryConfigsGetAuthorizationRuleResponse struct {
	DisasterRecoveryConfigsGetAuthorizationRuleResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DisasterRecoveryConfigsGetAuthorizationRuleResult contains the result from method DisasterRecoveryConfigs.GetAuthorizationRule.
type DisasterRecoveryConfigsGetAuthorizationRuleResult struct {
	AuthorizationRule
}

// DisasterRecoveryConfigsGetResponse contains the response from method DisasterRecoveryConfigs.Get.
type DisasterRecoveryConfigsGetResponse struct {
	DisasterRecoveryConfigsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DisasterRecoveryConfigsGetResult contains the result from method DisasterRecoveryConfigs.Get.
type DisasterRecoveryConfigsGetResult struct {
	ArmDisasterRecovery
}

// DisasterRecoveryConfigsListAuthorizationRulesResponse contains the response from method DisasterRecoveryConfigs.ListAuthorizationRules.
type DisasterRecoveryConfigsListAuthorizationRulesResponse struct {
	DisasterRecoveryConfigsListAuthorizationRulesResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DisasterRecoveryConfigsListAuthorizationRulesResult contains the result from method DisasterRecoveryConfigs.ListAuthorizationRules.
type DisasterRecoveryConfigsListAuthorizationRulesResult struct {
	AuthorizationRuleListResult
}

// DisasterRecoveryConfigsListKeysResponse contains the response from method DisasterRecoveryConfigs.ListKeys.
type DisasterRecoveryConfigsListKeysResponse struct {
	DisasterRecoveryConfigsListKeysResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DisasterRecoveryConfigsListKeysResult contains the result from method DisasterRecoveryConfigs.ListKeys.
type DisasterRecoveryConfigsListKeysResult struct {
	AccessKeys
}

// DisasterRecoveryConfigsListResponse contains the response from method DisasterRecoveryConfigs.List.
type DisasterRecoveryConfigsListResponse struct {
	DisasterRecoveryConfigsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DisasterRecoveryConfigsListResult contains the result from method DisasterRecoveryConfigs.List.
type DisasterRecoveryConfigsListResult struct {
	ArmDisasterRecoveryListResult
}

// EventHubsCreateOrUpdateAuthorizationRuleResponse contains the response from method EventHubs.CreateOrUpdateAuthorizationRule.
type EventHubsCreateOrUpdateAuthorizationRuleResponse struct {
	EventHubsCreateOrUpdateAuthorizationRuleResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// EventHubsCreateOrUpdateAuthorizationRuleResult contains the result from method EventHubs.CreateOrUpdateAuthorizationRule.
type EventHubsCreateOrUpdateAuthorizationRuleResult struct {
	AuthorizationRule
}

// EventHubsCreateOrUpdateResponse contains the response from method EventHubs.CreateOrUpdate.
type EventHubsCreateOrUpdateResponse struct {
	EventHubsCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// EventHubsCreateOrUpdateResult contains the result from method EventHubs.CreateOrUpdate.
type EventHubsCreateOrUpdateResult struct {
	Eventhub
}

// EventHubsDeleteAuthorizationRuleResponse contains the response from method EventHubs.DeleteAuthorizationRule.
type EventHubsDeleteAuthorizationRuleResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// EventHubsDeleteResponse contains the response from method EventHubs.Delete.
type EventHubsDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// EventHubsGetAuthorizationRuleResponse contains the response from method EventHubs.GetAuthorizationRule.
type EventHubsGetAuthorizationRuleResponse struct {
	EventHubsGetAuthorizationRuleResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// EventHubsGetAuthorizationRuleResult contains the result from method EventHubs.GetAuthorizationRule.
type EventHubsGetAuthorizationRuleResult struct {
	AuthorizationRule
}

// EventHubsGetResponse contains the response from method EventHubs.Get.
type EventHubsGetResponse struct {
	EventHubsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// EventHubsGetResult contains the result from method EventHubs.Get.
type EventHubsGetResult struct {
	Eventhub
}

// EventHubsListAuthorizationRulesResponse contains the response from method EventHubs.ListAuthorizationRules.
type EventHubsListAuthorizationRulesResponse struct {
	EventHubsListAuthorizationRulesResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// EventHubsListAuthorizationRulesResult contains the result from method EventHubs.ListAuthorizationRules.
type EventHubsListAuthorizationRulesResult struct {
	AuthorizationRuleListResult
}

// EventHubsListByNamespaceResponse contains the response from method EventHubs.ListByNamespace.
type EventHubsListByNamespaceResponse struct {
	EventHubsListByNamespaceResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// EventHubsListByNamespaceResult contains the result from method EventHubs.ListByNamespace.
type EventHubsListByNamespaceResult struct {
	EventHubListResult
}

// EventHubsListKeysResponse contains the response from method EventHubs.ListKeys.
type EventHubsListKeysResponse struct {
	EventHubsListKeysResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// EventHubsListKeysResult contains the result from method EventHubs.ListKeys.
type EventHubsListKeysResult struct {
	AccessKeys
}

// EventHubsRegenerateKeysResponse contains the response from method EventHubs.RegenerateKeys.
type EventHubsRegenerateKeysResponse struct {
	EventHubsRegenerateKeysResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// EventHubsRegenerateKeysResult contains the result from method EventHubs.RegenerateKeys.
type EventHubsRegenerateKeysResult struct {
	AccessKeys
}

// NamespacesCheckNameAvailabilityResponse contains the response from method Namespaces.CheckNameAvailability.
type NamespacesCheckNameAvailabilityResponse struct {
	NamespacesCheckNameAvailabilityResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesCheckNameAvailabilityResult contains the result from method Namespaces.CheckNameAvailability.
type NamespacesCheckNameAvailabilityResult struct {
	CheckNameAvailabilityResult
}

// NamespacesCreateOrUpdateAuthorizationRuleResponse contains the response from method Namespaces.CreateOrUpdateAuthorizationRule.
type NamespacesCreateOrUpdateAuthorizationRuleResponse struct {
	NamespacesCreateOrUpdateAuthorizationRuleResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesCreateOrUpdateAuthorizationRuleResult contains the result from method Namespaces.CreateOrUpdateAuthorizationRule.
type NamespacesCreateOrUpdateAuthorizationRuleResult struct {
	AuthorizationRule
}

// NamespacesCreateOrUpdateIPFilterRuleResponse contains the response from method Namespaces.CreateOrUpdateIPFilterRule.
type NamespacesCreateOrUpdateIPFilterRuleResponse struct {
	NamespacesCreateOrUpdateIPFilterRuleResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesCreateOrUpdateIPFilterRuleResult contains the result from method Namespaces.CreateOrUpdateIPFilterRule.
type NamespacesCreateOrUpdateIPFilterRuleResult struct {
	IPFilterRule
}

// NamespacesCreateOrUpdateNetworkRuleSetResponse contains the response from method Namespaces.CreateOrUpdateNetworkRuleSet.
type NamespacesCreateOrUpdateNetworkRuleSetResponse struct {
	NamespacesCreateOrUpdateNetworkRuleSetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesCreateOrUpdateNetworkRuleSetResult contains the result from method Namespaces.CreateOrUpdateNetworkRuleSet.
type NamespacesCreateOrUpdateNetworkRuleSetResult struct {
	NetworkRuleSet
}

// NamespacesCreateOrUpdatePollerResponse contains the response from method Namespaces.CreateOrUpdate.
type NamespacesCreateOrUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *NamespacesCreateOrUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
func (l NamespacesCreateOrUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (NamespacesCreateOrUpdateResponse, error) {
	respType := NamespacesCreateOrUpdateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.EHNamespace)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a NamespacesCreateOrUpdatePollerResponse from the provided client and resume token.
func (l *NamespacesCreateOrUpdatePollerResponse) Resume(ctx context.Context, client *NamespacesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("NamespacesClient.CreateOrUpdate", token, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return err
	}
	poller := &NamespacesCreateOrUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// NamespacesCreateOrUpdateResponse contains the response from method Namespaces.CreateOrUpdate.
type NamespacesCreateOrUpdateResponse struct {
	NamespacesCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesCreateOrUpdateResult contains the result from method Namespaces.CreateOrUpdate.
type NamespacesCreateOrUpdateResult struct {
	EHNamespace
}

// NamespacesCreateOrUpdateVirtualNetworkRuleResponse contains the response from method Namespaces.CreateOrUpdateVirtualNetworkRule.
type NamespacesCreateOrUpdateVirtualNetworkRuleResponse struct {
	NamespacesCreateOrUpdateVirtualNetworkRuleResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesCreateOrUpdateVirtualNetworkRuleResult contains the result from method Namespaces.CreateOrUpdateVirtualNetworkRule.
type NamespacesCreateOrUpdateVirtualNetworkRuleResult struct {
	VirtualNetworkRule
}

// NamespacesDeleteAuthorizationRuleResponse contains the response from method Namespaces.DeleteAuthorizationRule.
type NamespacesDeleteAuthorizationRuleResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesDeleteIPFilterRuleResponse contains the response from method Namespaces.DeleteIPFilterRule.
type NamespacesDeleteIPFilterRuleResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesDeletePollerResponse contains the response from method Namespaces.Delete.
type NamespacesDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *NamespacesDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
func (l NamespacesDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (NamespacesDeleteResponse, error) {
	respType := NamespacesDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a NamespacesDeletePollerResponse from the provided client and resume token.
func (l *NamespacesDeletePollerResponse) Resume(ctx context.Context, client *NamespacesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("NamespacesClient.Delete", token, client.pl, client.deleteHandleError)
	if err != nil {
		return err
	}
	poller := &NamespacesDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// NamespacesDeleteResponse contains the response from method Namespaces.Delete.
type NamespacesDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesDeleteVirtualNetworkRuleResponse contains the response from method Namespaces.DeleteVirtualNetworkRule.
type NamespacesDeleteVirtualNetworkRuleResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesGetAuthorizationRuleResponse contains the response from method Namespaces.GetAuthorizationRule.
type NamespacesGetAuthorizationRuleResponse struct {
	NamespacesGetAuthorizationRuleResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesGetAuthorizationRuleResult contains the result from method Namespaces.GetAuthorizationRule.
type NamespacesGetAuthorizationRuleResult struct {
	AuthorizationRule
}

// NamespacesGetIPFilterRuleResponse contains the response from method Namespaces.GetIPFilterRule.
type NamespacesGetIPFilterRuleResponse struct {
	NamespacesGetIPFilterRuleResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesGetIPFilterRuleResult contains the result from method Namespaces.GetIPFilterRule.
type NamespacesGetIPFilterRuleResult struct {
	IPFilterRule
}

// NamespacesGetNetworkRuleSetResponse contains the response from method Namespaces.GetNetworkRuleSet.
type NamespacesGetNetworkRuleSetResponse struct {
	NamespacesGetNetworkRuleSetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesGetNetworkRuleSetResult contains the result from method Namespaces.GetNetworkRuleSet.
type NamespacesGetNetworkRuleSetResult struct {
	NetworkRuleSet
}

// NamespacesGetResponse contains the response from method Namespaces.Get.
type NamespacesGetResponse struct {
	NamespacesGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesGetResult contains the result from method Namespaces.Get.
type NamespacesGetResult struct {
	EHNamespace
}

// NamespacesGetVirtualNetworkRuleResponse contains the response from method Namespaces.GetVirtualNetworkRule.
type NamespacesGetVirtualNetworkRuleResponse struct {
	NamespacesGetVirtualNetworkRuleResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesGetVirtualNetworkRuleResult contains the result from method Namespaces.GetVirtualNetworkRule.
type NamespacesGetVirtualNetworkRuleResult struct {
	VirtualNetworkRule
}

// NamespacesListAuthorizationRulesResponse contains the response from method Namespaces.ListAuthorizationRules.
type NamespacesListAuthorizationRulesResponse struct {
	NamespacesListAuthorizationRulesResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesListAuthorizationRulesResult contains the result from method Namespaces.ListAuthorizationRules.
type NamespacesListAuthorizationRulesResult struct {
	AuthorizationRuleListResult
}

// NamespacesListByResourceGroupResponse contains the response from method Namespaces.ListByResourceGroup.
type NamespacesListByResourceGroupResponse struct {
	NamespacesListByResourceGroupResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesListByResourceGroupResult contains the result from method Namespaces.ListByResourceGroup.
type NamespacesListByResourceGroupResult struct {
	EHNamespaceListResult
}

// NamespacesListIPFilterRulesResponse contains the response from method Namespaces.ListIPFilterRules.
type NamespacesListIPFilterRulesResponse struct {
	NamespacesListIPFilterRulesResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesListIPFilterRulesResult contains the result from method Namespaces.ListIPFilterRules.
type NamespacesListIPFilterRulesResult struct {
	IPFilterRuleListResult
}

// NamespacesListKeysResponse contains the response from method Namespaces.ListKeys.
type NamespacesListKeysResponse struct {
	NamespacesListKeysResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesListKeysResult contains the result from method Namespaces.ListKeys.
type NamespacesListKeysResult struct {
	AccessKeys
}

// NamespacesListResponse contains the response from method Namespaces.List.
type NamespacesListResponse struct {
	NamespacesListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesListResult contains the result from method Namespaces.List.
type NamespacesListResult struct {
	EHNamespaceListResult
}

// NamespacesListVirtualNetworkRulesResponse contains the response from method Namespaces.ListVirtualNetworkRules.
type NamespacesListVirtualNetworkRulesResponse struct {
	NamespacesListVirtualNetworkRulesResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesListVirtualNetworkRulesResult contains the result from method Namespaces.ListVirtualNetworkRules.
type NamespacesListVirtualNetworkRulesResult struct {
	VirtualNetworkRuleListResult
}

// NamespacesRegenerateKeysResponse contains the response from method Namespaces.RegenerateKeys.
type NamespacesRegenerateKeysResponse struct {
	NamespacesRegenerateKeysResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesRegenerateKeysResult contains the result from method Namespaces.RegenerateKeys.
type NamespacesRegenerateKeysResult struct {
	AccessKeys
}

// NamespacesUpdateResponse contains the response from method Namespaces.Update.
type NamespacesUpdateResponse struct {
	NamespacesUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesUpdateResult contains the result from method Namespaces.Update.
type NamespacesUpdateResult struct {
	EHNamespace
}

// OperationsListResponse contains the response from method Operations.List.
type OperationsListResponse struct {
	OperationsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// OperationsListResult contains the result from method Operations.List.
type OperationsListResult struct {
	OperationListResult
}

// PrivateEndpointConnectionsCreateOrUpdateResponse contains the response from method PrivateEndpointConnections.CreateOrUpdate.
type PrivateEndpointConnectionsCreateOrUpdateResponse struct {
	PrivateEndpointConnectionsCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateEndpointConnectionsCreateOrUpdateResult contains the result from method PrivateEndpointConnections.CreateOrUpdate.
type PrivateEndpointConnectionsCreateOrUpdateResult struct {
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsDeletePollerResponse contains the response from method PrivateEndpointConnections.Delete.
type PrivateEndpointConnectionsDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *PrivateEndpointConnectionsDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
func (l PrivateEndpointConnectionsDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (PrivateEndpointConnectionsDeleteResponse, error) {
	respType := PrivateEndpointConnectionsDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a PrivateEndpointConnectionsDeletePollerResponse from the provided client and resume token.
func (l *PrivateEndpointConnectionsDeletePollerResponse) Resume(ctx context.Context, client *PrivateEndpointConnectionsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("PrivateEndpointConnectionsClient.Delete", token, client.pl, client.deleteHandleError)
	if err != nil {
		return err
	}
	poller := &PrivateEndpointConnectionsDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// PrivateEndpointConnectionsDeleteResponse contains the response from method PrivateEndpointConnections.Delete.
type PrivateEndpointConnectionsDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateEndpointConnectionsGetResponse contains the response from method PrivateEndpointConnections.Get.
type PrivateEndpointConnectionsGetResponse struct {
	PrivateEndpointConnectionsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateEndpointConnectionsGetResult contains the result from method PrivateEndpointConnections.Get.
type PrivateEndpointConnectionsGetResult struct {
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsListResponse contains the response from method PrivateEndpointConnections.List.
type PrivateEndpointConnectionsListResponse struct {
	PrivateEndpointConnectionsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateEndpointConnectionsListResult contains the result from method PrivateEndpointConnections.List.
type PrivateEndpointConnectionsListResult struct {
	PrivateEndpointConnectionListResult
}

// PrivateLinkResourcesGetResponse contains the response from method PrivateLinkResources.Get.
type PrivateLinkResourcesGetResponse struct {
	PrivateLinkResourcesGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateLinkResourcesGetResult contains the result from method PrivateLinkResources.Get.
type PrivateLinkResourcesGetResult struct {
	PrivateLinkResourcesListResult
}

// RegionsListBySKUResponse contains the response from method Regions.ListBySKU.
type RegionsListBySKUResponse struct {
	RegionsListBySKUResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RegionsListBySKUResult contains the result from method Regions.ListBySKU.
type RegionsListBySKUResult struct {
	MessagingRegionsListResult
}
