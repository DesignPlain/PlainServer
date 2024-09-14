package types

type Appmesh_VirtualNodeSpecListenerTlsValidationSubjectAlternativeNames struct {
	// Criteria for determining a SAN's match.
	Match Appmesh_VirtualNodeSpecListenerTlsValidationSubjectAlternativeNamesMatch `json:"match,omitempty" yaml:"match,omitempty"`
}
