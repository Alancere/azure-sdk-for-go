package eventgridapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/eventgrid/mgmt/2021-12-01/eventgrid"
)

// DomainsClientAPI contains the set of methods on the DomainsClient type.
type DomainsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, domainName string, domainInfo eventgrid.Domain) (result eventgrid.DomainsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, domainName string) (result eventgrid.DomainsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, domainName string) (result eventgrid.Domain, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string, filter string, top *int32) (result eventgrid.DomainsListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, filter string, top *int32) (result eventgrid.DomainsListResultIterator, err error)
	ListBySubscription(ctx context.Context, filter string, top *int32) (result eventgrid.DomainsListResultPage, err error)
	ListBySubscriptionComplete(ctx context.Context, filter string, top *int32) (result eventgrid.DomainsListResultIterator, err error)
	ListSharedAccessKeys(ctx context.Context, resourceGroupName string, domainName string) (result eventgrid.DomainSharedAccessKeys, err error)
	RegenerateKey(ctx context.Context, resourceGroupName string, domainName string, regenerateKeyRequest eventgrid.DomainRegenerateKeyRequest) (result eventgrid.DomainSharedAccessKeys, err error)
	Update(ctx context.Context, resourceGroupName string, domainName string, domainUpdateParameters eventgrid.DomainUpdateParameters) (result eventgrid.DomainsUpdateFuture, err error)
}

var _ DomainsClientAPI = (*eventgrid.DomainsClient)(nil)

// DomainTopicsClientAPI contains the set of methods on the DomainTopicsClient type.
type DomainTopicsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, domainName string, domainTopicName string) (result eventgrid.DomainTopicsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, domainName string, domainTopicName string) (result eventgrid.DomainTopicsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, domainName string, domainTopicName string) (result eventgrid.DomainTopic, err error)
	ListByDomain(ctx context.Context, resourceGroupName string, domainName string, filter string, top *int32) (result eventgrid.DomainTopicsListResultPage, err error)
	ListByDomainComplete(ctx context.Context, resourceGroupName string, domainName string, filter string, top *int32) (result eventgrid.DomainTopicsListResultIterator, err error)
}

var _ DomainTopicsClientAPI = (*eventgrid.DomainTopicsClient)(nil)

// EventSubscriptionsClientAPI contains the set of methods on the EventSubscriptionsClient type.
type EventSubscriptionsClientAPI interface {
	CreateOrUpdate(ctx context.Context, scope string, eventSubscriptionName string, eventSubscriptionInfo eventgrid.EventSubscription) (result eventgrid.EventSubscriptionsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, scope string, eventSubscriptionName string) (result eventgrid.EventSubscriptionsDeleteFuture, err error)
	Get(ctx context.Context, scope string, eventSubscriptionName string) (result eventgrid.EventSubscription, err error)
	GetDeliveryAttributes(ctx context.Context, scope string, eventSubscriptionName string) (result eventgrid.DeliveryAttributeListResult, err error)
	GetFullURL(ctx context.Context, scope string, eventSubscriptionName string) (result eventgrid.EventSubscriptionFullURL, err error)
	ListByDomainTopic(ctx context.Context, resourceGroupName string, domainName string, topicName string, filter string, top *int32) (result eventgrid.EventSubscriptionsListResultPage, err error)
	ListByDomainTopicComplete(ctx context.Context, resourceGroupName string, domainName string, topicName string, filter string, top *int32) (result eventgrid.EventSubscriptionsListResultIterator, err error)
	ListByResource(ctx context.Context, resourceGroupName string, providerNamespace string, resourceTypeName string, resourceName string, filter string, top *int32) (result eventgrid.EventSubscriptionsListResultPage, err error)
	ListByResourceComplete(ctx context.Context, resourceGroupName string, providerNamespace string, resourceTypeName string, resourceName string, filter string, top *int32) (result eventgrid.EventSubscriptionsListResultIterator, err error)
	ListGlobalByResourceGroup(ctx context.Context, resourceGroupName string, filter string, top *int32) (result eventgrid.EventSubscriptionsListResultPage, err error)
	ListGlobalByResourceGroupComplete(ctx context.Context, resourceGroupName string, filter string, top *int32) (result eventgrid.EventSubscriptionsListResultIterator, err error)
	ListGlobalByResourceGroupForTopicType(ctx context.Context, resourceGroupName string, topicTypeName string, filter string, top *int32) (result eventgrid.EventSubscriptionsListResultPage, err error)
	ListGlobalByResourceGroupForTopicTypeComplete(ctx context.Context, resourceGroupName string, topicTypeName string, filter string, top *int32) (result eventgrid.EventSubscriptionsListResultIterator, err error)
	ListGlobalBySubscription(ctx context.Context, filter string, top *int32) (result eventgrid.EventSubscriptionsListResultPage, err error)
	ListGlobalBySubscriptionComplete(ctx context.Context, filter string, top *int32) (result eventgrid.EventSubscriptionsListResultIterator, err error)
	ListGlobalBySubscriptionForTopicType(ctx context.Context, topicTypeName string, filter string, top *int32) (result eventgrid.EventSubscriptionsListResultPage, err error)
	ListGlobalBySubscriptionForTopicTypeComplete(ctx context.Context, topicTypeName string, filter string, top *int32) (result eventgrid.EventSubscriptionsListResultIterator, err error)
	ListRegionalByResourceGroup(ctx context.Context, resourceGroupName string, location string, filter string, top *int32) (result eventgrid.EventSubscriptionsListResultPage, err error)
	ListRegionalByResourceGroupComplete(ctx context.Context, resourceGroupName string, location string, filter string, top *int32) (result eventgrid.EventSubscriptionsListResultIterator, err error)
	ListRegionalByResourceGroupForTopicType(ctx context.Context, resourceGroupName string, location string, topicTypeName string, filter string, top *int32) (result eventgrid.EventSubscriptionsListResultPage, err error)
	ListRegionalByResourceGroupForTopicTypeComplete(ctx context.Context, resourceGroupName string, location string, topicTypeName string, filter string, top *int32) (result eventgrid.EventSubscriptionsListResultIterator, err error)
	ListRegionalBySubscription(ctx context.Context, location string, filter string, top *int32) (result eventgrid.EventSubscriptionsListResultPage, err error)
	ListRegionalBySubscriptionComplete(ctx context.Context, location string, filter string, top *int32) (result eventgrid.EventSubscriptionsListResultIterator, err error)
	ListRegionalBySubscriptionForTopicType(ctx context.Context, location string, topicTypeName string, filter string, top *int32) (result eventgrid.EventSubscriptionsListResultPage, err error)
	ListRegionalBySubscriptionForTopicTypeComplete(ctx context.Context, location string, topicTypeName string, filter string, top *int32) (result eventgrid.EventSubscriptionsListResultIterator, err error)
	Update(ctx context.Context, scope string, eventSubscriptionName string, eventSubscriptionUpdateParameters eventgrid.EventSubscriptionUpdateParameters) (result eventgrid.EventSubscriptionsUpdateFuture, err error)
}

