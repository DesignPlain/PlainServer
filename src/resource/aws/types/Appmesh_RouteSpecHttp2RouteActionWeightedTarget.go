package types

type Appmesh_RouteSpecHttp2RouteActionWeightedTarget struct {
	// The targeted port of the weighted object.
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	// Virtual node to associate with the weighted target. Must be between 1 and 255 characters in length.
	VirtualNode string `json:"virtualNode,omitempty" yaml:"virtualNode,omitempty"`

	// Relative weight of the weighted target. An integer between 0 and 100.
	Weight int `json:"weight,omitempty" yaml:"weight,omitempty"`
}
