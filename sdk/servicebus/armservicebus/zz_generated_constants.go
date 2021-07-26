// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armservicebus

const telemetryInfo = "azsdk-go-armservicebus/v0.1.0"

type AccessRights string

const (
	AccessRightsManage AccessRights = "Manage"
	AccessRightsSend   AccessRights = "Send"
	AccessRightsListen AccessRights = "Listen"
)

// PossibleAccessRightsValues returns the possible values for the AccessRights const type.
func PossibleAccessRightsValues() []AccessRights {
	return []AccessRights{
		AccessRightsManage,
		AccessRightsSend,
		AccessRightsListen,
	}
}

// ToPtr returns a *AccessRights pointing to the current value.
func (c AccessRights) ToPtr() *AccessRights {
	return &c
}

// DefaultAction - Default Action for Network Rule Set
type DefaultAction string

const (
	DefaultActionAllow DefaultAction = "Allow"
	DefaultActionDeny  DefaultAction = "Deny"
)

// PossibleDefaultActionValues returns the possible values for the DefaultAction const type.
func PossibleDefaultActionValues() []DefaultAction {
	return []DefaultAction{
		DefaultActionAllow,
		DefaultActionDeny,
	}
}

// ToPtr returns a *DefaultAction pointing to the current value.
func (c DefaultAction) ToPtr() *DefaultAction {
	return &c
}

// EncodingCaptureDescription - Enumerates the possible values for the encoding format of capture description.
type EncodingCaptureDescription string

const (
	EncodingCaptureDescriptionAvro        EncodingCaptureDescription = "Avro"
	EncodingCaptureDescriptionAvroDeflate EncodingCaptureDescription = "AvroDeflate"
)

// PossibleEncodingCaptureDescriptionValues returns the possible values for the EncodingCaptureDescription const type.
func PossibleEncodingCaptureDescriptionValues() []EncodingCaptureDescription {
	return []EncodingCaptureDescription{
		EncodingCaptureDescriptionAvro,
		EncodingCaptureDescriptionAvroDeflate,
	}
}

// ToPtr returns a *EncodingCaptureDescription pointing to the current value.
func (c EncodingCaptureDescription) ToPtr() *EncodingCaptureDescription {
	return &c
}

// EndPointProvisioningState - Provisioning state of the Private Endpoint Connection.
type EndPointProvisioningState string

const (
	EndPointProvisioningStateCanceled  EndPointProvisioningState = "Canceled"
	EndPointProvisioningStateCreating  EndPointProvisioningState = "Creating"
	EndPointProvisioningStateDeleting  EndPointProvisioningState = "Deleting"
	EndPointProvisioningStateFailed    EndPointProvisioningState = "Failed"
	EndPointProvisioningStateSucceeded EndPointProvisioningState = "Succeeded"
	EndPointProvisioningStateUpdating  EndPointProvisioningState = "Updating"
)

// PossibleEndPointProvisioningStateValues returns the possible values for the EndPointProvisioningState const type.
func PossibleEndPointProvisioningStateValues() []EndPointProvisioningState {
	return []EndPointProvisioningState{
		EndPointProvisioningStateCanceled,
		EndPointProvisioningStateCreating,
		EndPointProvisioningStateDeleting,
		EndPointProvisioningStateFailed,
		EndPointProvisioningStateSucceeded,
		EndPointProvisioningStateUpdating,
	}
}

// ToPtr returns a *EndPointProvisioningState pointing to the current value.
func (c EndPointProvisioningState) ToPtr() *EndPointProvisioningState {
	return &c
}

// EntityStatus - Entity status.
type EntityStatus string

const (
	EntityStatusActive          EntityStatus = "Active"
	EntityStatusDisabled        EntityStatus = "Disabled"
	EntityStatusRestoring       EntityStatus = "Restoring"
	EntityStatusSendDisabled    EntityStatus = "SendDisabled"
	EntityStatusReceiveDisabled EntityStatus = "ReceiveDisabled"
	EntityStatusCreating        EntityStatus = "Creating"
	EntityStatusDeleting        EntityStatus = "Deleting"
	EntityStatusRenaming        EntityStatus = "Renaming"
	EntityStatusUnknown         EntityStatus = "Unknown"
)

// PossibleEntityStatusValues returns the possible values for the EntityStatus const type.
func PossibleEntityStatusValues() []EntityStatus {
	return []EntityStatus{
		EntityStatusActive,
		EntityStatusDisabled,
		EntityStatusRestoring,
		EntityStatusSendDisabled,
		EntityStatusReceiveDisabled,
		EntityStatusCreating,
		EntityStatusDeleting,
		EntityStatusRenaming,
		EntityStatusUnknown,
	}
}

// ToPtr returns a *EntityStatus pointing to the current value.
func (c EntityStatus) ToPtr() *EntityStatus {
	return &c
}

