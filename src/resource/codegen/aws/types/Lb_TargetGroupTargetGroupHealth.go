package types

type Lb_TargetGroupTargetGroupHealth struct {
	// Block to configure DNS Failover requirements. See DNS Failover below for details on attributes.
	DnsFailover Lb_TargetGroupTargetGroupHealthDnsFailover `json:"dnsFailover,omitempty" yaml:"dnsFailover,omitempty"`

	// Block to configure Unhealthy State Routing requirements. See Unhealthy State Routing below for details on attributes.
	UnhealthyStateRouting Lb_TargetGroupTargetGroupHealthUnhealthyStateRouting `json:"unhealthyStateRouting,omitempty" yaml:"unhealthyStateRouting,omitempty"`
}
