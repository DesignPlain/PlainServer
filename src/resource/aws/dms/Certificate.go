package dms

type Certificate struct {
	// The contents of the .pem X.509 certificate file for the certificate. Either `certificate_pem` or `certificate_wallet` must be set.
	CertificatePem string `json:"certificatePem,omitempty" yaml:"certificatePem,omitempty"`

	// The contents of the Oracle Wallet certificate for use with SSL, provided as a base64-encoded String. Either `certificate_pem` or `certificate_wallet` must be set.
	CertificateWallet string `json:"certificateWallet,omitempty" yaml:"certificateWallet,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	/*
	   The certificate identifier.

	   - Must contain from 1 to 255 alphanumeric characters and hyphens.
	*/
	CertificateId string `json:"certificateId,omitempty" yaml:"certificateId,omitempty"`
}
