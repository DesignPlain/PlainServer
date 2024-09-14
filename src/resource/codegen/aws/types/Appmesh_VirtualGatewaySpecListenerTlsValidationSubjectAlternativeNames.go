package types

type Appmesh_VirtualGatewaySpecListenerTlsValidationSubjectAlternativeNames struct {
	// Criteria for determining a SAN's match.
	Match Appmesh_VirtualGatewaySpecListenerTlsValidationSubjectAlternativeNamesMatch `json:"match,omitempty" yaml:"match,omitempty"`
}
