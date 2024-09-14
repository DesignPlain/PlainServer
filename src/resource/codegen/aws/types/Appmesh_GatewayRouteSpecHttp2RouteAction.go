package types

type Appmesh_GatewayRouteSpecHttp2RouteAction struct {
	// Gateway route action to rewrite.
	Rewrite Appmesh_GatewayRouteSpecHttp2RouteActionRewrite `json:"rewrite,omitempty" yaml:"rewrite,omitempty"`

	// Target that traffic is routed to when a request matches the gateway route.
	Target Appmesh_GatewayRouteSpecHttp2RouteActionTarget `json:"target,omitempty" yaml:"target,omitempty"`
}
