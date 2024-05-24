package types

type Appmesh_RouteSpecGrpcRoute struct {
	// Criteria for determining an gRPC request match.
	Match Appmesh_RouteSpecGrpcRouteMatch `json:"match,omitempty" yaml:"match,omitempty"`

	// Retry policy.
	RetryPolicy Appmesh_RouteSpecGrpcRouteRetryPolicy `json:"retryPolicy,omitempty" yaml:"retryPolicy,omitempty"`

	// Types of timeouts.
	Timeout Appmesh_RouteSpecGrpcRouteTimeout `json:"timeout,omitempty" yaml:"timeout,omitempty"`

	// Action to take if a match is determined.
	Action Appmesh_RouteSpecGrpcRouteAction `json:"action,omitempty" yaml:"action,omitempty"`
}
