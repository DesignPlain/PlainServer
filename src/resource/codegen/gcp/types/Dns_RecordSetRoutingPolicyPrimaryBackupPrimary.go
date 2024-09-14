package types

type Dns_RecordSetRoutingPolicyPrimaryBackupPrimary struct {
	/*
	   The list of internal load balancers to health check.
	   Structure is document below.
	*/
	InternalLoadBalancers []Dns_RecordSetRoutingPolicyPrimaryBackupPrimaryInternalLoadBalancer `json:"internalLoadBalancers,omitempty" yaml:"internalLoadBalancers,omitempty"`
}
