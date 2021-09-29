module github.com/Azure/azure-sdk-for-go/sdk/network/armnetwork

go 1.16

require (
	github.com/Azure/azure-sdk-for-go v57.1.0+incompatible
	github.com/Azure/azure-sdk-for-go/sdk/azcore v0.19.0
	github.com/Azure/azure-sdk-for-go/sdk/azidentity v0.10.0
)

retract [v0.1.0, v0.3.1]
