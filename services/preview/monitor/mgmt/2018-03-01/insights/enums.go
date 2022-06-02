package insights

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// AggregationType enumerates the values for aggregation type.
type AggregationType string

const (
	// Average ...
	Average AggregationType = "Average"
	// Count ...
	Count AggregationType = "Count"
	// Maximum ...
	Maximum AggregationType = "Maximum"
	// Minimum ...
	Minimum AggregationType = "Minimum"
	// None ...
	None AggregationType = "None"
	// Total ...
	Total AggregationType = "Total"
)

// PossibleAggregationTypeValues returns an array of possible values for the AggregationType const type.
func PossibleAggregationTypeValues() []AggregationType {
	return []AggregationType{Average, Count, Maximum, Minimum, None, Total}
}

// AlertSeverity enumerates the values for alert severity.
type AlertSeverity string

const (
	// Four ...
	Four AlertSeverity = "4"
	// One ...
	One AlertSeverity = "1"
	// Three ...
	Three AlertSeverity = "3"
	// Two ...
	Two AlertSeverity = "2"
	// Zero ...
	Zero AlertSeverity = "0"
)

// PossibleAlertSeverityValues returns an array of possible values for the AlertSeverity const type.
func PossibleAlertSeverityValues() []AlertSeverity {
	return []AlertSeverity{Four, One, Three, Two, Zero}
}

// CategoryType enumerates the values for category type.
type CategoryType string

const (
	// Logs ...
	Logs CategoryType = "Logs"
	// Metrics ...
	Metrics CategoryType = "Metrics"
)

// PossibleCategoryTypeValues returns an array of possible values for the CategoryType const type.
func PossibleCategoryTypeValues() []CategoryType {
	return []CategoryType{Logs, Metrics}
}

// ComparisonOperationType enumerates the values for comparison operation type.
type ComparisonOperationType string

const (
	// Equals ...
	Equals ComparisonOperationType = "Equals"
	// GreaterThan ...
	GreaterThan ComparisonOperationType = "GreaterThan"
	// GreaterThanOrEqual ...
	GreaterThanOrEqual ComparisonOperationType = "GreaterThanOrEqual"
	// LessThan ...
	LessThan ComparisonOperationType = "LessThan"
	// LessThanOrEqual ...
	LessThanOrEqual ComparisonOperationType = "LessThanOrEqual"
	// NotEquals ...
	NotEquals ComparisonOperationType = "NotEquals"
)

// PossibleComparisonOperationTypeValues returns an array of possible values for the ComparisonOperationType const type.
func PossibleComparisonOperationTypeValues() []ComparisonOperationType {
	return []ComparisonOperationType{Equals, GreaterThan, GreaterThanOrEqual, LessThan, LessThanOrEqual, NotEquals}
}

// ConditionalOperator enumerates the values for conditional operator.
type ConditionalOperator string

const (
	// ConditionalOperatorEqual ...
	ConditionalOperatorEqual ConditionalOperator = "Equal"
	// ConditionalOperatorGreaterThan ...
	ConditionalOperatorGreaterThan ConditionalOperator = "GreaterThan"
	// ConditionalOperatorLessThan ...
	ConditionalOperatorLessThan ConditionalOperator = "LessThan"
)

// PossibleConditionalOperatorValues returns an array of possible values for the ConditionalOperator const type.
func PossibleConditionalOperatorValues() []ConditionalOperator {
	return []ConditionalOperator{ConditionalOperatorEqual, ConditionalOperatorGreaterThan, ConditionalOperatorLessThan}
}

// ConditionOperator enumerates the values for condition operator.
type ConditionOperator string

const (
	// ConditionOperatorGreaterThan ...
	ConditionOperatorGreaterThan ConditionOperator = "GreaterThan"
	// ConditionOperatorGreaterThanOrEqual ...
	ConditionOperatorGreaterThanOrEqual ConditionOperator = "GreaterThanOrEqual"
	// ConditionOperatorLessThan ...
	ConditionOperatorLessThan ConditionOperator = "LessThan"
	// ConditionOperatorLessThanOrEqual ...
	ConditionOperatorLessThanOrEqual ConditionOperator = "LessThanOrEqual"
)

