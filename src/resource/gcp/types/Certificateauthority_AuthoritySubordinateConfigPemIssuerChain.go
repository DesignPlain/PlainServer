package types

type Certificateauthority_AuthoritySubordinateConfigPemIssuerChain struct {
	// Expected to be in leaf-to-root order according to RFC 5246.
	PemCertificates []string `json:"pemCertificates,omitempty" yaml:"pemCertificates,omitempty"`
}
