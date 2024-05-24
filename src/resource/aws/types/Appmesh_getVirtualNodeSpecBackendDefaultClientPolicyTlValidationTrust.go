package types

type Appmesh_getVirtualNodeSpecBackendDefaultClientPolicyTlValidationTrust struct {
	//
	Acms []Appmesh_getVirtualNodeSpecBackendDefaultClientPolicyTlValidationTrustAcm `json:"acms,omitempty" yaml:"acms,omitempty"`

	//
	Files []Appmesh_getVirtualNodeSpecBackendDefaultClientPolicyTlValidationTrustFile `json:"files,omitempty" yaml:"files,omitempty"`

	//
	Sds []Appmesh_getVirtualNodeSpecBackendDefaultClientPolicyTlValidationTrustSd `json:"sds,omitempty" yaml:"sds,omitempty"`
}
