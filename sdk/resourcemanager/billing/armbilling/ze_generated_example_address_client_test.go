//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/AddressInvalid.json
func ExampleAddressClient_Validate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armbilling.NewAddressClient(cred, nil)
	res, err := client.Validate(ctx,
		armbilling.AddressDetails{
			AddressLine1: to.StringPtr("<address-line1>"),
			City:         to.StringPtr("<city>"),
			Country:      to.StringPtr("<country>"),
			PostalCode:   to.StringPtr("<postal-code>"),
			Region:       to.StringPtr("<region>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.AddressClientValidateResult)
}
