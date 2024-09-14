package types

type Certificateauthority_getAuthorityConfigX509Config struct {
	// Describes the X.509 name constraints extension.
	NameConstraints []Certificateauthority_getAuthorityConfigX509ConfigNameConstraint `json:"nameConstraints,omitempty" yaml:"nameConstraints,omitempty"`

	// Describes the X.509 certificate policy object identifiers, per https://tools.ietf.org/html/rfc5280#section-4.2.1.4.
	PolicyIds []Certificateauthority_getAuthorityConfigX509ConfigPolicyId `json:"policyIds,omitempty" yaml:"policyIds,omitempty"`

	// Specifies an X.509 extension, which may be used in different parts of X.509 objects like certificates, CSRs, and CRLs.
	AdditionalExtensions []Certificateauthority_getAuthorityConfigX509ConfigAdditionalExtension `json:"additionalExtensions,omitempty" yaml:"additionalExtensions,omitempty"`

	/*
	   Describes Online Certificate Status Protocol (OCSP) endpoint addresses that appear in the
	   "Authority Information Access" extension in the certificate.
	*/
	AiaOcspServers []string `json:"aiaOcspServers,omitempty" yaml:"aiaOcspServers,omitempty"`

	// Describes values that are relevant in a CA certificate.
	CaOptions []Certificateauthority_getAuthorityConfigX509ConfigCaOption `json:"caOptions,omitempty" yaml:"caOptions,omitempty"`

	// Indicates the intended use for keys that correspond to a certificate.
	KeyUsages []Certificateauthority_getAuthorityConfigX509ConfigKeyUsage `json:"keyUsages,omitempty" yaml:"keyUsages,omitempty"`
}
