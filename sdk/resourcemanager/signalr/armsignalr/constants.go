//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsignalr

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/signalr/armsignalr"
	moduleVersion = "v2.0.0"
)

// ACLAction - Azure Networking ACL Action.
type ACLAction string

const (
	ACLActionAllow ACLAction = "Allow"
	ACLActionDeny  ACLAction = "Deny"
)

// PossibleACLActionValues returns the possible values for the ACLAction const type.
func PossibleACLActionValues() []ACLAction {
	return []ACLAction{
		ACLActionAllow,
		ACLActionDeny,
	}
}

// CreatedByType - The type of identity that created the resource.
type CreatedByType string

const (
	CreatedByTypeApplication     CreatedByType = "Application"
	CreatedByTypeKey             CreatedByType = "Key"
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	CreatedByTypeUser            CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns the possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{
		CreatedByTypeApplication,
		CreatedByTypeKey,
		CreatedByTypeManagedIdentity,
		CreatedByTypeUser,
	}
}

// FeatureFlags - FeatureFlags is the supported features of Azure SignalR service.
// * ServiceMode: Flag for backend server for SignalR service. Values allowed: "Default": have your own backend server; "Serverless":
// your application doesn't have a backend server; "Classic": for
// backward compatibility. Support both Default and Serverless mode but not recommended; "PredefinedOnly": for future use.
// * EnableConnectivityLogs: "true"/"false", to enable/disable the connectivity log category respectively.
// * EnableMessagingLogs: "true"/"false", to enable/disable the connectivity log category respectively.
// * EnableLiveTrace: Live Trace allows you to know what's happening inside Azure SignalR service, it will give you live traces
// in real time, it will be helpful when you developing your own Azure
// SignalR based web application or self-troubleshooting some issues. Please note that live traces are counted as outbound
// messages that will be charged. Values allowed: "true"/"false", to
// enable/disable live trace feature.
type FeatureFlags string

const (
	FeatureFlagsEnableConnectivityLogs FeatureFlags = "EnableConnectivityLogs"
	FeatureFlagsEnableLiveTrace        FeatureFlags = "EnableLiveTrace"
	FeatureFlagsEnableMessagingLogs    FeatureFlags = "EnableMessagingLogs"
	FeatureFlagsServiceMode            FeatureFlags = "ServiceMode"
)

// PossibleFeatureFlagsValues returns the possible values for the FeatureFlags const type.
func PossibleFeatureFlagsValues() []FeatureFlags {
	return []FeatureFlags{
		FeatureFlagsEnableConnectivityLogs,
		FeatureFlagsEnableLiveTrace,
		FeatureFlagsEnableMessagingLogs,
		FeatureFlagsServiceMode,
	}
}

// KeyType - The type of access key.
type KeyType string

const (
	KeyTypePrimary   KeyType = "Primary"
	KeyTypeSalt      KeyType = "Salt"
	KeyTypeSecondary KeyType = "Secondary"
)

// PossibleKeyTypeValues returns the possible values for the KeyType const type.
func PossibleKeyTypeValues() []KeyType {
	return []KeyType{
		KeyTypePrimary,
		KeyTypeSalt,
		KeyTypeSecondary,
	}
}

// ManagedIdentityType - Represents the identity type: systemAssigned, userAssigned, None
type ManagedIdentityType string

const (
	ManagedIdentityTypeNone           ManagedIdentityType = "None"
	ManagedIdentityTypeSystemAssigned ManagedIdentityType = "SystemAssigned"
	ManagedIdentityTypeUserAssigned   ManagedIdentityType = "UserAssigned"
)

// PossibleManagedIdentityTypeValues returns the possible values for the ManagedIdentityType const type.
func PossibleManagedIdentityTypeValues() []ManagedIdentityType {
	return []ManagedIdentityType{
		ManagedIdentityTypeNone,
		ManagedIdentityTypeSystemAssigned,
		ManagedIdentityTypeUserAssigned,
	}
}

// PrivateLinkServiceConnectionStatus - Indicates whether the connection has been Approved/Rejected/Removed by the owner of
// the service.
type PrivateLinkServiceConnectionStatus string

const (
	PrivateLinkServiceConnectionStatusApproved     PrivateLinkServiceConnectionStatus = "Approved"
	PrivateLinkServiceConnectionStatusDisconnected PrivateLinkServiceConnectionStatus = "Disconnected"
	PrivateLinkServiceConnectionStatusPending      PrivateLinkServiceConnectionStatus = "Pending"
	PrivateLinkServiceConnectionStatusRejected     PrivateLinkServiceConnectionStatus = "Rejected"
)

// PossiblePrivateLinkServiceConnectionStatusValues returns the possible values for the PrivateLinkServiceConnectionStatus const type.
func PossiblePrivateLinkServiceConnectionStatusValues() []PrivateLinkServiceConnectionStatus {
	return []PrivateLinkServiceConnectionStatus{
		PrivateLinkServiceConnectionStatusApproved,
		PrivateLinkServiceConnectionStatusDisconnected,
		PrivateLinkServiceConnectionStatusPending,
		PrivateLinkServiceConnectionStatusRejected,
	}
}

