package types

type Acmpca_getCertificateAuthorityRevocationConfiguration struct {
	//
	OcspConfigurations []Acmpca_getCertificateAuthorityRevocationConfigurationOcspConfiguration `json:"ocspConfigurations,omitempty" yaml:"ocspConfigurations,omitempty"`

	//
	CrlConfigurations []Acmpca_getCertificateAuthorityRevocationConfigurationCrlConfiguration `json:"crlConfigurations,omitempty" yaml:"crlConfigurations,omitempty"`
}