// FilterType - Rule filter types
type FilterType string

const (
	FilterTypeSQLFilter         FilterType = "SqlFilter"
	FilterTypeCorrelationFilter FilterType = "CorrelationFilter"
)

// PossibleFilterTypeValues returns the possible values for the FilterType const type.
func PossibleFilterTypeValues() []FilterType {
	return []FilterType{
		FilterTypeSQLFilter,
		FilterTypeCorrelationFilter,
	}
}

// ToPtr returns a *FilterType pointing to the current value.
func (c FilterType) ToPtr() *FilterType {
	return &c
}

// IPAction - The IP Filter Action
type IPAction string

const (
	IPActionAccept IPAction = "Accept"
	IPActionReject IPAction = "Reject"
)

// PossibleIPActionValues returns the possible values for the IPAction const type.
func PossibleIPActionValues() []IPAction {
	return []IPAction{
		IPActionAccept,
		IPActionReject,
	}
}

// ToPtr returns a *IPAction pointing to the current value.
func (c IPAction) ToPtr() *IPAction {
	return &c
}

// KeyType - The access key to regenerate.
type KeyType string

const (
	KeyTypePrimaryKey   KeyType = "PrimaryKey"
	KeyTypeSecondaryKey KeyType = "SecondaryKey"
)

// PossibleKeyTypeValues returns the possible values for the KeyType const type.
func PossibleKeyTypeValues() []KeyType {
	return []KeyType{
		KeyTypePrimaryKey,
		KeyTypeSecondaryKey,
	}
}

// ToPtr returns a *KeyType pointing to the current value.
func (c KeyType) ToPtr() *KeyType {
	return &c
}

type MigrationConfigurationName string

const (
	MigrationConfigurationNameDefault MigrationConfigurationName = "$default"
)

// PossibleMigrationConfigurationNameValues returns the possible values for the MigrationConfigurationName const type.
func PossibleMigrationConfigurationNameValues() []MigrationConfigurationName {
	return []MigrationConfigurationName{
		MigrationConfigurationNameDefault,
	}
}

// ToPtr returns a *MigrationConfigurationName pointing to the current value.
func (c MigrationConfigurationName) ToPtr() *MigrationConfigurationName {
	return &c
}

// NameSpaceType - Type of namespaces
type NameSpaceType string

const (
	NameSpaceTypeMessaging       NameSpaceType = "Messaging"
	NameSpaceTypeNotificationHub NameSpaceType = "NotificationHub"
	NameSpaceTypeMixed           NameSpaceType = "Mixed"
	NameSpaceTypeEventHub        NameSpaceType = "EventHub"
	NameSpaceTypeRelay           NameSpaceType = "Relay"
)

// PossibleNameSpaceTypeValues returns the possible values for the NameSpaceType const type.
func PossibleNameSpaceTypeValues() []NameSpaceType {
	return []NameSpaceType{
		NameSpaceTypeMessaging,
		NameSpaceTypeNotificationHub,
		NameSpaceTypeMixed,
		NameSpaceTypeEventHub,
		NameSpaceTypeRelay,
	}
}

// ToPtr returns a *NameSpaceType pointing to the current value.
func (c NameSpaceType) ToPtr() *NameSpaceType {
	return &c
}

// NetworkRuleIPAction - The IP Filter Action
type NetworkRuleIPAction string

const (
	NetworkRuleIPActionAllow NetworkRuleIPAction = "Allow"
)

// PossibleNetworkRuleIPActionValues returns the possible values for the NetworkRuleIPAction const type.
func PossibleNetworkRuleIPActionValues() []NetworkRuleIPAction {
	return []NetworkRuleIPAction{
		NetworkRuleIPActionAllow,
	}
}

// ToPtr returns a *NetworkRuleIPAction pointing to the current value.
func (c NetworkRuleIPAction) ToPtr() *NetworkRuleIPAction {
	return &c
}

// PrivateLinkConnectionStatus - Status of the connection.
type PrivateLinkConnectionStatus string

const (
	PrivateLinkConnectionStatusApproved     PrivateLinkConnectionStatus = "Approved"
	PrivateLinkConnectionStatusDisconnected PrivateLinkConnectionStatus = "Disconnected"
	PrivateLinkConnectionStatusPending      PrivateLinkConnectionStatus = "Pending"
	PrivateLinkConnectionStatusRejected     PrivateLinkConnectionStatus = "Rejected"
)

// PossiblePrivateLinkConnectionStatusValues returns the possible values for the PrivateLinkConnectionStatus const type.
func PossiblePrivateLinkConnectionStatusValues() []PrivateLinkConnectionStatus {
	return []PrivateLinkConnectionStatus{
		PrivateLinkConnectionStatusApproved,
		PrivateLinkConnectionStatusDisconnected,
		PrivateLinkConnectionStatusPending,
		PrivateLinkConnectionStatusRejected,
	}
}

