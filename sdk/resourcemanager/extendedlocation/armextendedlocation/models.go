//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armextendedlocation

import "time"

// CustomLocation - Custom Locations definition.
type CustomLocation struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// Identity for the resource.
	Identity *Identity

	// The set of properties specific to a Custom Location
	Properties *CustomLocationProperties

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// CustomLocationFindTargetResourceGroupProperties - The Find Target Resource Group operation request.
type CustomLocationFindTargetResourceGroupProperties struct {
	// Labels of the custom resource, this is a map of {key,value} pairs.
	Labels map[string]*string
}

// CustomLocationFindTargetResourceGroupResult - The Find Target Resource Group operation response.
type CustomLocationFindTargetResourceGroupResult struct {
	// READ-ONLY; The matching resource sync rule is the particular resource sync rule that matched the match expressions and
	// labels and had lowest priority. This is the rule responsible for mapping the target resource
	// to the target resource group.
	MatchedResourceSyncRule *string

	// READ-ONLY; The target resource group of matching resource sync rule. The labels from the request will be used to find out
	// matching resource sync rule against the selector property of the resource sync rule. The
	// one with highest priority will be returned if there are multiple matching rules.
	TargetResourceGroup *string
}

// CustomLocationListResult - The List Custom Locations operation response.
type CustomLocationListResult struct {
	// READ-ONLY; The URL to use for getting the next set of results.
	NextLink *string

	// READ-ONLY; The list of Custom Locations.
	Value []*CustomLocation
}

// CustomLocationOperation - Custom Locations operation.
type CustomLocationOperation struct {
	// Describes the properties of a Custom Locations Operation Value Display.
	Display *CustomLocationOperationValueDisplay

	// READ-ONLY; Is this Operation a data plane operation
	IsDataAction *bool

	// READ-ONLY; The name of the compute operation.
	Name *string

	// READ-ONLY; The origin of the compute operation.
	Origin *string
}

// CustomLocationOperationValueDisplay - Describes the properties of a Custom Locations Operation Value Display.
type CustomLocationOperationValueDisplay struct {
	// READ-ONLY; The description of the operation.
	Description *string

	// READ-ONLY; The display name of the compute operation.
	Operation *string

	// READ-ONLY; The resource provider for the operation.
	Provider *string

	// READ-ONLY; The display name of the resource the operation applies to.
	Resource *string
}

// CustomLocationOperationsList - Lists of Custom Locations operations.
type CustomLocationOperationsList struct {
	// REQUIRED; Array of customLocationOperation
	Value []*CustomLocationOperation

	// Next page of operations.
	NextLink *string
}

// CustomLocationProperties - Properties for a custom location.
type CustomLocationProperties struct {
	// This is optional input that contains the authentication that should be used to generate the namespace.
	Authentication *CustomLocationPropertiesAuthentication

	// Contains the reference to the add-on that contains charts to deploy CRDs and operators.
	ClusterExtensionIDs []*string

	// Display name for the Custom Locations location.
	DisplayName *string

	// Connected Cluster or AKS Cluster. The Custom Locations RP will perform a checkAccess API for listAdminCredentials permissions.
	HostResourceID *string

	// Type of host the Custom Locations is referencing (Kubernetes, etc…).
	HostType *HostType

	// Kubernetes namespace that will be created on the specified cluster.
	Namespace *string

	// Provisioning State for the Custom Location.
	ProvisioningState *string
}

// CustomLocationPropertiesAuthentication - This is optional input that contains the authentication that should be used to
// generate the namespace.
type CustomLocationPropertiesAuthentication struct {
	// The type of the Custom Locations authentication
	Type *string

	// The kubeconfig value.
	Value *string
}

// CustomLocationsClientBeginCreateOrUpdateOptions contains the optional parameters for the CustomLocationsClient.BeginCreateOrUpdate
// method.
type CustomLocationsClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// CustomLocationsClientBeginDeleteOptions contains the optional parameters for the CustomLocationsClient.BeginDelete method.
type CustomLocationsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// CustomLocationsClientFindTargetResourceGroupOptions contains the optional parameters for the CustomLocationsClient.FindTargetResourceGroup
// method.
type CustomLocationsClientFindTargetResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// CustomLocationsClientGetOptions contains the optional parameters for the CustomLocationsClient.Get method.
type CustomLocationsClientGetOptions struct {
	// placeholder for future optional parameters
}

