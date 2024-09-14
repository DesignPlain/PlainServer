package iot

import types "libds/aws/types"

type DomainConfiguration struct {
	// An object that specifies the authorization service for a domain. See the `authorizer_config` Block below for details.
	AuthorizerConfig types.Iot_DomainConfigurationAuthorizerConfig `json:"authorizerConfig,omitempty" yaml:"authorizerConfig,omitempty"`

	// Fully-qualified domain name.
	DomainName string `json:"domainName,omitempty" yaml:"domainName,omitempty"`

	// The name of the domain configuration. This value must be unique to a region.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The ARNs of the certificates that IoT passes to the device during the TLS handshake. Currently you can specify only one certificate ARN. This value is not required for Amazon Web Services-managed domains. When using a custom `domain_name`, the cert must include it.
	ServerCertificateArns []string `json:"serverCertificateArns,omitempty" yaml:"serverCertificateArns,omitempty"`

	// Map of tags to assign to this resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// An object that specifies the TLS configuration for a domain. See the `tls_config` Block below for details.
	TlsConfig types.Iot_DomainConfigurationTlsConfig `json:"tlsConfig,omitempty" yaml:"tlsConfig,omitempty"`

	// The type of service delivered by the endpoint. Note: Amazon Web Services IoT Core currently supports only the `DATA` service type.
	ServiceType string `json:"serviceType,omitempty" yaml:"serviceType,omitempty"`

	// The status to which the domain configuration should be set. Valid values are `ENABLED` and `DISABLED`.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`

	// The certificate used to validate the server certificate and prove domain name ownership. This certificate must be signed by a public certificate authority. This value is not required for Amazon Web Services-managed domains.
	ValidationCertificateArn string `json:"validationCertificateArn,omitempty" yaml:"validationCertificateArn,omitempty"`
}