// PossibleConditionOperatorValues returns an array of possible values for the ConditionOperator const type.
func PossibleConditionOperatorValues() []ConditionOperator {
	return []ConditionOperator{ConditionOperatorGreaterThan, ConditionOperatorGreaterThanOrEqual, ConditionOperatorLessThan, ConditionOperatorLessThanOrEqual}
}

// CriterionType enumerates the values for criterion type.
type CriterionType string

const (
	// CriterionTypeDynamicThresholdCriterion ...
	CriterionTypeDynamicThresholdCriterion CriterionType = "DynamicThresholdCriterion"
	// CriterionTypeMultiMetricCriteria ...
	CriterionTypeMultiMetricCriteria CriterionType = "MultiMetricCriteria"
	// CriterionTypeStaticThresholdCriterion ...
	CriterionTypeStaticThresholdCriterion CriterionType = "StaticThresholdCriterion"
)

// PossibleCriterionTypeValues returns an array of possible values for the CriterionType const type.
func PossibleCriterionTypeValues() []CriterionType {
	return []CriterionType{CriterionTypeDynamicThresholdCriterion, CriterionTypeMultiMetricCriteria, CriterionTypeStaticThresholdCriterion}
}

// DynamicThresholdOperator enumerates the values for dynamic threshold operator.
type DynamicThresholdOperator string

const (
	// DynamicThresholdOperatorGreaterOrLessThan ...
	DynamicThresholdOperatorGreaterOrLessThan DynamicThresholdOperator = "GreaterOrLessThan"
	// DynamicThresholdOperatorGreaterThan ...
	DynamicThresholdOperatorGreaterThan DynamicThresholdOperator = "GreaterThan"
	// DynamicThresholdOperatorLessThan ...
	DynamicThresholdOperatorLessThan DynamicThresholdOperator = "LessThan"
)

// PossibleDynamicThresholdOperatorValues returns an array of possible values for the DynamicThresholdOperator const type.
func PossibleDynamicThresholdOperatorValues() []DynamicThresholdOperator {
	return []DynamicThresholdOperator{DynamicThresholdOperatorGreaterOrLessThan, DynamicThresholdOperatorGreaterThan, DynamicThresholdOperatorLessThan}
}

// DynamicThresholdSensitivity enumerates the values for dynamic threshold sensitivity.
type DynamicThresholdSensitivity string

const (
	// High ...
	High DynamicThresholdSensitivity = "High"
	// Low ...
	Low DynamicThresholdSensitivity = "Low"
	// Medium ...
	Medium DynamicThresholdSensitivity = "Medium"
)

// PossibleDynamicThresholdSensitivityValues returns an array of possible values for the DynamicThresholdSensitivity const type.
func PossibleDynamicThresholdSensitivityValues() []DynamicThresholdSensitivity {
	return []DynamicThresholdSensitivity{High, Low, Medium}
}

// Enabled enumerates the values for enabled.
type Enabled string

const (
	// False ...
	False Enabled = "false"
	// True ...
	True Enabled = "true"
)

// PossibleEnabledValues returns an array of possible values for the Enabled const type.
func PossibleEnabledValues() []Enabled {
	return []Enabled{False, True}
}

// EventLevel enumerates the values for event level.
type EventLevel string

const (
	// Critical ...
	Critical EventLevel = "Critical"
	// Error ...
	Error EventLevel = "Error"
	// Informational ...
	Informational EventLevel = "Informational"
	// Verbose ...
	Verbose EventLevel = "Verbose"
	// Warning ...
	Warning EventLevel = "Warning"
)

// PossibleEventLevelValues returns an array of possible values for the EventLevel const type.
func PossibleEventLevelValues() []EventLevel {
	return []EventLevel{Critical, Error, Informational, Verbose, Warning}
}

// MetricStatisticType enumerates the values for metric statistic type.
type MetricStatisticType string

const (
	// MetricStatisticTypeAverage ...
	MetricStatisticTypeAverage MetricStatisticType = "Average"
	// MetricStatisticTypeMax ...
	MetricStatisticTypeMax MetricStatisticType = "Max"
	// MetricStatisticTypeMin ...
	MetricStatisticTypeMin MetricStatisticType = "Min"
	// MetricStatisticTypeSum ...
	MetricStatisticTypeSum MetricStatisticType = "Sum"
)

