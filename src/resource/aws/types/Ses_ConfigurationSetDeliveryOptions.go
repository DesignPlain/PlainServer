package types

type Ses_ConfigurationSetDeliveryOptions struct {
	// Whether messages that use the configuration set are required to use Transport Layer Security (TLS). If the value is `Require`, messages are only delivered if a TLS connection can be established. If the value is `Optional`, messages can be delivered in plain text if a TLS connection can't be established. Valid values: `Require` or `Optional`. Defaults to `Optional`.
	TlsPolicy string `json:"tlsPolicy,omitempty" yaml:"tlsPolicy,omitempty"`
}
