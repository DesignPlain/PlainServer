package types

type Appmesh_getRouteSpecHttpRoute struct {
	//
	Actions []Appmesh_getRouteSpecHttpRouteAction `json:"actions,omitempty" yaml:"actions,omitempty"`

	//
	Matches []Appmesh_getRouteSpecHttpRouteMatch `json:"matches,omitempty" yaml:"matches,omitempty"`

	//
	RetryPolicies []Appmesh_getRouteSpecHttpRouteRetryPolicy `json:"retryPolicies,omitempty" yaml:"retryPolicies,omitempty"`

	//
	Timeouts []Appmesh_getRouteSpecHttpRouteTimeout `json:"timeouts,omitempty" yaml:"timeouts,omitempty"`
}
