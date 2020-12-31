package mixedreality

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/preview/mixedreality/mgmt/2020-05-01-preview/mixedreality"

// AccountKeyRegenerateRequest request for account key regeneration
type AccountKeyRegenerateRequest struct {
	// Serial - serial of key to be regenerated
	Serial *int32 `json:"serial,omitempty"`
}

// AccountKeys developer Keys of account
type AccountKeys struct {
	autorest.Response `json:"-"`
	// PrimaryKey - READ-ONLY; value of primary key.
	PrimaryKey *string `json:"primaryKey,omitempty"`
	// SecondaryKey - READ-ONLY; value of secondary key.
	SecondaryKey *string `json:"secondaryKey,omitempty"`
}

// AccountProperties common Properties shared by Mixed Reality Accounts
type AccountProperties struct {
	// AccountID - READ-ONLY; unique id of certain account.
	AccountID *string `json:"accountId,omitempty"`
	// AccountDomain - READ-ONLY; Correspond domain name of certain Spatial Anchors Account
	AccountDomain *string `json:"accountDomain,omitempty"`
}

// AzureEntityResource the resource model definition for an Azure Resource Manager resource with an etag.
type AzureEntityResource struct {
	// Etag - READ-ONLY; Resource Etag.
	Etag *string `json:"etag,omitempty"`
	// ID - READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// CheckNameAvailabilityRequest check Name Availability Request
type CheckNameAvailabilityRequest struct {
	// Name - Resource Name To Verify
	Name *string `json:"name,omitempty"`
	// Type - Fully qualified resource type which includes provider namespace
	Type *string `json:"type,omitempty"`
}

// CheckNameAvailabilityResponse check Name Availability Response
type CheckNameAvailabilityResponse struct {
	autorest.Response `json:"-"`
	// NameAvailable - if name Available. Possible values include: 'True', 'False'
	NameAvailable NameAvailability `json:"nameAvailable,omitempty"`
	// Reason - Resource Name To Verify. Possible values include: 'Invalid', 'AlreadyExists'
	Reason NameUnavailableReason `json:"reason,omitempty"`
	// Message - detail message
	Message *string `json:"message,omitempty"`
}

// CloudError an Error response.
type CloudError struct {
	Error *CloudErrorBody `json:"error,omitempty"`
}

// CloudErrorBody an error response from Azure.
type CloudErrorBody struct {
	// Code - An identifier for the error. Codes are invariant and are intended to be consumed programmatically.
	Code *string `json:"code,omitempty"`
	// Message - A message describing the error, intended to be suitable for displaying in a user interface.
	Message *string `json:"message,omitempty"`
	// Target - The target of the particular error. For example, the name of the property in error.
	Target *string `json:"target,omitempty"`
	// Details - A list of additional details about the error.
	Details *[]CloudErrorBody `json:"details,omitempty"`
}

// Identity identity for the resource.
type Identity struct {
	// PrincipalID - READ-ONLY; The principal ID of resource identity.
	PrincipalID *string `json:"principalId,omitempty"`
	// TenantID - READ-ONLY; The tenant ID of resource.
	TenantID *string `json:"tenantId,omitempty"`
	// Type - The identity type. Possible values include: 'SystemAssigned'
	Type ResourceIdentityType `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for Identity.
func (i Identity) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if i.Type != "" {
		objectMap["type"] = i.Type
	}
	return json.Marshal(objectMap)
}

// Operation REST API operation
type Operation struct {
	// Name - Operation name: {provider}/{resource}/{operation}
	Name *string `json:"name,omitempty"`
	// Display - The object that represents the operation.
	Display *OperationDisplay `json:"display,omitempty"`
}

// OperationDisplay the object that represents the operation.
type OperationDisplay struct {
	// Provider - Service provider: Microsoft.ResourceProvider
	Provider *string `json:"provider,omitempty"`
	// Resource - Resource on which the operation is performed: Profile, endpoint, etc.
	Resource *string `json:"resource,omitempty"`
	// Operation - Operation type: Read, write, delete, etc.
	Operation *string `json:"operation,omitempty"`
	// Description - Description of operation
	Description *string `json:"description,omitempty"`
}

// OperationPage result of the request to list Resource Provider operations. It contains a list of
// operations and a URL link to get the next set of results.
type OperationPage struct {
	autorest.Response `json:"-"`
	// Value - List of operations supported by the Resource Provider.
	Value *[]Operation `json:"value,omitempty"`
	// NextLink - URL to get the next set of operation list results if there are any.
	NextLink *string `json:"nextLink,omitempty"`
}

// OperationPageIterator provides access to a complete listing of Operation values.
type OperationPageIterator struct {
	i    int
	page OperationPagePage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *OperationPageIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationPageIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *OperationPageIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter OperationPageIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter OperationPageIterator) Response() OperationPage {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter OperationPageIterator) Value() Operation {
	if !iter.page.NotDone() {
		return Operation{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the OperationPageIterator type.
func NewOperationPageIterator(page OperationPagePage) OperationPageIterator {
	return OperationPageIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (op OperationPage) IsEmpty() bool {
	return op.Value == nil || len(*op.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (op OperationPage) hasNextLink() bool {
	return op.NextLink != nil && len(*op.NextLink) != 0
}

// operationPagePreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (op OperationPage) operationPagePreparer(ctx context.Context) (*http.Request, error) {
	if !op.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(op.NextLink)))
}

// OperationPagePage contains a page of Operation values.
type OperationPagePage struct {
	fn func(context.Context, OperationPage) (OperationPage, error)
	op OperationPage
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *OperationPagePage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationPagePage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.op)
		if err != nil {
			return err
		}
		page.op = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *OperationPagePage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page OperationPagePage) NotDone() bool {
	return !page.op.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page OperationPagePage) Response() OperationPage {
	return page.op
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page OperationPagePage) Values() []Operation {
	if page.op.IsEmpty() {
		return nil
	}
	return *page.op.Value
}

// Creates a new instance of the OperationPagePage type.
func NewOperationPagePage(cur OperationPage, getNextPage func(context.Context, OperationPage) (OperationPage, error)) OperationPagePage {
	return OperationPagePage{
		fn: getNextPage,
		op: cur,
	}
}

// Plan plan for the resource.
type Plan struct {
	// Name - A user defined name of the 3rd Party Artifact that is being procured.
	Name *string `json:"name,omitempty"`
	// Publisher - The publisher of the 3rd Party Artifact that is being bought. E.g. NewRelic
	Publisher *string `json:"publisher,omitempty"`
	// Product - The 3rd Party artifact that is being procured. E.g. NewRelic. Product maps to the OfferID specified for the artifact at the time of Data Market onboarding.
	Product *string `json:"product,omitempty"`
	// PromotionCode - A publisher provided promotion code as provisioned in Data Market for the said product/artifact.
	PromotionCode *string `json:"promotionCode,omitempty"`
	// Version - The version of the desired product/artifact.
	Version *string `json:"version,omitempty"`
}

// ProxyResource the resource model definition for a Azure Resource Manager proxy resource. It will not
// have tags and a location
type ProxyResource struct {
	// ID - READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// RemoteRenderingAccount remoteRenderingAccount Response.
type RemoteRenderingAccount struct {
	autorest.Response `json:"-"`
	Identity          *RemoteRenderingAccountIdentity `json:"identity,omitempty"`
	// AccountProperties - Property bag.
	*AccountProperties `json:"properties,omitempty"`
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
	// Location - The geo-location where the resource lives
	Location *string `json:"location,omitempty"`
	// ID - READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for RemoteRenderingAccount.
func (rra RemoteRenderingAccount) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if rra.Identity != nil {
		objectMap["identity"] = rra.Identity
	}
	if rra.AccountProperties != nil {
		objectMap["properties"] = rra.AccountProperties
	}
	if rra.Tags != nil {
		objectMap["tags"] = rra.Tags
	}
	if rra.Location != nil {
		objectMap["location"] = rra.Location
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for RemoteRenderingAccount struct.
func (rra *RemoteRenderingAccount) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "identity":
			if v != nil {
				var identity RemoteRenderingAccountIdentity
				err = json.Unmarshal(*v, &identity)
				if err != nil {
					return err
				}
				rra.Identity = &identity
			}
		case "properties":
			if v != nil {
				var accountProperties AccountProperties
				err = json.Unmarshal(*v, &accountProperties)
				if err != nil {
					return err
				}
				rra.AccountProperties = &accountProperties
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				rra.Tags = tags
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				rra.Location = &location
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				rra.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				rra.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				rra.Type = &typeVar
			}
		}
	}

	return nil
}

// RemoteRenderingAccountIdentity ...
type RemoteRenderingAccountIdentity struct {
	// PrincipalID - READ-ONLY; The principal ID of resource identity.
	PrincipalID *string `json:"principalId,omitempty"`
	// TenantID - READ-ONLY; The tenant ID of resource.
	TenantID *string `json:"tenantId,omitempty"`
	// Type - The identity type. Possible values include: 'SystemAssigned'
	Type ResourceIdentityType `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for RemoteRenderingAccountIdentity.
func (rra RemoteRenderingAccountIdentity) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if rra.Type != "" {
		objectMap["type"] = rra.Type
	}
	return json.Marshal(objectMap)
}

// RemoteRenderingAccountPage result of the request to get resource collection. It contains a list of
// resources and a URL link to get the next set of results.
type RemoteRenderingAccountPage struct {
	autorest.Response `json:"-"`
	// Value - List of resources supported by the Resource Provider.
	Value *[]RemoteRenderingAccount `json:"value,omitempty"`
	// NextLink - URL to get the next set of resource list results if there are any.
	NextLink *string `json:"nextLink,omitempty"`
}

// RemoteRenderingAccountPageIterator provides access to a complete listing of RemoteRenderingAccount
// values.
type RemoteRenderingAccountPageIterator struct {
	i    int
	page RemoteRenderingAccountPagePage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *RemoteRenderingAccountPageIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RemoteRenderingAccountPageIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *RemoteRenderingAccountPageIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter RemoteRenderingAccountPageIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter RemoteRenderingAccountPageIterator) Response() RemoteRenderingAccountPage {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter RemoteRenderingAccountPageIterator) Value() RemoteRenderingAccount {
	if !iter.page.NotDone() {
		return RemoteRenderingAccount{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the RemoteRenderingAccountPageIterator type.
func NewRemoteRenderingAccountPageIterator(page RemoteRenderingAccountPagePage) RemoteRenderingAccountPageIterator {
	return RemoteRenderingAccountPageIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (rrap RemoteRenderingAccountPage) IsEmpty() bool {
	return rrap.Value == nil || len(*rrap.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (rrap RemoteRenderingAccountPage) hasNextLink() bool {
	return rrap.NextLink != nil && len(*rrap.NextLink) != 0
}

// remoteRenderingAccountPagePreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (rrap RemoteRenderingAccountPage) remoteRenderingAccountPagePreparer(ctx context.Context) (*http.Request, error) {
	if !rrap.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(rrap.NextLink)))
}

// RemoteRenderingAccountPagePage contains a page of RemoteRenderingAccount values.
type RemoteRenderingAccountPagePage struct {
	fn   func(context.Context, RemoteRenderingAccountPage) (RemoteRenderingAccountPage, error)
	rrap RemoteRenderingAccountPage
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *RemoteRenderingAccountPagePage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RemoteRenderingAccountPagePage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.rrap)
		if err != nil {
			return err
		}
		page.rrap = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *RemoteRenderingAccountPagePage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page RemoteRenderingAccountPagePage) NotDone() bool {
	return !page.rrap.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page RemoteRenderingAccountPagePage) Response() RemoteRenderingAccountPage {
	return page.rrap
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page RemoteRenderingAccountPagePage) Values() []RemoteRenderingAccount {
	if page.rrap.IsEmpty() {
		return nil
	}
	return *page.rrap.Value
}

// Creates a new instance of the RemoteRenderingAccountPagePage type.
func NewRemoteRenderingAccountPagePage(cur RemoteRenderingAccountPage, getNextPage func(context.Context, RemoteRenderingAccountPage) (RemoteRenderingAccountPage, error)) RemoteRenderingAccountPagePage {
	return RemoteRenderingAccountPagePage{
		fn:   getNextPage,
		rrap: cur,
	}
}

// Resource common fields that are returned in the response for all Azure Resource Manager resources
type Resource struct {
	// ID - READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// ResourceModelWithAllowedPropertySet the resource model definition containing the full set of allowed
// properties for a resource. Except properties bag, there cannot be a top level property outside of this
// set.
type ResourceModelWithAllowedPropertySet struct {
	// ID - READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
	// Location - The geo-location where the resource lives
	Location *string `json:"location,omitempty"`
	// ManagedBy - The fully qualified resource ID of the resource that manages this resource. Indicates if this resource is managed by another Azure resource. If this is present, complete mode deployment will not delete the resource if it is removed from the template since it is managed by another resource.
	ManagedBy *string `json:"managedBy,omitempty"`
	// Kind - Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type; e.g. ApiApps are a kind of Microsoft.Web/sites type.  If supported, the resource provider must validate and persist this value.
	Kind *string `json:"kind,omitempty"`
	// Etag - READ-ONLY; The etag field is *not* required. If it is provided in the response body, it must also be provided as a header per the normal etag convention.  Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields.
	Etag *string `json:"etag,omitempty"`
	// Tags - Resource tags.
	Tags     map[string]*string                           `json:"tags"`
	Identity *ResourceModelWithAllowedPropertySetIdentity `json:"identity,omitempty"`
	Sku      *ResourceModelWithAllowedPropertySetSku      `json:"sku,omitempty"`
	Plan     *ResourceModelWithAllowedPropertySetPlan     `json:"plan,omitempty"`
}

// MarshalJSON is the custom marshaler for ResourceModelWithAllowedPropertySet.
func (rmwaps ResourceModelWithAllowedPropertySet) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if rmwaps.Location != nil {
		objectMap["location"] = rmwaps.Location
	}
	if rmwaps.ManagedBy != nil {
		objectMap["managedBy"] = rmwaps.ManagedBy
	}
	if rmwaps.Kind != nil {
		objectMap["kind"] = rmwaps.Kind
	}
	if rmwaps.Tags != nil {
		objectMap["tags"] = rmwaps.Tags
	}
	if rmwaps.Identity != nil {
		objectMap["identity"] = rmwaps.Identity
	}
	if rmwaps.Sku != nil {
		objectMap["sku"] = rmwaps.Sku
	}
	if rmwaps.Plan != nil {
		objectMap["plan"] = rmwaps.Plan
	}
	return json.Marshal(objectMap)
}

// ResourceModelWithAllowedPropertySetIdentity ...
type ResourceModelWithAllowedPropertySetIdentity struct {
	// PrincipalID - READ-ONLY; The principal ID of resource identity.
	PrincipalID *string `json:"principalId,omitempty"`
	// TenantID - READ-ONLY; The tenant ID of resource.
	TenantID *string `json:"tenantId,omitempty"`
	// Type - The identity type. Possible values include: 'SystemAssigned'
	Type ResourceIdentityType `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for ResourceModelWithAllowedPropertySetIdentity.
func (rmwaps ResourceModelWithAllowedPropertySetIdentity) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if rmwaps.Type != "" {
		objectMap["type"] = rmwaps.Type
	}
	return json.Marshal(objectMap)
}

// ResourceModelWithAllowedPropertySetPlan ...
type ResourceModelWithAllowedPropertySetPlan struct {
	// Name - A user defined name of the 3rd Party Artifact that is being procured.
	Name *string `json:"name,omitempty"`
	// Publisher - The publisher of the 3rd Party Artifact that is being bought. E.g. NewRelic
	Publisher *string `json:"publisher,omitempty"`
	// Product - The 3rd Party artifact that is being procured. E.g. NewRelic. Product maps to the OfferID specified for the artifact at the time of Data Market onboarding.
	Product *string `json:"product,omitempty"`
	// PromotionCode - A publisher provided promotion code as provisioned in Data Market for the said product/artifact.
	PromotionCode *string `json:"promotionCode,omitempty"`
	// Version - The version of the desired product/artifact.
	Version *string `json:"version,omitempty"`
}

// ResourceModelWithAllowedPropertySetSku ...
type ResourceModelWithAllowedPropertySetSku struct {
	// Name - The name of the SKU. Ex - P3. It is typically a letter+number code
	Name *string `json:"name,omitempty"`
	// Tier - This field is required to be implemented by the Resource Provider if the service has more than one tier, but is not required on a PUT. Possible values include: 'Free', 'Basic', 'Standard', 'Premium'
	Tier SkuTier `json:"tier,omitempty"`
	// Size - The SKU size. When the name field is the combination of tier and some other value, this would be the standalone code.
	Size *string `json:"size,omitempty"`
	// Family - If the service has different generations of hardware, for the same SKU, then that can be captured here.
	Family *string `json:"family,omitempty"`
	// Capacity - If the SKU supports scale out/in then the capacity integer should be included. If scale out/in is not possible for the resource this may be omitted.
	Capacity *int32 `json:"capacity,omitempty"`
}

// Sku the resource model definition representing SKU
type Sku struct {
	// Name - The name of the SKU. Ex - P3. It is typically a letter+number code
	Name *string `json:"name,omitempty"`
	// Tier - This field is required to be implemented by the Resource Provider if the service has more than one tier, but is not required on a PUT. Possible values include: 'Free', 'Basic', 'Standard', 'Premium'
	Tier SkuTier `json:"tier,omitempty"`
	// Size - The SKU size. When the name field is the combination of tier and some other value, this would be the standalone code.
	Size *string `json:"size,omitempty"`
	// Family - If the service has different generations of hardware, for the same SKU, then that can be captured here.
	Family *string `json:"family,omitempty"`
	// Capacity - If the SKU supports scale out/in then the capacity integer should be included. If scale out/in is not possible for the resource this may be omitted.
	Capacity *int32 `json:"capacity,omitempty"`
}

// SpatialAnchorsAccount spatialAnchorsAccount Response.
type SpatialAnchorsAccount struct {
	autorest.Response `json:"-"`
	// AccountProperties - Property bag.
	*AccountProperties `json:"properties,omitempty"`
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
	// Location - The geo-location where the resource lives
	Location *string `json:"location,omitempty"`
	// ID - READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for SpatialAnchorsAccount.
func (saa SpatialAnchorsAccount) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if saa.AccountProperties != nil {
		objectMap["properties"] = saa.AccountProperties
	}
	if saa.Tags != nil {
		objectMap["tags"] = saa.Tags
	}
	if saa.Location != nil {
		objectMap["location"] = saa.Location
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for SpatialAnchorsAccount struct.
func (saa *SpatialAnchorsAccount) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var accountProperties AccountProperties
				err = json.Unmarshal(*v, &accountProperties)
				if err != nil {
					return err
				}
				saa.AccountProperties = &accountProperties
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				saa.Tags = tags
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				saa.Location = &location
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				saa.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				saa.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				saa.Type = &typeVar
			}
		}
	}

	return nil
}

// SpatialAnchorsAccountPage result of the request to get resource collection. It contains a list of
// resources and a URL link to get the next set of results.
type SpatialAnchorsAccountPage struct {
	autorest.Response `json:"-"`
	// Value - List of resources supported by the Resource Provider.
	Value *[]SpatialAnchorsAccount `json:"value,omitempty"`
	// NextLink - URL to get the next set of resource list results if there are any.
	NextLink *string `json:"nextLink,omitempty"`
}

// SpatialAnchorsAccountPageIterator provides access to a complete listing of SpatialAnchorsAccount values.
type SpatialAnchorsAccountPageIterator struct {
	i    int
	page SpatialAnchorsAccountPagePage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *SpatialAnchorsAccountPageIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SpatialAnchorsAccountPageIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *SpatialAnchorsAccountPageIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter SpatialAnchorsAccountPageIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter SpatialAnchorsAccountPageIterator) Response() SpatialAnchorsAccountPage {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter SpatialAnchorsAccountPageIterator) Value() SpatialAnchorsAccount {
	if !iter.page.NotDone() {
		return SpatialAnchorsAccount{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the SpatialAnchorsAccountPageIterator type.
func NewSpatialAnchorsAccountPageIterator(page SpatialAnchorsAccountPagePage) SpatialAnchorsAccountPageIterator {
	return SpatialAnchorsAccountPageIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (saap SpatialAnchorsAccountPage) IsEmpty() bool {
	return saap.Value == nil || len(*saap.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (saap SpatialAnchorsAccountPage) hasNextLink() bool {
	return saap.NextLink != nil && len(*saap.NextLink) != 0
}

// spatialAnchorsAccountPagePreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (saap SpatialAnchorsAccountPage) spatialAnchorsAccountPagePreparer(ctx context.Context) (*http.Request, error) {
	if !saap.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(saap.NextLink)))
}

// SpatialAnchorsAccountPagePage contains a page of SpatialAnchorsAccount values.
type SpatialAnchorsAccountPagePage struct {
	fn   func(context.Context, SpatialAnchorsAccountPage) (SpatialAnchorsAccountPage, error)
	saap SpatialAnchorsAccountPage
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *SpatialAnchorsAccountPagePage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SpatialAnchorsAccountPagePage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.saap)
		if err != nil {
			return err
		}
		page.saap = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *SpatialAnchorsAccountPagePage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page SpatialAnchorsAccountPagePage) NotDone() bool {
	return !page.saap.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page SpatialAnchorsAccountPagePage) Response() SpatialAnchorsAccountPage {
	return page.saap
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page SpatialAnchorsAccountPagePage) Values() []SpatialAnchorsAccount {
	if page.saap.IsEmpty() {
		return nil
	}
	return *page.saap.Value
}

// Creates a new instance of the SpatialAnchorsAccountPagePage type.
func NewSpatialAnchorsAccountPagePage(cur SpatialAnchorsAccountPage, getNextPage func(context.Context, SpatialAnchorsAccountPage) (SpatialAnchorsAccountPage, error)) SpatialAnchorsAccountPagePage {
	return SpatialAnchorsAccountPagePage{
		fn:   getNextPage,
		saap: cur,
	}
}

// TrackedResource the resource model definition for an Azure Resource Manager tracked top level resource
// which has 'tags' and a 'location'
type TrackedResource struct {
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
	// Location - The geo-location where the resource lives
	Location *string `json:"location,omitempty"`
	// ID - READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for TrackedResource.
func (tr TrackedResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if tr.Tags != nil {
		objectMap["tags"] = tr.Tags
	}
	if tr.Location != nil {
		objectMap["location"] = tr.Location
	}
	return json.Marshal(objectMap)
}
