package types

type Appmesh_GatewayRouteSpecHttp2RouteMatchQueryParameter struct {
	// The query parameter to match on.
	Match Appmesh_GatewayRouteSpecHttp2RouteMatchQueryParameterMatch `json:"match,omitempty" yaml:"match,omitempty"`

	// Name for the query parameter that will be matched on.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
