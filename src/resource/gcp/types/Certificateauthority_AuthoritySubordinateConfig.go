package types

type Certificateauthority_AuthoritySubordinateConfig struct {
	/*
	   This can refer to a CertificateAuthority that was used to create a
	   subordinate CertificateAuthority. This field is used for information
	   and usability purposes only. The resource name is in the format
	   `projects/-/locations/-/caPools/-/certificateAuthorities/-`.
	*/
	CertificateAuthority string `json:"certificateAuthority,omitempty" yaml:"certificateAuthority,omitempty"`

	/*
	   Contains the PEM certificate chain for the issuers of this CertificateAuthority,
	   but not pem certificate for this CA itself.
	   Structure is documented below.
	*/
	PemIssuerChain Certificateauthority_AuthoritySubordinateConfigPemIssuerChain `json:"pemIssuerChain,omitempty" yaml:"pemIssuerChain,omitempty"`
}
