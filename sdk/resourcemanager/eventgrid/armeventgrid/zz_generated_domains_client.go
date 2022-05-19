//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armeventgrid

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
	"strconv"
	"strings"
)

// DomainsClient contains the methods for the Domains group.
// Don't use this type directly, use NewDomainsClient() instead.
type DomainsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewDomainsClient creates a new instance of DomainsClient with the specified values.
// subscriptionID - Subscription credentials that uniquely identify a Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewDomainsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DomainsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &DomainsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Asynchronously creates or updates a new domain with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-10-15-preview
// resourceGroupName - The name of the resource group within the user's subscription.
// domainName - Name of the domain.
// domainInfo - Domain information.
// options - DomainsClientBeginCreateOrUpdateOptions contains the optional parameters for the DomainsClient.BeginCreateOrUpdate
// method.
func (client *DomainsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, domainName string, domainInfo Domain, options *DomainsClientBeginCreateOrUpdateOptions) (*runtime.Poller[DomainsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, domainName, domainInfo, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[DomainsClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[DomainsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Asynchronously creates or updates a new domain with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-10-15-preview
func (client *DomainsClient) createOrUpdate(ctx context.Context, resourceGroupName string, domainName string, domainInfo Domain, options *DomainsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, domainName, domainInfo, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DomainsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, domainName string, domainInfo Domain, options *DomainsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/domains/{domainName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if domainName == "" {
		return nil, errors.New("parameter domainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainName}", url.PathEscape(domainName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, domainInfo)
}

// BeginDelete - Delete existing domain.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-10-15-preview
// resourceGroupName - The name of the resource group within the user's subscription.
// domainName - Name of the domain.
// options - DomainsClientBeginDeleteOptions contains the optional parameters for the DomainsClient.BeginDelete method.
func (client *DomainsClient) BeginDelete(ctx context.Context, resourceGroupName string, domainName string, options *DomainsClientBeginDeleteOptions) (*runtime.Poller[DomainsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, domainName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[DomainsClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[DomainsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Delete existing domain.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-10-15-preview
func (client *DomainsClient) deleteOperation(ctx context.Context, resourceGroupName string, domainName string, options *DomainsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, domainName, options)
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
func (client *DomainsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, domainName string, options *DomainsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/domains/{domainName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if domainName == "" {
		return nil, errors.New("parameter domainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainName}", url.PathEscape(domainName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Get properties of a domain.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-10-15-preview
// resourceGroupName - The name of the resource group within the user's subscription.
// domainName - Name of the domain.
// options - DomainsClientGetOptions contains the optional parameters for the DomainsClient.Get method.
func (client *DomainsClient) Get(ctx context.Context, resourceGroupName string, domainName string, options *DomainsClientGetOptions) (DomainsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, domainName, options)
	if err != nil {
		return DomainsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DomainsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DomainsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DomainsClient) getCreateRequest(ctx context.Context, resourceGroupName string, domainName string, options *DomainsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/domains/{domainName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if domainName == "" {
		return nil, errors.New("parameter domainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainName}", url.PathEscape(domainName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DomainsClient) getHandleResponse(resp *http.Response) (DomainsClientGetResponse, error) {
	result := DomainsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Domain); err != nil {
		return DomainsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - List all the domains under a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-10-15-preview
// resourceGroupName - The name of the resource group within the user's subscription.
// options - DomainsClientListByResourceGroupOptions contains the optional parameters for the DomainsClient.ListByResourceGroup
// method.
func (client *DomainsClient) NewListByResourceGroupPager(resourceGroupName string, options *DomainsClientListByResourceGroupOptions) *runtime.Pager[DomainsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[DomainsClientListByResourceGroupResponse]{
		More: func(page DomainsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DomainsClientListByResourceGroupResponse) (DomainsClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DomainsClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DomainsClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DomainsClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *DomainsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *DomainsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/domains"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-15-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *DomainsClient) listByResourceGroupHandleResponse(resp *http.Response) (DomainsClientListByResourceGroupResponse, error) {
	result := DomainsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DomainsListResult); err != nil {
		return DomainsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - List all the domains under an Azure subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-10-15-preview
// options - DomainsClientListBySubscriptionOptions contains the optional parameters for the DomainsClient.ListBySubscription
// method.
func (client *DomainsClient) NewListBySubscriptionPager(options *DomainsClientListBySubscriptionOptions) *runtime.Pager[DomainsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[DomainsClientListBySubscriptionResponse]{
		More: func(page DomainsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DomainsClientListBySubscriptionResponse) (DomainsClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DomainsClientListBySubscriptionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DomainsClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DomainsClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *DomainsClient) listBySubscriptionCreateRequest(ctx context.Context, options *DomainsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.EventGrid/domains"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-15-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *DomainsClient) listBySubscriptionHandleResponse(resp *http.Response) (DomainsClientListBySubscriptionResponse, error) {
	result := DomainsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DomainsListResult); err != nil {
		return DomainsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// ListSharedAccessKeys - List the two keys used to publish to a domain.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-10-15-preview
// resourceGroupName - The name of the resource group within the user's subscription.
// domainName - Name of the domain.
// options - DomainsClientListSharedAccessKeysOptions contains the optional parameters for the DomainsClient.ListSharedAccessKeys
// method.
func (client *DomainsClient) ListSharedAccessKeys(ctx context.Context, resourceGroupName string, domainName string, options *DomainsClientListSharedAccessKeysOptions) (DomainsClientListSharedAccessKeysResponse, error) {
	req, err := client.listSharedAccessKeysCreateRequest(ctx, resourceGroupName, domainName, options)
	if err != nil {
		return DomainsClientListSharedAccessKeysResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DomainsClientListSharedAccessKeysResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DomainsClientListSharedAccessKeysResponse{}, runtime.NewResponseError(resp)
	}
	return client.listSharedAccessKeysHandleResponse(resp)
}

// listSharedAccessKeysCreateRequest creates the ListSharedAccessKeys request.
func (client *DomainsClient) listSharedAccessKeysCreateRequest(ctx context.Context, resourceGroupName string, domainName string, options *DomainsClientListSharedAccessKeysOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/domains/{domainName}/listKeys"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if domainName == "" {
		return nil, errors.New("parameter domainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainName}", url.PathEscape(domainName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listSharedAccessKeysHandleResponse handles the ListSharedAccessKeys response.
func (client *DomainsClient) listSharedAccessKeysHandleResponse(resp *http.Response) (DomainsClientListSharedAccessKeysResponse, error) {
	result := DomainsClientListSharedAccessKeysResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DomainSharedAccessKeys); err != nil {
		return DomainsClientListSharedAccessKeysResponse{}, err
	}
	return result, nil
}

// RegenerateKey - Regenerate a shared access key for a domain.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-10-15-preview
// resourceGroupName - The name of the resource group within the user's subscription.
// domainName - Name of the domain.
// regenerateKeyRequest - Request body to regenerate key.
// options - DomainsClientRegenerateKeyOptions contains the optional parameters for the DomainsClient.RegenerateKey method.
func (client *DomainsClient) RegenerateKey(ctx context.Context, resourceGroupName string, domainName string, regenerateKeyRequest DomainRegenerateKeyRequest, options *DomainsClientRegenerateKeyOptions) (DomainsClientRegenerateKeyResponse, error) {
	req, err := client.regenerateKeyCreateRequest(ctx, resourceGroupName, domainName, regenerateKeyRequest, options)
	if err != nil {
		return DomainsClientRegenerateKeyResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DomainsClientRegenerateKeyResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DomainsClientRegenerateKeyResponse{}, runtime.NewResponseError(resp)
	}
	return client.regenerateKeyHandleResponse(resp)
}

// regenerateKeyCreateRequest creates the RegenerateKey request.
func (client *DomainsClient) regenerateKeyCreateRequest(ctx context.Context, resourceGroupName string, domainName string, regenerateKeyRequest DomainRegenerateKeyRequest, options *DomainsClientRegenerateKeyOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/domains/{domainName}/regenerateKey"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if domainName == "" {
		return nil, errors.New("parameter domainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainName}", url.PathEscape(domainName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, regenerateKeyRequest)
}

// regenerateKeyHandleResponse handles the RegenerateKey response.
func (client *DomainsClient) regenerateKeyHandleResponse(resp *http.Response) (DomainsClientRegenerateKeyResponse, error) {
	result := DomainsClientRegenerateKeyResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DomainSharedAccessKeys); err != nil {
		return DomainsClientRegenerateKeyResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Asynchronously updates a domain with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-10-15-preview
// resourceGroupName - The name of the resource group within the user's subscription.
// domainName - Name of the domain.
// domainUpdateParameters - Domain update information.
// options - DomainsClientBeginUpdateOptions contains the optional parameters for the DomainsClient.BeginUpdate method.
func (client *DomainsClient) BeginUpdate(ctx context.Context, resourceGroupName string, domainName string, domainUpdateParameters DomainUpdateParameters, options *DomainsClientBeginUpdateOptions) (*runtime.Poller[DomainsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, domainName, domainUpdateParameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[DomainsClientUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[DomainsClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Asynchronously updates a domain with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-10-15-preview
func (client *DomainsClient) update(ctx context.Context, resourceGroupName string, domainName string, domainUpdateParameters DomainUpdateParameters, options *DomainsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, domainName, domainUpdateParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *DomainsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, domainName string, domainUpdateParameters DomainUpdateParameters, options *DomainsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/domains/{domainName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if domainName == "" {
		return nil, errors.New("parameter domainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainName}", url.PathEscape(domainName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, domainUpdateParameters)
}
