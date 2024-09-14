package types

type Appmesh_getVirtualNodeSpecBackendDefaultClientPolicyTl struct {
	//
	Certificates []Appmesh_getVirtualNodeSpecBackendDefaultClientPolicyTlCertificate `json:"certificates,omitempty" yaml:"certificates,omitempty"`

	//
	Enforce bool `json:"enforce,omitempty" yaml:"enforce,omitempty"`

	//
	Ports []int `json:"ports,omitempty" yaml:"ports,omitempty"`

	//
	Validations []Appmesh_getVirtualNodeSpecBackendDefaultClientPolicyTlValidation `json:"validations,omitempty" yaml:"validations,omitempty"`
}
