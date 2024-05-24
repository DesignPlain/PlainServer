package types

type Appmesh_VirtualGatewaySpecBackendDefaultsClientPolicyTls struct {
	// Listener's TLS certificate.
	Certificate Appmesh_VirtualGatewaySpecBackendDefaultsClientPolicyTlsCertificate `json:"certificate,omitempty" yaml:"certificate,omitempty"`

	// Whether the policy is enforced. Default is `true`.
	Enforce bool `json:"enforce,omitempty" yaml:"enforce,omitempty"`

	// One or more ports that the policy is enforced for.
	Ports []int `json:"ports,omitempty" yaml:"ports,omitempty"`

	// Listener's Transport Layer Security (TLS) validation context.
	Validation Appmesh_VirtualGatewaySpecBackendDefaultsClientPolicyTlsValidation `json:"validation,omitempty" yaml:"validation,omitempty"`
}
