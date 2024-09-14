package types

type Appmesh_getRouteSpecHttpRouteMatchQueryParameter struct {
	//
	Matches []Appmesh_getRouteSpecHttpRouteMatchQueryParameterMatch `json:"matches,omitempty" yaml:"matches,omitempty"`

	// Name of the route.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
