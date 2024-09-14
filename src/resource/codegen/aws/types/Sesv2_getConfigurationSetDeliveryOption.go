package types

type Sesv2_getConfigurationSetDeliveryOption struct {
	// The name of the dedicated IP pool to associate with the configuration set.
	SendingPoolName string `json:"sendingPoolName,omitempty" yaml:"sendingPoolName,omitempty"`

	// Specifies whether messages that use the configuration set are required to use Transport Layer Security (TLS).
	TlsPolicy string `json:"tlsPolicy,omitempty" yaml:"tlsPolicy,omitempty"`
}
