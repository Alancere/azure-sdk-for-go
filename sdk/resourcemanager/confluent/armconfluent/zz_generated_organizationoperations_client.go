//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armconfluent

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// OrganizationOperationsClient contains the methods for the OrganizationOperations group.
// Don't use this type directly, use NewOrganizationOperationsClient() instead.
type OrganizationOperationsClient struct {
	host string
	pl   runtime.Pipeline
}

// NewOrganizationOperationsClient creates a new instance of OrganizationOperationsClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewOrganizationOperationsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*OrganizationOperationsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublicCloud.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &OrganizationOperationsClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// List - List all operations provided by Microsoft.Confluent.
// If the operation fails it returns an *azcore.ResponseError type.
// options - OrganizationOperationsClientListOptions contains the optional parameters for the OrganizationOperationsClient.List
// method.
func (client *OrganizationOperationsClient) List(options *OrganizationOperationsClientListOptions) *runtime.Pager[OrganizationOperationsClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[OrganizationOperationsClientListResponse]{
		More: func(page OrganizationOperationsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *OrganizationOperationsClientListResponse) (OrganizationOperationsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return OrganizationOperationsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return OrganizationOperationsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return OrganizationOperationsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *OrganizationOperationsClient) listCreateRequest(ctx context.Context, options *OrganizationOperationsClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Confluent/operations"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *OrganizationOperationsClient) listHandleResponse(resp *http.Response) (OrganizationOperationsClientListResponse, error) {
	result := OrganizationOperationsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OperationListResult); err != nil {
		return OrganizationOperationsClientListResponse{}, err
	}
	return result, nil
}
