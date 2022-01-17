//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservicesbackup

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// BackupJobsClient contains the methods for the BackupJobs group.
// Don't use this type directly, use NewBackupJobsClient() instead.
type BackupJobsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewBackupJobsClient creates a new instance of BackupJobsClient with the specified values.
// subscriptionID - The subscription Id.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewBackupJobsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *BackupJobsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &BackupJobsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// List - Provides a pageable list of jobs.
// If the operation fails it returns an *azcore.ResponseError type.
// vaultName - The name of the recovery services vault.
// resourceGroupName - The name of the resource group where the recovery services vault is present.
// options - BackupJobsClientListOptions contains the optional parameters for the BackupJobsClient.List method.
func (client *BackupJobsClient) List(vaultName string, resourceGroupName string, options *BackupJobsClientListOptions) *BackupJobsClientListPager {
	return &BackupJobsClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, vaultName, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp BackupJobsClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.JobResourceList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *BackupJobsClient) listCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, options *BackupJobsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupJobs"
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *BackupJobsClient) listHandleResponse(resp *http.Response) (BackupJobsClientListResponse, error) {
	result := BackupJobsClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobResourceList); err != nil {
		return BackupJobsClientListResponse{}, err
	}
	return result, nil
}
