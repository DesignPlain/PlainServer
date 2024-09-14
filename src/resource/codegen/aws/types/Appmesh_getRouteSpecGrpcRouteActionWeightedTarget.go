package types

type Appmesh_getRouteSpecGrpcRouteActionWeightedTarget struct {
	//
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	//
	VirtualNode string `json:"virtualNode,omitempty" yaml:"virtualNode,omitempty"`

	//
	Weight int `json:"weight,omitempty" yaml:"weight,omitempty"`
}
