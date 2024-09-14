package types

type Appmesh_RouteSpecTcpRouteAction struct {
	/*
	   Targets that traffic is routed to when a request matches the route.
	   You can specify one or more targets and their relative weights with which to distribute traffic.
	*/
	WeightedTargets []Appmesh_RouteSpecTcpRouteActionWeightedTarget `json:"weightedTargets,omitempty" yaml:"weightedTargets,omitempty"`
}
