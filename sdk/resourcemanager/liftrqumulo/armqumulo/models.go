//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armqumulo

import "time"

// FileSystemResource - A FileSystem Resource by Qumulo
type FileSystemResource struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// REQUIRED; The resource-specific properties for this resource.
	Properties *FileSystemResourceProperties

	// The managed service identities assigned to this resource.
	Identity *ManagedServiceIdentity

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// FileSystemResourceListResult - The response of a FileSystemResource list operation.
type FileSystemResourceListResult struct {
	// REQUIRED; The FileSystemResource items on this page
	Value []*FileSystemResource

	// The link to the next page of items
	NextLink *string
}

// FileSystemResourceProperties - Properties specific to the Qumulo File System resource
type FileSystemResourceProperties struct {
	// REQUIRED; Initial administrator password of the resource
	AdminPassword *string

	// REQUIRED; Delegated subnet id for Vnet injection
	DelegatedSubnetID *string

	// REQUIRED; Storage capacity in TB
	InitialCapacity *int32

	// REQUIRED; Marketplace details
	MarketplaceDetails *MarketplaceDetails

	// REQUIRED; Storage Sku
	StorageSKU *StorageSKU

	// REQUIRED; User Details
	UserDetails *UserDetails

	// Availability zone
	AvailabilityZone *string

	// File system Id of the resource
	ClusterLoginURL *string

	// Private IPs of the resource
	PrivateIPs []*string

	// READ-ONLY; Provisioning State of the resource
	ProvisioningState *ProvisioningState
}

// FileSystemResourceUpdate - The type used for update operations of the FileSystemResource.
type FileSystemResourceUpdate struct {
	// The managed service identities assigned to this resource.
	Identity *ManagedServiceIdentity

	// The updatable properties of the FileSystemResource.
	Properties *FileSystemResourceUpdateProperties

	// Resource tags.
	Tags map[string]*string
}

// FileSystemResourceUpdateProperties - The updatable properties of the FileSystemResource.
type FileSystemResourceUpdateProperties struct {
	// File system Id of the resource
	ClusterLoginURL *string

	// Delegated subnet id for Vnet injection
	DelegatedSubnetID *string

	// Marketplace details
	MarketplaceDetails *MarketplaceDetails

	// Private IPs of the resource
	PrivateIPs []*string

	// User Details
	UserDetails *UserDetails
}

// FileSystemsClientBeginCreateOrUpdateOptions contains the optional parameters for the FileSystemsClient.BeginCreateOrUpdate
// method.
type FileSystemsClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// FileSystemsClientBeginDeleteOptions contains the optional parameters for the FileSystemsClient.BeginDelete method.
type FileSystemsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// FileSystemsClientGetOptions contains the optional parameters for the FileSystemsClient.Get method.
type FileSystemsClientGetOptions struct {
	// placeholder for future optional parameters
}

// FileSystemsClientListByResourceGroupOptions contains the optional parameters for the FileSystemsClient.NewListByResourceGroupPager
// method.
type FileSystemsClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// FileSystemsClientListBySubscriptionOptions contains the optional parameters for the FileSystemsClient.NewListBySubscriptionPager
// method.
type FileSystemsClientListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// FileSystemsClientUpdateOptions contains the optional parameters for the FileSystemsClient.Update method.
type FileSystemsClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// ManagedServiceIdentity - Managed service identity (system assigned and/or user assigned identities)
type ManagedServiceIdentity struct {
	// REQUIRED; Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
	Type *ManagedServiceIdentityType

	// The set of user assigned identities associated with the resource. The userAssignedIdentities dictionary keys will be ARM
	// resource ids in the form:
	// '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}.
	// The dictionary values can be empty objects ({}) in
	// requests.
	UserAssignedIdentities map[string]*UserAssignedIdentity

	// READ-ONLY; The service principal ID of the system assigned identity. This property will only be provided for a system assigned
	// identity.
	PrincipalID *string

	// READ-ONLY; The tenant ID of the system assigned identity. This property will only be provided for a system assigned identity.
	TenantID *string
}

// MarketplaceDetails of Qumulo FileSystem resource
type MarketplaceDetails struct {
	// REQUIRED; Offer Id
	OfferID *string

	// REQUIRED; Plan Id
	PlanID *string

	// REQUIRED; Publisher Id
	PublisherID *string

	// Marketplace Subscription Id
	MarketplaceSubscriptionID *string

	// READ-ONLY; Marketplace subscription status
	MarketplaceSubscriptionStatus *MarketplaceSubscriptionStatus
}

// Operation - Details of a REST API operation, returned from the Resource Provider Operations API
type Operation struct {
	// Localized display information for this particular operation.
	Display *OperationDisplay

	// READ-ONLY; Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
	ActionType *ActionType

	// READ-ONLY; Whether the operation applies to data-plane. This is "true" for data-plane operations and "false" for ARM/control-plane
	// operations.
	IsDataAction *bool

	// READ-ONLY; The name of the operation, as per Resource-Based Access Control (RBAC). Examples: "Microsoft.Compute/virtualMachines/write",
	// "Microsoft.Compute/virtualMachines/capture/action"
	Name *string

	// READ-ONLY; The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
	// value is "user,system"
	Origin *Origin
}

// OperationDisplay - Localized display information for this particular operation.
type OperationDisplay struct {
	// READ-ONLY; The short, localized friendly description of the operation; suitable for tool tips and detailed views.
	Description *string

	// READ-ONLY; The concise, localized friendly name for the operation; suitable for dropdowns. E.g. "Create or Update Virtual
	// Machine", "Restart Virtual Machine".
	Operation *string

	// READ-ONLY; The localized friendly form of the resource provider name, e.g. "Microsoft Monitoring Insights" or "Microsoft
	// Compute".
	Provider *string

	// READ-ONLY; The localized friendly name of the resource type related to this operation. E.g. "Virtual Machines" or "Job
	// Schedule Collections".
	Resource *string
}

// OperationListResult - A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to
// get the next set of results.
type OperationListResult struct {
	// READ-ONLY; URL to get the next set of operation list results (if there are any).
	NextLink *string

	// READ-ONLY; List of operations supported by the resource provider
	Value []*Operation
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.NewListPager method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time

	// The identity that created the resource.
	CreatedBy *string

	// The type of identity that created the resource.
	CreatedByType *CreatedByType

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time

	// The identity that last modified the resource.
	LastModifiedBy *string

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType
}

// UserAssignedIdentity - User assigned identity properties
type UserAssignedIdentity struct {
	// READ-ONLY; The client ID of the assigned identity.
	ClientID *string

	// READ-ONLY; The principal ID of the assigned identity.
	PrincipalID *string
}

// UserDetails - User Details of Qumulo FileSystem resource
type UserDetails struct {
	// REQUIRED; User Email
	Email *string
}
