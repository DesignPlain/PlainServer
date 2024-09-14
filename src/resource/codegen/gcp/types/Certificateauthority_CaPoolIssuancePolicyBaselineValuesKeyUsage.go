package types

type Certificateauthority_CaPoolIssuancePolicyBaselineValuesKeyUsage struct {
	/*
	   Describes high-level ways in which a key may be used.
	   Structure is documented below.
	*/
	BaseKeyUsage Certificateauthority_CaPoolIssuancePolicyBaselineValuesKeyUsageBaseKeyUsage `json:"baseKeyUsage,omitempty" yaml:"baseKeyUsage,omitempty"`

	/*
	   Describes high-level ways in which a key may be used.
	   Structure is documented below.
	*/
	ExtendedKeyUsage Certificateauthority_CaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage `json:"extendedKeyUsage,omitempty" yaml:"extendedKeyUsage,omitempty"`

	/*
	   An ObjectId specifies an object identifier (OID). These provide context and describe types in ASN.1 messages.
	   Structure is documented below.
	*/
	UnknownExtendedKeyUsages []Certificateauthority_CaPoolIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsage `json:"unknownExtendedKeyUsages,omitempty" yaml:"unknownExtendedKeyUsages,omitempty"`
}