// PossibleMetricStatisticTypeValues returns an array of possible values for the MetricStatisticType const type.
func PossibleMetricStatisticTypeValues() []MetricStatisticType {
	return []MetricStatisticType{MetricStatisticTypeAverage, MetricStatisticTypeMax, MetricStatisticTypeMin, MetricStatisticTypeSum}
}

// MetricTriggerType enumerates the values for metric trigger type.
type MetricTriggerType string

const (
	// MetricTriggerTypeConsecutive ...
	MetricTriggerTypeConsecutive MetricTriggerType = "Consecutive"
	// MetricTriggerTypeTotal ...
	MetricTriggerTypeTotal MetricTriggerType = "Total"
)

// PossibleMetricTriggerTypeValues returns an array of possible values for the MetricTriggerType const type.
func PossibleMetricTriggerTypeValues() []MetricTriggerType {
	return []MetricTriggerType{MetricTriggerTypeConsecutive, MetricTriggerTypeTotal}
}

// OdataType enumerates the values for odata type.
type OdataType string

const (
	// OdataTypeMicrosoftAzureManagementInsightsModelsRuleManagementEventDataSource ...
	OdataTypeMicrosoftAzureManagementInsightsModelsRuleManagementEventDataSource OdataType = "Microsoft.Azure.Management.Insights.Models.RuleManagementEventDataSource"
	// OdataTypeMicrosoftAzureManagementInsightsModelsRuleMetricDataSource ...
	OdataTypeMicrosoftAzureManagementInsightsModelsRuleMetricDataSource OdataType = "Microsoft.Azure.Management.Insights.Models.RuleMetricDataSource"
	// OdataTypeRuleDataSource ...
	OdataTypeRuleDataSource OdataType = "RuleDataSource"
)

// PossibleOdataTypeValues returns an array of possible values for the OdataType const type.
func PossibleOdataTypeValues() []OdataType {
	return []OdataType{OdataTypeMicrosoftAzureManagementInsightsModelsRuleManagementEventDataSource, OdataTypeMicrosoftAzureManagementInsightsModelsRuleMetricDataSource, OdataTypeRuleDataSource}
}

// OdataTypeBasicAction enumerates the values for odata type basic action.
type OdataTypeBasicAction string

const (
	// OdataTypeAction ...
	OdataTypeAction OdataTypeBasicAction = "Action"
	// OdataTypeMicrosoftWindowsAzureManagementMonitoringAlertsModelsMicrosoftAppInsightsNexusDataContractsResourcesScheduledQueryRulesAlertingAction ...
	OdataTypeMicrosoftWindowsAzureManagementMonitoringAlertsModelsMicrosoftAppInsightsNexusDataContractsResourcesScheduledQueryRulesAlertingAction OdataTypeBasicAction = "Microsoft.WindowsAzure.Management.Monitoring.Alerts.Models.Microsoft.AppInsights.Nexus.DataContracts.Resources.ScheduledQueryRules.AlertingAction"
	// OdataTypeMicrosoftWindowsAzureManagementMonitoringAlertsModelsMicrosoftAppInsightsNexusDataContractsResourcesScheduledQueryRulesLogToMetricAction ...
	OdataTypeMicrosoftWindowsAzureManagementMonitoringAlertsModelsMicrosoftAppInsightsNexusDataContractsResourcesScheduledQueryRulesLogToMetricAction OdataTypeBasicAction = "Microsoft.WindowsAzure.Management.Monitoring.Alerts.Models.Microsoft.AppInsights.Nexus.DataContracts.Resources.ScheduledQueryRules.LogToMetricAction"
)

// PossibleOdataTypeBasicActionValues returns an array of possible values for the OdataTypeBasicAction const type.
func PossibleOdataTypeBasicActionValues() []OdataTypeBasicAction {
	return []OdataTypeBasicAction{OdataTypeAction, OdataTypeMicrosoftWindowsAzureManagementMonitoringAlertsModelsMicrosoftAppInsightsNexusDataContractsResourcesScheduledQueryRulesAlertingAction, OdataTypeMicrosoftWindowsAzureManagementMonitoringAlertsModelsMicrosoftAppInsightsNexusDataContractsResourcesScheduledQueryRulesLogToMetricAction}
}

