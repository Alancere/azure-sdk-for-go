//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsupport

// ChatTranscriptsClientGetOptions contains the optional parameters for the ChatTranscriptsClient.Get method.
type ChatTranscriptsClientGetOptions struct {
	// placeholder for future optional parameters
}

// ChatTranscriptsClientListOptions contains the optional parameters for the ChatTranscriptsClient.NewListPager method.
type ChatTranscriptsClientListOptions struct {
	// placeholder for future optional parameters
}

// ChatTranscriptsNoSubscriptionClientGetOptions contains the optional parameters for the ChatTranscriptsNoSubscriptionClient.Get
// method.
type ChatTranscriptsNoSubscriptionClientGetOptions struct {
	// placeholder for future optional parameters
}

// CommunicationsClientBeginCreateOptions contains the optional parameters for the CommunicationsClient.BeginCreate method.
type CommunicationsClientBeginCreateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// CommunicationsClientCheckNameAvailabilityOptions contains the optional parameters for the CommunicationsClient.CheckNameAvailability
// method.
type CommunicationsClientCheckNameAvailabilityOptions struct {
	// placeholder for future optional parameters
}

// CommunicationsClientGetOptions contains the optional parameters for the CommunicationsClient.Get method.
type CommunicationsClientGetOptions struct {
	// placeholder for future optional parameters
}

// CommunicationsClientListOptions contains the optional parameters for the CommunicationsClient.NewListPager method.
type CommunicationsClientListOptions struct {
	// The filter to apply on the operation. You can filter by communicationType and createdDate properties. CommunicationType
	// supports Equals ('eq') operator and createdDate supports Greater Than ('gt') and
	// Greater Than or Equals ('ge') operators. You may combine the CommunicationType and CreatedDate filters by Logical And ('and')
	// operator.
	Filter *string

	// The number of values to return in the collection. Default is 10 and max is 10.
	Top *int32
}

// CommunicationsNoSubscriptionClientBeginCreateOptions contains the optional parameters for the CommunicationsNoSubscriptionClient.BeginCreate
// method.
type CommunicationsNoSubscriptionClientBeginCreateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// CommunicationsNoSubscriptionClientCheckNameAvailabilityOptions contains the optional parameters for the CommunicationsNoSubscriptionClient.CheckNameAvailability
// method.
type CommunicationsNoSubscriptionClientCheckNameAvailabilityOptions struct {
	// placeholder for future optional parameters
}

// CommunicationsNoSubscriptionClientGetOptions contains the optional parameters for the CommunicationsNoSubscriptionClient.Get
// method.
type CommunicationsNoSubscriptionClientGetOptions struct {
	// placeholder for future optional parameters
}

// FileWorkspacesClientCreateOptions contains the optional parameters for the FileWorkspacesClient.Create method.
type FileWorkspacesClientCreateOptions struct {
	// placeholder for future optional parameters
}

// FileWorkspacesClientGetOptions contains the optional parameters for the FileWorkspacesClient.Get method.
type FileWorkspacesClientGetOptions struct {
	// placeholder for future optional parameters
}

// FileWorkspacesNoSubscriptionClientCreateOptions contains the optional parameters for the FileWorkspacesNoSubscriptionClient.Create
// method.
type FileWorkspacesNoSubscriptionClientCreateOptions struct {
	// placeholder for future optional parameters
}

// FileWorkspacesNoSubscriptionClientGetOptions contains the optional parameters for the FileWorkspacesNoSubscriptionClient.Get
// method.
type FileWorkspacesNoSubscriptionClientGetOptions struct {
	// placeholder for future optional parameters
}

// FilesClientCreateOptions contains the optional parameters for the FilesClient.Create method.
type FilesClientCreateOptions struct {
	// placeholder for future optional parameters
}

// FilesClientGetOptions contains the optional parameters for the FilesClient.Get method.
type FilesClientGetOptions struct {
	// placeholder for future optional parameters
}

// FilesClientListOptions contains the optional parameters for the FilesClient.NewListPager method.
type FilesClientListOptions struct {
	// placeholder for future optional parameters
}

// FilesClientUploadOptions contains the optional parameters for the FilesClient.Upload method.
type FilesClientUploadOptions struct {
	// placeholder for future optional parameters
}

// FilesNoSubscriptionClientCreateOptions contains the optional parameters for the FilesNoSubscriptionClient.Create method.
type FilesNoSubscriptionClientCreateOptions struct {
	// placeholder for future optional parameters
}

// FilesNoSubscriptionClientGetOptions contains the optional parameters for the FilesNoSubscriptionClient.Get method.
type FilesNoSubscriptionClientGetOptions struct {
	// placeholder for future optional parameters
}

// FilesNoSubscriptionClientListOptions contains the optional parameters for the FilesNoSubscriptionClient.NewListPager method.
type FilesNoSubscriptionClientListOptions struct {
	// placeholder for future optional parameters
}

