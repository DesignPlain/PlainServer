package types

type Appmesh_getVirtualNodeSpecBackendDefaultClientPolicyTlValidation struct {
	//
	SubjectAlternativeNames []Appmesh_getVirtualNodeSpecBackendDefaultClientPolicyTlValidationSubjectAlternativeName `json:"subjectAlternativeNames,omitempty" yaml:"subjectAlternativeNames,omitempty"`

	//
	Trusts []Appmesh_getVirtualNodeSpecBackendDefaultClientPolicyTlValidationTrust `json:"trusts,omitempty" yaml:"trusts,omitempty"`
}
