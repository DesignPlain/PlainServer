package apigatewayv2

import types "DesignSphere_Server/src/resource/aws/types"

type DomainName struct {
	// Map of tags to assign to the domain name. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Domain name. Must be between 1 and 512 characters in length.
	DomainName string `json:"domainName,omitempty" yaml:"domainName,omitempty"`

	// Domain name configuration. See below.
	DomainNameConfiguration types.Apigatewayv2_DomainNameDomainNameConfiguration `json:"domainNameConfiguration,omitempty" yaml:"domainNameConfiguration,omitempty"`

	// Mutual TLS authentication configuration for the domain name.
	MutualTlsAuthentication types.Apigatewayv2_DomainNameMutualTlsAuthentication `json:"mutualTlsAuthentication,omitempty" yaml:"mutualTlsAuthentication,omitempty"`
}