// OdataTypeBasicMetricAlertCriteria enumerates the values for odata type basic metric alert criteria.
type OdataTypeBasicMetricAlertCriteria string

const (
	// OdataTypeMetricAlertCriteria ...
	OdataTypeMetricAlertCriteria OdataTypeBasicMetricAlertCriteria = "MetricAlertCriteria"
	// OdataTypeMicrosoftAzureMonitorMultipleResourceMultipleMetricCriteria ...
	OdataTypeMicrosoftAzureMonitorMultipleResourceMultipleMetricCriteria OdataTypeBasicMetricAlertCriteria = "Microsoft.Azure.Monitor.MultipleResourceMultipleMetricCriteria"
	// OdataTypeMicrosoftAzureMonitorSingleResourceMultipleMetricCriteria ...
	OdataTypeMicrosoftAzureMonitorSingleResourceMultipleMetricCriteria OdataTypeBasicMetricAlertCriteria = "Microsoft.Azure.Monitor.SingleResourceMultipleMetricCriteria"
	// OdataTypeMicrosoftAzureMonitorWebtestLocationAvailabilityCriteria ...
	OdataTypeMicrosoftAzureMonitorWebtestLocationAvailabilityCriteria OdataTypeBasicMetricAlertCriteria = "Microsoft.Azure.Monitor.WebtestLocationAvailabilityCriteria"
)

// PossibleOdataTypeBasicMetricAlertCriteriaValues returns an array of possible values for the OdataTypeBasicMetricAlertCriteria const type.
func PossibleOdataTypeBasicMetricAlertCriteriaValues() []OdataTypeBasicMetricAlertCriteria {
	return []OdataTypeBasicMetricAlertCriteria{OdataTypeMetricAlertCriteria, OdataTypeMicrosoftAzureMonitorMultipleResourceMultipleMetricCriteria, OdataTypeMicrosoftAzureMonitorSingleResourceMultipleMetricCriteria, OdataTypeMicrosoftAzureMonitorWebtestLocationAvailabilityCriteria}
}

// OdataTypeBasicRuleAction enumerates the values for odata type basic rule action.
type OdataTypeBasicRuleAction string

const (
	// OdataTypeMicrosoftAzureManagementInsightsModelsRuleEmailAction ...
	OdataTypeMicrosoftAzureManagementInsightsModelsRuleEmailAction OdataTypeBasicRuleAction = "Microsoft.Azure.Management.Insights.Models.RuleEmailAction"
	// OdataTypeMicrosoftAzureManagementInsightsModelsRuleWebhookAction ...
	OdataTypeMicrosoftAzureManagementInsightsModelsRuleWebhookAction OdataTypeBasicRuleAction = "Microsoft.Azure.Management.Insights.Models.RuleWebhookAction"
	// OdataTypeRuleAction ...
	OdataTypeRuleAction OdataTypeBasicRuleAction = "RuleAction"
)

// PossibleOdataTypeBasicRuleActionValues returns an array of possible values for the OdataTypeBasicRuleAction const type.
func PossibleOdataTypeBasicRuleActionValues() []OdataTypeBasicRuleAction {
	return []OdataTypeBasicRuleAction{OdataTypeMicrosoftAzureManagementInsightsModelsRuleEmailAction, OdataTypeMicrosoftAzureManagementInsightsModelsRuleWebhookAction, OdataTypeRuleAction}
}

// OdataTypeBasicRuleCondition enumerates the values for odata type basic rule condition.
type OdataTypeBasicRuleCondition string

