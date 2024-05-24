package types

type Acmpca_getCertificateAuthorityRevocationConfigurationOcspConfiguration struct {
	//
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	//
	OcspCustomCname string `json:"ocspCustomCname,omitempty" yaml:"ocspCustomCname,omitempty"`
}
