package types

type Appmesh_GatewayRouteSpecHttpRouteAction struct {
	// Gateway route action to rewrite.
	Rewrite Appmesh_GatewayRouteSpecHttpRouteActionRewrite `json:"rewrite,omitempty" yaml:"rewrite,omitempty"`

	// Target that traffic is routed to when a request matches the gateway route.
	Target Appmesh_GatewayRouteSpecHttpRouteActionTarget `json:"target,omitempty" yaml:"target,omitempty"`
}
