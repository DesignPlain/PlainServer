package types

type Dns_RecordSetRoutingPolicyWrr struct {
	/*
	   The list of targets to be health checked. Note that if DNSSEC is enabled for this zone, only one of `rrdatas` or `health_checked_targets` can be set.
	   Structure is document below.
	*/
	HealthCheckedTargets Dns_RecordSetRoutingPolicyWrrHealthCheckedTargets `json:"healthCheckedTargets,omitempty" yaml:"healthCheckedTargets,omitempty"`

	// Same as `rrdatas` above.
	Rrdatas []string `json:"rrdatas,omitempty" yaml:"rrdatas,omitempty"`

	// The ratio of traffic routed to the target.
	Weight float64 `json:"weight,omitempty" yaml:"weight,omitempty"`
}
