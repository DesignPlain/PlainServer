package finspace

import types "libds/aws/types"

type KxEnvironment struct {
	// List of DNS server name and server IP. This is used to set up Route-53 outbound resolvers. Defined below.
	CustomDnsConfigurations []types.Finspace_KxEnvironmentCustomDnsConfiguration `json:"customDnsConfigurations,omitempty" yaml:"customDnsConfigurations,omitempty"`

	// Description for the KX environment.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   KMS key ID to encrypt your data in the FinSpace environment.

	   The following arguments are optional:
	*/
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	// Name of the KX environment that you want to create.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Transit gateway and network configuration that is used to connect the KX environment to an internal network. Defined below.
	TransitGatewayConfiguration types.Finspace_KxEnvironmentTransitGatewayConfiguration `json:"transitGatewayConfiguration,omitempty" yaml:"transitGatewayConfiguration,omitempty"`
}
