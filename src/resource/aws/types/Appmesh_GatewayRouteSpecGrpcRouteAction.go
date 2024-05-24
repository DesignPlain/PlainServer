package types

type Appmesh_GatewayRouteSpecGrpcRouteAction struct {
	// Target that traffic is routed to when a request matches the gateway route.
	Target Appmesh_GatewayRouteSpecGrpcRouteActionTarget `json:"target,omitempty" yaml:"target,omitempty"`
}
