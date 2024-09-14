package types

type Appmesh_VirtualNodeSpecBackendDefaultsClientPolicyTlsValidation struct {
	// SANs for a TLS validation context.
	SubjectAlternativeNames Appmesh_VirtualNodeSpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNames `json:"subjectAlternativeNames,omitempty" yaml:"subjectAlternativeNames,omitempty"`

	// TLS validation context trust.
	Trust Appmesh_VirtualNodeSpecBackendDefaultsClientPolicyTlsValidationTrust `json:"trust,omitempty" yaml:"trust,omitempty"`
}
