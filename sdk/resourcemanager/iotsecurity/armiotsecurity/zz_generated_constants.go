//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armiotsecurity

const (
	moduleName    = "armiotsecurity"
	moduleVersion = "v0.2.1"
)

// AlertIntent - Kill chain related intent behind the alert. Could contain multiple enum values (separated by commas)
type AlertIntent string

const (
	AlertIntentCollection          AlertIntent = "Collection"
	AlertIntentCommandAndControl   AlertIntent = "CommandAndControl"
	AlertIntentCredentialAccess    AlertIntent = "CredentialAccess"
	AlertIntentDefenseEvasion      AlertIntent = "DefenseEvasion"
	AlertIntentDiscovery           AlertIntent = "Discovery"
	AlertIntentExecution           AlertIntent = "Execution"
	AlertIntentExfiltration        AlertIntent = "Exfiltration"
	AlertIntentExploitation        AlertIntent = "Exploitation"
	AlertIntentImpact              AlertIntent = "Impact"
	AlertIntentInitialAccess       AlertIntent = "InitialAccess"
	AlertIntentLateralMovement     AlertIntent = "LateralMovement"
	AlertIntentPersistence         AlertIntent = "Persistence"
	AlertIntentPreAttack           AlertIntent = "PreAttack"
	AlertIntentPrivilegeEscalation AlertIntent = "PrivilegeEscalation"
	AlertIntentProbing             AlertIntent = "Probing"
	AlertIntentUnknown             AlertIntent = "Unknown"
)

// PossibleAlertIntentValues returns the possible values for the AlertIntent const type.
func PossibleAlertIntentValues() []AlertIntent {
	return []AlertIntent{
		AlertIntentCollection,
		AlertIntentCommandAndControl,
		AlertIntentCredentialAccess,
		AlertIntentDefenseEvasion,
		AlertIntentDiscovery,
		AlertIntentExecution,
		AlertIntentExfiltration,
		AlertIntentExploitation,
		AlertIntentImpact,
		AlertIntentInitialAccess,
		AlertIntentLateralMovement,
		AlertIntentPersistence,
		AlertIntentPreAttack,
		AlertIntentPrivilegeEscalation,
		AlertIntentProbing,
		AlertIntentUnknown,
	}
}

// ToPtr returns a *AlertIntent pointing to the current value.
func (c AlertIntent) ToPtr() *AlertIntent {
	return &c
}

// AlertSeverity - Alert Severity
type AlertSeverity string

const (
	AlertSeverityHigh          AlertSeverity = "High"
	AlertSeverityInformational AlertSeverity = "Informational"
	AlertSeverityLow           AlertSeverity = "Low"
	AlertSeverityMedium        AlertSeverity = "Medium"
)

// PossibleAlertSeverityValues returns the possible values for the AlertSeverity const type.
func PossibleAlertSeverityValues() []AlertSeverity {
	return []AlertSeverity{
		AlertSeverityHigh,
		AlertSeverityInformational,
		AlertSeverityLow,
		AlertSeverityMedium,
	}
}

// ToPtr returns a *AlertSeverity pointing to the current value.
func (c AlertSeverity) ToPtr() *AlertSeverity {
	return &c
}

// AlertStatus - Alert Status
type AlertStatus string

const (
	AlertStatusClosed     AlertStatus = "Closed"
	AlertStatusInProgress AlertStatus = "InProgress"
	AlertStatusNew        AlertStatus = "New"
)

// PossibleAlertStatusValues returns the possible values for the AlertStatus const type.
func PossibleAlertStatusValues() []AlertStatus {
	return []AlertStatus{
		AlertStatusClosed,
		AlertStatusInProgress,
		AlertStatusNew,
	}
}

// ToPtr returns a *AlertStatus pointing to the current value.
func (c AlertStatus) ToPtr() *AlertStatus {
	return &c
}

// AuthorizedState - Authorized state of the device.
type AuthorizedState string

const (
	AuthorizedStateAuthorized   AuthorizedState = "Authorized"
	AuthorizedStateUnauthorized AuthorizedState = "Unauthorized"
)

// PossibleAuthorizedStateValues returns the possible values for the AuthorizedState const type.
func PossibleAuthorizedStateValues() []AuthorizedState {
	return []AuthorizedState{
		AuthorizedStateAuthorized,
		AuthorizedStateUnauthorized,
	}
}

