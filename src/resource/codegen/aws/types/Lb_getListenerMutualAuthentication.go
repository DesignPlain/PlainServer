package types

type Lb_getListenerMutualAuthentication struct {
	//
	IgnoreClientCertificateExpiry bool `json:"ignoreClientCertificateExpiry,omitempty" yaml:"ignoreClientCertificateExpiry,omitempty"`

	//
	Mode string `json:"mode,omitempty" yaml:"mode,omitempty"`

	//
	TrustStoreArn string `json:"trustStoreArn,omitempty" yaml:"trustStoreArn,omitempty"`
}