const (
	// OdataTypeMicrosoftAzureManagementInsightsModelsLocationThresholdRuleCondition ...
	OdataTypeMicrosoftAzureManagementInsightsModelsLocationThresholdRuleCondition OdataTypeBasicRuleCondition = "Microsoft.Azure.Management.Insights.Models.LocationThresholdRuleCondition"
	// OdataTypeMicrosoftAzureManagementInsightsModelsManagementEventRuleCondition ...
	OdataTypeMicrosoftAzureManagementInsightsModelsManagementEventRuleCondition OdataTypeBasicRuleCondition = "Microsoft.Azure.Management.Insights.Models.ManagementEventRuleCondition"
	// OdataTypeMicrosoftAzureManagementInsightsModelsThresholdRuleCondition ...
	OdataTypeMicrosoftAzureManagementInsightsModelsThresholdRuleCondition OdataTypeBasicRuleCondition = "Microsoft.Azure.Management.Insights.Models.ThresholdRuleCondition"
	// OdataTypeRuleCondition ...
	OdataTypeRuleCondition OdataTypeBasicRuleCondition = "RuleCondition"
)

// PossibleOdataTypeBasicRuleConditionValues returns an array of possible values for the OdataTypeBasicRuleCondition const type.
func PossibleOdataTypeBasicRuleConditionValues() []OdataTypeBasicRuleCondition {
	return []OdataTypeBasicRuleCondition{OdataTypeMicrosoftAzureManagementInsightsModelsLocationThresholdRuleCondition, OdataTypeMicrosoftAzureManagementInsightsModelsManagementEventRuleCondition, OdataTypeMicrosoftAzureManagementInsightsModelsThresholdRuleCondition, OdataTypeRuleCondition}
}

// Operator enumerates the values for operator.
type Operator string

const (
	// OperatorEquals ...
	OperatorEquals Operator = "Equals"
	// OperatorGreaterThan ...
	OperatorGreaterThan Operator = "GreaterThan"
	// OperatorGreaterThanOrEqual ...
	OperatorGreaterThanOrEqual Operator = "GreaterThanOrEqual"
	// OperatorLessThan ...
	OperatorLessThan Operator = "LessThan"
	// OperatorLessThanOrEqual ...
	OperatorLessThanOrEqual Operator = "LessThanOrEqual"
	// OperatorNotEquals ...
	OperatorNotEquals Operator = "NotEquals"
)

// PossibleOperatorValues returns an array of possible values for the Operator const type.
func PossibleOperatorValues() []Operator {
	return []Operator{OperatorEquals, OperatorGreaterThan, OperatorGreaterThanOrEqual, OperatorLessThan, OperatorLessThanOrEqual, OperatorNotEquals}
}

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// Canceled ...
	Canceled ProvisioningState = "Canceled"
	// Deploying ...
	Deploying ProvisioningState = "Deploying"
	// Failed ...
	Failed ProvisioningState = "Failed"
	// Succeeded ...
	Succeeded ProvisioningState = "Succeeded"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{Canceled, Deploying, Failed, Succeeded}
}

// QueryType enumerates the values for query type.
type QueryType string

const (
	// ResultCount ...
	ResultCount QueryType = "ResultCount"
)

// PossibleQueryTypeValues returns an array of possible values for the QueryType const type.
func PossibleQueryTypeValues() []QueryType {
	return []QueryType{ResultCount}
}

// ReceiverStatus enumerates the values for receiver status.
type ReceiverStatus string

const (
	// ReceiverStatusDisabled ...
	ReceiverStatusDisabled ReceiverStatus = "Disabled"
	// ReceiverStatusEnabled ...
	ReceiverStatusEnabled ReceiverStatus = "Enabled"
	// ReceiverStatusNotSpecified ...
	ReceiverStatusNotSpecified ReceiverStatus = "NotSpecified"
)

// PossibleReceiverStatusValues returns an array of possible values for the ReceiverStatus const type.
func PossibleReceiverStatusValues() []ReceiverStatus {
	return []ReceiverStatus{ReceiverStatusDisabled, ReceiverStatusEnabled, ReceiverStatusNotSpecified}
}

// RecurrenceFrequency enumerates the values for recurrence frequency.
type RecurrenceFrequency string

