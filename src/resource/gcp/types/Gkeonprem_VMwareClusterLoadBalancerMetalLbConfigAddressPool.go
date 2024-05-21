package types

type Gkeonprem_VMwareClusterLoadBalancerMetalLbConfigAddressPool struct {
	/*
	   The addresses that are part of this pool. Each address
	   must be either in the CIDR form (1.2.3.0/24) or range
	   form (1.2.3.1-1.2.3.5).
	*/
	Addresses []string `json:"addresses,omitempty" yaml:"addresses,omitempty"`

	/*
	   If true, avoid using IPs ending in .0 or .255.
	   This avoids buggy consumer devices mistakenly dropping IPv4 traffic for
	   those special IP addresses.
	*/
	AvoidBuggyIps bool `json:"avoidBuggyIps,omitempty" yaml:"avoidBuggyIps,omitempty"`

	/*
	   If true, prevent IP addresses from being automatically assigned.

	   <a name="nested_dataplane_v2"></a>The `dataplane_v2` block supports:
	*/
	ManualAssign bool `json:"manualAssign,omitempty" yaml:"manualAssign,omitempty"`

	// The name of the address pool.
	Pool string `json:"pool,omitempty" yaml:"pool,omitempty"`
}
