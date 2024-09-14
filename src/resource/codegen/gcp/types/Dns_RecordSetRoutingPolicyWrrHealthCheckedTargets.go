package types

type Dns_RecordSetRoutingPolicyWrrHealthCheckedTargets struct {
	/*
	   The list of internal load balancers to health check.
	   Structure is document below.
	*/
	InternalLoadBalancers []Dns_RecordSetRoutingPolicyWrrHealthCheckedTargetsInternalLoadBalancer `json:"internalLoadBalancers,omitempty" yaml:"internalLoadBalancers,omitempty"`
}
