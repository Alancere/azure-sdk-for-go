//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetapp

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp"
	moduleVersion = "v7.0.0"
)

// ActiveDirectoryStatus - Status of the Active Directory
type ActiveDirectoryStatus string

const (
	// ActiveDirectoryStatusCreated - Active Directory created but not in use
	ActiveDirectoryStatusCreated ActiveDirectoryStatus = "Created"
	// ActiveDirectoryStatusDeleted - Active Directory Deleted
	ActiveDirectoryStatusDeleted ActiveDirectoryStatus = "Deleted"
	// ActiveDirectoryStatusError - Error with the Active Directory
	ActiveDirectoryStatusError ActiveDirectoryStatus = "Error"
	// ActiveDirectoryStatusInUse - Active Directory in use by SMB Volume
	ActiveDirectoryStatusInUse ActiveDirectoryStatus = "InUse"
	// ActiveDirectoryStatusUpdating - Active Directory Updating
	ActiveDirectoryStatusUpdating ActiveDirectoryStatus = "Updating"
)

// PossibleActiveDirectoryStatusValues returns the possible values for the ActiveDirectoryStatus const type.
func PossibleActiveDirectoryStatusValues() []ActiveDirectoryStatus {
	return []ActiveDirectoryStatus{
		ActiveDirectoryStatusCreated,
		ActiveDirectoryStatusDeleted,
		ActiveDirectoryStatusError,
		ActiveDirectoryStatusInUse,
		ActiveDirectoryStatusUpdating,
	}
}

// ApplicationType - Application Type
type ApplicationType string

const (
	ApplicationTypeORACLE  ApplicationType = "ORACLE"
	ApplicationTypeSAPHANA ApplicationType = "SAP-HANA"
)

// PossibleApplicationTypeValues returns the possible values for the ApplicationType const type.
func PossibleApplicationTypeValues() []ApplicationType {
	return []ApplicationType{
		ApplicationTypeORACLE,
		ApplicationTypeSAPHANA,
	}
}

// AvsDataStore - Specifies whether the volume is enabled for Azure VMware Solution (AVS) datastore purpose
type AvsDataStore string

const (
	// AvsDataStoreDisabled - avsDataStore is disabled
	AvsDataStoreDisabled AvsDataStore = "Disabled"
	// AvsDataStoreEnabled - avsDataStore is enabled
	AvsDataStoreEnabled AvsDataStore = "Enabled"
)

// PossibleAvsDataStoreValues returns the possible values for the AvsDataStore const type.
func PossibleAvsDataStoreValues() []AvsDataStore {
	return []AvsDataStore{
		AvsDataStoreDisabled,
		AvsDataStoreEnabled,
	}
}

// BackupType - Type of backup Manual or Scheduled
type BackupType string

const (
	// BackupTypeManual - Manual backup
	BackupTypeManual BackupType = "Manual"
	// BackupTypeScheduled - Scheduled backup
	BackupTypeScheduled BackupType = "Scheduled"
)

// PossibleBackupTypeValues returns the possible values for the BackupType const type.
func PossibleBackupTypeValues() []BackupType {
	return []BackupType{
		BackupTypeManual,
		BackupTypeScheduled,
	}
}

// CheckNameResourceTypes - Resource type used for verification.
type CheckNameResourceTypes string

const (
	CheckNameResourceTypesMicrosoftNetAppNetAppAccounts                              CheckNameResourceTypes = "Microsoft.NetApp/netAppAccounts"
	CheckNameResourceTypesMicrosoftNetAppNetAppAccountsCapacityPools                 CheckNameResourceTypes = "Microsoft.NetApp/netAppAccounts/capacityPools"
	CheckNameResourceTypesMicrosoftNetAppNetAppAccountsCapacityPoolsVolumes          CheckNameResourceTypes = "Microsoft.NetApp/netAppAccounts/capacityPools/volumes"
	CheckNameResourceTypesMicrosoftNetAppNetAppAccountsCapacityPoolsVolumesSnapshots CheckNameResourceTypes = "Microsoft.NetApp/netAppAccounts/capacityPools/volumes/snapshots"
)

