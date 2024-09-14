package types

type Opensearch_DomainDomainEndpointOptions struct {
	// Whether to enable custom endpoint for the OpenSearch domain.
	CustomEndpointEnabled bool `json:"customEndpointEnabled,omitempty" yaml:"customEndpointEnabled,omitempty"`

	// Whether or not to require HTTPS. Defaults to `true`.
	EnforceHttps bool `json:"enforceHttps,omitempty" yaml:"enforceHttps,omitempty"`

	// Name of the TLS security policy that needs to be applied to the HTTPS endpoint. For valid values, refer to the [AWS documentation](https://docs.aws.amazon.com/opensearch-service/latest/APIReference/API_DomainEndpointOptions.html#opensearchservice-Type-DomainEndpointOptions-TLSSecurityPolicy). Pulumi will only perform drift detection if a configuration value is provided.
	TlsSecurityPolicy string `json:"tlsSecurityPolicy,omitempty" yaml:"tlsSecurityPolicy,omitempty"`

	// Fully qualified domain for your custom endpoint.
	CustomEndpoint string `json:"customEndpoint,omitempty" yaml:"customEndpoint,omitempty"`

	// ACM certificate ARN for your custom endpoint.
	CustomEndpointCertificateArn string `json:"customEndpointCertificateArn,omitempty" yaml:"customEndpointCertificateArn,omitempty"`
}