var _ EventSubscriptionsClientAPI = (*eventgrid.EventSubscriptionsClient)(nil)

// SystemTopicEventSubscriptionsClientAPI contains the set of methods on the SystemTopicEventSubscriptionsClient type.
type SystemTopicEventSubscriptionsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, systemTopicName string, eventSubscriptionName string, eventSubscriptionInfo eventgrid.EventSubscription) (result eventgrid.SystemTopicEventSubscriptionsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, systemTopicName string, eventSubscriptionName string) (result eventgrid.SystemTopicEventSubscriptionsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, systemTopicName string, eventSubscriptionName string) (result eventgrid.EventSubscription, err error)
	GetDeliveryAttributes(ctx context.Context, resourceGroupName string, systemTopicName string, eventSubscriptionName string) (result eventgrid.DeliveryAttributeListResult, err error)
	GetFullURL(ctx context.Context, resourceGroupName string, systemTopicName string, eventSubscriptionName string) (result eventgrid.EventSubscriptionFullURL, err error)
	ListBySystemTopic(ctx context.Context, resourceGroupName string, systemTopicName string, filter string, top *int32) (result eventgrid.EventSubscriptionsListResultPage, err error)
	ListBySystemTopicComplete(ctx context.Context, resourceGroupName string, systemTopicName string, filter string, top *int32) (result eventgrid.EventSubscriptionsListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, systemTopicName string, eventSubscriptionName string, eventSubscriptionUpdateParameters eventgrid.EventSubscriptionUpdateParameters) (result eventgrid.SystemTopicEventSubscriptionsUpdateFuture, err error)
}

var _ SystemTopicEventSubscriptionsClientAPI = (*eventgrid.SystemTopicEventSubscriptionsClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result eventgrid.OperationsListResult, err error)
}

var _ OperationsClientAPI = (*eventgrid.OperationsClient)(nil)