// ProvisioningState - Provisioning state of the resource.
type ProvisioningState string

const (
	ProvisioningStateCanceled  ProvisioningState = "Canceled"
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateMoving    ProvisioningState = "Moving"
	ProvisioningStateRunning   ProvisioningState = "Running"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateUnknown   ProvisioningState = "Unknown"
	ProvisioningStateUpdating  ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateCanceled,
		ProvisioningStateCreating,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateMoving,
		ProvisioningStateRunning,
		ProvisioningStateSucceeded,
		ProvisioningStateUnknown,
		ProvisioningStateUpdating,
	}
}

// ScaleType - The scale type applicable to the sku.
type ScaleType string

const (
	ScaleTypeAutomatic ScaleType = "Automatic"
	ScaleTypeManual    ScaleType = "Manual"
	ScaleTypeNone      ScaleType = "None"
)

// PossibleScaleTypeValues returns the possible values for the ScaleType const type.
func PossibleScaleTypeValues() []ScaleType {
	return []ScaleType{
		ScaleTypeAutomatic,
		ScaleTypeManual,
		ScaleTypeNone,
	}
}

// ServiceKind - The kind of the service, it can be SignalR or RawWebSockets
type ServiceKind string

const (
	ServiceKindRawWebSockets ServiceKind = "RawWebSockets"
	ServiceKindSignalR       ServiceKind = "SignalR"
)

// PossibleServiceKindValues returns the possible values for the ServiceKind const type.
func PossibleServiceKindValues() []ServiceKind {
	return []ServiceKind{
		ServiceKindRawWebSockets,
		ServiceKindSignalR,
	}
}

// SharedPrivateLinkResourceStatus - Status of the shared private link resource
type SharedPrivateLinkResourceStatus string

const (
	SharedPrivateLinkResourceStatusApproved     SharedPrivateLinkResourceStatus = "Approved"
	SharedPrivateLinkResourceStatusDisconnected SharedPrivateLinkResourceStatus = "Disconnected"
	SharedPrivateLinkResourceStatusPending      SharedPrivateLinkResourceStatus = "Pending"
	SharedPrivateLinkResourceStatusRejected     SharedPrivateLinkResourceStatus = "Rejected"
	SharedPrivateLinkResourceStatusTimeout      SharedPrivateLinkResourceStatus = "Timeout"
)

// PossibleSharedPrivateLinkResourceStatusValues returns the possible values for the SharedPrivateLinkResourceStatus const type.
func PossibleSharedPrivateLinkResourceStatusValues() []SharedPrivateLinkResourceStatus {
	return []SharedPrivateLinkResourceStatus{
		SharedPrivateLinkResourceStatusApproved,
		SharedPrivateLinkResourceStatusDisconnected,
		SharedPrivateLinkResourceStatusPending,
		SharedPrivateLinkResourceStatusRejected,
		SharedPrivateLinkResourceStatusTimeout,
	}
}

// SignalRRequestType - The incoming request type to the service
type SignalRRequestType string

const (
	SignalRRequestTypeClientConnection SignalRRequestType = "ClientConnection"
	SignalRRequestTypeRESTAPI          SignalRRequestType = "RESTAPI"
	SignalRRequestTypeServerConnection SignalRRequestType = "ServerConnection"
	SignalRRequestTypeTrace            SignalRRequestType = "Trace"
)

// PossibleSignalRRequestTypeValues returns the possible values for the SignalRRequestType const type.
func PossibleSignalRRequestTypeValues() []SignalRRequestType {
	return []SignalRRequestType{
		SignalRRequestTypeClientConnection,
		SignalRRequestTypeRESTAPI,
		SignalRRequestTypeServerConnection,
		SignalRRequestTypeTrace,
	}
}

// SignalRSKUTier - Optional tier of this particular SKU. 'Standard' or 'Free'.
// Basic is deprecated, use Standard instead.
type SignalRSKUTier string

const (
	SignalRSKUTierBasic    SignalRSKUTier = "Basic"
	SignalRSKUTierFree     SignalRSKUTier = "Free"
	SignalRSKUTierPremium  SignalRSKUTier = "Premium"
	SignalRSKUTierStandard SignalRSKUTier = "Standard"
)

// PossibleSignalRSKUTierValues returns the possible values for the SignalRSKUTier const type.
func PossibleSignalRSKUTierValues() []SignalRSKUTier {
	return []SignalRSKUTier{
		SignalRSKUTierBasic,
		SignalRSKUTierFree,
		SignalRSKUTierPremium,
		SignalRSKUTierStandard,
	}
}

// UpstreamAuthType - Upstream auth type enum.
type UpstreamAuthType string

const (
	UpstreamAuthTypeManagedIdentity UpstreamAuthType = "ManagedIdentity"
	UpstreamAuthTypeNone            UpstreamAuthType = "None"
)

// PossibleUpstreamAuthTypeValues returns the possible values for the UpstreamAuthType const type.
func PossibleUpstreamAuthTypeValues() []UpstreamAuthType {
	return []UpstreamAuthType{
		UpstreamAuthTypeManagedIdentity,
		UpstreamAuthTypeNone,
	}
}
