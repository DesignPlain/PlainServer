package types

type Certificateauthority_CertificateConfigX509Config struct {
	/*
	   Describes values that are relevant in a CA certificate.
	   Structure is documented below.
	*/
	CaOptions Certificateauthority_CertificateConfigX509ConfigCaOptions `json:"caOptions,omitempty" yaml:"caOptions,omitempty"`

	/*
	   Indicates the intended use for keys that correspond to a certificate.
	   Structure is documented below.
	*/
	KeyUsage Certificateauthority_CertificateConfigX509ConfigKeyUsage `json:"keyUsage,omitempty" yaml:"keyUsage,omitempty"`

	/*
	   Describes the X.509 name constraints extension.
	   Structure is documented below.
	*/
	NameConstraints Certificateauthority_CertificateConfigX509ConfigNameConstraints `json:"nameConstraints,omitempty" yaml:"nameConstraints,omitempty"`

	/*
	   Describes the X.509 certificate policy object identifiers, per https://tools.ietf.org/html/rfc5280#section-4.2.1.4.
	   Structure is documented below.
	*/
	PolicyIds []Certificateauthority_CertificateConfigX509ConfigPolicyId `json:"policyIds,omitempty" yaml:"policyIds,omitempty"`

	/*
	   Specifies an X.509 extension, which may be used in different parts of X.509 objects like certificates, CSRs, and CRLs.
	   Structure is documented below.
	*/
	AdditionalExtensions []Certificateauthority_CertificateConfigX509ConfigAdditionalExtension `json:"additionalExtensions,omitempty" yaml:"additionalExtensions,omitempty"`

	/*
	   Describes Online Certificate Status Protocol (OCSP) endpoint addresses that appear in the
	   "Authority Information Access" extension in the certificate.
	*/
	AiaOcspServers []string `json:"aiaOcspServers,omitempty" yaml:"aiaOcspServers,omitempty"`
}
