package types

type Appmesh_GatewayRouteSpecHttp2Route struct {
	// Action to take if a match is determined.
	Action Appmesh_GatewayRouteSpecHttp2RouteAction `json:"action,omitempty" yaml:"action,omitempty"`

	// Criteria for determining a request match.
	Match Appmesh_GatewayRouteSpecHttp2RouteMatch `json:"match,omitempty" yaml:"match,omitempty"`
}
