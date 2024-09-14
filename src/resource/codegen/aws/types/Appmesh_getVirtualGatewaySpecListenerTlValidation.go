package types

type Appmesh_getVirtualGatewaySpecListenerTlValidation struct {
	//
	Trusts []Appmesh_getVirtualGatewaySpecListenerTlValidationTrust `json:"trusts,omitempty" yaml:"trusts,omitempty"`

	//
	SubjectAlternativeNames []Appmesh_getVirtualGatewaySpecListenerTlValidationSubjectAlternativeName `json:"subjectAlternativeNames,omitempty" yaml:"subjectAlternativeNames,omitempty"`
}
