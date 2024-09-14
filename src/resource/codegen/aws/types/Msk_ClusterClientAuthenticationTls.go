package types

type Msk_ClusterClientAuthenticationTls struct {
	// List of ACM Certificate Authority Amazon Resource Names (ARNs).
	CertificateAuthorityArns []string `json:"certificateAuthorityArns,omitempty" yaml:"certificateAuthorityArns,omitempty"`
}
