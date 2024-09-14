package types

type Appmesh_GatewayRouteSpecGrpcRoute struct {
	// Action to take if a match is determined.
	Action Appmesh_GatewayRouteSpecGrpcRouteAction `json:"action,omitempty" yaml:"action,omitempty"`

	// Criteria for determining a request match.
	Match Appmesh_GatewayRouteSpecGrpcRouteMatch `json:"match,omitempty" yaml:"match,omitempty"`
}