// ToPtr returns a *AuthorizedState pointing to the current value.
func (c AuthorizedState) ToPtr() *AuthorizedState {
	return &c
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

// ToPtr returns a *CreatedByType pointing to the current value.
func (c CreatedByType) ToPtr() *CreatedByType {
	return &c
}

// Criticality - Device criticality.
type Criticality string

const (
	CriticalityHigh   Criticality = "High"
	CriticalityLow    Criticality = "Low"
	CriticalityNormal Criticality = "Normal"
)

// PossibleCriticalityValues returns the possible values for the Criticality const type.
func PossibleCriticalityValues() []Criticality {
	return []Criticality{
		CriticalityHigh,
		CriticalityLow,
		CriticalityNormal,
	}
}

// ToPtr returns a *Criticality pointing to the current value.
func (c Criticality) ToPtr() *Criticality {
	return &c
}

// DeviceDataSource - Device data source
type DeviceDataSource string

const (
	DeviceDataSourceCorpSensor DeviceDataSource = "CorpSensor"
	DeviceDataSourceMde        DeviceDataSource = "Mde"
	DeviceDataSourceMicroAgent DeviceDataSource = "MicroAgent"
	DeviceDataSourceOtSensor   DeviceDataSource = "OtSensor"
	DeviceDataSourceOther      DeviceDataSource = "Other"
)

// PossibleDeviceDataSourceValues returns the possible values for the DeviceDataSource const type.
func PossibleDeviceDataSourceValues() []DeviceDataSource {
	return []DeviceDataSource{
		DeviceDataSourceCorpSensor,
		DeviceDataSourceMde,
		DeviceDataSourceMicroAgent,
		DeviceDataSourceOtSensor,
		DeviceDataSourceOther,
	}
}

// ToPtr returns a *DeviceDataSource pointing to the current value.
func (c DeviceDataSource) ToPtr() *DeviceDataSource {
	return &c
}

// DeviceStatus - Device status.
type DeviceStatus string

const (
	DeviceStatusActive   DeviceStatus = "Active"
	DeviceStatusDeleted  DeviceStatus = "Deleted"
	DeviceStatusInactive DeviceStatus = "Inactive"
	DeviceStatusRemoved  DeviceStatus = "Removed"
)

// PossibleDeviceStatusValues returns the possible values for the DeviceStatus const type.
func PossibleDeviceStatusValues() []DeviceStatus {
	return []DeviceStatus{
		DeviceStatusActive,
		DeviceStatusDeleted,
		DeviceStatusInactive,
		DeviceStatusRemoved,
	}
}

// ToPtr returns a *DeviceStatus pointing to the current value.
func (c DeviceStatus) ToPtr() *DeviceStatus {
	return &c
}

// MacCertainty - Indicates whether the association of the mac to the ip address is certain or a guess.
type MacCertainty string

const (
	MacCertaintyCertain MacCertainty = "Certain"
	MacCertaintyGuess   MacCertainty = "Guess"
)

// PossibleMacCertaintyValues returns the possible values for the MacCertainty const type.
func PossibleMacCertaintyValues() []MacCertainty {
	return []MacCertainty{
		MacCertaintyCertain,
		MacCertaintyGuess,
	}
}

// ToPtr returns a *MacCertainty pointing to the current value.
func (c MacCertainty) ToPtr() *MacCertainty {
	return &c
}

// MdeIntegration - Integration status
type MdeIntegration string

const (
	MdeIntegrationDisabled MdeIntegration = "Disabled"
	MdeIntegrationEnabled  MdeIntegration = "Enabled"
)

// PossibleMdeIntegrationValues returns the possible values for the MdeIntegration const type.
func PossibleMdeIntegrationValues() []MdeIntegration {
	return []MdeIntegration{
		MdeIntegrationDisabled,
		MdeIntegrationEnabled,
	}
}

// ToPtr returns a *MdeIntegration pointing to the current value.
func (c MdeIntegration) ToPtr() *MdeIntegration {
	return &c
}

// OnboardingKind - The kind of onboarding for the subscription
type OnboardingKind string

const (
	OnboardingKindDefault         OnboardingKind = "Default"
	OnboardingKindEvaluation      OnboardingKind = "Evaluation"
	OnboardingKindMigratedToAzure OnboardingKind = "MigratedToAzure"
	OnboardingKindPurchased       OnboardingKind = "Purchased"
)

// PossibleOnboardingKindValues returns the possible values for the OnboardingKind const type.
func PossibleOnboardingKindValues() []OnboardingKind {
	return []OnboardingKind{
		OnboardingKindDefault,
		OnboardingKindEvaluation,
		OnboardingKindMigratedToAzure,
		OnboardingKindPurchased,
	}
}

// ToPtr returns a *OnboardingKind pointing to the current value.
func (c OnboardingKind) ToPtr() *OnboardingKind {
	return &c
}

// OnboardingStatus - Device onboarding status.
type OnboardingStatus string

const (
	OnboardingStatusInsufficientInfo OnboardingStatus = "InsufficientInfo"
	OnboardingStatusNotOnboarded     OnboardingStatus = "NotOnboarded"
	OnboardingStatusNotSupported     OnboardingStatus = "NotSupported"
	OnboardingStatusOnboarded        OnboardingStatus = "Onboarded"
)

// PossibleOnboardingStatusValues returns the possible values for the OnboardingStatus const type.
func PossibleOnboardingStatusValues() []OnboardingStatus {
	return []OnboardingStatus{
		OnboardingStatusInsufficientInfo,
		OnboardingStatusNotOnboarded,
		OnboardingStatusNotSupported,
		OnboardingStatusOnboarded,
	}
}

// ToPtr returns a *OnboardingStatus pointing to the current value.
func (c OnboardingStatus) ToPtr() *OnboardingStatus {
	return &c
}

// ProgrammingState - Indicates whether this device is programming
type ProgrammingState string

const (
	ProgrammingStateNotProgrammingDevice ProgrammingState = "NotProgrammingDevice"
	ProgrammingStateProgrammingDevice    ProgrammingState = "ProgrammingDevice"
)

// PossibleProgrammingStateValues returns the possible values for the ProgrammingState const type.
func PossibleProgrammingStateValues() []ProgrammingState {
	return []ProgrammingState{
		ProgrammingStateNotProgrammingDevice,
		ProgrammingStateProgrammingDevice,
	}
}

// ToPtr returns a *ProgrammingState pointing to the current value.
func (c ProgrammingState) ToPtr() *ProgrammingState {
	return &c
}

// PurdueLevel - Purdue level of the device.
type PurdueLevel string

const (
	PurdueLevelEnterprise     PurdueLevel = "Enterprise"
	PurdueLevelProcessControl PurdueLevel = "ProcessControl"
	PurdueLevelSupervisory    PurdueLevel = "Supervisory"
)

// PossiblePurdueLevelValues returns the possible values for the PurdueLevel const type.
func PossiblePurdueLevelValues() []PurdueLevel {
	return []PurdueLevel{
		PurdueLevelEnterprise,
		PurdueLevelProcessControl,
		PurdueLevelSupervisory,
	}
}

// ToPtr returns a *PurdueLevel pointing to the current value.
func (c PurdueLevel) ToPtr() *PurdueLevel {
	return &c
}

// RecommendationSeverity - The severity of the recommendation
type RecommendationSeverity string

const (
	RecommendationSeverityHealthy       RecommendationSeverity = "Healthy"
	RecommendationSeverityHigh          RecommendationSeverity = "High"
	RecommendationSeverityLow           RecommendationSeverity = "Low"
	RecommendationSeverityMedium        RecommendationSeverity = "Medium"
	RecommendationSeverityNotApplicable RecommendationSeverity = "NotApplicable"
	RecommendationSeverityOffByPolicy   RecommendationSeverity = "OffByPolicy"
	RecommendationSeverityUnknown       RecommendationSeverity = "Unknown"
)

// PossibleRecommendationSeverityValues returns the possible values for the RecommendationSeverity const type.
func PossibleRecommendationSeverityValues() []RecommendationSeverity {
	return []RecommendationSeverity{
		RecommendationSeverityHealthy,
		RecommendationSeverityHigh,
		RecommendationSeverityLow,
		RecommendationSeverityMedium,
		RecommendationSeverityNotApplicable,
		RecommendationSeverityOffByPolicy,
		RecommendationSeverityUnknown,
	}
}

// ToPtr returns a *RecommendationSeverity pointing to the current value.
func (c RecommendationSeverity) ToPtr() *RecommendationSeverity {
	return &c
}

// SensorStatus - Status of the IoT sensor
type SensorStatus string

const (
	SensorStatusDisconnected SensorStatus = "Disconnected"
	SensorStatusOk           SensorStatus = "Ok"
	SensorStatusUnavailable  SensorStatus = "Unavailable"
)

// PossibleSensorStatusValues returns the possible values for the SensorStatus const type.
func PossibleSensorStatusValues() []SensorStatus {
	return []SensorStatus{
		SensorStatusDisconnected,
		SensorStatusOk,
		SensorStatusUnavailable,
	}
}

// ToPtr returns a *SensorStatus pointing to the current value.
func (c SensorStatus) ToPtr() *SensorStatus {
	return &c
}

// SensorType - Type of sensor
type SensorType string

const (
	SensorTypeEnterprise SensorType = "Enterprise"
	SensorTypeOt         SensorType = "Ot"
)

// PossibleSensorTypeValues returns the possible values for the SensorType const type.
func PossibleSensorTypeValues() []SensorType {
	return []SensorType{
		SensorTypeEnterprise,
		SensorTypeOt,
	}
}

// ToPtr returns a *SensorType pointing to the current value.
func (c SensorType) ToPtr() *SensorType {
	return &c
}

// SeverityScore - The severity score of the vulnerability
type SeverityScore string

const (
	SeverityScoreCritical SeverityScore = "Critical"
	SeverityScoreHigh     SeverityScore = "High"
	SeverityScoreLow      SeverityScore = "Low"
	SeverityScoreMedium   SeverityScore = "Medium"
	SeverityScoreNone     SeverityScore = "None"
)

// PossibleSeverityScoreValues returns the possible values for the SeverityScore const type.
func PossibleSeverityScoreValues() []SeverityScore {
	return []SeverityScore{
		SeverityScoreCritical,
		SeverityScoreHigh,
		SeverityScoreLow,
		SeverityScoreMedium,
		SeverityScoreNone,
	}
}

// ToPtr returns a *SeverityScore pointing to the current value.
func (c SeverityScore) ToPtr() *SeverityScore {
	return &c
}

// SlotType - Slot type.
type SlotType string

const (
	SlotTypeAnalogIO    SlotType = "AnalogIO"
	SlotTypeCPU         SlotType = "Cpu"
	SlotTypeCommAdapter SlotType = "CommAdapter"
	SlotTypeDigitalIO   SlotType = "DigitalIO"
	SlotTypeGeneric     SlotType = "Generic"
	SlotTypeHmi         SlotType = "Hmi"
	SlotTypeSupply      SlotType = "Supply"
)

// PossibleSlotTypeValues returns the possible values for the SlotType const type.
func PossibleSlotTypeValues() []SlotType {
	return []SlotType{
		SlotTypeAnalogIO,
		SlotTypeCPU,
		SlotTypeCommAdapter,
		SlotTypeDigitalIO,
		SlotTypeGeneric,
		SlotTypeHmi,
		SlotTypeSupply,
	}
}

// ToPtr returns a *SlotType pointing to the current value.
func (c SlotType) ToPtr() *SlotType {
	return &c
}

// TiStatus - TI Status of the IoT sensor
type TiStatus string

const (
	TiStatusFailed          TiStatus = "Failed"
	TiStatusInProgress      TiStatus = "InProgress"
	TiStatusOk              TiStatus = "Ok"
	TiStatusUpdateAvailable TiStatus = "UpdateAvailable"
)

// PossibleTiStatusValues returns the possible values for the TiStatus const type.
func PossibleTiStatusValues() []TiStatus {
	return []TiStatus{
		TiStatusFailed,
		TiStatusInProgress,
		TiStatusOk,
		TiStatusUpdateAvailable,
	}
}

// ToPtr returns a *TiStatus pointing to the current value.
func (c TiStatus) ToPtr() *TiStatus {
	return &c
}

// VersionKind - Kind of the version
type VersionKind string

const (
	VersionKindLatest   VersionKind = "Latest"
	VersionKindPreview  VersionKind = "Preview"
	VersionKindPrevious VersionKind = "Previous"
)

// PossibleVersionKindValues returns the possible values for the VersionKind const type.
func PossibleVersionKindValues() []VersionKind {
	return []VersionKind{
		VersionKindLatest,
		VersionKindPreview,
		VersionKindPrevious,
	}
}

// ToPtr returns a *VersionKind pointing to the current value.
func (c VersionKind) ToPtr() *VersionKind {
	return &c
}
