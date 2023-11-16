//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappcomplianceautomation

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcomplianceautomation/armappcomplianceautomation"
	moduleVersion = "v0.2.1"
)

// ActionType - Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
type ActionType string

const (
	ActionTypeInternal ActionType = "Internal"
)

// PossibleActionTypeValues returns the possible values for the ActionType const type.
func PossibleActionTypeValues() []ActionType {
	return []ActionType{
		ActionTypeInternal,
	}
}

// AssessmentSeverity - Indicates the assessment severity.
type AssessmentSeverity string

const (
	AssessmentSeverityHigh   AssessmentSeverity = "High"
	AssessmentSeverityLow    AssessmentSeverity = "Low"
	AssessmentSeverityMedium AssessmentSeverity = "Medium"
)

// PossibleAssessmentSeverityValues returns the possible values for the AssessmentSeverity const type.
func PossibleAssessmentSeverityValues() []AssessmentSeverity {
	return []AssessmentSeverity{
		AssessmentSeverityHigh,
		AssessmentSeverityLow,
		AssessmentSeverityMedium,
	}
}

// CategoryStatus - Indicates the category status.
type CategoryStatus string

const (
	CategoryStatusHealthy   CategoryStatus = "Healthy"
	CategoryStatusUnhealthy CategoryStatus = "Unhealthy"
)

// PossibleCategoryStatusValues returns the possible values for the CategoryStatus const type.
func PossibleCategoryStatusValues() []CategoryStatus {
	return []CategoryStatus{
		CategoryStatusHealthy,
		CategoryStatusUnhealthy,
	}
}

// CategoryType - Indicates the compliance category type.
type CategoryType string

const (
	CategoryTypeFullyAutomated     CategoryType = "FullyAutomated"
	CategoryTypeManual             CategoryType = "Manual"
	CategoryTypePartiallyAutomated CategoryType = "PartiallyAutomated"
)

// PossibleCategoryTypeValues returns the possible values for the CategoryType const type.
func PossibleCategoryTypeValues() []CategoryType {
	return []CategoryType{
		CategoryTypeFullyAutomated,
		CategoryTypeManual,
		CategoryTypePartiallyAutomated,
	}
}

// ComplianceState - The compliance result's status.
type ComplianceState string

const (
	ComplianceStateHealthy   ComplianceState = "Healthy"
	ComplianceStateUnhealthy ComplianceState = "Unhealthy"
)

// PossibleComplianceStateValues returns the possible values for the ComplianceState const type.
func PossibleComplianceStateValues() []ComplianceState {
	return []ComplianceState{
		ComplianceStateHealthy,
		ComplianceStateUnhealthy,
	}
}

// ControlFamilyStatus - Indicates the control family status.
type ControlFamilyStatus string

const (
	ControlFamilyStatusHealthy   ControlFamilyStatus = "Healthy"
	ControlFamilyStatusUnhealthy ControlFamilyStatus = "Unhealthy"
)

// PossibleControlFamilyStatusValues returns the possible values for the ControlFamilyStatus const type.
func PossibleControlFamilyStatusValues() []ControlFamilyStatus {
	return []ControlFamilyStatus{
		ControlFamilyStatusHealthy,
		ControlFamilyStatusUnhealthy,
	}
}

// ControlFamilyType - Indicates the control family type.
type ControlFamilyType string

const (
	ControlFamilyTypeFullyAutomated     ControlFamilyType = "FullyAutomated"
	ControlFamilyTypeManual             ControlFamilyType = "Manual"
	ControlFamilyTypePartiallyAutomated ControlFamilyType = "PartiallyAutomated"
)

// PossibleControlFamilyTypeValues returns the possible values for the ControlFamilyType const type.
func PossibleControlFamilyTypeValues() []ControlFamilyType {
	return []ControlFamilyType{
		ControlFamilyTypeFullyAutomated,
		ControlFamilyTypeManual,
		ControlFamilyTypePartiallyAutomated,
	}
}

// ControlStatus - Indicates the control status.
type ControlStatus string

