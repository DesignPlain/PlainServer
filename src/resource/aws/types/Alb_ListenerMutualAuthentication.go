package types

type Alb_ListenerMutualAuthentication struct {
	// Whether client certificate expiry is ignored. Default is `false`.
	IgnoreClientCertificateExpiry bool `json:"ignoreClientCertificateExpiry,omitempty" yaml:"ignoreClientCertificateExpiry,omitempty"`

	// Valid values are `off`, `verify` and `passthrough`.
	Mode string `json:"mode,omitempty" yaml:"mode,omitempty"`

	// ARN of the elbv2 Trust Store.
	TrustStoreArn string `json:"trustStoreArn,omitempty" yaml:"trustStoreArn,omitempty"`
}
