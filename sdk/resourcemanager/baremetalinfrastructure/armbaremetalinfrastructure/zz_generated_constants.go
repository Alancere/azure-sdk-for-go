//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armbaremetalinfrastructure

const (
	moduleName    = "armbaremetalinfrastructure"
	moduleVersion = "v1.0.0"
)

// AzureBareMetalHardwareTypeNamesEnum - Name of the hardware type (vendor and/or their product name)
type AzureBareMetalHardwareTypeNamesEnum string

const (
	AzureBareMetalHardwareTypeNamesEnumCiscoUCS AzureBareMetalHardwareTypeNamesEnum = "Cisco_UCS"
	AzureBareMetalHardwareTypeNamesEnumHPE      AzureBareMetalHardwareTypeNamesEnum = "HPE"
)

// PossibleAzureBareMetalHardwareTypeNamesEnumValues returns the possible values for the AzureBareMetalHardwareTypeNamesEnum const type.
func PossibleAzureBareMetalHardwareTypeNamesEnumValues() []AzureBareMetalHardwareTypeNamesEnum {
	return []AzureBareMetalHardwareTypeNamesEnum{
		AzureBareMetalHardwareTypeNamesEnumCiscoUCS,
		AzureBareMetalHardwareTypeNamesEnumHPE,
	}
}

// AzureBareMetalInstancePowerStateEnum - Resource power state
type AzureBareMetalInstancePowerStateEnum string

const (
	AzureBareMetalInstancePowerStateEnumRestarting AzureBareMetalInstancePowerStateEnum = "restarting"
	AzureBareMetalInstancePowerStateEnumStarted    AzureBareMetalInstancePowerStateEnum = "started"
	AzureBareMetalInstancePowerStateEnumStarting   AzureBareMetalInstancePowerStateEnum = "starting"
	AzureBareMetalInstancePowerStateEnumStopped    AzureBareMetalInstancePowerStateEnum = "stopped"
	AzureBareMetalInstancePowerStateEnumStopping   AzureBareMetalInstancePowerStateEnum = "stopping"
	AzureBareMetalInstancePowerStateEnumUnknown    AzureBareMetalInstancePowerStateEnum = "unknown"
)

// PossibleAzureBareMetalInstancePowerStateEnumValues returns the possible values for the AzureBareMetalInstancePowerStateEnum const type.
func PossibleAzureBareMetalInstancePowerStateEnumValues() []AzureBareMetalInstancePowerStateEnum {
	return []AzureBareMetalInstancePowerStateEnum{
		AzureBareMetalInstancePowerStateEnumRestarting,
		AzureBareMetalInstancePowerStateEnumStarted,
		AzureBareMetalInstancePowerStateEnumStarting,
		AzureBareMetalInstancePowerStateEnumStopped,
		AzureBareMetalInstancePowerStateEnumStopping,
		AzureBareMetalInstancePowerStateEnumUnknown,
	}
}

// AzureBareMetalInstanceSizeNamesEnum - Specifies the AzureBareMetal instance SKU.
type AzureBareMetalInstanceSizeNamesEnum string

