package types

type Appmesh_VirtualGatewaySpecBackendDefaultsClientPolicyTlsValidation struct {
	// SANs for a virtual gateway's listener's Transport Layer Security (TLS) validation context.
	SubjectAlternativeNames Appmesh_VirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNames `json:"subjectAlternativeNames,omitempty" yaml:"subjectAlternativeNames,omitempty"`

	// TLS validation context trust.
	Trust Appmesh_VirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrust `json:"trust,omitempty" yaml:"trust,omitempty"`
}
