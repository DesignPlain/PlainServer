package types

type Appmesh_getGatewayRouteSpecHttpRouteMatchQueryParameter struct {
	//
	Matches []Appmesh_getGatewayRouteSpecHttpRouteMatchQueryParameterMatch `json:"matches,omitempty" yaml:"matches,omitempty"`

	// Name of the gateway route.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
