package types

type Appmesh_getVirtualGatewaySpecBackendDefaultClientPolicyTlValidation struct {
	//
	SubjectAlternativeNames []Appmesh_getVirtualGatewaySpecBackendDefaultClientPolicyTlValidationSubjectAlternativeName `json:"subjectAlternativeNames,omitempty" yaml:"subjectAlternativeNames,omitempty"`

	//
	Trusts []Appmesh_getVirtualGatewaySpecBackendDefaultClientPolicyTlValidationTrust `json:"trusts,omitempty" yaml:"trusts,omitempty"`
}