const (
	// RecurrenceFrequencyDay ...
	RecurrenceFrequencyDay RecurrenceFrequency = "Day"
	// RecurrenceFrequencyHour ...
	RecurrenceFrequencyHour RecurrenceFrequency = "Hour"
	// RecurrenceFrequencyMinute ...
	RecurrenceFrequencyMinute RecurrenceFrequency = "Minute"
	// RecurrenceFrequencyMonth ...
	RecurrenceFrequencyMonth RecurrenceFrequency = "Month"
	// RecurrenceFrequencyNone ...
	RecurrenceFrequencyNone RecurrenceFrequency = "None"
	// RecurrenceFrequencySecond ...
	RecurrenceFrequencySecond RecurrenceFrequency = "Second"
	// RecurrenceFrequencyWeek ...
	RecurrenceFrequencyWeek RecurrenceFrequency = "Week"
	// RecurrenceFrequencyYear ...
	RecurrenceFrequencyYear RecurrenceFrequency = "Year"
)

// PossibleRecurrenceFrequencyValues returns an array of possible values for the RecurrenceFrequency const type.
func PossibleRecurrenceFrequencyValues() []RecurrenceFrequency {
	return []RecurrenceFrequency{RecurrenceFrequencyDay, RecurrenceFrequencyHour, RecurrenceFrequencyMinute, RecurrenceFrequencyMonth, RecurrenceFrequencyNone, RecurrenceFrequencySecond, RecurrenceFrequencyWeek, RecurrenceFrequencyYear}
}

// ResultType enumerates the values for result type.
type ResultType string

const (
	// Data ...
	Data ResultType = "Data"
	// Metadata ...
	Metadata ResultType = "Metadata"
)

// PossibleResultTypeValues returns an array of possible values for the ResultType const type.
func PossibleResultTypeValues() []ResultType {
	return []ResultType{Data, Metadata}
}

// ScaleDirection enumerates the values for scale direction.
type ScaleDirection string

const (
	// ScaleDirectionDecrease ...
	ScaleDirectionDecrease ScaleDirection = "Decrease"
	// ScaleDirectionIncrease ...
	ScaleDirectionIncrease ScaleDirection = "Increase"
	// ScaleDirectionNone ...
	ScaleDirectionNone ScaleDirection = "None"
)

// PossibleScaleDirectionValues returns an array of possible values for the ScaleDirection const type.
func PossibleScaleDirectionValues() []ScaleDirection {
	return []ScaleDirection{ScaleDirectionDecrease, ScaleDirectionIncrease, ScaleDirectionNone}
}

// ScaleRuleMetricDimensionOperationType enumerates the values for scale rule metric dimension operation type.
type ScaleRuleMetricDimensionOperationType string

const (
	// ScaleRuleMetricDimensionOperationTypeEquals ...
	ScaleRuleMetricDimensionOperationTypeEquals ScaleRuleMetricDimensionOperationType = "Equals"
	// ScaleRuleMetricDimensionOperationTypeNotEquals ...
	ScaleRuleMetricDimensionOperationTypeNotEquals ScaleRuleMetricDimensionOperationType = "NotEquals"
)

// PossibleScaleRuleMetricDimensionOperationTypeValues returns an array of possible values for the ScaleRuleMetricDimensionOperationType const type.
func PossibleScaleRuleMetricDimensionOperationTypeValues() []ScaleRuleMetricDimensionOperationType {
	return []ScaleRuleMetricDimensionOperationType{ScaleRuleMetricDimensionOperationTypeEquals, ScaleRuleMetricDimensionOperationTypeNotEquals}
}

// ScaleType enumerates the values for scale type.
type ScaleType string

const (
	// ChangeCount ...
	ChangeCount ScaleType = "ChangeCount"
	// ExactCount ...
	ExactCount ScaleType = "ExactCount"
	// PercentChangeCount ...
	PercentChangeCount ScaleType = "PercentChangeCount"
)

// PossibleScaleTypeValues returns an array of possible values for the ScaleType const type.
func PossibleScaleTypeValues() []ScaleType {
	return []ScaleType{ChangeCount, ExactCount, PercentChangeCount}
}

// Sensitivity enumerates the values for sensitivity.
type Sensitivity string

const (
	// SensitivityHigh ...
	SensitivityHigh Sensitivity = "High"
	// SensitivityLow ...
	SensitivityLow Sensitivity = "Low"
	// SensitivityMedium ...
	SensitivityMedium Sensitivity = "Medium"
)

