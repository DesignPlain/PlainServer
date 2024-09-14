package types

type Cloudsearch_DomainEndpointOptions struct {
	// Enables or disables the requirement that all requests to the domain arrive over HTTPS.
	EnforceHttps bool `json:"enforceHttps,omitempty" yaml:"enforceHttps,omitempty"`

	// The minimum required TLS version. See the [AWS documentation](https://docs.aws.amazon.com/cloudsearch/latest/developerguide/API_DomainEndpointOptions.html) for valid values.
	TlsSecurityPolicy string `json:"tlsSecurityPolicy,omitempty" yaml:"tlsSecurityPolicy,omitempty"`
}
