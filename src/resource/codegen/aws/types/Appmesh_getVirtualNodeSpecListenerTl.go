package types

type Appmesh_getVirtualNodeSpecListenerTl struct {
	//
	Certificates []Appmesh_getVirtualNodeSpecListenerTlCertificate `json:"certificates,omitempty" yaml:"certificates,omitempty"`

	//
	Mode string `json:"mode,omitempty" yaml:"mode,omitempty"`

	//
	Validations []Appmesh_getVirtualNodeSpecListenerTlValidation `json:"validations,omitempty" yaml:"validations,omitempty"`
}