const (
	AzureBareMetalInstanceSizeNamesEnumS112    AzureBareMetalInstanceSizeNamesEnum = "S112"
	AzureBareMetalInstanceSizeNamesEnumS144    AzureBareMetalInstanceSizeNamesEnum = "S144"
	AzureBareMetalInstanceSizeNamesEnumS144M   AzureBareMetalInstanceSizeNamesEnum = "S144m"
	AzureBareMetalInstanceSizeNamesEnumS192    AzureBareMetalInstanceSizeNamesEnum = "S192"
	AzureBareMetalInstanceSizeNamesEnumS192M   AzureBareMetalInstanceSizeNamesEnum = "S192m"
	AzureBareMetalInstanceSizeNamesEnumS192Xm  AzureBareMetalInstanceSizeNamesEnum = "S192xm"
	AzureBareMetalInstanceSizeNamesEnumS224    AzureBareMetalInstanceSizeNamesEnum = "S224"
	AzureBareMetalInstanceSizeNamesEnumS224M   AzureBareMetalInstanceSizeNamesEnum = "S224m"
	AzureBareMetalInstanceSizeNamesEnumS224Om  AzureBareMetalInstanceSizeNamesEnum = "S224om"
	AzureBareMetalInstanceSizeNamesEnumS224Oo  AzureBareMetalInstanceSizeNamesEnum = "S224oo"
	AzureBareMetalInstanceSizeNamesEnumS224Oom AzureBareMetalInstanceSizeNamesEnum = "S224oom"
	AzureBareMetalInstanceSizeNamesEnumS224Ooo AzureBareMetalInstanceSizeNamesEnum = "S224ooo"
	AzureBareMetalInstanceSizeNamesEnumS384    AzureBareMetalInstanceSizeNamesEnum = "S384"
	AzureBareMetalInstanceSizeNamesEnumS384M   AzureBareMetalInstanceSizeNamesEnum = "S384m"
	AzureBareMetalInstanceSizeNamesEnumS384Xm  AzureBareMetalInstanceSizeNamesEnum = "S384xm"
	AzureBareMetalInstanceSizeNamesEnumS384Xxm AzureBareMetalInstanceSizeNamesEnum = "S384xxm"
	AzureBareMetalInstanceSizeNamesEnumS448    AzureBareMetalInstanceSizeNamesEnum = "S448"
	AzureBareMetalInstanceSizeNamesEnumS448M   AzureBareMetalInstanceSizeNamesEnum = "S448m"
	AzureBareMetalInstanceSizeNamesEnumS448Om  AzureBareMetalInstanceSizeNamesEnum = "S448om"
	AzureBareMetalInstanceSizeNamesEnumS448Oo  AzureBareMetalInstanceSizeNamesEnum = "S448oo"
	AzureBareMetalInstanceSizeNamesEnumS448Oom AzureBareMetalInstanceSizeNamesEnum = "S448oom"
	AzureBareMetalInstanceSizeNamesEnumS448Ooo AzureBareMetalInstanceSizeNamesEnum = "S448ooo"
	AzureBareMetalInstanceSizeNamesEnumS576M   AzureBareMetalInstanceSizeNamesEnum = "S576m"
	AzureBareMetalInstanceSizeNamesEnumS576Xm  AzureBareMetalInstanceSizeNamesEnum = "S576xm"
	AzureBareMetalInstanceSizeNamesEnumS672    AzureBareMetalInstanceSizeNamesEnum = "S672"
	AzureBareMetalInstanceSizeNamesEnumS672M   AzureBareMetalInstanceSizeNamesEnum = "S672m"
	AzureBareMetalInstanceSizeNamesEnumS672Om  AzureBareMetalInstanceSizeNamesEnum = "S672om"
	AzureBareMetalInstanceSizeNamesEnumS672Oo  AzureBareMetalInstanceSizeNamesEnum = "S672oo"
	AzureBareMetalInstanceSizeNamesEnumS672Oom AzureBareMetalInstanceSizeNamesEnum = "S672oom"
	AzureBareMetalInstanceSizeNamesEnumS672Ooo AzureBareMetalInstanceSizeNamesEnum = "S672ooo"
	AzureBareMetalInstanceSizeNamesEnumS72     AzureBareMetalInstanceSizeNamesEnum = "S72"
	AzureBareMetalInstanceSizeNamesEnumS72M    AzureBareMetalInstanceSizeNamesEnum = "S72m"
	AzureBareMetalInstanceSizeNamesEnumS768    AzureBareMetalInstanceSizeNamesEnum = "S768"
	AzureBareMetalInstanceSizeNamesEnumS768M   AzureBareMetalInstanceSizeNamesEnum = "S768m"
	AzureBareMetalInstanceSizeNamesEnumS768Xm  AzureBareMetalInstanceSizeNamesEnum = "S768xm"
	AzureBareMetalInstanceSizeNamesEnumS896    AzureBareMetalInstanceSizeNamesEnum = "S896"
	AzureBareMetalInstanceSizeNamesEnumS896M   AzureBareMetalInstanceSizeNamesEnum = "S896m"
	AzureBareMetalInstanceSizeNamesEnumS896Om  AzureBareMetalInstanceSizeNamesEnum = "S896om"
	AzureBareMetalInstanceSizeNamesEnumS896Oo  AzureBareMetalInstanceSizeNamesEnum = "S896oo"
	AzureBareMetalInstanceSizeNamesEnumS896Oom AzureBareMetalInstanceSizeNamesEnum = "S896oom"
	AzureBareMetalInstanceSizeNamesEnumS896Ooo AzureBareMetalInstanceSizeNamesEnum = "S896ooo"
	AzureBareMetalInstanceSizeNamesEnumS96     AzureBareMetalInstanceSizeNamesEnum = "S96"
	AzureBareMetalInstanceSizeNamesEnumS960M   AzureBareMetalInstanceSizeNamesEnum = "S960m"
)

