package types

type Appmesh_VirtualNodeSpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNames struct {
	// Criteria for determining a SAN's match.
	Match Appmesh_VirtualNodeSpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatch `json:"match,omitempty" yaml:"match,omitempty"`
}
