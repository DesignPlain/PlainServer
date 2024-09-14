package types

type Acmpca_CertificateAuthorityRevocationConfigurationOcspConfiguration struct {
	// Boolean value that specifies whether a custom OCSP responder is enabled.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// CNAME specifying a customized OCSP domain. Note: The value of the CNAME must not include a protocol prefix such as "http://" or "https://".
	OcspCustomCname string `json:"ocspCustomCname,omitempty" yaml:"ocspCustomCname,omitempty"`
}
