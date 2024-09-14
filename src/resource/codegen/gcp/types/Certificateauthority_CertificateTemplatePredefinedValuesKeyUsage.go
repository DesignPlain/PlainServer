package types

type Certificateauthority_CertificateTemplatePredefinedValuesKeyUsage struct {
	// Used to describe extended key usages that are not listed in the KeyUsage.ExtendedKeyUsageOptions message.
	UnknownExtendedKeyUsages []Certificateauthority_CertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsage `json:"unknownExtendedKeyUsages,omitempty" yaml:"unknownExtendedKeyUsages,omitempty"`

	// Describes high-level ways in which a key may be used.
	BaseKeyUsage Certificateauthority_CertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage `json:"baseKeyUsage,omitempty" yaml:"baseKeyUsage,omitempty"`

	// Detailed scenarios in which a key may be used.
	ExtendedKeyUsage Certificateauthority_CertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage `json:"extendedKeyUsage,omitempty" yaml:"extendedKeyUsage,omitempty"`
}
