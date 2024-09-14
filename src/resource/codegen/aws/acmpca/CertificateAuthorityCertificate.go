package acmpca

type CertificateAuthorityCertificate struct {
	// PEM-encoded certificate for the Certificate Authority.
	Certificate string `json:"certificate,omitempty" yaml:"certificate,omitempty"`

	// ARN of the Certificate Authority.
	CertificateAuthorityArn string `json:"certificateAuthorityArn,omitempty" yaml:"certificateAuthorityArn,omitempty"`

	// PEM-encoded certificate chain that includes any intermediate certificates and chains up to root CA. Required for subordinate Certificate Authorities. Not allowed for root Certificate Authorities.
	CertificateChain string `json:"certificateChain,omitempty" yaml:"certificateChain,omitempty"`
}
