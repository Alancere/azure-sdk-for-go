//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package armfeatures_test

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/internal/testutil"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armfeatures"
	"github.com/stretchr/testify/suite"
	"testing"
)

type FeaturesClientTestSuite struct {
	suite.Suite

	ctx               context.Context
	cred              azcore.TokenCredential
	options           *arm.ClientOptions
	location          string
	resourceGroupName string
	subscriptionID    string
}

func (testsuite *FeaturesClientTestSuite) SetupSuite() {
	testsuite.ctx = context.Background()
	testsuite.cred, testsuite.options = testutil.GetCredAndClientOptions(testsuite.T())
	testsuite.location = testutil.GetEnv("LOCATION", "eastus")
	testsuite.subscriptionID = testutil.GetEnv("AZURE_SUBSCRIPTION_ID", "00000000-0000-0000-0000-000000000000")
	testutil.StartRecording(testsuite.T(), "sdk/resourcemanager/resources/armfeatures/testdata")
	testsuite.resourceGroupName = *testutil.CreateResourceGroup(testsuite.T(), testsuite.ctx, testsuite.subscriptionID, testsuite.cred, testsuite.options, testsuite.location).Name
}

func (testsuite *FeaturesClientTestSuite) TearDownSuite() {
	testutil.DeleteResourceGroup(testsuite.T(), testsuite.ctx, testsuite.subscriptionID, testsuite.cred, testsuite.options, testsuite.resourceGroupName)
	testutil.StopRecording(testsuite.T())
}

func TestFeaturesClient(t *testing.T) {
	suite.Run(t, new(FeaturesClientTestSuite))
}

func (testsuite *FeaturesClientTestSuite) TestFeaturesCRUD() {
	featureClient := armfeatures.NewClient(testsuite.subscriptionID, testsuite.cred, testsuite.options)

	// register
	//featureName := "feature"
	//register, err := featureClient.Register(testsuite.ctx, "Microsoft.Network", featureName, nil)
	//testsuite.Require().NoError(err)
	//testsuite.Require().Equal(featureName, *register.Name)
	//
	//// unregister
	//unregister, err := featureClient.Unregister(testsuite.ctx, "Microsoft.Network", featureName, nil)
	//testsuite.Require().NoError(err)
	//testsuite.Require().Equal(featureName, *unregister.Name)

	// get
	//_, err := featureClient.Get(testsuite.ctx, "Microsoft.Network", featureName, nil)
	//testsuite.Require().NoError(err)

	// list
	pager := featureClient.List("Microsoft.Network", nil)
	testsuite.Require().NoError(pager.Err())

	// list all
	listAll := featureClient.ListAll(nil)
	testsuite.Require().NoError(listAll.Err())

	// list operation
	featuresClient := armfeatures.NewFeatureClient(testsuite.cred, testsuite.options)
	listOperations := featuresClient.ListOperations(nil)
	testsuite.Require().NoError(listOperations.Err())
}
