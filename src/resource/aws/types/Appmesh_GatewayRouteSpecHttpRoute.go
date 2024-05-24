package types

type Appmesh_GatewayRouteSpecHttpRoute struct {
	// Action to take if a match is determined.
	Action Appmesh_GatewayRouteSpecHttpRouteAction `json:"action,omitempty" yaml:"action,omitempty"`

	// Criteria for determining a request match.
	Match Appmesh_GatewayRouteSpecHttpRouteMatch `json:"match,omitempty" yaml:"match,omitempty"`
}