// PossibleAzureBareMetalInstanceSizeNamesEnumValues returns the possible values for the AzureBareMetalInstanceSizeNamesEnum const type.
func PossibleAzureBareMetalInstanceSizeNamesEnumValues() []AzureBareMetalInstanceSizeNamesEnum {
	return []AzureBareMetalInstanceSizeNamesEnum{
		AzureBareMetalInstanceSizeNamesEnumS112,
		AzureBareMetalInstanceSizeNamesEnumS144,
		AzureBareMetalInstanceSizeNamesEnumS144M,
		AzureBareMetalInstanceSizeNamesEnumS192,
		AzureBareMetalInstanceSizeNamesEnumS192M,
		AzureBareMetalInstanceSizeNamesEnumS192Xm,
		AzureBareMetalInstanceSizeNamesEnumS224,
		AzureBareMetalInstanceSizeNamesEnumS224M,
		AzureBareMetalInstanceSizeNamesEnumS224Om,
		AzureBareMetalInstanceSizeNamesEnumS224Oo,
		AzureBareMetalInstanceSizeNamesEnumS224Oom,
		AzureBareMetalInstanceSizeNamesEnumS224Ooo,
		AzureBareMetalInstanceSizeNamesEnumS384,
		AzureBareMetalInstanceSizeNamesEnumS384M,
		AzureBareMetalInstanceSizeNamesEnumS384Xm,
		AzureBareMetalInstanceSizeNamesEnumS384Xxm,
		AzureBareMetalInstanceSizeNamesEnumS448,
		AzureBareMetalInstanceSizeNamesEnumS448M,
		AzureBareMetalInstanceSizeNamesEnumS448Om,
		AzureBareMetalInstanceSizeNamesEnumS448Oo,
		AzureBareMetalInstanceSizeNamesEnumS448Oom,
		AzureBareMetalInstanceSizeNamesEnumS448Ooo,
		AzureBareMetalInstanceSizeNamesEnumS576M,
		AzureBareMetalInstanceSizeNamesEnumS576Xm,
		AzureBareMetalInstanceSizeNamesEnumS672,
		AzureBareMetalInstanceSizeNamesEnumS672M,
		AzureBareMetalInstanceSizeNamesEnumS672Om,
		AzureBareMetalInstanceSizeNamesEnumS672Oo,
		AzureBareMetalInstanceSizeNamesEnumS672Oom,
		AzureBareMetalInstanceSizeNamesEnumS672Ooo,
		AzureBareMetalInstanceSizeNamesEnumS72,
		AzureBareMetalInstanceSizeNamesEnumS72M,
		AzureBareMetalInstanceSizeNamesEnumS768,
		AzureBareMetalInstanceSizeNamesEnumS768M,
		AzureBareMetalInstanceSizeNamesEnumS768Xm,
		AzureBareMetalInstanceSizeNamesEnumS896,
		AzureBareMetalInstanceSizeNamesEnumS896M,
		AzureBareMetalInstanceSizeNamesEnumS896Om,
		AzureBareMetalInstanceSizeNamesEnumS896Oo,
		AzureBareMetalInstanceSizeNamesEnumS896Oom,
		AzureBareMetalInstanceSizeNamesEnumS896Ooo,
		AzureBareMetalInstanceSizeNamesEnumS96,
		AzureBareMetalInstanceSizeNamesEnumS960M,
	}
}

// AzureBareMetalProvisioningStatesEnum - State of provisioning of the AzureBareMetalInstance
type AzureBareMetalProvisioningStatesEnum string

const (
	AzureBareMetalProvisioningStatesEnumAccepted  AzureBareMetalProvisioningStatesEnum = "Accepted"
	AzureBareMetalProvisioningStatesEnumCreating  AzureBareMetalProvisioningStatesEnum = "Creating"
	AzureBareMetalProvisioningStatesEnumDeleting  AzureBareMetalProvisioningStatesEnum = "Deleting"
	AzureBareMetalProvisioningStatesEnumFailed    AzureBareMetalProvisioningStatesEnum = "Failed"
	AzureBareMetalProvisioningStatesEnumMigrating AzureBareMetalProvisioningStatesEnum = "Migrating"
	AzureBareMetalProvisioningStatesEnumSucceeded AzureBareMetalProvisioningStatesEnum = "Succeeded"
	AzureBareMetalProvisioningStatesEnumUpdating  AzureBareMetalProvisioningStatesEnum = "Updating"
)

// PossibleAzureBareMetalProvisioningStatesEnumValues returns the possible values for the AzureBareMetalProvisioningStatesEnum const type.
func PossibleAzureBareMetalProvisioningStatesEnumValues() []AzureBareMetalProvisioningStatesEnum {
	return []AzureBareMetalProvisioningStatesEnum{
		AzureBareMetalProvisioningStatesEnumAccepted,
		AzureBareMetalProvisioningStatesEnumCreating,
		AzureBareMetalProvisioningStatesEnumDeleting,
		AzureBareMetalProvisioningStatesEnumFailed,
		AzureBareMetalProvisioningStatesEnumMigrating,
		AzureBareMetalProvisioningStatesEnumSucceeded,
		AzureBareMetalProvisioningStatesEnumUpdating,
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
