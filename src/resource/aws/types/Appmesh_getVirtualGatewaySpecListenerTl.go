package types

type Appmesh_getVirtualGatewaySpecListenerTl struct {
	//
	Certificates []Appmesh_getVirtualGatewaySpecListenerTlCertificate `json:"certificates,omitempty" yaml:"certificates,omitempty"`

	//
	Mode string `json:"mode,omitempty" yaml:"mode,omitempty"`

	//
	Validations []Appmesh_getVirtualGatewaySpecListenerTlValidation `json:"validations,omitempty" yaml:"validations,omitempty"`
}
