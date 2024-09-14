package types

type Appmesh_VirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustSds struct {
	// Name of the secret for a virtual gateway's Transport Layer Security (TLS) Secret Discovery Service validation context trust.
	SecretName string `json:"secretName,omitempty" yaml:"secretName,omitempty"`
}
