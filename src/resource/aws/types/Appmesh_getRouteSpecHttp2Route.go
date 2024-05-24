package types

type Appmesh_getRouteSpecHttp2Route struct {
	//
	Actions []Appmesh_getRouteSpecHttp2RouteAction `json:"actions,omitempty" yaml:"actions,omitempty"`

	//
	Matches []Appmesh_getRouteSpecHttp2RouteMatch `json:"matches,omitempty" yaml:"matches,omitempty"`

	//
	RetryPolicies []Appmesh_getRouteSpecHttp2RouteRetryPolicy `json:"retryPolicies,omitempty" yaml:"retryPolicies,omitempty"`

	//
	Timeouts []Appmesh_getRouteSpecHttp2RouteTimeout `json:"timeouts,omitempty" yaml:"timeouts,omitempty"`
}
