package types

type Acmpca_getCertificateAuthorityRevocationConfiguration struct {
	// Nested attribute containing configuration of the certificate revocation list (CRL), if any, maintained by the certificate authority.
	CrlConfigurations []Acmpca_getCertificateAuthorityRevocationConfigurationCrlConfiguration `json:"crlConfigurations,omitempty" yaml:"crlConfigurations,omitempty"`

	//
	OcspConfigurations []Acmpca_getCertificateAuthorityRevocationConfigurationOcspConfiguration `json:"ocspConfigurations,omitempty" yaml:"ocspConfigurations,omitempty"`
}
