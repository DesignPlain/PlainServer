package types

type Appmesh_getVirtualGatewaySpecBackendDefaultClientPolicyTl struct {
	//
	Certificates []Appmesh_getVirtualGatewaySpecBackendDefaultClientPolicyTlCertificate `json:"certificates,omitempty" yaml:"certificates,omitempty"`

	//
	Enforce bool `json:"enforce,omitempty" yaml:"enforce,omitempty"`

	//
	Ports []int `json:"ports,omitempty" yaml:"ports,omitempty"`

	//
	Validations []Appmesh_getVirtualGatewaySpecBackendDefaultClientPolicyTlValidation `json:"validations,omitempty" yaml:"validations,omitempty"`
}
