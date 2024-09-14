package types

type Acmpca_CertificateAuthorityRevocationConfiguration struct {
	// Nested argument containing configuration of the certificate revocation list (CRL), if any, maintained by the certificate authority. Defined below.
	CrlConfiguration Acmpca_CertificateAuthorityRevocationConfigurationCrlConfiguration `json:"crlConfiguration,omitempty" yaml:"crlConfiguration,omitempty"`

	/*
	   Nested argument containing configuration of
	   the custom OCSP responder endpoint. Defined below.
	*/
	OcspConfiguration Acmpca_CertificateAuthorityRevocationConfigurationOcspConfiguration `json:"ocspConfiguration,omitempty" yaml:"ocspConfiguration,omitempty"`
}
