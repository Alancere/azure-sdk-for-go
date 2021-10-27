package main

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"log"
	"os"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/web/armweb"
)

var (
	subscriptionID     string
	location           = "westus"
	resourceGroupName  = "sample-resource-group3"
	appServicePlanName = "sample-web-plan"
)

func main() {
	subscriptionID = os.Getenv("AZURE_SUBSCRIPTION_ID")
	if len(subscriptionID) == 0 {
		log.Fatal("AZURE_SUBSCRIPTION_ID is not set.")
	}

	//cred, err := azidentity.NewDefaultAzureCredential(nil)
	cred, err := azidentity.NewEnvironmentCredential(nil)
	if err != nil {
		log.Fatal(err)
	}

	//conn := arm.NewDefaultConnection(cred, &arm.ConnectionOptions{
	//	Logging: policy.LogOptions{
	//		IncludeBody: true,
	//	},
	//})
	ctx := context.Background()

	resourceGroup, err := createResourceGroup(ctx,cred)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("resources group:", *resourceGroup.ID)

	appServicePlan, err := createAppServicePlan(ctx,cred)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("app service plan:", *appServicePlan.ID)

	//keepResource := os.Getenv("KEEP_RESOURCE")
	//if len(keepResource) == 0 {
	//	_, err := cleanup(ctx, conn)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	log.Println("cleaned up successfully.")
	//}
}

func createAppServicePlan(ctx context.Context,cred azcore.TokenCredential) (*armweb.AppServicePlan, error) {
	appServicePlansClient := armweb.NewAppServicePlansClient(subscriptionID,cred,nil)

	pollerResp, err := appServicePlansClient.BeginCreateOrUpdate(
		ctx,
		resourceGroupName,
		appServicePlanName,
		armweb.AppServicePlan{
			Resource: armweb.Resource{
				Location: to.StringPtr(location),
			},
			SKU: &armweb.SKUDescription{
				Name:     to.StringPtr("P1V2"),
				Capacity: to.Int32Ptr(1),
			},
			Properties: &armweb.AppServicePlanProperties{
				PerSiteScaling: to.BoolPtr(false),
				IsXenon:        to.BoolPtr(false),
				FreeOfferExpirationTime: to.TimePtr(time.Now().AddDate(0, 0, 7)),
			},
		},
		nil,
	)
	if err != nil {
		return nil, err
	}
	resp, err := pollerResp.PollUntilDone(ctx, 10*time.Second)
	if err != nil {
		return nil, err
	}
	return &resp.AppServicePlan, nil
}

func createResourceGroup(ctx context.Context,cred azcore.TokenCredential) (*armresources.ResourceGroup, error) {
	resourceGroupClient := armresources.NewResourceGroupsClient(subscriptionID,cred,nil)

	resourceGroupResp, err := resourceGroupClient.CreateOrUpdate(
		ctx,
		resourceGroupName,
		armresources.ResourceGroup{
			Location: to.StringPtr(location),
		},
		nil)
	if err != nil {
		return nil, err
	}
	return &resourceGroupResp.ResourceGroup, nil
}

//func cleanup(ctx context.Context, conn *arm.Connection) (*http.Response, error) {
//	resourceGroupClient := armresources.NewResourceGroupsClient(conn, subscriptionID)
//
//	pollerResp, err := resourceGroupClient.BeginDelete(ctx, resourceGroupName, nil)
//	if err != nil {
//		return nil, err
//	}
//
//	resp, err := pollerResp.PollUntilDone(ctx, 10*time.Second)
//	if err != nil {
//		return nil, err
//	}
//	return resp.RawResponse, nil
//}
