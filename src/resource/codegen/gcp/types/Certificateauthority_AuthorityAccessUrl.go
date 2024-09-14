package types

type Certificateauthority_AuthorityAccessUrl struct {
	/*
	   (Output)
	   The URL where this CertificateAuthority's CA certificate is published. This will only be
	   set for CAs that have been activated.
	*/
	CaCertificateAccessUrl string `json:"caCertificateAccessUrl,omitempty" yaml:"caCertificateAccessUrl,omitempty"`

	/*
	   (Output)
	   The URL where this CertificateAuthority's CRLs are published. This will only be set for
	   CAs that have been activated.
	*/
	CrlAccessUrls []string `json:"crlAccessUrls,omitempty" yaml:"crlAccessUrls,omitempty"`
}
