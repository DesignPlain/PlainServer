package types

type Acmpca_getCertificateAuthorityRevocationConfigurationOcspConfiguration struct {
	// Boolean value that specifies whether a custom OCSP responder is enabled.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// A CNAME specifying a customized OCSP domain.
	OcspCustomCname string `json:"ocspCustomCname,omitempty" yaml:"ocspCustomCname,omitempty"`
}
