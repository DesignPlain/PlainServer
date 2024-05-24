package types

type Appmesh_VirtualNodeSpecBackendDefaultsClientPolicyTlsValidationTrustAcm struct {
	// One or more ACM ARNs.
	CertificateAuthorityArns []string `json:"certificateAuthorityArns,omitempty" yaml:"certificateAuthorityArns,omitempty"`
}