// FilesNoSubscriptionClientUploadOptions contains the optional parameters for the FilesNoSubscriptionClient.Upload method.
type FilesNoSubscriptionClientUploadOptions struct {
	// placeholder for future optional parameters
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.NewListPager method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// ProblemClassificationsClientGetOptions contains the optional parameters for the ProblemClassificationsClient.Get method.
type ProblemClassificationsClientGetOptions struct {
	// placeholder for future optional parameters
}

// ProblemClassificationsClientListOptions contains the optional parameters for the ProblemClassificationsClient.NewListPager
// method.
type ProblemClassificationsClientListOptions struct {
	// placeholder for future optional parameters
}

// ServicesClientGetOptions contains the optional parameters for the ServicesClient.Get method.
type ServicesClientGetOptions struct {
	// placeholder for future optional parameters
}

// ServicesClientListOptions contains the optional parameters for the ServicesClient.NewListPager method.
type ServicesClientListOptions struct {
	// placeholder for future optional parameters
}

// TicketChatTranscriptsNoSubscriptionClientListOptions contains the optional parameters for the TicketChatTranscriptsNoSubscriptionClient.NewListPager
// method.
type TicketChatTranscriptsNoSubscriptionClientListOptions struct {
	// placeholder for future optional parameters
}

// TicketCommunicationsNoSubscriptionClientListOptions contains the optional parameters for the TicketCommunicationsNoSubscriptionClient.NewListPager
// method.
type TicketCommunicationsNoSubscriptionClientListOptions struct {
	// The filter to apply on the operation. You can filter by communicationType and createdDate properties. CommunicationType
	// supports Equals ('eq') operator and createdDate supports Greater Than ('gt') and
	// Greater Than or Equals ('ge') operators. You may combine the CommunicationType and CreatedDate filters by Logical And ('and')
	// operator.
	Filter *string

	// The number of values to return in the collection. Default is 10 and max is 10.
	Top *int32
}

// TicketsClientBeginCreateOptions contains the optional parameters for the TicketsClient.BeginCreate method.
type TicketsClientBeginCreateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// TicketsClientCheckNameAvailabilityOptions contains the optional parameters for the TicketsClient.CheckNameAvailability
// method.
type TicketsClientCheckNameAvailabilityOptions struct {
	// placeholder for future optional parameters
}

// TicketsClientGetOptions contains the optional parameters for the TicketsClient.Get method.
type TicketsClientGetOptions struct {
	// placeholder for future optional parameters
}

// TicketsClientListOptions contains the optional parameters for the TicketsClient.NewListPager method.
type TicketsClientListOptions struct {
	// The filter to apply on the operation. We support 'odata v4.0' filter semantics. Learn more [https://docs.microsoft.com/odata/concepts/queryoptions-overview].
	// Status, ServiceId, and
	// ProblemClassificationId filters can only be used with Equals ('eq') operator. For CreatedDate filter, the supported operators
	// are Greater Than ('gt') and Greater Than or Equals ('ge'). When using both
	// filters, combine them using the logical 'AND'.
	Filter *string

	// The number of values to return in the collection. Default is 25 and max is 100.
	Top *int32
}

// TicketsClientUpdateOptions contains the optional parameters for the TicketsClient.Update method.
type TicketsClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// TicketsNoSubscriptionClientBeginCreateOptions contains the optional parameters for the TicketsNoSubscriptionClient.BeginCreate
// method.
type TicketsNoSubscriptionClientBeginCreateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// TicketsNoSubscriptionClientCheckNameAvailabilityOptions contains the optional parameters for the TicketsNoSubscriptionClient.CheckNameAvailability
// method.
type TicketsNoSubscriptionClientCheckNameAvailabilityOptions struct {
	// placeholder for future optional parameters
}

// TicketsNoSubscriptionClientGetOptions contains the optional parameters for the TicketsNoSubscriptionClient.Get method.
type TicketsNoSubscriptionClientGetOptions struct {
	// placeholder for future optional parameters
}

// TicketsNoSubscriptionClientListOptions contains the optional parameters for the TicketsNoSubscriptionClient.NewListPager
// method.
type TicketsNoSubscriptionClientListOptions struct {
	// The filter to apply on the operation. We support 'odata v4.0' filter semantics. Learn more [https://docs.microsoft.com/odata/concepts/queryoptions-overview]
	// Status , ServiceId, and ProblemClassificationId filters can only be used with 'eq' operator. For CreatedDate filter, the
	// supported operators are 'gt' and 'ge'. When using both filters, combine them
	// using the logical 'AND'.
	Filter *string

	// The number of values to return in the collection. Default is 25 and max is 100.
	Top *int32
}

// TicketsNoSubscriptionClientUpdateOptions contains the optional parameters for the TicketsNoSubscriptionClient.Update method.
type TicketsNoSubscriptionClientUpdateOptions struct {
	// placeholder for future optional parameters
}