// PossibleCheckNameResourceTypesValues returns the possible values for the CheckNameResourceTypes const type.
func PossibleCheckNameResourceTypesValues() []CheckNameResourceTypes {
	return []CheckNameResourceTypes{
		CheckNameResourceTypesMicrosoftNetAppNetAppAccounts,
		CheckNameResourceTypesMicrosoftNetAppNetAppAccountsCapacityPools,
		CheckNameResourceTypesMicrosoftNetAppNetAppAccountsCapacityPoolsVolumes,
		CheckNameResourceTypesMicrosoftNetAppNetAppAccountsCapacityPoolsVolumesSnapshots,
	}
}

// CheckQuotaNameResourceTypes - Resource type used for verification.
type CheckQuotaNameResourceTypes string

const (
	CheckQuotaNameResourceTypesMicrosoftNetAppNetAppAccounts                              CheckQuotaNameResourceTypes = "Microsoft.NetApp/netAppAccounts"
	CheckQuotaNameResourceTypesMicrosoftNetAppNetAppAccountsCapacityPools                 CheckQuotaNameResourceTypes = "Microsoft.NetApp/netAppAccounts/capacityPools"
	CheckQuotaNameResourceTypesMicrosoftNetAppNetAppAccountsCapacityPoolsVolumes          CheckQuotaNameResourceTypes = "Microsoft.NetApp/netAppAccounts/capacityPools/volumes"
	CheckQuotaNameResourceTypesMicrosoftNetAppNetAppAccountsCapacityPoolsVolumesSnapshots CheckQuotaNameResourceTypes = "Microsoft.NetApp/netAppAccounts/capacityPools/volumes/snapshots"
)

// PossibleCheckQuotaNameResourceTypesValues returns the possible values for the CheckQuotaNameResourceTypes const type.
func PossibleCheckQuotaNameResourceTypesValues() []CheckQuotaNameResourceTypes {
	return []CheckQuotaNameResourceTypes{
		CheckQuotaNameResourceTypesMicrosoftNetAppNetAppAccounts,
		CheckQuotaNameResourceTypesMicrosoftNetAppNetAppAccountsCapacityPools,
		CheckQuotaNameResourceTypesMicrosoftNetAppNetAppAccountsCapacityPoolsVolumes,
		CheckQuotaNameResourceTypesMicrosoftNetAppNetAppAccountsCapacityPoolsVolumesSnapshots,
	}
}

// ChownMode - This parameter specifies who is authorized to change the ownership of a file. restricted - Only root user can
// change the ownership of the file. unrestricted - Non-root users can change ownership of
// files that they own.
type ChownMode string

const (
	ChownModeRestricted   ChownMode = "Restricted"
	ChownModeUnrestricted ChownMode = "Unrestricted"
)

// PossibleChownModeValues returns the possible values for the ChownMode const type.
func PossibleChownModeValues() []ChownMode {
	return []ChownMode{
		ChownModeRestricted,
		ChownModeUnrestricted,
	}
}

// CoolAccessRetrievalPolicy - coolAccessRetrievalPolicy determines the data retrieval behavior from the cool tier to standard
// storage based on the read pattern for cool access enabled volumes. The possible values for this field
// are: Default - Data will be pulled from cool tier to standard storage on random reads. This policy is the default. OnRead
// - All client-driven data read is pulled from cool tier to standard storage on
// both sequential and random reads. Never - No client-driven data is pulled from cool tier to standard storage.
type CoolAccessRetrievalPolicy string

const (
	CoolAccessRetrievalPolicyDefault CoolAccessRetrievalPolicy = "Default"
	CoolAccessRetrievalPolicyNever   CoolAccessRetrievalPolicy = "Never"
	CoolAccessRetrievalPolicyOnRead  CoolAccessRetrievalPolicy = "OnRead"
)

