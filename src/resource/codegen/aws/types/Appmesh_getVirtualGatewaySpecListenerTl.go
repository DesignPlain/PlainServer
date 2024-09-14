package types

type Appmesh_getVirtualGatewaySpecListenerTl struct {
	//
	Validations []Appmesh_getVirtualGatewaySpecListenerTlValidation `json:"validations,omitempty" yaml:"validations,omitempty"`

	//
	Certificates []Appmesh_getVirtualGatewaySpecListenerTlCertificate `json:"certificates,omitempty" yaml:"certificates,omitempty"`

	//
	Mode string `json:"mode,omitempty" yaml:"mode,omitempty"`
}
