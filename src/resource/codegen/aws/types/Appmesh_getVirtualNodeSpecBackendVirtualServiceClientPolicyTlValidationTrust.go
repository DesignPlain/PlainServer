package types

type Appmesh_getVirtualNodeSpecBackendVirtualServiceClientPolicyTlValidationTrust struct {
	//
	Acms []Appmesh_getVirtualNodeSpecBackendVirtualServiceClientPolicyTlValidationTrustAcm `json:"acms,omitempty" yaml:"acms,omitempty"`

	//
	Files []Appmesh_getVirtualNodeSpecBackendVirtualServiceClientPolicyTlValidationTrustFile `json:"files,omitempty" yaml:"files,omitempty"`

	//
	Sds []Appmesh_getVirtualNodeSpecBackendVirtualServiceClientPolicyTlValidationTrustSd `json:"sds,omitempty" yaml:"sds,omitempty"`
}
