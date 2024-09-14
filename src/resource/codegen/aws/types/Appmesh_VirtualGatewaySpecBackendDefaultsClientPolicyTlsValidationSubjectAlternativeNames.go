package types

type Appmesh_VirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNames struct {
	// Criteria for determining a SAN's match.
	Match Appmesh_VirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatch `json:"match,omitempty" yaml:"match,omitempty"`
}
