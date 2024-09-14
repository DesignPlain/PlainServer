package types

type Appmesh_VirtualNodeSpecBackendVirtualServiceClientPolicyTls struct {
	// One or more ports that the policy is enforced for.
	Ports []int `json:"ports,omitempty" yaml:"ports,omitempty"`

	// Listener's Transport Layer Security (TLS) validation context.
	Validation Appmesh_VirtualNodeSpecBackendVirtualServiceClientPolicyTlsValidation `json:"validation,omitempty" yaml:"validation,omitempty"`

	// Listener's TLS certificate.
	Certificate Appmesh_VirtualNodeSpecBackendVirtualServiceClientPolicyTlsCertificate `json:"certificate,omitempty" yaml:"certificate,omitempty"`

	// Whether the policy is enforced. Default is `true`.
	Enforce bool `json:"enforce,omitempty" yaml:"enforce,omitempty"`
}