// TopicsClientAPI contains the set of methods on the TopicsClient type.
type TopicsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, topicName string, topicInfo eventgrid.Topic) (result eventgrid.TopicsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, topicName string) (result eventgrid.TopicsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, topicName string) (result eventgrid.Topic, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string, filter string, top *int32) (result eventgrid.TopicsListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, filter string, top *int32) (result eventgrid.TopicsListResultIterator, err error)
	ListBySubscription(ctx context.Context, filter string, top *int32) (result eventgrid.TopicsListResultPage, err error)
	ListBySubscriptionComplete(ctx context.Context, filter string, top *int32) (result eventgrid.TopicsListResultIterator, err error)
	ListEventTypes(ctx context.Context, resourceGroupName string, providerNamespace string, resourceTypeName string, resourceName string) (result eventgrid.EventTypesListResult, err error)
	ListSharedAccessKeys(ctx context.Context, resourceGroupName string, topicName string) (result eventgrid.TopicSharedAccessKeys, err error)
	RegenerateKey(ctx context.Context, resourceGroupName string, topicName string, regenerateKeyRequest eventgrid.TopicRegenerateKeyRequest) (result eventgrid.TopicsRegenerateKeyFuture, err error)
	Update(ctx context.Context, resourceGroupName string, topicName string, topicUpdateParameters eventgrid.TopicUpdateParameters) (result eventgrid.TopicsUpdateFuture, err error)
}

var _ TopicsClientAPI = (*eventgrid.TopicsClient)(nil)

// PrivateEndpointConnectionsClientAPI contains the set of methods on the PrivateEndpointConnectionsClient type.
type PrivateEndpointConnectionsClientAPI interface {
	Delete(ctx context.Context, resourceGroupName string, parentType string, parentName string, privateEndpointConnectionName string) (result eventgrid.PrivateEndpointConnectionsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, parentType string, parentName string, privateEndpointConnectionName string) (result eventgrid.PrivateEndpointConnection, err error)
	ListByResource(ctx context.Context, resourceGroupName string, parentType string, parentName string, filter string, top *int32) (result eventgrid.PrivateEndpointConnectionListResultPage, err error)
	ListByResourceComplete(ctx context.Context, resourceGroupName string, parentType string, parentName string, filter string, top *int32) (result eventgrid.PrivateEndpointConnectionListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, parentType string, parentName string, privateEndpointConnectionName string, privateEndpointConnection eventgrid.PrivateEndpointConnection) (result eventgrid.PrivateEndpointConnectionsUpdateFuture, err error)
}

var _ PrivateEndpointConnectionsClientAPI = (*eventgrid.PrivateEndpointConnectionsClient)(nil)

// PrivateLinkResourcesClientAPI contains the set of methods on the PrivateLinkResourcesClient type.
type PrivateLinkResourcesClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, parentType string, parentName string, privateLinkResourceName string) (result eventgrid.PrivateLinkResource, err error)
	ListByResource(ctx context.Context, resourceGroupName string, parentType string, parentName string, filter string, top *int32) (result eventgrid.PrivateLinkResourcesListResultPage, err error)
	ListByResourceComplete(ctx context.Context, resourceGroupName string, parentType string, parentName string, filter string, top *int32) (result eventgrid.PrivateLinkResourcesListResultIterator, err error)
}

var _ PrivateLinkResourcesClientAPI = (*eventgrid.PrivateLinkResourcesClient)(nil)

// SystemTopicsClientAPI contains the set of methods on the SystemTopicsClient type.
type SystemTopicsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, systemTopicName string, systemTopicInfo eventgrid.SystemTopic) (result eventgrid.SystemTopicsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, systemTopicName string) (result eventgrid.SystemTopicsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, systemTopicName string) (result eventgrid.SystemTopic, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string, filter string, top *int32) (result eventgrid.SystemTopicsListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, filter string, top *int32) (result eventgrid.SystemTopicsListResultIterator, err error)
	ListBySubscription(ctx context.Context, filter string, top *int32) (result eventgrid.SystemTopicsListResultPage, err error)
	ListBySubscriptionComplete(ctx context.Context, filter string, top *int32) (result eventgrid.SystemTopicsListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, systemTopicName string, systemTopicUpdateParameters eventgrid.SystemTopicUpdateParameters) (result eventgrid.SystemTopicsUpdateFuture, err error)
}

var _ SystemTopicsClientAPI = (*eventgrid.SystemTopicsClient)(nil)

// ExtensionTopicsClientAPI contains the set of methods on the ExtensionTopicsClient type.
type ExtensionTopicsClientAPI interface {
	Get(ctx context.Context, scope string) (result eventgrid.ExtensionTopic, err error)
}

var _ ExtensionTopicsClientAPI = (*eventgrid.ExtensionTopicsClient)(nil)

// TopicTypesClientAPI contains the set of methods on the TopicTypesClient type.
type TopicTypesClientAPI interface {
	Get(ctx context.Context, topicTypeName string) (result eventgrid.TopicTypeInfo, err error)
	List(ctx context.Context) (result eventgrid.TopicTypesListResult, err error)
	ListEventTypes(ctx context.Context, topicTypeName string) (result eventgrid.EventTypesListResult, err error)
}

var _ TopicTypesClientAPI = (*eventgrid.TopicTypesClient)(nil)
