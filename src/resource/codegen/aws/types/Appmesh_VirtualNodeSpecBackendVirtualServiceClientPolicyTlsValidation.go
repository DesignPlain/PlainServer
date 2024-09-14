package types

type Appmesh_VirtualNodeSpecBackendVirtualServiceClientPolicyTlsValidation struct {
	// SANs for a TLS validation context.
	SubjectAlternativeNames Appmesh_VirtualNodeSpecBackendVirtualServiceClientPolicyTlsValidationSubjectAlternativeNames `json:"subjectAlternativeNames,omitempty" yaml:"subjectAlternativeNames,omitempty"`

	// TLS validation context trust.
	Trust Appmesh_VirtualNodeSpecBackendVirtualServiceClientPolicyTlsValidationTrust `json:"trust,omitempty" yaml:"trust,omitempty"`
}
