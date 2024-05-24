package types

type Appmesh_VirtualNodeSpecListenerTlsValidation struct {
	// SANs for a TLS validation context.
	SubjectAlternativeNames Appmesh_VirtualNodeSpecListenerTlsValidationSubjectAlternativeNames `json:"subjectAlternativeNames,omitempty" yaml:"subjectAlternativeNames,omitempty"`

	// TLS validation context trust.
	Trust Appmesh_VirtualNodeSpecListenerTlsValidationTrust `json:"trust,omitempty" yaml:"trust,omitempty"`
}
