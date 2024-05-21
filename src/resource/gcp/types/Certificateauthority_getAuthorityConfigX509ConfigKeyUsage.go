package types

type Certificateauthority_getAuthorityConfigX509ConfigKeyUsage struct {
	// An ObjectId specifies an object identifier (OID). These provide context and describe types in ASN.1 messages.
	UnknownExtendedKeyUsages []Certificateauthority_getAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsage `json:"unknownExtendedKeyUsages,omitempty" yaml:"unknownExtendedKeyUsages,omitempty"`

	// Describes high-level ways in which a key may be used.
	BaseKeyUsages []Certificateauthority_getAuthorityConfigX509ConfigKeyUsageBaseKeyUsage `json:"baseKeyUsages,omitempty" yaml:"baseKeyUsages,omitempty"`

	// Describes high-level ways in which a key may be used.
	ExtendedKeyUsages []Certificateauthority_getAuthorityConfigX509ConfigKeyUsageExtendedKeyUsage `json:"extendedKeyUsages,omitempty" yaml:"extendedKeyUsages,omitempty"`
}
