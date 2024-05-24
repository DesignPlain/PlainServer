package types

type Appmesh_VirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustAcm struct {
	// One or more ACM ARNs.
	CertificateAuthorityArns []string `json:"certificateAuthorityArns,omitempty" yaml:"certificateAuthorityArns,omitempty"`
}
