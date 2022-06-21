//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcommunication

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

// DomainsClient contains the methods for the Domains group.
// Don't use this type directly, use NewDomainsClient() instead.
type DomainsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewDomainsClient creates a new instance of DomainsClient with the specified values.
// subscriptionID - The ID of the target subscription.
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

// BeginCancelVerification - Cancel verification of DNS record.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-10-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// emailServiceName - The name of the EmailService resource.
// domainName - The name of the Domains resource.
// parameters - Type of verification to be canceled.
// options - DomainsClientBeginCancelVerificationOptions contains the optional parameters for the DomainsClient.BeginCancelVerification
// method.
func (client *DomainsClient) BeginCancelVerification(ctx context.Context, resourceGroupName string, emailServiceName string, domainName string, parameters VerificationParameter, options *DomainsClientBeginCancelVerificationOptions) (*runtime.Poller[DomainsClientCancelVerificationResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.cancelVerification(ctx, resourceGroupName, emailServiceName, domainName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[DomainsClientCancelVerificationResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[DomainsClientCancelVerificationResponse](options.ResumeToken, client.pl, nil)
	}
}

// CancelVerification - Cancel verification of DNS record.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-10-01-preview
func (client *DomainsClient) cancelVerification(ctx context.Context, resourceGroupName string, emailServiceName string, domainName string, parameters VerificationParameter, options *DomainsClientBeginCancelVerificationOptions) (*http.Response, error) {
	req, err := client.cancelVerificationCreateRequest(ctx, resourceGroupName, emailServiceName, domainName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// cancelVerificationCreateRequest creates the CancelVerification request.
func (client *DomainsClient) cancelVerificationCreateRequest(ctx context.Context, resourceGroupName string, emailServiceName string, domainName string, parameters VerificationParameter, options *DomainsClientBeginCancelVerificationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Communication/emailServices/{emailServiceName}/domains/{domainName}/cancelVerification"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if emailServiceName == "" {
		return nil, errors.New("parameter emailServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{emailServiceName}", url.PathEscape(emailServiceName))
	if domainName == "" {
		return nil, errors.New("parameter domainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainName}", url.PathEscape(domainName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginCreateOrUpdate - Add a new Domains resource under the parent EmailService resource or update an existing Domains resource.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-10-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// emailServiceName - The name of the EmailService resource.
// domainName - The name of the Domains resource.
// parameters - Parameters for the create or update operation
// options - DomainsClientBeginCreateOrUpdateOptions contains the optional parameters for the DomainsClient.BeginCreateOrUpdate
// method.
func (client *DomainsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, emailServiceName string, domainName string, parameters DomainResource, options *DomainsClientBeginCreateOrUpdateOptions) (*runtime.Poller[DomainsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, emailServiceName, domainName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[DomainsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[DomainsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Add a new Domains resource under the parent EmailService resource or update an existing Domains resource.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-10-01-preview
func (client *DomainsClient) createOrUpdate(ctx context.Context, resourceGroupName string, emailServiceName string, domainName string, parameters DomainResource, options *DomainsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, emailServiceName, domainName, parameters, options)
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

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DomainsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, emailServiceName string, domainName string, parameters DomainResource, options *DomainsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Communication/emailServices/{emailServiceName}/domains/{domainName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if emailServiceName == "" {
		return nil, errors.New("parameter emailServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{emailServiceName}", url.PathEscape(emailServiceName))
	if domainName == "" {
		return nil, errors.New("parameter domainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainName}", url.PathEscape(domainName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Operation to delete a Domains resource.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-10-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// emailServiceName - The name of the EmailService resource.
// domainName - The name of the Domains resource.
// options - DomainsClientBeginDeleteOptions contains the optional parameters for the DomainsClient.BeginDelete method.
func (client *DomainsClient) BeginDelete(ctx context.Context, resourceGroupName string, emailServiceName string, domainName string, options *DomainsClientBeginDeleteOptions) (*runtime.Poller[DomainsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, emailServiceName, domainName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[DomainsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[DomainsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Operation to delete a Domains resource.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-10-01-preview
func (client *DomainsClient) deleteOperation(ctx context.Context, resourceGroupName string, emailServiceName string, domainName string, options *DomainsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, emailServiceName, domainName, options)
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
func (client *DomainsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, emailServiceName string, domainName string, options *DomainsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Communication/emailServices/{emailServiceName}/domains/{domainName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if emailServiceName == "" {
		return nil, errors.New("parameter emailServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{emailServiceName}", url.PathEscape(emailServiceName))
	if domainName == "" {
		return nil, errors.New("parameter domainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainName}", url.PathEscape(domainName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get the Domains resource and its properties.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-10-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// emailServiceName - The name of the EmailService resource.
// domainName - The name of the Domains resource.
// options - DomainsClientGetOptions contains the optional parameters for the DomainsClient.Get method.
func (client *DomainsClient) Get(ctx context.Context, resourceGroupName string, emailServiceName string, domainName string, options *DomainsClientGetOptions) (DomainsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, emailServiceName, domainName, options)
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
func (client *DomainsClient) getCreateRequest(ctx context.Context, resourceGroupName string, emailServiceName string, domainName string, options *DomainsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Communication/emailServices/{emailServiceName}/domains/{domainName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if emailServiceName == "" {
		return nil, errors.New("parameter emailServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{emailServiceName}", url.PathEscape(emailServiceName))
	if domainName == "" {
		return nil, errors.New("parameter domainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainName}", url.PathEscape(domainName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DomainsClient) getHandleResponse(resp *http.Response) (DomainsClientGetResponse, error) {
	result := DomainsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DomainResource); err != nil {
		return DomainsClientGetResponse{}, err
	}
	return result, nil
}

// BeginInitiateVerification - Initiate verification of DNS record.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-10-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// emailServiceName - The name of the EmailService resource.
// domainName - The name of the Domains resource.
// parameters - Type of verification to be initiated.
// options - DomainsClientBeginInitiateVerificationOptions contains the optional parameters for the DomainsClient.BeginInitiateVerification
// method.
func (client *DomainsClient) BeginInitiateVerification(ctx context.Context, resourceGroupName string, emailServiceName string, domainName string, parameters VerificationParameter, options *DomainsClientBeginInitiateVerificationOptions) (*runtime.Poller[DomainsClientInitiateVerificationResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.initiateVerification(ctx, resourceGroupName, emailServiceName, domainName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[DomainsClientInitiateVerificationResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[DomainsClientInitiateVerificationResponse](options.ResumeToken, client.pl, nil)
	}
}

// InitiateVerification - Initiate verification of DNS record.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-10-01-preview
func (client *DomainsClient) initiateVerification(ctx context.Context, resourceGroupName string, emailServiceName string, domainName string, parameters VerificationParameter, options *DomainsClientBeginInitiateVerificationOptions) (*http.Response, error) {
	req, err := client.initiateVerificationCreateRequest(ctx, resourceGroupName, emailServiceName, domainName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// initiateVerificationCreateRequest creates the InitiateVerification request.
func (client *DomainsClient) initiateVerificationCreateRequest(ctx context.Context, resourceGroupName string, emailServiceName string, domainName string, parameters VerificationParameter, options *DomainsClientBeginInitiateVerificationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Communication/emailServices/{emailServiceName}/domains/{domainName}/initiateVerification"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if emailServiceName == "" {
		return nil, errors.New("parameter emailServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{emailServiceName}", url.PathEscape(emailServiceName))
	if domainName == "" {
		return nil, errors.New("parameter domainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainName}", url.PathEscape(domainName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// NewListByEmailServiceResourcePager - Handles requests to list all Domains resources under the parent EmailServices resource.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-10-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// emailServiceName - The name of the EmailService resource.
// options - DomainsClientListByEmailServiceResourceOptions contains the optional parameters for the DomainsClient.ListByEmailServiceResource
// method.
func (client *DomainsClient) NewListByEmailServiceResourcePager(resourceGroupName string, emailServiceName string, options *DomainsClientListByEmailServiceResourceOptions) *runtime.Pager[DomainsClientListByEmailServiceResourceResponse] {
	return runtime.NewPager(runtime.PagingHandler[DomainsClientListByEmailServiceResourceResponse]{
		More: func(page DomainsClientListByEmailServiceResourceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DomainsClientListByEmailServiceResourceResponse) (DomainsClientListByEmailServiceResourceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByEmailServiceResourceCreateRequest(ctx, resourceGroupName, emailServiceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DomainsClientListByEmailServiceResourceResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DomainsClientListByEmailServiceResourceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DomainsClientListByEmailServiceResourceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByEmailServiceResourceHandleResponse(resp)
		},
	})
}

// listByEmailServiceResourceCreateRequest creates the ListByEmailServiceResource request.
func (client *DomainsClient) listByEmailServiceResourceCreateRequest(ctx context.Context, resourceGroupName string, emailServiceName string, options *DomainsClientListByEmailServiceResourceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Communication/emailServices/{emailServiceName}/domains"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if emailServiceName == "" {
		return nil, errors.New("parameter emailServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{emailServiceName}", url.PathEscape(emailServiceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByEmailServiceResourceHandleResponse handles the ListByEmailServiceResource response.
func (client *DomainsClient) listByEmailServiceResourceHandleResponse(resp *http.Response) (DomainsClientListByEmailServiceResourceResponse, error) {
	result := DomainsClientListByEmailServiceResourceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DomainResourceList); err != nil {
		return DomainsClientListByEmailServiceResourceResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Operation to update an existing Domains resource.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-10-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// emailServiceName - The name of the EmailService resource.
// domainName - The name of the Domains resource.
// parameters - Parameters for the update operation
// options - DomainsClientBeginUpdateOptions contains the optional parameters for the DomainsClient.BeginUpdate method.
func (client *DomainsClient) BeginUpdate(ctx context.Context, resourceGroupName string, emailServiceName string, domainName string, parameters UpdateDomainRequestParameters, options *DomainsClientBeginUpdateOptions) (*runtime.Poller[DomainsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, emailServiceName, domainName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[DomainsClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[DomainsClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Operation to update an existing Domains resource.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-10-01-preview
func (client *DomainsClient) update(ctx context.Context, resourceGroupName string, emailServiceName string, domainName string, parameters UpdateDomainRequestParameters, options *DomainsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, emailServiceName, domainName, parameters, options)
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
func (client *DomainsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, emailServiceName string, domainName string, parameters UpdateDomainRequestParameters, options *DomainsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Communication/emailServices/{emailServiceName}/domains/{domainName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if emailServiceName == "" {
		return nil, errors.New("parameter emailServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{emailServiceName}", url.PathEscape(emailServiceName))
	if domainName == "" {
		return nil, errors.New("parameter domainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainName}", url.PathEscape(domainName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}