// PossibleSensitivityValues returns an array of possible values for the Sensitivity const type.
func PossibleSensitivityValues() []Sensitivity {
	return []Sensitivity{SensitivityHigh, SensitivityLow, SensitivityMedium}
}

// TimeAggregationOperator enumerates the values for time aggregation operator.
type TimeAggregationOperator string

const (
	// TimeAggregationOperatorAverage ...
	TimeAggregationOperatorAverage TimeAggregationOperator = "Average"
	// TimeAggregationOperatorLast ...
	TimeAggregationOperatorLast TimeAggregationOperator = "Last"
	// TimeAggregationOperatorMaximum ...
	TimeAggregationOperatorMaximum TimeAggregationOperator = "Maximum"
	// TimeAggregationOperatorMinimum ...
	TimeAggregationOperatorMinimum TimeAggregationOperator = "Minimum"
	// TimeAggregationOperatorTotal ...
	TimeAggregationOperatorTotal TimeAggregationOperator = "Total"
)

// PossibleTimeAggregationOperatorValues returns an array of possible values for the TimeAggregationOperator const type.
func PossibleTimeAggregationOperatorValues() []TimeAggregationOperator {
	return []TimeAggregationOperator{TimeAggregationOperatorAverage, TimeAggregationOperatorLast, TimeAggregationOperatorMaximum, TimeAggregationOperatorMinimum, TimeAggregationOperatorTotal}
}

// TimeAggregationType enumerates the values for time aggregation type.
type TimeAggregationType string

const (
	// TimeAggregationTypeAverage ...
	TimeAggregationTypeAverage TimeAggregationType = "Average"
	// TimeAggregationTypeCount ...
	TimeAggregationTypeCount TimeAggregationType = "Count"
	// TimeAggregationTypeLast ...
	TimeAggregationTypeLast TimeAggregationType = "Last"
	// TimeAggregationTypeMaximum ...
	TimeAggregationTypeMaximum TimeAggregationType = "Maximum"
	// TimeAggregationTypeMinimum ...
	TimeAggregationTypeMinimum TimeAggregationType = "Minimum"
	// TimeAggregationTypeTotal ...
	TimeAggregationTypeTotal TimeAggregationType = "Total"
)

// PossibleTimeAggregationTypeValues returns an array of possible values for the TimeAggregationType const type.
func PossibleTimeAggregationTypeValues() []TimeAggregationType {
	return []TimeAggregationType{TimeAggregationTypeAverage, TimeAggregationTypeCount, TimeAggregationTypeLast, TimeAggregationTypeMaximum, TimeAggregationTypeMinimum, TimeAggregationTypeTotal}
}

// Unit enumerates the values for unit.
type Unit string

const (
	// UnitBitsPerSecond ...
	UnitBitsPerSecond Unit = "BitsPerSecond"
	// UnitBytes ...
	UnitBytes Unit = "Bytes"
	// UnitByteSeconds ...
	UnitByteSeconds Unit = "ByteSeconds"
	// UnitBytesPerSecond ...
	UnitBytesPerSecond Unit = "BytesPerSecond"
	// UnitCores ...
	UnitCores Unit = "Cores"
	// UnitCount ...
	UnitCount Unit = "Count"
	// UnitCountPerSecond ...
	UnitCountPerSecond Unit = "CountPerSecond"
	// UnitMilliCores ...
	UnitMilliCores Unit = "MilliCores"
	// UnitMilliSeconds ...
	UnitMilliSeconds Unit = "MilliSeconds"
	// UnitNanoCores ...
	UnitNanoCores Unit = "NanoCores"
	// UnitPercent ...
	UnitPercent Unit = "Percent"
	// UnitSeconds ...
	UnitSeconds Unit = "Seconds"
	// UnitUnspecified ...
	UnitUnspecified Unit = "Unspecified"
)

// PossibleUnitValues returns an array of possible values for the Unit const type.
func PossibleUnitValues() []Unit {
	return []Unit{UnitBitsPerSecond, UnitBytes, UnitByteSeconds, UnitBytesPerSecond, UnitCores, UnitCount, UnitCountPerSecond, UnitMilliCores, UnitMilliSeconds, UnitNanoCores, UnitPercent, UnitSeconds, UnitUnspecified}
}
