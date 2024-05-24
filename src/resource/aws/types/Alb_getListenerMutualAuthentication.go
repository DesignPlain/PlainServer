package types

type Alb_getListenerMutualAuthentication struct {
	//
	Mode string `json:"mode,omitempty" yaml:"mode,omitempty"`

	//
	TrustStoreArn string `json:"trustStoreArn,omitempty" yaml:"trustStoreArn,omitempty"`

	//
	IgnoreClientCertificateExpiry bool `json:"ignoreClientCertificateExpiry,omitempty" yaml:"ignoreClientCertificateExpiry,omitempty"`
}