// PossibleCoolAccessRetrievalPolicyValues returns the possible values for the CoolAccessRetrievalPolicy const type.
func PossibleCoolAccessRetrievalPolicyValues() []CoolAccessRetrievalPolicy {
	return []CoolAccessRetrievalPolicy{
		CoolAccessRetrievalPolicyDefault,
		CoolAccessRetrievalPolicyNever,
		CoolAccessRetrievalPolicyOnRead,
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

// EnableSubvolumes - Flag indicating whether subvolume operations are enabled on the volume
type EnableSubvolumes string

const (
	// EnableSubvolumesDisabled - subvolumes are not enabled
	EnableSubvolumesDisabled EnableSubvolumes = "Disabled"
	// EnableSubvolumesEnabled - subvolumes are enabled
	EnableSubvolumesEnabled EnableSubvolumes = "Enabled"
)

// PossibleEnableSubvolumesValues returns the possible values for the EnableSubvolumes const type.
func PossibleEnableSubvolumesValues() []EnableSubvolumes {
	return []EnableSubvolumes{
		EnableSubvolumesDisabled,
		EnableSubvolumesEnabled,
	}
}

// EncryptionKeySource - Source of key used to encrypt data in volume. Applicable if NetApp account has encryption.keySource
// = 'Microsoft.KeyVault'. Possible values (case-insensitive) are: 'Microsoft.NetApp,
// Microsoft.KeyVault'
type EncryptionKeySource string

const (
	// EncryptionKeySourceMicrosoftKeyVault - Customer-managed key encryption
	EncryptionKeySourceMicrosoftKeyVault EncryptionKeySource = "Microsoft.KeyVault"
	// EncryptionKeySourceMicrosoftNetApp - Microsoft-managed key encryption
	EncryptionKeySourceMicrosoftNetApp EncryptionKeySource = "Microsoft.NetApp"
)

// PossibleEncryptionKeySourceValues returns the possible values for the EncryptionKeySource const type.
func PossibleEncryptionKeySourceValues() []EncryptionKeySource {
	return []EncryptionKeySource{
		EncryptionKeySourceMicrosoftKeyVault,
		EncryptionKeySourceMicrosoftNetApp,
	}
}

// EncryptionType - Encryption type of the capacity pool, set encryption type for data at rest for this pool and all volumes
// in it. This value can only be set when creating new pool.
type EncryptionType string

const (
	// EncryptionTypeDouble - EncryptionType Double, volumes will use double encryption at rest
	EncryptionTypeDouble EncryptionType = "Double"
	// EncryptionTypeSingle - EncryptionType Single, volumes will use single encryption at rest
	EncryptionTypeSingle EncryptionType = "Single"
)

// PossibleEncryptionTypeValues returns the possible values for the EncryptionType const type.
func PossibleEncryptionTypeValues() []EncryptionType {
	return []EncryptionType{
		EncryptionTypeDouble,
		EncryptionTypeSingle,
	}
}

// EndpointType - Indicates whether the local volume is the source or destination for the Volume Replication
type EndpointType string

const (
	EndpointTypeDst EndpointType = "dst"
	EndpointTypeSrc EndpointType = "src"
)

// PossibleEndpointTypeValues returns the possible values for the EndpointType const type.
func PossibleEndpointTypeValues() []EndpointType {
	return []EndpointType{
		EndpointTypeDst,
		EndpointTypeSrc,
	}
}

// FileAccessLogs - Flag indicating whether file access logs are enabled for the volume, based on active diagnostic settings
// present on the volume.
type FileAccessLogs string

const (
	// FileAccessLogsDisabled - fileAccessLogs are not enabled
	FileAccessLogsDisabled FileAccessLogs = "Disabled"
	// FileAccessLogsEnabled - fileAccessLogs are enabled
	FileAccessLogsEnabled FileAccessLogs = "Enabled"
)

// PossibleFileAccessLogsValues returns the possible values for the FileAccessLogs const type.
func PossibleFileAccessLogsValues() []FileAccessLogs {
	return []FileAccessLogs{
		FileAccessLogsDisabled,
		FileAccessLogsEnabled,
	}
}

// InAvailabilityReasonType - Invalid indicates the name provided does not match Azure App Service naming requirements. AlreadyExists
// indicates that the name is already in use and is therefore unavailable.
type InAvailabilityReasonType string

const (
	InAvailabilityReasonTypeAlreadyExists InAvailabilityReasonType = "AlreadyExists"
	InAvailabilityReasonTypeInvalid       InAvailabilityReasonType = "Invalid"
)

// PossibleInAvailabilityReasonTypeValues returns the possible values for the InAvailabilityReasonType const type.
func PossibleInAvailabilityReasonTypeValues() []InAvailabilityReasonType {
	return []InAvailabilityReasonType{
		InAvailabilityReasonTypeAlreadyExists,
		InAvailabilityReasonTypeInvalid,
	}
}

// KeySource - The encryption keySource (provider). Possible values (case-insensitive): Microsoft.NetApp, Microsoft.KeyVault
type KeySource string

const (
	// KeySourceMicrosoftKeyVault - Customer-managed key encryption
	KeySourceMicrosoftKeyVault KeySource = "Microsoft.KeyVault"
	// KeySourceMicrosoftNetApp - Microsoft-managed key encryption
	KeySourceMicrosoftNetApp KeySource = "Microsoft.NetApp"
)

// PossibleKeySourceValues returns the possible values for the KeySource const type.
func PossibleKeySourceValues() []KeySource {
	return []KeySource{
		KeySourceMicrosoftKeyVault,
		KeySourceMicrosoftNetApp,
	}
}

// KeyVaultStatus - Status of the KeyVault connection.
type KeyVaultStatus string

const (
	// KeyVaultStatusCreated - KeyVault connection created but not in use
	KeyVaultStatusCreated KeyVaultStatus = "Created"
	// KeyVaultStatusDeleted - KeyVault connection Deleted
	KeyVaultStatusDeleted KeyVaultStatus = "Deleted"
	// KeyVaultStatusError - Error with the KeyVault connection
	KeyVaultStatusError KeyVaultStatus = "Error"
	// KeyVaultStatusInUse - KeyVault connection in use by SMB Volume
	KeyVaultStatusInUse KeyVaultStatus = "InUse"
	// KeyVaultStatusUpdating - KeyVault connection Updating
	KeyVaultStatusUpdating KeyVaultStatus = "Updating"
)

// PossibleKeyVaultStatusValues returns the possible values for the KeyVaultStatus const type.
func PossibleKeyVaultStatusValues() []KeyVaultStatus {
	return []KeyVaultStatus{
		KeyVaultStatusCreated,
		KeyVaultStatusDeleted,
		KeyVaultStatusError,
		KeyVaultStatusInUse,
		KeyVaultStatusUpdating,
	}
}

// ManagedServiceIdentityType - Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
type ManagedServiceIdentityType string

const (
	ManagedServiceIdentityTypeNone                       ManagedServiceIdentityType = "None"
	ManagedServiceIdentityTypeSystemAssigned             ManagedServiceIdentityType = "SystemAssigned"
	ManagedServiceIdentityTypeSystemAssignedUserAssigned ManagedServiceIdentityType = "SystemAssigned,UserAssigned"
	ManagedServiceIdentityTypeUserAssigned               ManagedServiceIdentityType = "UserAssigned"
)

// PossibleManagedServiceIdentityTypeValues returns the possible values for the ManagedServiceIdentityType const type.
func PossibleManagedServiceIdentityTypeValues() []ManagedServiceIdentityType {
	return []ManagedServiceIdentityType{
		ManagedServiceIdentityTypeNone,
		ManagedServiceIdentityTypeSystemAssigned,
		ManagedServiceIdentityTypeSystemAssignedUserAssigned,
		ManagedServiceIdentityTypeUserAssigned,
	}
}

type MetricAggregationType string

const (
	MetricAggregationTypeAverage MetricAggregationType = "Average"
)

// PossibleMetricAggregationTypeValues returns the possible values for the MetricAggregationType const type.
func PossibleMetricAggregationTypeValues() []MetricAggregationType {
	return []MetricAggregationType{
		MetricAggregationTypeAverage,
	}
}

// MirrorState - The status of the replication
type MirrorState string

const (
	MirrorStateBroken        MirrorState = "Broken"
	MirrorStateMirrored      MirrorState = "Mirrored"
	MirrorStateUninitialized MirrorState = "Uninitialized"
)

// PossibleMirrorStateValues returns the possible values for the MirrorState const type.
func PossibleMirrorStateValues() []MirrorState {
	return []MirrorState{
		MirrorStateBroken,
		MirrorStateMirrored,
		MirrorStateUninitialized,
	}
}

// NetworkFeatures - Network features available to the volume, or current state of update.
type NetworkFeatures string

const (
	// NetworkFeaturesBasic - Basic network features.
	NetworkFeaturesBasic NetworkFeatures = "Basic"
	// NetworkFeaturesBasicStandard - Updating from Basic to Standard network features.
	NetworkFeaturesBasicStandard NetworkFeatures = "Basic_Standard"
	// NetworkFeaturesStandard - Standard network features.
	NetworkFeaturesStandard NetworkFeatures = "Standard"
	// NetworkFeaturesStandardBasic - Updating from Standard to Basic network features.
	NetworkFeaturesStandardBasic NetworkFeatures = "Standard_Basic"
)

// PossibleNetworkFeaturesValues returns the possible values for the NetworkFeatures const type.
func PossibleNetworkFeaturesValues() []NetworkFeatures {
	return []NetworkFeatures{
		NetworkFeaturesBasic,
		NetworkFeaturesBasicStandard,
		NetworkFeaturesStandard,
		NetworkFeaturesStandardBasic,
	}
}

// NetworkSiblingSetProvisioningState - Gets the status of the NetworkSiblingSet at the time the operation was called.
type NetworkSiblingSetProvisioningState string

const (
	NetworkSiblingSetProvisioningStateCanceled  NetworkSiblingSetProvisioningState = "Canceled"
	NetworkSiblingSetProvisioningStateFailed    NetworkSiblingSetProvisioningState = "Failed"
	NetworkSiblingSetProvisioningStateSucceeded NetworkSiblingSetProvisioningState = "Succeeded"
	NetworkSiblingSetProvisioningStateUpdating  NetworkSiblingSetProvisioningState = "Updating"
)

// PossibleNetworkSiblingSetProvisioningStateValues returns the possible values for the NetworkSiblingSetProvisioningState const type.
func PossibleNetworkSiblingSetProvisioningStateValues() []NetworkSiblingSetProvisioningState {
	return []NetworkSiblingSetProvisioningState{
		NetworkSiblingSetProvisioningStateCanceled,
		NetworkSiblingSetProvisioningStateFailed,
		NetworkSiblingSetProvisioningStateSucceeded,
		NetworkSiblingSetProvisioningStateUpdating,
	}
}

// ProvisioningState - Gets the status of the VolumeQuotaRule at the time the operation was called.
type ProvisioningState string

const (
	ProvisioningStateAccepted  ProvisioningState = "Accepted"
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateMoving    ProvisioningState = "Moving"
	ProvisioningStatePatching  ProvisioningState = "Patching"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateAccepted,
		ProvisioningStateCreating,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateMoving,
		ProvisioningStatePatching,
		ProvisioningStateSucceeded,
	}
}

// QosType - The qos type of the pool
type QosType string

const (
	// QosTypeAuto - qos type Auto
	QosTypeAuto QosType = "Auto"
	// QosTypeManual - qos type Manual
	QosTypeManual QosType = "Manual"
)

// PossibleQosTypeValues returns the possible values for the QosType const type.
func PossibleQosTypeValues() []QosType {
	return []QosType{
		QosTypeAuto,
		QosTypeManual,
	}
}

// RegionStorageToNetworkProximity - Provides storage to network proximity information in the region.
type RegionStorageToNetworkProximity string

const (
	// RegionStorageToNetworkProximityAcrossT2 - Standard AcrossT2 network connectivity.
	RegionStorageToNetworkProximityAcrossT2 RegionStorageToNetworkProximity = "AcrossT2"
	// RegionStorageToNetworkProximityDefault - Basic network connectivity.
	RegionStorageToNetworkProximityDefault RegionStorageToNetworkProximity = "Default"
	// RegionStorageToNetworkProximityT1 - Standard T1 network connectivity.
	RegionStorageToNetworkProximityT1 RegionStorageToNetworkProximity = "T1"
	// RegionStorageToNetworkProximityT1AndAcrossT2 - Standard T1 and AcrossT2 network connectivity.
	RegionStorageToNetworkProximityT1AndAcrossT2 RegionStorageToNetworkProximity = "T1AndAcrossT2"
	// RegionStorageToNetworkProximityT1AndT2 - Standard T1 and T2 network connectivity.
	RegionStorageToNetworkProximityT1AndT2 RegionStorageToNetworkProximity = "T1AndT2"
	// RegionStorageToNetworkProximityT1AndT2AndAcrossT2 - Standard T1, T2 and AcrossT2 network connectivity.
	RegionStorageToNetworkProximityT1AndT2AndAcrossT2 RegionStorageToNetworkProximity = "T1AndT2AndAcrossT2"
	// RegionStorageToNetworkProximityT2 - Standard T2 network connectivity.
	RegionStorageToNetworkProximityT2 RegionStorageToNetworkProximity = "T2"
	// RegionStorageToNetworkProximityT2AndAcrossT2 - Standard T2 and AcrossT2 network connectivity.
	RegionStorageToNetworkProximityT2AndAcrossT2 RegionStorageToNetworkProximity = "T2AndAcrossT2"
)

// PossibleRegionStorageToNetworkProximityValues returns the possible values for the RegionStorageToNetworkProximity const type.
func PossibleRegionStorageToNetworkProximityValues() []RegionStorageToNetworkProximity {
	return []RegionStorageToNetworkProximity{
		RegionStorageToNetworkProximityAcrossT2,
		RegionStorageToNetworkProximityDefault,
		RegionStorageToNetworkProximityT1,
		RegionStorageToNetworkProximityT1AndAcrossT2,
		RegionStorageToNetworkProximityT1AndT2,
		RegionStorageToNetworkProximityT1AndT2AndAcrossT2,
		RegionStorageToNetworkProximityT2,
		RegionStorageToNetworkProximityT2AndAcrossT2,
	}
}

// RelationshipStatus - Status of the mirror relationship
type RelationshipStatus string

const (
	RelationshipStatusFailed       RelationshipStatus = "Failed"
	RelationshipStatusIdle         RelationshipStatus = "Idle"
	RelationshipStatusTransferring RelationshipStatus = "Transferring"
	RelationshipStatusUnknown      RelationshipStatus = "Unknown"
)

// PossibleRelationshipStatusValues returns the possible values for the RelationshipStatus const type.
func PossibleRelationshipStatusValues() []RelationshipStatus {
	return []RelationshipStatus{
		RelationshipStatusFailed,
		RelationshipStatusIdle,
		RelationshipStatusTransferring,
		RelationshipStatusUnknown,
	}
}

// ReplicationSchedule - Schedule
type ReplicationSchedule string

const (
	ReplicationSchedule10Minutely ReplicationSchedule = "_10minutely"
	ReplicationScheduleDaily      ReplicationSchedule = "daily"
	ReplicationScheduleHourly     ReplicationSchedule = "hourly"
)

// PossibleReplicationScheduleValues returns the possible values for the ReplicationSchedule const type.
func PossibleReplicationScheduleValues() []ReplicationSchedule {
	return []ReplicationSchedule{
		ReplicationSchedule10Minutely,
		ReplicationScheduleDaily,
		ReplicationScheduleHourly,
	}
}

// SecurityStyle - The security style of volume, default unix, defaults to ntfs for dual protocol or CIFS protocol
type SecurityStyle string

const (
	SecurityStyleNtfs SecurityStyle = "ntfs"
	SecurityStyleUnix SecurityStyle = "unix"
)

// PossibleSecurityStyleValues returns the possible values for the SecurityStyle const type.
func PossibleSecurityStyleValues() []SecurityStyle {
	return []SecurityStyle{
		SecurityStyleNtfs,
		SecurityStyleUnix,
	}
}

// ServiceLevel - The service level of the file system
type ServiceLevel string

const (
	// ServiceLevelPremium - Premium service level
	ServiceLevelPremium ServiceLevel = "Premium"
	// ServiceLevelStandard - Standard service level
	ServiceLevelStandard ServiceLevel = "Standard"
	// ServiceLevelStandardZRS - Zone redundant storage service level
	ServiceLevelStandardZRS ServiceLevel = "StandardZRS"
	// ServiceLevelUltra - Ultra service level
	ServiceLevelUltra ServiceLevel = "Ultra"
)

// PossibleServiceLevelValues returns the possible values for the ServiceLevel const type.
func PossibleServiceLevelValues() []ServiceLevel {
	return []ServiceLevel{
		ServiceLevelPremium,
		ServiceLevelStandard,
		ServiceLevelStandardZRS,
		ServiceLevelUltra,
	}
}

// SmbAccessBasedEnumeration - Enables access-based enumeration share property for SMB Shares. Only applicable for SMB/DualProtocol
// volume
type SmbAccessBasedEnumeration string

const (
	// SmbAccessBasedEnumerationDisabled - smbAccessBasedEnumeration share setting is disabled
	SmbAccessBasedEnumerationDisabled SmbAccessBasedEnumeration = "Disabled"
	// SmbAccessBasedEnumerationEnabled - smbAccessBasedEnumeration share setting is enabled
	SmbAccessBasedEnumerationEnabled SmbAccessBasedEnumeration = "Enabled"
)

// PossibleSmbAccessBasedEnumerationValues returns the possible values for the SmbAccessBasedEnumeration const type.
func PossibleSmbAccessBasedEnumerationValues() []SmbAccessBasedEnumeration {
	return []SmbAccessBasedEnumeration{
		SmbAccessBasedEnumerationDisabled,
		SmbAccessBasedEnumerationEnabled,
	}
}

// SmbNonBrowsable - Enables non-browsable property for SMB Shares. Only applicable for SMB/DualProtocol volume
type SmbNonBrowsable string

const (
	// SmbNonBrowsableDisabled - smbNonBrowsable share setting is disabled
	SmbNonBrowsableDisabled SmbNonBrowsable = "Disabled"
	// SmbNonBrowsableEnabled - smbNonBrowsable share setting is enabled
	SmbNonBrowsableEnabled SmbNonBrowsable = "Enabled"
)

// PossibleSmbNonBrowsableValues returns the possible values for the SmbNonBrowsable const type.
func PossibleSmbNonBrowsableValues() []SmbNonBrowsable {
	return []SmbNonBrowsable{
		SmbNonBrowsableDisabled,
		SmbNonBrowsableEnabled,
	}
}

// Type - Type of quota
type Type string

const (
	// TypeDefaultGroupQuota - Default group quota
	TypeDefaultGroupQuota Type = "DefaultGroupQuota"
	// TypeDefaultUserQuota - Default user quota
	TypeDefaultUserQuota Type = "DefaultUserQuota"
	// TypeIndividualGroupQuota - Individual group quota
	TypeIndividualGroupQuota Type = "IndividualGroupQuota"
	// TypeIndividualUserQuota - Individual user quota
	TypeIndividualUserQuota Type = "IndividualUserQuota"
)

// PossibleTypeValues returns the possible values for the Type const type.
func PossibleTypeValues() []Type {
	return []Type{
		TypeDefaultGroupQuota,
		TypeDefaultUserQuota,
		TypeIndividualGroupQuota,
		TypeIndividualUserQuota,
	}
}

// VolumeStorageToNetworkProximity - Provides storage to network proximity information for the volume.
type VolumeStorageToNetworkProximity string

const (
	// VolumeStorageToNetworkProximityAcrossT2 - Standard AcrossT2 storage to network connectivity.
	VolumeStorageToNetworkProximityAcrossT2 VolumeStorageToNetworkProximity = "AcrossT2"
	// VolumeStorageToNetworkProximityDefault - Basic storage to network connectivity.
	VolumeStorageToNetworkProximityDefault VolumeStorageToNetworkProximity = "Default"
	// VolumeStorageToNetworkProximityT1 - Standard T1 storage to network connectivity.
	VolumeStorageToNetworkProximityT1 VolumeStorageToNetworkProximity = "T1"
	// VolumeStorageToNetworkProximityT2 - Standard T2 storage to network connectivity.
	VolumeStorageToNetworkProximityT2 VolumeStorageToNetworkProximity = "T2"
)

// PossibleVolumeStorageToNetworkProximityValues returns the possible values for the VolumeStorageToNetworkProximity const type.
func PossibleVolumeStorageToNetworkProximityValues() []VolumeStorageToNetworkProximity {
	return []VolumeStorageToNetworkProximity{
		VolumeStorageToNetworkProximityAcrossT2,
		VolumeStorageToNetworkProximityDefault,
		VolumeStorageToNetworkProximityT1,
		VolumeStorageToNetworkProximityT2,
	}
}
