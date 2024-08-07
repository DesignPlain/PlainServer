package types

type Compute_getRouterNatSubnetwork struct {
	/*
	   Name of the NAT service. The name must be 1-63 characters long and
	   comply with RFC1035.
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   List of the secondary ranges of the subnetwork that are allowed
	   to use NAT. This can be populated only if
	   'LIST_OF_SECONDARY_IP_RANGES' is one of the values in
	   sourceIpRangesToNat
	*/
	SecondaryIpRangeNames []string `json:"secondaryIpRangeNames,omitempty" yaml:"secondaryIpRangeNames,omitempty"`

	/*
	   List of options for which source IPs in the subnetwork
	   should have NAT enabled. Supported values include:
	   'ALL_IP_RANGES', 'LIST_OF_SECONDARY_IP_RANGES',
	   'PRIMARY_IP_RANGE'.
	*/
	SourceIpRangesToNats []string `json:"sourceIpRangesToNats,omitempty" yaml:"sourceIpRangesToNats,omitempty"`
}
