package types

type Certificateauthority_CaPoolPublishingOptions struct {
	/*
	   Specifies the encoding format of each CertificateAuthority's CA
	   certificate and CRLs. If this is omitted, CA certificates and CRLs
	   will be published in PEM.
	   Possible values are: `PEM`, `DER`.
	*/
	EncodingFormat string `json:"encodingFormat,omitempty" yaml:"encodingFormat,omitempty"`

	/*
	   When true, publishes each CertificateAuthority's CA certificate and includes its URL in the "Authority Information Access"
	   X.509 extension in all issued Certificates. If this is false, the CA certificate will not be published and the corresponding
	   X.509 extension will not be written in issued certificates.
	*/
	PublishCaCert bool `json:"publishCaCert,omitempty" yaml:"publishCaCert,omitempty"`

	/*
	   When true, publishes each CertificateAuthority's CRL and includes its URL in the "CRL Distribution Points" X.509 extension
	   in all issued Certificates. If this is false, CRLs will not be published and the corresponding X.509 extension will not
	   be written in issued certificates. CRLs will expire 7 days from their creation. However, we will rebuild daily. CRLs are
	   also rebuilt shortly after a certificate is revoked.
	*/
	PublishCrl bool `json:"publishCrl,omitempty" yaml:"publishCrl,omitempty"`
}