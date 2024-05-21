package types

type Dns_RecordSetRoutingPolicyGeoHealthCheckedTargets struct {
	/*
	   The list of internal load balancers to health check.
	   Structure is document below.
	*/
	InternalLoadBalancers []Dns_RecordSetRoutingPolicyGeoHealthCheckedTargetsInternalLoadBalancer `json:"internalLoadBalancers,omitempty" yaml:"internalLoadBalancers,omitempty"`
}
