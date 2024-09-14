package types

type Certificateauthority_CertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage struct {
	// Corresponds to OID 1.3.6.1.5.5.7.3.3. Officially described as "Signing of downloadable executable code client authentication".
	CodeSigning bool `json:"codeSigning,omitempty" yaml:"codeSigning,omitempty"`

	// Corresponds to OID 1.3.6.1.5.5.7.3.4. Officially described as "Email protection".
	EmailProtection bool `json:"emailProtection,omitempty" yaml:"emailProtection,omitempty"`

	// Corresponds to OID 1.3.6.1.5.5.7.3.9. Officially described as "Signing OCSP responses".
	OcspSigning bool `json:"ocspSigning,omitempty" yaml:"ocspSigning,omitempty"`

	// Corresponds to OID 1.3.6.1.5.5.7.3.1. Officially described as "TLS WWW server authentication", though regularly used for non-WWW TLS.
	ServerAuth bool `json:"serverAuth,omitempty" yaml:"serverAuth,omitempty"`

	// Corresponds to OID 1.3.6.1.5.5.7.3.8. Officially described as "Binding the hash of an object to a time".
	TimeStamping bool `json:"timeStamping,omitempty" yaml:"timeStamping,omitempty"`

	// Corresponds to OID 1.3.6.1.5.5.7.3.2. Officially described as "TLS WWW client authentication", though regularly used for non-WWW TLS.
	ClientAuth bool `json:"clientAuth,omitempty" yaml:"clientAuth,omitempty"`
}
