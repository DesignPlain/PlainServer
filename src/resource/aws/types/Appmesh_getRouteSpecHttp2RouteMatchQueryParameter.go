package types

type Appmesh_getRouteSpecHttp2RouteMatchQueryParameter struct {
	//
	Matches []Appmesh_getRouteSpecHttp2RouteMatchQueryParameterMatch `json:"matches,omitempty" yaml:"matches,omitempty"`

	// Name of the route.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