const (
	ControlStatusFailed        ControlStatus = "Failed"
	ControlStatusNotApplicable ControlStatus = "NotApplicable"
	ControlStatusPassed        ControlStatus = "Passed"
)

// PossibleControlStatusValues returns the possible values for the ControlStatus const type.
func PossibleControlStatusValues() []ControlStatus {
	return []ControlStatus{
		ControlStatusFailed,
		ControlStatusNotApplicable,
		ControlStatusPassed,
	}
}

// ControlType - Indicates the control type.
type ControlType string

const (
	ControlTypeFullyAutomated     ControlType = "FullyAutomated"
	ControlTypeManual             ControlType = "Manual"
	ControlTypePartiallyAutomated ControlType = "PartiallyAutomated"
)

// PossibleControlTypeValues returns the possible values for the ControlType const type.
func PossibleControlTypeValues() []ControlType {
	return []ControlType{
		ControlTypeFullyAutomated,
		ControlTypeManual,
		ControlTypePartiallyAutomated,
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

// DownloadType - Indicates the download type.
type DownloadType string

const (
	DownloadTypeComplianceDetailedPDFReport DownloadType = "ComplianceDetailedPdfReport"
	DownloadTypeCompliancePDFReport         DownloadType = "CompliancePdfReport"
	DownloadTypeComplianceReport            DownloadType = "ComplianceReport"
	DownloadTypeResourceList                DownloadType = "ResourceList"
)

// PossibleDownloadTypeValues returns the possible values for the DownloadType const type.
func PossibleDownloadTypeValues() []DownloadType {
	return []DownloadType{
		DownloadTypeComplianceDetailedPDFReport,
		DownloadTypeCompliancePDFReport,
		DownloadTypeComplianceReport,
		DownloadTypeResourceList,
	}
}

// IsPass - Indicates whether all the resource(s) are compliant.
type IsPass string

const (
	IsPassFalse IsPass = "False"
	IsPassTrue  IsPass = "True"
)

// PossibleIsPassValues returns the possible values for the IsPass const type.
func PossibleIsPassValues() []IsPass {
	return []IsPass{
		IsPassFalse,
		IsPassTrue,
	}
}

// Origin - The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
// value is "user,system"
type Origin string

const (
	OriginSystem     Origin = "system"
	OriginUser       Origin = "user"
	OriginUserSystem Origin = "user,system"
)

// PossibleOriginValues returns the possible values for the Origin const type.
func PossibleOriginValues() []Origin {
	return []Origin{
		OriginSystem,
		OriginUser,
		OriginUserSystem,
	}
}

// ProvisioningState - Resource provisioning states.
type ProvisioningState string

const (
	ProvisioningStateCanceled  ProvisioningState = "Canceled"
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateUpdating  ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateCanceled,
		ProvisioningStateCreating,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateSucceeded,
		ProvisioningStateUpdating,
	}
}

// ReportStatus - Report status.
type ReportStatus string

const (
	ReportStatusActive   ReportStatus = "Active"
	ReportStatusDisabled ReportStatus = "Disabled"
	ReportStatusFailed   ReportStatus = "Failed"
)

// PossibleReportStatusValues returns the possible values for the ReportStatus const type.
func PossibleReportStatusValues() []ReportStatus {
	return []ReportStatus{
		ReportStatusActive,
		ReportStatusDisabled,
		ReportStatusFailed,
	}
}

// ResourceStatus - Indicates the resource status.
type ResourceStatus string

const (
	ResourceStatusHealthy       ResourceStatus = "Healthy"
	ResourceStatusNotApplicable ResourceStatus = "NotApplicable"
	ResourceStatusUnhealthy     ResourceStatus = "Unhealthy"
)

// PossibleResourceStatusValues returns the possible values for the ResourceStatus const type.
func PossibleResourceStatusValues() []ResourceStatus {
	return []ResourceStatus{
		ResourceStatusHealthy,
		ResourceStatusNotApplicable,
		ResourceStatusUnhealthy,
	}
}