// CustomLocationsClientListByResourceGroupOptions contains the optional parameters for the CustomLocationsClient.NewListByResourceGroupPager
// method.
type CustomLocationsClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// CustomLocationsClientListBySubscriptionOptions contains the optional parameters for the CustomLocationsClient.NewListBySubscriptionPager
// method.
type CustomLocationsClientListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// CustomLocationsClientListEnabledResourceTypesOptions contains the optional parameters for the CustomLocationsClient.NewListEnabledResourceTypesPager
// method.
type CustomLocationsClientListEnabledResourceTypesOptions struct {
	// placeholder for future optional parameters
}

// CustomLocationsClientListOperationsOptions contains the optional parameters for the CustomLocationsClient.NewListOperationsPager
// method.
type CustomLocationsClientListOperationsOptions struct {
	// placeholder for future optional parameters
}

// CustomLocationsClientUpdateOptions contains the optional parameters for the CustomLocationsClient.Update method.
type CustomLocationsClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// EnabledResourceType definition.
type EnabledResourceType struct {
	// The set of properties for EnabledResourceType specific to a Custom Location
	Properties *EnabledResourceTypeProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// EnabledResourceTypeProperties - Properties for EnabledResourceType of a custom location.
type EnabledResourceTypeProperties struct {
	// Cluster Extension ID
	ClusterExtensionID *string

	// Cluster Extension Type
	ExtensionType *string

	// Metadata of the Resource Type
	TypesMetadata []*EnabledResourceTypePropertiesTypesMetadataItem
}

// EnabledResourceTypePropertiesTypesMetadataItem - Metadata of the Resource Type.
type EnabledResourceTypePropertiesTypesMetadataItem struct {
	// Api Version of Resource Type
	APIVersion *string

	// Resource Provider Namespace of Resource Type
	ResourceProviderNamespace *string

	// Resource Type
	ResourceType *string
}

// EnabledResourceTypesListResult - List of EnabledResourceTypes definition.
type EnabledResourceTypesListResult struct {
	// READ-ONLY; The URL to use for getting the next set of results.
	NextLink *string

	// READ-ONLY; The list of EnabledResourceTypes available for a customLocation.
	Value []*EnabledResourceType
}

// ErrorAdditionalInfo - The resource management error additional info.
type ErrorAdditionalInfo struct {
	// READ-ONLY; The additional info.
	Info any

	// READ-ONLY; The additional info type.
	Type *string
}

// ErrorDetail - The error detail.
type ErrorDetail struct {
	// READ-ONLY; The error additional info.
	AdditionalInfo []*ErrorAdditionalInfo

	// READ-ONLY; The error code.
	Code *string

	// READ-ONLY; The error details.
	Details []*ErrorDetail

	// READ-ONLY; The error message.
	Message *string

	// READ-ONLY; The error target.
	Target *string
}

// ErrorResponse - Common error response for all Azure Resource Manager APIs to return error details for failed operations.
// (This also follows the OData error response format.).
type ErrorResponse struct {
	// The error object.
	Error *ErrorDetail
}

// Identity for the resource.
type Identity struct {
	// The identity type.
	Type *ResourceIdentityType

	// READ-ONLY; The principal ID of resource identity.
	PrincipalID *string

	// READ-ONLY; The tenant ID of resource.
	TenantID *string
}

// MatchExpressionsProperties - Resource Sync Rules matchExpression property definition.
type MatchExpressionsProperties struct {
	// Key is the label key that the selector applies to.
	Key *string

	// The Operator field represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.
	Operator *string

	// The label value
	Values []*string
}

// PatchableCustomLocations - The Custom Locations patchable resource definition.
type PatchableCustomLocations struct {
	// Identity for the resource.
	Identity *Identity

	// The Custom Locations patchable properties.
	Properties *CustomLocationProperties

	// Resource tags
	Tags map[string]*string
}

// PatchableResourceSyncRule - The Resource Sync Rules patchable resource definition.
type PatchableResourceSyncRule struct {
	// The Resource Sync Rules patchable properties.
	Properties *ResourceSyncRuleProperties

	// Resource tags
	Tags map[string]*string
}

// ProxyResource - The resource model definition for a Azure Resource Manager proxy resource. It will not have tags and a
// location
type ProxyResource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// Resource - Common fields that are returned in the response for all Azure Resource Manager resources
type Resource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// ResourceSyncRule - Resource Sync Rules definition.
type ResourceSyncRule struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// The set of properties specific to a Resource Sync Rule
	Properties *ResourceSyncRuleProperties

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// ResourceSyncRuleListResult - The List Resource Sync Rules operation response.
type ResourceSyncRuleListResult struct {
	// READ-ONLY; The URL to use for getting the next set of results.
	NextLink *string

	// READ-ONLY; The list of Resource Sync Rules.
	Value []*ResourceSyncRule
}

// ResourceSyncRuleProperties - Properties for a resource sync rule. For an unmapped custom resource, its labels will be used
// to find out matching resource sync rules using the selector property of the resource sync rule. If this
// resource sync rule has highest priority among all matching rules, then the unmapped custom resource will be projected to
// the target resource group associated with this resource sync rule.
type ResourceSyncRuleProperties struct {
	// Priority represents a priority of the Resource Sync Rule
	Priority *int32

	// A label selector is composed of two parts, matchLabels and matchExpressions. The first part, matchLabels is a map of {key,value}
	// pairs. A single {key,value} in the matchLabels map is equivalent to an
	// element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'.
	// The second part, matchExpressions is a list of resource selector requirements.
	// Valid operators include In, NotIn, Exists, and DoesNotExist. The values set must be non-empty in the case of In and NotIn.
	// The values set must be empty in the case of Exists and DoesNotExist. All of
	// the requirements, from both matchLabels and matchExpressions must all be satisfied in order to match.
	Selector *ResourceSyncRulePropertiesSelector

	// For an unmapped custom resource, its labels will be used to find matching resource sync rules. If this resource sync rule
	// is one of the matching rules with highest priority, then the unmapped custom
	// resource will be projected to the target resource group associated with this resource sync rule. The user creating this
	// resource sync rule should have write permissions on the target resource group
	// and this write permission will be validated when creating the resource sync rule.
	TargetResourceGroup *string

	// READ-ONLY; Provisioning State for the Resource Sync Rule.
	ProvisioningState *string
}

// ResourceSyncRulePropertiesSelector - A label selector is composed of two parts, matchLabels and matchExpressions. The first
// part, matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an
// element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'.
// The second part, matchExpressions is a list of resource selector requirements.
// Valid operators include In, NotIn, Exists, and DoesNotExist. The values set must be non-empty in the case of In and NotIn.
// The values set must be empty in the case of Exists and DoesNotExist. All of
// the requirements, from both matchLabels and matchExpressions must all be satisfied in order to match.
type ResourceSyncRulePropertiesSelector struct {
	// MatchExpressions is a list of resource selector requirements. Valid operators include In, NotIn, Exists, and DoesNotExist.
	// The values set must be non-empty in the case of In and NotIn. The values set
	// must be empty in the case of Exists and DoesNotExist.
	MatchExpressions []*MatchExpressionsProperties

	// MatchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions,
	// whose key field is 'key', the operator is 'In', and the values
	// array contains only 'value'.
	MatchLabels map[string]*string
}

// ResourceSyncRulesClientBeginCreateOrUpdateOptions contains the optional parameters for the ResourceSyncRulesClient.BeginCreateOrUpdate
// method.
type ResourceSyncRulesClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ResourceSyncRulesClientBeginUpdateOptions contains the optional parameters for the ResourceSyncRulesClient.BeginUpdate
// method.
type ResourceSyncRulesClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ResourceSyncRulesClientDeleteOptions contains the optional parameters for the ResourceSyncRulesClient.Delete method.
type ResourceSyncRulesClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// ResourceSyncRulesClientGetOptions contains the optional parameters for the ResourceSyncRulesClient.Get method.
type ResourceSyncRulesClientGetOptions struct {
	// placeholder for future optional parameters
}

// ResourceSyncRulesClientListByCustomLocationIDOptions contains the optional parameters for the ResourceSyncRulesClient.NewListByCustomLocationIDPager
// method.
type ResourceSyncRulesClientListByCustomLocationIDOptions struct {
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

// TrackedResource - The resource model definition for an Azure Resource Manager tracked top level resource which has 'tags'
// and a 'location'
type TrackedResource struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}
