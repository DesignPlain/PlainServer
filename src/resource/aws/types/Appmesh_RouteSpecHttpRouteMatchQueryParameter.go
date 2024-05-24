package types

type Appmesh_RouteSpecHttpRouteMatchQueryParameter struct {
	// Criteria for determining an gRPC request match.
	Match Appmesh_RouteSpecHttpRouteMatchQueryParameterMatch `json:"match,omitempty" yaml:"match,omitempty"`

	// Name to use for the route. Must be between 1 and 255 characters in length.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
