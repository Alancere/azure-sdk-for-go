//go:build go1.16
// +build go1.16

// Code generated by Microsoft (R) AutoRest Code Generator (autorest: 3.4.3, generator: @autorest/go@4.0.0-preview.27)
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmysql

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// connectionOptions contains configuration settings for the connection's pipeline.
// All zero-value fields will be initialized with their default values.
type connectionOptions struct {
	// HTTPClient sets the transport for making HTTP requests.
	HTTPClient policy.Transporter
	// Retry configures the built-in retry policy behavior.
	Retry policy.RetryOptions
	// Telemetry configures the built-in telemetry policy behavior.
	Telemetry policy.TelemetryOptions
	// Logging configures the built-in logging policy behavior.
	Logging policy.LogOptions
	// PerCallPolicies contains custom policies to inject into the pipeline.
	// Each policy is executed once per request.
	PerCallPolicies []policy.Policy
	// PerRetryPolicies contains custom policies to inject into the pipeline.
	// Each policy is executed once per request, and for each retry request.
	PerRetryPolicies []policy.Policy
}

// connection - The Microsoft Azure management API provides create, read, update, and delete functionality for Azure MySQL resources including servers, databases, firewall rules, VNET rules, log files and configurations with new business model.
type connection struct {
	u string
	p runtime.Pipeline
}

// defaultEndpoint is the default service endpoint.
const defaultEndpoint = "https://management.azure.com"

// newDefaultConnection creates an instance of the connection type using the defaultEndpoint.
// Pass nil to accept the default options; this is the same as passing a zero-value options.
func newDefaultConnection(cred azcore.Credential, options *connectionOptions) *connection {
	return newConnection(defaultEndpoint, cred, options)
}

// newConnection creates an instance of the connection type with the specified endpoint.
// Pass nil to accept the default options; this is the same as passing a zero-value options.
func newConnection(endpoint string, cred azcore.Credential, options *connectionOptions) *connection {
	if options == nil {
		options = &connectionOptions{}
	}
	policies := []policy.Policy{}
	if !options.Telemetry.Disabled {
		policies = append(policies, runtime.NewTelemetryPolicy(module, version, &options.Telemetry))
	}
	policies = append(policies, options.PerCallPolicies...)
	policies = append(policies, runtime.NewRetryPolicy(&options.Retry))
	policies = append(policies, options.PerRetryPolicies...)
	policies = append(policies, cred.NewAuthenticationPolicy(runtime.AuthenticationOptions{TokenRequest: policy.TokenRequestOptions{}}))
	policies = append(policies, runtime.NewLogPolicy(&options.Logging))
	return &connection{u: endpoint, p: runtime.NewPipeline(options.HTTPClient, policies...)}
}

// Endpoint returns the connection's endpoint.
func (c *connection) Endpoint() string {
	return c.u
}

// Pipeline returns the connection's pipeline.
func (c *connection) Pipeline() runtime.Pipeline {
	return c.p
}
