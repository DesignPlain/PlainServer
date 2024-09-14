package types

type Appmesh_VirtualGatewaySpecBackendDefaultsClientPolicy struct {
	// Transport Layer Security (TLS) client policy.
	Tls Appmesh_VirtualGatewaySpecBackendDefaultsClientPolicyTls `json:"tls,omitempty" yaml:"tls,omitempty"`
}
