//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// GalleryImageVersionsClient contains the methods for the GalleryImageVersions group.
// Don't use this type directly, use NewGalleryImageVersionsClient() instead.
type GalleryImageVersionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewGalleryImageVersionsClient creates a new instance of GalleryImageVersionsClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewGalleryImageVersionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*GalleryImageVersionsClient, error) {
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
	client := &GalleryImageVersionsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or update a gallery image version.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// galleryName - The name of the Shared Image Gallery in which the Image Definition resides.
// galleryImageName - The name of the gallery image definition in which the Image Version is to be created.
// galleryImageVersionName - The name of the gallery image version to be created. Needs to follow semantic version name pattern:
// The allowed characters are digit and period. Digits must be within the range of a 32-bit integer.
// Format: ..
// galleryImageVersion - Parameters supplied to the create or update gallery image version operation.
// options - GalleryImageVersionsClientBeginCreateOrUpdateOptions contains the optional parameters for the GalleryImageVersionsClient.BeginCreateOrUpdate
// method.
func (client *GalleryImageVersionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string, galleryImageVersionName string, galleryImageVersion GalleryImageVersion, options *GalleryImageVersionsClientBeginCreateOrUpdateOptions) (*armruntime.Poller[GalleryImageVersionsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, galleryName, galleryImageName, galleryImageVersionName, galleryImageVersion, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[GalleryImageVersionsClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[GalleryImageVersionsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Create or update a gallery image version.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *GalleryImageVersionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string, galleryImageVersionName string, galleryImageVersion GalleryImageVersion, options *GalleryImageVersionsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, galleryName, galleryImageName, galleryImageVersionName, galleryImageVersion, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *GalleryImageVersionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string, galleryImageVersionName string, galleryImageVersion GalleryImageVersion, options *GalleryImageVersionsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/images/{galleryImageName}/versions/{galleryImageVersionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if galleryName == "" {
		return nil, errors.New("parameter galleryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	if galleryImageName == "" {
		return nil, errors.New("parameter galleryImageName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryImageName}", url.PathEscape(galleryImageName))
	if galleryImageVersionName == "" {
		return nil, errors.New("parameter galleryImageVersionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryImageVersionName}", url.PathEscape(galleryImageVersionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, galleryImageVersion)
}

// BeginDelete - Delete a gallery image version.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// galleryName - The name of the Shared Image Gallery in which the Image Definition resides.
// galleryImageName - The name of the gallery image definition in which the Image Version resides.
// galleryImageVersionName - The name of the gallery image version to be deleted.
// options - GalleryImageVersionsClientBeginDeleteOptions contains the optional parameters for the GalleryImageVersionsClient.BeginDelete
// method.
func (client *GalleryImageVersionsClient) BeginDelete(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string, galleryImageVersionName string, options *GalleryImageVersionsClientBeginDeleteOptions) (*armruntime.Poller[GalleryImageVersionsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, galleryName, galleryImageName, galleryImageVersionName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[GalleryImageVersionsClientDeleteResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[GalleryImageVersionsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Delete a gallery image version.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *GalleryImageVersionsClient) deleteOperation(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string, galleryImageVersionName string, options *GalleryImageVersionsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, galleryName, galleryImageName, galleryImageVersionName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *GalleryImageVersionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string, galleryImageVersionName string, options *GalleryImageVersionsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/images/{galleryImageName}/versions/{galleryImageVersionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if galleryName == "" {
		return nil, errors.New("parameter galleryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	if galleryImageName == "" {
		return nil, errors.New("parameter galleryImageName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryImageName}", url.PathEscape(galleryImageName))
	if galleryImageVersionName == "" {
		return nil, errors.New("parameter galleryImageVersionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryImageVersionName}", url.PathEscape(galleryImageVersionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Retrieves information about a gallery image version.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// galleryName - The name of the Shared Image Gallery in which the Image Definition resides.
// galleryImageName - The name of the gallery image definition in which the Image Version resides.
// galleryImageVersionName - The name of the gallery image version to be retrieved.
// options - GalleryImageVersionsClientGetOptions contains the optional parameters for the GalleryImageVersionsClient.Get
// method.
func (client *GalleryImageVersionsClient) Get(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string, galleryImageVersionName string, options *GalleryImageVersionsClientGetOptions) (GalleryImageVersionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, galleryName, galleryImageName, galleryImageVersionName, options)
	if err != nil {
		return GalleryImageVersionsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return GalleryImageVersionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return GalleryImageVersionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *GalleryImageVersionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string, galleryImageVersionName string, options *GalleryImageVersionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/images/{galleryImageName}/versions/{galleryImageVersionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if galleryName == "" {
		return nil, errors.New("parameter galleryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	if galleryImageName == "" {
		return nil, errors.New("parameter galleryImageName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryImageName}", url.PathEscape(galleryImageName))
	if galleryImageVersionName == "" {
		return nil, errors.New("parameter galleryImageVersionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryImageVersionName}", url.PathEscape(galleryImageVersionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", string(*options.Expand))
	}
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *GalleryImageVersionsClient) getHandleResponse(resp *http.Response) (GalleryImageVersionsClientGetResponse, error) {
	result := GalleryImageVersionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GalleryImageVersion); err != nil {
		return GalleryImageVersionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByGalleryImagePager - List gallery image versions in a gallery image definition.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// galleryName - The name of the Shared Image Gallery in which the Image Definition resides.
// galleryImageName - The name of the Shared Image Gallery Image Definition from which the Image Versions are to be listed.
// options - GalleryImageVersionsClientListByGalleryImageOptions contains the optional parameters for the GalleryImageVersionsClient.ListByGalleryImage
// method.
func (client *GalleryImageVersionsClient) NewListByGalleryImagePager(resourceGroupName string, galleryName string, galleryImageName string, options *GalleryImageVersionsClientListByGalleryImageOptions) *runtime.Pager[GalleryImageVersionsClientListByGalleryImageResponse] {
	return runtime.NewPager(runtime.PageProcessor[GalleryImageVersionsClientListByGalleryImageResponse]{
		More: func(page GalleryImageVersionsClientListByGalleryImageResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *GalleryImageVersionsClientListByGalleryImageResponse) (GalleryImageVersionsClientListByGalleryImageResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByGalleryImageCreateRequest(ctx, resourceGroupName, galleryName, galleryImageName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return GalleryImageVersionsClientListByGalleryImageResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return GalleryImageVersionsClientListByGalleryImageResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return GalleryImageVersionsClientListByGalleryImageResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByGalleryImageHandleResponse(resp)
		},
	})
}

// listByGalleryImageCreateRequest creates the ListByGalleryImage request.
func (client *GalleryImageVersionsClient) listByGalleryImageCreateRequest(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string, options *GalleryImageVersionsClientListByGalleryImageOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/images/{galleryImageName}/versions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if galleryName == "" {
		return nil, errors.New("parameter galleryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	if galleryImageName == "" {
		return nil, errors.New("parameter galleryImageName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryImageName}", url.PathEscape(galleryImageName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByGalleryImageHandleResponse handles the ListByGalleryImage response.
func (client *GalleryImageVersionsClient) listByGalleryImageHandleResponse(resp *http.Response) (GalleryImageVersionsClientListByGalleryImageResponse, error) {
	result := GalleryImageVersionsClientListByGalleryImageResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GalleryImageVersionList); err != nil {
		return GalleryImageVersionsClientListByGalleryImageResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update a gallery image version.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// galleryName - The name of the Shared Image Gallery in which the Image Definition resides.
// galleryImageName - The name of the gallery image definition in which the Image Version is to be updated.
// galleryImageVersionName - The name of the gallery image version to be updated. Needs to follow semantic version name pattern:
// The allowed characters are digit and period. Digits must be within the range of a 32-bit integer.
// Format: ..
// galleryImageVersion - Parameters supplied to the update gallery image version operation.
// options - GalleryImageVersionsClientBeginUpdateOptions contains the optional parameters for the GalleryImageVersionsClient.BeginUpdate
// method.
func (client *GalleryImageVersionsClient) BeginUpdate(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string, galleryImageVersionName string, galleryImageVersion GalleryImageVersionUpdate, options *GalleryImageVersionsClientBeginUpdateOptions) (*armruntime.Poller[GalleryImageVersionsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, galleryName, galleryImageName, galleryImageVersionName, galleryImageVersion, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[GalleryImageVersionsClientUpdateResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[GalleryImageVersionsClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Update a gallery image version.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *GalleryImageVersionsClient) update(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string, galleryImageVersionName string, galleryImageVersion GalleryImageVersionUpdate, options *GalleryImageVersionsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, galleryName, galleryImageName, galleryImageVersionName, galleryImageVersion, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *GalleryImageVersionsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string, galleryImageVersionName string, galleryImageVersion GalleryImageVersionUpdate, options *GalleryImageVersionsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/images/{galleryImageName}/versions/{galleryImageVersionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if galleryName == "" {
		return nil, errors.New("parameter galleryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	if galleryImageName == "" {
		return nil, errors.New("parameter galleryImageName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryImageName}", url.PathEscape(galleryImageName))
	if galleryImageVersionName == "" {
		return nil, errors.New("parameter galleryImageVersionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryImageVersionName}", url.PathEscape(galleryImageVersionName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, galleryImageVersion)
}