// ToPtr returns a *PrivateLinkConnectionStatus pointing to the current value.
func (c PrivateLinkConnectionStatus) ToPtr() *PrivateLinkConnectionStatus {
	return &c
}

// ProvisioningStateDR - Provisioning state of the Alias(Disaster Recovery configuration) - possible values 'Accepted' or 'Succeeded' or 'Failed'
type ProvisioningStateDR string

const (
	ProvisioningStateDRAccepted  ProvisioningStateDR = "Accepted"
	ProvisioningStateDRSucceeded ProvisioningStateDR = "Succeeded"
	ProvisioningStateDRFailed    ProvisioningStateDR = "Failed"
)

// PossibleProvisioningStateDRValues returns the possible values for the ProvisioningStateDR const type.
func PossibleProvisioningStateDRValues() []ProvisioningStateDR {
	return []ProvisioningStateDR{
		ProvisioningStateDRAccepted,
		ProvisioningStateDRSucceeded,
		ProvisioningStateDRFailed,
	}
}

// ToPtr returns a *ProvisioningStateDR pointing to the current value.
func (c ProvisioningStateDR) ToPtr() *ProvisioningStateDR {
	return &c
}

// RoleDisasterRecovery - role of namespace in GEO DR - possible values 'Primary' or 'PrimaryNotReplicating' or 'Secondary'
type RoleDisasterRecovery string

const (
	RoleDisasterRecoveryPrimary               RoleDisasterRecovery = "Primary"
	RoleDisasterRecoveryPrimaryNotReplicating RoleDisasterRecovery = "PrimaryNotReplicating"
	RoleDisasterRecoverySecondary             RoleDisasterRecovery = "Secondary"
)

// PossibleRoleDisasterRecoveryValues returns the possible values for the RoleDisasterRecovery const type.
func PossibleRoleDisasterRecoveryValues() []RoleDisasterRecovery {
	return []RoleDisasterRecovery{
		RoleDisasterRecoveryPrimary,
		RoleDisasterRecoveryPrimaryNotReplicating,
		RoleDisasterRecoverySecondary,
	}
}

// ToPtr returns a *RoleDisasterRecovery pointing to the current value.
func (c RoleDisasterRecovery) ToPtr() *RoleDisasterRecovery {
	return &c
}

// SKUName - Name of this SKU.
type SKUName string

const (
	SKUNameBasic    SKUName = "Basic"
	SKUNameStandard SKUName = "Standard"
	SKUNamePremium  SKUName = "Premium"
)

// PossibleSKUNameValues returns the possible values for the SKUName const type.
func PossibleSKUNameValues() []SKUName {
	return []SKUName{
		SKUNameBasic,
		SKUNameStandard,
		SKUNamePremium,
	}
}

// ToPtr returns a *SKUName pointing to the current value.
func (c SKUName) ToPtr() *SKUName {
	return &c
}

// SKUTier - The billing tier of this particular SKU.
type SKUTier string

const (
	SKUTierBasic    SKUTier = "Basic"
	SKUTierStandard SKUTier = "Standard"
	SKUTierPremium  SKUTier = "Premium"
)

// PossibleSKUTierValues returns the possible values for the SKUTier const type.
func PossibleSKUTierValues() []SKUTier {
	return []SKUTier{
		SKUTierBasic,
		SKUTierStandard,
		SKUTierPremium,
	}
}

// ToPtr returns a *SKUTier pointing to the current value.
func (c SKUTier) ToPtr() *SKUTier {
	return &c
}

// UnavailableReason - Specifies the reason for the unavailability of the service.
type UnavailableReason string

const (
	UnavailableReasonNone                                  UnavailableReason = "None"
	UnavailableReasonInvalidName                           UnavailableReason = "InvalidName"
	UnavailableReasonSubscriptionIsDisabled                UnavailableReason = "SubscriptionIsDisabled"
	UnavailableReasonNameInUse                             UnavailableReason = "NameInUse"
	UnavailableReasonNameInLockdown                        UnavailableReason = "NameInLockdown"
	UnavailableReasonTooManyNamespaceInCurrentSubscription UnavailableReason = "TooManyNamespaceInCurrentSubscription"
)

// PossibleUnavailableReasonValues returns the possible values for the UnavailableReason const type.
func PossibleUnavailableReasonValues() []UnavailableReason {
	return []UnavailableReason{
		UnavailableReasonNone,
		UnavailableReasonInvalidName,
		UnavailableReasonSubscriptionIsDisabled,
		UnavailableReasonNameInUse,
		UnavailableReasonNameInLockdown,
		UnavailableReasonTooManyNamespaceInCurrentSubscription,
	}
}

// ToPtr returns a *UnavailableReason pointing to the current value.
func (c UnavailableReason) ToPtr() *UnavailableReason {
	return &c
}
