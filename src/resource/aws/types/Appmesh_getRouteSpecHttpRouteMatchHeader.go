package types

type Appmesh_getRouteSpecHttpRouteMatchHeader struct {
	//
	Invert bool `json:"invert,omitempty" yaml:"invert,omitempty"`

	//
	Matches []Appmesh_getRouteSpecHttpRouteMatchHeaderMatch `json:"matches,omitempty" yaml:"matches,omitempty"`

	// Name of the route.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
