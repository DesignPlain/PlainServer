package types

type Appmesh_getVirtualGatewaySpecBackendDefaultClientPolicyTlValidationTrust struct {
	//
	Files []Appmesh_getVirtualGatewaySpecBackendDefaultClientPolicyTlValidationTrustFile `json:"files,omitempty" yaml:"files,omitempty"`

	//
	Sds []Appmesh_getVirtualGatewaySpecBackendDefaultClientPolicyTlValidationTrustSd `json:"sds,omitempty" yaml:"sds,omitempty"`

	//
	Acms []Appmesh_getVirtualGatewaySpecBackendDefaultClientPolicyTlValidationTrustAcm `json:"acms,omitempty" yaml:"acms,omitempty"`
}
