package types

type Dns_RecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargets struct {
	/*
	   The list of internal load balancers to health check.
	   Structure is document below.
	*/
	InternalLoadBalancers []Dns_RecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargetsInternalLoadBalancer `json:"internalLoadBalancers,omitempty" yaml:"internalLoadBalancers,omitempty"`
}
