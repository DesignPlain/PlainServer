package types

type Appmesh_getRouteSpecHttp2RouteMatchHeader struct {
	//
	Invert bool `json:"invert,omitempty" yaml:"invert,omitempty"`

	//
	Matches []Appmesh_getRouteSpecHttp2RouteMatchHeaderMatch `json:"matches,omitempty" yaml:"matches,omitempty"`

	// Name of the route.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
