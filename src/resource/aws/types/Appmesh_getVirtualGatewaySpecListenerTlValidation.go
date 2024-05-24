package types

type Appmesh_getVirtualGatewaySpecListenerTlValidation struct {
	//
	SubjectAlternativeNames []Appmesh_getVirtualGatewaySpecListenerTlValidationSubjectAlternativeName `json:"subjectAlternativeNames,omitempty" yaml:"subjectAlternativeNames,omitempty"`

	//
	Trusts []Appmesh_getVirtualGatewaySpecListenerTlValidationTrust `json:"trusts,omitempty" yaml:"trusts,omitempty"`
}
