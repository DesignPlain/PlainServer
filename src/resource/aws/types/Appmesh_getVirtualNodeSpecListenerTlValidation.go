package types

type Appmesh_getVirtualNodeSpecListenerTlValidation struct {
	//
	SubjectAlternativeNames []Appmesh_getVirtualNodeSpecListenerTlValidationSubjectAlternativeName `json:"subjectAlternativeNames,omitempty" yaml:"subjectAlternativeNames,omitempty"`

	//
	Trusts []Appmesh_getVirtualNodeSpecListenerTlValidationTrust `json:"trusts,omitempty" yaml:"trusts,omitempty"`
}
