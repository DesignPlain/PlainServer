package types

type Certificateauthority_CertificateCertificateDescriptionX509DescriptionKeyUsage struct {
	/*
	   Describes high-level ways in which a key may be used.
	   Structure is documented below.
	*/
	BaseKeyUsages []Certificateauthority_CertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage `json:"baseKeyUsages,omitempty" yaml:"baseKeyUsages,omitempty"`

	/*
	   Describes high-level ways in which a key may be used.
	   Structure is documented below.
	*/
	ExtendedKeyUsages []Certificateauthority_CertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsage `json:"extendedKeyUsages,omitempty" yaml:"extendedKeyUsages,omitempty"`

	/*
	   An ObjectId specifies an object identifier (OID). These provide context and describe types in ASN.1 messages.
	   Structure is documented below.
	*/
	UnknownExtendedKeyUsages []Certificateauthority_CertificateCertificateDescriptionX509DescriptionKeyUsageUnknownExtendedKeyUsage `json:"unknownExtendedKeyUsages,omitempty" yaml:"unknownExtendedKeyUsages,omitempty"`
}
