package types

type Certificateauthority_CertificateTemplatePredefinedValues struct {
	// Optional. Describes Online Certificate Status Protocol (OCSP) endpoint addresses that appear in the "Authority Information Access" extension in the certificate.
	AiaOcspServers []string `json:"aiaOcspServers,omitempty" yaml:"aiaOcspServers,omitempty"`

	// Optional. Describes options in this X509Parameters that are relevant in a CA certificate.
	CaOptions Certificateauthority_CertificateTemplatePredefinedValuesCaOptions `json:"caOptions,omitempty" yaml:"caOptions,omitempty"`

	// Optional. Indicates the intended use for keys that correspond to a certificate.
	KeyUsage Certificateauthority_CertificateTemplatePredefinedValuesKeyUsage `json:"keyUsage,omitempty" yaml:"keyUsage,omitempty"`

	// Optional. Describes the X.509 certificate policy object identifiers, per https://tools.ietf.org/html/rfc5280#section-4.2.1.4.
	PolicyIds []Certificateauthority_CertificateTemplatePredefinedValuesPolicyId `json:"policyIds,omitempty" yaml:"policyIds,omitempty"`

	// Optional. Describes custom X.509 extensions.
	AdditionalExtensions []Certificateauthority_CertificateTemplatePredefinedValuesAdditionalExtension `json:"additionalExtensions,omitempty" yaml:"additionalExtensions,omitempty"`
}
