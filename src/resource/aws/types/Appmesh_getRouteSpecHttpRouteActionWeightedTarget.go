package types

type Appmesh_getRouteSpecHttpRouteActionWeightedTarget struct {
	//
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	//
	VirtualNode string `json:"virtualNode,omitempty" yaml:"virtualNode,omitempty"`

	//
	Weight int `json:"weight,omitempty" yaml:"weight,omitempty"`
}
