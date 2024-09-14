package types

type Appmesh_VirtualNodeSpecBackendVirtualServiceClientPolicyTlsValidationSubjectAlternativeNames struct {
	// Criteria for determining a SAN's match.
	Match Appmesh_VirtualNodeSpecBackendVirtualServiceClientPolicyTlsValidationSubjectAlternativeNamesMatch `json:"match,omitempty" yaml:"match,omitempty"`
}
