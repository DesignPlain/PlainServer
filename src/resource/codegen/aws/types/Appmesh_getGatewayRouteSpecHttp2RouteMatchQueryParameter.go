package types

type Appmesh_getGatewayRouteSpecHttp2RouteMatchQueryParameter struct {
	//
	Matches []Appmesh_getGatewayRouteSpecHttp2RouteMatchQueryParameterMatch `json:"matches,omitempty" yaml:"matches,omitempty"`

	// Name of the gateway route.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
