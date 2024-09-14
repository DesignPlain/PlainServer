package types

type Elasticsearch_DomainDomainEndpointOptions struct {
	// Whether to enable custom endpoint for the Elasticsearch domain.
	CustomEndpointEnabled bool `json:"customEndpointEnabled,omitempty" yaml:"customEndpointEnabled,omitempty"`

	// Whether or not to require HTTPS. Defaults to `true`.
	EnforceHttps bool `json:"enforceHttps,omitempty" yaml:"enforceHttps,omitempty"`

	// Name of the TLS security policy that needs to be applied to the HTTPS endpoint. Valid values:  `Policy-Min-TLS-1-0-2019-07` and `Policy-Min-TLS-1-2-2019-07`. The provider will only perform drift detection if a configuration value is provided.
	TlsSecurityPolicy string `json:"tlsSecurityPolicy,omitempty" yaml:"tlsSecurityPolicy,omitempty"`

	// Fully qualified domain for your custom endpoint.
	CustomEndpoint string `json:"customEndpoint,omitempty" yaml:"customEndpoint,omitempty"`

	// ACM certificate ARN for your custom endpoint.
	CustomEndpointCertificateArn string `json:"customEndpointCertificateArn,omitempty" yaml:"customEndpointCertificateArn,omitempty"`
}
