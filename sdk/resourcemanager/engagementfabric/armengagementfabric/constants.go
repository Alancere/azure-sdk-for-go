//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armengagementfabric

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/engagementfabric/armengagementfabric"
	moduleVersion = "v0.3.0"
)

// CheckNameUnavailableReason - The reason of name availability result
type CheckNameUnavailableReason string

const (
	CheckNameUnavailableReasonAlreadyExists CheckNameUnavailableReason = "AlreadyExists"
	CheckNameUnavailableReasonInvalid       CheckNameUnavailableReason = "Invalid"
)

// PossibleCheckNameUnavailableReasonValues returns the possible values for the CheckNameUnavailableReason const type.
func PossibleCheckNameUnavailableReasonValues() []CheckNameUnavailableReason {
	return []CheckNameUnavailableReason{
		CheckNameUnavailableReasonAlreadyExists,
		CheckNameUnavailableReasonInvalid,
	}
}

// KeyRank - The rank of the EngagementFabric account key
type KeyRank string

const (
	KeyRankPrimaryKey   KeyRank = "PrimaryKey"
	KeyRankSecondaryKey KeyRank = "SecondaryKey"
)

// PossibleKeyRankValues returns the possible values for the KeyRank const type.
func PossibleKeyRankValues() []KeyRank {
	return []KeyRank{
		KeyRankPrimaryKey,
		KeyRankSecondaryKey,
	}
}
