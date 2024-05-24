package types

type Appmesh_RouteSpecHttp2Route struct {
	// Action to take if a match is determined.
	Action Appmesh_RouteSpecHttp2RouteAction `json:"action,omitempty" yaml:"action,omitempty"`

	// Criteria for determining an gRPC request match.
	Match Appmesh_RouteSpecHttp2RouteMatch `json:"match,omitempty" yaml:"match,omitempty"`

	// Retry policy.
	RetryPolicy Appmesh_RouteSpecHttp2RouteRetryPolicy `json:"retryPolicy,omitempty" yaml:"retryPolicy,omitempty"`

	// Types of timeouts.
	Timeout Appmesh_RouteSpecHttp2RouteTimeout `json:"timeout,omitempty" yaml:"timeout,omitempty"`
}
