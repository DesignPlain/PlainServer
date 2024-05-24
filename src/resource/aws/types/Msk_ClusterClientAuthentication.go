package types

type Msk_ClusterClientAuthentication struct {
	// Configuration block for specifying SASL client authentication. See below.
	Sasl Msk_ClusterClientAuthenticationSasl `json:"sasl,omitempty" yaml:"sasl,omitempty"`

	// Configuration block for specifying TLS client authentication. See below.
	Tls Msk_ClusterClientAuthenticationTls `json:"tls,omitempty" yaml:"tls,omitempty"`

	// Enables unauthenticated access.
	Unauthenticated bool `json:"unauthenticated,omitempty" yaml:"unauthenticated,omitempty"`
}
