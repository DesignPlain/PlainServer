package types

type Appmesh_RouteSpecTcpRoute struct {
	// Action to take if a match is determined.
	Action Appmesh_RouteSpecTcpRouteAction `json:"action,omitempty" yaml:"action,omitempty"`

	//
	Match Appmesh_RouteSpecTcpRouteMatch `json:"match,omitempty" yaml:"match,omitempty"`

	// Types of timeouts.
	Timeout Appmesh_RouteSpecTcpRouteTimeout `json:"timeout,omitempty" yaml:"timeout,omitempty"`
}
