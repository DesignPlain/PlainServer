package types

type Appmesh_getVirtualNodeSpecBackendVirtualServiceClientPolicyTl struct {
	//
	Validations []Appmesh_getVirtualNodeSpecBackendVirtualServiceClientPolicyTlValidation `json:"validations,omitempty" yaml:"validations,omitempty"`

	//
	Certificates []Appmesh_getVirtualNodeSpecBackendVirtualServiceClientPolicyTlCertificate `json:"certificates,omitempty" yaml:"certificates,omitempty"`

	//
	Enforce bool `json:"enforce,omitempty" yaml:"enforce,omitempty"`

	//
	Ports []int `json:"ports,omitempty" yaml:"ports,omitempty"`
}
