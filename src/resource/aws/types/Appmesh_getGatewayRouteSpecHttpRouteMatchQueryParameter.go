package types

type Appmesh_getGatewayRouteSpecHttpRouteMatchQueryParameter struct {
	// Name of the gateway route.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	//
	Matches []Appmesh_getGatewayRouteSpecHttpRouteMatchQueryParameterMatch `json:"matches,omitempty" yaml:"matches,omitempty"`
}
