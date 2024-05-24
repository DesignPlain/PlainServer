package types

type Appmesh_RouteSpecHttpRoute struct {
	// Action to take if a match is determined.
	Action Appmesh_RouteSpecHttpRouteAction `json:"action,omitempty" yaml:"action,omitempty"`

	// Criteria for determining an HTTP request match.
	Match Appmesh_RouteSpecHttpRouteMatch `json:"match,omitempty" yaml:"match,omitempty"`

	// Retry policy.
	RetryPolicy Appmesh_RouteSpecHttpRouteRetryPolicy `json:"retryPolicy,omitempty" yaml:"retryPolicy,omitempty"`

	// Types of timeouts.
	Timeout Appmesh_RouteSpecHttpRouteTimeout `json:"timeout,omitempty" yaml:"timeout,omitempty"`
}
