package types

type Appmesh_getRouteSpecGrpcRoute struct {
	//
	Actions []Appmesh_getRouteSpecGrpcRouteAction `json:"actions,omitempty" yaml:"actions,omitempty"`

	//
	Matches []Appmesh_getRouteSpecGrpcRouteMatch `json:"matches,omitempty" yaml:"matches,omitempty"`

	//
	RetryPolicies []Appmesh_getRouteSpecGrpcRouteRetryPolicy `json:"retryPolicies,omitempty" yaml:"retryPolicies,omitempty"`

	//
	Timeouts []Appmesh_getRouteSpecGrpcRouteTimeout `json:"timeouts,omitempty" yaml:"timeouts,omitempty"`
}
